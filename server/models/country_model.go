package models

import (
	"context"
	"mapeleven/models/ent"
)

// CreateCountryInput holds the input data needed to create a new country record.
type CreateCountryInput struct {
	Name string
	Code string
	Flag string
}

// UpdateCountryInput holds the input data needed to update an existing country record.
type UpdateCountryInput struct {
	ID   int
	Flag *string
}

// CountryModel handles the CRUD operations for the Country entity.
type CountryModel struct {
	client *ent.Client
}

// NewCountryModel creates a new CountryModel instance.
func NewCountryModel(client *ent.Client) *CountryModel {
	return &CountryModel{client: client}
}

// CreateCountry creates a new country record.
func (cm *CountryModel) CreateCountry(input CreateCountryInput) (*ent.Country, error) {
	return cm.client.Country.
		Create().
		SetName(input.Name).
		SetCode(input.Code).
		SetFlag(input.Flag).
		Save(context.Background())
}

// UpdateCountry updates an existing country record.
func (cm *CountryModel) UpdateCountry(input UpdateCountryInput) (*ent.Country, error) {
	updater := cm.client.Country.UpdateOneID(input.ID)
	if input.Flag != nil {
		updater.SetFlag(*input.Flag)
	}
	return updater.Save(context.Background())
}

// DeleteCountry deletes a country record by ID.
func (cm *CountryModel) DeleteCountry(id int) error {
	return cm.client.Country.
		DeleteOneID(id).
		Exec(context.Background())
}

// GetCountry retrieves a country record by ID.
func (cm *CountryModel) GetCountry(id int) (*ent.Country, error) {
	return cm.client.Country.
		Get(context.Background(), id)
}
