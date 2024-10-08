// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tsgoals"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TSGoals is the model entity for the TSGoals schema.
type TSGoals struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GoalsForTotalHome holds the value of the "goalsForTotalHome" field.
	GoalsForTotalHome int `json:"goalsForTotalHome,omitempty"`
	// GoalsForTotalAway holds the value of the "goalsForTotalAway" field.
	GoalsForTotalAway int `json:"goalsForTotalAway,omitempty"`
	// GoalsForTotal holds the value of the "goalsForTotal" field.
	GoalsForTotal int `json:"goalsForTotal,omitempty"`
	// GoalsForAverageHome holds the value of the "goalsForAverageHome" field.
	GoalsForAverageHome string `json:"goalsForAverageHome,omitempty"`
	// GoalsForAverageAway holds the value of the "goalsForAverageAway" field.
	GoalsForAverageAway string `json:"goalsForAverageAway,omitempty"`
	// GoalsForAverageTotal holds the value of the "goalsForAverageTotal" field.
	GoalsForAverageTotal string `json:"goalsForAverageTotal,omitempty"`
	// GoalsForMinute0To15Total holds the value of the "goalsForMinute0To15Total" field.
	GoalsForMinute0To15Total int `json:"goalsForMinute0To15Total,omitempty"`
	// GoalsForMinute0To15Percentage holds the value of the "goalsForMinute0To15Percentage" field.
	GoalsForMinute0To15Percentage string `json:"goalsForMinute0To15Percentage,omitempty"`
	// GoalsForMinute16To30Total holds the value of the "goalsForMinute16To30Total" field.
	GoalsForMinute16To30Total int `json:"goalsForMinute16To30Total,omitempty"`
	// GoalsForMinute16To30Percentage holds the value of the "goalsForMinute16To30Percentage" field.
	GoalsForMinute16To30Percentage string `json:"goalsForMinute16To30Percentage,omitempty"`
	// GoalsForMinute31To45Total holds the value of the "goalsForMinute31To45Total" field.
	GoalsForMinute31To45Total int `json:"goalsForMinute31To45Total,omitempty"`
	// GoalsForMinute31To45Percentage holds the value of the "goalsForMinute31To45Percentage" field.
	GoalsForMinute31To45Percentage string `json:"goalsForMinute31To45Percentage,omitempty"`
	// GoalsForMinute46To60Total holds the value of the "goalsForMinute46To60Total" field.
	GoalsForMinute46To60Total int `json:"goalsForMinute46To60Total,omitempty"`
	// GoalsForMinute46To60Percentage holds the value of the "goalsForMinute46To60Percentage" field.
	GoalsForMinute46To60Percentage string `json:"goalsForMinute46To60Percentage,omitempty"`
	// GoalsForMinute61To75Total holds the value of the "goalsForMinute61To75Total" field.
	GoalsForMinute61To75Total int `json:"goalsForMinute61To75Total,omitempty"`
	// GoalsForMinute61To75Percentage holds the value of the "goalsForMinute61To75Percentage" field.
	GoalsForMinute61To75Percentage string `json:"goalsForMinute61To75Percentage,omitempty"`
	// GoalsForMinute76To90Total holds the value of the "goalsForMinute76To90Total" field.
	GoalsForMinute76To90Total int `json:"goalsForMinute76To90Total,omitempty"`
	// GoalsForMinute76To90Percentage holds the value of the "goalsForMinute76To90Percentage" field.
	GoalsForMinute76To90Percentage string `json:"goalsForMinute76To90Percentage,omitempty"`
	// GoalsForMinute91To105Total holds the value of the "goalsForMinute91To105Total" field.
	GoalsForMinute91To105Total int `json:"goalsForMinute91To105Total,omitempty"`
	// GoalsForMinute91To105Percentage holds the value of the "goalsForMinute91To105Percentage" field.
	GoalsForMinute91To105Percentage string `json:"goalsForMinute91To105Percentage,omitempty"`
	// GoalsForMinute106To120Total holds the value of the "goalsForMinute106To120Total" field.
	GoalsForMinute106To120Total int `json:"goalsForMinute106To120Total,omitempty"`
	// GoalsForMinute106To120Percentage holds the value of the "goalsForMinute106To120Percentage" field.
	GoalsForMinute106To120Percentage string `json:"goalsForMinute106To120Percentage,omitempty"`
	// GoalsAgainstTotalHome holds the value of the "goalsAgainstTotalHome" field.
	GoalsAgainstTotalHome int `json:"goalsAgainstTotalHome,omitempty"`
	// GoalsAgainstTotalAway holds the value of the "goalsAgainstTotalAway" field.
	GoalsAgainstTotalAway int `json:"goalsAgainstTotalAway,omitempty"`
	// GoalsAgainstTotal holds the value of the "goalsAgainstTotal" field.
	GoalsAgainstTotal int `json:"goalsAgainstTotal,omitempty"`
	// GoalsAgainstAverageHome holds the value of the "goalsAgainstAverageHome" field.
	GoalsAgainstAverageHome string `json:"goalsAgainstAverageHome,omitempty"`
	// GoalsAgainstAverageAway holds the value of the "goalsAgainstAverageAway" field.
	GoalsAgainstAverageAway string `json:"goalsAgainstAverageAway,omitempty"`
	// GoalsAgainstAverageTotal holds the value of the "goalsAgainstAverageTotal" field.
	GoalsAgainstAverageTotal string `json:"goalsAgainstAverageTotal,omitempty"`
	// GoalsAgainstMinute0To15Total holds the value of the "goalsAgainstMinute0To15Total" field.
	GoalsAgainstMinute0To15Total int `json:"goalsAgainstMinute0To15Total,omitempty"`
	// GoalsAgainstMinute0To15Percentage holds the value of the "goalsAgainstMinute0To15Percentage" field.
	GoalsAgainstMinute0To15Percentage string `json:"goalsAgainstMinute0To15Percentage,omitempty"`
	// GoalsAgainstMinute16To30Total holds the value of the "goalsAgainstMinute16To30Total" field.
	GoalsAgainstMinute16To30Total int `json:"goalsAgainstMinute16To30Total,omitempty"`
	// GoalsAgainstMinute16To30Percentage holds the value of the "goalsAgainstMinute16To30Percentage" field.
	GoalsAgainstMinute16To30Percentage string `json:"goalsAgainstMinute16To30Percentage,omitempty"`
	// GoalsAgainstMinute31To45Total holds the value of the "goalsAgainstMinute31To45Total" field.
	GoalsAgainstMinute31To45Total int `json:"goalsAgainstMinute31To45Total,omitempty"`
	// GoalsAgainstMinute31To45Percentage holds the value of the "goalsAgainstMinute31To45Percentage" field.
	GoalsAgainstMinute31To45Percentage string `json:"goalsAgainstMinute31To45Percentage,omitempty"`
	// GoalsAgainstMinute46To60Total holds the value of the "goalsAgainstMinute46To60Total" field.
	GoalsAgainstMinute46To60Total int `json:"goalsAgainstMinute46To60Total,omitempty"`
	// GoalsAgainstMinute46To60Percentage holds the value of the "goalsAgainstMinute46To60Percentage" field.
	GoalsAgainstMinute46To60Percentage string `json:"goalsAgainstMinute46To60Percentage,omitempty"`
	// GoalsAgainstMinute61To75Total holds the value of the "goalsAgainstMinute61To75Total" field.
	GoalsAgainstMinute61To75Total int `json:"goalsAgainstMinute61To75Total,omitempty"`
	// GoalsAgainstMinute61To75Percentage holds the value of the "goalsAgainstMinute61To75Percentage" field.
	GoalsAgainstMinute61To75Percentage string `json:"goalsAgainstMinute61To75Percentage,omitempty"`
	// GoalsAgainstMinute76To90Total holds the value of the "goalsAgainstMinute76To90Total" field.
	GoalsAgainstMinute76To90Total int `json:"goalsAgainstMinute76To90Total,omitempty"`
	// GoalsAgainstMinute76To90Percentage holds the value of the "goalsAgainstMinute76To90Percentage" field.
	GoalsAgainstMinute76To90Percentage string `json:"goalsAgainstMinute76To90Percentage,omitempty"`
	// GoalsAgainstMinute91To105Total holds the value of the "goalsAgainstMinute91To105Total" field.
	GoalsAgainstMinute91To105Total int `json:"goalsAgainstMinute91To105Total,omitempty"`
	// GoalsAgainstMinute91To105Percentage holds the value of the "goalsAgainstMinute91To105Percentage" field.
	GoalsAgainstMinute91To105Percentage string `json:"goalsAgainstMinute91To105Percentage,omitempty"`
	// GoalsAgainstMinute106To120Total holds the value of the "goalsAgainstMinute106To120Total" field.
	GoalsAgainstMinute106To120Total int `json:"goalsAgainstMinute106To120Total,omitempty"`
	// GoalsAgainstMinute106To120Percentage holds the value of the "goalsAgainstMinute106To120Percentage" field.
	GoalsAgainstMinute106To120Percentage string `json:"goalsAgainstMinute106To120Percentage,omitempty"`
	// LastUpdated holds the value of the "lastUpdated" field.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TSGoalsQuery when eager-loading is set.
	Edges            TSGoalsEdges `json:"edges"`
	team_goals_stats *int
	selectValues     sql.SelectValues
}

