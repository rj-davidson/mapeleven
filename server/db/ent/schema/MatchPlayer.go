package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// MatchPlayer holds the schema definition for the MatchPlayer entity.
type MatchPlayer struct {
	ent.Schema
}

// Fields of the MatchPlayer.
func (MatchPlayer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number"),
		field.String("position"),
		field.String("x"),
		field.String("y"),

		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the MatchPlayer.
func (MatchPlayer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("matchPlayer").
			Unique().
			Required(),
		edge.From("lineup", FixtureLineups.Type).
			Ref("lineupPlayer").
			Unique().
			Required(),
	}
}
