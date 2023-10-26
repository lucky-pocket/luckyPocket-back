package usecase

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/pkg/errors"
)

func (uc *userUseCase) GetMyDetail(ctx context.Context) (*output.MyDetailOutput, error) {
	userInfo := auth.MustExtract(ctx)

	user, err := uc.UserRepository.FindByID(ctx, userInfo.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	if user == nil {
		return nil, status.NewError(http.StatusNotFound, "user not found")
	}

	noticeExists, err := uc.NoticeRepository.ExistsByUserID(ctx, user.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	return mapper.ToMyDetailOutput(*user, noticeExists), nil
}
