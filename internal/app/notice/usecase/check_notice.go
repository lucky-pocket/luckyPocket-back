package usecase

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/pkg/errors"
)

func (uc *noticeUseCase) CheckNotice(ctx context.Context, noticeID uint64) error {
	info := auth.MustExtract(ctx)

	notice, err := uc.NoticeRepository.FindByID(ctx, noticeID)
	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	if notice == nil {
		return status.NewError(http.StatusNotFound, "notice not found")
	}

	if notice.UserID != info.UserID {
		return status.NewError(http.StatusForbidden, "you are not the owner")
	}

	if err := uc.NoticeRepository.SetChecked(ctx, notice.NoticeID, true); err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	return nil
}
