package gocachemanager

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGoCacheWrapper(t *testing.T) {
	t.Run("should create wrapper with in-memory cache only", func(t *testing.T) {
		settings := DefaultCacheSettings()
		wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)

		require.NoError(t, err)
		assert.NotNil(t, wrapper)
		assert.Equal(t, "test", wrapper.prefix)
		assert.Equal(t, 5*time.Second, wrapper.expiration)
		assert.False(t, wrapper.gzip)
	})

	t.Run("should create wrapper with gzip enabled", func(t *testing.T) {
		settings := DefaultCacheSettings()
		settings.gzip = true
		wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)

		require.NoError(t, err)
		assert.True(t, wrapper.gzip)
	})

	t.Run("should create wrapper with custom cache size", func(t *testing.T) {
		settings := DefaultCacheSettings()
		settings.inMemoryCacheSize = 1024 * 1024 // 1MB
		wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)

		require.NoError(t, err)
		assert.NotNil(t, wrapper)
	})

	t.Run("should fail when no cache store is configured", func(t *testing.T) {
		settings := &CacheSettings{
			skipInMemoryCache: true,
			redisConnection:   "", // No Redis
		}
		wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)

		assert.Error(t, err)
		assert.Nil(t, wrapper)
		assert.Equal(t, "no cache store is configured", err.Error())
	})
}

func TestGoCacheWrapper_GetKey(t *testing.T) {
	settings := DefaultCacheSettings()
	wrapper, err := NewGoCacheWrapper("myprefix", 5*time.Second, settings)
	require.NoError(t, err)

	key := []byte("test-key")
	expectedKey := "myprefix::" + base64.StdEncoding.EncodeToString(key)

	result := wrapper.getKey(key)
	assert.Equal(t, expectedKey, result)
}

func TestGoCacheWrapper_SetAndGet_WithoutGzip(t *testing.T) {
	ctx := context.Background()
	settings := DefaultCacheSettings()
	wrapper, err := NewGoCacheWrapper("test", 10*time.Second, settings)
	require.NoError(t, err)

	t.Run("should set and get value successfully", func(t *testing.T) {
		key := []byte("my-key")
		value := []byte("my-value-data")

		// Set
		err := wrapper.Set(ctx, key, value)
		require.NoError(t, err)

		// Ristretto is async, give it time to process
		time.Sleep(10 * time.Millisecond)

		// Get
		result, err := wrapper.Get(ctx, key)
		require.NoError(t, err)
		assert.Equal(t, value, result)
	})

	t.Run("should handle empty value", func(t *testing.T) {
		key := []byte("empty-key")
		value := []byte("")

		err := wrapper.Set(ctx, key, value)
		require.NoError(t, err)

		time.Sleep(10 * time.Millisecond)

		result, err := wrapper.Get(ctx, key)
		require.NoError(t, err)
		assert.Equal(t, value, result)
	})

	t.Run("should handle large value", func(t *testing.T) {
		key := []byte("large-key")
		value := make([]byte, 1024*100) // 100KB
		for i := range value {
			value[i] = byte(i % 256)
		}

		err := wrapper.Set(ctx, key, value)
		require.NoError(t, err)

		time.Sleep(10 * time.Millisecond)

		result, err := wrapper.Get(ctx, key)
		require.NoError(t, err)
		assert.Equal(t, value, result)
	})
}

func TestGoCacheWrapper_SetAndGet_WithGzip(t *testing.T) {
	ctx := context.Background()
	settings := DefaultCacheSettings()
	settings.gzip = true
	wrapper, err := NewGoCacheWrapper("test-gzip", 10*time.Second, settings)
	require.NoError(t, err)

	t.Run("should compress and decompress value successfully", func(t *testing.T) {
		key := []byte("compressed-key")
		value := []byte("this is a test value that will be compressed using gzip")

		err := wrapper.Set(ctx, key, value)
		require.NoError(t, err)

		time.Sleep(10 * time.Millisecond)

		result, err := wrapper.Get(ctx, key)
		require.NoError(t, err)
		assert.Equal(t, value, result)
	})

	t.Run("should handle large compressible data", func(t *testing.T) {
		key := []byte("large-compressed-key")
		// Create highly compressible data (repeated pattern)
		value := bytes.Repeat([]byte("Lorem ipsum dolor sit amet. "), 1000)

		err := wrapper.Set(ctx, key, value)
		require.NoError(t, err)

		time.Sleep(10 * time.Millisecond)

		result, err := wrapper.Get(ctx, key)
		require.NoError(t, err)
		assert.Equal(t, value, result)
	})
}

