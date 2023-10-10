package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Club holds the schema definition for the Club entity.
type Club struct {
	ent.Schema
}

// Fields of the Club.
func (Club) Fields() []ent.Field {
	return []ent.Field{
		field.Int("apiFootballId").
			Unique().
			Immutable(),
		field.String("slug").
			Unique().
			Immutable(),
		field.String("name"),
		field.String("code").
			MaxLen(3),
		field.Int("founded"),
		field.Bool("national"),
		field.String("logo"),
	}
}

// Edges of the Club.
func (Club) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("country", Country.Type).
			Ref("clubs").
			Unique(),
		edge.To("team", Team.Type),
		edge.To("playerSeasons", PlayerSeason.Type),
	}
}
