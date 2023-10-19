package input

import "github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"

type UserInput struct {
	SortType constant.SortType
	UserType constant.UserType
	Grade    *int
	Class    *int
	Name     *string
}
