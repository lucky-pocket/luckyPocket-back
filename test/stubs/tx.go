package stubs

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
)

type stubTx struct{}

func (stx *stubTx) Transactor() any {
	return ""
}

func (stx *stubTx) Begin() (any, error) {
	return nil, nil
}

func (stx *stubTx) Commit() error {
	return nil
}

func (stx *stubTx) Rollback() error {
	return nil
}

type stubTransactor struct{}

func (*stubTransactor) NewTx() tx.Tx {
	return &stubTx{}
}
