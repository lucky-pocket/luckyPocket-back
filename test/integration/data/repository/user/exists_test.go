package user_test

import (
	"context"

	"github.com/onee-only/gauth-go"
)

func (s *UserRepositoryTestSuite) TestExists() {
	info := gauth.UserInfo{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: gauth.GenderFemale,
		Role:   gauth.RoleTeacher,
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
	info := gauth.UserInfo{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: gauth.GenderFemale,
		Role:   gauth.RoleTeacher,
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
