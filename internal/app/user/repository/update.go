package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/user"
)

func (r *userRepository) UpdateCoin(ctx context.Context, userID uint64, coin int) error {
	return r.getClient(ctx).User.Update().
		SetCoins(coin).
		Where(user.ID(userID)).Exec(ctx)
}
