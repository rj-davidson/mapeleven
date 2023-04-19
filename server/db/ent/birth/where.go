// Code generated by ent, DO NOT EDIT.

package birth

import (
	"mapeleven/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Birth {
	return predicate.Birth(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Birth {
	return predicate.Birth(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Birth {
	return predicate.Birth(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Birth {
	return predicate.Birth(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Birth {
	return predicate.Birth(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Birth {
	return predicate.Birth(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Birth {
	return predicate.Birth(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Birth {
	return predicate.Birth(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Birth {
	return predicate.Birth(sql.FieldLTE(FieldID, id))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldEQ(FieldDate, v))
}

// Place applies equality check predicate on the "place" field. It's identical to PlaceEQ.
func Place(v string) predicate.Birth {
	return predicate.Birth(sql.FieldEQ(FieldPlace, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Birth {
	return predicate.Birth(sql.FieldEQ(FieldCountry, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Birth {
	return predicate.Birth(sql.FieldLTE(FieldDate, v))
}

// PlaceEQ applies the EQ predicate on the "place" field.
func PlaceEQ(v string) predicate.Birth {
	return predicate.Birth(sql.FieldEQ(FieldPlace, v))
}

// PlaceNEQ applies the NEQ predicate on the "place" field.
func PlaceNEQ(v string) predicate.Birth {
	return predicate.Birth(sql.FieldNEQ(FieldPlace, v))
}

// PlaceIn applies the In predicate on the "place" field.
func PlaceIn(vs ...string) predicate.Birth {
	return predicate.Birth(sql.FieldIn(FieldPlace, vs...))
}

// PlaceNotIn applies the NotIn predicate on the "place" field.
func PlaceNotIn(vs ...string) predicate.Birth {
	return predicate.Birth(sql.FieldNotIn(FieldPlace, vs...))
}

// PlaceGT applies the GT predicate on the "place" field.
func PlaceGT(v string) predicate.Birth {
	return predicate.Birth(sql.FieldGT(FieldPlace, v))
}

// PlaceGTE applies the GTE predicate on the "place" field.
func PlaceGTE(v string) predicate.Birth {
	return predicate.Birth(sql.FieldGTE(FieldPlace, v))
}

// PlaceLT applies the LT predicate on the "place" field.
func PlaceLT(v string) predicate.Birth {
	return predicate.Birth(sql.FieldLT(FieldPlace, v))
}

// PlaceLTE applies the LTE predicate on the "place" field.
func PlaceLTE(v string) predicate.Birth {
	return predicate.Birth(sql.FieldLTE(FieldPlace, v))
}

// PlaceContains applies the Contains predicate on the "place" field.
func PlaceContains(v string) predicate.Birth {
	return predicate.Birth(sql.FieldContains(FieldPlace, v))
}

// PlaceHasPrefix applies the HasPrefix predicate on the "place" field.
func PlaceHasPrefix(v string) predicate.Birth {
	return predicate.Birth(sql.FieldHasPrefix(FieldPlace, v))
}

// PlaceHasSuffix applies the HasSuffix predicate on the "place" field.
func PlaceHasSuffix(v string) predicate.Birth {
	return predicate.Birth(sql.FieldHasSuffix(FieldPlace, v))
}

// PlaceEqualFold applies the EqualFold predicate on the "place" field.
func PlaceEqualFold(v string) predicate.Birth {
	return predicate.Birth(sql.FieldEqualFold(FieldPlace, v))
}

// PlaceContainsFold applies the ContainsFold predicate on the "place" field.
func PlaceContainsFold(v string) predicate.Birth {
	return predicate.Birth(sql.FieldContainsFold(FieldPlace, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Birth {
	return predicate.Birth(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Birth {
	return predicate.Birth(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Birth {
	return predicate.Birth(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Birth {
	return predicate.Birth(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Birth {
	return predicate.Birth(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Birth {
	return predicate.Birth(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Birth {
	return predicate.Birth(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Birth {
	return predicate.Birth(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Birth {
	return predicate.Birth(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Birth {
	return predicate.Birth(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Birth {
	return predicate.Birth(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Birth {
	return predicate.Birth(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Birth {
	return predicate.Birth(sql.FieldContainsFold(FieldCountry, v))
}

// HasPlayer applies the HasEdge predicate on the "player" edge.
func HasPlayer() predicate.Birth {
	return predicate.Birth(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PlayerTable, PlayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayerWith applies the HasEdge predicate on the "player" edge with a given conditions (other predicates).
func HasPlayerWith(preds ...predicate.Player) predicate.Birth {
	return predicate.Birth(func(s *sql.Selector) {
		step := newPlayerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Birth) predicate.Birth {
	return predicate.Birth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Birth) predicate.Birth {
	return predicate.Birth(func(s *sql.Selector) {
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
func Not(p predicate.Birth) predicate.Birth {
	return predicate.Birth(func(s *sql.Selector) {
		p(s.Not())
	})
}
