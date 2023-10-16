// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mapeleven/db/ent/club"
	"mapeleven/db/ent/season"
	"mapeleven/db/ent/team"
	"mapeleven/db/ent/tsbiggest"
	"mapeleven/db/ent/tscards"
	"mapeleven/db/ent/tscleansheet"
	"mapeleven/db/ent/tsfailedtoscore"
	"mapeleven/db/ent/tsfixtures"
	"mapeleven/db/ent/tsgoals"
	"mapeleven/db/ent/tspenalty"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Team is the model entity for the Team schema.
type Team struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Form holds the value of the "form" field.
	Form string `json:"form,omitempty"`
	// LastUpdated holds the value of the "lastUpdated" field.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeamQuery when eager-loading is set.
	Edges        TeamEdges `json:"edges"`
	club_team    *int
	season_teams *int
	selectValues sql.SelectValues
}

// TeamEdges holds the relations/edges for other nodes in the graph.
type TeamEdges struct {
	// Season holds the value of the season edge.
	Season *Season `json:"season,omitempty"`
	// Club holds the value of the club edge.
	Club *Club `json:"club,omitempty"`
	// Standings holds the value of the standings edge.
	Standings []*Standings `json:"standings,omitempty"`
	// HomeFixtures holds the value of the homeFixtures edge.
	HomeFixtures []*Fixture `json:"homeFixtures,omitempty"`
	// AwayFixtures holds the value of the awayFixtures edge.
	AwayFixtures []*Fixture `json:"awayFixtures,omitempty"`
	// FixtureEvents holds the value of the fixtureEvents edge.
	FixtureEvents []*FixtureEvents `json:"fixtureEvents,omitempty"`
	// FixtureLineups holds the value of the fixtureLineups edge.
	FixtureLineups []*FixtureLineups `json:"fixtureLineups,omitempty"`
	// Players holds the value of the players edge.
	Players []*Player `json:"players,omitempty"`
	// Squad holds the value of the squad edge.
	Squad []*Squad `json:"squad,omitempty"`
	// BiggestStats holds the value of the biggest_stats edge.
	BiggestStats *TSBiggest `json:"biggest_stats,omitempty"`
	// CardsStats holds the value of the cards_stats edge.
	CardsStats *TSCards `json:"cards_stats,omitempty"`
	// CleanSheetStats holds the value of the clean_sheet_stats edge.
	CleanSheetStats *TSCleanSheet `json:"clean_sheet_stats,omitempty"`
	// FailedToScoreStats holds the value of the failed_to_score_stats edge.
	FailedToScoreStats *TSFailedToScore `json:"failed_to_score_stats,omitempty"`
	// FixturesStats holds the value of the fixtures_stats edge.
	FixturesStats *TSFixtures `json:"fixtures_stats,omitempty"`
	// GoalsStats holds the value of the goals_stats edge.
	GoalsStats *TSGoals `json:"goals_stats,omitempty"`
	// Lineups holds the value of the lineups edge.
	Lineups []*TSLineups `json:"lineups,omitempty"`
	// PenaltyStats holds the value of the penalty_stats edge.
	PenaltyStats *TSPenalty `json:"penalty_stats,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [17]bool
}

// SeasonOrErr returns the Season value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) SeasonOrErr() (*Season, error) {
	if e.loadedTypes[0] {
		if e.Season == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: season.Label}
		}
		return e.Season, nil
	}
	return nil, &NotLoadedError{edge: "season"}
}

// ClubOrErr returns the Club value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) ClubOrErr() (*Club, error) {
	if e.loadedTypes[1] {
		if e.Club == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: club.Label}
		}
		return e.Club, nil
	}
	return nil, &NotLoadedError{edge: "club"}
}

// StandingsOrErr returns the Standings value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) StandingsOrErr() ([]*Standings, error) {
	if e.loadedTypes[2] {
		return e.Standings, nil
	}
	return nil, &NotLoadedError{edge: "standings"}
}

// HomeFixturesOrErr returns the HomeFixtures value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) HomeFixturesOrErr() ([]*Fixture, error) {
	if e.loadedTypes[3] {
		return e.HomeFixtures, nil
	}
	return nil, &NotLoadedError{edge: "homeFixtures"}
}

// AwayFixturesOrErr returns the AwayFixtures value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) AwayFixturesOrErr() ([]*Fixture, error) {
	if e.loadedTypes[4] {
		return e.AwayFixtures, nil
	}
	return nil, &NotLoadedError{edge: "awayFixtures"}
}

// FixtureEventsOrErr returns the FixtureEvents value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) FixtureEventsOrErr() ([]*FixtureEvents, error) {
	if e.loadedTypes[5] {
		return e.FixtureEvents, nil
	}
	return nil, &NotLoadedError{edge: "fixtureEvents"}
}

// FixtureLineupsOrErr returns the FixtureLineups value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) FixtureLineupsOrErr() ([]*FixtureLineups, error) {
	if e.loadedTypes[6] {
		return e.FixtureLineups, nil
	}
	return nil, &NotLoadedError{edge: "fixtureLineups"}
}

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[7] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// SquadOrErr returns the Squad value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) SquadOrErr() ([]*Squad, error) {
	if e.loadedTypes[8] {
		return e.Squad, nil
	}
	return nil, &NotLoadedError{edge: "squad"}
}

// BiggestStatsOrErr returns the BiggestStats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) BiggestStatsOrErr() (*TSBiggest, error) {
	if e.loadedTypes[9] {
		if e.BiggestStats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tsbiggest.Label}
		}
		return e.BiggestStats, nil
	}
	return nil, &NotLoadedError{edge: "biggest_stats"}
}

// CardsStatsOrErr returns the CardsStats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) CardsStatsOrErr() (*TSCards, error) {
	if e.loadedTypes[10] {
		if e.CardsStats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tscards.Label}
		}
		return e.CardsStats, nil
	}
	return nil, &NotLoadedError{edge: "cards_stats"}
}

// CleanSheetStatsOrErr returns the CleanSheetStats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) CleanSheetStatsOrErr() (*TSCleanSheet, error) {
	if e.loadedTypes[11] {
		if e.CleanSheetStats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tscleansheet.Label}
		}
		return e.CleanSheetStats, nil
	}
	return nil, &NotLoadedError{edge: "clean_sheet_stats"}
}

// FailedToScoreStatsOrErr returns the FailedToScoreStats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) FailedToScoreStatsOrErr() (*TSFailedToScore, error) {
	if e.loadedTypes[12] {
		if e.FailedToScoreStats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tsfailedtoscore.Label}
		}
		return e.FailedToScoreStats, nil
	}
	return nil, &NotLoadedError{edge: "failed_to_score_stats"}
}

// FixturesStatsOrErr returns the FixturesStats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) FixturesStatsOrErr() (*TSFixtures, error) {
	if e.loadedTypes[13] {
		if e.FixturesStats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tsfixtures.Label}
		}
		return e.FixturesStats, nil
	}
	return nil, &NotLoadedError{edge: "fixtures_stats"}
}

// GoalsStatsOrErr returns the GoalsStats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) GoalsStatsOrErr() (*TSGoals, error) {
	if e.loadedTypes[14] {
		if e.GoalsStats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tsgoals.Label}
		}
		return e.GoalsStats, nil
	}
	return nil, &NotLoadedError{edge: "goals_stats"}
}

// LineupsOrErr returns the Lineups value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) LineupsOrErr() ([]*TSLineups, error) {
	if e.loadedTypes[15] {
		return e.Lineups, nil
	}
	return nil, &NotLoadedError{edge: "lineups"}
}

// PenaltyStatsOrErr returns the PenaltyStats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) PenaltyStatsOrErr() (*TSPenalty, error) {
	if e.loadedTypes[16] {
		if e.PenaltyStats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tspenalty.Label}
		}
		return e.PenaltyStats, nil
	}
	return nil, &NotLoadedError{edge: "penalty_stats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Team) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			values[i] = new(sql.NullInt64)
		case team.FieldForm:
			values[i] = new(sql.NullString)
		case team.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case team.ForeignKeys[0]: // club_team
			values[i] = new(sql.NullInt64)
		case team.ForeignKeys[1]: // season_teams
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Team fields.
func (t *Team) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case team.FieldForm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field form", values[i])
			} else if value.Valid {
				t.Form = value.String
			}
		case team.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastUpdated", values[i])
			} else if value.Valid {
				t.LastUpdated = value.Time
			}
		case team.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field club_team", value)
			} else if value.Valid {
				t.club_team = new(int)
				*t.club_team = int(value.Int64)
			}
		case team.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field season_teams", value)
			} else if value.Valid {
				t.season_teams = new(int)
				*t.season_teams = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Team.
// This includes values selected through modifiers, order, etc.
func (t *Team) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QuerySeason queries the "season" edge of the Team entity.
func (t *Team) QuerySeason() *SeasonQuery {
	return NewTeamClient(t.config).QuerySeason(t)
}

// QueryClub queries the "club" edge of the Team entity.
func (t *Team) QueryClub() *ClubQuery {
	return NewTeamClient(t.config).QueryClub(t)
}

// QueryStandings queries the "standings" edge of the Team entity.
func (t *Team) QueryStandings() *StandingsQuery {
	return NewTeamClient(t.config).QueryStandings(t)
}

// QueryHomeFixtures queries the "homeFixtures" edge of the Team entity.
func (t *Team) QueryHomeFixtures() *FixtureQuery {
	return NewTeamClient(t.config).QueryHomeFixtures(t)
}

// QueryAwayFixtures queries the "awayFixtures" edge of the Team entity.
func (t *Team) QueryAwayFixtures() *FixtureQuery {
	return NewTeamClient(t.config).QueryAwayFixtures(t)
}

// QueryFixtureEvents queries the "fixtureEvents" edge of the Team entity.
func (t *Team) QueryFixtureEvents() *FixtureEventsQuery {
	return NewTeamClient(t.config).QueryFixtureEvents(t)
}

// QueryFixtureLineups queries the "fixtureLineups" edge of the Team entity.
func (t *Team) QueryFixtureLineups() *FixtureLineupsQuery {
	return NewTeamClient(t.config).QueryFixtureLineups(t)
}

// QueryPlayers queries the "players" edge of the Team entity.
func (t *Team) QueryPlayers() *PlayerQuery {
	return NewTeamClient(t.config).QueryPlayers(t)
}

// QuerySquad queries the "squad" edge of the Team entity.
func (t *Team) QuerySquad() *SquadQuery {
	return NewTeamClient(t.config).QuerySquad(t)
}

// QueryBiggestStats queries the "biggest_stats" edge of the Team entity.
func (t *Team) QueryBiggestStats() *TSBiggestQuery {
	return NewTeamClient(t.config).QueryBiggestStats(t)
}

// QueryCardsStats queries the "cards_stats" edge of the Team entity.
func (t *Team) QueryCardsStats() *TSCardsQuery {
	return NewTeamClient(t.config).QueryCardsStats(t)
}

// QueryCleanSheetStats queries the "clean_sheet_stats" edge of the Team entity.
func (t *Team) QueryCleanSheetStats() *TSCleanSheetQuery {
	return NewTeamClient(t.config).QueryCleanSheetStats(t)
}

// QueryFailedToScoreStats queries the "failed_to_score_stats" edge of the Team entity.
func (t *Team) QueryFailedToScoreStats() *TSFailedToScoreQuery {
	return NewTeamClient(t.config).QueryFailedToScoreStats(t)
}

// QueryFixturesStats queries the "fixtures_stats" edge of the Team entity.
func (t *Team) QueryFixturesStats() *TSFixturesQuery {
	return NewTeamClient(t.config).QueryFixturesStats(t)
}

// QueryGoalsStats queries the "goals_stats" edge of the Team entity.
func (t *Team) QueryGoalsStats() *TSGoalsQuery {
	return NewTeamClient(t.config).QueryGoalsStats(t)
}

// QueryLineups queries the "lineups" edge of the Team entity.
func (t *Team) QueryLineups() *TSLineupsQuery {
	return NewTeamClient(t.config).QueryLineups(t)
}

// QueryPenaltyStats queries the "penalty_stats" edge of the Team entity.
func (t *Team) QueryPenaltyStats() *TSPenaltyQuery {
	return NewTeamClient(t.config).QueryPenaltyStats(t)
}

// Update returns a builder for updating this Team.
// Note that you need to call Team.Unwrap() before calling this method if this Team
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Team) Update() *TeamUpdateOne {
	return NewTeamClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Team entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Team) Unwrap() *Team {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Team is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Team) String() string {
	var builder strings.Builder
	builder.WriteString("Team(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("form=")
	builder.WriteString(t.Form)
	builder.WriteString(", ")
	builder.WriteString("lastUpdated=")
	builder.WriteString(t.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Teams is a parsable slice of Team.
type Teams []*Team
