package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PSPenalty holds the schema definition for the PSPenalty entity.
type PSPenalty struct {
	ent.Schema
}

// Fields of the PSPenalty.
func (PSPenalty) Fields() []ent.Field {
	return []ent.Field{
		field.Int("FoulsDrawn"),
		field.Int("FoulsCommitted"),
		field.Int("CardsYellow"),
		field.Int("CardYellowRed"),
		field.Int("CardsRed"),
		field.Int("PenaltyWon").Default(0),
		field.Int("PenaltyCommitted").Default(0),
		field.Int("PenaltyScored").Default(0),
		field.Int("PenaltyMissed"),
		field.Int("PenaltySaved").Default(0),
	}
}

// Edges of the PSPenalty.
func (PSPenalty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).
			Ref("pspenalty").
			Unique(),
	}
}
