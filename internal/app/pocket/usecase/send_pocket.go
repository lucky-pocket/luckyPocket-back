package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/pkg/errors"
)

func (uc *pocketUseCase) SendPocket(ctx context.Context, input *input.PocketInput) error {
	userInfo := auth.MustExtract(ctx)

	return uc.TxManager.WithTx(ctx, func(ctx context.Context) error {
		receiver, err := uc.UserRepository.FindByID(ctx, input.ReceiverID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		if receiver == nil {
			return errors.New("user not found")
		}

		coins, err := uc.UserRepository.CountCoinsByUserID(ctx, userInfo.UserID)
		if err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		pocket := domain.Pocket{
			Receiver: receiver,
			Sender: &domain.User{
				UserID: userInfo.UserID,
				Role:   userInfo.Role,
			},
			Content:  input.Message,
			Coins:    input.Coins,
			IsPublic: input.IsPublic,
		}

		price := pocket.Coins + constant.CostSendPocket
		if coins < price {
			return errors.New("you don't have enough money")
		}

		if err := uc.PocketRepository.Create(ctx, &pocket); err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		if input.IsPublic {
			if err := uc.PocketRepository.CreateReveal(ctx, userInfo.UserID, pocket.PocketID); err != nil {
				return errors.Wrap(err, "unexpected db error")
			}
		}

		if err := uc.UserRepository.UpdateCoin(ctx, userInfo.UserID, coins-price); err != nil {
			return errors.Wrap(err, "unexpected db error")
		}

		return nil
	}, uc.PocketRepository.(tx.Transactor).NewTx())
}
