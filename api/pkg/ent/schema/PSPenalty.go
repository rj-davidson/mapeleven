package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PSPenalty holds the schema definition for the PSPenalty entity.
type PSPenalty struct {
	ent.Schema
}

// Fields of the PSPenalty.
func (PSPenalty) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Won").Default(0),
		field.Int("Scored").Default(0),
		field.Int("Missed").Default(0),
		field.Int("Saved").Default(0),
		field.Int("Committed").Default(0),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the PSPenalty.
func (PSPenalty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("psPenalty").
			Unique(),
	}
}
