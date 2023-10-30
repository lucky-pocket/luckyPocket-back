package user_test

import (
	"context"

	"github.com/onee-only/gauth-go"
)

func (s *UserRepositoryTestSuite) TestFindByEmail() {
	info := gauth.UserInfo{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: gauth.GenderFemale,
		Role:   gauth.RoleTeacher,
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
