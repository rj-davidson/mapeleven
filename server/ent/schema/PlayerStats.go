package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlayerStats holds the schema definition for the PlayerStats entity.
type PlayerStats struct {
	ent.Schema
}

// Fields of the PlayerStats.
func (PlayerStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("goals"),
		field.Int("assists"),
		field.Int("playing_time"),
	}
}

// Edges of the PlayerStats.
func (PlayerStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("playerStats").
			Unique().
			Required(),
		edge.From("fixture", Fixture.Type).
			Ref("playerStats").
			Unique().
			Required(),
	}
}
