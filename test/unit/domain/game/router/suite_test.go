package router_test

import (
	"testing"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lucky-pocket/luckyPocket-back/internal/app/game/delivery"
	v "github.com/lucky-pocket/luckyPocket-back/internal/global/validator"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/suite"
)

type GameRouterTestSuite struct {
	suite.Suite
	engine          *gin.Engine
	r               *delivery.GameRouter
	mockGameUseCase *mocks.GameUseCase
}

func TestGameRouterSuite(t *testing.T) {
	suite.Run(t, new(GameRouterTestSuite))
}

func (s *GameRouterTestSuite) SetupSuite() {
	err := v.Initialize(binding.Validator.Engine().(*validator.Validate))
	if err != nil {
		return
	}

	s.mockGameUseCase = mocks.NewGameUseCase(s.T())

	s.r = delivery.NewGameRouter(s.mockGameUseCase)

	s.engine = gin.Default()
	s.engine.Use(filter.NewErrorFilter().Register())
	s.r.Register(s.engine)
}
