package models

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/league"
	"mapeleven/db/ent/season"
	"time"
)

// CreateSeasonInput holds the required fields to create a new season.
type CreateSeasonInput struct {
	Year    int
	Start   string
	End     string
	Current bool
}

// UpdateSeasonInput holds the fields to update an existing season.
type UpdateSeasonInput struct {
	ID      int
	Start   string
	End     string
	Current bool
}

// SeasonModel defines the methods for managing the season data.
type SeasonModel struct {
	client *ent.Client
}

// NewSeasonModel creates a new SeasonModel.
func NewSeasonModel(client *ent.Client) *SeasonModel {
	return &SeasonModel{client: client}
}

// CreateSeason creates a new season.
func (m *SeasonModel) CreateSeason(ctx context.Context, input CreateSeasonInput, slug string, leagueID int) (*ent.Season, error) {
	startDate, err := time.Parse("2006-01-02", input.Start)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse("2006-01-02", input.End)
	if err != nil {
		return nil, err
	}

	return m.client.Season.
		Create().
		SetSlug(slug).
		SetYear(input.Year).
		SetStartDate(startDate).
		SetEndDate(endDate).
		SetCurrent(input.Current).
		SetLeagueID(leagueID).
		Save(ctx)
}

// UpdateSeason updates an existing season.
func (m *SeasonModel) UpdateSeason(ctx context.Context, input UpdateSeasonInput) (*ent.Season, error) {
	startDate, err := time.Parse("2006-01-02", input.Start)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse("2006-01-02", input.End)
	if err != nil {
		return nil, err
	}

	return m.client.Season.
		UpdateOneID(input.ID).
		SetStartDate(startDate).
		SetEndDate(endDate).
		SetCurrent(input.Current).
		Save(ctx)
}

// GetCurrent retrieves the seasons that are currently active.
func (m *SeasonModel) GetCurrent() ([]*ent.Season, error) {
	return m.client.Season.
		Query().
		Where(season.CurrentEQ(true)).
		All(context.Background())
}

// SeasonExistsByLeagueIdAndYear checks if a season exists by its league ID and year.
func (m *SeasonModel) SeasonExistsByLeagueIdAndYear(ctx context.Context, leagueID int, year int) *ent.Season {
	s, err := m.client.Season.
		Query().
		WithLeague().
		Where(
			season.HasLeagueWith(league.IDEQ(leagueID)),
			season.YearEQ(year),
		).
		Only(ctx)
	if err != nil {
		return nil
	}
	return s
}
