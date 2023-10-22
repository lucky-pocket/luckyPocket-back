package domain

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

type Pocket struct {
	PocketID uint64
	Receiver *User
	Sender   *User
	Content  string
	Coins    int
}

func (p Pocket) IsEmpty() bool {
	return p.Coins == 0
}

type PocketUseCase interface {
	SendPocket(ctx context.Context, input *input.PocketInput) error
	RevealSender(ctx context.Context, input *input.UserIDInput) error
	GetUserPockets(ctx context.Context, input *input.UserIDInput) (*output.PocketListOutput, error)
	GetPocketDetail(ctx context.Context, input *input.PocketIDInput) (*output.PocketOutput, error)
	SetVisibility(ctx context.Context, input *input.VisibilityInput) error
}

type PocketRepository interface {
	Create(ctx context.Context, pocket *Pocket) error
	FindByID(ctx context.Context, pocketID uint64) (pocket *Pocket, err error)
	FindPocketListByUserID(ctx context.Context, userID uint64, offset, limit int) (pocketList []*Pocket, err error)
}
