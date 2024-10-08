// Code generated by ent, DO NOT EDIT.

package psfairplay

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the psfairplay type in the database.
	Label = "ps_fairplay"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFoulsCommitted holds the string denoting the foulscommitted field in the database.
	FieldFoulsCommitted = "fouls_committed"
	// FieldYellow holds the string denoting the yellow field in the database.
	FieldYellow = "yellow"
	// FieldYellowRed holds the string denoting the yellowred field in the database.
	FieldYellowRed = "yellow_red"
	// FieldRed holds the string denoting the red field in the database.
	FieldRed = "red"
	// FieldPenaltyConceded holds the string denoting the penaltyconceded field in the database.
	FieldPenaltyConceded = "penalty_conceded"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgePlayerStats holds the string denoting the playerstats edge name in mutations.
	EdgePlayerStats = "playerStats"
	// Table holds the table name of the psfairplay in the database.
	Table = "ps_fairplays"
	// PlayerStatsTable is the table that holds the playerStats relation/edge.
	PlayerStatsTable = "ps_fairplays"
	// PlayerStatsInverseTable is the table name for the PlayerStats entity.
	// It exists in this package in order to avoid circular dependency with the "playerstats" package.
	PlayerStatsInverseTable = "player_stats"
	// PlayerStatsColumn is the table column denoting the playerStats relation/edge.
	PlayerStatsColumn = "player_stats_ps_fairplay"
)

// Columns holds all SQL columns for psfairplay fields.
var Columns = []string{
	FieldID,
	FieldFoulsCommitted,
	FieldYellow,
	FieldYellowRed,
	FieldRed,
	FieldPenaltyConceded,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ps_fairplays"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"player_stats_ps_fairplay",
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
	// DefaultFoulsCommitted holds the default value on creation for the "FoulsCommitted" field.
	DefaultFoulsCommitted int
	// DefaultYellow holds the default value on creation for the "Yellow" field.
	DefaultYellow int
	// DefaultYellowRed holds the default value on creation for the "YellowRed" field.
	DefaultYellowRed int
	// DefaultRed holds the default value on creation for the "Red" field.
	DefaultRed int
	// DefaultPenaltyConceded holds the default value on creation for the "PenaltyConceded" field.
	DefaultPenaltyConceded int
	// DefaultLastUpdated holds the default value on creation for the "lastUpdated" field.
	DefaultLastUpdated func() time.Time
	// UpdateDefaultLastUpdated holds the default value on update for the "lastUpdated" field.
	UpdateDefaultLastUpdated func() time.Time
)

// OrderOption defines the ordering options for the PSFairplay queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFoulsCommitted orders the results by the FoulsCommitted field.
func ByFoulsCommitted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFoulsCommitted, opts...).ToFunc()
}

// ByYellow orders the results by the Yellow field.
func ByYellow(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYellow, opts...).ToFunc()
}

// ByYellowRed orders the results by the YellowRed field.
func ByYellowRed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYellowRed, opts...).ToFunc()
}

// ByRed orders the results by the Red field.
func ByRed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRed, opts...).ToFunc()
}

// ByPenaltyConceded orders the results by the PenaltyConceded field.
func ByPenaltyConceded(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPenaltyConceded, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
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
		sqlgraph.Edge(sqlgraph.O2O, true, PlayerStatsTable, PlayerStatsColumn),
	)
}
