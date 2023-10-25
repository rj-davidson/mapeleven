package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// FixtureLineups holds the schema definition for the FixtureLineups entity.
type FixtureLineups struct {
	ent.Schema
}

// Fields of the FixtureLineups.
func (FixtureLineups) Fields() []ent.Field {
	return []ent.Field{
		field.String("formation"),

		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the FixtureLineups.
func (FixtureLineups) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("fixtureLineups").
			Unique().
			Required(),
		edge.From("fixture", Fixture.Type).
			Ref("lineups").
			Unique().
			Required(),

		edge.To("lineupPlayer", MatchPlayer.Type),
	}
}
