package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
)

func (s *UserRepositoryTestSuite) TestCountCoinsByUserID() {
	info := domain.GAuthUser{
		Email:    "l",
		Name:     ptr("aef"),
		Gender:   constant.GenderFemale,
		Role:     constant.TypeStudent,
		Grade:    ptr(1),
		ClassNum: ptr(1),
		Num:      ptr(1),
	}

	user, err := s.r.Create(context.Background(), info)
	s.NoError(err)

	err = s.r.UpdateCoin(context.Background(), user.UserID, 40)
	s.NoError(err)

	// Don't want to test it. 
	// Becasue SQLite doesn't support FOR UPDATE/SHARE statement.
	// If you want to test it, then make the integration test run on MySQL.
	return

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
