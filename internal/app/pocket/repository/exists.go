package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/pocket"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/user"
)

func (r *pocketRepository) RevealExists(ctx context.Context, userID uint64, pocketID uint64) (bool, error) {
	return r.getClient(ctx).Pocket.Query().
		Where(
			pocket.And(
				pocket.ID(pocketID),
				pocket.HasRevealersWith(user.ID(userID)),
			),
		).Exist(ctx)
}
