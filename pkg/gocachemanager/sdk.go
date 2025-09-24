package gocachemanager

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/metrics"
	"github.com/eko/gocache/lib/v4/store"
	redis_store "github.com/eko/gocache/store/redis/v4"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
	redis "github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

const (
	pingRedisConnectionInterval   = 5 * time.Second
	pingRedisConnectionMaxRetries = 3
)

var ErrCacheMiss = errors.New("cache miss")

type GoCacheWrapper struct {
	prefix       string
	expiration   time.Duration
	cacheManager cache.CacheInterface[string]
	gzip         bool
}

func NewGoCacheWrapper(
	prefix string,
	expiration time.Duration,
	settings *CacheSettings,
) (*GoCacheWrapper, error) {
	caches := []cache.SetterCacheInterface[string]{}

	if !settings.skipInMemoryCache {
		maxSize := settings.inMemoryCacheSize
		if maxSize == 0 {
			maxSize = 256_000_000
		}

		ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
			NumCounters: 1e7, // number of keys to track frequency of (10M).
			MaxCost:     maxSize,
			BufferItems: 64,
		})
		if err != nil {
			return nil, fmt.Errorf("creating ristretto instance: %w", err)
		}

		ristrettoStore := ristretto_store.NewRistretto(ristrettoCache)
		caches = append(caches, cache.New[string](ristrettoStore))
	}

	// Handle Redis connection (either simple connection string or full URL)
	if settings.redisConnection != "" || settings.redisConnectionStr != "" {
		var redisClient *redis.Client

		if settings.redisConnectionStr != "" {
			opts, err := redis.ParseURL(settings.redisConnectionStr)
			if err != nil {
				return nil, fmt.Errorf("parsing Redis connection string: %w", err)
			}
			redisClient = redis.NewClient(opts)
		} else {
			// Use legacy simple connection format for backward compatibility
			redisClient = redis.NewClient(&redis.Options{Addr: settings.redisConnection})
		}

		if settings.expiration != 0 {
			expiration = settings.expiration
		}

		redisStore := redis_store.NewRedis(redisClient, store.WithExpiration(expiration))

		if err := pingRedisConnectionLoop(redisClient); err != nil {
			return nil, err
		}

		caches = append(caches, cache.New[string](redisStore))
	}

	if len(caches) == 0 {
		return nil, errors.New("no cache store is configured")
	}

	// Initialize chained cache
	var cacheManager cache.CacheInterface[string]

	if settings.prometheusPrefix == "" {
		cacheManager = cache.NewChain(caches...)
	} else {
		var promOpts []metrics.PrometheusOption

		if len(settings.prometheusNamespace) > 0 {
			promOpts = append(promOpts, metrics.WithNamespace(settings.prometheusNamespace))
		}

		promMetrics := metrics.NewPrometheus(settings.prometheusPrefix, promOpts...)

		cacheManager = cache.NewMetric(
			promMetrics,
			cache.NewChain(caches...),
		)
	}

	return &GoCacheWrapper{
		prefix:       prefix,
		cacheManager: cacheManager,
		expiration:   expiration,
		gzip:         settings.gzip,
	}, nil
}

func (gcw *GoCacheWrapper) getKey(key []byte) string {
	b64 := base64.StdEncoding.EncodeToString(key)

	return fmt.Sprintf("%s::%s", gcw.prefix, b64)
}

func (gcw *GoCacheWrapper) Get(ctx context.Context, key []byte) ([]byte, error) {
	strKey := gcw.getKey(key)

	value, err := gcw.cacheManager.Get(ctx, strKey)
	if err != nil {
		if _, isError := lo.ErrorsAs[*store.NotFound](err); isError {
			return nil, ErrCacheMiss
		}

		return nil, fmt.Errorf("getting data: %w", err)
	}

	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, fmt.Errorf("decoding cache value: %w", err)
	}

	var resultBuffer bytes.Buffer

	if err = gunzipWrite(&resultBuffer, data); err != nil {
		log.Printf("decompressing data: %s\n", err.Error())
		return data, nil
	}

	return resultBuffer.Bytes(), nil
}

func (gcw *GoCacheWrapper) Set(ctx context.Context, key []byte, value []byte) error {
	strKey := gcw.getKey(key)

	var buffer bytes.Buffer

	if gcw.gzip {
		w := gzip.NewWriter(&buffer)
		defer w.Close()

		if err := gzipWrite(&buffer, value); err != nil {
			return fmt.Errorf("writting value to buffer: %w", err)
		}
	} else {
		if _, err := buffer.Write(value); err != nil {
			return fmt.Errorf("writting value to buffer: %w", err)
		}
	}

	return gcw.cacheManager.Set(
		ctx,
		strKey,
		base64.StdEncoding.EncodeToString(buffer.Bytes()),
		store.WithExpiration(gcw.expiration),
	)
}

func (gcw *GoCacheWrapper) Delete(ctx context.Context, key []byte) error {
	strKey := gcw.getKey(key)

	return gcw.cacheManager.Delete(ctx, strKey)
}

func gunzipWrite(w io.Writer, data []byte) error {
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("creating gzip reader: %w", err)
	}

	defer gr.Close()

	if _, err = io.Copy(w, gr); err != nil {
		return fmt.Errorf("coping gzip bytes: %w", err)
	}

	return nil
}

func gzipWrite(w io.Writer, data []byte) error {
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	if err != nil {
		return fmt.Errorf("initializing gzip: %w", err)
	}

	defer gw.Close()

	if _, err = gw.Write(data); err != nil {
		return fmt.Errorf("writing gzip data: %w", err)
	}

	return nil
}

func pingRedisConnectionLoop(redisClient *redis.Client) error {
	var errList []error

	for i := 0; i < pingRedisConnectionMaxRetries; i++ {
		err := pingRedisConnection(redisClient)
		if err != nil {
			errList = append(errList, err)
			time.Sleep(pingRedisConnectionInterval)

			continue
		}

		return nil
	}

	return fmt.Errorf("pinging redis: %w", errors.Join(errList...))
}

func pingRedisConnection(redisClient *redis.Client) error {
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return fmt.Errorf("pinging redis: %w", err)
	}

	return nil
}
