// Code generated by ent, DO NOT EDIT.

package playerstats

import (
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldLTE(FieldID, id))
}

// LastUpdated applies equality check predicate on the "lastUpdated" field. It's identical to LastUpdatedEQ.
func LastUpdated(v time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldEQ(FieldLastUpdated, v))
}

// LastUpdatedEQ applies the EQ predicate on the "lastUpdated" field.
func LastUpdatedEQ(v time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldEQ(FieldLastUpdated, v))
}

// LastUpdatedNEQ applies the NEQ predicate on the "lastUpdated" field.
func LastUpdatedNEQ(v time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldNEQ(FieldLastUpdated, v))
}

// LastUpdatedIn applies the In predicate on the "lastUpdated" field.
func LastUpdatedIn(vs ...time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldIn(FieldLastUpdated, vs...))
}

// LastUpdatedNotIn applies the NotIn predicate on the "lastUpdated" field.
func LastUpdatedNotIn(vs ...time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldNotIn(FieldLastUpdated, vs...))
}

// LastUpdatedGT applies the GT predicate on the "lastUpdated" field.
func LastUpdatedGT(v time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldGT(FieldLastUpdated, v))
}

// LastUpdatedGTE applies the GTE predicate on the "lastUpdated" field.
func LastUpdatedGTE(v time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldGTE(FieldLastUpdated, v))
}

// LastUpdatedLT applies the LT predicate on the "lastUpdated" field.
func LastUpdatedLT(v time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldLT(FieldLastUpdated, v))
}

// LastUpdatedLTE applies the LTE predicate on the "lastUpdated" field.
func LastUpdatedLTE(v time.Time) predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldLTE(FieldLastUpdated, v))
}

// LastUpdatedIsNil applies the IsNil predicate on the "lastUpdated" field.
func LastUpdatedIsNil() predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldIsNull(FieldLastUpdated))
}

// LastUpdatedNotNil applies the NotNil predicate on the "lastUpdated" field.
func LastUpdatedNotNil() predicate.PlayerStats {
	return predicate.PlayerStats(sql.FieldNotNull(FieldLastUpdated))
}

// HasPlayer applies the HasEdge predicate on the "player" edge.
func HasPlayer() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlayerTable, PlayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayerWith applies the HasEdge predicate on the "player" edge with a given conditions (other predicates).
func HasPlayerWith(preds ...predicate.Player) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPlayerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeamTable, TeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Team) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newTeamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlayerEvents applies the HasEdge predicate on the "playerEvents" edge.
func HasPlayerEvents() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PlayerEventsTable, PlayerEventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayerEventsWith applies the HasEdge predicate on the "playerEvents" edge with a given conditions (other predicates).
func HasPlayerEventsWith(preds ...predicate.FixtureEvents) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPlayerEventsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMatchPlayer applies the HasEdge predicate on the "matchPlayer" edge.
func HasMatchPlayer() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MatchPlayerTable, MatchPlayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchPlayerWith applies the HasEdge predicate on the "matchPlayer" edge with a given conditions (other predicates).
func HasMatchPlayerWith(preds ...predicate.MatchPlayer) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newMatchPlayerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssistEvents applies the HasEdge predicate on the "assistEvents" edge.
func HasAssistEvents() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssistEventsTable, AssistEventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssistEventsWith applies the HasEdge predicate on the "assistEvents" edge with a given conditions (other predicates).
func HasAssistEventsWith(preds ...predicate.FixtureEvents) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newAssistEventsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPsGames applies the HasEdge predicate on the "psGames" edge.
func HasPsGames() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PsGamesTable, PsGamesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPsGamesWith applies the HasEdge predicate on the "psGames" edge with a given conditions (other predicates).
func HasPsGamesWith(preds ...predicate.PSGames) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPsGamesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPsShooting applies the HasEdge predicate on the "psShooting" edge.
func HasPsShooting() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PsShootingTable, PsShootingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPsShootingWith applies the HasEdge predicate on the "psShooting" edge with a given conditions (other predicates).
func HasPsShootingWith(preds ...predicate.PSShooting) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPsShootingStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPsDefense applies the HasEdge predicate on the "psDefense" edge.
func HasPsDefense() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PsDefenseTable, PsDefenseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPsDefenseWith applies the HasEdge predicate on the "psDefense" edge with a given conditions (other predicates).
func HasPsDefenseWith(preds ...predicate.PSDefense) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPsDefenseStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPsTechnical applies the HasEdge predicate on the "psTechnical" edge.
func HasPsTechnical() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PsTechnicalTable, PsTechnicalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPsTechnicalWith applies the HasEdge predicate on the "psTechnical" edge with a given conditions (other predicates).
func HasPsTechnicalWith(preds ...predicate.PSTechnical) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPsTechnicalStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPsPenalty applies the HasEdge predicate on the "psPenalty" edge.
func HasPsPenalty() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PsPenaltyTable, PsPenaltyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPsPenaltyWith applies the HasEdge predicate on the "psPenalty" edge with a given conditions (other predicates).
func HasPsPenaltyWith(preds ...predicate.PSPenalty) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPsPenaltyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPsSubstitutes applies the HasEdge predicate on the "psSubstitutes" edge.
func HasPsSubstitutes() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PsSubstitutesTable, PsSubstitutesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPsSubstitutesWith applies the HasEdge predicate on the "psSubstitutes" edge with a given conditions (other predicates).
func HasPsSubstitutesWith(preds ...predicate.PSSubstitutes) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPsSubstitutesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSeason applies the HasEdge predicate on the "season" edge.
func HasSeason() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SeasonTable, SeasonPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSeasonWith applies the HasEdge predicate on the "season" edge with a given conditions (other predicates).
func HasSeasonWith(preds ...predicate.Season) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newSeasonStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPsFairplay applies the HasEdge predicate on the "psFairplay" edge.
func HasPsFairplay() predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PsFairplayTable, PsFairplayColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPsFairplayWith applies the HasEdge predicate on the "psFairplay" edge with a given conditions (other predicates).
func HasPsFairplayWith(preds ...predicate.PSFairplay) predicate.PlayerStats {
	return predicate.PlayerStats(func(s *sql.Selector) {
		step := newPsFairplayStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlayerStats) predicate.PlayerStats {
	return predicate.PlayerStats(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlayerStats) predicate.PlayerStats {
	return predicate.PlayerStats(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlayerStats) predicate.PlayerStats {
	return predicate.PlayerStats(sql.NotPredicates(p))
}
