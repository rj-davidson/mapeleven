package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
		field.Enum("type").
			Values("Goal", "Assist", "Card", "Substitution", "Other"),
		field.Enum("option").
			Values("Yellow", "Red", "OwnGoal", "In", "Out", "None"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("events").
			Unique().
			Required(),
		edge.From("player", Player.Type).
			Ref("events").
			Unique().
			Required(),
		//edge.From("player", Player.Type).
		//	Ref("events_player2").
		//	Unique(),
		edge.From("fixture", Fixture.Type).
			Ref("events").
			Unique().
			Required(),
	}
}
