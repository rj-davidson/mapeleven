package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PSDefense holds the schema definition for the PSDefense entity.
type PSDefense struct {
	ent.Schema
}

// Fields of the PSDefense.
func (PSDefense) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tacklesTotal"),
		field.Int("blocks").Default(0),
		field.Int("interceptions"),
		field.Int("totalDuels"),
		field.Int("wonDuels"),
	}
}

// Edges of the PSDefense.
func (PSDefense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).Ref("psdefense"),
	}
}
