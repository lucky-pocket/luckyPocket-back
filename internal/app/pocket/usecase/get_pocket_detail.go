package usecase

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/pkg/errors"
)

func (uc *pocketUseCase) GetPocketDetail(ctx context.Context, input *input.PocketIDInput) (*output.PocketOutput, error) {
	userInfo, _ := auth.Extract(ctx)

	pocket, err := uc.PocketRepository.FindByID(ctx, input.PocketID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	if pocket == nil {
		return nil, status.NewError(http.StatusNotFound, "pocket not found")
	}

	isReceiver := userInfo != nil && pocket.Receiver.UserID == userInfo.UserID
	if !(pocket.IsPublic || isReceiver) {
		return nil, status.NewError(http.StatusForbidden, "you cannot open this pocket")
	}

	exists, err := uc.PocketRepository.RevealExists(ctx, pocket.Receiver.UserID, pocket.PocketID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	var sender *domain.User
	if exists {
		sender, err = uc.UserRepository.FindByID(ctx, pocket.Sender.UserID)
		if err != nil {
			return nil, errors.Wrap(err, "unexpected db error")
		}
	}

	return mapper.ToPocketOutput(pocket, sender), nil
}
