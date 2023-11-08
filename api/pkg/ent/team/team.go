// Code generated by ent, DO NOT EDIT.

package team

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldForm holds the string denoting the form field in the database.
	FieldForm = "form"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeSeason holds the string denoting the season edge name in mutations.
	EdgeSeason = "season"
	// EdgeClub holds the string denoting the club edge name in mutations.
	EdgeClub = "club"
	// EdgePlayerStats holds the string denoting the playerstats edge name in mutations.
	EdgePlayerStats = "playerStats"
	// EdgeStandings holds the string denoting the standings edge name in mutations.
	EdgeStandings = "standings"
	// EdgeHomeFixtures holds the string denoting the homefixtures edge name in mutations.
	EdgeHomeFixtures = "homeFixtures"
	// EdgeAwayFixtures holds the string denoting the awayfixtures edge name in mutations.
	EdgeAwayFixtures = "awayFixtures"
	// EdgeTeamFixtureEvents holds the string denoting the teamfixtureevents edge name in mutations.
	EdgeTeamFixtureEvents = "teamFixtureEvents"
	// EdgeFixtureLineups holds the string denoting the fixturelineups edge name in mutations.
	EdgeFixtureLineups = "fixtureLineups"
	// EdgeSquad holds the string denoting the squad edge name in mutations.
	EdgeSquad = "squad"
	// EdgeBiggestStats holds the string denoting the biggest_stats edge name in mutations.
	EdgeBiggestStats = "biggest_stats"
	// EdgeCardsStats holds the string denoting the cards_stats edge name in mutations.
	EdgeCardsStats = "cards_stats"
	// EdgeCleanSheetStats holds the string denoting the clean_sheet_stats edge name in mutations.
	EdgeCleanSheetStats = "clean_sheet_stats"
	// EdgeFailedToScoreStats holds the string denoting the failed_to_score_stats edge name in mutations.
	EdgeFailedToScoreStats = "failed_to_score_stats"
	// EdgeFixturesStats holds the string denoting the fixtures_stats edge name in mutations.
	EdgeFixturesStats = "fixtures_stats"
	// EdgeGoalsStats holds the string denoting the goals_stats edge name in mutations.
	EdgeGoalsStats = "goals_stats"
	// EdgeLineups holds the string denoting the lineups edge name in mutations.
	EdgeLineups = "lineups"
	// EdgePenaltyStats holds the string denoting the penalty_stats edge name in mutations.
	EdgePenaltyStats = "penalty_stats"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// SeasonTable is the table that holds the season relation/edge.
	SeasonTable = "teams"
	// SeasonInverseTable is the table name for the Season entity.
	// It exists in this package in order to avoid circular dependency with the "season" package.
	SeasonInverseTable = "seasons"
	// SeasonColumn is the table column denoting the season relation/edge.
	SeasonColumn = "season_teams"
	// ClubTable is the table that holds the club relation/edge.
	ClubTable = "teams"
	// ClubInverseTable is the table name for the Club entity.
	// It exists in this package in order to avoid circular dependency with the "club" package.
	ClubInverseTable = "clubs"
	// ClubColumn is the table column denoting the club relation/edge.
	ClubColumn = "club_team"
	// PlayerStatsTable is the table that holds the playerStats relation/edge.
	PlayerStatsTable = "teams"
	// PlayerStatsInverseTable is the table name for the PlayerStats entity.
	// It exists in this package in order to avoid circular dependency with the "playerstats" package.
	PlayerStatsInverseTable = "player_stats"
	// PlayerStatsColumn is the table column denoting the playerStats relation/edge.
	PlayerStatsColumn = "player_stats_team"
	// StandingsTable is the table that holds the standings relation/edge.
	StandingsTable = "standings"
	// StandingsInverseTable is the table name for the Standings entity.
	// It exists in this package in order to avoid circular dependency with the "standings" package.
	StandingsInverseTable = "standings"
	// StandingsColumn is the table column denoting the standings relation/edge.
	StandingsColumn = "team_standings"
	// HomeFixturesTable is the table that holds the homeFixtures relation/edge.
	HomeFixturesTable = "fixtures"
	// HomeFixturesInverseTable is the table name for the Fixture entity.
	// It exists in this package in order to avoid circular dependency with the "fixture" package.
	HomeFixturesInverseTable = "fixtures"
	// HomeFixturesColumn is the table column denoting the homeFixtures relation/edge.
	HomeFixturesColumn = "team_home_fixtures"
	// AwayFixturesTable is the table that holds the awayFixtures relation/edge.
	AwayFixturesTable = "fixtures"
	// AwayFixturesInverseTable is the table name for the Fixture entity.
	// It exists in this package in order to avoid circular dependency with the "fixture" package.
	AwayFixturesInverseTable = "fixtures"
	// AwayFixturesColumn is the table column denoting the awayFixtures relation/edge.
	AwayFixturesColumn = "team_away_fixtures"
	// TeamFixtureEventsTable is the table that holds the teamFixtureEvents relation/edge.
	TeamFixtureEventsTable = "fixture_events"
	// TeamFixtureEventsInverseTable is the table name for the FixtureEvents entity.
	// It exists in this package in order to avoid circular dependency with the "fixtureevents" package.
	TeamFixtureEventsInverseTable = "fixture_events"
	// TeamFixtureEventsColumn is the table column denoting the teamFixtureEvents relation/edge.
	TeamFixtureEventsColumn = "team_team_fixture_events"
	// FixtureLineupsTable is the table that holds the fixtureLineups relation/edge.
	FixtureLineupsTable = "fixture_lineups"
	// FixtureLineupsInverseTable is the table name for the FixtureLineups entity.
	// It exists in this package in order to avoid circular dependency with the "fixturelineups" package.
	FixtureLineupsInverseTable = "fixture_lineups"
	// FixtureLineupsColumn is the table column denoting the fixtureLineups relation/edge.
	FixtureLineupsColumn = "team_fixture_lineups"
	// SquadTable is the table that holds the squad relation/edge.
	SquadTable = "squads"
	// SquadInverseTable is the table name for the Squad entity.
	// It exists in this package in order to avoid circular dependency with the "squad" package.
	SquadInverseTable = "squads"
	// SquadColumn is the table column denoting the squad relation/edge.
	SquadColumn = "team_squad"
	// BiggestStatsTable is the table that holds the biggest_stats relation/edge.
	BiggestStatsTable = "ts_biggests"
	// BiggestStatsInverseTable is the table name for the TSBiggest entity.
	// It exists in this package in order to avoid circular dependency with the "tsbiggest" package.
	BiggestStatsInverseTable = "ts_biggests"
	// BiggestStatsColumn is the table column denoting the biggest_stats relation/edge.
	BiggestStatsColumn = "team_biggest_stats"
	// CardsStatsTable is the table that holds the cards_stats relation/edge.
	CardsStatsTable = "ts_cards"
	// CardsStatsInverseTable is the table name for the TSCards entity.
	// It exists in this package in order to avoid circular dependency with the "tscards" package.
	CardsStatsInverseTable = "ts_cards"
	// CardsStatsColumn is the table column denoting the cards_stats relation/edge.
	CardsStatsColumn = "team_cards_stats"
	// CleanSheetStatsTable is the table that holds the clean_sheet_stats relation/edge.
	CleanSheetStatsTable = "ts_clean_sheets"
	// CleanSheetStatsInverseTable is the table name for the TSCleanSheet entity.
	// It exists in this package in order to avoid circular dependency with the "tscleansheet" package.
	CleanSheetStatsInverseTable = "ts_clean_sheets"
	// CleanSheetStatsColumn is the table column denoting the clean_sheet_stats relation/edge.
	CleanSheetStatsColumn = "team_clean_sheet_stats"
	// FailedToScoreStatsTable is the table that holds the failed_to_score_stats relation/edge.
	FailedToScoreStatsTable = "ts_failed_to_scores"
	// FailedToScoreStatsInverseTable is the table name for the TSFailedToScore entity.
	// It exists in this package in order to avoid circular dependency with the "tsfailedtoscore" package.
	FailedToScoreStatsInverseTable = "ts_failed_to_scores"
	// FailedToScoreStatsColumn is the table column denoting the failed_to_score_stats relation/edge.
	FailedToScoreStatsColumn = "team_failed_to_score_stats"
	// FixturesStatsTable is the table that holds the fixtures_stats relation/edge.
	FixturesStatsTable = "ts_fixtures"
	// FixturesStatsInverseTable is the table name for the TSFixtures entity.
	// It exists in this package in order to avoid circular dependency with the "tsfixtures" package.
	FixturesStatsInverseTable = "ts_fixtures"
	// FixturesStatsColumn is the table column denoting the fixtures_stats relation/edge.
	FixturesStatsColumn = "team_fixtures_stats"
	// GoalsStatsTable is the table that holds the goals_stats relation/edge.
	GoalsStatsTable = "ts_goals"
	// GoalsStatsInverseTable is the table name for the TSGoals entity.
	// It exists in this package in order to avoid circular dependency with the "tsgoals" package.
	GoalsStatsInverseTable = "ts_goals"
	// GoalsStatsColumn is the table column denoting the goals_stats relation/edge.
	GoalsStatsColumn = "team_goals_stats"
	// LineupsTable is the table that holds the lineups relation/edge.
	LineupsTable = "ts_lineups"
	// LineupsInverseTable is the table name for the TSLineups entity.
	// It exists in this package in order to avoid circular dependency with the "tslineups" package.
	LineupsInverseTable = "ts_lineups"
	// LineupsColumn is the table column denoting the lineups relation/edge.
	LineupsColumn = "team_id"
	// PenaltyStatsTable is the table that holds the penalty_stats relation/edge.
	PenaltyStatsTable = "ts_penalties"
	// PenaltyStatsInverseTable is the table name for the TSPenalty entity.
	// It exists in this package in order to avoid circular dependency with the "tspenalty" package.
	PenaltyStatsInverseTable = "ts_penalties"
	// PenaltyStatsColumn is the table column denoting the penalty_stats relation/edge.
	PenaltyStatsColumn = "team_penalty_stats"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldForm,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "teams"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"club_team",
	"player_stats_team",
	"season_teams",
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

