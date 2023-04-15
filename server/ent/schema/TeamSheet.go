package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TeamSheet holds the schema definition for the TeamSheet entity.
type TeamSheet struct {
	ent.Schema
}

// Fields of the TeamSheet.
func (TeamSheet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("x"),
		field.Int("y"),
		field.Enum("type").Values("Strt", "Bnch", "MF", "Sus", "Inj"),
	}
}

// Edges of the TeamSheet.
func (TeamSheet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("teamsheet").
			Unique().
			Required(),
		edge.From("Lineup", Lineup.Type).Ref("teamSheet").Required(),
	}
}
