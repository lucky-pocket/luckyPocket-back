package domain

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

type Notice struct {
	NoticeID  uint64
	User      *User
	UserID    uint64
	Pocket    *Pocket
	PocketID  uint64
	Type      constant.NoticeType
	Checked   bool
	CreatedAt time.Time
}

type NoticeUseCase interface {
	GetNotice(ctx context.Context) (*output.NoticeListOutput, error)
	CheckNotice(ctx context.Context, noticeID uint64) error
	CheckAllNotices(ctx context.Context) error
}

type NoticeRepository interface {
	CreateBulk(ctx context.Context, notices []*Notice) error
	FindAllByUserID(ctx context.Context, userID uint64) ([]*Notice, error)
	FindByID(ctx context.Context, noticeID uint64) (*Notice, error)
	ExistsByUserID(ctx context.Context, userID uint64) (bool, error)
	SetChecked(ctx context.Context, noticeID uint64, checked bool) error
	SetCheckedByUserID(ctx context.Context, userID uint64, checked bool) error
}

type NoticePool interface {
	Put(ctx context.Context, notice *Notice) error
	Take(ctx context.Context, count int) ([]*Notice, error)
}
