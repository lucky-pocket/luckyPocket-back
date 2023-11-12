package batch

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"go.uber.org/zap"
)

type NoticeSenderDeps struct {
	NoticeRepository domain.NoticeRepository
	NoticePool       domain.NoticePool
	Logger           *zap.Logger
}

type noticeSender struct {
	*NoticeSenderDeps
	batchSize int
}

func NewNoticeSender(deps *NoticeSenderDeps) Processor {
	return &noticeSender{deps, 100}
}

func (n *noticeSender) Do() {
	defer n.Logger.Sync()

	ctx := context.Background()
	start := time.Now()

	notices, err := n.NoticePool.Take(ctx, n.batchSize)
	if err != nil {
		n.Logger.Error("notice pool access failed",
			zap.Error(err),
		)
		return
	}

	if len(notices) > 0 {
		err = n.NoticeRepository.CreateBulk(ctx, notices)
		if err != nil {
			n.Logger.Error("bulk creation failed",
				zap.Error(err),
			)
			return
		}

		n.Logger.Info("bulk creation success",
			zap.Duration("took", time.Since(start)),
			zap.Int("count", len(notices)),
		)
		return
	}

	n.Logger.Info("notice not found")
}
