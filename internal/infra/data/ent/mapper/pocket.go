package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
)

func ToPocketDomain(entity *ent.Pocket) *domain.Pocket {
	return &domain.Pocket{
		PocketID:   entity.ID,
		ReceiverID: entity.ReceiverID,
		SenderID:   entity.SenderID,
		Content:    entity.Content,
		Coins:      entity.Coins,
		IsPublic:   entity.IsPublic,
	}
}
