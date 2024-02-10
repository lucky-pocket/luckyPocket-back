package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/pkg/errors"
)

func (uc *pocketUseCase) GetUserPockets(ctx context.Context, input *input.PocketQueryInput) (*output.PocketListOutput, error) {
	info, _ := auth.Extract(ctx)

	pockets, err := uc.PocketRepository.FindListByUserID(ctx, input.UserID, input.Offset, input.Limit)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	if info != nil {
		err = uc.PocketRepository.FillSenderNameOnRevealed(ctx, pockets, input.UserID, info.UserID)
		if err != nil {
			return nil, errors.Wrap(err, "unexpected db error")
		}
	}

	return mapper.ToPocketListOutput(pockets), nil
}
