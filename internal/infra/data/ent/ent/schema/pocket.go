package schema

import (
	"time"

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
			Annotations(entsql.Annotation{Incremental: &t}),
		field.Uint64("receiverID"),
		field.Uint64("senderID"),
		field.String("content").MaxLen(300),
		field.Int("coins").Min(0),
		field.Bool("isPublic"),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the Pocket.
func (Pocket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("receiver", User.Type).
			Ref("received").
			Field("receiverID").
			Unique().Required(),
		edge.From("sender", User.Type).
			Ref("sent").
			Field("senderID").
			Unique().Required(),
		edge.From("revealers", User.Type).Ref("revealed"),
	}
}
