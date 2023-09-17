package models

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/club"
	"mapeleven/db/ent/fixture"
	"mapeleven/db/ent/season"
	"mapeleven/db/ent/team"
	"time"
)

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

func (fm *FixtureModel) getTeam(ctx context.Context, seasonID, apiFootabllClubId int) (*ent.Team, error) {
	return fm.client.Team.
		Query().
		Where(
			team.HasSeasonWith(
				season.IDEQ(seasonID),
			),
			team.HasClubWith(
				club.ApiFootballIdEQ(apiFootabllClubId),
			),
		).
		Only(ctx)
}

func (fm *FixtureModel) CreateFixture(ctx context.Context, input CreateFixtureInput) (*ent.Fixture, error) {
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

func (fm *FixtureModel) UpdateFixture(ctx context.Context, input UpdateFixtureInput) (*ent.Fixture, error) {
	updater := fm.client.Fixture.UpdateOneID(input.ApiFootballId)
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
		All(ctx)
}
