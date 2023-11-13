package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PlayerStats holds the schema definition for the PlayerStats entity.
type PlayerStats struct {
	ent.Schema
}

// Fields of the PlayerStats.
func (PlayerStats) Fields() []ent.Field {
	return []ent.Field{
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the PlayerStats.
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
		edge.To("psGames", PSGames.Type),
		edge.To("psShooting", PSShooting.Type),
		edge.To("psDefense", PSDefense.Type),
		edge.To("psTechnical", PSTechnical.Type),
		edge.To("psPenalty", PSPenalty.Type),
		edge.To("psSubstitutes", PSSubstitutes.Type),
		edge.To("season", Season.Type),
		edge.To("psFairplay", PSFairplay.Type),
	}
}
