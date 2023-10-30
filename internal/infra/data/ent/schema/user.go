package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	t := true
	return []ent.Field{
		field.Uint64("id").
			Positive().
			Immutable().
			Annotations(entsql.Annotation{
				Incremental: &t,
			}),
		field.String("email").NotEmpty().MaxLen(100).Unique(),
		field.String("name").MinLen(2).MaxLen(10),
		field.Int("coins").Min(0),
		field.Int("grade").Range(1, 3).Nillable().Optional(),
		field.Int("class").Range(1, 4).Nillable().Optional(),
		field.Int("number").Positive().Nillable().Optional(),
		field.Enum("gender").GoType(constant.Gender("")),
		field.Enum("userType").GoType(constant.UserType("")),
		field.Enum("role").GoType(constant.Role("")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notices", Notice.Type),
		edge.To("received", Pocket.Type),
		edge.To("sent", Pocket.Type),
		edge.To("revealed", Pocket.Type),
	}
}
