package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Standings struct {
	ent.Schema
}

// Fields of the Standings.
func (Standings) Fields() []ent.Field {
	return []ent.Field{
		field.Int("rank"),
		field.Int("points").Default(0),
		field.Int("goalsDiff").Default(0),
		field.String("group").Default(""),
		field.String("form").Default(""),
		field.String("status").Optional(),
		field.String("description").Optional(),

		// All
		field.Int("allPlayed"),
		field.Int("allWin"),
		field.Int("allDraw"),
		field.Int("allLose"),
		field.Int("allGoalsFor"),
		field.Int("allGoalsAgainst"),

		// Home
		field.Int("homePlayed"),
		field.Int("homeWin"),
		field.Int("homeDraw"),
		field.Int("homeLose"),
		field.Int("homeGoalsFor"),
		field.Int("homeGoalsAgainst"),

		// Away
		field.Int("awayPlayed"),
		field.Int("awayWin"),
		field.Int("awayDraw"),
		field.Int("awayLose"),
		field.Int("awayGoalsFor"),
		field.Int("awayGoalsAgainst"),

		field.Time("LastUpdated").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Standings.
func (Standings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("standings").
			Unique().
			Required(),
		edge.From("season", Season.Type).
			Ref("standings").
			Unique().
			Required(),
	}
}
