// Code generated by ent, DO NOT EDIT.

package squad

import (
	"mapeleven/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Squad {
	return predicate.Squad(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Squad {
	return predicate.Squad(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Squad {
	return predicate.Squad(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Squad {
	return predicate.Squad(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Squad {
	return predicate.Squad(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Squad {
	return predicate.Squad(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Squad {
	return predicate.Squad(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Squad {
	return predicate.Squad(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Squad {
	return predicate.Squad(sql.FieldLTE(FieldID, id))
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.Squad {
	return predicate.Squad(sql.FieldEQ(FieldPosition, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int) predicate.Squad {
	return predicate.Squad(sql.FieldEQ(FieldNumber, v))
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.Squad {
	return predicate.Squad(sql.FieldEQ(FieldPosition, v))
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.Squad {
	return predicate.Squad(sql.FieldNEQ(FieldPosition, v))
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.Squad {
	return predicate.Squad(sql.FieldIn(FieldPosition, vs...))
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.Squad {
	return predicate.Squad(sql.FieldNotIn(FieldPosition, vs...))
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.Squad {
	return predicate.Squad(sql.FieldGT(FieldPosition, v))
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.Squad {
	return predicate.Squad(sql.FieldGTE(FieldPosition, v))
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.Squad {
	return predicate.Squad(sql.FieldLT(FieldPosition, v))
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.Squad {
	return predicate.Squad(sql.FieldLTE(FieldPosition, v))
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.Squad {
	return predicate.Squad(sql.FieldContains(FieldPosition, v))
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.Squad {
	return predicate.Squad(sql.FieldHasPrefix(FieldPosition, v))
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.Squad {
	return predicate.Squad(sql.FieldHasSuffix(FieldPosition, v))
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.Squad {
	return predicate.Squad(sql.FieldEqualFold(FieldPosition, v))
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.Squad {
	return predicate.Squad(sql.FieldContainsFold(FieldPosition, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int) predicate.Squad {
	return predicate.Squad(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int) predicate.Squad {
	return predicate.Squad(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int) predicate.Squad {
	return predicate.Squad(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int) predicate.Squad {
	return predicate.Squad(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int) predicate.Squad {
	return predicate.Squad(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int) predicate.Squad {
	return predicate.Squad(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int) predicate.Squad {
	return predicate.Squad(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int) predicate.Squad {
	return predicate.Squad(sql.FieldLTE(FieldNumber, v))
}

// HasPlayer applies the HasEdge predicate on the "player" edge.
func HasPlayer() predicate.Squad {
	return predicate.Squad(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlayerTable, PlayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayerWith applies the HasEdge predicate on the "player" edge with a given conditions (other predicates).
func HasPlayerWith(preds ...predicate.Player) predicate.Squad {
	return predicate.Squad(func(s *sql.Selector) {
		step := newPlayerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.Squad {
	return predicate.Squad(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeamTable, TeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Team) predicate.Squad {
	return predicate.Squad(func(s *sql.Selector) {
		step := newTeamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Squad) predicate.Squad {
	return predicate.Squad(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Squad) predicate.Squad {
	return predicate.Squad(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Squad) predicate.Squad {
	return predicate.Squad(func(s *sql.Selector) {
		p(s.Not())
	})
}
