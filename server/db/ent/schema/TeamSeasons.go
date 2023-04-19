// team_season.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TeamSeason struct {
	ent.Schema
}

func (TeamSeason) Fields() []ent.Field {
	return []ent.Field{
		field.Int("team_season_id").Unique(),
	}
}

func (TeamSeason) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("teamSeasons").
			Unique().
			Required(),
		edge.From("season", Season.Type).
			Ref("teamSeasons").
			Unique().
			Required(),
		edge.To("playerTeamSeasons", PlayerTeamSeason.Type),
	}
}
