package test

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/NSXBet/protoc-gen-go-cache-manager/gen/go/nsx/testapp"
	"github.com/NSXBet/protoc-gen-go-cache-manager/pkg/gocachemanager"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func userCacheManager(
	t *testing.T,
	redisEndpoint string,
	options ...gocachemanager.CacheOption,
) *testapp.UserCacheManager {
	t.Helper()

	options = append(
		options,
		gocachemanager.WithRedisConnection(redisEndpoint),
	)

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
		options...,
	)
	require.NoError(t, err)

	return userCacheManager
}

func tournamentCacheManager(t *testing.T, redisEndpoint string) *testapp.TournamentCacheManager {
	t.Helper()

	manager, err := testapp.NewTournamentCacheManager(
		func(_ context.Context, input *testapp.MainTournamentsRequest) (*testapp.MainTournamentsResponse, error) {
			return &testapp.MainTournamentsResponse{
				Tournaments: []*testapp.Tournament{
					testTournament(t, true),
					testTournament(t, false),
				},
			}, nil
		},
		gocachemanager.WithRedisConnection(redisEndpoint),
		gocachemanager.WithPrometheusPrefix("test"),
	)
	require.NoError(t, err)

	return manager
}

func testTournament(t *testing.T, useStringPrize bool) *testapp.Tournament {
	t.Helper()
	dt := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	dt2 := time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)

	tourn := &testapp.Tournament{
		Id:       "1",
		Name:     "Test Tournament",
		ImageUrl: "image.jpg",
		Url:      "https://test.com",
		Dbl:      1.1,
		Flt:      1.2,
		Num32:    1,
		Num64:    2,
		Unum32:   3,
		Unum64:   4,
		Snum32:   5,
		Snum64:   6,
		Fnum32:   7,
		Fnum64:   8,
		Sfnum32:  9,
		Sfnum64:  10,
		IsActive: true,
		Data:     []byte("data"),
		Type:     testapp.TournamentType_TOURNAMENT_TYPE_DAILY,
		Events: []*testapp.Event{
			{
				StartTime: timestamppb.New(dt),
				Name:      "Event 1",
				Players:   []string{"Player 1", "Player 2"},
			},
			{
				StartTime: timestamppb.New(dt2),
				Name:      "Event 2",
				Players:   []string{"Player 3", "Player 4"},
			},
		},
		Metadata: map[string]string{
			"key": "value",
		},
	}

	if useStringPrize {
		tourn.Prize = &testapp.Tournament_PrizeVal{PrizeVal: "prize"}
	} else {
		tourn.Prize = &testapp.Tournament_PrizeNum{PrizeNum: 10}
	}

	return tourn
}

func redisKey[TInput proto.Message](t *testing.T, input TInput) string {
	t.Helper()

	key, err := proto.MarshalOptions{
		Deterministic: true,
	}.Marshal(input)
	require.NoError(t, err)

	b64 := base64.StdEncoding.EncodeToString(key)
	cacheKey := fmt.Sprintf("usercache::userdetails::%s", b64)
	return cacheKey
}
