// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Season is the model entity for the Season schema.
type Season struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Year holds the value of the "year" field.
	Year int `json:"year,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// Current holds the value of the "current" field.
	Current bool `json:"current,omitempty"`
	// LastUpdated holds the value of the "lastUpdated" field.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SeasonQuery when eager-loading is set.
	Edges         SeasonEdges `json:"edges"`
	league_season *int
	selectValues  sql.SelectValues
}

// SeasonEdges holds the relations/edges for other nodes in the graph.
type SeasonEdges struct {
	// League holds the value of the league edge.
	League *League `json:"league,omitempty"`
	// Fixtures holds the value of the fixtures edge.
	Fixtures []*Fixture `json:"fixtures,omitempty"`
	// Standings holds the value of the standings edge.
	Standings []*Standings `json:"standings,omitempty"`
	// Teams holds the value of the teams edge.
	Teams []*Team `json:"teams,omitempty"`
	// Squad holds the value of the squad edge.
	Squad []*Squad `json:"squad,omitempty"`
	// PlayerStats holds the value of the playerStats edge.
	PlayerStats []*PlayerStats `json:"playerStats,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// LeagueOrErr returns the League value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SeasonEdges) LeagueOrErr() (*League, error) {
	if e.loadedTypes[0] {
		if e.League == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: league.Label}
		}
		return e.League, nil
	}
	return nil, &NotLoadedError{edge: "league"}
}

// FixturesOrErr returns the Fixtures value or an error if the edge
// was not loaded in eager-loading.
func (e SeasonEdges) FixturesOrErr() ([]*Fixture, error) {
	if e.loadedTypes[1] {
		return e.Fixtures, nil
	}
	return nil, &NotLoadedError{edge: "fixtures"}
}

// StandingsOrErr returns the Standings value or an error if the edge
// was not loaded in eager-loading.
func (e SeasonEdges) StandingsOrErr() ([]*Standings, error) {
	if e.loadedTypes[2] {
		return e.Standings, nil
	}
	return nil, &NotLoadedError{edge: "standings"}
}

// TeamsOrErr returns the Teams value or an error if the edge
// was not loaded in eager-loading.
func (e SeasonEdges) TeamsOrErr() ([]*Team, error) {
	if e.loadedTypes[3] {
		return e.Teams, nil
	}
	return nil, &NotLoadedError{edge: "teams"}
}

// SquadOrErr returns the Squad value or an error if the edge
// was not loaded in eager-loading.
func (e SeasonEdges) SquadOrErr() ([]*Squad, error) {
	if e.loadedTypes[4] {
		return e.Squad, nil
	}
	return nil, &NotLoadedError{edge: "squad"}
}

// PlayerStatsOrErr returns the PlayerStats value or an error if the edge
// was not loaded in eager-loading.
func (e SeasonEdges) PlayerStatsOrErr() ([]*PlayerStats, error) {
	if e.loadedTypes[5] {
		return e.PlayerStats, nil
	}
	return nil, &NotLoadedError{edge: "playerStats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Season) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case season.FieldCurrent:
			values[i] = new(sql.NullBool)
		case season.FieldID, season.FieldYear:
			values[i] = new(sql.NullInt64)
		case season.FieldSlug:
			values[i] = new(sql.NullString)
		case season.FieldStartDate, season.FieldEndDate, season.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case season.ForeignKeys[0]: // league_season
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Season fields.
func (s *Season) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case season.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case season.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				s.Slug = value.String
			}
		case season.FieldYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field year", values[i])
			} else if value.Valid {
				s.Year = int(value.Int64)
			}
		case season.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				s.StartDate = value.Time
			}
		case season.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				s.EndDate = value.Time
			}
		case season.FieldCurrent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field current", values[i])
			} else if value.Valid {
				s.Current = value.Bool
			}
		case season.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastUpdated", values[i])
			} else if value.Valid {
				s.LastUpdated = value.Time
			}
		case season.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field league_season", value)
			} else if value.Valid {
				s.league_season = new(int)
				*s.league_season = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Season.
// This includes values selected through modifiers, order, etc.
func (s *Season) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryLeague queries the "league" edge of the Season entity.
func (s *Season) QueryLeague() *LeagueQuery {
	return NewSeasonClient(s.config).QueryLeague(s)
}

// QueryFixtures queries the "fixtures" edge of the Season entity.
func (s *Season) QueryFixtures() *FixtureQuery {
	return NewSeasonClient(s.config).QueryFixtures(s)
}

// QueryStandings queries the "standings" edge of the Season entity.
func (s *Season) QueryStandings() *StandingsQuery {
	return NewSeasonClient(s.config).QueryStandings(s)
}

// QueryTeams queries the "teams" edge of the Season entity.
func (s *Season) QueryTeams() *TeamQuery {
	return NewSeasonClient(s.config).QueryTeams(s)
}

// QuerySquad queries the "squad" edge of the Season entity.
func (s *Season) QuerySquad() *SquadQuery {
	return NewSeasonClient(s.config).QuerySquad(s)
}

// QueryPlayerStats queries the "playerStats" edge of the Season entity.
func (s *Season) QueryPlayerStats() *PlayerStatsQuery {
	return NewSeasonClient(s.config).QueryPlayerStats(s)
}

// Update returns a builder for updating this Season.
// Note that you need to call Season.Unwrap() before calling this method if this Season
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Season) Update() *SeasonUpdateOne {
	return NewSeasonClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Season entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Season) Unwrap() *Season {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Season is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Season) String() string {
	var builder strings.Builder
	builder.WriteString("Season(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("slug=")
	builder.WriteString(s.Slug)
	builder.WriteString(", ")
	builder.WriteString("year=")
	builder.WriteString(fmt.Sprintf("%v", s.Year))
	builder.WriteString(", ")
	builder.WriteString("start_date=")
	builder.WriteString(s.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_date=")
	builder.WriteString(s.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("current=")
	builder.WriteString(fmt.Sprintf("%v", s.Current))
	builder.WriteString(", ")
	builder.WriteString("lastUpdated=")
	builder.WriteString(s.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Seasons is a parsable slice of Season.
type Seasons []*Season
