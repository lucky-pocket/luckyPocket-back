package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestFindByID() {
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

	s.Run("found", func() {
		found, err := s.r.FindByID(context.Background(), user.UserID)

		s.NoError(err)
		s.Equal(user, *found)
	})

	s.Run("not found", func() {
		found, err := s.r.FindByID(context.Background(), user.UserID+1)

		s.NoError(err)
		s.Nil(found)
	})
}
