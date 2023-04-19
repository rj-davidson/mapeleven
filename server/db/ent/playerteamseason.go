// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mapeleven/db/ent/player"
	"mapeleven/db/ent/playerteamseason"
	"mapeleven/db/ent/teamseason"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PlayerTeamSeason is the model entity for the PlayerTeamSeason schema.
type PlayerTeamSeason struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PlayerTeamSeasonID holds the value of the "player_team_season_id" field.
	PlayerTeamSeasonID int `json:"player_team_season_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerTeamSeasonQuery when eager-loading is set.
	Edges                           PlayerTeamSeasonEdges `json:"edges"`
	player_player_team_seasons      *int
	team_season_player_team_seasons *int
	selectValues                    sql.SelectValues
}

// PlayerTeamSeasonEdges holds the relations/edges for other nodes in the graph.
type PlayerTeamSeasonEdges struct {
	// Player holds the value of the player edge.
	Player *Player `json:"player,omitempty"`
	// TeamSeason holds the value of the teamSeason edge.
	TeamSeason *TeamSeason `json:"teamSeason,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PlayerOrErr returns the Player value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerTeamSeasonEdges) PlayerOrErr() (*Player, error) {
	if e.loadedTypes[0] {
		if e.Player == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: player.Label}
		}
		return e.Player, nil
	}
	return nil, &NotLoadedError{edge: "player"}
}

// TeamSeasonOrErr returns the TeamSeason value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerTeamSeasonEdges) TeamSeasonOrErr() (*TeamSeason, error) {
	if e.loadedTypes[1] {
		if e.TeamSeason == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: teamseason.Label}
		}
		return e.TeamSeason, nil
	}
	return nil, &NotLoadedError{edge: "teamSeason"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlayerTeamSeason) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case playerteamseason.FieldID, playerteamseason.FieldPlayerTeamSeasonID:
			values[i] = new(sql.NullInt64)
		case playerteamseason.ForeignKeys[0]: // player_player_team_seasons
			values[i] = new(sql.NullInt64)
		case playerteamseason.ForeignKeys[1]: // team_season_player_team_seasons
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlayerTeamSeason fields.
func (pts *PlayerTeamSeason) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case playerteamseason.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pts.ID = int(value.Int64)
		case playerteamseason.FieldPlayerTeamSeasonID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field player_team_season_id", values[i])
			} else if value.Valid {
				pts.PlayerTeamSeasonID = int(value.Int64)
			}
		case playerteamseason.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field player_player_team_seasons", value)
			} else if value.Valid {
				pts.player_player_team_seasons = new(int)
				*pts.player_player_team_seasons = int(value.Int64)
			}
		case playerteamseason.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_season_player_team_seasons", value)
			} else if value.Valid {
				pts.team_season_player_team_seasons = new(int)
				*pts.team_season_player_team_seasons = int(value.Int64)
			}
		default:
			pts.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PlayerTeamSeason.
// This includes values selected through modifiers, order, etc.
func (pts *PlayerTeamSeason) Value(name string) (ent.Value, error) {
	return pts.selectValues.Get(name)
}

// QueryPlayer queries the "player" edge of the PlayerTeamSeason entity.
func (pts *PlayerTeamSeason) QueryPlayer() *PlayerQuery {
	return NewPlayerTeamSeasonClient(pts.config).QueryPlayer(pts)
}

// QueryTeamSeason queries the "teamSeason" edge of the PlayerTeamSeason entity.
func (pts *PlayerTeamSeason) QueryTeamSeason() *TeamSeasonQuery {
	return NewPlayerTeamSeasonClient(pts.config).QueryTeamSeason(pts)
}

// Update returns a builder for updating this PlayerTeamSeason.
// Note that you need to call PlayerTeamSeason.Unwrap() before calling this method if this PlayerTeamSeason
// was returned from a transaction, and the transaction was committed or rolled back.
func (pts *PlayerTeamSeason) Update() *PlayerTeamSeasonUpdateOne {
	return NewPlayerTeamSeasonClient(pts.config).UpdateOne(pts)
}

// Unwrap unwraps the PlayerTeamSeason entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pts *PlayerTeamSeason) Unwrap() *PlayerTeamSeason {
	_tx, ok := pts.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlayerTeamSeason is not a transactional entity")
	}
	pts.config.driver = _tx.drv
	return pts
}

// String implements the fmt.Stringer.
func (pts *PlayerTeamSeason) String() string {
	var builder strings.Builder
	builder.WriteString("PlayerTeamSeason(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pts.ID))
	builder.WriteString("player_team_season_id=")
	builder.WriteString(fmt.Sprintf("%v", pts.PlayerTeamSeasonID))
	builder.WriteByte(')')
	return builder.String()
}

// PlayerTeamSeasons is a parsable slice of PlayerTeamSeason.
type PlayerTeamSeasons []*PlayerTeamSeason
