package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

// Notice holds the schema definition for the Notice entity.
type Notice struct {
	ent.Schema
}

// Fields of the Notice.
func (Notice) Fields() []ent.Field {
	t := true
	return []ent.Field{
		field.Uint64("id").
			Positive().
			Immutable().
			Annotations(entsql.Annotation{
				Incremental: &t,
			}),
		field.Enum("type").GoType(constant.NoticeType("")),
		field.Bool("checked"),
	}
}

// Edges of the Notice.
func (Notice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("notices").Unique(),
		edge.To("pocket", Pocket.Type).Unique(),
	}
}
