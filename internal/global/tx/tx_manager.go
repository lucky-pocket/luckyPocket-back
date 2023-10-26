package tx

import (
	"context"

	errs "github.com/pkg/errors"
)

type Manager interface {
	WithTx(ctx context.Context, f func(ctx context.Context) error, txs ...Tx) error
}

type txManager struct{}

func NewTxManager() Manager {
	return &txManager{}
}

// WithTx makes function f to be atomic.
// You have to register specific tx.Tx implementation to this function.
func (tm *txManager) WithTx(ctx context.Context, f func(ctx context.Context) error, txs ...Tx) error {
	mtx := mergeTx(txs...)

	ctx, err := mtx.begin(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to start transaction")
	}

	if err := f(ctx); err != nil {
		if e := mtx.rollback(); e != nil {
			return errs.Wrapf(err, "transaction rollback failed: %s", e.Error())
		}
		return err
	}

	if err := mtx.commit(); err != nil {
		return errs.Wrap(err, "transaction commit failed")
	}

	return nil
}
