package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestExists() {
	info := domain.GAuthUser{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: constant.GenderFemale,
		Role:   constant.TypeStudent,
	}

	user, err := s.r.Create(context.Background(), info)
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

func (s *UserRepositoryTestSuite) TestExistsByEmail() {
	info := domain.GAuthUser{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: constant.GenderFemale,
		Role:   constant.TypeStudent,
	}

	_, err := s.r.Create(context.Background(), info)
	s.NoError(err)

	s.Run("found", func() {
		exists, err := s.r.ExistsByEmail(context.Background(), info.Email)

		s.NoError(err)
		s.True(exists)
	})

	s.Run("not found", func() {
		exists, err := s.r.ExistsByEmail(context.Background(), info.Email+"l")

		s.NoError(err)
		s.False(exists)
	})
}
