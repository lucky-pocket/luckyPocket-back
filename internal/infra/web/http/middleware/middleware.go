package middleware

import (
	"fmt"
	"net/http"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
)

type RateLimiter struct {
	limiter func(ctx *gin.Context)
	storage ratelimit.Store
}

func NewRateLimiter() *RateLimiter {
	l := &RateLimiter{}

	l.storage = ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})

	l.limiter = ratelimit.RateLimiter(l.storage, &ratelimit.Options{
		ErrorHandler: func(ctx *gin.Context, i ratelimit.Info) {
			ctx.Error(status.NewError(
				http.StatusTooManyRequests,
				fmt.Sprintf("rate limit exceeded. try again in %dms", time.Until(i.ResetTime).Milliseconds()),
			))
		},
		KeyFunc: func(ctx *gin.Context) string { return ctx.ClientIP() },
	})

	return l
}

func (r *RateLimiter) Register() gin.HandlerFunc {
	return r.limiter
}
