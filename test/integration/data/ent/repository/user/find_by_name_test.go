package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestFindByNameContains() {
	info := domain.GAuthUser{
		Email:  "l",
		Name:   ptr("aef"),
		Gender: constant.GenderFemale,
		Role:   constant.RoleMember,
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
