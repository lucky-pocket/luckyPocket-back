package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
)

type Token struct {
	jwt.RegisteredClaims
	auth.Info
}
