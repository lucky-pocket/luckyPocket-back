package gamelog_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	repository "github.com/lucky-pocket/luckyPocket-back/internal/app/game/repository/gamelog"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/client"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/suite"
)

type GameLogRepositoryTestSuite struct {
	suite.Suite
	client   *ent.Client
	r        domain.GameLogRepository
	deferred []func()
}

func TestGameLogRepository(t *testing.T) {
	suite.Run(t, new(GameLogRepositoryTestSuite))
}

func (s *GameLogRepositoryTestSuite) SetupSuite() {
	c, closeFunc, err := integration.CreateTestEntClient()
	if err != nil {
		s.T().Fatal(err)
	}
	s.deferred = append(s.deferred, closeFunc)

	s.client = c

	if err = client.Migrate(context.Background(), s.client); err != nil {
		s.T().Fatal(err)
	}

	s.r = repository.NewGameLogRepository(s.client)
}

func (s *GameLogRepositoryTestSuite) TearDownSuite() {
	for _, f := range s.deferred {
		f()
	}
}

func (s *GameLogRepositoryTestSuite) TearDownTest() {
	ctx := context.Background()

	_, _ = s.client.Notice.Delete().Exec(ctx)
	_, _ = s.client.Pocket.Delete().Exec(ctx)
	_, _ = s.client.User.Delete().Exec(ctx)
	_, _ = s.client.GameLog.Delete().Exec(ctx)
}
