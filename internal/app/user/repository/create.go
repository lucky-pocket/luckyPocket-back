package repository

import (
	"context"
	"strings"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/mapper"
)

func (r *userRepository) Create(ctx context.Context, userInfo domain.GAuthUser) (*domain.User, error) {
	usrType, _ := strings.CutPrefix(string(userInfo.Role), "ROLE_")

	userType := constant.UserType(usrType)
	gender := constant.Gender(userInfo.Gender)

	builder := r.getClient(ctx).User.Create().
		SetEmail(userInfo.Email).
		SetName(*userInfo.Name).
		SetRole(constant.RoleMember).
		SetUserType(userType).
		SetGender(gender).
		SetCoins(0)

	if userType == constant.TypeStudent {
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
