package pocket_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/pocket"
)

func (s *PocketRepositoryTestSuite) TestRevealExists() {
	user1ID, user2ID := createTwoPeople(s.client, s.T())

	pockt, err := s.client.Pocket.Create().
		SetCoins(0).
		SetContent("haha go pocket").
		SetIsPublic(false).
		SetReceiverID(user1ID).
		SetSenderID(user2ID).
		Save(context.Background())

	if !s.NoError(err) {
		return
	}

	err = s.client.Pocket.
		Update().
		Where(pocket.ID(pockt.ID)).
		AddRevealerIDs(user1ID).
		Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	exists, err := s.r.RevealExists(context.Background(), user1ID, pockt.ID)

	s.NoError(err)
	s.True(exists)
}
