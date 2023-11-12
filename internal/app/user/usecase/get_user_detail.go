package usecase

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/pkg/errors"
)

func (uc *userUseCase) GetUserDetail(ctx context.Context, input *input.UserIDInput) (*output.UserInfo, error) {
	user, err := uc.UserRepository.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	if user == nil {
		return nil, status.NewError(http.StatusNotFound, "user not found")
	}

	return mapper.ToUserInfo(*user), nil
}
