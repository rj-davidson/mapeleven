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
		field.String("location"),
		field.String("champion"),
		field.Int("num_teams"),
	}
}

// Edges of the League.
func (League) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teams", Team.Type).Unique(),
		edge.To("achievements", Achievements.Type).Unique(),
	}
}
