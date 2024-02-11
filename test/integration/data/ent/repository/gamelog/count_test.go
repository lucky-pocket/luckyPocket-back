package gamelog_test

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *GameLogRepositoryTestSuite) TestCountByUserID() {
	user, err := s.client.User.Create().
		SetEmail("12321").
		SetName("hi").
		SetCoins(0).
		SetGender(constant.GenderFemale).
		SetRole(constant.RoleMember).
		SetUserType(constant.TypeTeacher).
		Save(context.Background())

	if !s.Nil(err) {
		return
	}

	now := time.Now()

	err = s.client.GameLog.Create().
		SetGameType("yut").
		SetTimestamp(now).
		SetUserID(user.ID).Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	err = s.client.GameLog.Create().
		SetGameType("yut").
		SetTimestamp(time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).Add(time.Hour)).
		SetUserID(user.ID).Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	count, err := s.r.CountByUserID(context.Background(), user.ID)
	if s.NoError(err) {
		s.Equal(1, count)
	}
}
