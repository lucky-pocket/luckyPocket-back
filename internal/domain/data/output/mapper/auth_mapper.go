package mapper

import (
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToTokenOutput(access, refresh string) *output.TokenOutput {
	return &output.TokenOutput{
		Access: output.TokenElem{
			Token:     access,
			ExpiresAt: time.Now().Add(constant.JwtAccessTTL),
		},
		Refresh: output.TokenElem{
			Token:     refresh,
			ExpiresAt: time.Now().Add(constant.JwtRefreshTTL),
		},
	}
}
