package notice_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/notice/repository"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/suite"
)

type NoticePoolTestSuite struct {
	suite.Suite
	client   *redis.Client
	p        domain.NoticePool
	deferred []func()
}

func TestNoticePool(t *testing.T) {
	suite.Run(t, new(NoticePoolTestSuite))
}

func (s *NoticePoolTestSuite) SetupSuite() {
	c, closeFunc, err := integration.CreateTestRedisClient()
	if err != nil {
		s.T().Fatal(err)
	}
	s.deferred = append(s.deferred, closeFunc)

	s.client = c

	s.p = repository.NewNoticePool(s.client)
}

func (s *NoticePoolTestSuite) TearDownSuite() {
	for _, f := range s.deferred {
		f()
	}
}

func (s *NoticePoolTestSuite) TearDownTest() {
	ctx := context.Background()

	s.client.FlushAll(ctx)
}
