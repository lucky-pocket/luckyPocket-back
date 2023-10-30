package pocket_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/pocket/repository"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/client"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent"
	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PocketRepositoryTestSuite struct {
	suite.Suite
	client   *ent.Client
	r        domain.PocketRepository
	deferred []func()
}

func TestPocketRepository(t *testing.T) {
	suite.Run(t, new(PocketRepositoryTestSuite))
}

func (s *PocketRepositoryTestSuite) SetupSuite() {
	c, closeFunc, err := integration.CreateTestClient()
	if err != nil {
		s.T().Fatal(err)
	}
	s.deferred = append(s.deferred, closeFunc)

	s.client = c

	if err = client.Migrate(context.Background(), s.client); err != nil {
		s.T().Fatal(err)
	}

	s.r = repository.NewPocketRepository(s.client)
}

func (s *PocketRepositoryTestSuite) TearDownSuite() {
	for _, f := range s.deferred {
		f()
	}
}

func (s *PocketRepositoryTestSuite) TearDownTest() {
	ctx := context.Background()

	_, _ = s.client.Notice.Delete().Exec(ctx)
	_, _ = s.client.Pocket.Delete().Exec(ctx)
	_, _ = s.client.User.Delete().Exec(ctx)
}

func createTwoPeople(c *ent.Client, t *testing.T) (_, _ uint64) {
	user1, err := c.User.Create().
		SetEmail("1").
		SetName("hi").
		SetCoins(0).
		SetGender(constant.GenderFemale).
		SetRole(constant.RoleMember).
		SetUserType(constant.TypeTeacher).
		Save(context.Background())

	if !assert.Nil(t, err) {
		return 0, 0
	}

	user2, err := c.User.Create().
		SetEmail("2").
		SetName("hei").
		SetCoins(0).
		SetGender(constant.GenderFemale).
		SetRole(constant.RoleMember).
		SetUserType(constant.TypeTeacher).
		Save(context.Background())

	if !assert.Nil(t, err) {
		return 0, 0
	}

	return user1.ID, user2.ID
}
