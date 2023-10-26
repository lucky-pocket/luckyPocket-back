package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/pkg/errors"
)

func (uc *userUseCase) CountCoins(ctx context.Context) (*output.CoinOutput, error) {
	userInfo := auth.MustExtract(ctx)

	coins, err := uc.UserRepository.CountCoinsByUserID(ctx, userInfo.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	return mapper.ToCoinOutput(coins), nil
}
