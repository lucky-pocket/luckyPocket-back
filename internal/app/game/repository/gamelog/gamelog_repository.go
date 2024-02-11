package repository

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/gamelog"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/user"
	ent_tx "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/tx"
)

type gameLogRepository struct{ client *ent.Client }

func NewGameLogRepository(client *ent.Client) domain.GameLogRepository {
	return &gameLogRepository{client}
}

func (r *gameLogRepository) NewTx() tx.Tx { return ent_tx.New(r.client) }

func (r *gameLogRepository) getClient(ctx context.Context) *ent.Client {
	tx, err := ent_tx.FromContext(ctx)
	if err != nil {
		return r.client
	}
	return tx.Client()
}

func (r *gameLogRepository) Create(ctx context.Context, log domain.GameLog) error {
	return r.getClient(ctx).GameLog.Create().
		SetUserID(log.User.UserID).
		SetGameType(log.GameType).
		Exec(ctx)
}

func (r *gameLogRepository) CountByUserID(ctx context.Context, userID uint64) (int, error) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.Add(24 * time.Hour)

	return r.getClient(ctx).GameLog.Query().
		Where(
			gamelog.And(
				gamelog.HasUserWith(user.ID(userID)),
				gamelog.TimestampGTE(start),
				gamelog.TimestampLTE(end),
			),
		).Count(ctx)
}
