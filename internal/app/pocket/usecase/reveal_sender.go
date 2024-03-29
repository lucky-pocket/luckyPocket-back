package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/event"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/pkg/errors"
)

func (uc *pocketUseCase) RevealSender(ctx context.Context, input *input.PocketIDInput) (*output.UserInfo, error) {
	userInfo := auth.MustExtract(ctx)
	var sender *domain.User

	err := uc.TxManager.WithTx(ctx, func(ctx context.Context) error {
		pocket, err := uc.PocketRepository.FindByID(ctx, input.PocketID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		if pocket == nil {
			return status.NewError(http.StatusNotFound, "user not found")
		}

		if pocket.ReceiverID != userInfo.UserID && !pocket.IsPublic {
			return status.NewError(http.StatusForbidden, "you have no permission to reveal this pocket")
		}

		coins, err := uc.UserRepository.CountCoinsByUserID(ctx, userInfo.UserID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		exists, err := uc.PocketRepository.RevealExists(ctx, userInfo.UserID, pocket.PocketID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		if exists {
			return status.NewError(http.StatusConflict, "reveal exists")
		}

		if coins < constant.CostRevealSender {
			return status.NewError(http.StatusForbidden, "you don't have enough coins")
		}

		if err := uc.PocketRepository.CreateReveal(ctx, userInfo.UserID, pocket.PocketID); err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		err = uc.UserRepository.UpdateCoin(ctx, userInfo.UserID, coins-constant.CostRevealSender)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		sender, err = uc.UserRepository.FindByID(ctx, pocket.SenderID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		notice := domain.Notice{
			UserID:    pocket.SenderID,
			PocketID:  pocket.PocketID,
			Type:      constant.NoticeTypeRevealed,
			Checked:   false,
			CreatedAt: time.Now(),
		}

		if err := uc.EventManager.Publish(ctx, string(event.TopicRevealCreated), &notice); err != nil {
			return errors.Wrap(err, "event publishing failed")
		}

		return nil
	},
		uc.PocketRepository.(tx.Transactor).NewTx(),
		uc.UserRepository.(tx.Transactor).NewTx(),
	)
	if err != nil {
		return nil, err
	}

	return mapper.ToUserInfo(*sender), nil
}
