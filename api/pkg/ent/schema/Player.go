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
		field.Int("GameAppearances"),
		field.Int("GameLineups"),
		field.Int("GameMinutes"),
		field.Int("TotalShots"),
		field.Int("ShotsOnTarget"),
		field.Int("TotalGoals"),
		field.Int("Assists"),
		field.Int("GoalsConceded"),
		field.Int("Saves"),

		field.String("position"),
		field.String("rating"),
		field.String("teamName"),
		field.Int("teamID"),
		field.Int("leagueID"),
		field.String("leagueName"),
		field.Int("passTotal"),
		field.Int("passKey"),
		field.Int("passAccuracy"),
		field.Int("totalTackle"),
		field.Int("blocks").Default(0),
		field.Int("interceptions").Default(0),
		field.Int("duelsTotal"),
		field.Int("duelsWon"),
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
		edge.To("playerStats", PlayerStats.Type),
	}
}
