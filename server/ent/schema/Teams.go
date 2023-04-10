package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coaches", Coach.Type),
		edge.To("players", Player.Type),
		edge.To("achievements", Achievements.Type),
		edge.To("transfers", Transfers.Type),
		edge.To("team_statistics", TeamStats.Type),
		edge.To("player_statistics", PlayerStats.Type),
		edge.From("league", League.Type).
			Ref("teams").
			Unique().
			Required(),
	}
}
