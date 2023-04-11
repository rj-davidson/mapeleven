package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.String("short").
			NotEmpty().
			Unique().
			MaxLen(4),
		field.String("long").
			NotEmpty(),
		field.String("description").
			Optional(),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fixtures", Fixture.Type).
			Unique(),
	}
}
