package filter_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/stretchr/testify/assert"
)

func TestErrorFilter(t *testing.T) {
	r := gin.Default()

	r.Use(filter.NewErrorFilter().Register())
	r.GET("/pass", func(ctx *gin.Context) {})
	r.GET("/err", func(ctx *gin.Context) {
		ctx.Error(errors.New("haha new error"))
	})
	r.GET("/status", func(ctx *gin.Context) {
		ctx.Error(status.NewError(http.StatusBadRequest, ""))
	})

	testcases := []struct {
		desc   string
		method string
		path   string
		body   io.Reader

		statusCode int
	}{
		{
			desc:       "ok",
			method:     "GET",
			path:       "/pass",
			body:       nil,
			statusCode: http.StatusOK,
		},
		{
			desc:       "custom status",
			method:     "GET",
			path:       "/status",
			body:       nil,
			statusCode: http.StatusBadRequest,
		},
		{
			desc:       "unexpected",
			method:     "GET",
			path:       "/err",
			body:       nil,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, tc.path, tc.body)

			r.ServeHTTP(w, req)

			assert.Equal(t, tc.statusCode, w.Code)
		})
	}
}
