package notice_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/notice/repository"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/client"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/suite"
)

type NoticeRepositoryTestSuite struct {
	suite.Suite
	client   *ent.Client
	r        domain.NoticeRepository
	deferred []func()
}

func TestNoticeRepository(t *testing.T) {
	suite.Run(t, new(NoticeRepositoryTestSuite))
}

func (s *NoticeRepositoryTestSuite) SetupSuite() {
	c, closeFunc, err := integration.CreateTestEntClient()
	if err != nil {
		s.T().Fatal(err)
	}
	s.deferred = append(s.deferred, closeFunc)

	s.client = c

	if err = client.Migrate(context.Background(), s.client); err != nil {
		s.T().Fatal(err)
	}

	s.r = repository.NewNoticeRepository(s.client)
}

func (s *NoticeRepositoryTestSuite) TearDownSuite() {
	for _, f := range s.deferred {
		f()
	}
}

func (s *NoticeRepositoryTestSuite) TearDownTest() {
	ctx := context.Background()

	_, _ = s.client.Notice.Delete().Exec(ctx)
	_, _ = s.client.Pocket.Delete().Exec(ctx)
	_, _ = s.client.User.Delete().Exec(ctx)
	_, _ = s.client.GameLog.Delete().Exec(ctx)
}

func (s *NoticeRepositoryTestSuite) createUserAndPocket() (_, _ uint64) {
	user, err := s.client.User.Create().
		SetEmail("221312").
		SetName("hi").
		SetCoins(0).
		SetGender(constant.GenderFemale).
		SetRole(constant.RoleMember).
		SetUserType(constant.TypeTeacher).
		Save(context.Background())

	if !s.Nil(err) {
		return
	}

	pocket, err := s.client.Pocket.Create().
		SetCoins(0).
		SetContent("").
		SetIsPublic(false).
		SetReceiverID(user.ID).
		SetSenderID(user.ID).
		Save(context.Background())

	if !s.Nil(err) {
		return
	}

	return user.ID, pocket.ID
}
