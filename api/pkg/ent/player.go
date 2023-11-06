// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/birth"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/country"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Player is the model entity for the Player schema.
type Player struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// ApiFootballId holds the value of the "ApiFootballId" field.
	ApiFootballId int `json:"ApiFootballId,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Firstname holds the value of the "firstname" field.
	Firstname string `json:"firstname,omitempty"`
	// Lastname holds the value of the "lastname" field.
	Lastname string `json:"lastname,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Height holds the value of the "height" field.
	Height string `json:"height,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight string `json:"weight,omitempty"`
	// Injured holds the value of the "injured" field.
	Injured bool `json:"injured,omitempty"`
	// Photo holds the value of the "photo" field.
	Photo string `json:"photo,omitempty"`
	// LastUpdated holds the value of the "lastUpdated" field.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Form holds the value of the "form" field.
	Form string `json:"form,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerQuery when eager-loading is set.
	Edges           PlayerEdges `json:"edges"`
	birth_player    *int
	club_player     *int
	country_players *int
	league_player   *int
	season_player   *int
	team_players    *int
	selectValues    sql.SelectValues
}

// PlayerEdges holds the relations/edges for other nodes in the graph.
type PlayerEdges struct {
	// Birth holds the value of the birth edge.
	Birth *Birth `json:"birth,omitempty"`
	// Nationality holds the value of the nationality edge.
	Nationality *Country `json:"nationality,omitempty"`
	// Squad holds the value of the squad edge.
	Squad []*Squad `json:"squad,omitempty"`
	// PlayerEvents holds the value of the playerEvents edge.
	PlayerEvents []*FixtureEvents `json:"playerEvents,omitempty"`
	// MatchPlayer holds the value of the matchPlayer edge.
	MatchPlayer []*MatchPlayer `json:"matchPlayer,omitempty"`
	// AssistEvents holds the value of the assistEvents edge.
	AssistEvents []*FixtureEvents `json:"assistEvents,omitempty"`
	// Psgames holds the value of the psgames edge.
	Psgames []*PSGames `json:"psgames,omitempty"`
	// Psgoals holds the value of the psgoals edge.
	Psgoals []*PSGoals `json:"psgoals,omitempty"`
	// Psdefense holds the value of the psdefense edge.
	Psdefense []*PSDefense `json:"psdefense,omitempty"`
	// Psoffense holds the value of the psoffense edge.
	Psoffense []*PSOffense `json:"psoffense,omitempty"`
	// Pspenalty holds the value of the pspenalty edge.
	Pspenalty []*PSPenalty `json:"pspenalty,omitempty"`
	// Season holds the value of the season edge.
	Season *Season `json:"season,omitempty"`
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// Club holds the value of the club edge.
	Club *Club `json:"club,omitempty"`
	// League holds the value of the league edge.
	League *League `json:"league,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [15]bool
}

// BirthOrErr returns the Birth value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerEdges) BirthOrErr() (*Birth, error) {
	if e.loadedTypes[0] {
		if e.Birth == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: birth.Label}
		}
		return e.Birth, nil
	}
	return nil, &NotLoadedError{edge: "birth"}
}

// NationalityOrErr returns the Nationality value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerEdges) NationalityOrErr() (*Country, error) {
	if e.loadedTypes[1] {
		if e.Nationality == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: country.Label}
		}
		return e.Nationality, nil
	}
	return nil, &NotLoadedError{edge: "nationality"}
}

// SquadOrErr returns the Squad value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) SquadOrErr() ([]*Squad, error) {
	if e.loadedTypes[2] {
		return e.Squad, nil
	}
	return nil, &NotLoadedError{edge: "squad"}
}

// PlayerEventsOrErr returns the PlayerEvents value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) PlayerEventsOrErr() ([]*FixtureEvents, error) {
	if e.loadedTypes[3] {
		return e.PlayerEvents, nil
	}
	return nil, &NotLoadedError{edge: "playerEvents"}
}

// MatchPlayerOrErr returns the MatchPlayer value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) MatchPlayerOrErr() ([]*MatchPlayer, error) {
	if e.loadedTypes[4] {
		return e.MatchPlayer, nil
	}
	return nil, &NotLoadedError{edge: "matchPlayer"}
}

// AssistEventsOrErr returns the AssistEvents value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) AssistEventsOrErr() ([]*FixtureEvents, error) {
	if e.loadedTypes[5] {
		return e.AssistEvents, nil
	}
	return nil, &NotLoadedError{edge: "assistEvents"}
}

// PsgamesOrErr returns the Psgames value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) PsgamesOrErr() ([]*PSGames, error) {
	if e.loadedTypes[6] {
		return e.Psgames, nil
	}
	return nil, &NotLoadedError{edge: "psgames"}
}

// PsgoalsOrErr returns the Psgoals value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) PsgoalsOrErr() ([]*PSGoals, error) {
	if e.loadedTypes[7] {
		return e.Psgoals, nil
	}
	return nil, &NotLoadedError{edge: "psgoals"}
}

// PsdefenseOrErr returns the Psdefense value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) PsdefenseOrErr() ([]*PSDefense, error) {
	if e.loadedTypes[8] {
		return e.Psdefense, nil
	}
	return nil, &NotLoadedError{edge: "psdefense"}
}

// PsoffenseOrErr returns the Psoffense value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) PsoffenseOrErr() ([]*PSOffense, error) {
	if e.loadedTypes[9] {
		return e.Psoffense, nil
	}
	return nil, &NotLoadedError{edge: "psoffense"}
}

// PspenaltyOrErr returns the Pspenalty value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) PspenaltyOrErr() ([]*PSPenalty, error) {
	if e.loadedTypes[10] {
		return e.Pspenalty, nil
	}
	return nil, &NotLoadedError{edge: "pspenalty"}
}

// SeasonOrErr returns the Season value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerEdges) SeasonOrErr() (*Season, error) {
	if e.loadedTypes[11] {
		if e.Season == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: season.Label}
		}
		return e.Season, nil
	}
	return nil, &NotLoadedError{edge: "season"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[12] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// ClubOrErr returns the Club value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerEdges) ClubOrErr() (*Club, error) {
	if e.loadedTypes[13] {
		if e.Club == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: club.Label}
		}
		return e.Club, nil
	}
	return nil, &NotLoadedError{edge: "club"}
}

// LeagueOrErr returns the League value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerEdges) LeagueOrErr() (*League, error) {
	if e.loadedTypes[14] {
		if e.League == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: league.Label}
		}
		return e.League, nil
	}
	return nil, &NotLoadedError{edge: "league"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Player) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case player.FieldInjured:
			values[i] = new(sql.NullBool)
		case player.FieldID, player.FieldApiFootballId, player.FieldAge:
			values[i] = new(sql.NullInt64)
		case player.FieldSlug, player.FieldName, player.FieldFirstname, player.FieldLastname, player.FieldHeight, player.FieldWeight, player.FieldPhoto, player.FieldForm:
			values[i] = new(sql.NullString)
		case player.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case player.ForeignKeys[0]: // birth_player
			values[i] = new(sql.NullInt64)
		case player.ForeignKeys[1]: // club_player
			values[i] = new(sql.NullInt64)
		case player.ForeignKeys[2]: // country_players
			values[i] = new(sql.NullInt64)
		case player.ForeignKeys[3]: // league_player
			values[i] = new(sql.NullInt64)
		case player.ForeignKeys[4]: // season_player
			values[i] = new(sql.NullInt64)
		case player.ForeignKeys[5]: // team_players
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Player fields.
func (pl *Player) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case player.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case player.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				pl.Slug = value.String
			}
		case player.FieldApiFootballId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ApiFootballId", values[i])
			} else if value.Valid {
				pl.ApiFootballId = int(value.Int64)
			}
		case player.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case player.FieldFirstname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstname", values[i])
			} else if value.Valid {
				pl.Firstname = value.String
			}
		case player.FieldLastname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastname", values[i])
			} else if value.Valid {
				pl.Lastname = value.String
			}
		case player.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				pl.Age = int(value.Int64)
			}
		case player.FieldHeight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				pl.Height = value.String
			}
		case player.FieldWeight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				pl.Weight = value.String
			}
		case player.FieldInjured:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field injured", values[i])
			} else if value.Valid {
				pl.Injured = value.Bool
			}
		case player.FieldPhoto:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo", values[i])
			} else if value.Valid {
				pl.Photo = value.String
			}
		case player.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastUpdated", values[i])
			} else if value.Valid {
				pl.LastUpdated = value.Time
			}
		case player.FieldForm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field form", values[i])
			} else if value.Valid {
				pl.Form = value.String
			}
		case player.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field birth_player", value)
			} else if value.Valid {
				pl.birth_player = new(int)
				*pl.birth_player = int(value.Int64)
			}
		case player.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field club_player", value)
			} else if value.Valid {
				pl.club_player = new(int)
				*pl.club_player = int(value.Int64)
			}
		case player.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field country_players", value)
			} else if value.Valid {
				pl.country_players = new(int)
				*pl.country_players = int(value.Int64)
			}
		case player.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field league_player", value)
			} else if value.Valid {
				pl.league_player = new(int)
				*pl.league_player = int(value.Int64)
			}
		case player.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field season_player", value)
			} else if value.Valid {
				pl.season_player = new(int)
				*pl.season_player = int(value.Int64)
			}
		case player.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_players", value)
			} else if value.Valid {
				pl.team_players = new(int)
				*pl.team_players = int(value.Int64)
			}
		default:
			pl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Player.
// This includes values selected through modifiers, order, etc.
func (pl *Player) Value(name string) (ent.Value, error) {
	return pl.selectValues.Get(name)
}

// QueryBirth queries the "birth" edge of the Player entity.
func (pl *Player) QueryBirth() *BirthQuery {
	return NewPlayerClient(pl.config).QueryBirth(pl)
}

// QueryNationality queries the "nationality" edge of the Player entity.
func (pl *Player) QueryNationality() *CountryQuery {
	return NewPlayerClient(pl.config).QueryNationality(pl)
}

// QuerySquad queries the "squad" edge of the Player entity.
func (pl *Player) QuerySquad() *SquadQuery {
	return NewPlayerClient(pl.config).QuerySquad(pl)
}

// QueryPlayerEvents queries the "playerEvents" edge of the Player entity.
func (pl *Player) QueryPlayerEvents() *FixtureEventsQuery {
	return NewPlayerClient(pl.config).QueryPlayerEvents(pl)
}

// QueryMatchPlayer queries the "matchPlayer" edge of the Player entity.
func (pl *Player) QueryMatchPlayer() *MatchPlayerQuery {
	return NewPlayerClient(pl.config).QueryMatchPlayer(pl)
}

// QueryAssistEvents queries the "assistEvents" edge of the Player entity.
func (pl *Player) QueryAssistEvents() *FixtureEventsQuery {
	return NewPlayerClient(pl.config).QueryAssistEvents(pl)
}

// QueryPsgames queries the "psgames" edge of the Player entity.
func (pl *Player) QueryPsgames() *PSGamesQuery {
	return NewPlayerClient(pl.config).QueryPsgames(pl)
}

// QueryPsgoals queries the "psgoals" edge of the Player entity.
func (pl *Player) QueryPsgoals() *PSGoalsQuery {
	return NewPlayerClient(pl.config).QueryPsgoals(pl)
}

// QueryPsdefense queries the "psdefense" edge of the Player entity.
func (pl *Player) QueryPsdefense() *PSDefenseQuery {
	return NewPlayerClient(pl.config).QueryPsdefense(pl)
}

// QueryPsoffense queries the "psoffense" edge of the Player entity.
func (pl *Player) QueryPsoffense() *PSOffenseQuery {
	return NewPlayerClient(pl.config).QueryPsoffense(pl)
}

// QueryPspenalty queries the "pspenalty" edge of the Player entity.
func (pl *Player) QueryPspenalty() *PSPenaltyQuery {
	return NewPlayerClient(pl.config).QueryPspenalty(pl)
}

// QuerySeason queries the "season" edge of the Player entity.
func (pl *Player) QuerySeason() *SeasonQuery {
	return NewPlayerClient(pl.config).QuerySeason(pl)
}

// QueryTeam queries the "team" edge of the Player entity.
func (pl *Player) QueryTeam() *TeamQuery {
	return NewPlayerClient(pl.config).QueryTeam(pl)
}

// QueryClub queries the "club" edge of the Player entity.
func (pl *Player) QueryClub() *ClubQuery {
	return NewPlayerClient(pl.config).QueryClub(pl)
}

// QueryLeague queries the "league" edge of the Player entity.
func (pl *Player) QueryLeague() *LeagueQuery {
	return NewPlayerClient(pl.config).QueryLeague(pl)
}

// Update returns a builder for updating this Player.
// Note that you need to call Player.Unwrap() before calling this method if this Player
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Player) Update() *PlayerUpdateOne {
	return NewPlayerClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Player entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Player) Unwrap() *Player {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Player is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Player) String() string {
	var builder strings.Builder
	builder.WriteString("Player(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("slug=")
	builder.WriteString(pl.Slug)
	builder.WriteString(", ")
	builder.WriteString("ApiFootballId=")
	builder.WriteString(fmt.Sprintf("%v", pl.ApiFootballId))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("firstname=")
	builder.WriteString(pl.Firstname)
	builder.WriteString(", ")
	builder.WriteString("lastname=")
	builder.WriteString(pl.Lastname)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", pl.Age))
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(pl.Height)
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(pl.Weight)
	builder.WriteString(", ")
	builder.WriteString("injured=")
	builder.WriteString(fmt.Sprintf("%v", pl.Injured))
	builder.WriteString(", ")
	builder.WriteString("photo=")
	builder.WriteString(pl.Photo)
	builder.WriteString(", ")
	builder.WriteString("lastUpdated=")
	builder.WriteString(pl.LastUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("form=")
	builder.WriteString(pl.Form)
	builder.WriteByte(')')
	return builder.String()
}

// Players is a parsable slice of Player.
type Players []*Player
