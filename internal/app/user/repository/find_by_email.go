package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/user"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/mapper"
)

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.getClient(ctx).User.
		Query().
		Where(user.Email(email)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ToUserDomain(user), nil
}
