package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"net/http"
	"time"
)

type CodeQuery struct {
	Code string `form:"code" binding:"required"`
}

type AccessResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   string `json:"expiresAt"`
}

type AuthRouter struct {
	authUseCase domain.AuthUseCase
}

func NewAuthRouter(ac domain.AuthUseCase) *AuthRouter {
	return &AuthRouter{ac}
}

func (a *AuthRouter) Login(c *gin.Context) {
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

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    tokenOutput.Refresh.Token,
		Domain:   "",
		Expires:  tokenOutput.Refresh.ExpiresAt,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, AccessResponse{
		AccessToken: tokenOutput.Access.Token,
		ExpiresAt:   tokenOutput.Access.ExpiresAt.Local().Format(time.RFC3339),
	})
}

func (a *AuthRouter) Logout(c *gin.Context) {
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

func (a *AuthRouter) RefreshToken(c *gin.Context) {
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

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    tokenOutput.Refresh.Token,
		Domain:   "",
		Expires:  tokenOutput.Refresh.ExpiresAt,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, AccessResponse{
		AccessToken: tokenOutput.Access.Token,
		ExpiresAt:   tokenOutput.Access.ExpiresAt.Local().Format(time.RFC3339),
	})
}
