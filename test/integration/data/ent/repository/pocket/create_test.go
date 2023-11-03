package pocket_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
)

func (s *PocketRepositoryTestSuite) TestCreate() {
	user1ID, user2ID := createTwoPeople(s.client, s.T())

	err := s.r.Create(context.Background(), &domain.Pocket{
		ReceiverID: user1ID,
		SenderID:   user2ID,
		Content:    "haha go pocket",
		Coins:      0,
		IsPublic:   false,
	})
	s.NoError(err)
}

func (s *PocketRepositoryTestSuite) TestCreateReveal() {
	user1ID, user2ID := createTwoPeople(s.client, s.T())

	pocket, err := s.client.Pocket.Create().
		SetCoins(0).
		SetContent("haha go pocket").
		SetIsPublic(false).
		SetReceiverID(user1ID).
		SetSenderID(user2ID).
		Save(context.Background())

	if !s.NoError(err) {
		return
	}

	err = s.r.CreateReveal(context.Background(), user1ID, pocket.ID)
	s.NoError(err)
}
