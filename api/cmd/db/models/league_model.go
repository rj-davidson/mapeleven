package models

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/utils"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/country"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"context"
)

// CreateLeagueInput holds the input data needed to create a new league record.
type CreateLeagueInput struct {
	FootballApiId int         `json:"id"`
	Name          string      `json:"name"`
	Type          league.Type `json:"type"`
	Logo          string      `json:"logo"`
	Country       string      `json:"country"`
}

// UpdateLeagueInput holds the input data needed to update an existing league record.
type UpdateLeagueInput struct {
	ID      int
	Name    *string
	Type    *league.Type
	Logo    *string
	Country *string
}

// LeagueModel handles the CRUD operations for the SeasonID entity.
type LeagueModel struct {
	client *ent.Client
}

// NewLeagueModel creates a new LeagueModel instance.
func NewLeagueModel(client *ent.Client) *LeagueModel {
	return &LeagueModel{client: client}
}

// CreateLeague creates a new league record.
func (lm *LeagueModel) CreateLeague(ctx context.Context, input CreateLeagueInput) (*ent.League, error) {
	c, err := lm.client.Country.
		Query().
		Where(country.NameEQ(input.Country)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	// Use the country's FootballApiId instead of the country code
	return lm.client.League.
		Create().
		SetFootballApiId(input.FootballApiId).
		SetSlug(utils.Slugify(input.Name)).
		SetName(input.Name).
		SetType(input.Type).
		SetLogo(input.Logo).
		SetCountryID(c.ID).
		Save(ctx)
}

// UpdateLeague updates an existing league record.
func (lm *LeagueModel) UpdateLeague(ctx context.Context, input UpdateLeagueInput) (*ent.League, error) {
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
		// Find the country by its name
		c, err := lm.client.Country.
			Query().
			Where(country.NameEQ(*input.Country)).
			Only(ctx)

		if err != nil {
			// Handle error properly, maybe country not found or other issue
			return nil, err
		}

		updater.SetCountryID(c.ID)
	}

	return updater.Save(ctx)
}

// GetLeagueByID retrieves a league record by FootballApiId.
func (lm *LeagueModel) GetLeagueByID(ctx context.Context, id int) (*ent.League, error) {
	return lm.client.League.
		Get(ctx, id)
}

// ListLeagues retrieves a list of all leagues.
func (lm *LeagueModel) ListLeagues(ctx context.Context) ([]*ent.League, error) {
	return lm.client.League.
		Query().
		All(ctx)
}

// LeagueExistsByFootballApiId checks if a league exists by FootballApiId.
func (lm *LeagueModel) LeagueExistsByFootballApiId(ctx context.Context, id int) *ent.League {
	l, _ := lm.client.League.
		Query().
		Where(league.FootballApiIdEQ(id)).
		Only(ctx)
	return l
}

// GetCurrentSeasonForLeague retrieves the current season for a league.
func (lm *LeagueModel) GetCurrentSeasonForLeague(ctx context.Context, l *ent.League) (*ent.Season, error) {
	return lm.client.Season.
		Query().
		Where(
			season.HasLeagueWith(league.IDEQ(l.ID)),
			season.CurrentEQ(true),
		).
		WithLeague().
		Only(ctx)
}

func (lm *LeagueModel) ListAll(ctx context.Context) ([]*ent.League, error) {
	return lm.client.League.
		Query().
		All(ctx)
}
