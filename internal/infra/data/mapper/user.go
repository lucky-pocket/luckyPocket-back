package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent"
)

func ToUserDomain(entity *ent.User) *domain.User {
	return &domain.User{
		UserID:   entity.ID,
		Name:     entity.Name,
		Coins:    entity.Coins,
		Gender:   entity.Gender,
		UserType: entity.UserType,
		Role:     entity.Role,

		Grade:  entity.Grade,
		Class:  entity.Class,
		Number: entity.Number,
	}
}
