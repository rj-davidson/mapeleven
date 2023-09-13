package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("slug").Unique().Immutable(),
		field.String("name"),
		field.String("firstname"),
		field.String("lastname"),
		field.Int("age"),
		field.Float("height"),
		field.Float("weight"),
		field.Bool("injured"),
		field.String("photo"),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("birth", Birth.Type).
			Ref("player").Unique(),
	}
}
