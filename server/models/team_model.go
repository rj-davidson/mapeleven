package models

import (
	"context"
	"fmt"
	"mapeleven/db/ent"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/team"
)

// CreateTeamInput holds the required fields to create a new team.
type CreateTeamInput struct {
	ID       int
	Name     string
	Code     string
	Founded  int
	National bool
	Logo     string
	Country  string
}

// UpdateTeamInput holds the fields to update an existing team.
type UpdateTeamInput struct {
	ID       int
	Name     *string
	Code     *string
	Founded  *int
	National *bool
	Logo     *string
	Country  *string
}

// TeamModel defines the methods for managing the team data.
type TeamModel struct {
	client *ent.Client
}

// NewTeamModel creates a new TeamModel.
func NewTeamModel(client *ent.Client) *TeamModel {
	return &TeamModel{client: client}
}

func (m *TeamModel) CreateTeam(ctx context.Context, input CreateTeamInput) (*ent.Team, error) {
	// Get the country by code
	// Find the country by its code
	c, err := m.client.Country.
		Query().
		Where(country.NameEQ(input.Country)).
		Only(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to find country: %w", err)
	}

	return m.client.Team.
		Create().
		SetID(input.ID).
		SetName(input.Name).
		SetCode(input.Code).
		SetFounded(input.Founded).
		SetNational(input.National).
		SetLogo(input.Logo).
		SetCountryID(c.ID).
		Save(ctx)
}

// UpdateTeam updates an existing team.
func (m *TeamModel) UpdateTeam(ctx context.Context, input UpdateTeamInput) (*ent.Team, error) {
	// Find the country by its code
	c, err := m.client.Country.
		Query().
		Where(country.NameEQ(*input.Country)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	updater := m.client.Team.UpdateOneID(input.ID)
	if input.Name != nil {
		updater.SetName(*input.Name)
	}
	if input.Code != nil {
		updater.SetCode(*input.Code)
	}
	if input.Founded != nil {
		updater.SetFounded(*input.Founded)
	}
	if input.National != nil {
		updater.SetNational(*input.National)
	}
	if input.Logo != nil {
		updater.SetLogo(*input.Logo)
	}
	if input.Country != nil {
		updater.SetCountryID(c.ID)
	}

	return updater.Save(ctx)
}

// DeleteTeam deletes a team.
func (m *TeamModel) DeleteTeam(ctx context.Context, teamID int) error {
	return m.client.Team.
		DeleteOneID(teamID).
		Exec(ctx)
}

// GetTeamByID retrieves a team by ID.
func (m *TeamModel) GetTeamByID(ctx context.Context, teamID int) (*ent.Team, error) {
	return m.client.Team.
		Get(ctx, teamID)
}

// ListTeams retrieves a list of all teams.
func (m *TeamModel) ListTeams(ctx context.Context) ([]*ent.Team, error) {
	return m.client.Team.
		Query().
		All(ctx)
}

// GetTeamCountry retrieves the country of a team.
func (m *TeamModel) GetTeamCountry(ctx context.Context, teamID int) (*ent.Country, error) {
	return m.client.Team.
		Query().
		Where(team.ID(teamID)).
		QueryCountry().
		Only(ctx)
}

// GetTeamLeagues retrieves the leagues of a team.
func (m *TeamModel) GetTeamLeagues(ctx context.Context, teamID int) ([]*ent.League, error) {
	return m.client.Team.
		Query().
		Where(team.ID(teamID)).
		QueryLeagues().
		All(ctx)
}

// GetTeamStandings retrieves the standings of a team.
func (m *TeamModel) GetTeamStandings(ctx context.Context, teamID int) ([]*ent.Standings, error) {
	return m.client.Team.
		Query().
		Where(team.ID(teamID)).
		QueryStandings().
		All(ctx)
}

// GetTeamPlayers retrieves the players of a team.
func (m *TeamModel) GetTeamPlayers(ctx context.Context, teamID int) ([]*ent.Player, error) {
	return m.client.Team.
		Query().
		Where(team.ID(teamID)).
		QueryPlayers().
		All(ctx)
}
