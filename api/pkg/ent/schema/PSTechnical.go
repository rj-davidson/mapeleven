package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PSTechnical holds the schema definition for the PSTechnical entity.
type PSTechnical struct {
	ent.Schema
}

// Fields of the PSTechnical.
func (PSTechnical) Fields() []ent.Field {
	return []ent.Field{
		field.Int("FoulsDrawn"),
		field.Int("DribbleAttempts"),
		field.Int("DribbleSuccess"),
		field.Int("DribblePast").Default(0),
		field.Int("PassesTotal"),
		field.Int("PassesKey"),
		field.Int("PassesAccuracy"),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the PSTechnical.
func (PSTechnical) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("psTechnical").
			Unique(),
	}
}
