package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/mapper"
)

func (r *userRepository) FindByID(ctx context.Context, userID uint64) (*domain.User, error) {
	user, err := r.getClient(ctx).User.Get(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ToUserDomain(user), nil
}
