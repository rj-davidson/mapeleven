package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// League holds the schema definition for the League entity.
type League struct {
	ent.Schema
}

// Fields of the League.
func (League) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
		field.Enum("type").
			Values("League", "Cup", "Tournament", "Friendly"),
		field.String("logo"),
	}
}

// Edges of the League.
func (League) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("seasons", Season.Type),
		edge.To("standings", Standings.Type),
		edge.To("teams", Team.Type),
		edge.From("country", Country.Type).
			Ref("leagues").
			Unique(),
	}
}
