package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Birth holds the schema definition for the Birth entity.
type Birth struct {
	ent.Schema
}

// Fields of the Birth.
func (Birth) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.String("place"),
		field.String("country"),
	}
}

// Edges of the Birth.
func (Birth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("player", Player.Type),
	}
}
