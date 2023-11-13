package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PSFairplay holds the schema definition for the PSFairplay entity.
type PSFairplay struct {
	ent.Schema
}

// Fields of the PSFairplay.
func (PSFairplay) Fields() []ent.Field {
	return []ent.Field{
		field.Int("FoulsCommitted").Default(0),
		field.Int("Yellow").Default(0),
		field.Int("YellowRed").Default(0),
		field.Int("Red").Default(0),
		field.Int("PenaltyConceded").Default(0),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the PSFairplay.
func (PSFairplay) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("psFairplay").
			Unique(),
	}
}