// OrderOption defines the ordering options for the Team queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByForm orders the results by the form field.
func ByForm(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForm, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// BySeasonField orders the results by season field.
func BySeasonField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSeasonStep(), sql.OrderByField(field, opts...))
	}
}

// ByClubField orders the results by club field.
func ByClubField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClubStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlayerStatsField orders the results by playerStats field.
func ByPlayerStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayerStatsStep(), sql.OrderByField(field, opts...))
	}
}

// ByStandingsCount orders the results by standings count.
func ByStandingsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStandingsStep(), opts...)
	}
}

// ByStandings orders the results by standings terms.
func ByStandings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStandingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByHomeFixturesCount orders the results by homeFixtures count.
func ByHomeFixturesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHomeFixturesStep(), opts...)
	}
}

// ByHomeFixtures orders the results by homeFixtures terms.
func ByHomeFixtures(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHomeFixturesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAwayFixturesCount orders the results by awayFixtures count.
func ByAwayFixturesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAwayFixturesStep(), opts...)
	}
}

// ByAwayFixtures orders the results by awayFixtures terms.
func ByAwayFixtures(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAwayFixturesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTeamFixtureEventsCount orders the results by teamFixtureEvents count.
func ByTeamFixtureEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamFixtureEventsStep(), opts...)
	}
}

