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

func (r *NoticeRouter) GetNotice(c *gin.Context) {
	notices, err := r.noticeUseCase.GetNotice(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, notices)
}

func (r *NoticeRouter) CheckAllNotices(c *gin.Context) {
	err := r.noticeUseCase.CheckAllNotices(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

func (r *NoticeRouter) CheckNotice(c *gin.Context) {
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
