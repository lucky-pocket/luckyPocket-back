package repository

import (
	"context"
	"fmt"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

var justForTest = 1

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	builder := r.getClient(ctx).User.Create().
		SetEmail(fmt.Sprint(justForTest)).
		SetName(user.Name).
		SetRole(constant.RoleMember).
		SetUserType(user.UserType).
		SetGender(user.Gender).
		SetCoins(0)

	justForTest++

	if user.UserType == constant.TypeStudent {
		builder = builder.
			SetGrade(*user.Grade).
			SetClass(*user.Class).
			SetNumber(*user.Number)
	}

	_, err := builder.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
