package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PSOffense holds the schema definition for the PSOffense entity.
type PSOffense struct {
	ent.Schema
}

// Fields of the PSOffense.
func (PSOffense) Fields() []ent.Field {
	return []ent.Field{
		field.Int("DribbleAttempts"),
		field.Int("DribbleSuccess"),
		field.Int("DribblePast").Default(0),
		field.Int("PassesTotal"),
		field.Int("PassesKey"),
		field.Int("PassesAccuracy"),
	}
}

// Edges of the PSOffense.
func (PSOffense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("psoffense").
			Unique(),
	}
}
