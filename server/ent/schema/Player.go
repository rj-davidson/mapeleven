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
		field.String("name"),
		field.Int("age"),
		field.String("photo"),
		field.String("position"),
		field.String("footed"),
		field.Int("salary"),
		field.String("country"),
		field.String("team"),
		field.String("league"),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("players").
			Unique(),
		edge.From("league", League.Type).
			Ref("players").
			Unique(),
		edge.From("achievements", Achievements.Type).
			Ref("player").
			Unique(),
		edge.From("playerStats", PlayerStats.Type).
			Ref("player").
			Unique(),
	}
}
