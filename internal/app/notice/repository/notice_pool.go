package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	redis_tx "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/redis/tx"
	"github.com/pkg/errors"

	"github.com/redis/go-redis/v9"
)

const poolKey = "noticePool"

type jsonNotice struct {
	NoticeID  uint64    `json:"noticeID"`
	UserID    uint64    `json:"userID"`
	PocketID  uint64    `json:"pocketID"`
	Type      string    `json:"type"`
	Checked   bool      `json:"checked"`
	CreatedAt time.Time `json:"createdAt"`
}

type noticePool struct{ client *redis.Client }

func NewNoticePool(client *redis.Client) domain.NoticePool {
	return &noticePool{client}
}

func (n *noticePool) NewTx() tx.Tx { return redis_tx.New(n.client) }

func (r *noticePool) getClient(ctx context.Context) redis.Cmdable {
	tx, err := redis_tx.FromContext(ctx)
	if err != nil {
		return r.client
	}
	return tx
}

func (n *noticePool) Put(ctx context.Context, notice *domain.Notice) error {
	noticeJSON := jsonNotice{
		NoticeID:  notice.NoticeID,
		UserID:    notice.UserID,
		PocketID:  notice.PocketID,
		Type:      string(notice.Type),
		Checked:   notice.Checked,
		CreatedAt: notice.CreatedAt,
	}

	bytes, err := json.Marshal(noticeJSON)
	if err != nil {
		return errors.Wrap(err, "error marshaling notice")
	}

	cmd := n.getClient(ctx).SAdd(ctx, poolKey, bytes)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (n *noticePool) Take(ctx context.Context, count int) ([]*domain.Notice, error) {
	cmd := n.getClient(ctx).SRandMemberN(ctx, poolKey, int64(count))
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	notices := make([]*domain.Notice, 0, count)
	lines := cmd.Val()
	for _, line := range lines {
		var notice domain.Notice
		if err := json.Unmarshal([]byte(line), &notice); err != nil {
			return nil, errors.Wrap(err, "error unmarshaling notice")
		}
		notices = append(notices, &notice)
	}
	return notices, nil
}
