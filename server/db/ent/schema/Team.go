package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("form").
			Optional(),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("season", Season.Type).
			Ref("teams").
			Unique(),
		edge.From("club", Club.Type).
			Ref("team").
			Unique().
			Required(),
		edge.To("standings", Standings.Type),
		edge.To("homeFixtures", Fixture.Type),
		edge.To("awayFixtures", Fixture.Type),
		edge.To("fixtureEvents", FixtureEvents.Type),
		edge.To("fixtureLineups", FixtureLineups.Type),
		edge.To("players", Player.Type),
		edge.To("squad", Squad.Type),

		// Team Statistics
		edge.To("biggest_stats", TSBiggest.Type).
			Unique(),
		edge.To("cards_stats", TSCards.Type).
			Unique(),
		edge.To("clean_sheet_stats", TSCleanSheet.Type).
			Unique(),
		edge.To("failed_to_score_stats", TSFailedToScore.Type).
			Unique(),
		edge.To("fixtures_stats", TSFixtures.Type).
			Unique(),
		edge.To("goals_stats", TSGoals.Type).
			Unique(),
		edge.To("lineups", TSLineups.Type),
		edge.To("penalty_stats", TSPenalty.Type).
			Unique(),
	}
}
