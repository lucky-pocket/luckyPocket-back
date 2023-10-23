package domain

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

type User struct {
	UserID   uint64
	Name     string
	Coins    int
	Gender   constant.Gender
	UserType constant.UserType
	Role     constant.Role

	Grade  *int
	Class  *int
	Number *int
}

func (u User) IsAdmin() bool {
	return u.Role == constant.RoleAdmin
}

type UserUseCase interface {
	GetMyDetail(ctx context.Context) (*output.MyDetailOutput, error)
	CountCoins(ctx context.Context) (*output.CoinOutput, error)
	GetUserDetail(ctx context.Context) (*output.UserInfo, error)
	GetRanking(ctx context.Context, input *input.UserInput) (*output.RankOutput, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, userID uint64) (*User, error)
	Exists(ctx context.Context, userID uint64) (bool, error)
	FindStudentsByFilter(ctx context.Context, sortType constant.SortType, name *string, grade, class *int) ([]output.RankElem, error)
	FindNonStudentByFilter(ctx context.Context, sortType constant.SortType, name *string) ([]output.RankElem, error)
}
