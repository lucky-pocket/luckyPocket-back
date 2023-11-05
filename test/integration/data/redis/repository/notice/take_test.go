package notice_test

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *NoticePoolTestSuite) TestTake() {
	notice := domain.Notice{
		UserID:    1,
		PocketID:  1,
		Type:      constant.NoticeTypeReceived,
		Checked:   true,
		CreatedAt: time.Now(),
	}

	err := s.p.Put(context.Background(), &notice)
	if !s.NoError(err) {
		return
	}

	notices, err := s.p.Take(context.Background(), 2)
	if !s.NoError(err) {
		return
	}

	if s.Len(notices, 1) {
		s.True(notice.CreatedAt.Equal(notices[0].CreatedAt))
	}
}
