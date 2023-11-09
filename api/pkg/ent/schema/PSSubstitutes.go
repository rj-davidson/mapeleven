package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PSSubstitutes holds the schema definition for the PSSubstitutes entity.
type PSSubstitutes struct {
	ent.Schema
}

// Fields of the PSSubstitutes.
func (PSSubstitutes) Fields() []ent.Field {
	return []ent.Field{
		field.Int("In"),
		field.Int("Out"),
		field.Int("Bench"),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the PSSubstitutes.
func (PSSubstitutes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).Ref("psSubstitutes").Unique(),
	}
}
