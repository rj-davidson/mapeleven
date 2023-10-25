package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TSFixtures holds the schema definition for the TSFixtures entity.
type TSFixtures struct {
	ent.Schema
}

// Fields of the TSFixtures.
func (TSFixtures) Fields() []ent.Field {
	return []ent.Field{
		field.Int("playedHome").
			Optional().
			Default(0),
		field.Int("playedAway").
			Optional().
			Default(0),
		field.Int("playedTotal").
			Optional().
			Default(0),
		field.Int("winsHome").
			Optional().
			Default(0),
		field.Int("winsAway").
			Optional().
			Default(0),
		field.Int("winsTotal").
			Optional().
			Default(0),
		field.Int("drawsHome").
			Optional().
			Default(0),
		field.Int("drawsAway").
			Optional().
			Default(0),
		field.Int("drawsTotal").
			Optional().
			Default(0),
		field.Int("lossesHome").
			Optional().
			Default(0),
		field.Int("lossesAway").
			Optional().
			Default(0),
		field.Int("lossesTotal").
			Optional().
			Default(0),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the TSFixtures.
func (TSFixtures) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("fixtures_stats").
			Unique().
			Required(),
	}
}
