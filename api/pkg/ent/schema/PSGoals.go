package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PSGoals holds the schema definition for the PSGoals entity.
type PSGoals struct {
	ent.Schema
}

// Fields of the PSGoals.
func (PSGoals) Fields() []ent.Field {
	return []ent.Field{
		field.Int("TotalGoals"),
		field.Int("ConcededGoals"),
		field.Int("AssistGoals"),
		field.Int("SaveGoals").Default(0),
		field.Int("ShotsTotal"),
		field.Int("ShotsOn"),
	}
}

// Edges of the PSGoals.
func (PSGoals) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("psgoals").
			Unique(),
	}
}
