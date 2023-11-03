package notice_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *NoticeRepositoryTestSuite) TestCreateBulk() {
	userID, pocketID := s.createUserAndPocket()

	notices := []*domain.Notice{
		{
			UserID:   userID,
			PocketID: pocketID,
			Type:     constant.NoticeTypeReceived,
			Checked:  false,
		},
		{
			UserID:   userID,
			PocketID: pocketID,
			Type:     constant.NoticeTypeRevealed,
			Checked:  false,
		},
		{
			UserID:   userID,
			PocketID: pocketID,
			Type:     constant.NoticeTypeReceived,
			Checked:  true,
		},
	}

	err := s.r.CreateBulk(context.Background(), notices)

	if !s.NoError(err) {
		return
	}

	count, err := s.client.Notice.
		Query().
		Count(context.Background())

	if s.NoError(err) {
		s.Equal(len(notices), count)
	}
}
