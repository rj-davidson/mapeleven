// Code generated by ent, DO NOT EDIT.

package psgoals

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the psgoals type in the database.
	Label = "ps_goals"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTotalGoals holds the string denoting the totalgoals field in the database.
	FieldTotalGoals = "total_goals"
	// FieldConcededGoals holds the string denoting the concededgoals field in the database.
	FieldConcededGoals = "conceded_goals"
	// FieldAssistGoals holds the string denoting the assistgoals field in the database.
	FieldAssistGoals = "assist_goals"
	// FieldSaveGoals holds the string denoting the savegoals field in the database.
	FieldSaveGoals = "save_goals"
	// FieldShotsTotal holds the string denoting the shotstotal field in the database.
	FieldShotsTotal = "shots_total"
	// FieldShotsOn holds the string denoting the shotson field in the database.
	FieldShotsOn = "shots_on"
	// EdgePlayer holds the string denoting the player edge name in mutations.
	EdgePlayer = "player"
	// Table holds the table name of the psgoals in the database.
	Table = "ps_goals"
	// PlayerTable is the table that holds the player relation/edge.
	PlayerTable = "ps_goals"
	// PlayerInverseTable is the table name for the Player entity.
	// It exists in this package in order to avoid circular dependency with the "player" package.
	PlayerInverseTable = "players"
	// PlayerColumn is the table column denoting the player relation/edge.
	PlayerColumn = "player_psgoals"
)

// Columns holds all SQL columns for psgoals fields.
var Columns = []string{
	FieldID,
	FieldTotalGoals,
	FieldConcededGoals,
	FieldAssistGoals,
	FieldSaveGoals,
	FieldShotsTotal,
	FieldShotsOn,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ps_goals"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"player_psgoals",
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
	// DefaultSaveGoals holds the default value on creation for the "saveGoals" field.
	DefaultSaveGoals int
)

// OrderOption defines the ordering options for the PSGoals queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTotalGoals orders the results by the totalGoals field.
func ByTotalGoals(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalGoals, opts...).ToFunc()
}

// ByConcededGoals orders the results by the concededGoals field.
func ByConcededGoals(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConcededGoals, opts...).ToFunc()
}

// ByAssistGoals orders the results by the assistGoals field.
func ByAssistGoals(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAssistGoals, opts...).ToFunc()
}

// BySaveGoals orders the results by the saveGoals field.
func BySaveGoals(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSaveGoals, opts...).ToFunc()
}

// ByShotsTotal orders the results by the shotsTotal field.
func ByShotsTotal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShotsTotal, opts...).ToFunc()
}

// ByShotsOn orders the results by the shotsOn field.
func ByShotsOn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShotsOn, opts...).ToFunc()
}

// ByPlayerField orders the results by player field.
func ByPlayerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayerStep(), sql.OrderByField(field, opts...))
	}
}
func newPlayerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlayerTable, PlayerColumn),
	)
}
