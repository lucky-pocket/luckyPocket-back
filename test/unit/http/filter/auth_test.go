package filter_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth/jwt"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/stretchr/testify/assert"
)

func TestAuthFilter(t *testing.T) {
	jwtIssuer := jwt.NewIssuer([]byte("secret"))
	authFilter := filter.NewAuthFilter(jwt.NewParser([]byte("secret")))

	info := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	expiredToken, _ := jwtIssuer.Issue(info, -1*time.Minute)
	validToken, _ := jwtIssuer.Issue(info, 5*time.Minute)

	r := gin.Default()
	r.Use(filter.NewErrorFilter())
	r.GET("/true", authFilter.WithRequired(true), func(ctx *gin.Context) { auth.MustExtract(ctx.Request.Context()) })
	r.GET("/false", authFilter.WithRequired(false))

	testcases := []struct {
		desc   string
		method string
		path   string
		token  string

		statusCode int
	}{
		{
			desc:       "required and valid",
			method:     "GET",
			path:       "/true",
			token:      validToken,
			statusCode: http.StatusOK,
		},
		{
			desc:       "required and not given",
			method:     "GET",
			path:       "/true",
			token:      "",
			statusCode: http.StatusUnauthorized,
		},
		{
			desc:       "required and invalid",
			method:     "GET",
			path:       "/true",
			token:      "aefaefaekljfh",
			statusCode: http.StatusUnauthorized,
		},
		{
			desc:       "not required and valid",
			method:     "GET",
			path:       "/false",
			token:      validToken,
			statusCode: http.StatusOK,
		},
		{
			desc:       "not required and expired",
			method:     "GET",
			path:       "/false",
			token:      expiredToken,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, tc.path, nil)

			if tc.token != "" {
				req.Header.Set("Authorization", fmt.Sprint("Bearer ", tc.token))
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tc.statusCode, w.Code, req)
		})
	}
}
