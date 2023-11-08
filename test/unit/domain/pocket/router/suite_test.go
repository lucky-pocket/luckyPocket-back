package router_test

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lucky-pocket/luckyPocket-back/internal/app/pocket/delivery"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"

	v "github.com/lucky-pocket/luckyPocket-back/internal/global/validator"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PocketRouterTestSuite struct {
	suite.Suite
	engine            *gin.Engine
	r                 *delivery.PocketRouter
	mockPocketUseCase *mocks.PocketUseCase
}

func TestPocketRouterSuite(t *testing.T) {
	suite.Run(t, new(PocketRouterTestSuite))
}

func (p *PocketRouterTestSuite) SetupSuite() {
	err := v.Initialize(binding.Validator.Engine().(*validator.Validate))
	if err != nil {
		return
	}

	p.mockPocketUseCase = mocks.NewPocketUseCase(p.T())

	p.r = delivery.NewPocketRouter(p.mockPocketUseCase)

	p.engine = gin.Default()
	p.engine.Use(filter.NewErrorFilter().Register())
	p.r.Register(p.engine)
}
