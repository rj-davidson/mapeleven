// Code generated by ent, DO NOT EDIT.

package player

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the player type in the database.
	Label = "player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldApiFootballId holds the string denoting the apifootballid field in the database.
	FieldApiFootballId = "api_football_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldInjured holds the string denoting the injured field in the database.
	FieldInjured = "injured"
	// FieldPhoto holds the string denoting the photo field in the database.
	FieldPhoto = "photo"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// FieldForm holds the string denoting the form field in the database.
	FieldForm = "form"
	// EdgeBirth holds the string denoting the birth edge name in mutations.
	EdgeBirth = "birth"
	// EdgeNationality holds the string denoting the nationality edge name in mutations.
	EdgeNationality = "nationality"
	// EdgeSquad holds the string denoting the squad edge name in mutations.
	EdgeSquad = "squad"
	// EdgePlayerEvents holds the string denoting the playerevents edge name in mutations.
	EdgePlayerEvents = "playerEvents"
	// EdgeMatchPlayer holds the string denoting the matchplayer edge name in mutations.
	EdgeMatchPlayer = "matchPlayer"
	// EdgeAssistEvents holds the string denoting the assistevents edge name in mutations.
	EdgeAssistEvents = "assistEvents"
	// EdgePsgames holds the string denoting the psgames edge name in mutations.
	EdgePsgames = "psgames"
	// EdgePsgoals holds the string denoting the psgoals edge name in mutations.
	EdgePsgoals = "psgoals"
	// EdgePsdefense holds the string denoting the psdefense edge name in mutations.
	EdgePsdefense = "psdefense"
	// EdgePsoffense holds the string denoting the psoffense edge name in mutations.
	EdgePsoffense = "psoffense"
	// EdgePspenalty holds the string denoting the pspenalty edge name in mutations.
	EdgePspenalty = "pspenalty"
	// Table holds the table name of the player in the database.
	Table = "players"
	// BirthTable is the table that holds the birth relation/edge.
	BirthTable = "players"
	// BirthInverseTable is the table name for the Birth entity.
	// It exists in this package in order to avoid circular dependency with the "birth" package.
	BirthInverseTable = "births"
	// BirthColumn is the table column denoting the birth relation/edge.
	BirthColumn = "birth_player"
	// NationalityTable is the table that holds the nationality relation/edge.
	NationalityTable = "players"
	// NationalityInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	NationalityInverseTable = "countries"
	// NationalityColumn is the table column denoting the nationality relation/edge.
	NationalityColumn = "country_players"
	// SquadTable is the table that holds the squad relation/edge.
	SquadTable = "squads"
	// SquadInverseTable is the table name for the Squad entity.
	// It exists in this package in order to avoid circular dependency with the "squad" package.
	SquadInverseTable = "squads"
	// SquadColumn is the table column denoting the squad relation/edge.
	SquadColumn = "player_squad"
	// PlayerEventsTable is the table that holds the playerEvents relation/edge.
	PlayerEventsTable = "fixture_events"
	// PlayerEventsInverseTable is the table name for the FixtureEvents entity.
	// It exists in this package in order to avoid circular dependency with the "fixtureevents" package.
	PlayerEventsInverseTable = "fixture_events"
	// PlayerEventsColumn is the table column denoting the playerEvents relation/edge.
	PlayerEventsColumn = "player_player_events"
	// MatchPlayerTable is the table that holds the matchPlayer relation/edge.
	MatchPlayerTable = "match_players"
	// MatchPlayerInverseTable is the table name for the MatchPlayer entity.
	// It exists in this package in order to avoid circular dependency with the "matchplayer" package.
	MatchPlayerInverseTable = "match_players"
	// MatchPlayerColumn is the table column denoting the matchPlayer relation/edge.
	MatchPlayerColumn = "player_match_player"
	// AssistEventsTable is the table that holds the assistEvents relation/edge.
	AssistEventsTable = "fixture_events"
	// AssistEventsInverseTable is the table name for the FixtureEvents entity.
	// It exists in this package in order to avoid circular dependency with the "fixtureevents" package.
	AssistEventsInverseTable = "fixture_events"
	// AssistEventsColumn is the table column denoting the assistEvents relation/edge.
	AssistEventsColumn = "player_assist_events"
	// PsgamesTable is the table that holds the psgames relation/edge.
	PsgamesTable = "ps_games"
	// PsgamesInverseTable is the table name for the PSGames entity.
	// It exists in this package in order to avoid circular dependency with the "psgames" package.
	PsgamesInverseTable = "ps_games"
	// PsgamesColumn is the table column denoting the psgames relation/edge.
	PsgamesColumn = "player_psgames"
	// PsgoalsTable is the table that holds the psgoals relation/edge.
	PsgoalsTable = "ps_goals"
	// PsgoalsInverseTable is the table name for the PSGoals entity.
	// It exists in this package in order to avoid circular dependency with the "psgoals" package.
	PsgoalsInverseTable = "ps_goals"
	// PsgoalsColumn is the table column denoting the psgoals relation/edge.
	PsgoalsColumn = "player_psgoals"
	// PsdefenseTable is the table that holds the psdefense relation/edge.
	PsdefenseTable = "ps_defenses"
	// PsdefenseInverseTable is the table name for the PSDefense entity.
	// It exists in this package in order to avoid circular dependency with the "psdefense" package.
	PsdefenseInverseTable = "ps_defenses"
	// PsdefenseColumn is the table column denoting the psdefense relation/edge.
	PsdefenseColumn = "player_psdefense"
	// PsoffenseTable is the table that holds the psoffense relation/edge.
	PsoffenseTable = "ps_offenses"
	// PsoffenseInverseTable is the table name for the PSOffense entity.
	// It exists in this package in order to avoid circular dependency with the "psoffense" package.
	PsoffenseInverseTable = "ps_offenses"
	// PsoffenseColumn is the table column denoting the psoffense relation/edge.
	PsoffenseColumn = "player_psoffense"
	// PspenaltyTable is the table that holds the pspenalty relation/edge.
	PspenaltyTable = "ps_penalties"
	// PspenaltyInverseTable is the table name for the PSPenalty entity.
	// It exists in this package in order to avoid circular dependency with the "pspenalty" package.
	PspenaltyInverseTable = "ps_penalties"
	// PspenaltyColumn is the table column denoting the pspenalty relation/edge.
	PspenaltyColumn = "player_pspenalty"
)

