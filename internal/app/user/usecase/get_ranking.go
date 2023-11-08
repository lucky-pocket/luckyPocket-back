package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/pkg/errors"
)

func (uc *userUseCase) GetRanking(ctx context.Context, input *input.RankQueryInput) (*output.RankOutput, error) {
	var (
		users []output.RankElem
		err   error
	)

	if input.UserType != nil && *input.UserType == constant.TypeStudent {
		users, err = uc.UserRepository.RankStudents(ctx, input.SortType, input.Name, input.Grade, input.Class)
	} else {
		users, err = uc.UserRepository.RankNonStudents(ctx, input.SortType, input.UserType, input.Name)
	}

	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	return mapper.ToRankOutput(users), nil
}
