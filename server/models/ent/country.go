// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mapeleven/models/ent/country"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Country is the model entity for the Country schema.
type Country struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Flag holds the value of the "flag" field.
	Flag string `json:"flag,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CountryQuery when eager-loading is set.
	Edges        CountryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CountryEdges holds the relations/edges for other nodes in the graph.
type CountryEdges struct {
	// Players holds the value of the players edge.
	Players []*Player `json:"players,omitempty"`
	// Leagues holds the value of the leagues edge.
	Leagues []*League `json:"leagues,omitempty"`
	// Teams holds the value of the teams edge.
	Teams []*Team `json:"teams,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[0] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// LeaguesOrErr returns the Leagues value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) LeaguesOrErr() ([]*League, error) {
	if e.loadedTypes[1] {
		return e.Leagues, nil
	}
	return nil, &NotLoadedError{edge: "leagues"}
}

// TeamsOrErr returns the Teams value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) TeamsOrErr() ([]*Team, error) {
	if e.loadedTypes[2] {
		return e.Teams, nil
	}
	return nil, &NotLoadedError{edge: "teams"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			values[i] = new(sql.NullInt64)
		case country.FieldName, country.FieldCode, country.FieldFlag:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country fields.
func (c *Country) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case country.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case country.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				c.Code = value.String
			}
		case country.FieldFlag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flag", values[i])
			} else if value.Valid {
				c.Flag = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Country.
// This includes values selected through modifiers, order, etc.
func (c *Country) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryPlayers queries the "players" edge of the Country entity.
func (c *Country) QueryPlayers() *PlayerQuery {
	return NewCountryClient(c.config).QueryPlayers(c)
}

// QueryLeagues queries the "leagues" edge of the Country entity.
func (c *Country) QueryLeagues() *LeagueQuery {
	return NewCountryClient(c.config).QueryLeagues(c)
}

// QueryTeams queries the "teams" edge of the Country entity.
func (c *Country) QueryTeams() *TeamQuery {
	return NewCountryClient(c.config).QueryTeams(c)
}

// Update returns a builder for updating this Country.
// Note that you need to call Country.Unwrap() before calling this method if this Country
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Country) Update() *CountryUpdateOne {
	return NewCountryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Country entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Country) Unwrap() *Country {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Country is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Country) String() string {
	var builder strings.Builder
	builder.WriteString("Country(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(c.Code)
	builder.WriteString(", ")
	builder.WriteString("flag=")
	builder.WriteString(c.Flag)
	builder.WriteByte(')')
	return builder.String()
}

// Countries is a parsable slice of Country.
type Countries []*Country
