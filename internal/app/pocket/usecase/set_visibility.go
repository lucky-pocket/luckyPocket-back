package usecase

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/pkg/errors"
)

func (uc *pocketUseCase) SetVisibility(ctx context.Context, input *input.VisibilityInput) error {
	userInfo := auth.MustExtract(ctx)

	return uc.TxManager.WithTx(ctx, func(ctx context.Context) error {
		pocket, err := uc.PocketRepository.FindByID(ctx, input.PocketID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		if pocket == nil {
			return status.NewError(http.StatusNotFound, "pocket not found")
		}

		if pocket.Receiver.UserID != userInfo.UserID {
			return status.NewError(http.StatusForbidden, "you have no permission to set visibility of this pocket")
		}

		if err := uc.PocketRepository.UpdateVisibility(ctx, pocket.PocketID, input.Visible); err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		return nil
	}, uc.PocketRepository.(tx.Transactor).NewTx())
}
