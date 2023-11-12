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

	err = s.client.GameLog.Create().
		SetGameType("yut").
		SetTimestamp(time.Now()).
		SetUserID(user.ID).Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	err = s.client.GameLog.Create().
		SetGameType("yut").
		SetTimestamp(time.Now().Truncate(24 * time.Hour).Add(24*time.Hour + 1)).
		SetUserID(user.ID).Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	count, err := s.r.CountByUserID(context.Background(), user.ID)
	if s.NoError(err) {
		s.Equal(1, count)
	}
}
