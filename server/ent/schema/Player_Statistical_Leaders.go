package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlayerStatisticalLeader holds the schema definition for the PlayerStatisticalLeader entity.
type PlayerStatisticalLeader struct {
	ent.Schema
}

// Fields of the PlayerStatisticalLeader.
func (PlayerStatisticalLeader) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("type"),
		field.Int("value"),
	}
}

// Edges of the PlayerStatisticalLeader.
func (PlayerStatisticalLeader) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("competition", Competition.Type).
			Ref("competition").
			Required(),
		edge.From("player", Player.Type).
			Ref("name").
			Unique().
			Required(),
	}
}
