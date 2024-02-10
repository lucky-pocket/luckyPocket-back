package pocket_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/pocket"
)

func (s *PocketRepositoryTestSuite) TestFillSender() {
	user1ID, user2ID := createTwoPeople(s.client, s.T())

	entity, err := s.client.Pocket.Create().
		SetCoins(0).
		SetContent("haha go pocket").
		SetIsPublic(false).
		SetReceiverID(user1ID).
		SetSenderID(user2ID).
		Save(context.Background())

	if !s.NoError(err) {
		return
	}

	err = s.client.Pocket.Update().
		AddRevealerIDs(user1ID).
		Where(pocket.ID(entity.ID)).
		Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	s.Run("found", func() {
		pockets, err := s.r.FindListByUserID(context.Background(), user1ID, 0, 30)
		s.NoError(err)
		err = s.r.FillSenderNameOnRevealed(context.Background(), pockets, user1ID, user1ID)
		s.NoError(err)
		s.NotNil(pockets[0].Sender)
	})
}
