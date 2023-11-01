package gamelog_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
)

func (s *GameLogRepositoryTestSuite) TestCreate() {
	userID := createPerson(s.client, s.T())

	err := s.r.Create(context.Background(), domain.GameLog{
		User:     &domain.User{UserID: userID},
		GameType: "yut",
	})
	s.NoError(err)
}
