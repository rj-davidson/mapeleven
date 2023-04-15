package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name").MaxLen(255),
		field.String("code").MaxLen(2),
		field.String("flag").MaxLen(255),
		field.String("flag_icon").MaxLen(255),
		field.String("search").MaxLen(3),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("competition", Competition.Type).
			Unique().
			Required(),
	}
}
