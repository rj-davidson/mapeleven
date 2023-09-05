package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Standings holds the schema definition for the Standings entity.
type Standings struct {
	ent.Schema
}

// Fields of the Standings.
func (Standings) Fields() []ent.Field {
	return []ent.Field{
		field.Int("rank"),
		field.String("description"),
	}
}

// Edges of the Standings.
func (Standings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("league", League.Type).
			Ref("standings").
			Unique(),
		edge.From("team", Team.Type).
			Ref("standings").
			Unique(),
	}
}
