package test

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/NSXBet/protoc-gen-go-cache-manager/gen/go/nsx/testapp"
	"github.com/NSXBet/protoc-gen-go-cache-manager/pkg/gocachemanager"
	"github.com/google/uuid"
	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (suite *TestSuite) TestCanGetDataFromCache() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	manager := userCacheManager(t, redisEndpoint)
	keyInput := &testapp.UserDetailsRequest{
		UserId: "1",
	}

	// ACT
	userDetails, err := manager.GetUserDetails(
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
	redisKey := redisKey(t, keyInput)
	data, err := suite.redisClient.Get(context.Background(), redisKey).Result()
	require.NoError(t, err)
	require.NotNil(t, data)

	bytes, err := base64.StdEncoding.DecodeString(data)
	require.NoError(t, err)

	userDetailsResponse := &testapp.UserDetailsResponse{}
	err = proto.Unmarshal(bytes, userDetailsResponse)
	require.NoError(t, err)

	require.NotNil(t, userDetailsResponse.User)
	require.Equal(t, "Test User", userDetailsResponse.User.GetName())
	require.Equal(t, "1", userDetailsResponse.User.GetUserId())
	require.Equal(t, "test@user.com", userDetailsResponse.User.GetEmail())

	// delete from redis and shout get from memory
	err = suite.redisClient.Del(context.Background(), redisKey).Err()
	require.NoError(t, err)

	userDetails, err = manager.GetUserDetails(
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

func (suite *TestSuite) TestCanRefreshDataFromCache() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	manager := userCacheManager(t, redisEndpoint)
	keyInput := &testapp.UserDetailsRequest{
		UserId: uuid.NewString(),
	}

	// ACT
	userDetails, err := manager.RefreshUserDetails(
		context.Background(),
		keyInput,
	)

	// ASSERT
	require.NoError(t, err)
	require.NotNil(t, userDetails)
	require.NotNil(t, userDetails.User)

	user := userDetails.User

	require.Equal(t, "Test User", user.GetName())
	require.Equal(t, keyInput.UserId, user.GetUserId())
	require.Equal(t, "test@user.com", user.GetEmail())

	// get from redis
	redisKey := redisKey(t, keyInput)
	data, err := suite.redisClient.Get(context.Background(), redisKey).Result()
	require.NoError(t, err)
	require.NotNil(t, data)

	bytes, err := base64.StdEncoding.DecodeString(data)
	require.NoError(t, err)

	userDetailsResponse := &testapp.UserDetailsResponse{}
	err = proto.Unmarshal(bytes, userDetailsResponse)
	require.NoError(t, err)

	require.NotNil(t, userDetailsResponse.User)
	require.Equal(t, "Test User", userDetailsResponse.User.GetName())
	require.Equal(t, keyInput.UserId, userDetailsResponse.User.GetUserId())
	require.Equal(t, "test@user.com", userDetailsResponse.User.GetEmail())
}

func (suite *TestSuite) TestCanGetTournament() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	loader := tournamentLoader(t)
	manager := tournamentCacheManager(t, redisEndpoint)
	keyInput := &testapp.MainTournamentsRequest{
		Empty: &emptypb.Empty{},
	}

	// ACT
	tourn, err := manager.GetMainTournaments(
		context.Background(),
		keyInput,
		map[string]any{
			"loader": loader,
		},
	)
	require.NoError(t, err)
	require.NotNil(t, tourn)

	// ASSERT
	require.Len(t, tourn.Tournaments, 2)

	for _, tournament := range tourn.Tournaments {
		require.Equal(t, "Test Tournament", tournament.GetName())
		require.Equal(t, "image.jpg", tournament.GetImageUrl())
		require.Equal(t, "https://test.com", tournament.GetUrl())
		require.Equal(t, 1.1, tournament.GetDbl())
		require.Equal(t, float32(1.2), tournament.GetFlt())
		require.Equal(t, int32(1), tournament.GetNum32())
		require.Equal(t, int64(2), tournament.GetNum64())
		require.Equal(t, uint32(3), tournament.GetUnum32())
		require.Equal(t, uint64(4), tournament.GetUnum64())
		require.Equal(t, int32(5), tournament.GetSnum32())
		require.Equal(t, int64(6), tournament.GetSnum64())
		require.Equal(t, uint32(7), tournament.GetFnum32())
		require.Equal(t, uint64(8), tournament.GetFnum64())
		require.Equal(t, int32(9), tournament.GetSfnum32())
		require.Equal(t, int64(10), tournament.GetSfnum64())
		require.True(t, tournament.GetIsActive())
		require.Equal(t, []byte("data"), tournament.GetData())
		require.Equal(t, testapp.TournamentType_TOURNAMENT_TYPE_DAILY, tournament.GetType())
		require.Equal(t, map[string]string{"key": "value"}, tournament.GetMetadata())

		require.Len(t, tournament.GetEvents(), 2)

		for index, event := range tournament.GetEvents() {
			dt := time.Date(2021, 1, index+1, 0, 0, 0, 0, time.UTC)

			require.Equal(t, fmt.Sprintf("Event %d", index+1), event.GetName())
			require.Equal(t, []string{
				fmt.Sprintf("Player %d", index*2+1),
				fmt.Sprintf("Player %d", index*2+2),
			}, event.GetPlayers())
			require.Equal(t, dt, event.GetStartTime().AsTime())
		}
	}
}

