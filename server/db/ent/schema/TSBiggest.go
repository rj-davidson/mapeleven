package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TSBiggest holds the schema definition for the TSBiggest entity.
type TSBiggest struct {
	ent.Schema
}

// Fields of the TSBiggest.
func (TSBiggest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("streakWins").
			Optional(),
		field.Int("streakLosses").
			Optional(),
		field.Int("streakDraws").
			Optional(),
		field.String("winsHome").
			Optional(),
		field.String("winsAway").
			Optional(),
		field.String("lossesHome").
			Optional(),
		field.String("lossesAway").
			Optional(),
		field.Int("goalsForHome").
			Optional(),
		field.Int("goalsForAway").
			Optional(),
		field.Int("goalsAgainstHome").
			Optional(),
		field.Int("goalsAgainstAway").
			Optional(),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the TSBiggest.
func (TSBiggest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("biggest_stats").
			Unique().
			Required(),
	}
}
