package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Squad holds the schema definition for the Squad entity.
type Squad struct {
	ent.Schema
}

// Fields of the Squad.
func (Squad) Fields() []ent.Field {
	return []ent.Field{
		field.String("position"),
		field.Int("number"),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the Squad.
func (Squad) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("squad").
			Unique(),
		edge.From("team", Team.Type).
			Ref("squad").
			Unique(),
		edge.From("season", Season.Type).
			Ref("squad").
			Unique(),
	}
}
