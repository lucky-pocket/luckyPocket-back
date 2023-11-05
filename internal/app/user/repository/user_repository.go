package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	ent_tx "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/tx"
)

type userRepository struct{ client *ent.Client }

func NewUserRepository(client *ent.Client) domain.UserRepository {
	return &userRepository{client}
}

func (r *userRepository) NewTx() tx.Tx { return ent_tx.New(r.client) }

func (r *userRepository) getClient(ctx context.Context) *ent.Client {
	tx, err := ent_tx.FromContext(ctx)
	if err != nil {
		return r.client
	}
	return tx.Client()
}
