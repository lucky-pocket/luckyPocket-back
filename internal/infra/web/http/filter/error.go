package filter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
)

type ErrorFilter struct{}

func NewErrorFilter() *ErrorFilter {
	return &ErrorFilter{}
}

func (f *ErrorFilter) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// It should iterate once.
		for _, e := range c.Errors {
			statusErr, ok := e.Err.(*status.Err)
			if !ok {
				statusErr = status.NewError(http.StatusInternalServerError, "internal server error")
			}

			c.AbortWithStatusJSON(statusErr.Code, gin.H{
				"message": statusErr.Message,
			})
		}
	}
}
