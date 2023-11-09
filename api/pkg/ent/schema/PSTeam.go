package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PSDefense holds the schema definition for the PSDefense entity.
type PSTeam struct {
	ent.Schema
}

// Fields of the PSDefense.
func (PSTeam) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ApiFootballId"),
		field.String("Name"),
		field.String("Logo"),
	}
}

// Edges of the PSDefense.
func (PSTeam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("psteam").
			Unique(),
	}
}
