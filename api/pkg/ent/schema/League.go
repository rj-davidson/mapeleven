package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// League holds the schema definition for the League entity.
type League struct {
	ent.Schema
}

// Fields of the League.
func (League) Fields() []ent.Field {
	return []ent.Field{
		field.Int("footballApiId").
			Unique().
			Immutable(),
		field.String("slug").
			Unique().
			Immutable(),
		field.String("name"),
		field.Enum("type").
			Values("League", "Cup", "Tournament", "Friendly"),
		field.String("logo"),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the League.
func (League) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("country", Country.Type).
			Ref("leagues").
			Unique(),
		edge.To("season", Season.Type),
	}
}
