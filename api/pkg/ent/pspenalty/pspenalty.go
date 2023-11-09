// Code generated by ent, DO NOT EDIT.

package pspenalty

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the pspenalty type in the database.
	Label = "ps_penalty"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFoulsDrawn holds the string denoting the foulsdrawn field in the database.
	FieldFoulsDrawn = "fouls_drawn"
	// FieldFoulsCommitted holds the string denoting the foulscommitted field in the database.
	FieldFoulsCommitted = "fouls_committed"
	// FieldCardsYellow holds the string denoting the cardsyellow field in the database.
	FieldCardsYellow = "cards_yellow"
	// FieldCardYellowRed holds the string denoting the cardyellowred field in the database.
	FieldCardYellowRed = "card_yellow_red"
	// FieldCardsRed holds the string denoting the cardsred field in the database.
	FieldCardsRed = "cards_red"
	// FieldPenaltyWon holds the string denoting the penaltywon field in the database.
	FieldPenaltyWon = "penalty_won"
	// FieldPenaltyCommitted holds the string denoting the penaltycommitted field in the database.
	FieldPenaltyCommitted = "penalty_committed"
	// FieldPenaltyScored holds the string denoting the penaltyscored field in the database.
	FieldPenaltyScored = "penalty_scored"
	// FieldPenaltyMissed holds the string denoting the penaltymissed field in the database.
	FieldPenaltyMissed = "penalty_missed"
	// FieldPenaltySaved holds the string denoting the penaltysaved field in the database.
	FieldPenaltySaved = "penalty_saved"
	// EdgePlayerStats holds the string denoting the playerstats edge name in mutations.
	EdgePlayerStats = "playerStats"
	// Table holds the table name of the pspenalty in the database.
	Table = "ps_penalties"
	// PlayerStatsTable is the table that holds the playerStats relation/edge.
	PlayerStatsTable = "ps_penalties"
	// PlayerStatsInverseTable is the table name for the PlayerStats entity.
	// It exists in this package in order to avoid circular dependency with the "playerstats" package.
	PlayerStatsInverseTable = "player_stats"
	// PlayerStatsColumn is the table column denoting the playerStats relation/edge.
	PlayerStatsColumn = "player_stats_pspenalty"
)

// Columns holds all SQL columns for pspenalty fields.
var Columns = []string{
	FieldID,
	FieldFoulsDrawn,
	FieldFoulsCommitted,
	FieldCardsYellow,
	FieldCardYellowRed,
	FieldCardsRed,
	FieldPenaltyWon,
	FieldPenaltyCommitted,
	FieldPenaltyScored,
	FieldPenaltyMissed,
	FieldPenaltySaved,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ps_penalties"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"player_stats_pspenalty",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPenaltyWon holds the default value on creation for the "PenaltyWon" field.
	DefaultPenaltyWon int
	// DefaultPenaltyCommitted holds the default value on creation for the "PenaltyCommitted" field.
	DefaultPenaltyCommitted int
	// DefaultPenaltyScored holds the default value on creation for the "PenaltyScored" field.
	DefaultPenaltyScored int
	// DefaultPenaltySaved holds the default value on creation for the "PenaltySaved" field.
	DefaultPenaltySaved int
)

// OrderOption defines the ordering options for the PSPenalty queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFoulsDrawn orders the results by the FoulsDrawn field.
func ByFoulsDrawn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFoulsDrawn, opts...).ToFunc()
}

// ByFoulsCommitted orders the results by the FoulsCommitted field.
func ByFoulsCommitted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFoulsCommitted, opts...).ToFunc()
}

// ByCardsYellow orders the results by the CardsYellow field.
func ByCardsYellow(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCardsYellow, opts...).ToFunc()
}

// ByCardYellowRed orders the results by the CardYellowRed field.
func ByCardYellowRed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCardYellowRed, opts...).ToFunc()
}

// ByCardsRed orders the results by the CardsRed field.
func ByCardsRed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCardsRed, opts...).ToFunc()
}

// ByPenaltyWon orders the results by the PenaltyWon field.
func ByPenaltyWon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPenaltyWon, opts...).ToFunc()
}

// ByPenaltyCommitted orders the results by the PenaltyCommitted field.
func ByPenaltyCommitted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPenaltyCommitted, opts...).ToFunc()
}

// ByPenaltyScored orders the results by the PenaltyScored field.
func ByPenaltyScored(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPenaltyScored, opts...).ToFunc()
}

// ByPenaltyMissed orders the results by the PenaltyMissed field.
func ByPenaltyMissed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPenaltyMissed, opts...).ToFunc()
}

// ByPenaltySaved orders the results by the PenaltySaved field.
func ByPenaltySaved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPenaltySaved, opts...).ToFunc()
}

// ByPlayerStatsField orders the results by playerStats field.
func ByPlayerStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayerStatsStep(), sql.OrderByField(field, opts...))
	}
}
func newPlayerStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayerStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlayerStatsTable, PlayerStatsColumn),
	)
}
