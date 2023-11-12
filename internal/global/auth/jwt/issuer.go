package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
)

type Issuer interface {
	IssueAccess(userInfo auth.Info) (string, time.Time)
	IssueRefresh(userInfo auth.Info) (string, time.Time)
	Issue(userInfo auth.Info, expiresIn time.Duration) (string, time.Time)
}

type issuer struct {
	secret        []byte
	signingMethod jwt.SigningMethod

	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewIssuer(secret []byte) Issuer {
	return &issuer{
		secret:        secret,
		signingMethod: jwt.GetSigningMethod(constant.JwtSigningMethod),

		accessTTL:  constant.JwtAccessTTL,
		refreshTTL: constant.JwtRefreshTTL,
	}
}

// IssueAccess issues token with pre-defined accessTTL.
func (i *issuer) IssueAccess(userInfo auth.Info) (string, time.Time) {
	return i.Issue(userInfo, i.accessTTL)
}

// IssueRefresh issues token with pre-defined refreshTTL.
func (i *issuer) IssueRefresh(userInfo auth.Info) (string, time.Time) {
	return i.Issue(userInfo, i.refreshTTL)
}

// Issue issues jwt token with given userInfo and expiresIn.
func (i *issuer) Issue(userInfo auth.Info, expiresIn time.Duration) (string, time.Time) {
	expiresAt := time.Now().Add(expiresIn)
	claims := Token{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Info: userInfo,
	}

	token, _ := jwt.NewWithClaims(i.signingMethod, claims).SignedString(i.secret)
	return token, expiresAt
}
