package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	_ "entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	_ "entgo.io/ent/schema/field"
	_ "time"
)

// PSGames holds the schema definition for the PSGames entity.
type PSGames struct {
	ent.Schema
}

// Fields of the PSGames.
func (PSGames) Fields() []ent.Field {
	return []ent.Field{
		field.Int("appearences"),
		field.Int("lineups"),
		field.Int("minutes"),
		field.Int("number"),
		field.String("position"),
		field.String("rating"),
		field.Bool("captain"),
	}
}

// Edges of the PSGames.
func (PSGames) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).Ref("psgames").Unique(),
	}
}
