// Code generated by ent, DO NOT EDIT.

package tsfixtures

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tsfixtures type in the database.
	Label = "ts_fixtures"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPlayedHome holds the string denoting the playedhome field in the database.
	FieldPlayedHome = "played_home"
	// FieldPlayedAway holds the string denoting the playedaway field in the database.
	FieldPlayedAway = "played_away"
	// FieldPlayedTotal holds the string denoting the playedtotal field in the database.
	FieldPlayedTotal = "played_total"
	// FieldWinsHome holds the string denoting the winshome field in the database.
	FieldWinsHome = "wins_home"
	// FieldWinsAway holds the string denoting the winsaway field in the database.
	FieldWinsAway = "wins_away"
	// FieldWinsTotal holds the string denoting the winstotal field in the database.
	FieldWinsTotal = "wins_total"
	// FieldDrawsHome holds the string denoting the drawshome field in the database.
	FieldDrawsHome = "draws_home"
	// FieldDrawsAway holds the string denoting the drawsaway field in the database.
	FieldDrawsAway = "draws_away"
	// FieldDrawsTotal holds the string denoting the drawstotal field in the database.
	FieldDrawsTotal = "draws_total"
	// FieldLossesHome holds the string denoting the losseshome field in the database.
	FieldLossesHome = "losses_home"
	// FieldLossesAway holds the string denoting the lossesaway field in the database.
	FieldLossesAway = "losses_away"
	// FieldLossesTotal holds the string denoting the lossestotal field in the database.
	FieldLossesTotal = "losses_total"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// Table holds the table name of the tsfixtures in the database.
	Table = "ts_fixtures"
	// TeamTable is the table that holds the team relation/edge.
	TeamTable = "ts_fixtures"
	// TeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamInverseTable = "teams"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "team_fixtures_stats"
)

// Columns holds all SQL columns for tsfixtures fields.
var Columns = []string{
	FieldID,
	FieldPlayedHome,
	FieldPlayedAway,
	FieldPlayedTotal,
	FieldWinsHome,
	FieldWinsAway,
	FieldWinsTotal,
	FieldDrawsHome,
	FieldDrawsAway,
	FieldDrawsTotal,
	FieldLossesHome,
	FieldLossesAway,
	FieldLossesTotal,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ts_fixtures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"team_fixtures_stats",
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
	// DefaultPlayedHome holds the default value on creation for the "playedHome" field.
	DefaultPlayedHome int
	// DefaultPlayedAway holds the default value on creation for the "playedAway" field.
	DefaultPlayedAway int
	// DefaultPlayedTotal holds the default value on creation for the "playedTotal" field.
	DefaultPlayedTotal int
	// DefaultWinsHome holds the default value on creation for the "winsHome" field.
	DefaultWinsHome int
	// DefaultWinsAway holds the default value on creation for the "winsAway" field.
	DefaultWinsAway int
	// DefaultWinsTotal holds the default value on creation for the "winsTotal" field.
	DefaultWinsTotal int
	// DefaultDrawsHome holds the default value on creation for the "drawsHome" field.
	DefaultDrawsHome int
	// DefaultDrawsAway holds the default value on creation for the "drawsAway" field.
	DefaultDrawsAway int
	// DefaultDrawsTotal holds the default value on creation for the "drawsTotal" field.
	DefaultDrawsTotal int
	// DefaultLossesHome holds the default value on creation for the "lossesHome" field.
	DefaultLossesHome int
	// DefaultLossesAway holds the default value on creation for the "lossesAway" field.
	DefaultLossesAway int
	// DefaultLossesTotal holds the default value on creation for the "lossesTotal" field.
	DefaultLossesTotal int
	// DefaultLastUpdated holds the default value on creation for the "lastUpdated" field.
	DefaultLastUpdated func() time.Time
	// UpdateDefaultLastUpdated holds the default value on update for the "lastUpdated" field.
	UpdateDefaultLastUpdated func() time.Time
)

// Order defines the ordering method for the TSFixtures queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPlayedHome orders the results by the playedHome field.
func ByPlayedHome(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPlayedHome, opts...).ToFunc()
}

// ByPlayedAway orders the results by the playedAway field.
func ByPlayedAway(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPlayedAway, opts...).ToFunc()
}

// ByPlayedTotal orders the results by the playedTotal field.
func ByPlayedTotal(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPlayedTotal, opts...).ToFunc()
}

// ByWinsHome orders the results by the winsHome field.
func ByWinsHome(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldWinsHome, opts...).ToFunc()
}

// ByWinsAway orders the results by the winsAway field.
func ByWinsAway(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldWinsAway, opts...).ToFunc()
}

// ByWinsTotal orders the results by the winsTotal field.
func ByWinsTotal(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldWinsTotal, opts...).ToFunc()
}

// ByDrawsHome orders the results by the drawsHome field.
func ByDrawsHome(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDrawsHome, opts...).ToFunc()
}

// ByDrawsAway orders the results by the drawsAway field.
func ByDrawsAway(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDrawsAway, opts...).ToFunc()
}

// ByDrawsTotal orders the results by the drawsTotal field.
func ByDrawsTotal(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDrawsTotal, opts...).ToFunc()
}

// ByLossesHome orders the results by the lossesHome field.
func ByLossesHome(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLossesHome, opts...).ToFunc()
}

// ByLossesAway orders the results by the lossesAway field.
func ByLossesAway(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLossesAway, opts...).ToFunc()
}

// ByLossesTotal orders the results by the lossesTotal field.
func ByLossesTotal(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLossesTotal, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByTeamField orders the results by team field.
func ByTeamField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamStep(), sql.OrderByField(field, opts...))
	}
}
func newTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, TeamTable, TeamColumn),
	)
}
