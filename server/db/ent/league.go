// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/league"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// League is the model entity for the League schema.
type League struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FootballApiId holds the value of the "footballApiId" field.
	FootballApiId int `json:"footballApiId,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type league.Type `json:"type,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
	// LastUpdated holds the value of the "lastUpdated" field.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LeagueQuery when eager-loading is set.
	Edges           LeagueEdges `json:"edges"`
	country_leagues *int
	selectValues    sql.SelectValues
}

// LeagueEdges holds the relations/edges for other nodes in the graph.
type LeagueEdges struct {
	// Country holds the value of the country edge.
	Country *Country `json:"country,omitempty"`
	// Season holds the value of the season edge.
	Season []*Season `json:"season,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LeagueEdges) CountryOrErr() (*Country, error) {
	if e.loadedTypes[0] {
		if e.Country == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: country.Label}
		}
		return e.Country, nil
	}
	return nil, &NotLoadedError{edge: "country"}
}

// SeasonOrErr returns the Season value or an error if the edge
// was not loaded in eager-loading.
func (e LeagueEdges) SeasonOrErr() ([]*Season, error) {
	if e.loadedTypes[1] {
		return e.Season, nil
	}
	return nil, &NotLoadedError{edge: "season"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*League) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case league.FieldID, league.FieldFootballApiId:
			values[i] = new(sql.NullInt64)
		case league.FieldSlug, league.FieldName, league.FieldType, league.FieldLogo:
			values[i] = new(sql.NullString)
		case league.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case league.ForeignKeys[0]: // country_leagues
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the League fields.
func (l *League) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case league.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case league.FieldFootballApiId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field footballApiId", values[i])
			} else if value.Valid {
				l.FootballApiId = int(value.Int64)
			}
		case league.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				l.Slug = value.String
			}
		case league.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case league.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				l.Type = league.Type(value.String)
			}
		case league.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				l.Logo = value.String
			}
		case league.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastUpdated", values[i])
			} else if value.Valid {
				l.LastUpdated = value.Time
			}
		case league.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field country_leagues", value)
			} else if value.Valid {
				l.country_leagues = new(int)
				*l.country_leagues = int(value.Int64)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the League.
// This includes values selected through modifiers, order, etc.
func (l *League) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryCountry queries the "country" edge of the League entity.
func (l *League) QueryCountry() *CountryQuery {
	return NewLeagueClient(l.config).QueryCountry(l)
}

// QuerySeason queries the "season" edge of the League entity.
func (l *League) QuerySeason() *SeasonQuery {
	return NewLeagueClient(l.config).QuerySeason(l)
}

// Update returns a builder for updating this League.
// Note that you need to call League.Unwrap() before calling this method if this League
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *League) Update() *LeagueUpdateOne {
	return NewLeagueClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the League entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *League) Unwrap() *League {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: League is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *League) String() string {
	var builder strings.Builder
	builder.WriteString("League(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("footballApiId=")
	builder.WriteString(fmt.Sprintf("%v", l.FootballApiId))
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(l.Slug)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", l.Type))
	builder.WriteString(", ")
	builder.WriteString("logo=")
	builder.WriteString(l.Logo)
	builder.WriteString(", ")
	builder.WriteString("lastUpdated=")
	builder.WriteString(l.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Leagues is a parsable slice of League.
type Leagues []*League
