package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TSCards holds the schema definition for the TSCards entity.
type TSCards struct {
	ent.Schema
}

// Fields of the TSCards.
func (TSCards) Fields() []ent.Field {
	return []ent.Field{
		field.Int("yellow0To15Total").
			Default(0).
			Optional(),
		field.String("yellow0To15Percentage").
			Default("0%").
			Optional(),
		field.Int("yellow16To30Total").
			Default(0).
			Optional(),
		field.String("yellow16To30Percentage").
			Default("0%").
			Optional(),
		field.Int("yellow31To45Total").
			Default(0).
			Optional(),
		field.String("yellow31To45Percentage").
			Default("0%").
			Optional(),
		field.Int("yellow46To60Total").
			Default(0).
			Optional(),
		field.String("yellow46To60Percentage").
			Default("0%").
			Optional(),
		field.Int("yellow61To75Total").
			Default(0).
			Optional(),
		field.String("yellow61To75Percentage").
			Default("0%").
			Optional(),
		field.Int("yellow76To90Total").
			Default(0).
			Optional(),
		field.String("yellow76To90Percentage").
			Default("0%").
			Optional(),
		field.Int("yellow91to105Total").
			Default(0).
			Optional(),
		field.String("yellow91to105Percentage").
			Default("0%").
			Optional(),
		field.Int("yellow106To120Total").
			Default(0).
			Optional(),
		field.String("yellow106To120Percentage").
			Default("0%").
			Optional(),
		field.Int("red0To15Total").
			Default(0).
			Optional(),
		field.String("red0To15Percentage").
			Default("0%").
			Optional(),
		field.Int("red16To30Total").
			Default(0).
			Optional(),
		field.String("red16To30Percentage").
			Default("0%").
			Optional(),
		field.Int("red31To45Total").
			Default(0).
			Optional(),
		field.String("red31To45Percentage").
			Default("0%").
			Optional(),
		field.Int("red46To60Total").
			Default(0).
			Optional(),
		field.String("red46To60Percentage").
			Default("0%").
			Optional(),
		field.Int("red61To75Total").
			Default(0).
			Optional(),
		field.String("red61To75Percentage").
			Default("0%").
			Optional(),
		field.Int("red76To90Total").
			Default(0).
			Optional(),
		field.String("red76To90Percentage").
			Default("0%").
			Optional(),
		field.Int("red91to105Total").
			Default(0).
			Optional(),
		field.String("red91to105Percentage").
			Default("0%").
			Optional(),
		field.Int("red106To120Total").
			Default(0).
			Optional(),
		field.String("red106To120Percentage").
			Default("0%").
			Optional(),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the TSCards.
func (TSCards) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("cards_stats").
			Unique().
			Required(),
	}
}
