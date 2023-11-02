package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestCreate() {
	grade, class, num := 1, 1, 1

	user := domain.GAuthUser{
		Email:    "l",
		Name:     ptr("aef"),
		Gender:   constant.GenderFemale,
		Role:     constant.TypeStudent,
		Grade:    &grade,
		ClassNum: &class,
		Num:      &num,
	}

	_, err := s.r.Create(context.Background(), user)
	s.NoError(err)
}
