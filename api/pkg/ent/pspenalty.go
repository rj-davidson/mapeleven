// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PSPenalty is the model entity for the PSPenalty schema.
type PSPenalty struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Won holds the value of the "Won" field.
	Won int `json:"Won,omitempty"`
	// Scored holds the value of the "Scored" field.
	Scored int `json:"Scored,omitempty"`
	// Missed holds the value of the "Missed" field.
	Missed int `json:"Missed,omitempty"`
	// Saved holds the value of the "Saved" field.
	Saved int `json:"Saved,omitempty"`
	// LastUpdated holds the value of the "lastUpdated" field.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PSPenaltyQuery when eager-loading is set.
	Edges                   PSPenaltyEdges `json:"edges"`
	player_stats_ps_penalty *int
	selectValues            sql.SelectValues
}

// PSPenaltyEdges holds the relations/edges for other nodes in the graph.
type PSPenaltyEdges struct {
	// PlayerStats holds the value of the playerStats edge.
	PlayerStats *PlayerStats `json:"playerStats,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlayerStatsOrErr returns the PlayerStats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PSPenaltyEdges) PlayerStatsOrErr() (*PlayerStats, error) {
	if e.loadedTypes[0] {
		if e.PlayerStats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: playerstats.Label}
		}
		return e.PlayerStats, nil
	}
	return nil, &NotLoadedError{edge: "playerStats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PSPenalty) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pspenalty.FieldID, pspenalty.FieldWon, pspenalty.FieldScored, pspenalty.FieldMissed, pspenalty.FieldSaved:
			values[i] = new(sql.NullInt64)
		case pspenalty.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case pspenalty.ForeignKeys[0]: // player_stats_ps_penalty
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PSPenalty fields.
func (pp *PSPenalty) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pspenalty.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pp.ID = int(value.Int64)
		case pspenalty.FieldWon:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Won", values[i])
			} else if value.Valid {
				pp.Won = int(value.Int64)
			}
		case pspenalty.FieldScored:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Scored", values[i])
			} else if value.Valid {
				pp.Scored = int(value.Int64)
			}
		case pspenalty.FieldMissed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Missed", values[i])
			} else if value.Valid {
				pp.Missed = int(value.Int64)
			}
		case pspenalty.FieldSaved:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Saved", values[i])
			} else if value.Valid {
				pp.Saved = int(value.Int64)
			}
		case pspenalty.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastUpdated", values[i])
			} else if value.Valid {
				pp.LastUpdated = value.Time
			}
		case pspenalty.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field player_stats_ps_penalty", value)
			} else if value.Valid {
				pp.player_stats_ps_penalty = new(int)
				*pp.player_stats_ps_penalty = int(value.Int64)
			}
		default:
			pp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PSPenalty.
// This includes values selected through modifiers, order, etc.
func (pp *PSPenalty) Value(name string) (ent.Value, error) {
	return pp.selectValues.Get(name)
}

// QueryPlayerStats queries the "playerStats" edge of the PSPenalty entity.
func (pp *PSPenalty) QueryPlayerStats() *PlayerStatsQuery {
	return NewPSPenaltyClient(pp.config).QueryPlayerStats(pp)
}

// Update returns a builder for updating this PSPenalty.
// Note that you need to call PSPenalty.Unwrap() before calling this method if this PSPenalty
// was returned from a transaction, and the transaction was committed or rolled back.
func (pp *PSPenalty) Update() *PSPenaltyUpdateOne {
	return NewPSPenaltyClient(pp.config).UpdateOne(pp)
}

// Unwrap unwraps the PSPenalty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pp *PSPenalty) Unwrap() *PSPenalty {
	_tx, ok := pp.config.driver.(*txDriver)
	if !ok {
		panic("ent: PSPenalty is not a transactional entity")
	}
	pp.config.driver = _tx.drv
	return pp
}

// String implements the fmt.Stringer.
func (pp *PSPenalty) String() string {
	var builder strings.Builder
	builder.WriteString("PSPenalty(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pp.ID))
	builder.WriteString("Won=")
	builder.WriteString(fmt.Sprintf("%v", pp.Won))
	builder.WriteString(", ")
	builder.WriteString("Scored=")
	builder.WriteString(fmt.Sprintf("%v", pp.Scored))
	builder.WriteString(", ")
	builder.WriteString("Missed=")
	builder.WriteString(fmt.Sprintf("%v", pp.Missed))
	builder.WriteString(", ")
	builder.WriteString("Saved=")
	builder.WriteString(fmt.Sprintf("%v", pp.Saved))
	builder.WriteString(", ")
	builder.WriteString("lastUpdated=")
	builder.WriteString(pp.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PSPenalties is a parsable slice of PSPenalty.
type PSPenalties []*PSPenalty