// TSGoalsEdges holds the relations/edges for other nodes in the graph.
type TSGoalsEdges struct {
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TSGoalsEdges) TeamOrErr() (*Team, error) {
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
func (*TSGoals) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tsgoals.FieldID, tsgoals.FieldGoalsForTotalHome, tsgoals.FieldGoalsForTotalAway, tsgoals.FieldGoalsForTotal, tsgoals.FieldGoalsForMinute0To15Total, tsgoals.FieldGoalsForMinute16To30Total, tsgoals.FieldGoalsForMinute31To45Total, tsgoals.FieldGoalsForMinute46To60Total, tsgoals.FieldGoalsForMinute61To75Total, tsgoals.FieldGoalsForMinute76To90Total, tsgoals.FieldGoalsForMinute91To105Total, tsgoals.FieldGoalsForMinute106To120Total, tsgoals.FieldGoalsAgainstTotalHome, tsgoals.FieldGoalsAgainstTotalAway, tsgoals.FieldGoalsAgainstTotal, tsgoals.FieldGoalsAgainstMinute0To15Total, tsgoals.FieldGoalsAgainstMinute16To30Total, tsgoals.FieldGoalsAgainstMinute31To45Total, tsgoals.FieldGoalsAgainstMinute46To60Total, tsgoals.FieldGoalsAgainstMinute61To75Total, tsgoals.FieldGoalsAgainstMinute76To90Total, tsgoals.FieldGoalsAgainstMinute91To105Total, tsgoals.FieldGoalsAgainstMinute106To120Total:
			values[i] = new(sql.NullInt64)
		case tsgoals.FieldGoalsForAverageHome, tsgoals.FieldGoalsForAverageAway, tsgoals.FieldGoalsForAverageTotal, tsgoals.FieldGoalsForMinute0To15Percentage, tsgoals.FieldGoalsForMinute16To30Percentage, tsgoals.FieldGoalsForMinute31To45Percentage, tsgoals.FieldGoalsForMinute46To60Percentage, tsgoals.FieldGoalsForMinute61To75Percentage, tsgoals.FieldGoalsForMinute76To90Percentage, tsgoals.FieldGoalsForMinute91To105Percentage, tsgoals.FieldGoalsForMinute106To120Percentage, tsgoals.FieldGoalsAgainstAverageHome, tsgoals.FieldGoalsAgainstAverageAway, tsgoals.FieldGoalsAgainstAverageTotal, tsgoals.FieldGoalsAgainstMinute0To15Percentage, tsgoals.FieldGoalsAgainstMinute16To30Percentage, tsgoals.FieldGoalsAgainstMinute31To45Percentage, tsgoals.FieldGoalsAgainstMinute46To60Percentage, tsgoals.FieldGoalsAgainstMinute61To75Percentage, tsgoals.FieldGoalsAgainstMinute76To90Percentage, tsgoals.FieldGoalsAgainstMinute91To105Percentage, tsgoals.FieldGoalsAgainstMinute106To120Percentage:
			values[i] = new(sql.NullString)
		case tsgoals.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case tsgoals.ForeignKeys[0]: // team_goals_stats
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TSGoals fields.
func (tg *TSGoals) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tsgoals.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tg.ID = int(value.Int64)
		case tsgoals.FieldGoalsForTotalHome:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForTotalHome", values[i])
			} else if value.Valid {
				tg.GoalsForTotalHome = int(value.Int64)
			}
		case tsgoals.FieldGoalsForTotalAway:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForTotalAway", values[i])
			} else if value.Valid {
				tg.GoalsForTotalAway = int(value.Int64)
			}
		case tsgoals.FieldGoalsForTotal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForTotal", values[i])
			} else if value.Valid {
				tg.GoalsForTotal = int(value.Int64)
			}
		case tsgoals.FieldGoalsForAverageHome:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForAverageHome", values[i])
			} else if value.Valid {
				tg.GoalsForAverageHome = value.String
			}
		case tsgoals.FieldGoalsForAverageAway:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForAverageAway", values[i])
			} else if value.Valid {
				tg.GoalsForAverageAway = value.String
			}
		case tsgoals.FieldGoalsForAverageTotal:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForAverageTotal", values[i])
			} else if value.Valid {
				tg.GoalsForAverageTotal = value.String
			}
		case tsgoals.FieldGoalsForMinute0To15Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute0To15Total", values[i])
			} else if value.Valid {
				tg.GoalsForMinute0To15Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsForMinute0To15Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute0To15Percentage", values[i])
			} else if value.Valid {
				tg.GoalsForMinute0To15Percentage = value.String
			}
		case tsgoals.FieldGoalsForMinute16To30Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute16To30Total", values[i])
			} else if value.Valid {
				tg.GoalsForMinute16To30Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsForMinute16To30Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute16To30Percentage", values[i])
			} else if value.Valid {
				tg.GoalsForMinute16To30Percentage = value.String
			}
		case tsgoals.FieldGoalsForMinute31To45Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute31To45Total", values[i])
			} else if value.Valid {
				tg.GoalsForMinute31To45Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsForMinute31To45Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute31To45Percentage", values[i])
			} else if value.Valid {
				tg.GoalsForMinute31To45Percentage = value.String
			}
		case tsgoals.FieldGoalsForMinute46To60Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute46To60Total", values[i])
			} else if value.Valid {
				tg.GoalsForMinute46To60Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsForMinute46To60Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute46To60Percentage", values[i])
			} else if value.Valid {
				tg.GoalsForMinute46To60Percentage = value.String
			}
		case tsgoals.FieldGoalsForMinute61To75Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute61To75Total", values[i])
			} else if value.Valid {
				tg.GoalsForMinute61To75Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsForMinute61To75Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute61To75Percentage", values[i])
			} else if value.Valid {
				tg.GoalsForMinute61To75Percentage = value.String
			}
		case tsgoals.FieldGoalsForMinute76To90Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute76To90Total", values[i])
			} else if value.Valid {
				tg.GoalsForMinute76To90Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsForMinute76To90Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute76To90Percentage", values[i])
			} else if value.Valid {
				tg.GoalsForMinute76To90Percentage = value.String
			}
		case tsgoals.FieldGoalsForMinute91To105Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute91To105Total", values[i])
			} else if value.Valid {
				tg.GoalsForMinute91To105Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsForMinute91To105Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute91To105Percentage", values[i])
			} else if value.Valid {
				tg.GoalsForMinute91To105Percentage = value.String
			}
		case tsgoals.FieldGoalsForMinute106To120Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute106To120Total", values[i])
			} else if value.Valid {
				tg.GoalsForMinute106To120Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsForMinute106To120Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsForMinute106To120Percentage", values[i])
			} else if value.Valid {
				tg.GoalsForMinute106To120Percentage = value.String
			}
		case tsgoals.FieldGoalsAgainstTotalHome:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstTotalHome", values[i])
			} else if value.Valid {
				tg.GoalsAgainstTotalHome = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstTotalAway:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstTotalAway", values[i])
			} else if value.Valid {
				tg.GoalsAgainstTotalAway = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstTotal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstTotal", values[i])
			} else if value.Valid {
				tg.GoalsAgainstTotal = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstAverageHome:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstAverageHome", values[i])
			} else if value.Valid {
				tg.GoalsAgainstAverageHome = value.String
			}
		case tsgoals.FieldGoalsAgainstAverageAway:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstAverageAway", values[i])
			} else if value.Valid {
				tg.GoalsAgainstAverageAway = value.String
			}
		case tsgoals.FieldGoalsAgainstAverageTotal:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstAverageTotal", values[i])
			} else if value.Valid {
				tg.GoalsAgainstAverageTotal = value.String
			}
		case tsgoals.FieldGoalsAgainstMinute0To15Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute0To15Total", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute0To15Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstMinute0To15Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute0To15Percentage", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute0To15Percentage = value.String
			}
		case tsgoals.FieldGoalsAgainstMinute16To30Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute16To30Total", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute16To30Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstMinute16To30Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute16To30Percentage", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute16To30Percentage = value.String
			}
		case tsgoals.FieldGoalsAgainstMinute31To45Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute31To45Total", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute31To45Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstMinute31To45Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute31To45Percentage", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute31To45Percentage = value.String
			}
		case tsgoals.FieldGoalsAgainstMinute46To60Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute46To60Total", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute46To60Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstMinute46To60Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute46To60Percentage", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute46To60Percentage = value.String
			}
		case tsgoals.FieldGoalsAgainstMinute61To75Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute61To75Total", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute61To75Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstMinute61To75Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute61To75Percentage", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute61To75Percentage = value.String
			}
		case tsgoals.FieldGoalsAgainstMinute76To90Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute76To90Total", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute76To90Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstMinute76To90Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute76To90Percentage", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute76To90Percentage = value.String
			}
		case tsgoals.FieldGoalsAgainstMinute91To105Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute91To105Total", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute91To105Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstMinute91To105Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute91To105Percentage", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute91To105Percentage = value.String
			}
		case tsgoals.FieldGoalsAgainstMinute106To120Total:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute106To120Total", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute106To120Total = int(value.Int64)
			}
		case tsgoals.FieldGoalsAgainstMinute106To120Percentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goalsAgainstMinute106To120Percentage", values[i])
			} else if value.Valid {
				tg.GoalsAgainstMinute106To120Percentage = value.String
			}
		case tsgoals.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastUpdated", values[i])
			} else if value.Valid {
				tg.LastUpdated = value.Time
			}
		case tsgoals.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_goals_stats", value)
			} else if value.Valid {
				tg.team_goals_stats = new(int)
				*tg.team_goals_stats = int(value.Int64)
			}
		default:
			tg.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TSGoals.
