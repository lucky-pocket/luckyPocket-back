package router_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lucky-pocket/luckyPocket-back/internal/app/user/delivery"
	v "github.com/lucky-pocket/luckyPocket-back/internal/global/validator"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/suite"
)

type UserRouterTestSuite struct {
	suite.Suite
	engine          *gin.Engine
	r               *delivery.UserRouter
	mockUserUseCase *mocks.UserUseCase
}

func TestUserRouterSuite(t *testing.T) {
	suite.Run(t, new(UserRouterTestSuite))
}

func (s *UserRouterTestSuite) SetupSuite() {
	err := v.Initialize(binding.Validator.Engine().(*validator.Validate))
	if err != nil {
		return
	}

	s.mockUserUseCase = mocks.NewUserUseCase(s.T())

	s.r = delivery.NewUserRouter(s.mockUserUseCase)

	s.engine = gin.Default()
	s.engine.Use(filter.NewErrorFilter().Register())

	s.engine.GET("/users/:userID", s.r.GetUserDetail)
	s.engine.GET("/users/me/coins", s.r.CountCoins)
	s.engine.GET("/users/me", s.r.GetMyDetail)
	s.engine.GET("/rank", s.r.GetRanking)
	s.engine.GET("/users", s.r.Search)
}