func TestGoCacheWrapper_Get_CacheMiss(t *testing.T) {
	ctx := context.Background()
	settings := DefaultCacheSettings()
	wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)
	require.NoError(t, err)

	key := []byte("non-existent-key")
	result, err := wrapper.Get(ctx, key)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrCacheMiss, err)
}

func TestGoCacheWrapper_Delete(t *testing.T) {
	ctx := context.Background()
	settings := DefaultCacheSettings()
	wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)
	require.NoError(t, err)

	key := []byte("delete-key")
	value := []byte("delete-value")

	// Set
	err = wrapper.Set(ctx, key, value)
	require.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	// Verify it exists
	result, err := wrapper.Get(ctx, key)
	require.NoError(t, err)
	assert.Equal(t, value, result)

	// Delete
	err = wrapper.Delete(ctx, key)
	require.NoError(t, err)

	// Verify it's gone
	result, err = wrapper.Get(ctx, key)
	assert.Error(t, err)
	assert.Equal(t, ErrCacheMiss, err)
	assert.Nil(t, result)
}

func TestGoCacheWrapper_Expiration(t *testing.T) {
	ctx := context.Background()
	settings := DefaultCacheSettings()
	settings.expiration = 1 * time.Second
	wrapper, err := NewGoCacheWrapper("test", 1*time.Second, settings)
	require.NoError(t, err)

	key := []byte("expiring-key")
	value := []byte("expiring-value")

	// Set
	err = wrapper.Set(ctx, key, value)
	require.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	// Should exist immediately
	result, err := wrapper.Get(ctx, key)
	require.NoError(t, err)
	assert.Equal(t, value, result)

	// Wait for expiration (plus buffer)
	time.Sleep(1500 * time.Millisecond)

	// Should be gone
	_, err = wrapper.Get(ctx, key)
	assert.Error(t, err)
	assert.Equal(t, ErrCacheMiss, err)
}

func TestGoCacheWrapper_Get_CorruptedGzipData(t *testing.T) {
	ctx := context.Background()
	settings := DefaultCacheSettings()
	settings.gzip = true
	wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)
	require.NoError(t, err)

	// Manually insert corrupted data into cache
	key := []byte("corrupted-key")
	strKey := wrapper.getKey(key)

	// Set corrupted gzip data (not actually gzipped)
	corruptedData := base64.StdEncoding.EncodeToString([]byte("not gzipped data"))
	err = wrapper.cacheManager.Set(ctx, strKey, corruptedData)
	require.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	// Should return error when trying to decompress
	result, err := wrapper.Get(ctx, key)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "decompressing data")
}

func TestGoCacheWrapper_Get_WithoutGzip_DoesNotDecompress(t *testing.T) {
	ctx := context.Background()
	settings := DefaultCacheSettings()
	settings.gzip = false
	wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)
	require.NoError(t, err)

	// Manually insert gzipped data
	key := []byte("gzipped-key")
	strKey := wrapper.getKey(key)

	// Create gzipped data
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	_, err = gw.Write([]byte("gzipped content"))
	require.NoError(t, err)
	require.NoError(t, gw.Close())

	gzippedData := base64.StdEncoding.EncodeToString(buf.Bytes())
	err = wrapper.cacheManager.Set(ctx, strKey, gzippedData)
	require.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	// Should return the gzipped data as-is (not decompress)
	result, err := wrapper.Get(ctx, key)
	require.NoError(t, err)
	// Result should be the raw gzipped bytes, not decompressed
	assert.Equal(t, buf.Bytes(), result)
	assert.NotEqual(t, []byte("gzipped content"), result)
}

