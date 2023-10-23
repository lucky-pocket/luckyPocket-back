package output

import "github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"

type UserInfo struct {
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
}

type MyDetailOutput struct {
	UserInfo
	UserRole           constant.Role `json:"userRole"`
	HasNewNotification bool          `json:"hasNewNotification"`
}

type CoinOutput struct {
	Coins int `json:"coins"`
}

type RankElem struct {
	UserInfo
	Gender constant.Gender `json:"gender"`
	Amount int             `json:"amount"`
}

type RankOutput struct {
	Users []RankElem `json:"users"`
}