func (suite *TestSuite) TestSkipInMemoryCache() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	manager := userCacheManager(
		t,
		redisEndpoint,
		gocachemanager.WithSkipInMemoryCache(),
	)
	keyInput := &testapp.UserDetailsRequest{
		UserId: "1",
	}

	redisKey := redisKey(t, keyInput)

	// cleanup redis
	err = suite.redisClient.Del(context.Background(), redisKey).Err()
	require.NoError(t, err)

	// ACT
	userDetails, err := manager.GetUserDetails(
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
	data, err := suite.redisClient.Get(context.Background(), redisKey).Result()
	require.NoError(t, err)
	require.NotNil(t, data)

	bytes, err := base64.StdEncoding.DecodeString(data)
	require.NoError(t, err)

	userDetailsResponse := &testapp.UserDetailsResponse{}
	err = proto.Unmarshal(bytes, userDetailsResponse)
	require.NoError(t, err)

	require.NotNil(t, userDetailsResponse.User)
	require.Equal(t, "Test User", userDetailsResponse.User.GetName())
	require.Equal(t, "1", userDetailsResponse.User.GetUserId())
	require.Equal(t, "test@user.com", userDetailsResponse.User.GetEmail())
}

func (suite *TestSuite) TestCanDeleteDataFromCache() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	manager := userCacheManager(t, redisEndpoint)
	keyInput := &testapp.UserDetailsRequest{
		UserId: "1",
	}

	userDetails, err := manager.GetUserDetails(
		context.Background(),
		keyInput,
	)
	require.NoError(t, err)
	require.NotNil(t, userDetails)
	require.NotNil(t, userDetails.User)

	// get from redis
	rk := redisKey(t, keyInput)
	data, err := suite.redisClient.Get(context.Background(), rk).Result()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	// ACT
	err = manager.DeleteUserDetails(
		context.Background(),
		keyInput,
	)

	// ASSERT
	require.NoError(t, err)

	// get from redis
	rk = redisKey(t, keyInput)
	data, err = suite.redisClient.Get(context.Background(), rk).Result()
	require.ErrorIs(t, err, redis.Nil)
	require.Empty(t, data)
}

