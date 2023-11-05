package notice_test

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *NoticePoolTestSuite) TestPut() {
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

	cmd := s.client.SPop(context.Background(), "noticePool")
	if !s.NoError(cmd.Err()) {
		return
	}
}
