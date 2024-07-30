package test

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/NSXBet/go-cache-manager/gen/go/nsx/testapp"
	"github.com/NSXBet/go-cache-manager/pkg/gocachemanager"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (suite *TestSuite) TestCanGetDataFromCache() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	userCacheManager, err := testapp.NewUserCacheManager(
		func(_ context.Context, input *testapp.UserDetailsRequest) (*testapp.UserDetailsResponse, error) {
			return &testapp.UserDetailsResponse{
				User: &testapp.User{
					UserId: input.UserId,
					Name:   "Test User",
					Email:  "test@user.com",
				},
			}, nil
		},
		gocachemanager.WithRedisConnection(redisEndpoint),
	)
	require.NoError(t, err)
	keyInput := &testapp.UserDetailsRequest{
		UserId: "1",
	}

	// ACT
	userDetails, err := userCacheManager.GetUserDetails(
		context.Background(),
		keyInput,
	)

	// ASSERT
	require.NoError(t, err)
	require.NotNil(t, userDetails)
	require.NotNil(t, userDetails.User)

	user := userDetails.User

	require.Equal(t, "Test User", user.GetName())
	require.Equal(t, "1", user.GetUserId())
	require.Equal(t, "test@user.com", user.GetEmail())

	// get from redis
	key, err := proto.MarshalOptions{
		Deterministic: true,
	}.Marshal(keyInput)
	require.NoError(t, err)

	b64 := base64.StdEncoding.EncodeToString(key)
	redisKey := fmt.Sprintf("userdetails:%s", b64)
	data, err := suite.redisClient.Get(context.Background(), redisKey).Result()
	require.NoError(t, err)
	require.NotNil(t, data)

	userDetailsResponse := &testapp.UserDetailsResponse{}
	err = proto.Unmarshal([]byte(data), userDetailsResponse)
	require.NoError(t, err)

	require.NotNil(t, userDetailsResponse.User)
	require.Equal(t, "Test User", userDetailsResponse.User.GetName())
	require.Equal(t, "1", userDetailsResponse.User.GetUserId())
	require.Equal(t, "test@user.com", userDetailsResponse.User.GetEmail())

	// delete from redis and shout get from memory
	err = suite.redisClient.Del(context.Background(), redisKey).Err()
	require.NoError(t, err)

	userDetails, err = userCacheManager.GetUserDetails(
		context.Background(),
		keyInput,
	)
	require.NoError(t, err)
	require.NotNil(t, userDetails)
	require.NotNil(t, userDetails.User)

	user = userDetails.User
	require.Equal(t, "Test User", user.GetName())
	require.Equal(t, "1", user.GetUserId())
}
