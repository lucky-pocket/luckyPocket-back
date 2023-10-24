package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pocket holds the schema definition for the Pocket entity.
type Pocket struct {
	ent.Schema
}

// Fields of the Pocket.
func (Pocket) Fields() []ent.Field {
	t := true
	return []ent.Field{
		field.Uint64("id").
			Positive().
			Immutable().
			Annotations(entsql.Annotation{
				Incremental: &t,
			}),
		field.String("content").MaxLen(1000),
		field.Uint64("coins"),
	}
}

// Edges of the Pocket.
func (Pocket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("receiver", User.Type).Ref("received").Unique(),
		edge.From("sender", User.Type).Ref("sent").Unique(),
		edge.From("revealers", User.Type).Ref("revealed"),
	}
}
