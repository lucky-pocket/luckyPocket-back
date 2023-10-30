package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/pocket"
)

func (r *pocketRepository) UpdateVisibility(ctx context.Context, pocketID uint64, visible bool) error {
	return r.getClient(ctx).Pocket.
		Update().
		Where(pocket.ID(pocketID)).
		SetIsPublic(visible).Exec(ctx)
}
