// Code generated by ent, DO NOT EDIT.

package fixture

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the fixture type in the database.
	Label = "fixture"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldApiFootballId holds the string denoting the apifootballid field in the database.
	FieldApiFootballId = "api_football_id"
	// FieldReferee holds the string denoting the referee field in the database.
	FieldReferee = "referee"
	// FieldTimezone holds the string denoting the timezone field in the database.
	FieldTimezone = "timezone"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldElapsed holds the string denoting the elapsed field in the database.
	FieldElapsed = "elapsed"
	// FieldRound holds the string denoting the round field in the database.
	FieldRound = "round"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldHomeTeamScore holds the string denoting the hometeamscore field in the database.
	FieldHomeTeamScore = "home_team_score"
	// FieldAwayTeamScore holds the string denoting the awayteamscore field in the database.
	FieldAwayTeamScore = "away_team_score"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeHomeTeam holds the string denoting the hometeam edge name in mutations.
	EdgeHomeTeam = "homeTeam"
	// EdgeAwayTeam holds the string denoting the awayteam edge name in mutations.
	EdgeAwayTeam = "awayTeam"
	// EdgeSeason holds the string denoting the season edge name in mutations.
	EdgeSeason = "season"
	// EdgeLineups holds the string denoting the lineups edge name in mutations.
	EdgeLineups = "lineups"
	// EdgeFixtureEvents holds the string denoting the fixtureevents edge name in mutations.
	EdgeFixtureEvents = "fixtureEvents"
	// Table holds the table name of the fixture in the database.
	Table = "fixtures"
	// HomeTeamTable is the table that holds the homeTeam relation/edge.
	HomeTeamTable = "fixtures"
	// HomeTeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	HomeTeamInverseTable = "teams"
	// HomeTeamColumn is the table column denoting the homeTeam relation/edge.
	HomeTeamColumn = "team_home_fixtures"
	// AwayTeamTable is the table that holds the awayTeam relation/edge.
	AwayTeamTable = "fixtures"
	// AwayTeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	AwayTeamInverseTable = "teams"
	// AwayTeamColumn is the table column denoting the awayTeam relation/edge.
	AwayTeamColumn = "team_away_fixtures"
	// SeasonTable is the table that holds the season relation/edge.
	SeasonTable = "fixtures"
	// SeasonInverseTable is the table name for the Season entity.
	// It exists in this package in order to avoid circular dependency with the "season" package.
	SeasonInverseTable = "seasons"
	// SeasonColumn is the table column denoting the season relation/edge.
	SeasonColumn = "season_fixtures"
	// LineupsTable is the table that holds the lineups relation/edge.
	LineupsTable = "fixture_lineups"
	// LineupsInverseTable is the table name for the FixtureLineups entity.
	// It exists in this package in order to avoid circular dependency with the "fixturelineups" package.
	LineupsInverseTable = "fixture_lineups"
	// LineupsColumn is the table column denoting the lineups relation/edge.
	LineupsColumn = "fixture_lineups"
	// FixtureEventsTable is the table that holds the fixtureEvents relation/edge.
	FixtureEventsTable = "fixture_events"
	// FixtureEventsInverseTable is the table name for the FixtureEvents entity.
	// It exists in this package in order to avoid circular dependency with the "fixtureevents" package.
	FixtureEventsInverseTable = "fixture_events"
	// FixtureEventsColumn is the table column denoting the fixtureEvents relation/edge.
	FixtureEventsColumn = "fixture_fixture_events"
)

// Columns holds all SQL columns for fixture fields.
var Columns = []string{
	FieldID,
	FieldSlug,
	FieldApiFootballId,
	FieldReferee,
	FieldTimezone,
	FieldDate,
	FieldElapsed,
	FieldRound,
	FieldStatus,
	FieldHomeTeamScore,
	FieldAwayTeamScore,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "fixtures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"season_fixtures",
	"team_home_fixtures",
	"team_away_fixtures",
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

// Order defines the ordering method for the Fixture queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByApiFootballId orders the results by the apiFootballId field.
func ByApiFootballId(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldApiFootballId, opts...).ToFunc()
}

// ByReferee orders the results by the referee field.
func ByReferee(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldReferee, opts...).ToFunc()
}

// ByTimezone orders the results by the timezone field.
func ByTimezone(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTimezone, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByElapsed orders the results by the elapsed field.
func ByElapsed(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldElapsed, opts...).ToFunc()
}

// ByRound orders the results by the round field.
func ByRound(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldRound, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByHomeTeamScore orders the results by the homeTeamScore field.
func ByHomeTeamScore(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldHomeTeamScore, opts...).ToFunc()
}

// ByAwayTeamScore orders the results by the awayTeamScore field.
func ByAwayTeamScore(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAwayTeamScore, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByHomeTeamField orders the results by homeTeam field.
func ByHomeTeamField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHomeTeamStep(), sql.OrderByField(field, opts...))
	}
}

// ByAwayTeamField orders the results by awayTeam field.
func ByAwayTeamField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAwayTeamStep(), sql.OrderByField(field, opts...))
	}
}

// BySeasonField orders the results by season field.
func BySeasonField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSeasonStep(), sql.OrderByField(field, opts...))
	}
}

// ByLineupsCount orders the results by lineups count.
func ByLineupsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLineupsStep(), opts...)
	}
}

// ByLineups orders the results by lineups terms.
func ByLineups(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLineupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFixtureEventsCount orders the results by fixtureEvents count.
func ByFixtureEventsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFixtureEventsStep(), opts...)
	}
}

// ByFixtureEvents orders the results by fixtureEvents terms.
func ByFixtureEvents(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFixtureEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHomeTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HomeTeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, HomeTeamTable, HomeTeamColumn),
	)
}
func newAwayTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AwayTeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AwayTeamTable, AwayTeamColumn),
	)
}
func newSeasonStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SeasonInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SeasonTable, SeasonColumn),
	)
}
func newLineupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LineupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LineupsTable, LineupsColumn),
	)
}
func newFixtureEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FixtureEventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FixtureEventsTable, FixtureEventsColumn),
	)
}
