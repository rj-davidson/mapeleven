package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// FixtureEvents holds the schema definition for the FixtureEvents entity.
type FixtureEvents struct {
	ent.Schema
}

// Fields of the FixtureEvents.
func (FixtureEvents) Fields() []ent.Field {
	return []ent.Field{
		field.Int("elapsedTime"),
		field.Int("extraTime").Optional(),
		field.String("type"),
		field.String("detail"),
		field.String("comments").Optional(),

		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the FixtureEvents.
func (FixtureEvents) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("playerEvents").
			Unique().
			Required(),
		edge.From("assist", Player.Type).
			Ref("assistEvents").
			Unique(),
		edge.From("team", Team.Type).
			Ref("teamFixtureEvents").
			Unique().
			Required(),
		edge.From("fixture", Fixture.Type).
			Ref("fixtureEvents").
			Unique().
			Required(),
		edge.From("playerStats", PlayerStats.Type).
			Ref("playerEvents").
			Unique(),
	}
}