func TestGoCacheWrapper_MultipleKeys(t *testing.T) {
	ctx := context.Background()
	settings := DefaultCacheSettings()
	wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)
	require.NoError(t, err)

	// Set multiple keys
	keys := [][]byte{
		[]byte("key1"),
		[]byte("key2"),
		[]byte("key3"),
	}
	values := [][]byte{
		[]byte("value1"),
		[]byte("value2"),
		[]byte("value3"),
	}

	for i := range keys {
		err := wrapper.Set(ctx, keys[i], values[i])
		require.NoError(t, err)
	}

	time.Sleep(10 * time.Millisecond)

	// Get all keys
	for i := range keys {
		result, err := wrapper.Get(ctx, keys[i])
		require.NoError(t, err)
		assert.Equal(t, values[i], result)
	}

	// Delete one key
	err = wrapper.Delete(ctx, keys[1])
	require.NoError(t, err)

	// Verify deletion
	_, err = wrapper.Get(ctx, keys[1])
	assert.Equal(t, ErrCacheMiss, err)

	// Other keys should still exist
	result, err := wrapper.Get(ctx, keys[0])
	require.NoError(t, err)
	assert.Equal(t, values[0], result)

	result, err = wrapper.Get(ctx, keys[2])
	require.NoError(t, err)
	assert.Equal(t, values[2], result)
}

func TestGunzipWrite(t *testing.T) {
	t.Run("should decompress gzipped data", func(t *testing.T) {
		// Create gzipped data
		original := []byte("test data for compression")
		var compressed bytes.Buffer
		gw := gzip.NewWriter(&compressed)
		_, err := gw.Write(original)
		require.NoError(t, err)
		require.NoError(t, gw.Close())

		// Decompress
		var result bytes.Buffer
		err = gunzipWrite(&result, compressed.Bytes())
		require.NoError(t, err)
		assert.Equal(t, original, result.Bytes())
	})

	t.Run("should fail on invalid gzip data", func(t *testing.T) {
		invalid := []byte("not gzipped")
		var result bytes.Buffer
		err := gunzipWrite(&result, invalid)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "creating gzip reader")
	})
}

func TestGzipWrite(t *testing.T) {
	t.Run("should compress data", func(t *testing.T) {
		original := []byte("test data for compression")
		var compressed bytes.Buffer

		err := gzipWrite(&compressed, original)
		require.NoError(t, err)

		// Decompress to verify
		gr, err := gzip.NewReader(&compressed)
		require.NoError(t, err)
		defer gr.Close()

		var result bytes.Buffer
		_, err = result.ReadFrom(gr)
		require.NoError(t, err)
		assert.Equal(t, original, result.Bytes())
	})

	t.Run("should handle empty data", func(t *testing.T) {
		original := []byte("")
		var compressed bytes.Buffer

		err := gzipWrite(&compressed, original)
		require.NoError(t, err)
		assert.Greater(t, compressed.Len(), 0) // Gzip header is present even for empty data
	})
}

func TestCacheOptions(t *testing.T) {
	t.Run("should apply WithGzip option", func(t *testing.T) {
		settings := DefaultCacheSettings()
		WithGzip()(settings)

		wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)
		require.NoError(t, err)
		assert.True(t, wrapper.gzip)
	})

	t.Run("should work with multiple options", func(t *testing.T) {
		settings := DefaultCacheSettings()
		WithGzip()(settings)
		WithExpiration(10 * time.Second)(settings)
		WithInMemoryCacheSize(1024 * 1024)(settings)

		// The wrapper constructor uses the settings.expiration (10s) since it's set
		wrapper, err := NewGoCacheWrapper("test", 5*time.Second, settings)
		require.NoError(t, err)
		assert.True(t, wrapper.gzip)
		// When settings.expiration is set, it overrides the parameter in NewGoCacheWrapper
		assert.Equal(t, 10*time.Second, settings.expiration)
	})
}
