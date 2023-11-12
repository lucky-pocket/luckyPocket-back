package tx

type Tx interface {
	Transactor() any
	Begin() (any, error)
	Commit() error
	Rollback() error
}

type Transactor interface {
	NewTx() Tx
}
