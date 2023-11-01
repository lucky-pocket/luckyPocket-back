package pocket_test

import "context"

func (s *PocketRepositoryTestSuite) TestUpdateVisibility() {
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

	err = s.r.UpdateVisibility(context.Background(), pocket.ID, true)
	s.NoError(err)
}
