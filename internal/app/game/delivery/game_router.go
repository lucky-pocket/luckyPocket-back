package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
)

type FreeRequest struct {
	Free *bool `json:"free" binding:"required"`
}

type GameRouter struct {
	gameUseCase domain.GameUseCase
}

func NewGameRouter(uc domain.GameUseCase) *GameRouter {
	return &GameRouter{uc}
}

func (r *GameRouter) Register(engine *gin.Engine) {
	engine.GET("/games/free-ticket", r.getTicketInfo)
	engine.POST("/games/yut", r.playYut)
}

func (r *GameRouter) getTicketInfo(c *gin.Context) {
	ticket, err := r.gameUseCase.GetTicketInfo(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func (r *GameRouter) playYut(c *gin.Context) {
	var freeRequest FreeRequest

	if err := c.ShouldBindJSON(&freeRequest); err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid"))
		return
	}

	result, err := r.gameUseCase.PlayYut(c.Request.Context(), &input.FreeInput{Free: *freeRequest.Free})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
