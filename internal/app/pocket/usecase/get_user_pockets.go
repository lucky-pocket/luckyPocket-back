package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/pkg/errors"
)

func (uc *pocketUseCase) GetUserPockets(ctx context.Context, input *input.PocketQueryInput) (*output.PocketListOutput, error) {
	pockets, err := uc.PocketRepository.FindListByUserID(ctx, input.UserID, input.Offset, input.Limit)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	return mapper.ToPocketListOutput(pockets), nil
}
