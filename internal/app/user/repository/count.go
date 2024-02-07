package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/user"
)

func (r *userRepository) CountCoinsByUserID(ctx context.Context, userID uint64) (int, error) {
	user, err := r.getClient(ctx).User.
		Query().
		Where(user.ID(userID)).
		Select(user.FieldCoins).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return 0, nil
		}
		return 0, err
	}

	return user.Coins, nil
}
