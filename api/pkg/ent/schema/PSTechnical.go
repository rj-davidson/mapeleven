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
		field.Int("FoulsDrawn").Default(0),
		field.Int("DribbleAttempts").Default(0),
		field.Int("DribbleSuccess").Default(0),
		field.Int("DribblePast").Default(0),
		field.Int("PassesTotal").Default(0),
		field.Int("PassesKey").Default(0),
		field.Int("PassesAccuracy").Default(0),
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
