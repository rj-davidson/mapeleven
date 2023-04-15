// Code generated by ent, DO NOT EDIT.

package season

import (
	"mapeleven/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldID, id))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldYear, v))
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldStart, v))
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldEnd, v))
}

// Current applies equality check predicate on the "current" field. It's identical to CurrentEQ.
func Current(v bool) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldCurrent, v))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldYear, v))
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldStart, v))
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldStart, v))
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldStart, vs...))
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldStart, vs...))
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldStart, v))
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldStart, v))
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldStart, v))
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldStart, v))
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldEnd, v))
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldEnd, v))
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldEnd, vs...))
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldEnd, vs...))
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldEnd, v))
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldEnd, v))
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldEnd, v))
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldEnd, v))
}

// CurrentEQ applies the EQ predicate on the "current" field.
func CurrentEQ(v bool) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldCurrent, v))
}

// CurrentNEQ applies the NEQ predicate on the "current" field.
func CurrentNEQ(v bool) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldCurrent, v))
}

// HasLeague applies the HasEdge predicate on the "league" edge.
func HasLeague() predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LeagueTable, LeagueColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLeagueWith applies the HasEdge predicate on the "league" edge with a given conditions (other predicates).
func HasLeagueWith(preds ...predicate.League) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		step := newLeagueStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Season) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Season) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
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
func Not(p predicate.Season) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		p(s.Not())
	})
}
