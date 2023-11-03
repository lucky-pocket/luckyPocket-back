package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/pocket"
)

func (r *pocketRepository) Create(ctx context.Context, pocket *domain.Pocket) error {
	return r.getClient(ctx).Pocket.
		Create().
		SetCoins(pocket.Coins).
		SetContent(pocket.Content).
		SetIsPublic(pocket.IsPublic).
		SetReceiverID(pocket.ReceiverID).
		SetSenderID(pocket.SenderID).
		Exec(ctx)
}

func (r *pocketRepository) CreateReveal(ctx context.Context, userID uint64, pocketID uint64) error {
	return r.getClient(ctx).Pocket.
		Update().
		Where(pocket.ID(pocketID)).
		AddRevealerIDs(userID).
		Exec(ctx)
}
