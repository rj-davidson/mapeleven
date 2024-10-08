// Code generated by ent, DO NOT EDIT.

package pstechnical

import (
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLTE(FieldID, id))
}

// FoulsDrawn applies equality check predicate on the "FoulsDrawn" field. It's identical to FoulsDrawnEQ.
func FoulsDrawn(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldFoulsDrawn, v))
}

// DribbleAttempts applies equality check predicate on the "DribbleAttempts" field. It's identical to DribbleAttemptsEQ.
func DribbleAttempts(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldDribbleAttempts, v))
}

// DribbleSuccess applies equality check predicate on the "DribbleSuccess" field. It's identical to DribbleSuccessEQ.
func DribbleSuccess(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldDribbleSuccess, v))
}

// PassesTotal applies equality check predicate on the "PassesTotal" field. It's identical to PassesTotalEQ.
func PassesTotal(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldPassesTotal, v))
}

// PassesKey applies equality check predicate on the "PassesKey" field. It's identical to PassesKeyEQ.
func PassesKey(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldPassesKey, v))
}

// PassesAccuracy applies equality check predicate on the "PassesAccuracy" field. It's identical to PassesAccuracyEQ.
func PassesAccuracy(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldPassesAccuracy, v))
}

// LastUpdated applies equality check predicate on the "lastUpdated" field. It's identical to LastUpdatedEQ.
func LastUpdated(v time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldLastUpdated, v))
}

// FoulsDrawnEQ applies the EQ predicate on the "FoulsDrawn" field.
func FoulsDrawnEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldFoulsDrawn, v))
}

// FoulsDrawnNEQ applies the NEQ predicate on the "FoulsDrawn" field.
func FoulsDrawnNEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNEQ(FieldFoulsDrawn, v))
}

// FoulsDrawnIn applies the In predicate on the "FoulsDrawn" field.
func FoulsDrawnIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIn(FieldFoulsDrawn, vs...))
}

// FoulsDrawnNotIn applies the NotIn predicate on the "FoulsDrawn" field.
func FoulsDrawnNotIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotIn(FieldFoulsDrawn, vs...))
}

// FoulsDrawnGT applies the GT predicate on the "FoulsDrawn" field.
func FoulsDrawnGT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGT(FieldFoulsDrawn, v))
}

// FoulsDrawnGTE applies the GTE predicate on the "FoulsDrawn" field.
func FoulsDrawnGTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGTE(FieldFoulsDrawn, v))
}

// FoulsDrawnLT applies the LT predicate on the "FoulsDrawn" field.
func FoulsDrawnLT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLT(FieldFoulsDrawn, v))
}

// FoulsDrawnLTE applies the LTE predicate on the "FoulsDrawn" field.
func FoulsDrawnLTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLTE(FieldFoulsDrawn, v))
}

// DribbleAttemptsEQ applies the EQ predicate on the "DribbleAttempts" field.
func DribbleAttemptsEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldDribbleAttempts, v))
}

// DribbleAttemptsNEQ applies the NEQ predicate on the "DribbleAttempts" field.
func DribbleAttemptsNEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNEQ(FieldDribbleAttempts, v))
}

// DribbleAttemptsIn applies the In predicate on the "DribbleAttempts" field.
func DribbleAttemptsIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIn(FieldDribbleAttempts, vs...))
}

// DribbleAttemptsNotIn applies the NotIn predicate on the "DribbleAttempts" field.
func DribbleAttemptsNotIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotIn(FieldDribbleAttempts, vs...))
}

// DribbleAttemptsGT applies the GT predicate on the "DribbleAttempts" field.
func DribbleAttemptsGT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGT(FieldDribbleAttempts, v))
}

// DribbleAttemptsGTE applies the GTE predicate on the "DribbleAttempts" field.
func DribbleAttemptsGTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGTE(FieldDribbleAttempts, v))
}

// DribbleAttemptsLT applies the LT predicate on the "DribbleAttempts" field.
func DribbleAttemptsLT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLT(FieldDribbleAttempts, v))
}

