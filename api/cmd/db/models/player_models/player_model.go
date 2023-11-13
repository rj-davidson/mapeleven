package player_models

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/utils"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"context"
	"errors"
)

// CreatePlayerInput holds the required fields to create a new player.
type CreatePlayerInput struct {
	ApiFootballId int
	Name          string
	Firstname     string
	Lastname      string
	Age           int
	Height        string
	Weight        string
	Injured       bool
	Photo         string
	LeagueID      int
	Season        int
}

// UpdatePlayerInput holds the fields to update an existing player.
type UpdatePlayerInput struct {
	ApiFootballId          int
	Name                   *string
	Firstname              *string
	Lastname               *string
	Age                    *int
	Height                 *string
	Weight                 *string
	Injured                *bool
	Photo                  *string
	LeagueID               *int
	Season                 *int
	UpdatePlayerStatsInput *UpdatePlayerStatsInput
}

type UpdatePlayerStatsInput struct {
	LeagueID int
	Season   int
}

// DeletePlayerInput holds the required fields to delete a player.
type DeletePlayerInput struct {
	ID int
}

// PlayerModel defines the methods for managing the player data.
type PlayerModel struct {
	client *ent.Client
}

// NewPlayerModel creates a new PlayerModel.
func NewPlayerModel(client *ent.Client) *PlayerModel {
	return &PlayerModel{client: client}
}

// CreatePlayer creates a new player.
func (m *PlayerModel) CreatePlayer(ctx context.Context, input CreatePlayerInput) (*ent.Player, error) {
	p, err := m.client.Player.
		Create().
		SetApiFootballId(input.ApiFootballId).
		SetSlug(utils.Slugify(input.Name)).
		SetName(input.Name).
		SetFirstname(input.Firstname).
		SetLastname(input.Lastname).
		SetAge(input.Age).
		SetHeight(input.Height).
		SetWeight(input.Weight).
		SetInjured(input.Injured).
		SetPhoto(input.Photo).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return p, nil
}

// UpdatePlayer updates an existing player.
func (m *PlayerModel) UpdatePlayer(ctx context.Context, input UpdatePlayerInput) (*ent.Player, error) {
	pl, err := m.client.Player.Query().Where(player.ApiFootballIdEQ(input.ApiFootballId)).Only(ctx)
	if err != nil {
		return nil, err
	}
	update := m.client.Player.UpdateOne(pl)

	if input.Name != nil {
		update.SetName(*input.Name)
	}
	if input.Firstname != nil {
		update.SetFirstname(*input.Firstname)
	}
	if input.Lastname != nil {
		update.SetLastname(*input.Lastname)
	}
	if input.Age != nil {
		update.SetAge(*input.Age)
	}
	if input.Height != nil {
		update.SetHeight(*input.Height)
	}
	if input.Weight != nil {
		update.SetWeight(*input.Weight)
	}
	if input.Injured != nil {
		update.SetInjured(*input.Injured)
	}
	if input.Photo != nil {
		update.SetPhoto(*input.Photo)
	}

	p, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// DeletePlayer deletes a player.
func (m *PlayerModel) DeletePlayer(ctx context.Context, input DeletePlayerInput) error {
	err := m.client.Player.
		DeleteOneID(input.ID).
		Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("player not found")
		}
		return err
	}
	return nil
}

// GetPlayerByID retrieves a player by FootballApiId.
func (m *PlayerModel) GetPlayerByID(ctx context.Context, id int) (*ent.Player, error) {
	r, err := m.client.Player.
		Query().
		Where(player.ID(id)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("r not found")
		}
		return nil, err
	}
	return r, nil
}

// ListPlayers retrieves a list of all players.
func (m *PlayerModel) ListPlayers(ctx context.Context) ([]*ent.Player, error) {
	players, err := m.client.Player.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return players, nil
}

// GetPlayerBirth retrieves the birth of a player.
func (m *PlayerModel) GetPlayerBirth(ctx context.Context, playerID int) (*ent.Birth, error) {
	birth, err := m.client.Player.
		Query().
		Where(player.ID(playerID)).
		QueryBirth().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("birth not found")
		}
		return nil, err
	}
	return birth, nil
}

// Exists checks if a player exists.
func (m *PlayerModel) Exists(ctx context.Context, id int) bool {
	count, _ := m.client.Player.
		Query().
		Where(player.IDEQ(id)).
		Count(ctx)

	return count > 0
}

// WithTransaction method to execute a transaction
func (m *PlayerModel) WithTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error {
	// Start a new transaction
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return err
	}
	// Execute logic
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
