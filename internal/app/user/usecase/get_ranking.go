package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/pkg/errors"
)

func (uc *userUseCase) GetRanking(ctx context.Context, input *input.UserInput) (*output.RankOutput, error) {
	var (
		users []output.RankElem
		err   error
	)

	if input.UserType == constant.TypeStudent {
		users, err = uc.UserRepository.FindStudentsWithFilter(ctx, input.SortType, input.Name, input.Class, input.Grade)
	} else {
		users, err = uc.UserRepository.FindNonStudentWithFilter(ctx, input.SortType, input.Name)
	}

	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	return mapper.ToRankOutput(users), nil
}
