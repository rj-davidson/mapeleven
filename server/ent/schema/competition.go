package schema

import "entgo.io/ent"

// Competition holds the schema definition for the Competition entity.
type Competition struct {
	ent.Schema
}

// Fields of the Competition.
func (Competition) Fields() []ent.Field {
	return nil
}

// Edges of the Competition.
func (Competition) Edges() []ent.Edge {
	return nil
}
