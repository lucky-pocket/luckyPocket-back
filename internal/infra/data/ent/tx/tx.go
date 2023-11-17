package tx

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/pkg/errors"
)

type entTX struct {
	c  *ent.Client
	tx *ent.Tx
}

type entTxKey struct{}

func New(client *ent.Client) *entTX {
	return &entTX{c: client}
}

func FromContext(ctx context.Context) (tx *ent.Tx, err error) {
	tx, ok := ctx.Value(entTxKey{}).(*ent.Tx)
	if !ok {
		return nil, errors.New("tx not found")
	}
	return tx, nil
}

func (entTX) Transactor() any {
	return entTxKey{}
}

func (etx *entTX) Begin() (any, error) {
	// might have to inject ctx.
	tranx, err := etx.c.Tx(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "ent transaciton begin failed")
	}
	etx.tx = tranx
	return tranx, nil
}

func (etx *entTX) Commit() error {
	return etx.tx.Commit()
}

func (etx *entTX) Rollback() error {
	return etx.tx.Rollback()
}
