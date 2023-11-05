// Code generated by ent, DO NOT EDIT.

package psgames

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLTE(FieldID, id))
}

// Appearences applies equality check predicate on the "appearences" field. It's identical to AppearencesEQ.
func Appearences(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldAppearences, v))
}

// Lineups applies equality check predicate on the "lineups" field. It's identical to LineupsEQ.
func Lineups(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldLineups, v))
}

// Minutes applies equality check predicate on the "minutes" field. It's identical to MinutesEQ.
func Minutes(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldMinutes, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldNumber, v))
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldPosition, v))
}

// Rating applies equality check predicate on the "rating" field. It's identical to RatingEQ.
func Rating(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldRating, v))
}

// Captain applies equality check predicate on the "captain" field. It's identical to CaptainEQ.
func Captain(v bool) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldCaptain, v))
}

// AppearencesEQ applies the EQ predicate on the "appearences" field.
func AppearencesEQ(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldAppearences, v))
}

// AppearencesNEQ applies the NEQ predicate on the "appearences" field.
func AppearencesNEQ(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNEQ(FieldAppearences, v))
}

// AppearencesIn applies the In predicate on the "appearences" field.
func AppearencesIn(vs ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldIn(FieldAppearences, vs...))
}

// AppearencesNotIn applies the NotIn predicate on the "appearences" field.
func AppearencesNotIn(vs ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNotIn(FieldAppearences, vs...))
}

// AppearencesGT applies the GT predicate on the "appearences" field.
func AppearencesGT(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGT(FieldAppearences, v))
}

// AppearencesGTE applies the GTE predicate on the "appearences" field.
func AppearencesGTE(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGTE(FieldAppearences, v))
}

// AppearencesLT applies the LT predicate on the "appearences" field.
func AppearencesLT(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLT(FieldAppearences, v))
}

// AppearencesLTE applies the LTE predicate on the "appearences" field.
func AppearencesLTE(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLTE(FieldAppearences, v))
}

// LineupsEQ applies the EQ predicate on the "lineups" field.
func LineupsEQ(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldLineups, v))
}

// LineupsNEQ applies the NEQ predicate on the "lineups" field.
func LineupsNEQ(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNEQ(FieldLineups, v))
}

// LineupsIn applies the In predicate on the "lineups" field.
func LineupsIn(vs ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldIn(FieldLineups, vs...))
}

// LineupsNotIn applies the NotIn predicate on the "lineups" field.
func LineupsNotIn(vs ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNotIn(FieldLineups, vs...))
}

// LineupsGT applies the GT predicate on the "lineups" field.
func LineupsGT(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGT(FieldLineups, v))
}

// LineupsGTE applies the GTE predicate on the "lineups" field.
func LineupsGTE(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGTE(FieldLineups, v))
}

// LineupsLT applies the LT predicate on the "lineups" field.
func LineupsLT(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLT(FieldLineups, v))
}

// LineupsLTE applies the LTE predicate on the "lineups" field.
func LineupsLTE(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLTE(FieldLineups, v))
}

// MinutesEQ applies the EQ predicate on the "minutes" field.
func MinutesEQ(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldMinutes, v))
}

// MinutesNEQ applies the NEQ predicate on the "minutes" field.
func MinutesNEQ(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNEQ(FieldMinutes, v))
}

// MinutesIn applies the In predicate on the "minutes" field.
func MinutesIn(vs ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldIn(FieldMinutes, vs...))
}

// MinutesNotIn applies the NotIn predicate on the "minutes" field.
func MinutesNotIn(vs ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNotIn(FieldMinutes, vs...))
}

// MinutesGT applies the GT predicate on the "minutes" field.
func MinutesGT(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGT(FieldMinutes, v))
}

// MinutesGTE applies the GTE predicate on the "minutes" field.
func MinutesGTE(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGTE(FieldMinutes, v))
}

// MinutesLT applies the LT predicate on the "minutes" field.
func MinutesLT(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLT(FieldMinutes, v))
}

