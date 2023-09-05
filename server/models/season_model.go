package models

import (
	"context"
	"errors"
	"mapeleven/db/ent"
	"mapeleven/db/ent/league"
	"mapeleven/db/ent/season"
	"time"
)

// CreateSeasonInput holds the required fields to create a new season.
type CreateSeasonInput struct {
	Year    int
	Start   time.Time
	End     time.Time
	Current bool
	League  int
}

// UpdateSeasonInput holds the fields to update an existing season.
type UpdateSeasonInput struct {
	ID      int
	Year    *int
	Start   *time.Time
	End     *time.Time
	Current *bool
	League  *int
}

// DeleteSeasonInput holds the required fields to delete a season.
type DeleteSeasonInput struct {
	ID int
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
func (m *SeasonModel) CreateSeason(ctx context.Context, input CreateSeasonInput) (*ent.Season, error) {
	leagueEnt, err := m.client.League.
		Get(ctx, input.League)

	if err != nil {
		return nil, err
	}

	se, err := m.client.Season.
		Create().
		SetYear(input.Year).
		SetStart(input.Start).
		SetEnd(input.End).
		SetCurrent(input.Current).
		SetLeague(leagueEnt).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return se, nil
}

// UpdateSeason updates an existing season.
func (m *SeasonModel) UpdateSeason(ctx context.Context, input UpdateSeasonInput) (*ent.Season, error) {
	update := m.client.Season.UpdateOneID(input.ID)

	if input.Year != nil {
		update.SetYear(*input.Year)
	}
	if input.Start != nil {
		update.SetStart(*input.Start)
	}
	if input.End != nil {
		update.SetEnd(*input.End)
	}
	if input.Current != nil {
		update.SetCurrent(*input.Current)
	}
	if input.League != nil {
		leagueEnt, err := m.client.League.
			Get(ctx, *input.League)

		if err != nil {
			return nil, err
		}

		update.SetLeague(leagueEnt)
	}

	se, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return se, nil
}

// DeleteSeason deletes a season.
func (m *SeasonModel) DeleteSeason(ctx context.Context, input DeleteSeasonInput) error {
	err := m.client.Season.
		DeleteOneID(input.ID).
		Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("season not found")
		}
		return err
	}
	return nil
}

// GetSeasonByID retrieves a season by ID.
func (m *SeasonModel) GetSeasonByID(ctx context.Context, id int) (*ent.Season, error) {
	se, err := m.client.Season.
		Query().
		Where(season.ID(id)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("se not found")
		}
		return nil, err
	}
	return se, nil
}

// ListSeasons retrieves a list of all seasons.
func (m *SeasonModel) ListSeasons(ctx context.Context) ([]*ent.Season, error) {
	ses, err := m.client.Season.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return ses, nil
}

// GetSeasonLeague retrieves the league of a season.
func (m *SeasonModel) GetSeasonLeague(ctx context.Context, seasonID int) (*ent.League, error) {
	league, err := m.client.Season.
		Query().
		Where(season.ID(seasonID)).
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

// GetLeagueSeason retrieves the season of a league.
func (m *SeasonModel) GetLeagueSeason(ctx context.Context, leagueID int) ([]*ent.Season, error) {
	seasons, err := m.client.League.
		Query().
		Where(league.ID(leagueID)).
		QuerySeason().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return seasons, nil
}

// ListSeasonIDs retrieves a list of all season IDs.
func (m *SeasonModel) ListSeasonIDs(ctx context.Context) ([]int, error) {
	seasons, err := m.client.Season.
		Query().
		Select(season.FieldID).
		All(ctx)

	if err != nil {
		return nil, err
	}

	seasonIDs := make([]int, len(seasons))
	for i, s := range seasons {
		seasonIDs[i] = s.ID
	}
	return seasonIDs, nil
}
