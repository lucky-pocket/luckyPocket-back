package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
)

type PocketQuery struct {
	Offset int `form:"offset" binding:"required,number"`
	Limit  int `form:"limit" binding:"required,number"`
}

type PocketRequest struct {
	ReceiverID uint64 `json:"receiverID" binding:"required"`
	Coins      int    `json:"coins" binding:"required,number"`
	Message    string `json:"message" binding:"required"`
	IsPublic   bool   `json:"isPublic" binding:"required"`
}

type VisibilityRequest struct {
	Visible bool `json:"visible" binding:"required"`
}

type PocketRouter struct {
	pocketUseCase domain.PocketUseCase
}

func NewPocketRouter(pc domain.PocketUseCase) *PocketRouter {
	return &PocketRouter{pc}
}

func (p *PocketRouter) Register(engine *gin.Engine) {
	engine.POST("/pockets", p.sendPocket)
	engine.GET("/pockets/:pocketID", p.getPocketDetail)
	engine.PATCH("/users/me/pockets/:pocketID/visibility", p.setVisibility)
	engine.POST("/users/me/pockets/:pocketID/sender", p.revealSender)
	engine.GET("/users/me/pockets/received", p.getMyPockets)
	engine.GET("/users/:userID/pockets", p.getUserPockets)
}

func (p *PocketRouter) getMyPockets(c *gin.Context) {
	var pocketQuery PocketQuery

	userInfo := auth.MustExtract(c.Request.Context())

	if err := c.ShouldBindQuery(&pocketQuery); err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid param"))
		return
	}

	pocketsList, err := p.pocketUseCase.GetUserPockets(c.Request.Context(), &input.PocketQueryInput{
		UserID: userInfo.UserID,
		Offset: pocketQuery.Offset,
		Limit:  pocketQuery.Limit,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, pocketsList)
}

func (p *PocketRouter) getUserPockets(c *gin.Context) {
	var pocketQuery PocketQuery

	id := c.Param("userID")

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid"))
		return
	}

	if err := c.ShouldBindQuery(&pocketQuery); err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid param"))
		return
	}

	pocketsList, err := p.pocketUseCase.GetUserPockets(c.Request.Context(), &input.PocketQueryInput{
		UserID: userID,
		Offset: pocketQuery.Offset,
		Limit:  pocketQuery.Limit,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, pocketsList)
}

func (p *PocketRouter) sendPocket(c *gin.Context) {
	var pocketRequest PocketRequest

	if err := c.ShouldBindJSON(&pocketRequest); err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "bad request"))
		return
	}

	err := p.pocketUseCase.SendPocket(c.Request.Context(), &input.PocketInput{
		ReceiverID: pocketRequest.ReceiverID,
		Coins:      pocketRequest.Coins,
		Message:    pocketRequest.Message,
		IsPublic:   pocketRequest.IsPublic,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusCreated)
}

func (p *PocketRouter) getPocketDetail(c *gin.Context) {
	id := c.Param("pocketID")

	pocketID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid"))
		return
	}

	pocketOutput, err := p.pocketUseCase.GetPocketDetail(c.Request.Context(), &input.PocketIDInput{PocketID: pocketID})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, pocketOutput)
}

func (p *PocketRouter) setVisibility(c *gin.Context) {
	var visible VisibilityRequest

	id := c.Param("pocketID")

	pocketID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid"))
		return
	}

	if err := c.ShouldBindJSON(&visible); err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid request"))
		return
	}

	err = p.pocketUseCase.SetVisibility(c.Request.Context(),
		&input.VisibilityInput{PocketID: pocketID, Visible: visible.Visible})
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

func (p *PocketRouter) revealSender(c *gin.Context) {
	id := c.Param("pocketID")

	pocketID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid"))
		return
	}

	err = p.pocketUseCase.RevealSender(c.Request.Context(), &input.PocketIDInput{PocketID: pocketID})
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
