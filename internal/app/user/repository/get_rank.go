package repository

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/pocket"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/user"
)

func (r *userRepository) RankStudents(ctx context.Context, sortType constant.SortType, name *string, grade, class *int) ([]output.RankElem, error) {
	builder := r.getClient(ctx).User.Query()

	if name != nil {
		builder = builder.Where(user.NameContains(*name))
	}

	if grade != nil {
		builder = builder.Where(user.Grade(*grade))
	}

	if class != nil {
		builder = builder.Where(user.Class(*class))
	}

	builder = builder.Where(user.UserTypeEQ(constant.TypeStudent))
	return r.getRank(ctx, builder, sortType)
}

func (r *userRepository) RankNonStudents(ctx context.Context, sortType constant.SortType, userType *constant.UserType, name *string) ([]output.RankElem, error) {
	builder := r.getClient(ctx).User.Query()

	if userType != nil {
		builder = builder.Where(user.UserTypeEQ(*userType))
	}

	if name != nil {
		builder = builder.Where(user.NameContains(*name))
	}

	builder = builder.Where(user.UserTypeNEQ(constant.TypeStudent))
	return r.getRank(ctx, builder, sortType)
}

func (r *userRepository) getRank(ctx context.Context, builder *ent.UserQuery, sortType constant.SortType) ([]output.RankElem, error) {
	switch sortType {
	case constant.SortTypeCoins:
		builder = builder.Modify(func(s *sql.Selector) {
			s.Select(sql.As(user.FieldCoins, "amount"))
		}).UserQuery
	case constant.SortTypePocket:
		subQuery := sql.Select(sql.Count("*")).
			From(sql.Table(pocket.Table)).
			Where(sql.ColumnsEQ(
				fmt.Sprintf("`%s`.`%s`", pocket.Table, pocket.ReceiverColumn),
				fmt.Sprintf("`%s`.`%s`", user.Table, user.FieldID),
			))
		builder = builder.Modify(func(s *sql.Selector) {
			query, _ := subQuery.Query()
			s.Select(sql.As(fmt.Sprintf("(%s)", query), "amount"))
		}).UserQuery
	}

	var elems []output.RankElem
	err := builder.
		Modify(func(s *sql.Selector) {
			s.AppendSelectAs(user.FieldID, "userId")
			s.AppendSelectAs(user.FieldUserType, "userType")
			s.AppendSelect(user.FieldName, user.FieldGender, user.FieldGrade, user.FieldClass)
			s.OrderBy(sql.Desc("amount"))
		}).
		Scan(ctx, &elems)

	return elems, err
}
