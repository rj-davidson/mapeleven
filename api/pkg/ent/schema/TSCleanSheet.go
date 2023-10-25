package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TSCleanSheet holds the schema definition for the TSCleanSheet entity.
type TSCleanSheet struct {
	ent.Schema
}

// Fields of the TSCleanSheet.
func (TSCleanSheet) Fields() []ent.Field {
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

// Edges of the TSCleanSheet.
func (TSCleanSheet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("clean_sheet_stats").
			Unique().
			Required(),
	}
}
