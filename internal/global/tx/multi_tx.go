package tx

import (
	"context"
	"errors"
)

type multiTx struct {
	txs map[any]Tx
}

func mergeTx(txs ...Tx) *multiTx {
	newTx := &multiTx{
		txs: make(map[any]Tx, len(txs)),
	}

	for _, tx := range txs {
		newTx.txs[tx.Transactor()] = tx
	}

	return newTx
}

func (mtx *multiTx) begin(ctx context.Context) (_ context.Context, err error) {
	for transactor, tx := range mtx.txs {
		trans, e := tx.Begin()
		if e == nil {
			ctx = context.WithValue(ctx, transactor, trans)
		}
		err = errors.Join(err, e)
	}
	return ctx, err
}

func (mtx *multiTx) commit() (err error) {
	for _, tx := range mtx.txs {
		err = errors.Join(err, tx.Commit())
	}
	return err
}

func (mtx *multiTx) rollback() (err error) {
	for _, tx := range mtx.txs {
		err = errors.Join(err, tx.Rollback())
	}
	return err
}
