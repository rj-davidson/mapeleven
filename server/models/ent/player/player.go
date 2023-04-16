// Code generated by ent, DO NOT EDIT.

package player

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the player type in the database.
	Label = "player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldInjured holds the string denoting the injured field in the database.
	FieldInjured = "injured"
	// FieldPhoto holds the string denoting the photo field in the database.
	FieldPhoto = "photo"
	// EdgeBirth holds the string denoting the birth edge name in mutations.
	EdgeBirth = "birth"
	// EdgeTeams holds the string denoting the teams edge name in mutations.
	EdgeTeams = "teams"
	// Table holds the table name of the player in the database.
	Table = "players"
	// BirthTable is the table that holds the birth relation/edge.
	BirthTable = "births"
	// BirthInverseTable is the table name for the Birth entity.
	// It exists in this package in order to avoid circular dependency with the "birth" package.
	BirthInverseTable = "births"
	// BirthColumn is the table column denoting the birth relation/edge.
	BirthColumn = "player_birth"
	// TeamsTable is the table that holds the teams relation/edge. The primary key declared below.
	TeamsTable = "team_players"
	// TeamsInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamsInverseTable = "teams"
)

// Columns holds all SQL columns for player fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFirstname,
	FieldLastname,
	FieldAge,
	FieldHeight,
	FieldWeight,
	FieldInjured,
	FieldPhoto,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "players"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"country_players",
}

var (
	// TeamsPrimaryKey and TeamsColumn2 are the table columns denoting the
	// primary key for the teams relation (M2M).
	TeamsPrimaryKey = []string{"team_id", "player_id"}
)

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

// Order defines the ordering method for the Player queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFirstname orders the results by the firstname field.
func ByFirstname(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldFirstname, opts...).ToFunc()
}

// ByLastname orders the results by the lastname field.
func ByLastname(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLastname, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByInjured orders the results by the injured field.
func ByInjured(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldInjured, opts...).ToFunc()
}

// ByPhoto orders the results by the photo field.
func ByPhoto(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPhoto, opts...).ToFunc()
}

// ByBirthField orders the results by birth field.
func ByBirthField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBirthStep(), sql.OrderByField(field, opts...))
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
func newBirthStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BirthInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, BirthTable, BirthColumn),
	)
}
func newTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TeamsTable, TeamsPrimaryKey...),
	)
}
