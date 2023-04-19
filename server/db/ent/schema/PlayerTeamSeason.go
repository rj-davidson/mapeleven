// player_team_season.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PlayerTeamSeason struct {
	ent.Schema
}

func (PlayerTeamSeason) Fields() []ent.Field {
	return []ent.Field{
		field.Int("player_team_season_id").Unique(),
	}
}

func (PlayerTeamSeason) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("playerTeamSeasons").
			Unique().
			Required(),
		edge.From("teamSeason", TeamSeason.Type).
			Ref("playerTeamSeasons").
			Unique().
			Required(),
	}
}