// DribbleAttemptsLTE applies the LTE predicate on the "DribbleAttempts" field.
func DribbleAttemptsLTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLTE(FieldDribbleAttempts, v))
}

// DribbleSuccessEQ applies the EQ predicate on the "DribbleSuccess" field.
func DribbleSuccessEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldDribbleSuccess, v))
}

// DribbleSuccessNEQ applies the NEQ predicate on the "DribbleSuccess" field.
func DribbleSuccessNEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNEQ(FieldDribbleSuccess, v))
}

// DribbleSuccessIn applies the In predicate on the "DribbleSuccess" field.
func DribbleSuccessIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIn(FieldDribbleSuccess, vs...))
}

// DribbleSuccessNotIn applies the NotIn predicate on the "DribbleSuccess" field.
func DribbleSuccessNotIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotIn(FieldDribbleSuccess, vs...))
}

// DribbleSuccessGT applies the GT predicate on the "DribbleSuccess" field.
func DribbleSuccessGT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGT(FieldDribbleSuccess, v))
}

// DribbleSuccessGTE applies the GTE predicate on the "DribbleSuccess" field.
func DribbleSuccessGTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGTE(FieldDribbleSuccess, v))
}

// DribbleSuccessLT applies the LT predicate on the "DribbleSuccess" field.
func DribbleSuccessLT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLT(FieldDribbleSuccess, v))
}

// DribbleSuccessLTE applies the LTE predicate on the "DribbleSuccess" field.
func DribbleSuccessLTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLTE(FieldDribbleSuccess, v))
}

// PassesTotalEQ applies the EQ predicate on the "PassesTotal" field.
func PassesTotalEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldPassesTotal, v))
}

// PassesTotalNEQ applies the NEQ predicate on the "PassesTotal" field.
func PassesTotalNEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNEQ(FieldPassesTotal, v))
}

// PassesTotalIn applies the In predicate on the "PassesTotal" field.
func PassesTotalIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIn(FieldPassesTotal, vs...))
}

// PassesTotalNotIn applies the NotIn predicate on the "PassesTotal" field.
func PassesTotalNotIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotIn(FieldPassesTotal, vs...))
}

// PassesTotalGT applies the GT predicate on the "PassesTotal" field.
func PassesTotalGT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGT(FieldPassesTotal, v))
}

// PassesTotalGTE applies the GTE predicate on the "PassesTotal" field.
func PassesTotalGTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGTE(FieldPassesTotal, v))
}

// PassesTotalLT applies the LT predicate on the "PassesTotal" field.
func PassesTotalLT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLT(FieldPassesTotal, v))
}

// PassesTotalLTE applies the LTE predicate on the "PassesTotal" field.
func PassesTotalLTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLTE(FieldPassesTotal, v))
}

// PassesKeyEQ applies the EQ predicate on the "PassesKey" field.
func PassesKeyEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldPassesKey, v))
}

// PassesKeyNEQ applies the NEQ predicate on the "PassesKey" field.
func PassesKeyNEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNEQ(FieldPassesKey, v))
}

// PassesKeyIn applies the In predicate on the "PassesKey" field.
func PassesKeyIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIn(FieldPassesKey, vs...))
}

// PassesKeyNotIn applies the NotIn predicate on the "PassesKey" field.
func PassesKeyNotIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotIn(FieldPassesKey, vs...))
}

// PassesKeyGT applies the GT predicate on the "PassesKey" field.
func PassesKeyGT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGT(FieldPassesKey, v))
}

// PassesKeyGTE applies the GTE predicate on the "PassesKey" field.
func PassesKeyGTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGTE(FieldPassesKey, v))
}

// PassesKeyLT applies the LT predicate on the "PassesKey" field.
func PassesKeyLT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLT(FieldPassesKey, v))
}

// PassesKeyLTE applies the LTE predicate on the "PassesKey" field.
func PassesKeyLTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLTE(FieldPassesKey, v))
}

