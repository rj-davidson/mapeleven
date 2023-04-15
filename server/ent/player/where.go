// Code generated by ent, DO NOT EDIT.

package player

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldName, v))
}

// Firstname applies equality check predicate on the "firstname" field. It's identical to FirstnameEQ.
func Firstname(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldFirstname, v))
}

// Lastname applies equality check predicate on the "lastname" field. It's identical to LastnameEQ.
func Lastname(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldLastname, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldAge, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v float64) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldHeight, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v float64) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldWeight, v))
}

// Injured applies equality check predicate on the "injured" field. It's identical to InjuredEQ.
func Injured(v bool) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldInjured, v))
}

// Photo applies equality check predicate on the "photo" field. It's identical to PhotoEQ.
func Photo(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldPhoto, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Player {
	return predicate.Player(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Player {
	return predicate.Player(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Player {
	return predicate.Player(sql.FieldContainsFold(FieldName, v))
}

// FirstnameEQ applies the EQ predicate on the "firstname" field.
func FirstnameEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldFirstname, v))
}

// FirstnameNEQ applies the NEQ predicate on the "firstname" field.
func FirstnameNEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldFirstname, v))
}

// FirstnameIn applies the In predicate on the "firstname" field.
func FirstnameIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldFirstname, vs...))
}

// FirstnameNotIn applies the NotIn predicate on the "firstname" field.
func FirstnameNotIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldFirstname, vs...))
}

// FirstnameGT applies the GT predicate on the "firstname" field.
func FirstnameGT(v string) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldFirstname, v))
}

// FirstnameGTE applies the GTE predicate on the "firstname" field.
func FirstnameGTE(v string) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldFirstname, v))
}

// FirstnameLT applies the LT predicate on the "firstname" field.
func FirstnameLT(v string) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldFirstname, v))
}

// FirstnameLTE applies the LTE predicate on the "firstname" field.
func FirstnameLTE(v string) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldFirstname, v))
}

// FirstnameContains applies the Contains predicate on the "firstname" field.
func FirstnameContains(v string) predicate.Player {
	return predicate.Player(sql.FieldContains(FieldFirstname, v))
}

// FirstnameHasPrefix applies the HasPrefix predicate on the "firstname" field.
func FirstnameHasPrefix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasPrefix(FieldFirstname, v))
}

// FirstnameHasSuffix applies the HasSuffix predicate on the "firstname" field.
func FirstnameHasSuffix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasSuffix(FieldFirstname, v))
}

// FirstnameEqualFold applies the EqualFold predicate on the "firstname" field.
func FirstnameEqualFold(v string) predicate.Player {
	return predicate.Player(sql.FieldEqualFold(FieldFirstname, v))
}

// FirstnameContainsFold applies the ContainsFold predicate on the "firstname" field.
func FirstnameContainsFold(v string) predicate.Player {
	return predicate.Player(sql.FieldContainsFold(FieldFirstname, v))
}

// LastnameEQ applies the EQ predicate on the "lastname" field.
func LastnameEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldLastname, v))
}

// LastnameNEQ applies the NEQ predicate on the "lastname" field.
func LastnameNEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldLastname, v))
}

// LastnameIn applies the In predicate on the "lastname" field.
func LastnameIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldLastname, vs...))
}

// LastnameNotIn applies the NotIn predicate on the "lastname" field.
func LastnameNotIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldLastname, vs...))
}

// LastnameGT applies the GT predicate on the "lastname" field.
func LastnameGT(v string) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldLastname, v))
}

// LastnameGTE applies the GTE predicate on the "lastname" field.
func LastnameGTE(v string) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldLastname, v))
}

// LastnameLT applies the LT predicate on the "lastname" field.
func LastnameLT(v string) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldLastname, v))
}

// LastnameLTE applies the LTE predicate on the "lastname" field.
func LastnameLTE(v string) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldLastname, v))
}

