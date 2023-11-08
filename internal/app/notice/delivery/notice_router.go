package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
)

type NoticeRouter struct {
	noticeUseCase domain.NoticeUseCase
}

func NewNoticeRouter(uc domain.NoticeUseCase) *NoticeRouter {
	return &NoticeRouter{uc}
}

func (r *NoticeRouter) Register(engine *gin.Engine) {
	engine.GET("/users/me/notices", r.getNotice)
	engine.PATCH("/users/me/notices/:noticeID", r.checkNotice)
}

func (r *NoticeRouter) getNotice(c *gin.Context) {
	notices, err := r.noticeUseCase.GetNotice(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, notices)
}

func (r *NoticeRouter) checkNotice(c *gin.Context) {
	id := c.Param("noticeID")

	noticeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid"))
		return
	}

	err = r.noticeUseCase.CheckNotice(c.Request.Context(), noticeID)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
