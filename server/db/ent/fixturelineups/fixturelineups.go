// Code generated by ent, DO NOT EDIT.

package fixturelineups

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the fixturelineups type in the database.
	Label = "fixture_lineups"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFormation holds the string denoting the formation field in the database.
	FieldFormation = "formation"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// EdgeFixture holds the string denoting the fixture edge name in mutations.
	EdgeFixture = "fixture"
	// EdgeLineupPlayer holds the string denoting the lineupplayer edge name in mutations.
	EdgeLineupPlayer = "lineupPlayer"
	// Table holds the table name of the fixturelineups in the database.
	Table = "fixture_lineups"
	// TeamTable is the table that holds the team relation/edge.
	TeamTable = "fixture_lineups"
	// TeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamInverseTable = "teams"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "team_fixture_lineups"
	// FixtureTable is the table that holds the fixture relation/edge.
	FixtureTable = "fixture_lineups"
	// FixtureInverseTable is the table name for the Fixture entity.
	// It exists in this package in order to avoid circular dependency with the "fixture" package.
	FixtureInverseTable = "fixtures"
	// FixtureColumn is the table column denoting the fixture relation/edge.
	FixtureColumn = "fixture_lineups"
	// LineupPlayerTable is the table that holds the lineupPlayer relation/edge.
	LineupPlayerTable = "match_players"
	// LineupPlayerInverseTable is the table name for the MatchPlayer entity.
	// It exists in this package in order to avoid circular dependency with the "matchplayer" package.
	LineupPlayerInverseTable = "match_players"
	// LineupPlayerColumn is the table column denoting the lineupPlayer relation/edge.
	LineupPlayerColumn = "fixture_lineups_lineup_player"
)

// Columns holds all SQL columns for fixturelineups fields.
var Columns = []string{
	FieldID,
	FieldFormation,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "fixture_lineups"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"fixture_lineups",
	"team_fixture_lineups",
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

// Order defines the ordering method for the FixtureLineups queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFormation orders the results by the formation field.
func ByFormation(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldFormation, opts...).ToFunc()
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

// ByFixtureField orders the results by fixture field.
func ByFixtureField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFixtureStep(), sql.OrderByField(field, opts...))
	}
}

// ByLineupPlayerCount orders the results by lineupPlayer count.
func ByLineupPlayerCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLineupPlayerStep(), opts...)
	}
}

// ByLineupPlayer orders the results by lineupPlayer terms.
func ByLineupPlayer(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLineupPlayerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TeamTable, TeamColumn),
	)
}
func newFixtureStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FixtureInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FixtureTable, FixtureColumn),
	)
}
func newLineupPlayerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LineupPlayerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LineupPlayerTable, LineupPlayerColumn),
	)
}
