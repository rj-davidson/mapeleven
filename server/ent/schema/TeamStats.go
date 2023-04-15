package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TeamStats holds the schema definition for the TeamStats entity.
type TeamStats struct {
	ent.Schema
}

// Fields of the TeamStats.
func (TeamStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
		field.String("country"),
		field.Int("wins"),
		field.Int("losses"),
		field.String("record"),
		field.Int("goals_for"),
		field.Int("goals_against"),
		field.Int("time_pos"),
		field.Int("assists"),
		field.Int("saves"),
	}
}

// Edges of the TeamStats.
func (TeamStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("teamStats").
			Unique().
			Required(),
		edge.To("achievements", Achievements.Type),
		edge.To("transfers", Transfers.Type),
		edge.From("coach", Coach.Type).Ref("teamStats"),
	}
}
