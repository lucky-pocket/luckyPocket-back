package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestUpdateCoin() {
	user := domain.User{
		UserID:   1,
		Name:     "aef",
		Coins:    0,
		Gender:   constant.GenderFemale,
		UserType: constant.TypeTeacher,
		Role:     constant.RoleMember,
	}

	err := s.r.Create(context.Background(), &user)
	s.NoError(err)

	err = s.r.UpdateCoin(context.Background(), user.UserID, 30)
	s.NoError(err)
}
