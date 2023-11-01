package domain

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/onee-only/gauth-go"
)

type AuthUseCase interface {
	Login(ctx context.Context, input *input.CodeInput) (*output.TokenOutput, error)
	Logout(ctx context.Context, input *input.RefreshInput) error
	RefreshToken(ctx context.Context, input *input.RefreshInput) (*output.TokenOutput, error)
}

type GAuthClient interface {
	IssueToken(code string) (access, refresh string, err error)
	GetUserInfo(accessToken string) (*gauth.UserInfo, error)
}

type BlackListRepository interface {
	Exists(ctx context.Context, refreshToken string) (bool, error)
	Save(ctx context.Context, refreshToken string, expiresAt time.Time) error
}
