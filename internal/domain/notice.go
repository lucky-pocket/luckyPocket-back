package domain

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

type Notice struct {
	NoticeID uint64
	User     *User
	Pocket   *Pocket
	Type     constant.NoticeType
	Checked  bool
}

type NoticeUseCase interface {
	GetNotice(ctx context.Context) (*output.NoticeListOutput, error)
}

type NoticeRepository interface {
	Create(ctx context.Context, notices []*Notice) error
	FindAllByUserID(ctx context.Context, userID uint64) (*output.NoticeListOutput, error)
	FindByID(ctx context.Context, noticeID uint64) (*Notice, error)
}