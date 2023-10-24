package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TSLineups holds the schema definition for the TSLineups entity.
type TSLineups struct {
	ent.Schema
}

// Fields of the TSLineups.
func (TSLineups) Fields() []ent.Field {
	return []ent.Field{
		field.String("formation").
			Immutable(),
		field.Int("played"),
		field.Int("team_id").
			Comment("This field is a foreign key which refers to the team this lineup belongs to"),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now)}
}

// Edges of the TSLineups.
func (TSLineups) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("lineups").
			Field("team_id").
			Unique().
			Required(),
	}
}
