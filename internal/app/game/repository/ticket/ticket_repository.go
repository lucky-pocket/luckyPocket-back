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

const prefix = "ticket"

func buildKey(key string) string { return fmt.Sprintf("%s:%s", prefix, key) }

type ticketRepository struct{ client *redis.Client }

func NewTicketRepository(client *redis.Client) domain.TicketRepository {
	return &ticketRepository{client}
}

func (r *ticketRepository) NewTx() tx.Tx { return redis_tx.New(r.client) }

func (r *ticketRepository) getClient(ctx context.Context) redis.Cmdable {
	tx, err := redis_tx.FromContext(ctx)
	if err != nil {
		return r.client
	}
	return tx
}

func (r *ticketRepository) GetRefillAt(ctx context.Context, userID uint64) (time.Time, error) {
	cmd := r.getClient(ctx).ExpireTime(ctx, buildKey(fmt.Sprint(userID)))
	if err := cmd.Err(); err != nil {
		return time.Time{}, err
	}

	if cmd.Val() == -2 {
		return time.Now(), nil
	}

	expiresAt := time.Unix(int64(cmd.Val().Seconds()), 0)

	return expiresAt, nil
}
func (r *ticketRepository) CountByUserID(ctx context.Context, userID uint64) (int, error) {
	cmd := r.getClient(ctx).Get(ctx, buildKey(fmt.Sprint(userID)))
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return -1, err
	}

	return cmd.Int()
}
func (r *ticketRepository) Increase(ctx context.Context, userID uint64) error {
	key := buildKey(fmt.Sprint(userID))
	client := r.getClient(ctx)

	cmd := client.Incr(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}

	now := time.Now()
	midNight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())

	err := client.ExpireNX(ctx, key, time.Until(midNight)).Err()
	if err != nil {
		return err
	}

	return nil
}
