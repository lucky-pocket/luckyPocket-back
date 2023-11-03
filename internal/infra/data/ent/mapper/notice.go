package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
)

func ToNoticeDomain(entity *ent.Notice) *domain.Notice {
	return &domain.Notice{
		NoticeID:  entity.ID,
		UserID:    entity.Edges.User.ID,
		PocketID:  entity.Edges.Pocket.ID,
		Type:      entity.Type,
		Checked:   entity.Checked,
		CreatedAt: entity.CreatedAt,
	}
}
