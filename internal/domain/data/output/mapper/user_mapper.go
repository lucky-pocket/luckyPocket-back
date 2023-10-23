package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToUserToOutput(user domain.User, hasNewNotification bool) *output.MyDetailOutput {
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

func CheckUserNil(user *domain.User, err error) {
	if err != nil {
		return
	}

	if user == nil {
		return
	}
}

func RankOutput(users []output.RankElem) *output.RankOutput {
	rankElems := make([]output.RankElem, 0, len(users))

	for _, user := range users {
		rankElems = append(rankElems, output.RankElem{
			UserInfo: user.UserInfo,
			Gender:   user.Gender,
			Amount:   user.Amount,
		})
	}

	return &output.RankOutput{
		Users: rankElems,
	}
}
