package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PlayerSeason struct {
	ent.Schema
}

// Fields of the PlayerSeason.
func (PlayerSeason) Fields() []ent.Field {
	return []ent.Field{
		field.Int("season"),
		field.Int("pID"),
		field.String("nationality"),
		field.String("position"),
		field.String("team"),
		field.Int("year"),
		field.Int("appearances"),
		field.Int("goals"),
		field.Int("assists"),
		field.Int("saves"),
	}
}

// Edges of the PlayerSeason.
func (PlayerSeason) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("player", Player.Type).
		//	Ref("playerSeasons").
		//	Unique(),
		edge.From("club", Club.Type).
			Ref("playerSeasons"),
	}
}
