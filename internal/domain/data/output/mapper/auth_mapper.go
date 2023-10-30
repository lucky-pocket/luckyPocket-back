package mapper

import (
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToTokenOutput(access, refresh string, accessExp, refreshExp time.Time) *output.TokenOutput {
	return &output.TokenOutput{
		Access: output.TokenElem{
			Token:     access,
			ExpiresAt: accessExp,
		},
		Refresh: output.TokenElem{
			Token:     refresh,
			ExpiresAt: refreshExp,
		},
	}
}
