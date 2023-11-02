package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestFindByID() {
	info := domain.GAuthUser{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: constant.GenderFemale,
		Role:   constant.TypeStudent,
	}

	user, err := s.r.Create(context.Background(), info)
	s.NoError(err)

	s.Run("found", func() {
		found, err := s.r.FindByID(context.Background(), user.UserID)

		s.NoError(err)
		s.Equal(user, found)
	})

	s.Run("not found", func() {
		found, err := s.r.FindByID(context.Background(), user.UserID+1)

		s.NoError(err)
		s.Nil(found)
	})
}
