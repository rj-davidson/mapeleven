package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("season", Season.Type).
			Ref("teams").
			Unique(),
		edge.From("club", Club.Type).
			Ref("team").
			Unique().
			Required(),
		edge.To("standings", Standings.Type),
		edge.To("homeFixtures", Fixture.Type),
		edge.To("awayFixtures", Fixture.Type),
		edge.To("players", Player.Type),
	}
}
