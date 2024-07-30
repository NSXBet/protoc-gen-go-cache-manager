package test

import (
	"context"
	"testing"

	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	redis_container "github.com/testcontainers/testcontainers-go/modules/redis"
)

type TestSuite struct {
	suite.Suite

	redisContainer *redis_container.RedisContainer
	redisClient    *redis.Client
}

func (suite *TestSuite) SetupSuite() {
	container, err := redis_container.Run(context.Background(), "docker.io/redis:7")
	require.NoError(suite.T(), err)

	suite.redisContainer = container

	redisEndpoint, err := suite.redisContainer.Endpoint(context.Background(), "")
	require.NoError(suite.T(), err)

	suite.redisClient = redis.NewClient(
		&redis.Options{Addr: redisEndpoint},
	)
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
