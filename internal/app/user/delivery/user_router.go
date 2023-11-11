package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/middleware"
)

type RankRequest struct {
	SortType constant.SortType  `form:"sortType" binding:"required,enum"`
	UserType *constant.UserType `form:"userType" binding:"enum"`
	grade    *int               `form:"grade" binding:"number"`
	class    *int               `form:"class" binding:"number"`
	name     *string            `form:"name"`
}

type QueryRequest struct {
	Query string
}

type UserRouter struct {
	userUseCase domain.UserUseCase
}

func NewUserRouter(uc domain.UserUseCase) *UserRouter {
	return &UserRouter{uc}
}

func (r *UserRouter) Register(engine *gin.Engine, m middleware.Middlewares) {
	engine.Use(m.AuthFilter.WithRequired(true), m.LogInterceptor.Register())

	engine.GET("/users/:userID", r.getUserDetail)
	engine.GET("/users/me/coins", r.countCoins)
	engine.GET("/users/me", r.getMyDetail)
	engine.GET("/rank", r.getRanking)
	engine.GET("/users", r.search)
}

func (r *UserRouter) getUserDetail(c *gin.Context) {
	id := c.Param("userID")

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid"))
		return
	}

	userInfo, err := r.userUseCase.GetUserDetail(c.Request.Context(), &input.UserIDInput{UserID: userID})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func (r *UserRouter) countCoins(c *gin.Context) {
	coins, err := r.userUseCase.CountCoins(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, coins)
}

func (r *UserRouter) getMyDetail(c *gin.Context) {
	userInfo, err := r.userUseCase.GetMyDetail(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func (r *UserRouter) getRanking(c *gin.Context) {
	var rank RankRequest

	if err := c.ShouldBindQuery(&rank); err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid param"))
		return
	}

	rankOutput, err := r.userUseCase.GetRanking(c.Request.Context(), &input.RankQueryInput{
		SortType: rank.SortType,
		UserType: rank.UserType,
		Grade:    rank.grade,
		Class:    rank.class,
		Name:     rank.name,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rankOutput)
}

func (r *UserRouter) search(c *gin.Context) {
	var query QueryRequest

	if err := c.ShouldBindQuery(&query); err != nil {
		c.Error(status.NewError(http.StatusBadRequest, "not valid param"))
		return
	}

	res, err := r.userUseCase.Search(c.Request.Context(), &input.SearchInput{
		Query: query.Query,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}
