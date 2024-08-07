package test

import (
	"context"
	"fmt"
	"time"

	"github.com/NSXBet/protoc-gen-go-cache-manager/gen/go/nsx/testapp"
	"github.com/NSXBet/protoc-gen-go-cache-manager/pkg/gocachemanager"
	"github.com/google/uuid"
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

	userDetailsResponse := &testapp.UserDetailsResponse{}
	err = proto.Unmarshal([]byte(data), userDetailsResponse)
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

	manager := tournamentCacheManager(t, redisEndpoint)
	keyInput := &testapp.MainTournamentsRequest{
		Empty: &emptypb.Empty{},
	}

	// ACT
	tourn, err := manager.GetMainTournaments(
		context.Background(),
		keyInput,
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

	userDetailsResponse := &testapp.UserDetailsResponse{}
	err = proto.Unmarshal([]byte(data), userDetailsResponse)
	require.NoError(t, err)

	require.NotNil(t, userDetailsResponse.User)
	require.Equal(t, "Test User", userDetailsResponse.User.GetName())
	require.Equal(t, "1", userDetailsResponse.User.GetUserId())
	require.Equal(t, "test@user.com", userDetailsResponse.User.GetEmail())
}
