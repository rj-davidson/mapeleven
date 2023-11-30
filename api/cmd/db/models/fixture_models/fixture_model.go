package fixture_models

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixture"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"context"
	"log"
	"time"
)

type FixtureDetails struct {
	Teams   TeamDetails  `json:"teams"`
	Events  []Event      `json:"events"`
	Lineups []LineupInfo `json:"lineups"`
}

type TeamDetails struct {
	Home TeamInfo `json:"home"`
	Away TeamInfo `json:"away"`
}

type TeamInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner bool   `json:"winner"`
}

type Event struct {
	Time     TimeDetail `json:"time"`
	Team     TeamInfo   `json:"team"`
	Player   Player     `json:"player"`
	Assist   *Player    `json:"assist"`
	Type     string     `json:"type"`
	Detail   string     `json:"detail"`
	Comments *string    `json:"comments"`
}

type TimeDetail struct {
	Elapsed int  `json:"elapsed"`
	Extra   *int `json:"extra"`
}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LineupInfo struct {
	Team        TeamInfo       `json:"team"`
	Coach       Coach          `json:"coach"`
	Formation   string         `json:"formation"`
	StartXI     []PlayerDetail `json:"startXI"`
	Substitutes []PlayerDetail `json:"substitutes"`
}

type Coach struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type PlayerDetail struct {
	Player PlayerDetailInfo `json:"player"`
}

type PlayerDetailInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
	Pos    string `json:"pos"`
	Grid   string `json:"grid"`
}

// CreateFixtureInput holds the input data needed to create a new fixture record.
type CreateFixtureInput struct {
	ApiFootballId int
	Slug          string
	Referee       *string
	Timezone      *string
	Date          time.Time
	Elapsed       *int
	Round         *int
	Status        string
	HomeTeamScore *int
	AwayTeamScore *int
	Season        *ent.Season
	HomeTeamID    int
	AwayTeamID    int
}

// UpdateFixtureInput holds the input data needed to update an existing fixture record.
type UpdateFixtureInput struct {
	ApiFootballId int
	Slug          string
	Referee       *string
	Timezone      *string
	Date          time.Time
	Elapsed       *int
	Round         *int
	Status        *string
	HomeTeamScore *int
	AwayTeamScore *int
	HomeTeamID    *int
	AwayTeamID    *int
}

type FixtureModel struct {
	client *ent.Client
}

func NewFixtureModel(client *ent.Client) *FixtureModel {
	return &FixtureModel{client: client}
}

func (fm *FixtureModel) getTeam(ctx context.Context, seasonID, apiFootballClubId int) (*ent.Team, error) {
	return fm.client.Team.
		Query().
		Where(
			team.HasSeasonWith(
				season.IDEQ(seasonID),
			),
			team.HasClubWith(
				club.ApiFootballIdEQ(apiFootballClubId),
			),
		).
		Only(ctx)
}

func (fm *FixtureModel) CreateBaseFixture(ctx context.Context, input CreateFixtureInput) (*ent.Fixture, error) {
	awayTeam, err := fm.getTeam(ctx, input.Season.ID, input.AwayTeamID)
	if err != nil {
		return nil, err
	}
	homeTeam, err := fm.getTeam(ctx, input.Season.ID, input.HomeTeamID)
	if err != nil {
		return nil, err
	}
	return fm.client.Fixture.
		Create().
		SetSlug(input.Slug).
		SetApiFootballId(input.ApiFootballId).
		SetNillableReferee(input.Referee).
		SetNillableTimezone(input.Timezone).
		SetDate(input.Date).
		SetNillableElapsed(input.Elapsed).
		SetStatus(input.Status).
		SetNillableRound(input.Round).
		SetNillableHomeTeamScore(input.HomeTeamScore).
		SetNillableAwayTeamScore(input.AwayTeamScore).
		SetAwayTeam(awayTeam).
		SetHomeTeam(homeTeam).
		SetSeason(input.Season).
		Save(ctx)
}

