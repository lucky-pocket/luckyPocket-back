package user_test

import (
	"context"

	"github.com/onee-only/gauth-go"
)

func (s *UserRepositoryTestSuite) TestCreate() {
	grade, class, num := 1, 1, 1

	user := gauth.UserInfo{
		Email:    "l",
		Name:     ptr("aef"),
		Gender:   gauth.GenderFemale,
		Role:     gauth.RoleStudent,
		Grade:    &grade,
		ClassNum: &class,
		Num:      &num,
	}

	_, err := s.r.Create(context.Background(), user)
	s.NoError(err)
}
