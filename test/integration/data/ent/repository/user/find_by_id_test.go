package user_test

import (
	"context"

	"github.com/onee-only/gauth-go"
)

func (s *UserRepositoryTestSuite) TestFindByID() {
	info := gauth.UserInfo{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: gauth.GenderFemale,
		Role:   gauth.RoleTeacher,
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
