package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Lineup holds the schema definition for the Lineup entity.
type Lineup struct {
	ent.Schema
}

// Fields of the Lineup.
func (Lineup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
		field.String("formation").
			Optional(),
	}
}

// Edges of the Lineup.
func (Lineup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coach", Coach.Type).
			Ref("lineup").
			Unique().
			Required(),
		edge.To("teamsheet", TeamSheet.Type),
		edge.From("team", Team.Type).
			Ref("lineups").
			Field("team_id").
			Unique(),
		edge.From("fixture", Fixture.Type).
			Ref("lineups").
			Unique().
			Required(),
	}
}
