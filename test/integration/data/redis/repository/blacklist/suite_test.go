package blacklist_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"

	repository "github.com/lucky-pocket/luckyPocket-back/internal/app/auth/repository/blacklist"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/suite"
)

type BlacklistRepositoryTestSuite struct {
	suite.Suite
	client   *redis.Client
	r        domain.BlackListRepository
	deferred []func()
}

func TestPocketRepository(t *testing.T) {
	suite.Run(t, new(BlacklistRepositoryTestSuite))
}

func (s *BlacklistRepositoryTestSuite) SetupSuite() {
	c, closeFunc, err := integration.CreateTestRedisClient()
	if err != nil {
		s.T().Fatal(err)
	}
	s.deferred = append(s.deferred, closeFunc)

	s.client = c

	s.r = repository.NewBlackListRepository(s.client)
}

func (s *BlacklistRepositoryTestSuite) TearDownSuite() {
	for _, f := range s.deferred {
		f()
	}
}

func (s *BlacklistRepositoryTestSuite) TearDownTest() {
	ctx := context.Background()

	s.client.FlushAll(ctx)
}
