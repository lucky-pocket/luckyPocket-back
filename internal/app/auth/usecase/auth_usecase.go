package usecase

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth/jwt"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/pkg/errors"
)

type Deps struct {
	UserRepository      domain.UserRepository
	GAuthClient         domain.GAuthClient
	BlackListRepository domain.BlackListRepository
	JwtParser           jwt.Parser
	JwtIssuer           jwt.Issuer
	TxManager           tx.Manager
}

type authUseCase struct{ *Deps }

func NewAuthUseCase(deps *Deps) domain.AuthUseCase {
	return &authUseCase{deps}
}

func (a *authUseCase) Login(ctx context.Context, input *input.CodeInput) (*output.TokenOutput, error) {
	access, _, err := a.GAuthClient.IssueToken(input.Code)
	if err != nil {
		return nil, status.NewError(http.StatusBadRequest, "code is invalid")
	}

	userInfo, err := a.GAuthClient.GetUserInfo(access)
	if err != nil {
		return nil, errors.Wrap(err, "gauth client error")
	}

	exists, err := a.UserRepository.ExistsByEmail(ctx, userInfo.Email)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error")
	}

	var user *domain.User
	if !exists {
		user, err = a.UserRepository.Create(ctx, *userInfo)
	} else {
		user, err = a.UserRepository.FindByEmail(ctx, userInfo.Email)
	}

	if err != nil {
		return nil, errors.Wrap(err, "unexpected error")
	}

	info := auth.Info{UserID: user.UserID, Role: user.Role}

	access, accessExp := a.JwtIssuer.IssueAccess(info)
	refresh, refreshExp := a.JwtIssuer.IssueRefresh(info)

	return mapper.ToTokenOutput(
		access, refresh,
		accessExp, refreshExp,
	), nil
}

func (a *authUseCase) Logout(ctx context.Context, input *input.RefreshInput) error {
	token, err := a.JwtParser.Parse(input.RefreshToken)
	if err != nil {
		return status.NewError(http.StatusUnauthorized, "token is invalid")
	}

	exist, err := a.BlackListRepository.Exists(ctx, input.RefreshToken)
	if err != nil {
		return errors.Wrap(err, "unexpected error")
	}

	if exist {
		return status.NewError(http.StatusUnauthorized, "token is blacklisted")
	}

	err = a.BlackListRepository.Save(ctx, input.RefreshToken, token.ExpiresAt.Time)
	if err != nil {
		return errors.Wrap(err, "unexpected error")
	}

	return nil
}

func (a *authUseCase) RefreshToken(ctx context.Context, input *input.RefreshInput) (*output.TokenOutput, error) {
	token, err := a.JwtParser.Parse(input.RefreshToken)
	if err != nil {
		return nil, status.NewError(http.StatusUnauthorized, "token is not valid")
	}

	exist, err := a.BlackListRepository.Exists(ctx, input.RefreshToken)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error")
	}

	if exist {
		return nil, status.NewError(http.StatusUnauthorized, "token is blacklisted")
	}

	err = a.BlackListRepository.Save(ctx, input.RefreshToken, token.ExpiresAt.Time)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error")
	}

	info := auth.Info{
		UserID: token.UserID,
		Role:   token.Role,
	}

	access, accessExp := a.JwtIssuer.IssueAccess(info)
	refresh, refreshExp := a.JwtIssuer.IssueRefresh(info)

	return mapper.ToTokenOutput(
		access, refresh,
		accessExp, refreshExp,
	), nil
}
