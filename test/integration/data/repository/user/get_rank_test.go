package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestGetRank() {
	ptr := func(i int) *int {
		return &i
	}

	users := []domain.User{
		{
			Name:     "aef",
			Coins:    0,
			Gender:   constant.GenderFemale,
			UserType: constant.TypeTeacher,
			Role:     constant.RoleMember,
		},
		{
			Name:     "hi?",
			Coins:    0,
			Gender:   constant.GenderFemale,
			UserType: constant.TypeStudent,
			Grade:    ptr(1),
			Class:    ptr(3),
			Number:   ptr(2),
		},
		{
			Name:     "aef",
			Coins:    0,
			Gender:   constant.GenderFemale,
			UserType: constant.TypeStudent,
			Grade:    ptr(1),
			Class:    ptr(2),
			Number:   ptr(2),
		},
	}

	for _, user := range users {
		err := s.r.Create(context.Background(), &user)
		s.NoError(err)
	}

	err := s.r.UpdateCoin(context.Background(), 3, 50)
	s.NoError(err)

	_, err = s.client.Pocket.Create().
		SetCoins(0).
		SetContent("").
		SetIsPublic(false).
		SetSenderID(1).
		SetReceiverID(2).
		Save(context.Background())
	s.NoError(err)

	name := "ae"
	rank, err := s.r.FindNonStudentWithFilter(context.Background(), constant.SortTypeCoins, &name)
	if s.NoError(err) && s.Len(rank, 1) {
		s.Equal(uint64(1), rank[0].UserID)
	}

	grade := 1
	rank, err = s.r.FindStudentsWithFilter(context.Background(), constant.SortTypePocket, nil, &grade, nil)
	if s.NoError(err) && s.Len(rank, 2) {
		s.Equal(uint64(2), rank[0].UserID)
	}
}