// PassesAccuracyEQ applies the EQ predicate on the "PassesAccuracy" field.
func PassesAccuracyEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldPassesAccuracy, v))
}

// PassesAccuracyNEQ applies the NEQ predicate on the "PassesAccuracy" field.
func PassesAccuracyNEQ(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNEQ(FieldPassesAccuracy, v))
}

// PassesAccuracyIn applies the In predicate on the "PassesAccuracy" field.
func PassesAccuracyIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIn(FieldPassesAccuracy, vs...))
}

// PassesAccuracyNotIn applies the NotIn predicate on the "PassesAccuracy" field.
func PassesAccuracyNotIn(vs ...int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotIn(FieldPassesAccuracy, vs...))
}

// PassesAccuracyGT applies the GT predicate on the "PassesAccuracy" field.
func PassesAccuracyGT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGT(FieldPassesAccuracy, v))
}

// PassesAccuracyGTE applies the GTE predicate on the "PassesAccuracy" field.
func PassesAccuracyGTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGTE(FieldPassesAccuracy, v))
}

// PassesAccuracyLT applies the LT predicate on the "PassesAccuracy" field.
func PassesAccuracyLT(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLT(FieldPassesAccuracy, v))
}

// PassesAccuracyLTE applies the LTE predicate on the "PassesAccuracy" field.
func PassesAccuracyLTE(v int) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLTE(FieldPassesAccuracy, v))
}

// LastUpdatedEQ applies the EQ predicate on the "lastUpdated" field.
func LastUpdatedEQ(v time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldEQ(FieldLastUpdated, v))
}

// LastUpdatedNEQ applies the NEQ predicate on the "lastUpdated" field.
func LastUpdatedNEQ(v time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNEQ(FieldLastUpdated, v))
}

// LastUpdatedIn applies the In predicate on the "lastUpdated" field.
func LastUpdatedIn(vs ...time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIn(FieldLastUpdated, vs...))
}

// LastUpdatedNotIn applies the NotIn predicate on the "lastUpdated" field.
func LastUpdatedNotIn(vs ...time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotIn(FieldLastUpdated, vs...))
}

// LastUpdatedGT applies the GT predicate on the "lastUpdated" field.
func LastUpdatedGT(v time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGT(FieldLastUpdated, v))
}

// LastUpdatedGTE applies the GTE predicate on the "lastUpdated" field.
func LastUpdatedGTE(v time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldGTE(FieldLastUpdated, v))
}

// LastUpdatedLT applies the LT predicate on the "lastUpdated" field.
func LastUpdatedLT(v time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLT(FieldLastUpdated, v))
}

// LastUpdatedLTE applies the LTE predicate on the "lastUpdated" field.
func LastUpdatedLTE(v time.Time) predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldLTE(FieldLastUpdated, v))
}

// LastUpdatedIsNil applies the IsNil predicate on the "lastUpdated" field.
func LastUpdatedIsNil() predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldIsNull(FieldLastUpdated))
}

// LastUpdatedNotNil applies the NotNil predicate on the "lastUpdated" field.
func LastUpdatedNotNil() predicate.PSTechnical {
	return predicate.PSTechnical(sql.FieldNotNull(FieldLastUpdated))
}

// HasPlayerStats applies the HasEdge predicate on the "playerStats" edge.
func HasPlayerStats() predicate.PSTechnical {
	return predicate.PSTechnical(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PlayerStatsTable, PlayerStatsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayerStatsWith applies the HasEdge predicate on the "playerStats" edge with a given conditions (other predicates).
func HasPlayerStatsWith(preds ...predicate.PlayerStats) predicate.PSTechnical {
	return predicate.PSTechnical(func(s *sql.Selector) {
		step := newPlayerStatsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PSTechnical) predicate.PSTechnical {
	return predicate.PSTechnical(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PSTechnical) predicate.PSTechnical {
	return predicate.PSTechnical(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PSTechnical) predicate.PSTechnical {
	return predicate.PSTechnical(sql.NotPredicates(p))
}
