package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GameLog holds the schema definition for the GameLog entity.
type GameLog struct {
	ent.Schema
}

// Fields of the GameLog.
func (GameLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("gameType"),
		field.Time("timestamp").Default(time.Now),
	}
}

// Edges of the GameLog.
func (GameLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("logs").Unique().Required(),
	}
}
