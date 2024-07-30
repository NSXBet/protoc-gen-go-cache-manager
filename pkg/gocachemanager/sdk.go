package gocachemanager

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	redis_store "github.com/eko/gocache/store/redis/v4"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
	redis "github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

var ErrCacheMiss = errors.New("cache miss")

type GoCacheWrapper struct {
	prefix       string
	expiration   time.Duration
	cacheManager *cache.ChainCache[[]byte]
}

func NewGoCacheWrapper(
	prefix string,
	expiration time.Duration,
	settings *CacheSettings,
) (*GoCacheWrapper, error) {
	var redisStore *redis_store.RedisStore

	if settings.redisConnection != "" {
		redisClient := redis.NewClient(&redis.Options{Addr: settings.redisConnection})

		// Initialize stores
		redisStore = redis_store.NewRedis(redisClient, store.WithExpiration(5*time.Second))
	}

	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000,
		MaxCost:     100,
		BufferItems: 64,
	})
	if err != nil {
		return nil, fmt.Errorf("creating ristretto instance: %w", err)
	}

	ristrettoStore := ristretto_store.NewRistretto(ristrettoCache)

	caches := []cache.SetterCacheInterface[[]byte]{
		cache.New[[]byte](ristrettoStore),
	}

	if redisStore != nil {
		caches = append(caches, cache.New[[]byte](redisStore))
	}

	// Initialize chained cache
	cacheManager := cache.NewChain[[]byte](caches...)

	return &GoCacheWrapper{
		prefix:       prefix,
		cacheManager: cacheManager,
		expiration:   expiration,
	}, nil
}

func (gcw *GoCacheWrapper) getKey(key []byte) string {
	b64 := base64.StdEncoding.EncodeToString(key)

	return fmt.Sprintf("%s:%s", gcw.prefix, b64)
}

func (gcw *GoCacheWrapper) Get(ctx context.Context, key []byte) ([]byte, error) {
	strKey := gcw.getKey(key)
	data, err := gcw.cacheManager.Get(ctx, strKey)
	if err != nil {
		if _, isError := lo.ErrorsAs[*store.NotFound](err); isError {
			return nil, ErrCacheMiss
		}

		return nil, fmt.Errorf("getting data: %w", err)
	}

	return data, nil
}

func (gcw *GoCacheWrapper) Set(ctx context.Context, key []byte, value []byte) error {
	strKey := gcw.getKey(key)

	return gcw.cacheManager.Set(ctx, strKey, value, store.WithExpiration(gcw.expiration))
}
