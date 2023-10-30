package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/user"
)

func (r *userRepository) Exists(ctx context.Context, userID uint64) (bool, error) {
	return r.getClient(ctx).User.
		Query().
		Where(user.ID(userID)).
		Exist(ctx)
}
