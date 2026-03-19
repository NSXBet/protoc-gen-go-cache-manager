package v2

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"crypto/tls"

	"github.com/dgraph-io/ristretto"
	redis "github.com/redis/go-redis/v9"
)

// ErrCacheMiss is returned when the key is not found in any store.
var ErrCacheMiss = errors.New("cache miss")

// GoCacheWrapperV2 is a cache wrapper that stores raw []byte (no base64, no gzip).
// It chains an optional in-memory Ristretto cache with an optional Redis backend.
type GoCacheWrapperV2 struct {
	prefix     string
	expiration time.Duration
	memory     *ristretto.Cache
	redis      *redis.Client
	metrics    *prometheusMetrics
}

// NewGoCacheWrapper creates a v2 cache wrapper. Values are stored as []byte directly.
func NewGoCacheWrapper(
	prefix string,
	expiration time.Duration,
	settings *CacheSettings,
) (*GoCacheWrapperV2, error) {
	w := &GoCacheWrapperV2{
		prefix:     prefix,
		expiration: expiration,
	}

	if settings.expiration != 0 {
		w.expiration = settings.expiration
	}

	if !settings.skipInMemoryCache {
		maxSize := settings.inMemoryCacheSize
		if maxSize == 0 {
			maxSize = 256_000_000
		}
		ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
			NumCounters: 1e7,
			MaxCost:     maxSize,
			BufferItems: 64,
		})
		if err != nil {
			return nil, fmt.Errorf("creating ristretto instance: %w", err)
		}
		w.memory = ristrettoCache
	}

	if settings.redisConnection != "" {
		redisOpts := &redis.Options{Addr: settings.redisConnection}
		if settings.redisPassword != "" {
			redisOpts.Password = settings.redisPassword
		}
		if settings.redisTLS {
			redisOpts.TLSConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
		}
		w.redis = redis.NewClient(redisOpts)
	}

	if w.memory == nil && w.redis == nil {
		return nil, errors.New("no cache store is configured")
	}

	if settings.prometheusPrefix != "" {
		w.metrics = newPrometheusMetrics(settings.prometheusPrefix, settings.prometheusNamespace)
	}

	return w, nil
}

func (w *GoCacheWrapperV2) getKey(key []byte) string {
	return w.prefix + "::" + hex.EncodeToString(key)
}

// Get returns the value for key, or ErrCacheMiss if not found.
// L1 (Ristretto) is tried first, then L2 (Redis). Returned slice is a copy.
func (w *GoCacheWrapperV2) Get(ctx context.Context, key []byte) ([]byte, error) {
	strKey := w.getKey(key)

	if w.memory != nil {
		if val, ok := w.memory.Get(strKey); ok {
			b, ok := val.([]byte)
			if ok && len(b) > 0 {
				if w.metrics != nil {
					w.metrics.recordHit()
				}
				out := make([]byte, len(b))
				copy(out, b)
				return out, nil
			}
		}
	}

	if w.redis != nil {
		data, err := w.redis.Get(ctx, strKey).Bytes()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				if w.metrics != nil {
					w.metrics.recordMiss()
				}

				return nil, ErrCacheMiss
			}

			return nil, fmt.Errorf("redis get: %w", err)
		}

		if w.metrics != nil {
			w.metrics.recordHit()
		}

		// Backfill L1 so next hit is from memory
		if w.memory != nil && len(data) > 0 {
			copied := make([]byte, len(data))
			copy(copied, data)
			_ = w.memory.SetWithTTL(strKey, copied, int64(len(copied)), w.expiration)
		}
		out := make([]byte, len(data))
		copy(out, data)
		return out, nil
	}

	if w.metrics != nil {
		w.metrics.recordMiss()
	}

	return nil, ErrCacheMiss
}

// Set stores the value for key in all configured stores (memory and/or Redis).
func (w *GoCacheWrapperV2) Set(ctx context.Context, key []byte, value []byte) error {
	strKey := w.getKey(key)
	cost := int64(len(value))

	if w.memory != nil {
		copied := make([]byte, len(value))
		copy(copied, value)
		if !w.memory.SetWithTTL(strKey, copied, cost, w.expiration) {
			// Ristretto may drop; continue to Redis
		}
	}

	if w.redis != nil {
		if err := w.redis.Set(ctx, strKey, value, w.expiration).Err(); err != nil {
			if w.metrics != nil {
				w.metrics.recordSetError()
			}

			return fmt.Errorf("redis set: %w", err)
		}
	}

	if w.metrics != nil {
		w.metrics.recordSetSuccess()
	}

	return nil
}

// Delete removes the key from all configured stores.
func (w *GoCacheWrapperV2) Delete(ctx context.Context, key []byte) error {
	strKey := w.getKey(key)

	if w.memory != nil {
		w.memory.Del(strKey)
	}

	if w.redis != nil {
		if err := w.redis.Del(ctx, strKey).Err(); err != nil {
			if w.metrics != nil {
				w.metrics.recordDeleteError()
			}

			return fmt.Errorf("redis del: %w", err)
		}
	}

	if w.metrics != nil {
		w.metrics.recordDeleteSuccess()
	}

	return nil
}
