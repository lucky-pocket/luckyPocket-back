package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/pkg/errors"
)

type Deps struct {
	UserRepository   domain.UserRepository
	PocketRepository domain.PocketRepository
}

type pocketUseCase struct{ *Deps }

func NewPocketUseCase(deps *Deps) domain.PocketUseCase {
	return &pocketUseCase{deps}
}

func (p *pocketUseCase) GetUserPockets(ctx context.Context, input *input.Input) (*output.PocketListOutput, error) {
	user, err := p.UserRepository.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	if user == nil {
		return nil, errors.Wrap(err, "user not found")
	}

	pockets, err := p.PocketRepository.FindListByUserID(ctx, input.UserID, input.Offset, input.Limit)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	return mapper.ToPocketListOutput(pockets), nil
}

func (p *pocketUseCase) SendPocket(ctx context.Context, input *input.PocketInput) error {
	userInfo := auth.MustExtract(ctx)

	receiver, err := p.UserRepository.FindByName(ctx, input.Receiver)
	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	if receiver == nil {
		return errors.Wrap(err, "user not found")
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

	if err := p.PocketRepository.Create(ctx, &pocket); err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	if input.IsPublic {
		if err := p.UserRepository.SaveReveal(ctx, userInfo.UserID, pocket.PocketID); err != nil {
			return errors.Wrap(err, "unexpected db error")
		}
	}

	return nil
}

func (p *pocketUseCase) RevealSender(ctx context.Context, input *input.PocketIDInput) error {
	userInfo := auth.MustExtract(ctx)

	currentUser, err := p.UserRepository.FindByID(ctx, userInfo.UserID)
	pocket, err := p.PocketRepository.FindByID(ctx, input.PocketID)

	if currentUser == nil && pocket == nil {
		return errors.Wrap(err, "user not found")
	}

	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	if pocket.Receiver != currentUser {
		return errors.Wrap(err, "not allowed permission this pocket")
	}

	exists, err := p.UserRepository.ExistsReveal(ctx, currentUser.UserID, pocket.PocketID)

	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	if exists {
		return errors.Wrap(err, "exists reveal")
	}

	err = p.UserRepository.SaveReveal(ctx, currentUser.UserID, pocket.PocketID)

	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	coin := currentUser.Coins
	err = p.UserRepository.UpdateCoin(ctx, currentUser.UserID, coin-1)

	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	return nil
}

func (p *pocketUseCase) GetPocketDetail(ctx context.Context, input *input.PocketIDInput) (*output.PocketOutput, error) {
	pocket, err := p.PocketRepository.FindByID(ctx, input.PocketID)

	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	if pocket == nil {
		return nil, errors.Wrap(err, "pocket not found")
	}

	isPublic, err2 := p.UserRepository.ExistsReveal(ctx, pocket.Receiver.UserID, pocket.PocketID)

	if err2 != nil {
		return nil, errors.Wrap(err2, "unexpected db error")
	}

	return mapper.ToPocketOutput(pocket, isPublic), nil
}

func (p *pocketUseCase) SetVisibility(ctx context.Context, input *input.VisibilityInput) error {
	userInfo := auth.MustExtract(ctx)

	pocket, err := p.PocketRepository.FindByID(ctx, input.PocketID)

	if pocket == nil {
		return errors.Wrap(err, "not found error")
	}

	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}

	if pocket.Receiver.UserID != userInfo.UserID {
		return errors.Wrap(err, "not allowed permission this pocket")
	}

	err = p.PocketRepository.UpdateVisibility(ctx, pocket.PocketID, input.Visible)

	if err != nil {
		return errors.Wrap(err, "unexpected db error")
	}
	return nil
}
