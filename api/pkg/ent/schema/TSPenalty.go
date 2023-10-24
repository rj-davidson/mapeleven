package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TSPenalty holds the schema definition for the TSPenalty entity.
type TSPenalty struct {
	ent.Schema
}

// Fields of the TSPenalty.
func (TSPenalty) Fields() []ent.Field {
	return []ent.Field{
		field.Int("total").
			Optional().
			Default(0),
		field.Int("scoredTotal").
			Default(0).
			Optional(),
		field.Int("missedTotal").
			Default(0).
			Optional(),
		field.String("scoredPercentage").
			Default("0%").
			Optional(),
		field.String("missedPercentage").
			Default("0%").
			Optional(),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now)}
}

// Edges of the TSPenalty.
func (TSPenalty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("penalty_stats").
			Unique().
			Required(),
	}
}
