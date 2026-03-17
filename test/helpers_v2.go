package test

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/NSXBet/protoc-gen-go-cache-manager/gen/go/nsx/testapp"
	"github.com/NSXBet/protoc-gen-go-cache-manager/pkg/gocachemanager/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

const v2UserDetailsPrefix = "usercache::userdetails"

// userCacheManagerV2 creates a v2 cache manager for UserDetails (raw []byte in Redis, no base64/gzip).
func userCacheManagerV2(
	t *testing.T,
	redisEndpoint string,
	options ...v2.CacheOption,
) *v2.CacheManagerV2[*testapp.UserDetailsRequest, *testapp.UserDetailsResponse] {
	t.Helper()

	options = append(options, v2.WithRedisConnection(redisEndpoint))

	manager, err := v2.NewCacheManager(
		v2UserDetailsPrefix,
		func() *testapp.UserDetailsResponse { return &testapp.UserDetailsResponse{} },
		func(_ context.Context, input *testapp.UserDetailsRequest, _ map[string]any) (*testapp.UserDetailsResponse, error) {
			return &testapp.UserDetailsResponse{
				User: &testapp.User{
					UserId: input.UserId,
					Name:   "Test User",
					Email:  "test@user.com",
				},
			}, nil
		},
		options...,
	)
	require.NoError(t, err)
	return manager
}

// redisKeyV2 returns the cache key string used by v2 (prefix::hex(marshalled input)).
func redisKeyV2[TInput proto.Message](t *testing.T, prefix string, input TInput) string {
	t.Helper()

	key, err := proto.MarshalOptions{
		Deterministic: true,
	}.Marshal(input)
	require.NoError(t, err)

	return fmt.Sprintf("%s::%s", prefix, hex.EncodeToString(key))
}
