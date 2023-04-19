package models

import (
	"context"
	"fmt"
	"mapeleven/db/ent"
	"mapeleven/db/ent/country"
)

// CreateCountryInput holds the input data needed to create a new country record.
type CreateCountryInput struct {
	Name string
	Code string
	Flag string
}

// UpdateCountryInput holds the input data needed to update an existing country record.
type UpdateCountryInput struct {
	Name string
	Code string
	Flag string
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
	c, err := cm.client.Country.
		Query().
		Where(country.CodeEQ(input.Code)).
		First(context.Background())

	if err != nil {
		return nil, err
	}

	updatedCountry, err := c.Update().
		SetFlag(input.Flag).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return updatedCountry, nil
}

// DeleteCountry deletes a country record by code.
func (cm *CountryModel) DeleteCountry(code string) error {
	numDeleted, err := cm.client.Country.
		Delete().
		Where(country.CodeEQ(code)).
		Exec(context.Background())
	if err != nil {
		return err
	}
	if numDeleted == 0 {
		return fmt.Errorf("no country found with code %s", code)
	}
	return nil
}

// GetCountryByCode retrieves a country record by country code.
func (cm *CountryModel) GetCountryByCode(code string) (*ent.Country, error) {
	return cm.client.Country.
		Query().
		Where(country.CodeEQ(code)).
		First(context.Background())
}
