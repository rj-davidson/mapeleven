package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transfers holds the schema definition for the Transfers entity.
type Transfers struct {
	ent.Schema
}

// Fields of the Transfers.
func (Transfers) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("info"),
	}
}

// Edges of the Transfers.
func (Transfers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team1", Team.Type).
			Ref("transfers_outgoing").
			Unique().
			Required(),
		edge.From("team2", Team.Type).
			Ref("transfers_incoming").
			Unique().
			Required(),
		edge.From("player", Player.Type).
			Ref("transfers").
			Unique().
			Required(),
	}
}
