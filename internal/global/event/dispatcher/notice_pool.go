package dispatcher

import (
	"context"
	"fmt"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/event"
	"github.com/pkg/errors"
)

type NoticePoolDumperDeps struct {
	NoticePool domain.NoticePool
}

type noticePoolDumper struct{ *NoticePoolDumperDeps }

func NewNoticePoolDumper(deps *NoticePoolDumperDeps) event.Dispatcher {
	return &noticePoolDumper{deps}
}

func (d *noticePoolDumper) Dispatch(ctx context.Context, _ string, payload any) error {
	notice, ok := payload.(*domain.Notice)
	if !ok {
		return errors.New(fmt.Sprintf("payload assertion failed. expected: %T actual: %T", domain.Notice{}, payload))
	}

	if err := d.NoticePool.Put(ctx, notice); err != nil {
		return errors.Wrap(err, "putting notice to pool failed")
	}

	return nil
}