// Columns holds all SQL columns for player fields.
var Columns = []string{
	FieldID,
	FieldSlug,
	FieldApiFootballId,
	FieldName,
	FieldFirstname,
	FieldLastname,
	FieldAge,
	FieldHeight,
	FieldWeight,
	FieldInjured,
	FieldPhoto,
	FieldLastUpdated,
	FieldForm,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "players"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"birth_player",
	"country_players",
	"team_players",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultLastUpdated holds the default value on creation for the "lastUpdated" field.
	DefaultLastUpdated func() time.Time
	// UpdateDefaultLastUpdated holds the default value on update for the "lastUpdated" field.
	UpdateDefaultLastUpdated func() time.Time
)

// OrderOption defines the ordering options for the Player queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByApiFootballId orders the results by the ApiFootballId field.
func ByApiFootballId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApiFootballId, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFirstname orders the results by the firstname field.
func ByFirstname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstname, opts...).ToFunc()
}

// ByLastname orders the results by the lastname field.
func ByLastname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastname, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByInjured orders the results by the injured field.
func ByInjured(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInjured, opts...).ToFunc()
}

// ByPhoto orders the results by the photo field.
func ByPhoto(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoto, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByForm orders the results by the form field.
func ByForm(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForm, opts...).ToFunc()
}

// ByBirthField orders the results by birth field.
func ByBirthField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBirthStep(), sql.OrderByField(field, opts...))
	}
}

// ByNationalityField orders the results by nationality field.
func ByNationalityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNationalityStep(), sql.OrderByField(field, opts...))
	}
}

// BySquadCount orders the results by squad count.
func BySquadCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSquadStep(), opts...)
	}
}

// BySquad orders the results by squad terms.
func BySquad(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSquadStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlayerEventsCount orders the results by playerEvents count.
func ByPlayerEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlayerEventsStep(), opts...)
	}
}

// ByPlayerEvents orders the results by playerEvents terms.
func ByPlayerEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayerEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMatchPlayerCount orders the results by matchPlayer count.
func ByMatchPlayerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMatchPlayerStep(), opts...)
	}
}

// ByMatchPlayer orders the results by matchPlayer terms.
func ByMatchPlayer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMatchPlayerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAssistEventsCount orders the results by assistEvents count.
func ByAssistEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssistEventsStep(), opts...)
	}
}

// ByAssistEvents orders the results by assistEvents terms.
func ByAssistEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssistEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPsgamesCount orders the results by psgames count.
func ByPsgamesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPsgamesStep(), opts...)
	}
}

// ByPsgames orders the results by psgames terms.
func ByPsgames(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPsgamesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPsgoalsCount orders the results by psgoals count.
func ByPsgoalsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPsgoalsStep(), opts...)
	}
}

// ByPsgoals orders the results by psgoals terms.
func ByPsgoals(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPsgoalsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPsdefenseCount orders the results by psdefense count.
func ByPsdefenseCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPsdefenseStep(), opts...)
	}
}

// ByPsdefense orders the results by psdefense terms.
func ByPsdefense(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPsdefenseStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPsoffenseCount orders the results by psoffense count.
func ByPsoffenseCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPsoffenseStep(), opts...)
	}
}

// ByPsoffense orders the results by psoffense terms.
func ByPsoffense(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPsoffenseStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPspenaltyCount orders the results by pspenalty count.
func ByPspenaltyCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPspenaltyStep(), opts...)
	}
}

// ByPspenalty orders the results by pspenalty terms.
func ByPspenalty(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPspenaltyStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBirthStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BirthInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BirthTable, BirthColumn),
	)
}
func newNationalityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NationalityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, NationalityTable, NationalityColumn),
	)
}
func newSquadStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SquadInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SquadTable, SquadColumn),
	)
}
func newPlayerEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayerEventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PlayerEventsTable, PlayerEventsColumn),
	)
}
func newMatchPlayerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MatchPlayerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MatchPlayerTable, MatchPlayerColumn),
	)
}
func newAssistEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssistEventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AssistEventsTable, AssistEventsColumn),
	)
}
func newPsgamesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PsgamesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PsgamesTable, PsgamesColumn),
	)
}
func newPsgoalsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PsgoalsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PsgoalsTable, PsgoalsColumn),
	)
}
func newPsdefenseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PsdefenseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PsdefenseTable, PsdefenseColumn),
	)
}
func newPsoffenseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PsoffenseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PsoffenseTable, PsoffenseColumn),
	)
}
func newPspenaltyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PspenaltyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PspenaltyTable, PspenaltyColumn),
	)
}
