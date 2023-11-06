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
		field.String("form").Optional(),
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
		edge.To("playerEvents", FixtureEvents.Type),
		edge.To("matchPlayer", MatchPlayer.Type),
		edge.To("assistEvents", FixtureEvents.Type),
		edge.To("psgames", PSGames.Type),
		edge.To("psgoals", PSGoals.Type),
		edge.To("psdefense", PSDefense.Type),
		edge.To("psoffense", PSOffense.Type),
		edge.To("pspenalty", PSPenalty.Type),
		edge.From("season", Season.Type).
			Ref("player").
			Unique(),
		edge.From("team", Team.Type).
			Ref("players").
			Unique(),
		edge.From("club", Club.Type).
			Ref("player").
			Unique(),
		edge.From("league", League.Type).
			Ref("player").
			Unique(),
	}
}