// MinutesLTE applies the LTE predicate on the "minutes" field.
func MinutesLTE(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLTE(FieldMinutes, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int) predicate.PSGames {
	return predicate.PSGames(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int) predicate.PSGames {
	return predicate.PSGames(sql.FieldLTE(FieldNumber, v))
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldPosition, v))
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldNEQ(FieldPosition, v))
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.PSGames {
	return predicate.PSGames(sql.FieldIn(FieldPosition, vs...))
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.PSGames {
	return predicate.PSGames(sql.FieldNotIn(FieldPosition, vs...))
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldGT(FieldPosition, v))
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldGTE(FieldPosition, v))
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldLT(FieldPosition, v))
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldLTE(FieldPosition, v))
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldContains(FieldPosition, v))
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldHasPrefix(FieldPosition, v))
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldHasSuffix(FieldPosition, v))
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldEqualFold(FieldPosition, v))
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldContainsFold(FieldPosition, v))
}

// RatingEQ applies the EQ predicate on the "rating" field.
func RatingEQ(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldRating, v))
}

// RatingNEQ applies the NEQ predicate on the "rating" field.
func RatingNEQ(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldNEQ(FieldRating, v))
}

// RatingIn applies the In predicate on the "rating" field.
func RatingIn(vs ...string) predicate.PSGames {
	return predicate.PSGames(sql.FieldIn(FieldRating, vs...))
}

// RatingNotIn applies the NotIn predicate on the "rating" field.
func RatingNotIn(vs ...string) predicate.PSGames {
	return predicate.PSGames(sql.FieldNotIn(FieldRating, vs...))
}

// RatingGT applies the GT predicate on the "rating" field.
func RatingGT(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldGT(FieldRating, v))
}

// RatingGTE applies the GTE predicate on the "rating" field.
func RatingGTE(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldGTE(FieldRating, v))
}

// RatingLT applies the LT predicate on the "rating" field.
func RatingLT(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldLT(FieldRating, v))
}

// RatingLTE applies the LTE predicate on the "rating" field.
func RatingLTE(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldLTE(FieldRating, v))
}

// RatingContains applies the Contains predicate on the "rating" field.
func RatingContains(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldContains(FieldRating, v))
}

// RatingHasPrefix applies the HasPrefix predicate on the "rating" field.
func RatingHasPrefix(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldHasPrefix(FieldRating, v))
}

// RatingHasSuffix applies the HasSuffix predicate on the "rating" field.
func RatingHasSuffix(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldHasSuffix(FieldRating, v))
}

// RatingEqualFold applies the EqualFold predicate on the "rating" field.
func RatingEqualFold(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldEqualFold(FieldRating, v))
}

// RatingContainsFold applies the ContainsFold predicate on the "rating" field.
func RatingContainsFold(v string) predicate.PSGames {
	return predicate.PSGames(sql.FieldContainsFold(FieldRating, v))
}

// CaptainEQ applies the EQ predicate on the "captain" field.
func CaptainEQ(v bool) predicate.PSGames {
	return predicate.PSGames(sql.FieldEQ(FieldCaptain, v))
}

// CaptainNEQ applies the NEQ predicate on the "captain" field.
func CaptainNEQ(v bool) predicate.PSGames {
	return predicate.PSGames(sql.FieldNEQ(FieldCaptain, v))
}

// HasPlayer applies the HasEdge predicate on the "player" edge.
func HasPlayer() predicate.PSGames {
	return predicate.PSGames(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlayerTable, PlayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayerWith applies the HasEdge predicate on the "player" edge with a given conditions (other predicates).
func HasPlayerWith(preds ...predicate.Player) predicate.PSGames {
	return predicate.PSGames(func(s *sql.Selector) {
		step := newPlayerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PSGames) predicate.PSGames {
	return predicate.PSGames(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PSGames) predicate.PSGames {
	return predicate.PSGames(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PSGames) predicate.PSGames {
	return predicate.PSGames(sql.NotPredicates(p))
}
