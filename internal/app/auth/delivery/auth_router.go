package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
)

type CodeQuery struct {
	Code string `form:"code" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `header:"Cookie"`
}

type AuthRouter struct {
	authUseCase domain.AuthUseCase
}

func NewAuthRouter(ac domain.AuthUseCase) *AuthRouter {
	return &AuthRouter{ac}
}

func (a *AuthRouter) Register(engine *gin.Engine) {
	engine.GET("/auth/gauth", a.login)
	engine.POST("/auth/logout", a.logout)
	engine.POST("/auth/refresh", a.refreshToken)
}

func (a *AuthRouter) login(c *gin.Context) {
	var codeQuery CodeQuery

	if err := c.ShouldBindQuery(&codeQuery); err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid code"))
		return
	}

	tokenOutput, err := a.authUseCase.Login(c.Request.Context(), &input.CodeInput{
		Code: codeQuery.Code,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tokenOutput)
}

func (a *AuthRouter) logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.Error(status.NewError(http.StatusUnauthorized, "no cookie"))
		return
	}

	err = a.authUseCase.Logout(c.Request.Context(), &input.RefreshInput{
		RefreshToken: refreshToken,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

func (a *AuthRouter) refreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.Error(status.NewError(http.StatusUnauthorized, "no cookie"))
		return
	}

	tokenOutput, err := a.authUseCase.RefreshToken(c.Request.Context(), &input.RefreshInput{
		RefreshToken: refreshToken,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tokenOutput)
}
