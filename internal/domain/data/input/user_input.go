package input

import "github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"

type RankQueryInput struct {
	SortType constant.SortType
	UserType constant.UserType
	Grade    *int
	Class    *int
	Name     *string
}

type SearchInput struct {
	SearchQuery string
}
