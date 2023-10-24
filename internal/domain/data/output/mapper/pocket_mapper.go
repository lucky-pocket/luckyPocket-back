package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToPocketListOutput(pocketList []*domain.Pocket) *output.PocketListOutput {
	pockets := &output.PocketListOutput{}

	for _, p := range pocketList {
		isEmpty := true
		if p.Coins == 0 {
			isEmpty = true
		} else {
			isEmpty = false
		}

		elem := output.PocketElem{
			PocketID: p.PocketID,
			IsEmpty:  isEmpty,
			IsPublic: p.IsPublic,
		}

		pockets.Pockets = append(pockets.Pockets, elem)
	}

	return pockets
}

func ToPocketOutput(pocket *domain.Pocket, isPublic bool) *output.PocketOutput {
	var userInfo *output.UserInfo

	if isPublic {
		sender := pocket.Sender.Name
		userInfo = &output.UserInfo{
			UserID: pocket.Sender.UserID,
			Name:   sender,
		}
	} else {
		userInfo = nil
	}

	return &output.PocketOutput{
		Content: pocket.Content,
		Coins:   pocket.Coins,
		Sender:  userInfo,
	}
}
