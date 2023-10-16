package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").
			Unique().
			Immutable(),
		field.Int("ApiFootballId").
			Unique().
			Immutable(),
		field.String("name"),
		field.String("firstname"),
		field.String("lastname"),
		field.Int("age"),
		field.String("height"),
		field.String("weight"),
		field.Bool("injured"),
		field.String("photo"),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("birth", Birth.Type).
			Ref("player").
			Unique(),
		edge.From("nationality", Country.Type).
			Ref("players").
			Unique(),
		edge.To("squad", Squad.Type),
	}
}
