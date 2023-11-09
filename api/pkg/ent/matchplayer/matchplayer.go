// Code generated by ent, DO NOT EDIT.

package matchplayer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the matchplayer type in the database.
	Label = "match_player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldX holds the string denoting the x field in the database.
	FieldX = "x"
	// FieldY holds the string denoting the y field in the database.
	FieldY = "y"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgePlayer holds the string denoting the player edge name in mutations.
	EdgePlayer = "player"
	// EdgeLineup holds the string denoting the lineup edge name in mutations.
	EdgeLineup = "lineup"
	// EdgePlayerStats holds the string denoting the playerstats edge name in mutations.
	EdgePlayerStats = "playerStats"
	// Table holds the table name of the matchplayer in the database.
	Table = "match_players"
	// PlayerTable is the table that holds the player relation/edge.
	PlayerTable = "match_players"
	// PlayerInverseTable is the table name for the Player entity.
	// It exists in this package in order to avoid circular dependency with the "player" package.
	PlayerInverseTable = "players"
	// PlayerColumn is the table column denoting the player relation/edge.
	PlayerColumn = "player_match_player"
	// LineupTable is the table that holds the lineup relation/edge.
	LineupTable = "match_players"
	// LineupInverseTable is the table name for the FixtureLineups entity.
	// It exists in this package in order to avoid circular dependency with the "fixturelineups" package.
	LineupInverseTable = "fixture_lineups"
	// LineupColumn is the table column denoting the lineup relation/edge.
	LineupColumn = "fixture_lineups_lineup_player"
	// PlayerStatsTable is the table that holds the playerStats relation/edge.
	PlayerStatsTable = "match_players"
	// PlayerStatsInverseTable is the table name for the PlayerStats entity.
	// It exists in this package in order to avoid circular dependency with the "playerstats" package.
	PlayerStatsInverseTable = "player_stats"
	// PlayerStatsColumn is the table column denoting the playerStats relation/edge.
	PlayerStatsColumn = "player_stats_match_player"
)

// Columns holds all SQL columns for matchplayer fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldPosition,
	FieldX,
	FieldY,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "match_players"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"fixture_lineups_lineup_player",
	"player_match_player",
	"player_stats_match_player",
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

// OrderOption defines the ordering options for the MatchPlayer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByPosition orders the results by the position field.
func ByPosition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPosition, opts...).ToFunc()
}

// ByX orders the results by the x field.
func ByX(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldX, opts...).ToFunc()
}

// ByY orders the results by the y field.
func ByY(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldY, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByPlayerField orders the results by player field.
func ByPlayerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayerStep(), sql.OrderByField(field, opts...))
	}
}

// ByLineupField orders the results by lineup field.
func ByLineupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLineupStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlayerStatsField orders the results by playerStats field.
func ByPlayerStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayerStatsStep(), sql.OrderByField(field, opts...))
	}
}
func newPlayerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlayerTable, PlayerColumn),
	)
}
func newLineupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LineupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LineupTable, LineupColumn),
	)
}
func newPlayerStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayerStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlayerStatsTable, PlayerStatsColumn),
	)
}
