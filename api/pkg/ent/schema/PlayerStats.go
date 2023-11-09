package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Player holds the schema definition for the Player entity.
type PlayerStats struct {
	ent.Schema
}

// Fields of the Player.
func (PlayerStats) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").
			Unique().
			Immutable(),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the Player.
func (PlayerStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("playerStats").
			Unique(),
		edge.From("team", Team.Type).
			Ref("playerStats").
			Unique(),
		edge.To("playerEvents", FixtureEvents.Type),
		edge.To("matchPlayer", MatchPlayer.Type),
		edge.To("assistEvents", FixtureEvents.Type),
		edge.To("psgames", PSGames.Type),
		edge.To("psgoals", PSGoals.Type),
		edge.To("psdefense", PSDefense.Type),
		edge.To("psoffense", PSOffense.Type),
		edge.To("pspenalty", PSPenalty.Type),
		edge.To("season", Season.Type),
	}
}
