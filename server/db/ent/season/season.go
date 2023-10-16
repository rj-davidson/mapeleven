// Code generated by ent, DO NOT EDIT.

package season

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the season type in the database.
	Label = "season"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldCurrent holds the string denoting the current field in the database.
	FieldCurrent = "current"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeLeague holds the string denoting the league edge name in mutations.
	EdgeLeague = "league"
	// EdgeFixtures holds the string denoting the fixtures edge name in mutations.
	EdgeFixtures = "fixtures"
	// EdgeStandings holds the string denoting the standings edge name in mutations.
	EdgeStandings = "standings"
	// EdgeTeams holds the string denoting the teams edge name in mutations.
	EdgeTeams = "teams"
	// Table holds the table name of the season in the database.
	Table = "seasons"
	// LeagueTable is the table that holds the league relation/edge.
	LeagueTable = "seasons"
	// LeagueInverseTable is the table name for the League entity.
	// It exists in this package in order to avoid circular dependency with the "league" package.
	LeagueInverseTable = "leagues"
	// LeagueColumn is the table column denoting the league relation/edge.
	LeagueColumn = "league_season"
	// FixturesTable is the table that holds the fixtures relation/edge.
	FixturesTable = "fixtures"
	// FixturesInverseTable is the table name for the Fixture entity.
	// It exists in this package in order to avoid circular dependency with the "fixture" package.
	FixturesInverseTable = "fixtures"
	// FixturesColumn is the table column denoting the fixtures relation/edge.
	FixturesColumn = "season_fixtures"
	// StandingsTable is the table that holds the standings relation/edge.
	StandingsTable = "standings"
	// StandingsInverseTable is the table name for the Standings entity.
	// It exists in this package in order to avoid circular dependency with the "standings" package.
	StandingsInverseTable = "standings"
	// StandingsColumn is the table column denoting the standings relation/edge.
	StandingsColumn = "season_standings"
	// TeamsTable is the table that holds the teams relation/edge.
	TeamsTable = "teams"
	// TeamsInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamsInverseTable = "teams"
	// TeamsColumn is the table column denoting the teams relation/edge.
	TeamsColumn = "season_teams"
)

// Columns holds all SQL columns for season fields.
var Columns = []string{
	FieldID,
	FieldSlug,
	FieldYear,
	FieldStartDate,
	FieldEndDate,
	FieldCurrent,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "seasons"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"league_season",
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
	// DefaultLastUpdated holds the default value on creation for the "lastUpdated" field.
	DefaultLastUpdated func() time.Time
	// UpdateDefaultLastUpdated holds the default value on update for the "lastUpdated" field.
	UpdateDefaultLastUpdated func() time.Time
)

// Order defines the ordering method for the Season queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByCurrent orders the results by the current field.
func ByCurrent(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCurrent, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByLeagueField orders the results by league field.
func ByLeagueField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLeagueStep(), sql.OrderByField(field, opts...))
	}
}

// ByFixturesCount orders the results by fixtures count.
func ByFixturesCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFixturesStep(), opts...)
	}
}

// ByFixtures orders the results by fixtures terms.
func ByFixtures(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFixturesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStandingsCount orders the results by standings count.
func ByStandingsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStandingsStep(), opts...)
	}
}

// ByStandings orders the results by standings terms.
func ByStandings(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStandingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTeamsCount orders the results by teams count.
func ByTeamsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamsStep(), opts...)
	}
}

// ByTeams orders the results by teams terms.
func ByTeams(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLeagueStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LeagueInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LeagueTable, LeagueColumn),
	)
}
func newFixturesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FixturesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FixturesTable, FixturesColumn),
	)
}
func newStandingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StandingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StandingsTable, StandingsColumn),
	)
}
func newTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TeamsTable, TeamsColumn),
	)
}
