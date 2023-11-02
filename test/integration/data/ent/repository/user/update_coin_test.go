package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestUpdateCoin() {
	info := domain.GAuthUser{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: constant.GenderFemale,
		Role:   constant.TypeStudent,
	}

	user, err := s.r.Create(context.Background(), info)
	s.NoError(err)

	err = s.r.UpdateCoin(context.Background(), user.UserID, 30)
	s.NoError(err)
}
