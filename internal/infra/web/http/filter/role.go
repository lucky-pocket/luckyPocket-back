package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"net/http"
)

var (
	errNoPermission = status.NewError(http.StatusForbidden, "permission denied")
)

type RoleFilter struct{}

func NewRoleFilter() *RoleFilter {
	return &RoleFilter{}
}

func (f *RoleFilter) Register(role constant.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		info := auth.MustExtract(c.Request.Context())

		if info.Role != role {
			c.Error(errNoPermission)
			return
		}

		c.Next()
	}
}
