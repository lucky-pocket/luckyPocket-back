package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent"
)

func ToPocketDomain(entity *ent.Pocket) *domain.Pocket {
	return &domain.Pocket{
		PocketID: entity.ID,
		Receiver: &domain.User{UserID: entity.Edges.Receiver.ID},
		Sender:   &domain.User{UserID: entity.Edges.Sender.ID},
		Content:  entity.Content,
		Coins:    entity.Coins,
		IsPublic: entity.IsPublic,
	}
}
