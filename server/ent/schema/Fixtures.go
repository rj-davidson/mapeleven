package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Fixture holds the schema definition for the Fixture entity.
type Fixture struct {
	ent.Schema
}

// Fields of the Fixture.
func (Fixture) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable(),
		field.Time("date"),
		field.Int("time_zone"),
	}
}

// Edges of the Fixture.
func (Fixture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("competition", Competition.Type).
			Ref("fixtures").
			Unique().
			Required(),
		edge.To("match_statistics", MatchStatistics.Type),
		edge.To("predictions", Predictions.Type),
		edge.To("status", Status.Type).
			Unique(),
		edge.To("round", Round.Type),
		edge.To("lineups", Lineup.Type),
		edge.To("events", Event.Type),
		edge.To("h2h", H2H.Type),
		edge.To("playerStats", PlayerStats.Type).
			Field("fixture_id").
			Unique(),
	}
}
