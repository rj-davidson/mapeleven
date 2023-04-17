package models

import (
	"context"
	"errors"
	"mapeleven/models/ent"
	"mapeleven/models/ent/standings"
)

// CreateStandingsInput holds the required fields to create a new standings entry.
type CreateStandingsInput struct {
	Rank        int
	Description string
	League      int
	Team        int
}

// UpdateStandingsInput holds the fields to update an existing standings entry.
type UpdateStandingsInput struct {
	ID          int
	Rank        *int
	Description *string
	League      *int
	Team        *int
}

// DeleteStandingsInput holds the required fields to delete a standings entry.
type DeleteStandingsInput struct {
	ID int
}

// StandingsModel defines the methods for managing the standings data.
type StandingsModel struct {
	client *ent.Client
}

// NewStandingsModel creates a new StandingsModel.
func NewStandingsModel(client *ent.Client) *StandingsModel {
	return &StandingsModel{client: client}
}

// CreateStandings creates a new standings entry.
func (m *StandingsModel) CreateStandings(ctx context.Context, input CreateStandingsInput) (*ent.Standings, error) {
	stands, err := m.client.Standings.
		Create().
		SetRank(input.Rank).
		SetDescription(input.Description).
		SetLeagueID(input.League).
		SetTeamID(input.Team).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return stands, nil
}

// UpdateStandings updates an existing standings entry.
func (m *StandingsModel) UpdateStandings(ctx context.Context, input UpdateStandingsInput) (*ent.Standings, error) {
	update := m.client.Standings.UpdateOneID(input.ID)

	if input.Rank != nil {
		update.SetRank(*input.Rank)
	}
	if input.Description != nil {
		update.SetDescription(*input.Description)
	}
	if input.League != nil {
		update.SetLeagueID(*input.League)
	}
	if input.Team != nil {
		update.SetTeamID(*input.Team)
	}

	stands, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return stands, nil
}

// DeleteStandings deletes a standings entry.
func (m *StandingsModel) DeleteStandings(ctx context.Context, input DeleteStandingsInput) error {
	err := m.client.Standings.
		DeleteOneID(input.ID).
		Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("standings entry not found")
		}
		return err
	}
	return nil
}

// GetStandingsByID retrieves a standings entry by ID.
func (m *StandingsModel) GetStandingsByID(ctx context.Context, id int) (*ent.Standings, error) {
	stands, err := m.client.Standings.
		Query().
		Where(standings.ID(id)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("stands entry not found")
		}
		return nil, err
	}
	return stands, nil
}

// ListStandings retrieves a list of all standings entries.
func (m *StandingsModel) ListStandings(ctx context.Context) ([]*ent.Standings, error) {
	standingsEntries, err := m.client.Standings.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return standingsEntries, nil
}

// GetStandingsLeague retrieves the league of a standings entry.
func (m *StandingsModel) GetStandingsLeague(ctx context.Context, standingsID int) (*ent.League, error) {
	league, err := m.client.Standings.
		Query().
		Where(standings.ID(standingsID)).
		QueryLeague().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("league not found")
		}
		return nil, err
	}
	return league, nil
}

// GetStandingsTeam retrieves the team of a standings entry.
func (m *StandingsModel) GetStandingsTeam(ctx context.Context, standingsID int) (*ent.Team, error) {
	team, err := m.client.Standings.
		Query().
		Where(standings.ID(standingsID)).
		QueryTeam().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}
	return team, nil
}
