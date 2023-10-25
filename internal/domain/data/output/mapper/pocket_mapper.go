package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToPocketListOutput(pockets []*domain.Pocket) *output.PocketListOutput {
	out := &output.PocketListOutput{
		Pockets: make([]output.PocketElem, 0, len(pockets)),
	}

	for _, pocket := range pockets {
		out.Pockets = append(out.Pockets, output.PocketElem{
			PocketID: pocket.PocketID,
			IsEmpty:  pocket.IsEmpty(),
			IsPublic: pocket.IsPublic,
		})
	}

	return out
}

func ToPocketOutput(pocket *domain.Pocket, sender *domain.User) *output.PocketOutput {
	out := &output.PocketOutput{
		Content: pocket.Content,
		Coins:   pocket.Coins,
	}

	if sender != nil {
		out.Sender = ToUserInfo(*sender)
	}

	return out
}
