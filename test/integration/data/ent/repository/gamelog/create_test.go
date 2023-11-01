package gamelog_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *GameLogRepositoryTestSuite) TestCreate() {
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

	err = s.r.Create(context.Background(), domain.GameLog{
		User:     &domain.User{UserID: user.ID},
		GameType: "yut",
	})
	s.NoError(err)
}