func (fm *FixtureModel) UpdateBaseFixture(ctx context.Context, input UpdateFixtureInput) (*ent.Fixture, error) {
	f, _ := fm.client.Fixture.Query().
		Where(fixture.ApiFootballIdEQ(input.ApiFootballId)).
		Only(ctx)
	updater := fm.client.Fixture.UpdateOne(f)
	if input.Round != nil {
		updater.SetRound(*input.Round)
	}
	if input.Status != nil {
		updater.SetStatus(*input.Status)
	}
	if input.HomeTeamScore != nil {
		updater.SetHomeTeamScore(*input.HomeTeamScore)
	}
	if input.AwayTeamScore != nil {
		updater.SetAwayTeamScore(*input.AwayTeamScore)
	}
	return updater.Save(ctx)
}

func (fm *FixtureModel) GetFixtureByApiFootballId(ctx context.Context, apiFootballId int) (*ent.Fixture, error) {
	return fm.client.Fixture.
		Query().
		Where(fixture.ApiFootballIdEQ(apiFootballId)).
		WithHomeTeam().
		WithAwayTeam().
		Only(ctx)
}

func (fm *FixtureModel) FixtureExistsByApiFootballId(ctx context.Context, apiFootballId int) bool {
	f, _ := fm.client.Fixture.
		Query().
		Where(fixture.ApiFootballIdEQ(apiFootballId)).
		Only(ctx)
	return f != nil
}

func (fm *FixtureModel) ListFixtures(ctx context.Context) ([]*ent.Fixture, error) {
	return fm.client.Fixture.
		Query().
		WithSeason(func(sq *ent.SeasonQuery) {
			sq.WithLeague()
		}).
		All(ctx)
}

func (fm *FixtureModel) UpsertFixtureData(apiFootballId int, data FixtureDetails, ctx context.Context) error {
	f, err := fm.GetFixtureByApiFootballId(ctx, apiFootballId)
	homeTeam := f.Edges.HomeTeam
	awayTeam := f.Edges.AwayTeam

	if len(data.Lineups) > 0 {
		_, err = UpsertLineup(f, homeTeam, data.Lineups[0], fm.client, ctx)
		if err != nil {
			log.Printf("Error upserting lineup: %v", err)
		}

		_, err = UpsertLineup(f, awayTeam, data.Lineups[1], fm.client, ctx)
		if err != nil {
			log.Printf("Error upserting lineup: %v", err)
		}
	}

	err = RepopulateEvents(f, fm.client, data.Events, ctx)
	if err != nil {
		log.Printf("Error repopulating events: %v", err)
	}

	return nil
}

// ListTodaysFixtures List fixtures that will be played today or tomorrow and do not have the status of FT, AET, PEN, or CANC
func (fm *FixtureModel) ListTodaysFixtures(ctx context.Context) ([]*ent.Fixture, error) {
	today := time.Now().UTC()
	tomorrow := today.AddDate(0, 0, 1)
	return fm.client.Fixture.
		Query().
		Where(
			fixture.DateGTE(today),
			fixture.DateLT(tomorrow),
			fixture.StatusNEQ("FT"),
			fixture.StatusNEQ("AET"),
			fixture.StatusNEQ("PEN"),
			fixture.StatusNEQ("CANC"),
		).
		WithSeason(func(sq *ent.SeasonQuery) {
			sq.WithLeague()
		}).
		All(ctx)
}

// ListLiveFixtures lists fixtures that are scheduled for today or tomorrow and have the status of 1H, HT, 2H, ET, BT, or P
func (fm *FixtureModel) ListLiveFixtures(ctx context.Context) ([]*ent.Fixture, error) {
	today := time.Now().UTC()
	tomorrow := today.AddDate(0, 0, 1)
	return fm.client.Fixture.
		Query().
		Where(
			fixture.DateGTE(today),
			fixture.DateLT(tomorrow),
			fixture.StatusEQ("1H"),
			fixture.StatusEQ("HT"),
			fixture.StatusEQ("2H"),
			fixture.StatusEQ("ET"),
			fixture.StatusEQ("BT"),
			fixture.StatusEQ("P"),
		).
		WithSeason(func(sq *ent.SeasonQuery) {
			sq.WithLeague()
		}).
		All(ctx)
}
