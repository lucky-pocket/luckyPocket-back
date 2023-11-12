package router_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lucky-pocket/luckyPocket-back/internal/app/pocket/delivery"
	v "github.com/lucky-pocket/luckyPocket-back/internal/global/validator"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/suite"
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

	p.engine.POST("/pockets", p.r.SendPocket)
	p.engine.GET("/pockets/:pocketID", p.r.GetPocketDetail)
	p.engine.PATCH("/users/me/pockets/:pocketID/visibility", p.r.SetVisibility)
	p.engine.POST("/users/me/pockets/:pocketID/sender", p.r.RevealSender)
	p.engine.GET("/users/me/pockets/received", p.r.GetMyPockets)
	p.engine.GET("/users/:userID/pockets", p.r.GetUserPockets)
}
