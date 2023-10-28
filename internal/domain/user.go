package domain

import (
	"context"
	"github.com/onee-only/gauth-go"

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
	GetUserDetail(ctx context.Context, input *input.UserIDInput) (*output.UserInfo, error)
	CountCoins(ctx context.Context) (*output.CoinOutput, error)
	GetRanking(ctx context.Context, input *input.RankQueryInput) (*output.RankOutput, error)
}

type UserRepository interface {
	Create(ctx context.Context, userInfo gauth.UserInfo) (user *User, err error)
	FindByID(ctx context.Context, userID uint64) (*User, error)
	Exists(ctx context.Context, userID uint64) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindStudentsWithFilter(ctx context.Context, sortType constant.SortType, name *string, grade, class *int) ([]output.RankElem, error)
	FindNonStudentWithFilter(ctx context.Context, sortType constant.SortType, name *string) ([]output.RankElem, error)
	CountCoinsByUserID(ctx context.Context, userID uint64) (int, error)
	UpdateCoin(ctx context.Context, userID uint64, coin int) error
}
