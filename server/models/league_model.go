package models

import (
	"context"
	"mapeleven/models/ent"
	"mapeleven/models/ent/league"
)

// CreateLeagueInput holds the input data needed to create a new league record.
type CreateLeagueInput struct {
	ID      int
	Name    string
	Type    league.Type
	Logo    string
	Country int
}

// UpdateLeagueInput holds the input data needed to update an existing league record.
type UpdateLeagueInput struct {
	ID      int
	Name    *string
	Type    *league.Type
	Logo    *string
	Country *int
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
	return lm.client.League.
		Create().
		SetID(input.ID).
		SetName(input.Name).
		SetType(input.Type).
		SetLogo(input.Logo).
		SetCountryID(input.Country).
		Save(context.Background())
}

// UpdateLeague updates an existing league record.
func (lm *LeagueModel) UpdateLeague(input UpdateLeagueInput) (*ent.League, error) {
	updater := lm.client.League.UpdateOneID(input.ID)
	if input.Name != nil {
		updater.SetName(*input.Name)
	}
	if input.Type != nil {
		updater.SetType(*input.Type)
	}
	if input.Logo != nil {
		updater.SetLogo("public/images/leagues/" + *input.Logo)
	}
	if input.Country != nil {
		updater.SetCountryID(*input.Country)
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
func (lm *LeagueModel) GetLeague(id int) (*ent.League, error) {
	return lm.client.League.
		Get(context.Background(), id)
}
