package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestCreate() {
	grade, class, num := 1, 1, 1

	user := domain.User{
		Name:     "aef",
		Coins:    0,
		Gender:   constant.GenderFemale,
		UserType: constant.TypeStudent,
		Grade:    &grade,
		Class:    &class,
		Number:   &num,
	}

	err := s.r.Create(context.Background(), &user)

	s.NoError(err)
}
