package user_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/onee-only/gauth-go"
)

func (s *UserRepositoryTestSuite) TestGetRank() {
	infos := []gauth.UserInfo{
		{
			Email:  "1",
			Name:   ptr("aef"),
			Gender: gauth.GenderFemale,
			Role:   gauth.RoleTeacher,
		},
		{
			Email:    "2",
			Name:     ptr("hi?"),
			Gender:   gauth.GenderFemale,
			Role:     gauth.RoleStudent,
			Grade:    ptr(1),
			ClassNum: ptr(3),
			Num:      ptr(2),
		},
		{
			Email:    "3",
			Name:     ptr("aef"),
			Gender:   gauth.GenderFemale,
			Role:     gauth.RoleStudent,
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
	rank, err := s.r.RankNonStudents(context.Background(), constant.SortTypeCoins, &name)
	if s.NoError(err) && s.Len(rank, 1) {
		s.Equal("aef", rank[0].Name)
	}

	grade := 1
	rank, err = s.r.RankStudents(context.Background(), constant.SortTypePocket, nil, &grade, nil)
	if s.NoError(err) && s.Len(rank, 2) {
		s.Equal("hi?", rank[0].Name)
	}
}
