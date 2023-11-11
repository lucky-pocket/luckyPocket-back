package middleware

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/interceptor"
)

type Middlewares struct {
	AuthFilter     *filter.AuthFilter
	LogInterceptor *interceptor.Logger
}
