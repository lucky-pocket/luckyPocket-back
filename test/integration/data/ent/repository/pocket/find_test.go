package pocket_test

import "context"

func (s *PocketRepositoryTestSuite) TestFindByID() {
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

	s.Run("found", func() {
		pocket, err := s.r.FindByID(context.Background(), entity.ID)
		s.NoError(err)
		s.Equal(entity.ID, pocket.PocketID)
	})
	s.Run("not found", func() {
		pocket, err := s.r.FindByID(context.Background(), entity.ID+1)
		s.NoError(err)
		s.Nil(pocket)
	})
}

func (s *PocketRepositoryTestSuite) TestFindListByUserID() {
	user1ID, user2ID := createTwoPeople(s.client, s.T())

	for i := 0; i < 5; i++ {
		_, err := s.client.Pocket.Create().
			SetCoins(0).
			SetContent("haha go pocket").
			SetIsPublic(false).
			SetReceiverID(user1ID).
			SetSenderID(user2ID).
			Save(context.Background())

		if !s.NoError(err) {
			return
		}
	}

	pockets, err := s.r.FindListByUserID(context.Background(), user1ID, 0, 10)
	s.NoError(err)
	s.Len(pockets, 5)
}
