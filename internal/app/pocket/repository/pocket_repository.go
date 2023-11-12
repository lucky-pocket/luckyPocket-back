package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	ent_tx "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/tx"
)

type pocketRepository struct{ client *ent.Client }

func NewPocketRepository(client *ent.Client) domain.PocketRepository {
	return &pocketRepository{client}
}

func (r *pocketRepository) NewTx() tx.Tx { return ent_tx.New(r.client) }

func (r *pocketRepository) getClient(ctx context.Context) *ent.Client {
	tx, err := ent_tx.FromContext(ctx)
	if err != nil {
		return r.client
	}
	return tx.Client()
}
