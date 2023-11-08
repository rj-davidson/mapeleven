// schemas/psoffense.go
package schema

import (
	"entgo.io/ent"
	_ "entgo.io/ent/schema"
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
		field.Int("dribbleAttempts"),
		field.Int("dribbleSuccess"),
		field.Int("dribblePast").Default(0),
		field.Int("passesTotal"),
		field.Int("passesKey"),
		field.Int("passesAccuracy"),
	}
}

// Edges of the PSOffense.
func (PSOffense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).Ref("psoffense"),
	}
}
