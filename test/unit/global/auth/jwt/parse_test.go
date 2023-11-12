package jwt_test

import (
	"testing"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth/jwt"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	secret := []byte("hi new secret")
	info := auth.Info{
		UserID: 1,
		Role:   constant.RoleAdmin,
	}

	issuer := jwt.NewIssuer(secret)
	parser := jwt.NewParser(secret)

	t.Run("success", func(t *testing.T) {
		token, _ := issuer.Issue(info, time.Minute)

		tokn, err := parser.Parse(token)

		assert.NoError(t, err)
		assert.Equal(t, info.UserID, tokn.UserID)
	})

	t.Run("expired", func(t *testing.T) {
		token, _ := issuer.Issue(info, -time.Minute)

		_, err := parser.Parse(token)

		assert.Error(t, err)
	})
}
