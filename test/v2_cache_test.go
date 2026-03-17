package test

import (
	"context"

	"github.com/NSXBet/protoc-gen-go-cache-manager/gen/go/nsx/testapp"
	"github.com/NSXBet/protoc-gen-go-cache-manager/pkg/gocachemanager/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

// TestV2_RedisStoresRawBytes ensures v2 stores raw protobuf bytes in Redis (no base64).
func (suite *TestSuite) TestV2_RedisStoresRawBytes() {
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	manager := userCacheManagerV2(t, redisEndpoint)
	keyInput := &testapp.UserDetailsRequest{UserId: "v2-raw-bytes"}

	_, err = manager.Get(context.Background(), keyInput)
	require.NoError(t, err)

	rk := redisKeyV2(t, v2UserDetailsPrefix, keyInput)
	data, err := suite.redisClient.Get(context.Background(), rk).Bytes()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	// v2 stores raw proto bytes: unmarshal directly (no base64 decode).
	resp := &testapp.UserDetailsResponse{}
	err = proto.Unmarshal(data, resp)
	require.NoError(t, err)
	require.NotNil(t, resp.User)
	require.Equal(t, "v2-raw-bytes", resp.User.GetUserId())
	require.Equal(t, "Test User", resp.User.GetName())
}

// TestV2_GetRefreshReplaceDelete runs the same flow as v1 tests using v2 SDK.
func (suite *TestSuite) TestV2_GetRefreshReplaceDelete() {
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	manager := userCacheManagerV2(t, redisEndpoint)
	keyInput := &testapp.UserDetailsRequest{UserId: uuid.NewString()}

	got, err := manager.Get(context.Background(), keyInput)
	require.NoError(t, err)
	require.NotNil(t, got.User)
	require.Equal(t, keyInput.UserId, got.User.GetUserId())

	_, err = manager.Refresh(context.Background(), keyInput)
	require.NoError(t, err)

	newVal := &testapp.UserDetailsResponse{
		User: &testapp.User{UserId: keyInput.UserId, Name: "Updated", Email: "updated@test.com"},
	}
	replaced, err := manager.Replace(context.Background(), keyInput, newVal)
	require.NoError(t, err)
	require.Equal(t, "Updated", replaced.User.GetName())

	got2, err := manager.Get(context.Background(), keyInput)
	require.NoError(t, err)
	require.Equal(t, "Updated", got2.User.GetName())

	err = manager.Delete(context.Background(), keyInput)
	require.NoError(t, err)

	rk := redisKeyV2(t, v2UserDetailsPrefix, keyInput)
	_, err = suite.redisClient.Get(context.Background(), rk).Result()
	require.Error(t, err) // key should be gone
}

// TestV2_MemoryOnly ensures v2 works with in-memory cache only (no Redis).
func (suite *TestSuite) TestV2_MemoryOnly() {
	t := suite.T()

	manager, err := v2.NewCacheManager(
		"memory_only",
		func() *testapp.UserDetailsResponse { return &testapp.UserDetailsResponse{} },
		func(_ context.Context, input *testapp.UserDetailsRequest, _ map[string]any) (*testapp.UserDetailsResponse, error) {
			return &testapp.UserDetailsResponse{
				User: &testapp.User{UserId: input.UserId, Name: "InMem", Email: "mem@test.com"},
			}, nil
		},
	)
	require.NoError(t, err)

	keyInput := &testapp.UserDetailsRequest{UserId: "mem-1"}
	got, err := manager.Get(context.Background(), keyInput)
	require.NoError(t, err)
	require.Equal(t, "InMem", got.User.GetName())
}
