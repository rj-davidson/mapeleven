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
		field.Int("ApiFootballID").
			Unique().
			Immutable(),
		field.String("slug").
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
		//field.Int("leagueID"),
		//field.Int("season"),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("birth", Birth.Type).
			Ref("player").
			Unique(),
		//edge.From("team", Team.Type).
		//	Ref("players").
		//	Required(),
		//edge.To("playerSeasons", PlayerSeason.Type),
	}
}
