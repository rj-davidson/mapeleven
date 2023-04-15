// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/birth"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/player"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Player is the model entity for the Player schema.
type Player struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Firstname holds the value of the "firstname" field.
	Firstname string `json:"firstname,omitempty"`
	// Lastname holds the value of the "lastname" field.
	Lastname string `json:"lastname,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Height holds the value of the "height" field.
	Height float64 `json:"height,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight float64 `json:"weight,omitempty"`
	// Injured holds the value of the "injured" field.
	Injured bool `json:"injured,omitempty"`
	// Photo holds the value of the "photo" field.
	Photo string `json:"photo,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerQuery when eager-loading is set.
	Edges           PlayerEdges `json:"edges"`
	country_players *int
	selectValues    sql.SelectValues
}

// PlayerEdges holds the relations/edges for other nodes in the graph.
type PlayerEdges struct {
	// Birth holds the value of the birth edge.
	Birth *Birth `json:"birth,omitempty"`
	// Nationality holds the value of the nationality edge.
	Nationality []*Country `json:"nationality,omitempty"`
	// Teams holds the value of the teams edge.
	Teams []*Team `json:"teams,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
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
// was not loaded in eager-loading.
func (e PlayerEdges) NationalityOrErr() ([]*Country, error) {
	if e.loadedTypes[1] {
		return e.Nationality, nil
	}
	return nil, &NotLoadedError{edge: "nationality"}
}

// TeamsOrErr returns the Teams value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) TeamsOrErr() ([]*Team, error) {
	if e.loadedTypes[2] {
		return e.Teams, nil
	}
	return nil, &NotLoadedError{edge: "teams"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Player) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case player.FieldInjured:
			values[i] = new(sql.NullBool)
		case player.FieldHeight, player.FieldWeight:
			values[i] = new(sql.NullFloat64)
		case player.FieldID, player.FieldAge:
			values[i] = new(sql.NullInt64)
		case player.FieldName, player.FieldFirstname, player.FieldLastname, player.FieldPhoto:
			values[i] = new(sql.NullString)
		case player.ForeignKeys[0]: // country_players
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
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				pl.Height = value.Float64
			}
		case player.FieldWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				pl.Weight = value.Float64
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
		case player.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field country_players", value)
			} else if value.Valid {
				pl.country_players = new(int)
				*pl.country_players = int(value.Int64)
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

// QueryTeams queries the "teams" edge of the Player entity.
func (pl *Player) QueryTeams() *TeamQuery {
	return NewPlayerClient(pl.config).QueryTeams(pl)
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
	builder.WriteString(fmt.Sprintf("%v", pl.Height))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", pl.Weight))
	builder.WriteString(", ")
	builder.WriteString("injured=")
	builder.WriteString(fmt.Sprintf("%v", pl.Injured))
	builder.WriteString(", ")
	builder.WriteString("photo=")
	builder.WriteString(pl.Photo)
	builder.WriteByte(')')
	return builder.String()
}

// Players is a parsable slice of Player.
type Players []*Player
