package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Achievements holds the schema definition for the Achievements entity.
type Achievements struct {
	ent.Schema
}

// Fields of the Achievements.
func (Achievements) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
		field.String("recipient"),
		field.String("award"),
		field.Int("year"),
	}
}

// Edges of the Achievements.
func (Achievements) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("achievements").
			Unique(),
		edge.From("coach", Coach.Type).
			Ref("achievements").
			Unique(),
		edge.From("player", Player.Type).
			Ref("achievements").
			Unique(),
	}
}
