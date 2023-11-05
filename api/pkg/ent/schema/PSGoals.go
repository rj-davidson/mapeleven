// schemas/psgoals.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PSGoals holds the schema definition for the PSGoals entity.
type PSGoals struct {
	ent.Schema
}

// Fields of the PSGoals.
func (PSGoals) Fields() []ent.Field {
	return []ent.Field{
		field.Int("totalGoals"),
		field.Int("concededGoals"),
		field.Int("assistGoals"),
		field.Int("saveGoals").Default(0),
		field.Int("shotsTotal"),
		field.Int("shotsOn"),
	}
}

// Edges of the PSGoals.
func (PSGoals) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).Ref("psgoals").Unique(),
	}
}
