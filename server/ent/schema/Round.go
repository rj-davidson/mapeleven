package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Round holds the schema definition for the Round entity.
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
		field.Bool("current").
			Optional(),
	}
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("season", Season.Type).
			Ref("rounds").
			Unique().
			Required(),
		edge.To("fixtures", Fixture.Type),
		edge.From("league", League.Type).
			Ref("rounds").
			Field("league_id").
			Unique(),
	}
}
