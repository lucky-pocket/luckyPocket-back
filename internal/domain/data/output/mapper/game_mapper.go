package mapper

import (
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToYutOutput(marked bool, yutPieces [3]bool, coinsEarned int, result string) *output.YutOutput {
	return &output.YutOutput{
		Result: &output.YutResultElem{
			Marked:    marked,
			YutPieces: yutPieces,
		},
		CoinsEarned: coinsEarned,
		Output:      result,
	}
}

func ToFreeTicketOutput(count int, refillAt time.Time) *output.TicketOutput {
	return &output.TicketOutput{
		RefillAt:    refillAt,
		TicketCount: count,
	}
}

func ToPlayCountOutput(count int) *output.PlayCountOutput {
	return &output.PlayCountOutput{
		Count: count,
	}
}
