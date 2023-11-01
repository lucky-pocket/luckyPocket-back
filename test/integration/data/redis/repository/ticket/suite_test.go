package ticket_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"

	repository "github.com/lucky-pocket/luckyPocket-back/internal/app/game/repository/ticket"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/suite"
)

type TicketRepositoryTestSuite struct {
	suite.Suite
	client   *redis.Client
	r        domain.TicketRepository
	deferred []func()
}

func TestTicketRepository(t *testing.T) {
	suite.Run(t, new(TicketRepositoryTestSuite))
}

func (s *TicketRepositoryTestSuite) SetupSuite() {
	c, closeFunc, err := integration.CreateTestRedisClient()
	if err != nil {
		s.T().Fatal(err)
	}
	s.deferred = append(s.deferred, closeFunc)

	s.client = c

	s.r = repository.NewTicketRepository(s.client)
}

func (s *TicketRepositoryTestSuite) TearDownSuite() {
	for _, f := range s.deferred {
		f()
	}
}

func (s *TicketRepositoryTestSuite) TearDownTest() {
	ctx := context.Background()

	s.client.FlushAll(ctx)
}
