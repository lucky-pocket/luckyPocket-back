package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToMyDetailOutput(user domain.User, hasNewNotification bool) *output.MyDetailOutput {
	return &output.MyDetailOutput{
		UserInfo: output.UserInfo{
			UserID: user.UserID,
			Name:   user.Name,
		},
		UserRole:           user.Role,
		HasNewNotification: hasNewNotification,
	}
}

func ToUserInfo(user domain.User) *output.UserInfo {
	return &output.UserInfo{
		UserID: user.UserID,
		Name:   user.Name,
	}
}

func ToRankOutput(users []output.RankElem) *output.RankOutput {
	out := &output.RankOutput{
		Users: make([]output.RankElem, 0, len(users)),
	}

	for _, user := range users {
		out.Users = append(out.Users, output.RankElem{
			UserInfo: user.UserInfo,
			Gender:   user.Gender,
			Amount:   user.Amount,
		})
	}

	return out
}

func ToCoinOutput(coins int) *output.CoinOutput {
	return &output.CoinOutput{
		Coins: coins,
	}
}
