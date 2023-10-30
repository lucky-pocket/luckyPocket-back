package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestCountCoinsByUserID() {
	user := domain.User{
		UserID:   1,
		Name:     "aef",
		Coins:    50,
		Gender:   constant.GenderFemale,
		UserType: constant.TypeGraduate,
	}

	err := s.r.Create(context.Background(), &user)
	s.NoError(err)

	err = s.r.UpdateCoin(context.Background(), user.UserID, 40)
	s.NoError(err)

	s.Run("found", func() {
		coins, err := s.r.CountCoinsByUserID(context.Background(), user.UserID)

		s.NoError(err)
		s.Equal(40, coins)
	})

	s.Run("not found", func() {
		coins, err := s.r.CountCoinsByUserID(context.Background(), user.UserID+1)

		s.NoError(err)
		s.Zero(coins)
	})
}
