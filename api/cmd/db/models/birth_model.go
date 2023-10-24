package models

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"time"
)

// CreateBirthInput holds the input data needed to create a new birth record.
type CreateBirthInput struct {
	Date    time.Time
	Place   string
	Country string
}

// UpdateBirthInput holds the input data needed to update an existing birth record.
type UpdateBirthInput struct {
	ID      int
	Date    *time.Time
	Place   *string
	Country *string
}

// BirthModel handles the CRUD operations for the Birth entity.
type BirthModel struct {
	client *ent.Client
}

// NewBirthModel creates a new BirthModel instance.
func NewBirthModel(client *ent.Client) *BirthModel {
	return &BirthModel{client: client}
}

// GetBirth retrieves a birth record by FootballApiId.
func (bm *BirthModel) GetBirth(id int) (*ent.Birth, error) {
	return bm.client.Birth.
		Get(context.Background(), id)
}

// UpdateBirth updates an existing birth record.
func (bm *BirthModel) UpdateBirth(input UpdateBirthInput) (*ent.Birth, error) {
	updater := bm.client.Birth.UpdateOneID(input.ID)
	if input.Date != nil {
		updater.SetDate(*input.Date)
	}
	if input.Place != nil {
		updater.SetPlace(*input.Place)
	}
	if input.Country != nil {
		updater.SetCountry(*input.Country)
	}
	return updater.Save(context.Background())
}

// DeleteBirth deletes a birth record by FootballApiId.
func (bm *BirthModel) DeleteBirth(id int) error {
	return bm.client.Birth.
		DeleteOneID(id).
		Exec(context.Background())
}

// CreateBirth creates a new birth record.
func (bm *BirthModel) CreateBirth(input CreateBirthInput) (*ent.Birth, error) {
	return bm.client.Birth.
		Create().
		SetDate(input.Date).
		SetPlace(input.Place).
		SetCountry(input.Country).
		Save(context.Background())
}
