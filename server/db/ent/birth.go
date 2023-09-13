// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mapeleven/db/ent/birth"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Birth is the model entity for the Birth schema.
type Birth struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Place holds the value of the "place" field.
	Place string `json:"place,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BirthQuery when eager-loading is set.
	Edges        BirthEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BirthEdges holds the relations/edges for other nodes in the graph.
type BirthEdges struct {
	// Player holds the value of the player edge.
	Player []*Player `json:"player,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlayerOrErr returns the Player value or an error if the edge
// was not loaded in eager-loading.
func (e BirthEdges) PlayerOrErr() ([]*Player, error) {
	if e.loadedTypes[0] {
		return e.Player, nil
	}
	return nil, &NotLoadedError{edge: "player"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Birth) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case birth.FieldID:
			values[i] = new(sql.NullInt64)
		case birth.FieldPlace, birth.FieldCountry:
			values[i] = new(sql.NullString)
		case birth.FieldDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Birth fields.
func (b *Birth) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case birth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case birth.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				b.Date = value.Time
			}
		case birth.FieldPlace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place", values[i])
			} else if value.Valid {
				b.Place = value.String
			}
		case birth.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				b.Country = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Birth.
// This includes values selected through modifiers, order, etc.
func (b *Birth) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryPlayer queries the "player" edge of the Birth entity.
func (b *Birth) QueryPlayer() *PlayerQuery {
	return NewBirthClient(b.config).QueryPlayer(b)
}

// Update returns a builder for updating this Birth.
// Note that you need to call Birth.Unwrap() before calling this method if this Birth
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Birth) Update() *BirthUpdateOne {
	return NewBirthClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Birth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Birth) Unwrap() *Birth {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Birth is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Birth) String() string {
	var builder strings.Builder
	builder.WriteString("Birth(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("date=")
	builder.WriteString(b.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("place=")
	builder.WriteString(b.Place)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(b.Country)
	builder.WriteByte(')')
	return builder.String()
}

// Births is a parsable slice of Birth.
type Births []*Birth
