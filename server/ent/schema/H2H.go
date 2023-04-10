package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// H2H holds the schema definition for the H2H entity.
type H2H struct {
	ent.Schema
}

// Fields of the H2H.
func (H2H) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
	}
}

// Edges of the H2H.
func (H2H) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("away_team", Team.Type).
			Ref("h2h_away").
			Unique().
			Required(),
		edge.From("home_team", Team.Type).
			Ref("h2h_home").
			Unique().
			Required(),
		edge.From("fixture", Fixture.Type).
			Ref("h2h").
			Unique().
			Required(),
	}
}
