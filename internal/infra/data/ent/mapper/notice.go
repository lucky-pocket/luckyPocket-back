package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
)

func ToNoticeDomain(entity *ent.Notice) *domain.Notice {
	return &domain.Notice{
		NoticeID:  entity.ID,
		UserID:    entity.UserID,
		PocketID:  entity.PocketID,
		Type:      entity.Type,
		Checked:   entity.Checked,
		CreatedAt: entity.CreatedAt,
	}
}
