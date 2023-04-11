package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Competition holds the schema definition for the Competition entity.
type Competition struct {
	ent.Schema
}

// Fields of the Competition.
func (Competition) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
		field.String("name"),
		field.Int("season"),
		field.Enum("type").
			Values("league", "cup", "exhibition"),
		field.Bool("current"),
		field.String("logo"),
		field.String("search").
			MaxLen(3),
	}
}

// Edges of the Competition.
func (Competition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("country", Country.Type).
			Required().
			Field("country_id"),
		edge.To("seasons", Season.Type).
			Unique().
			Field("season"),
		edge.To("player_statistical_leaders", PlayerStatisticalLeader.Type).
			Unique().
			Field("competition"),
	}
}
