package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Coach holds the schema definition for the Birth entity.
type Coach struct {
	ent.Schema
}

// Fields of the Coach.
func (Coach) Fields() []ent.Field {
	return []ent.Field{
		field.Int("footballApiId").
			Unique().
			Immutable(),
		field.String("name"),
		field.String("photo"),

		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the Coach.
func (Coach) Edges() []ent.Edge {
	return []ent.Edge{}
}
