package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Coach holds the schema definition for the Coach entity.
type Coach struct {
	ent.Schema
}

// Fields of the Coach.
func (Coach) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
		field.String("position"),
		field.String("record"),
		field.String("salary"),
	}
}

// Edges of the Coach.
func (Coach) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("coaches").
			Unique().
			Required(),
		edge.To("achievements", Achievements.Type),
		edge.To("transfers", Transfers.Type),
		edge.To("teamStats", TeamStats.Type).Required(),
		edge.To("lineup", Lineup.Type),
	}
}
