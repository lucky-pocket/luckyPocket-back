package input

import "github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"

type UserInput struct {
	SortType constant.SortType
	UserType constant.UserType
	grade    *int
	class    *int
	name     *string
}
