package models

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/utils"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/country"
	"context"
	"fmt"
)

// Club holds the required fields for the club schema.
type Club struct {
	ApiFootballId int    `json:"id"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Slug          string `json:"-"`
	Founded       int    `json:"founded"`
	National      bool   `json:"national"`
	Logo          string `json:"logo"`
	Country       string `json:"country"`
}

// ClubModel defines the methods for managing the team data.
type ClubModel struct {
	client *ent.Client
}

// NewClubModel creates a new TeamModel.
func NewClubModel(client *ent.Client) *ClubModel {
	return &ClubModel{client: client}
}

func (m *ClubModel) CreateClub(ctx context.Context, input Club) (*ent.Club, error) {
	// Get the country by code
	// Find the country by its code
	c, err := m.client.Country.
		Query().
		Where(country.NameEQ(input.Country)).
		Only(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to find country: %w", err)
	}
	if input.Slug == "" {
		input.Slug = utils.Slugify(input.Name)
	}

	return m.client.Club.
		Create().
		SetApiFootballId(input.ApiFootballId).
		SetSlug(utils.Slugify(input.Name)).
		SetName(input.Name).
		SetCode(input.Code).
		SetFounded(input.Founded).
		SetNational(input.National).
		SetLogo(input.Logo).
		SetCountry(c).
		Save(ctx)
}

// UpdateClub updates an existing team.
func (m *ClubModel) UpdateClub(ctx context.Context, input Club) error {
	// Find the country by its code
	c, err := m.client.Country.
		Query().
		Where(country.NameEQ(input.Country)).
		Only(context.Background())

	if err != nil {
		return err
	}

	updater := m.client.Club.Update().Where(club.ApiFootballIdEQ(input.ApiFootballId))

	updater.SetName(input.Name)
	updater.SetCode(input.Code)
	updater.SetFounded(input.Founded)
	updater.SetNational(input.National)
	updater.SetLogo(input.Logo)
	updater.SetCountry(c)

	_, err = updater.Save(ctx)

	if err != nil {
		return err
	} else {
		return nil
	}
}

// GetByApiFootballId retrieves a club by FootballApiId.
func (m *ClubModel) GetByApiFootballId(ctx context.Context, apiFootballId int) (*ent.Club, error) {
	return m.client.Club.
		Query().
		Where(club.ApiFootballIdEQ(apiFootballId)).
		Only(ctx)
}

// ListAll retrieves a list of all Clubs.
func (m *ClubModel) ListAll(ctx context.Context) ([]*ent.Club, error) {
	return m.client.Club.
		Query().
		All(ctx)
}

// Exists checks if a club exists.
func (m *ClubModel) Exists(ctx context.Context, apiFootballId int) bool {
	count, _ := m.client.Club.
		Query().
		Where(club.ApiFootballIdEQ(apiFootballId)).
		Count(ctx)

	return count > 0
}
