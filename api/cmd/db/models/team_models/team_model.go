package team_models

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"context"
)

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
func (m *TeamModel) ListAll(ctx context.Context) ([]*ent.Team, error) {
	return m.client.Team.
		Query().
		WithSeason(
			func(s *ent.SeasonQuery) {
				s.WithLeague()
			}).
		WithSquad().
		WithClub().
		All(ctx)
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
