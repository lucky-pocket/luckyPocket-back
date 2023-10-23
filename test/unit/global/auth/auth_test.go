package auth_test

import (
	"context"
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleAdmin,
	}

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		ctx := auth.Inject(ctx, userInfo)

		uInfo, err := auth.Extract(ctx)

		assert.NoError(t, err)
		assert.Equal(t, userInfo, *uInfo)
	})

	t.Run("not found", func(t *testing.T) {
		uInfo, err := auth.Extract(ctx)

		assert.Error(t, err)
		assert.Nil(t, uInfo)
	})

	t.Run("not found with panic", func(t *testing.T) {
		assert.Panics(t, func() {
			auth.MustExtract(ctx)
		})
	})

}
