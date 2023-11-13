package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PSDefense holds the schema definition for the PSDefense entity.
type PSDefense struct {
	ent.Schema
}

// Fields of the PSDefense.
func (PSDefense) Fields() []ent.Field {
	return []ent.Field{
		field.Int("TacklesTotal").Default(0),
		field.Int("Blocks").Default(0),
		field.Int("Interceptions").Default(0),
		field.Int("DuelsTotal").Default(0),
		field.Int("WonDuels").Default(0),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the PSDefense.
func (PSDefense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("psDefense").
			Unique(),
	}
}
