package domain

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/onee-only/gauth-go"
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
	Search(ctx context.Context, input *input.SearchInput) ([]*User, error)
}

type UserRepository interface {
	Create(ctx context.Context, userInfo gauth.UserInfo) (*User, error)
	FindByID(ctx context.Context, userID uint64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByNameContains(ctx context.Context, name string) ([]*User, error)
	Exists(ctx context.Context, userID uint64) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	RankStudents(ctx context.Context, sortType constant.SortType, name *string, grade, class *int) ([]output.RankElem, error)
	RankNonStudents(ctx context.Context, sortType constant.SortType, name *string) ([]output.RankElem, error)
	CountCoinsByUserID(ctx context.Context, userID uint64) (int, error)
	UpdateCoin(ctx context.Context, userID uint64, coin int) error
}
