package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MatchStatistics holds the schema definition for the MatchStatistics entity.
type MatchStatistics struct {
	ent.Schema
}

// Fields of the MatchStatistics.
func (MatchStatistics) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable(),
		field.Int("shots_on_goal"),
		field.Int("shots_off_goal"),
		field.Int("shots_insidebox"),
		field.Int("shots_outsidebox"),
		field.Int("total_shots"),
		field.Int("blocked_shots"),
		field.Int("fouls"),
		field.Int("corner_kicks"),
		field.Int("offsides"),
		field.Int("ball_possession"),
		field.Int("yellow_cards"),
		field.Int("red_cards"),
		field.Int("goalkeeper_saves"),
		field.Int("total_passes"),
		field.Int("passes_accurate"),
		field.Float("passes_percent"),
	}
}

// Edges of the MatchStatistics.
func (MatchStatistics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fixture", Fixture.Type).
			Ref("match_statistics").
			Unique().
			Required(),
		edge.From("team", Team.Type).
			Ref("match_statistics").
			Field("team_id").
			Unique(),
	}
}
