package router_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lucky-pocket/luckyPocket-back/internal/app/auth/delivery"
	v "github.com/lucky-pocket/luckyPocket-back/internal/global/validator"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/suite"
)

type AuthRouterTestSuite struct {
	suite.Suite
	engine          *gin.Engine
	r               *delivery.AuthRouter
	mockAuthUseCase *mocks.AuthUseCase
}

func TestAuthRouterSuite(t *testing.T) {
	suite.Run(t, new(AuthRouterTestSuite))
}

func (s *AuthRouterTestSuite) SetupSuite() {
	err := v.Initialize(binding.Validator.Engine().(*validator.Validate))
	if err != nil {
		return
	}

	s.mockAuthUseCase = mocks.NewAuthUseCase(s.T())

	s.r = delivery.NewAuthRouter(s.mockAuthUseCase)

	s.engine = gin.Default()
	s.engine.Use(filter.NewErrorFilter().Register())
	s.r.Register(s.engine)
}
