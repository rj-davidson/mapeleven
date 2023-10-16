package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Season holds the schema definition for the Season entity.
type Season struct {
	ent.Schema
}

// Fields of the Season.
func (Season) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").
			Unique().
			Immutable(),
		field.Int("year").
			Immutable(),
		field.Time("start_date"),
		field.Time("end_date"),
		field.Bool("current"),

		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the Season.
func (Season) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("league", League.Type).
			Ref("season").
			Unique(),
		edge.To("fixtures", Fixture.Type),
		edge.To("standings", Standings.Type),
		edge.To("teams", Team.Type),
	}
}
