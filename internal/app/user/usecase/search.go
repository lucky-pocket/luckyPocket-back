package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/pkg/errors"
)

func (uc *userUseCase) Search(ctx context.Context, input *input.SearchInput) (*output.SearchOutput, error) {
	users, err := uc.UserRepository.FindByNameContains(ctx, input.SearchQuery)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	return mapper.ToSearchOutput(users), nil
}
