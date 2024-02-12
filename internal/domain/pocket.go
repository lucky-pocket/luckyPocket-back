package domain

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

type Pocket struct {
	PocketID   uint64
	Receiver   *User
	ReceiverID uint64
	Sender     *User
	SenderID   uint64
	Content    string
	Coins      int
	IsPublic   bool
}

func (p Pocket) IsEmpty() bool {
	return p.Coins == 0
}

type PocketUseCase interface {
	SendPocket(ctx context.Context, input *input.PocketInput) error
	RevealSender(ctx context.Context, input *input.PocketIDInput) (*output.UserInfo, error)
	GetUserPockets(ctx context.Context, input *input.PocketQueryInput) (*output.PocketListOutput, error)
	GetPocketDetail(ctx context.Context, input *input.PocketIDInput) (*output.PocketOutput, error)
	SetVisibility(ctx context.Context, input *input.VisibilityInput) error
}

type PocketRepository interface {
	Create(ctx context.Context, pocket *Pocket) error
	FindByID(ctx context.Context, pocketID uint64) (*Pocket, error)
	FindListByUserID(ctx context.Context, userID uint64, offset, limit int) ([]*Pocket, error)
	FillSenderNameOnRevealed(ctx context.Context, pockets []*Pocket, receiverID, userID uint64) error
	UpdateVisibility(ctx context.Context, pocketID uint64, visible bool) error
	CreateReveal(ctx context.Context, userID uint64, pocketID uint64) error
	RevealExists(ctx context.Context, userID uint64, pocketID uint64) (bool, error)
	CountBySenderIdAndReceiverId(ctx context.Context, senderID uint64, receiverID uint64) (int, error)
}
