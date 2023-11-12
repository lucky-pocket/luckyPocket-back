package jwt_test

import (
	"testing"

	jwtgo "github.com/golang-jwt/jwt/v5"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth/jwt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestIssue(t *testing.T) {
	secret := []byte("hi new secret")
	info := auth.Info{
		UserID: 1,
		Role:   constant.RoleAdmin,
	}

	issuer := jwt.NewIssuer(secret)

	validate := func(token string) error {
		_, err := jwtgo.ParseWithClaims(token, &jwt.Token{}, func(t *jwtgo.Token) (interface{}, error) {
			return secret, nil
		})

		return errors.Wrap(err, "error parsing jwt token")
	}

	t.Run("access", func(t *testing.T) {
		accessToken, _ := issuer.IssueAccess(info)

		err := validate(accessToken)

		assert.NoError(t, err)
	})

	t.Run("refresh", func(t *testing.T) {
		refreshToken, _ := issuer.IssueRefresh(info)

		err := validate(refreshToken)

		assert.NoError(t, err)
	})
}
