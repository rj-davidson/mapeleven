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
		field.String("slug").Unique().Immutable(),
		field.String("name"),
		field.String("code").MaxLen(3),
		field.Int("founded"),
		field.Bool("national"),
		field.String("logo"),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("standings", Standings.Type),
		edge.From("country", Country.Type).
			Ref("teams").
			Unique(),
		edge.From("leagues", League.Type).
			Ref("teams"),
		edge.To("players", Player.Type),
		edge.To("teamSeasons", TeamSeason.Type),
		edge.To("homeFixtures", Fixture.Type),
		edge.To("awayFixtures", Fixture.Type),
	}
}
