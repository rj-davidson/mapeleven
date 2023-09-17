package models

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/club"
	"mapeleven/db/ent/league"
	"mapeleven/db/ent/season"
	"mapeleven/db/ent/team"
)

type Team struct {
}

type TeamModel struct {
	client *ent.Client
}

func NewTeamModel(client *ent.Client) *TeamModel {
	return &TeamModel{client: client}
}

// CreateTeam creates a new team.
func (m *TeamModel) CreateTeam(ctx context.Context, year, apiFootballLeagueId, apiFootballClubId int) (*ent.Team, error) {
	// Find the season by its year
	s, err := m.client.Season.
		Query().
		Where(
			season.YearEQ(year),
			season.HasLeagueWith(
				league.FootballApiIdEQ(apiFootballLeagueId),
			),
		).
		WithLeague().
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	// Find the club by its apiFootballId
	c, err := m.client.Club.
		Query().
		Where(club.ApiFootballIdEQ(apiFootballClubId)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return m.client.Team.
		Create().
		SetSeason(s).
		SetClub(c).
		Save(ctx)
}

// UpdateTeam updates an existing team.
func (m *TeamModel) UpdateTeam(ctx context.Context, year, apiFootballLeagueId, apiFootballClubId int) error {
	_, err := m.client.Team.
		Update().
		Where(
			team.HasSeasonWith(
				season.YearEQ(year),
				season.HasLeagueWith(
					league.FootballApiIdEQ(apiFootballLeagueId),
				),
			),
			team.HasClubWith(
				club.ApiFootballIdEQ(apiFootballClubId),
			),
		).
		SetSeasonID(year).
		SetClubID(apiFootballClubId).
		Save(ctx)

	return err
}

// ListAll returns a list of all teams.
func (m *TeamModel) ListAll() ([]*ent.Team, error) {
	return m.client.Team.Query().All(context.Background())
}

// Exists checks if a team exists.
func (m *TeamModel) Exists(ctx context.Context, year, apiFootballLeagueId, apiFootballClubId int) bool {
	count, _ := m.client.Team.
		Query().
		Where(
			team.HasSeasonWith(
				season.YearEQ(year),
				season.HasLeagueWith(
					league.FootballApiIdEQ(apiFootballLeagueId),
				),
			),
			team.HasClubWith(
				club.ApiFootballIdEQ(apiFootballClubId),
			),
		).
		Count(ctx)

	return count > 0
}

// GetTeam retrieves a team by its year, apiFootballLeagueId and apiFootballClubId.
func (m *TeamModel) GetTeam(ctx context.Context, year, apiFootballLeagueId, apiFootballClubId int) (*ent.Team, error) {
	return m.client.Team.
		Query().
		Where(
			team.HasSeasonWith(
				season.YearEQ(year),
				season.HasLeagueWith(
					league.FootballApiIdEQ(apiFootballLeagueId),
				),
			),
			team.HasClubWith(
				club.ApiFootballIdEQ(apiFootballClubId),
			),
		).
		Only(ctx)
}
