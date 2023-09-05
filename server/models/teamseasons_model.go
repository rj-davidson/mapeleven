package models

import (
	"context"
	"mapeleven/db/ent"
	_ "mapeleven/db/ent/teamseason"
)

// CreateTeamSeasonInput holds the required fields to create a new team season.
type CreateTeamSeasonInput struct {
	TeamSeasonID int
	TeamID       int
	SeasonID     int
}

// UpdateTeamSeasonInput holds the fields to update an existing team season.
type UpdateTeamSeasonInput struct {
	ID       int
	TeamID   *int
	SeasonID *int
}

// TeamSeasonsModel defines the methods for managing the team season data.
type TeamSeasonsModel struct {
	client *ent.Client
}

// NewTeamSeasonsModel creates a new TeamSeasonsModel.
func NewTeamSeasonsModel(client *ent.Client) *TeamSeasonsModel {
	return &TeamSeasonsModel{client: client}
}

// CreateTeamSeason creates a new team season.
func (m *TeamSeasonsModel) CreateTeamSeason(ctx context.Context, input CreateTeamSeasonInput) (*ent.TeamSeason, error) {
	teamSeason, err := m.client.TeamSeason.
		Create().
		SetTeamSeasonID(input.TeamSeasonID).
		SetTeamID(input.TeamID).
		SetSeasonID(input.SeasonID).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return teamSeason, nil
}

// UpdateTeamSeason updates an existing team season.
func (m *TeamSeasonsModel) UpdateTeamSeason(ctx context.Context, input UpdateTeamSeasonInput) (*ent.TeamSeason, error) {
	updater := m.client.TeamSeason.UpdateOneID(input.ID)
	if input.TeamID != nil {
		updater.SetTeamID(*input.TeamID)
	}
	if input.SeasonID != nil {
		updater.SetSeasonID(*input.SeasonID)
	}
	return updater.Save(ctx)
}

// DeleteTeamSeason deletes a team season.
func (m *TeamSeasonsModel) DeleteTeamSeason(ctx context.Context, teamSeasonID int) error {
	return m.client.TeamSeason.
		DeleteOneID(teamSeasonID).
		Exec(ctx)
}

// GetTeamSeasonByID retrieves a team season by ID.
func (m *TeamSeasonsModel) GetTeamSeasonByID(ctx context.Context, teamSeasonID int) (*ent.TeamSeason, error) {
	return m.client.TeamSeason.
		Get(ctx, teamSeasonID)
}

// ListTeamSeasons retrieves a list of all team seasons.
func (m *TeamSeasonsModel) ListTeamSeasons(ctx context.Context) ([]*ent.TeamSeason, error) {
	return m.client.TeamSeason.
		Query().
		All(ctx)
}
