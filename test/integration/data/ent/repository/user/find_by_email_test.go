package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestFindByEmail() {
	info := domain.GAuthUser{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: constant.GenderFemale,
		Role:   constant.RoleMember,
	}

	user, err := s.r.Create(context.Background(), info)
	s.NoError(err)

	s.Run("found", func() {
		found, err := s.r.FindByEmail(context.Background(), info.Email)

		if s.NoError(err) && s.NotNil(found) {
			s.Equal(user.UserID, found.UserID)
		}
	})

	s.Run("not found", func() {
		found, err := s.r.FindByEmail(context.Background(), info.Email+"1")

		s.NoError(err)
		s.Nil(found)
	})
}
