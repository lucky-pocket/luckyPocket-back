package notice_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *NoticeRepositoryTestSuite) TestFindAllByUserID() {
	userID, pocketID := s.createUserAndPocket()

	err := s.client.Notice.Create().
		SetChecked(false).
		SetPocketID(pocketID).
		SetType(constant.NoticeTypeReceived).
		SetUserID(userID).Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	notices, err := s.r.FindAllByUserID(context.Background(), userID)

	if s.NoError(err) {
		s.Len(notices, 1)
	}
}

func (s *NoticeRepositoryTestSuite) TestFindByID() {
	userID, pocketID := s.createUserAndPocket()

	notice, err := s.client.Notice.Create().
		SetChecked(false).
		SetPocketID(pocketID).
		SetType(constant.NoticeTypeReceived).
		SetUserID(userID).Save(context.Background())

	if !s.NoError(err) {
		return
	}

	s.Run("found", func() {
		n, err := s.r.FindByID(context.Background(), notice.ID)

		if s.NoError(err) {
			s.Equal(notice.ID, n.NoticeID)
		}
	})
	s.Run("not found", func() {
		n, err := s.r.FindByID(context.Background(), notice.ID+1)

		if s.NoError(err) {
			s.Nil(n)
		}
	})
}
