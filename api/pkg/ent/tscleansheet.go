// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tscleansheet"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TSCleanSheet is the model entity for the TSCleanSheet schema.
type TSCleanSheet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Home holds the value of the "home" field.
	Home int `json:"home,omitempty"`
	// Away holds the value of the "away" field.
	Away int `json:"away,omitempty"`
	// Total holds the value of the "total" field.
	Total int `json:"total,omitempty"`
	// LastUpdated holds the value of the "lastUpdated" field.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TSCleanSheetQuery when eager-loading is set.
	Edges                  TSCleanSheetEdges `json:"edges"`
	team_clean_sheet_stats *int
	selectValues           sql.SelectValues
}

// TSCleanSheetEdges holds the relations/edges for other nodes in the graph.
type TSCleanSheetEdges struct {
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TSCleanSheetEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[0] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TSCleanSheet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tscleansheet.FieldID, tscleansheet.FieldHome, tscleansheet.FieldAway, tscleansheet.FieldTotal:
			values[i] = new(sql.NullInt64)
		case tscleansheet.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case tscleansheet.ForeignKeys[0]: // team_clean_sheet_stats
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TSCleanSheet fields.
func (tcs *TSCleanSheet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tscleansheet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tcs.ID = int(value.Int64)
		case tscleansheet.FieldHome:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field home", values[i])
			} else if value.Valid {
				tcs.Home = int(value.Int64)
			}
		case tscleansheet.FieldAway:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field away", values[i])
			} else if value.Valid {
				tcs.Away = int(value.Int64)
			}
		case tscleansheet.FieldTotal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[i])
			} else if value.Valid {
				tcs.Total = int(value.Int64)
			}
		case tscleansheet.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastUpdated", values[i])
			} else if value.Valid {
				tcs.LastUpdated = value.Time
			}
		case tscleansheet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_clean_sheet_stats", value)
			} else if value.Valid {
				tcs.team_clean_sheet_stats = new(int)
				*tcs.team_clean_sheet_stats = int(value.Int64)
			}
		default:
			tcs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TSCleanSheet.
// This includes values selected through modifiers, order, etc.
func (tcs *TSCleanSheet) Value(name string) (ent.Value, error) {
	return tcs.selectValues.Get(name)
}

// QueryTeam queries the "team" edge of the TSCleanSheet entity.
func (tcs *TSCleanSheet) QueryTeam() *TeamQuery {
	return NewTSCleanSheetClient(tcs.config).QueryTeam(tcs)
}

// Update returns a builder for updating this TSCleanSheet.
// Note that you need to call TSCleanSheet.Unwrap() before calling this method if this TSCleanSheet
// was returned from a transaction, and the transaction was committed or rolled back.
func (tcs *TSCleanSheet) Update() *TSCleanSheetUpdateOne {
	return NewTSCleanSheetClient(tcs.config).UpdateOne(tcs)
}

// Unwrap unwraps the TSCleanSheet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tcs *TSCleanSheet) Unwrap() *TSCleanSheet {
	_tx, ok := tcs.config.driver.(*txDriver)
	if !ok {
		panic("ent: TSCleanSheet is not a transactional entity")
	}
	tcs.config.driver = _tx.drv
	return tcs
}

// String implements the fmt.Stringer.
func (tcs *TSCleanSheet) String() string {
	var builder strings.Builder
	builder.WriteString("TSCleanSheet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tcs.ID))
	builder.WriteString("home=")
	builder.WriteString(fmt.Sprintf("%v", tcs.Home))
	builder.WriteString(", ")
	builder.WriteString("away=")
	builder.WriteString(fmt.Sprintf("%v", tcs.Away))
	builder.WriteString(", ")
	builder.WriteString("total=")
	builder.WriteString(fmt.Sprintf("%v", tcs.Total))
	builder.WriteString(", ")
	builder.WriteString("lastUpdated=")
	builder.WriteString(tcs.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TSCleanSheets is a parsable slice of TSCleanSheet.
type TSCleanSheets []*TSCleanSheet