func (suite *TestSuite) TestCanReplaceDataFromCache() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	manager := userCacheManager(t, redisEndpoint)
	keyInput := &testapp.UserDetailsRequest{
		UserId: "1",
	}

	userDetails, err := manager.GetUserDetails(
		context.Background(),
		keyInput,
	)
	require.NoError(t, err)
	require.NotNil(t, userDetails)
	require.NotNil(t, userDetails.User)

	// get from redis
	rk := redisKey(t, keyInput)
	data, err := suite.redisClient.Get(context.Background(), rk).Result()
	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.NotEqual(t, "Replaced User", userDetails.User.GetName())
	require.NotEqual(t, "replaced@user.com", userDetails.User.GetEmail())

	// ACT
	replacedUser, err := manager.ReplaceUserDetails(
		context.Background(),
		keyInput,
		&testapp.UserDetailsResponse{
			User: &testapp.User{
				UserId: keyInput.UserId,
				Name:   "Replaced User",
				Email:  "replaced@user.com",
			},
		},
	)

	// ASSERT
	require.NoError(t, err)
	require.Equal(t, "Replaced User", replacedUser.User.GetName())
	require.Equal(t, "1", replacedUser.User.GetUserId())
	require.Equal(t, "replaced@user.com", replacedUser.User.GetEmail())

	// get from redis
	rk = redisKey(t, keyInput)
	data, err = suite.redisClient.Get(context.Background(), rk).Result()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	bytes, err := base64.StdEncoding.DecodeString(data)
	require.NoError(t, err)

	var userDetailsResponse testapp.UserDetailsResponse

	err = proto.Unmarshal(bytes, &userDetailsResponse)
	require.NoError(t, err)

	require.NotNil(t, userDetailsResponse.User)
	require.Equal(t, "Replaced User", userDetailsResponse.User.GetName())
	require.Equal(t, "1", userDetailsResponse.User.GetUserId())
	require.Equal(t, "replaced@user.com", userDetailsResponse.User.GetEmail())
}

func (suite *TestSuite) TestShouldStoreWithoutGzipAndReadWithGzipFromCache() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	// setting user into cache without gzip option
	manager := userCacheManager(
		t,
		redisEndpoint,
	)
	keyInput := &testapp.UserDetailsRequest{
		UserId: "2",
	}

	userDetails, err := manager.GetUserDetails(
		context.Background(),
		keyInput,
	)
	require.NoError(t, err)
	require.NotNil(t, userDetails)
	require.NotNil(t, userDetails.User)

	// reading user from cache with gzip enabled
	newmanager := userCacheManager(
		t,
		redisEndpoint,
		gocachemanager.WithGzip(),
	)

	newUserDetails, err := newmanager.GetUserDetails(
		context.Background(),
		keyInput,
	)
	require.NoError(t, err)
	require.Equal(t, userDetails.GetUser().GetName(), newUserDetails.GetUser().GetName())
	require.Equal(t, userDetails.GetUser().GetUserId(), newUserDetails.GetUser().GetUserId())
	require.Equal(t, userDetails.GetUser().GetEmail(), newUserDetails.GetUser().GetEmail())
}

func (suite *TestSuite) TestShouldStoreWithGzipAndReadWithoutGzipFromCache() {
	// ARRANGE
	t := suite.T()
	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(t, err)

	// setting user into cache with gzip enabled
	manager := userCacheManager(
		t,
		redisEndpoint,
		gocachemanager.WithGzip(),
	)
	keyInput := &testapp.UserDetailsRequest{
		UserId: "2",
	}

	userDetails, err := manager.GetUserDetails(
		context.Background(),
		keyInput,
	)
	require.NoError(t, err)
	require.NotNil(t, userDetails)
	require.NotNil(t, userDetails.User)

	// reading user from cache with gzip disabled
	newmanager := userCacheManager(
		t,
		redisEndpoint,
	)

	newUserDetails, err := newmanager.GetUserDetails(
		context.Background(),
		keyInput,
	)
	require.NoError(t, err)
	require.Equal(t, userDetails.GetUser().GetName(), newUserDetails.GetUser().GetName())
	require.Equal(t, userDetails.GetUser().GetUserId(), newUserDetails.GetUser().GetUserId())
	require.Equal(t, userDetails.GetUser().GetEmail(), newUserDetails.GetUser().GetEmail())
}
