package user_test

import (
	"context"

	"github.com/onee-only/gauth-go"
)

func (s *UserRepositoryTestSuite) TestCountCoinsByUserID() {
	info := gauth.UserInfo{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: gauth.GenderFemale,
		Role:   gauth.RoleTeacher,
	}

	user, err := s.r.Create(context.Background(), info)
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