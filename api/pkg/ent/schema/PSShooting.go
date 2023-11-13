package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PSShooting holds the schema definition for the PSShooting entity.
type PSShooting struct {
	ent.Schema
}

// Fields of the PSShooting.
func (PSShooting) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Goals").Default(0),
		field.Int("Conceded").Default(0),
		field.Int("Assists").Default(0),
		field.Int("Saves").Default(0),
		field.Int("Shots").Default(0),
		field.Int("OnTarget").Default(0),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the PSShooting.
func (PSShooting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("psShooting").
			Unique(),
	}
}
