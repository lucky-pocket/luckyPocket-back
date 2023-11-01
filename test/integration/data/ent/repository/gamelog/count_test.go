package gamelog_test

import (
	"context"
	"time"
)

func (s *GameLogRepositoryTestSuite) TestCountByUserID() {
	userID := createPerson(s.client, s.T())

	err := s.client.GameLog.Create().
		SetGameType("yut").
		SetTimestamp(time.Now()).
		SetUserID(userID).Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	err = s.client.GameLog.Create().
		SetGameType("yut").
		SetTimestamp(time.Now().Truncate(24 * time.Hour).Add(24*time.Hour + 1)).
		SetUserID(userID).Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	count, err := s.r.CountByUserID(context.Background(), userID)
	if s.NoError(err) {
		s.Equal(1, count)
	}
}
