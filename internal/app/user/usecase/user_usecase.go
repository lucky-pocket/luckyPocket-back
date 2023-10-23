package usecase

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
)

type Deps struct {
	UserRepository   domain.UserRepository
	PocketRepository domain.PocketRepository
}

type userUseCase struct{ *Deps }

func NewUserUseCase(deps *Deps) domain.UserUseCase {
	return &userUseCase{deps}
}

func (uc *userUseCase) GetMyDetail(ctx context.Context) (*output.MyDetailOutput, error) {
	userInfo := auth.MustExtract(ctx)

	user, err := uc.UserRepository.FindByID(ctx, userInfo.UserID)

	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	if user == nil {
		return nil, errors.Wrap(err, "user notfound")
	}

	// TODO : Notice Service 작성시 hasNewNotification 에 대한 로직 추가.
	return mapper.ToMyDetailOutput(*user, true), nil
}

func (uc *userUseCase) CountCoins(ctx context.Context) (*output.CoinOutput, error) {
	userInfo := auth.MustExtract(ctx)

	coins, err := uc.UserRepository.CountCoinsByUserID(ctx, userInfo.UserID)

	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	return mapper.ToCoinOutput(coins), nil
}

func (uc *userUseCase) GetUserDetail(ctx context.Context) (*output.UserInfo, error) {
	userInfo := auth.MustExtract(ctx)

	user, err := uc.UserRepository.FindByID(ctx, userInfo.UserID)

	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	if user == nil {
		return nil, errors.Wrap(err, "user notfound")
	}

	return mapper.ToUserInfo(*user), nil
}

func (uc *userUseCase) GetRanking(ctx context.Context, input *input.UserInput) (*output.RankOutput, error) {
	var (
		users []output.RankElem
		err   error
	)

	switch input.UserType {
	case constant.TypeStudent:
		users, err = uc.UserRepository.FindStudentsWithFilter(ctx, input.SortType, input.Name, input.Class, input.Grade)
	default:
		users, err = uc.UserRepository.FindNonStudentWithFilter(ctx, input.SortType, input.Name)
	}

	if err != nil {
		return nil, errors.Wrap(err, "unexpected error occurred")
	}

	return mapper.ToRankOutput(users), nil
}
