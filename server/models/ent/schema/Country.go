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
		field.String("code").
			Unique().
			Immutable().
			MaxLen(3).
			StorageKey("code"),
		field.String("name").
			Unique().
			Immutable(),
		field.String("flag"),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("players", Player.Type),
		edge.To("leagues", League.Type),
		edge.To("teams", Team.Type),
	}
}
