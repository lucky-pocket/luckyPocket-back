package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/user"
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

func (r *userRepository) FindByNameContains(ctx context.Context, name string) ([]*domain.User, error) {
	entities, err := r.getClient(ctx).User.
		Query().
		Where(user.NameContains(name)).
		Order(user.ByGrade(sql.OrderAsc()), user.ByClass(sql.OrderAsc())).
		All(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, 0, len(entities))
	for _, entity := range entities {
		users = append(users, mapper.ToUserDomain(entity))
	}

	return users, nil
}
