package router_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lucky-pocket/luckyPocket-back/internal/app/notice/delivery"
	v "github.com/lucky-pocket/luckyPocket-back/internal/global/validator"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/suite"
)

type NoticeRouterTestSuite struct {
	suite.Suite
	engine            *gin.Engine
	r                 *delivery.NoticeRouter
	mockNoticeUseCase *mocks.NoticeUseCase
}

func TestNoticeRouterSuite(t *testing.T) {
	suite.Run(t, new(NoticeRouterTestSuite))
}

func (s *NoticeRouterTestSuite) SetupSuite() {
	err := v.Initialize(binding.Validator.Engine().(*validator.Validate))
	if err != nil {
		return
	}

	s.mockNoticeUseCase = mocks.NewNoticeUseCase(s.T())

	s.r = delivery.NewNoticeRouter(s.mockNoticeUseCase)

	s.engine = gin.Default()
	s.engine.Use(filter.NewErrorFilter().Register())

	s.engine.GET("/users/me/notices", s.r.GetNotice)
	s.engine.PATCH("/users/me/notices/", s.r.CheckAllNotices)
	s.engine.PATCH("/users/me/notices/:noticeID", s.r.CheckNotice)
}
