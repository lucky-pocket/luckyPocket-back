package domain

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

type GameLog struct {
	User      *User
	UserID    uint64
	TimeStamp time.Time
	GameType  string
}

type GameUseCase interface {
	PlayYut(ctx context.Context, input *input.FreeInput) (*output.YutOutput, error)
	GetTicketInfo(ctx context.Context) (*output.TicketOutput, error)
	CountPlays(ctx context.Context) (*output.PlayCountOutput, error)
}

type TicketRepository interface {
	GetRefillAt(ctx context.Context, userID uint64) (time.Time, error)
	CountByUserID(ctx context.Context, userID uint64) (int, error)
	Increase(ctx context.Context, userID uint64) error
}

type GameLogRepository interface {
	Create(ctx context.Context, log GameLog) error
	CountByUserID(ctx context.Context, userID uint64) (int, error)
}
