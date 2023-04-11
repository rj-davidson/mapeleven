package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Standing holds the schema definition for the Standing entity.
type Standing struct {
	ent.Schema
}

// Fields of the Standing.
func (Standing) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("points"),
		field.Int("wins"),
		field.Int("draws"),
		field.Int("losses"),
		field.Int("goals_for"),
		field.Int("goals_against"),
	}
}

// Edges of the Standing.
func (Standing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("competition", Competition.Type).
			Ref("name").
			Unique(),
	}
}
