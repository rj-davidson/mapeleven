package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Predictions holds the schema definition for the Predictions entity.
type Predictions struct {
	ent.Schema
}

// Fields of the Predictions.
func (Predictions) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable().
			Positive(),
		field.String("prediction").
			NotEmpty(),
	}
}

// Edges of the Predictions.
func (Predictions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fixture", Fixture.Type).
			Ref("predictions").
			Unique().
			Required(),
	}
}
