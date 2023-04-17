package models

import (
	"context"
	"errors"
	"mapeleven/models/ent"
	"mapeleven/models/ent/team"
)

// CreateTeamInput holds the required fields to create a new team.
type CreateTeamInput struct {
	ID       int
	Name     string
	Code     string
	Founded  int
	National bool
	Logo     string
	Country  int
}

// UpdateTeamInput holds the fields to update an existing team.
type UpdateTeamInput struct {
	ID       int
	Name     *string
	Code     *string
	Founded  *int
	National *bool
	Logo     *string
	Country  *int
}

// DeleteTeamInput holds the required fields to delete a team.
type DeleteTeamInput struct {
	ID int
}

// TeamModel defines the methods for managing the team data.
type TeamModel struct {
	client *ent.Client
}

// NewTeamModel creates a new TeamModel.
func NewTeamModel(client *ent.Client) *TeamModel {
	return &TeamModel{client: client}
}

// CreateTeam creates a new team.
func (m *TeamModel) CreateTeam(ctx context.Context, input CreateTeamInput) (*ent.Team, error) {
	t, err := m.client.Team.
		Create().
		SetID(input.ID).
		SetName(input.Name).
		SetCode(input.Code).
		SetFounded(input.Founded).
		SetNational(input.National).
		SetLogo(input.Logo).
		SetCountryID(input.Country).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return t, nil
}

// UpdateTeam updates an existing team.
func (m *TeamModel) UpdateTeam(ctx context.Context, input UpdateTeamInput) (*ent.Team, error) {
	update := m.client.Team.UpdateOneID(input.ID)

	if input.Name != nil {
		update.SetName(*input.Name)
	}
	if input.Code != nil {
		update.SetCode(*input.Code)
	}
	if input.Founded != nil {
		update.SetFounded(*input.Founded)
	}
	if input.National != nil {
		update.SetNational(*input.National)
	}
	if input.Logo != nil {
		update.SetLogo(*input.Logo)
	}
	if input.Country != nil {
		update.SetCountryID(*input.Country)
	}

	t, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// DeleteTeam deletes a team.
func (m *TeamModel) DeleteTeam(ctx context.Context, input DeleteTeamInput) error {
	err := m.client.Team.
		DeleteOneID(input.ID).
		Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("team not found")
		}
		return err
	}
	return nil
}

// GetTeamByID retrieves a team by ID.
func (m *TeamModel) GetTeamByID(ctx context.Context, id int) (*ent.Team, error) {
	t, err := m.client.Team.
		Query().
		Where(team.ID(id)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("t not found")
		}
		return nil, err
	}
	return t, nil
}

// ListTeams retrieves a list of all teams.
func (m *TeamModel) ListTeams(ctx context.Context) ([]*ent.Team, error) {
	teams, err := m.client.Team.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return teams, nil
}

// GetTeamCountry retrieves the country of a team.
func (m *TeamModel) GetTeamCountry(ctx context.Context, teamID int) (*ent.Country, error) {
	country, err := m.client.Team.
		Query().
		Where(team.ID(teamID)).
		QueryCountry().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("country not found")
		}
		return nil, err
	}
	return country, nil
}

// GetTeamLeagues retrieves the leagues of a team.
func (m *TeamModel) GetTeamLeagues(ctx context.Context, teamID int) ([]*ent.League, error) {
	leagues, err := m.client.Team.
		Query().
		Where(team.ID(teamID)).
		QueryLeagues().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return leagues, nil
}

// GetTeamStandings retrieves the standings of a team.
func (m *TeamModel) GetTeamStandings(ctx context.Context, teamID int) ([]*ent.Standings, error) {
	standings, err := m.client.Team.
		Query().
		Where(team.ID(teamID)).
		QueryStandings().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return standings, nil
}

// GetTeamPlayers retrieves the players of a team.
func (m *TeamModel) GetTeamPlayers(ctx context.Context, teamID int) ([]*ent.Player, error) {
	players, err := m.client.Team.
		Query().
		Where(team.ID(teamID)).
		QueryPlayers().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return players, nil
}
