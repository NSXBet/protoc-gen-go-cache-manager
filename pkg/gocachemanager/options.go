package gocachemanager

// CacheSettings contains the configuration for building a cache manager.
type CacheSettings struct {
	// RedisConnection is the connection string for the Redis server.
	// Defaults to empty string, meaning no Redis connection. So no cache will be used in Redis.
	redisConnection string
}

// DefaultCacheSettings returns the default cache settings.
func DefaultCacheSettings() *CacheSettings {
	return &CacheSettings{
		redisConnection: "", // No Redis connection by default
	}
}

// CacheOption is an interface for applying cache settings.
type CacheOption func(*CacheSettings)

// WithRedisConnection is a cache option for setting the Redis connection string.
func WithRedisConnection(redisConnection string) CacheOption {
	return func(settings *CacheSettings) {
		settings.redisConnection = redisConnection
	}
}
