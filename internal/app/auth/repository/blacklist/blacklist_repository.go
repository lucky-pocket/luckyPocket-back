package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	redis_tx "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/redis/tx"
	"github.com/redis/go-redis/v9"
)

const prefix = "blacklist"

func buildKey(key string) string { return fmt.Sprintf("%s:%s", prefix, key) }

type blackListRepository struct{ client *redis.Client }

func NewBlackListRepository(client *redis.Client) domain.BlackListRepository {
	return &blackListRepository{client}
}

func (r *blackListRepository) NewTx() tx.Tx { return redis_tx.New(r.client) }

func (r *blackListRepository) getClient(ctx context.Context) redis.Cmdable {
	tx, err := redis_tx.FromContext(ctx)
	if err != nil {
		return r.client
	}
	return tx
}

func (r *blackListRepository) Exists(ctx context.Context, refreshToken string) (bool, error) {
	cmd := r.getClient(ctx).Exists(ctx, buildKey(refreshToken))
	if err := cmd.Err(); err != nil {
		return false, err
	}

	return cmd.Val() != 0, nil
}

func (r *blackListRepository) Save(ctx context.Context, refreshToken string, expiresAt time.Time) error {
	cmd := r.getClient(ctx).Set(
		ctx, buildKey(refreshToken),
		nil, time.Until(expiresAt),
	)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}
