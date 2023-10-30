package user_test

import (
	"context"

	"github.com/onee-only/gauth-go"
)

func (s *UserRepositoryTestSuite) TestUpdateCoin() {
	info := gauth.UserInfo{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: gauth.GenderFemale,
		Role:   gauth.RoleTeacher,
	}

	user, err := s.r.Create(context.Background(), info)
	s.NoError(err)

	err = s.r.UpdateCoin(context.Background(), user.UserID, 30)
	s.NoError(err)
}
