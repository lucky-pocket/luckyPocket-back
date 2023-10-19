package domain

import (
	"context"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

type GameUseCase interface {
	PlayYut(ctx context.Context, input *input.FreeInput) (*output.YutOutput, error)
	GetTicketInfo(ctx context.Context) (*output.TicketOutput, error)
}
