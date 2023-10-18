package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Fixture holds the schema definition for the Fixture entity.
type Fixture struct {
	ent.Schema
}

// Fields of the Fixture.
func (Fixture) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").
			Unique().
			Immutable(),
		field.Int("apiFootballId").
			Unique().
			Immutable(),
		field.String("referee").
			Optional(),
		field.String("timezone").
			Optional(),
		field.Time("date"),
		field.Int("elapsed").
			Optional(),
		field.Int("round").
			Optional(),
		field.String("status"),
		field.Int("homeTeamScore").
			Optional(),
		field.Int("awayTeamScore").
			Optional(),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the Fixture.
func (Fixture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("homeTeam", Team.Type).
			Ref("homeFixtures").
			Unique().
			Required(),
		edge.From("awayTeam", Team.Type).
			Ref("awayFixtures").
			Unique().
			Required(),
		edge.From("season", Season.Type).
			Ref("fixtures").
			Unique().
			Required(),

		edge.To("lineups", FixtureLineups.Type),
		edge.To("fixtureEvents", FixtureEvents.Type),
	}
}
