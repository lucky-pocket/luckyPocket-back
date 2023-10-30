package mapper

import (
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToTokenOutput(access, refresh string, accessExpiredAt, refreshExpiredAT time.Time) *output.TokenOutput {
	return &output.TokenOutput{
		Access: output.TokenElem{
			Token:     access,
			ExpiresAt: accessExpiredAt,
		},
		Refresh: output.TokenElem{
			Token:     refresh,
			ExpiresAt: refreshExpiredAT,
		},
	}
}
