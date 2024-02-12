package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/event"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/pkg/errors"
)

func (uc *pocketUseCase) SendPocket(ctx context.Context, input *input.PocketInput) error {
	userInfo := auth.MustExtract(ctx)
	if userInfo.UserID == input.ReceiverID {
		return status.NewError(http.StatusForbidden, "자기 자신은 영원한 친구입니다.")
	}

	count, err := uc.PocketRepository.CountBySenderIdAndReceiverId(ctx, userInfo.UserID, input.ReceiverID)
	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	if count >= constant.SameSendLimit {
		return status.NewError(http.StatusTeapot, "You cannot send more than five times a day to the same user.")
	}

	return uc.TxManager.WithTx(ctx, func(ctx context.Context) error {
		receiver, err := uc.UserRepository.FindByID(ctx, input.ReceiverID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		if receiver == nil {
			return status.NewError(http.StatusNotFound, "user not found")
		}

		coins, err := uc.UserRepository.CountCoinsByUserID(ctx, userInfo.UserID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		pocket := domain.Pocket{
			ReceiverID: receiver.UserID,
			SenderID:   userInfo.UserID,
			Content:    input.Message,
			Coins:      input.Coins,
			IsPublic:   true,
		}

		price := pocket.Coins + constant.CostSendPocket
		if coins < price {
			return status.NewError(http.StatusForbidden, "you don't have enough money")
		}

		if err := uc.PocketRepository.Create(ctx, &pocket); err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		if input.IsPublic {
			if err := uc.PocketRepository.CreateReveal(ctx, receiver.UserID, pocket.PocketID); err != nil {
				return errors.Wrap(err, "unexpected db error")
			}
		}

		if err := uc.UserRepository.UpdateCoin(ctx, userInfo.UserID, coins-price); err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		if err := uc.UserRepository.UpdateCoin(ctx, receiver.UserID, receiver.Coins+input.Coins); err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		notice := domain.Notice{
			UserID:    receiver.UserID,
			PocketID:  pocket.PocketID,
			Type:      constant.NoticeTypeReceived,
			Checked:   false,
			CreatedAt: time.Now(),
		}

		if err := uc.EventManager.Publish(ctx, string(event.TopicPocketReceived), &notice); err != nil {
			return errors.Wrap(err, "event publishing failed")
		}

		return nil
	},
		uc.PocketRepository.(tx.Transactor).NewTx(),
		uc.UserRepository.(tx.Transactor).NewTx(),
	)
}
