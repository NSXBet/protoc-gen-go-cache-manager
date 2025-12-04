package gocachemanager

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
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

	if settings.redisConnection != "" {
		redisOptions := &redis.Options{Addr: settings.redisConnection}
		if settings.redisPassword != "" {
			redisOptions.Password = settings.redisPassword
		}

		if settings.redisTLS {
			redisOptions.TLSConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
		}

		redisClient := redis.NewClient(redisOptions)

		if settings.expiration != 0 {
			expiration = settings.expiration
		}

		redisStore := redis_store.NewRedis(redisClient, store.WithExpiration(expiration))

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

	if isGzipCompressed(data) {
		var resultBuffer bytes.Buffer

		if err = gunzipWrite(&resultBuffer, data); err != nil {
			return nil, fmt.Errorf("decompressing data: %w", err)
		}

		return resultBuffer.Bytes(), nil
	}

	return data, nil
}

func (gcw *GoCacheWrapper) Set(ctx context.Context, key []byte, value []byte) error {
	strKey := gcw.getKey(key)

	var buffer bytes.Buffer

	if gcw.gzip {
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

// isGzipCompressed checks if data is gzip-compressed by verifying the magic number (0x1f 0x8b)
func isGzipCompressed(data []byte) bool {
	return len(data) >= 2 && data[0] == 0x1f && data[1] == 0x8b
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
