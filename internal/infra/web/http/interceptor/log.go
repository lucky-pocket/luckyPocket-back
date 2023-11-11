package interceptor

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.Logger
}

func NewLogger(log *zap.Logger) *Logger {
	return &Logger{log}
}

func (l *Logger) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			defer l.log.Sync()
			for _, e := range c.Errors {
				if _, ok := e.Err.(*status.Err); !ok {
					l.log.Error("unexpected internal error",
						zap.String("clientIP", c.ClientIP()),
						zap.String("method", c.Request.Method),
						zap.String("path", c.FullPath()),
						zap.String("handler", c.HandlerName()),
						zap.Any("ctx", c.Request.Context()),
						zap.Stack("stack"),
						zap.Error(e.Err),
					)
				}
			}
		}
	}
}
