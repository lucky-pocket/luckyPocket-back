package filter

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth/jwt"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
)

var (
	errInvalidAuthToken = status.NewError(http.StatusUnauthorized, "invalid auth token")
)

type AuthFilter struct {
	jwtParser jwt.Parser
}

func NewAuthFilter(jwtParser jwt.Parser) *AuthFilter {
	return &AuthFilter{jwtParser: jwtParser}
}

func (f *AuthFilter) WithRequired(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		info, err := f.extractInfo(c)
		if err != nil {
			if required {
				c.Error(errInvalidAuthToken)
				c.Abort()
				return
			}
		} else {
			c.Request = c.Request.WithContext(auth.Inject(c, info))
		}

		c.Next()
	}
}

func (f *AuthFilter) extractInfo(c *gin.Context) (auth.Info, error) {
	authorization, ok := c.Request.Header["Authorization"]
	if !ok || len(authorization) != 1 {
		return auth.Info{}, errInvalidAuthToken
	}

	bearerToken, found := strings.CutPrefix(authorization[0], "Bearer ")
	if !found {
		return auth.Info{}, errInvalidAuthToken
	}

	token, err := f.jwtParser.Parse(bearerToken)
	if err != nil {
		return auth.Info{}, errInvalidAuthToken
	}

	return token.Info, nil
}
