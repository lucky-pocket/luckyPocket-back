package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/pkg/errors"
)

func (uc *userUseCase) GetUserDetail(ctx context.Context) (*output.UserInfo, error) {
	userInfo := auth.MustExtract(ctx)

	user, err := uc.UserRepository.FindByID(ctx, userInfo.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	if user == nil {
		return nil, errors.Wrap(err, "user not found")
	}

	return mapper.ToUserInfo(*user), nil
}