// This includes values selected through modifiers, order, etc.
func (tg *TSGoals) Value(name string) (ent.Value, error) {
	return tg.selectValues.Get(name)
}

// QueryTeam queries the "team" edge of the TSGoals entity.
func (tg *TSGoals) QueryTeam() *TeamQuery {
	return NewTSGoalsClient(tg.config).QueryTeam(tg)
}

// Update returns a builder for updating this TSGoals.
// Note that you need to call TSGoals.Unwrap() before calling this method if this TSGoals
// was returned from a transaction, and the transaction was committed or rolled back.
func (tg *TSGoals) Update() *TSGoalsUpdateOne {
	return NewTSGoalsClient(tg.config).UpdateOne(tg)
}

// Unwrap unwraps the TSGoals entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tg *TSGoals) Unwrap() *TSGoals {
	_tx, ok := tg.config.driver.(*txDriver)
	if !ok {
		panic("ent: TSGoals is not a transactional entity")
	}
	tg.config.driver = _tx.drv
	return tg
}

// String implements the fmt.Stringer.
func (tg *TSGoals) String() string {
	var builder strings.Builder
	builder.WriteString("TSGoals(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tg.ID))
	builder.WriteString("goalsForTotalHome=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForTotalHome))
	builder.WriteString(", ")
	builder.WriteString("goalsForTotalAway=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForTotalAway))
	builder.WriteString(", ")
	builder.WriteString("goalsForTotal=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForTotal))
	builder.WriteString(", ")
	builder.WriteString("goalsForAverageHome=")
	builder.WriteString(tg.GoalsForAverageHome)
	builder.WriteString(", ")
	builder.WriteString("goalsForAverageAway=")
	builder.WriteString(tg.GoalsForAverageAway)
	builder.WriteString(", ")
	builder.WriteString("goalsForAverageTotal=")
	builder.WriteString(tg.GoalsForAverageTotal)
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute0To15Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForMinute0To15Total))
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute0To15Percentage=")
	builder.WriteString(tg.GoalsForMinute0To15Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute16To30Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForMinute16To30Total))
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute16To30Percentage=")
	builder.WriteString(tg.GoalsForMinute16To30Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute31To45Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForMinute31To45Total))
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute31To45Percentage=")
	builder.WriteString(tg.GoalsForMinute31To45Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute46To60Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForMinute46To60Total))
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute46To60Percentage=")
	builder.WriteString(tg.GoalsForMinute46To60Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute61To75Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForMinute61To75Total))
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute61To75Percentage=")
	builder.WriteString(tg.GoalsForMinute61To75Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute76To90Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForMinute76To90Total))
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute76To90Percentage=")
	builder.WriteString(tg.GoalsForMinute76To90Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute91To105Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForMinute91To105Total))
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute91To105Percentage=")
	builder.WriteString(tg.GoalsForMinute91To105Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute106To120Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsForMinute106To120Total))
	builder.WriteString(", ")
	builder.WriteString("goalsForMinute106To120Percentage=")
	builder.WriteString(tg.GoalsForMinute106To120Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstTotalHome=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstTotalHome))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstTotalAway=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstTotalAway))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstTotal=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstTotal))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstAverageHome=")
	builder.WriteString(tg.GoalsAgainstAverageHome)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstAverageAway=")
	builder.WriteString(tg.GoalsAgainstAverageAway)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstAverageTotal=")
	builder.WriteString(tg.GoalsAgainstAverageTotal)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute0To15Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstMinute0To15Total))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute0To15Percentage=")
	builder.WriteString(tg.GoalsAgainstMinute0To15Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute16To30Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstMinute16To30Total))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute16To30Percentage=")
	builder.WriteString(tg.GoalsAgainstMinute16To30Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute31To45Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstMinute31To45Total))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute31To45Percentage=")
	builder.WriteString(tg.GoalsAgainstMinute31To45Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute46To60Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstMinute46To60Total))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute46To60Percentage=")
	builder.WriteString(tg.GoalsAgainstMinute46To60Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute61To75Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstMinute61To75Total))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute61To75Percentage=")
	builder.WriteString(tg.GoalsAgainstMinute61To75Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute76To90Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstMinute76To90Total))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute76To90Percentage=")
	builder.WriteString(tg.GoalsAgainstMinute76To90Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute91To105Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstMinute91To105Total))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute91To105Percentage=")
	builder.WriteString(tg.GoalsAgainstMinute91To105Percentage)
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute106To120Total=")
	builder.WriteString(fmt.Sprintf("%v", tg.GoalsAgainstMinute106To120Total))
	builder.WriteString(", ")
	builder.WriteString("goalsAgainstMinute106To120Percentage=")
	builder.WriteString(tg.GoalsAgainstMinute106To120Percentage)
	builder.WriteString(", ")
	builder.WriteString("lastUpdated=")
	builder.WriteString(tg.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TSGoalsSlice is a parsable slice of TSGoals.
type TSGoalsSlice []*TSGoals
