package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TSFailedToScore holds the schema definition for the TSFailedToScore entity.
type TSFailedToScore struct {
	ent.Schema
}

// Fields of the TSFailedToScore.
func (TSFailedToScore) Fields() []ent.Field {
	return []ent.Field{
		field.Int("home").
			Optional().
			Default(0),
		field.Int("away").
			Optional().
			Default(0),
		field.Int("total").
			Optional().
			Default(0),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the TSFailedToScore.
func (TSFailedToScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("failed_to_score_stats").
			Unique().
			Required(),
	}
}
