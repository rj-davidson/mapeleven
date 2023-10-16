// Code generated by ent, DO NOT EDIT.

package league

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the league type in the database.
	Label = "league"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFootballApiId holds the string denoting the footballapiid field in the database.
	FieldFootballApiId = "football_api_id"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeCountry holds the string denoting the country edge name in mutations.
	EdgeCountry = "country"
	// EdgeSeason holds the string denoting the season edge name in mutations.
	EdgeSeason = "season"
	// Table holds the table name of the league in the database.
	Table = "leagues"
	// CountryTable is the table that holds the country relation/edge.
	CountryTable = "leagues"
	// CountryInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	CountryInverseTable = "countries"
	// CountryColumn is the table column denoting the country relation/edge.
	CountryColumn = "country_leagues"
	// SeasonTable is the table that holds the season relation/edge.
	SeasonTable = "seasons"
	// SeasonInverseTable is the table name for the Season entity.
	// It exists in this package in order to avoid circular dependency with the "season" package.
	SeasonInverseTable = "seasons"
	// SeasonColumn is the table column denoting the season relation/edge.
	SeasonColumn = "league_season"
)

// Columns holds all SQL columns for league fields.
var Columns = []string{
	FieldID,
	FieldFootballApiId,
	FieldSlug,
	FieldName,
	FieldType,
	FieldLogo,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "leagues"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"country_leagues",
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeLeague     Type = "League"
	TypeCup        Type = "Cup"
	TypeTournament Type = "Tournament"
	TypeFriendly   Type = "Friendly"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeLeague, TypeCup, TypeTournament, TypeFriendly:
		return nil
	default:
		return fmt.Errorf("league: invalid enum value for type field: %q", _type)
	}
}

// Order defines the ordering method for the League queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFootballApiId orders the results by the footballApiId field.
func ByFootballApiId(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldFootballApiId, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByLogo orders the results by the logo field.
func ByLogo(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLogo, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByCountryField orders the results by country field.
func ByCountryField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCountryStep(), sql.OrderByField(field, opts...))
	}
}

// BySeasonCount orders the results by season count.
func BySeasonCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSeasonStep(), opts...)
	}
}

// BySeason orders the results by season terms.
func BySeason(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSeasonStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCountryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CountryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CountryTable, CountryColumn),
	)
}
func newSeasonStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SeasonInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SeasonTable, SeasonColumn),
	)
}
