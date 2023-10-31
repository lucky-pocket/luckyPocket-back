package user_test

import (
	"context"

	"github.com/onee-only/gauth-go"
)

func (s *UserRepositoryTestSuite) TestFindByNameContains() {
	info := gauth.UserInfo{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: gauth.GenderFemale,
		Role:   gauth.RoleTeacher,
	}

	_, err := s.r.Create(context.Background(), info)
	s.NoError(err)

	s.Run("found", func() {
		list, err := s.r.FindByNameContains(context.Background(), "a")

		s.NoError(err)
		s.Len(list, 1)
	})

	s.Run("not found", func() {
		list, err := s.r.FindByNameContains(context.Background(), "efefe")

		s.NoError(err)
		s.Len(list, 0)
	})
}
