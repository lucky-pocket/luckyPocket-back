package domain

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

type AuthUseCase interface {
	Login(ctx context.Context, input *input.CodeInput) (*output.TokenOutput, error)
	Logout(ctx context.Context, input *input.RefreshInput) error
	RefreshToken(ctx context.Context, input *input.RefreshInput) (*output.TokenOutput, error)
}

type AuthRepository interface {
	ExistByRefreshToken(ctx context.Context, refreshToken *string) (bool, error)
}