// LastnameContains applies the Contains predicate on the "lastname" field.
func LastnameContains(v string) predicate.Player {
	return predicate.Player(sql.FieldContains(FieldLastname, v))
}

// LastnameHasPrefix applies the HasPrefix predicate on the "lastname" field.
func LastnameHasPrefix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasPrefix(FieldLastname, v))
}

// LastnameHasSuffix applies the HasSuffix predicate on the "lastname" field.
func LastnameHasSuffix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasSuffix(FieldLastname, v))
}

// LastnameEqualFold applies the EqualFold predicate on the "lastname" field.
func LastnameEqualFold(v string) predicate.Player {
	return predicate.Player(sql.FieldEqualFold(FieldLastname, v))
}

// LastnameContainsFold applies the ContainsFold predicate on the "lastname" field.
func LastnameContainsFold(v string) predicate.Player {
	return predicate.Player(sql.FieldContainsFold(FieldLastname, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldAge, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v float64) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v float64) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...float64) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...float64) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v float64) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v float64) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v float64) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v float64) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldHeight, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v float64) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v float64) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...float64) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...float64) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v float64) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v float64) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v float64) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v float64) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldWeight, v))
}

// InjuredEQ applies the EQ predicate on the "injured" field.
func InjuredEQ(v bool) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldInjured, v))
}

// InjuredNEQ applies the NEQ predicate on the "injured" field.
func InjuredNEQ(v bool) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldInjured, v))
}

// PhotoEQ applies the EQ predicate on the "photo" field.
func PhotoEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldPhoto, v))
}

// PhotoNEQ applies the NEQ predicate on the "photo" field.
func PhotoNEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldPhoto, v))
}

// PhotoIn applies the In predicate on the "photo" field.
func PhotoIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldPhoto, vs...))
}

// PhotoNotIn applies the NotIn predicate on the "photo" field.
func PhotoNotIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldPhoto, vs...))
}

// PhotoGT applies the GT predicate on the "photo" field.
func PhotoGT(v string) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldPhoto, v))
}

// PhotoGTE applies the GTE predicate on the "photo" field.
func PhotoGTE(v string) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldPhoto, v))
}

// PhotoLT applies the LT predicate on the "photo" field.
func PhotoLT(v string) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldPhoto, v))
}

// PhotoLTE applies the LTE predicate on the "photo" field.
func PhotoLTE(v string) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldPhoto, v))
}

// PhotoContains applies the Contains predicate on the "photo" field.
func PhotoContains(v string) predicate.Player {
	return predicate.Player(sql.FieldContains(FieldPhoto, v))
}

// PhotoHasPrefix applies the HasPrefix predicate on the "photo" field.
func PhotoHasPrefix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasPrefix(FieldPhoto, v))
}

// PhotoHasSuffix applies the HasSuffix predicate on the "photo" field.
func PhotoHasSuffix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasSuffix(FieldPhoto, v))
}

// PhotoEqualFold applies the EqualFold predicate on the "photo" field.
func PhotoEqualFold(v string) predicate.Player {
	return predicate.Player(sql.FieldEqualFold(FieldPhoto, v))
}

// PhotoContainsFold applies the ContainsFold predicate on the "photo" field.
func PhotoContainsFold(v string) predicate.Player {
	return predicate.Player(sql.FieldContainsFold(FieldPhoto, v))
}

// HasBirth applies the HasEdge predicate on the "birth" edge.
func HasBirth() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BirthTable, BirthColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBirthWith applies the HasEdge predicate on the "birth" edge with a given conditions (other predicates).
func HasBirthWith(preds ...predicate.Birth) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newBirthStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNationality applies the HasEdge predicate on the "nationality" edge.
func HasNationality() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NationalityTable, NationalityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNationalityWith applies the HasEdge predicate on the "nationality" edge with a given conditions (other predicates).
func HasNationalityWith(preds ...predicate.Country) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newNationalityStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TeamsTable, TeamsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.Team) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newTeamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Player) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Player) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
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
func Not(p predicate.Player) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		p(s.Not())
	})
}
