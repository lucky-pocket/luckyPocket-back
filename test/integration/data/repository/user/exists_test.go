package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestExists() {
	user := domain.User{
		UserID:   1,
		Name:     "aef",
		Coins:    50,
		Gender:   constant.GenderFemale,
		UserType: constant.TypeTeacher,
	}

	err := s.r.Create(context.Background(), &user)
	s.NoError(err)

	s.Run("found", func() {
		exists, err := s.r.Exists(context.Background(), user.UserID)

		s.NoError(err)
		s.True(exists)
	})

	s.Run("not found", func() {
		exists, err := s.r.Exists(context.Background(), user.UserID+1)

		s.NoError(err)
		s.False(exists)
	})
}
