package batch

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
)

type NoticeSenderDeps struct {
	NoticeRepository domain.NoticeRepository
	NoticePool       domain.NoticePool
}

type noticeSender struct {
	*NoticeSenderDeps
	batchSize int
}

func NewNoticeSender(deps *NoticeSenderDeps) Processor {
	return &noticeSender{deps, 100}
}

func (n *noticeSender) Do() {
	ctx := context.Background()

	notices, err := n.NoticePool.Take(ctx, n.batchSize)
	if err != nil {
		// TODO: log the error.
		return
	}

	err = n.NoticeRepository.CreateBulk(ctx, notices)
	if err != nil {
		// TODO: log the error.
		return
	}

	// TODO: log the result.
}
