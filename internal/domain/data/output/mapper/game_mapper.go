package mapper

import (
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToYutOutput(marked bool, yutPieces [3]bool, coinsEarned int) *output.YutOutput {
	return &output.YutOutput{
		Result: output.YutResultElem{
			Marked:    marked,
			YutPieces: yutPieces,
		},
		CoinsEarned: coinsEarned,
	}
}

func ToFreeTicketOutput(count int, refillAt time.Time) *output.TicketOutput {
	return &output.TicketOutput{
		RefillAt:    refillAt,
		TicketCount: count,
	}
}
