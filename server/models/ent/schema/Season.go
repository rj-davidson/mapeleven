package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Season holds the schema definition for the Season entity.
type Season struct {
	ent.Schema
}

// Fields of the Season.
func (Season) Fields() []ent.Field {
	return []ent.Field{
		field.Int("year"),
		field.Time("start"),
		field.Time("end"),
		field.Bool("current").
			Default(false),
	}
}

// Edges of the Season.
func (Season) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("league", League.Type).
			Unique(),
	}
}
