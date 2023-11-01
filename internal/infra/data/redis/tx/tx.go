package tx

import (
	"context"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// cannot implement tx we want with redis. just creating stub for it.
type redisTx struct{}

type redisTxKey struct{}

func New(client *redis.Client) *redisTx {
	return &redisTx{}
}

func FromContext(ctx context.Context) (tx redis.Pipeliner, err error) {
	tx, ok := ctx.Value(redisTxKey{}).(redis.Pipeliner)
	if !ok {
		return nil, errors.New("tx not found")
	}
	return tx, nil
}

func (redisTx) Transactor() any {
	return redisTxKey{}
}

func (rtx *redisTx) Begin() (any, error) {
	return nil, nil
}

func (rtx *redisTx) Commit() error {
	return nil
}

func (rtx *redisTx) Rollback() error {
	return nil
}
