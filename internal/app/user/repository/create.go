package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/mapper"
)

func (r *userRepository) Create(ctx context.Context, userInfo domain.GAuthUser) (*domain.User, error) {
	builder := r.getClient(ctx).User.Create().
		SetEmail(userInfo.Email).
		SetName(*userInfo.Name).
		SetRole(constant.RoleMember).
		SetUserType(userInfo.Role).
		SetGender(userInfo.Gender).
		SetCoins(0)

	if userInfo.Role == constant.TypeStudent {
		builder = builder.
			SetGrade(*userInfo.Grade).
			SetClass(*userInfo.ClassNum).
			SetNumber(*userInfo.Num)
	}

	user, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserDomain(user), err
}
