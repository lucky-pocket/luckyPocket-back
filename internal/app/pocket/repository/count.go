package repository

import (
	"context"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/pocket"
	"time"
)

func (r *pocketRepository) CountBySenderIdAndReceiverId(ctx context.Context, senderID uint64, receiverID uint64) (int, error) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.Add(24 * time.Hour)

	return r.getClient(ctx).Pocket.Query().
		Where(
			pocket.And(
				pocket.SenderID(senderID),
				pocket.ReceiverID(receiverID),
				pocket.CreatedAtGTE(start),
				pocket.CreatedAtLTE(end),
			),
		).Count(ctx)
}
