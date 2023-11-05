// schemas/pspenalty.go
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
		field.Int("foulsDrawn"),
		field.Int("foulsCommitted"),
		field.Int("cardsYellow"),
		field.Int("cardYellowred"),
		field.Int("cardsRed"),
		field.Int("penaltyWon").Default(0),
		field.Int("penaltyCommited").Default(0),
		field.Int("penaltyScored").Default(0),
		field.Int("penaltyMissed"),
		field.Int("penaltySaved").Default(0),
	}
}

// Edges of the PSPenalty.
func (PSPenalty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).Ref("pspenalty").Unique(),
	}
}
