package v2

import "time"

// CacheSettings contains the configuration for building a v2 cache manager.
// Values are stored as raw []byte (no base64, no gzip) for optimal protobuf caching.
type CacheSettings struct {
	redisConnection     string
	skipInMemoryCache   bool
	inMemoryCacheSize   int64
	prometheusPrefix    string
	prometheusNamespace string
	expiration          time.Duration
	redisPassword       string
	redisTLS            bool
}

// DefaultCacheSettings returns the default cache settings for v2.
func DefaultCacheSettings() *CacheSettings {
	return &CacheSettings{
		redisConnection:   "",
		skipInMemoryCache: false,
	}
}

// CacheOption applies a setting to CacheSettings.
type CacheOption func(*CacheSettings)

// WithRedisConnection sets the Redis connection string.
func WithRedisConnection(redisConnection string) CacheOption {
	return func(s *CacheSettings) {
		s.redisConnection = redisConnection
	}
}

// WithSkipInMemoryCache skips the in-memory cache and uses Redis only.
func WithSkipInMemoryCache() CacheOption {
	return func(s *CacheSettings) {
		s.skipInMemoryCache = true
	}
}

// WithPrometheusPrefix sets the Prometheus metrics prefix (service label). When set, cache operations are instrumented.
func WithPrometheusPrefix(prometheusPrefix string) CacheOption {
	return func(s *CacheSettings) {
		s.prometheusPrefix = prometheusPrefix
	}
}

// WithPrometheusNamespace sets the Prometheus metrics namespace (default "cache").
func WithPrometheusNamespace(prometheusNamespace string) CacheOption {
	return func(s *CacheSettings) {
		s.prometheusNamespace = prometheusNamespace
	}
}

// WithInMemoryCacheSize sets the in-memory cache size in bytes (default 256MB).
func WithInMemoryCacheSize(size int64) CacheOption {
	return func(s *CacheSettings) {
		s.inMemoryCacheSize = size
	}
}

// WithExpiration sets the cache expiration duration.
func WithExpiration(expiration time.Duration) CacheOption {
	return func(s *CacheSettings) {
		s.expiration = expiration
	}
}

// WithRedisPassword sets the Redis password.
func WithRedisPassword(redisPassword string) CacheOption {
	return func(s *CacheSettings) {
		s.redisPassword = redisPassword
	}
}

// WithRedisTLS enables TLS for Redis.
func WithRedisTLS(redisTLS bool) CacheOption {
	return func(s *CacheSettings) {
		s.redisTLS = redisTLS
	}
}
