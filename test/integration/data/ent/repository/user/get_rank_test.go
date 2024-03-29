package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *UserRepositoryTestSuite) TestGetRank() {
	infos := []domain.GAuthUser{
		{
			Email:  "1",
			Name:   ptr("aef"),
			Gender: constant.GenderFemale,
			Role:   constant.TypeTeacher,
		},
		{
			Email:    "2",
			Name:     ptr("aefas"),
			Gender:   constant.GenderFemale,
			Role:     constant.TypeStudent,
			Grade:    ptr(1),
			ClassNum: ptr(3),
			Num:      ptr(2),
		},
		{
			Email:    "3",
			Name:     ptr("aefas"),
			Gender:   constant.GenderFemale,
			Role:     constant.TypeStudent,
			Grade:    ptr(1),
			ClassNum: ptr(2),
			Num:      ptr(2),
		},
	}

	for idx, info := range infos {
		user, err := s.r.Create(context.Background(), info)
		if idx == 1 {
			_, err = s.client.Pocket.Create().
				SetCoins(0).
				SetContent("").
				SetIsPublic(false).
				SetSenderID(user.UserID - 1).
				SetReceiverID(user.UserID).
				Save(context.Background())
			s.NoError(err)
		}
		s.NoError(err)
	}

	err := s.r.UpdateCoin(context.Background(), 3, 50)
	s.NoError(err)

	name := "ae"
	rank, err := s.r.RankNonStudents(context.Background(), constant.SortTypeCoins, nil, &name)
	if s.NoError(err) && s.Len(rank, 3) {
		s.Equal("aef", rank[0].Name)
	}

	grade := 1
	rank, err = s.r.RankStudents(context.Background(), constant.SortTypePocket, nil, &grade, nil)
	if s.NoError(err) && s.Len(rank, 2) {
		s.Equal("aefas", rank[0].Name)
	}
}