// ByTeamFixtureEvents orders the results by teamFixtureEvents terms.
func ByTeamFixtureEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamFixtureEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFixtureLineupsCount orders the results by fixtureLineups count.
func ByFixtureLineupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFixtureLineupsStep(), opts...)
	}
}

// ByFixtureLineups orders the results by fixtureLineups terms.
func ByFixtureLineups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFixtureLineupsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByBiggestStatsField orders the results by biggest_stats field.
func ByBiggestStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBiggestStatsStep(), sql.OrderByField(field, opts...))
	}
}

// ByCardsStatsField orders the results by cards_stats field.
func ByCardsStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCardsStatsStep(), sql.OrderByField(field, opts...))
	}
}

// ByCleanSheetStatsField orders the results by clean_sheet_stats field.
func ByCleanSheetStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCleanSheetStatsStep(), sql.OrderByField(field, opts...))
	}
}

// ByFailedToScoreStatsField orders the results by failed_to_score_stats field.
func ByFailedToScoreStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFailedToScoreStatsStep(), sql.OrderByField(field, opts...))
	}
}

// ByFixturesStatsField orders the results by fixtures_stats field.
func ByFixturesStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFixturesStatsStep(), sql.OrderByField(field, opts...))
	}
}

// ByGoalsStatsField orders the results by goals_stats field.
func ByGoalsStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGoalsStatsStep(), sql.OrderByField(field, opts...))
	}
}

// ByLineupsCount orders the results by lineups count.
func ByLineupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLineupsStep(), opts...)
	}
}

// ByLineups orders the results by lineups terms.
func ByLineups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLineupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPenaltyStatsField orders the results by penalty_stats field.
func ByPenaltyStatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPenaltyStatsStep(), sql.OrderByField(field, opts...))
	}
}
func newSeasonStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SeasonInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SeasonTable, SeasonColumn),
	)
}
func newClubStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClubInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ClubTable, ClubColumn),
	)
}
func newPlayerStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayerStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlayerStatsTable, PlayerStatsColumn),
	)
}
func newStandingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StandingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StandingsTable, StandingsColumn),
	)
}
func newHomeFixturesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HomeFixturesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HomeFixturesTable, HomeFixturesColumn),
	)
}
func newAwayFixturesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AwayFixturesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AwayFixturesTable, AwayFixturesColumn),
	)
}
func newTeamFixtureEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamFixtureEventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TeamFixtureEventsTable, TeamFixtureEventsColumn),
	)
}
func newFixtureLineupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FixtureLineupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FixtureLineupsTable, FixtureLineupsColumn),
	)
}
func newSquadStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SquadInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SquadTable, SquadColumn),
	)
}
func newBiggestStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BiggestStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, BiggestStatsTable, BiggestStatsColumn),
	)
}
func newCardsStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CardsStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CardsStatsTable, CardsStatsColumn),
	)
}
func newCleanSheetStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CleanSheetStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CleanSheetStatsTable, CleanSheetStatsColumn),
	)
}
func newFailedToScoreStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FailedToScoreStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, FailedToScoreStatsTable, FailedToScoreStatsColumn),
	)
}
func newFixturesStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FixturesStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, FixturesStatsTable, FixturesStatsColumn),
	)
}
func newGoalsStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GoalsStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, GoalsStatsTable, GoalsStatsColumn),
	)
}
func newLineupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LineupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LineupsTable, LineupsColumn),
	)
}
func newPenaltyStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PenaltyStatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, PenaltyStatsTable, PenaltyStatsColumn),
	)
}
