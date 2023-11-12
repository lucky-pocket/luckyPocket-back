package user_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/user/repository"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/client"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	client   *ent.Client
	r        domain.UserRepository
	deferred []func()
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	c, closeFunc, err := integration.CreateTestEntClient()
	if err != nil {
		s.T().Fatal(err)
	}
	s.deferred = append(s.deferred, closeFunc)

	s.client = c

	if err = client.Migrate(context.Background(), s.client); err != nil {
		s.T().Fatal(err)
	}

	s.r = repository.NewUserRepository(s.client)
}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	for _, f := range s.deferred {
		f()
	}
}

func (s *UserRepositoryTestSuite) TearDownTest() {
	ctx := context.Background()

	_, _ = s.client.Notice.Delete().Exec(ctx)
	_, _ = s.client.Pocket.Delete().Exec(ctx)
	_, _ = s.client.User.Delete().Exec(ctx)
	_, _ = s.client.GameLog.Delete().Exec(ctx)
}

func ptr[T any](i T) *T {
	return &i
}
