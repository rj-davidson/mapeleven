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
		field.Int("FoulsCommitted"),
		field.Int("Yellow"),
		field.Int("YellowRed"),
		field.Int("Red"),
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
