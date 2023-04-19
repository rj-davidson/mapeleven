package models

import (
	"context"
	"errors"
	"mapeleven/db/ent"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/league"
)

// CreateLeagueInput holds the input data needed to create a new league record.
type CreateLeagueInput struct {
	ID      int
	Name    string
	Type    league.Type
	Logo    string
	Country string
}

// UpdateLeagueInput holds the input data needed to update an existing league record.
type UpdateLeagueInput struct {
	ID      int
	Name    *string
	Type    *league.Type
	Logo    *string
	Country *string
}

// LeagueModel handles the CRUD operations for the League entity.
type LeagueModel struct {
	client *ent.Client
}

// NewLeagueModel creates a new LeagueModel instance.
func NewLeagueModel(client *ent.Client) *LeagueModel {
	return &LeagueModel{client: client}
}

// CreateLeague creates a new league record.
func (lm *LeagueModel) CreateLeague(input CreateLeagueInput) (*ent.League, error) {
	// Find the country by its code
	c, err := lm.client.Country.
		Query().
		Where(country.CodeEQ(input.Country)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	// Use the country's ID instead of the country code
	return lm.client.League.
		Create().
		SetID(input.ID).
		SetName(input.Name).
		SetType(input.Type).
		SetLogo(input.Logo).
		SetCountryID(c.ID). // Set the country ID here
		Save(context.Background())
}

// UpdateLeague updates an existing league record.
func (lm *LeagueModel) UpdateLeague(input UpdateLeagueInput) (*ent.League, error) {
	// Find the country by its code
	c, err := lm.client.Country.
		Query().
		Where(country.CodeEQ(*input.Country)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	updater := lm.client.League.UpdateOneID(input.ID)
	if input.Name != nil {
		updater.SetName(*input.Name)
	}
	if input.Type != nil {
		updater.SetType(*input.Type)
	}
	if input.Logo != nil {
		updater.SetLogo(*input.Logo)
	}
	if input.Country != nil {
		updater.SetCountryID(c.ID)
	}
	return updater.Save(context.Background())
}

// DeleteLeague deletes a league record by ID.
func (lm *LeagueModel) DeleteLeague(id int) error {
	return lm.client.League.
		DeleteOneID(id).
		Exec(context.Background())
}

// GetLeague retrieves a league record by ID.
func (lm *LeagueModel) GetLeague(ctx context.Context, id int) (*ent.League, error) {
	return lm.client.League.
		Get(ctx, id)
}

// GetLeagueSeason retrieves the season of a league.
func (lm *LeagueModel) GetLeagueSeason(ctx context.Context, leagueID int) (*ent.Season, error) {
	season, err := lm.client.League.
		Query().
		Where(league.ID(leagueID)).
		QuerySeason().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("season not found")
		}
		return nil, err
	}
	return season, nil
}

// ListLeagues retrieves a list of all leagues.
func (lm *LeagueModel) ListLeagues(ctx context.Context) ([]*ent.League, error) {
	leagues, err := lm.client.League.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return leagues, nil
}
