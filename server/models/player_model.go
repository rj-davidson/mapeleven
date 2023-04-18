package models

import (
	"context"
	"errors"
	"mapeleven/models/ent"
	"mapeleven/models/ent/player"
	"mapeleven/models/ent/team"
)

// CreatePlayerInput holds the required fields to create a new player.
type CreatePlayerInput struct {
	ID        int
	Name      string
	Firstname string
	Lastname  string
	Age       int
	Height    float64
	Weight    float64
	Injured   bool
	Photo     string
}

// UpdatePlayerInput holds the fields to update an existing player.
type UpdatePlayerInput struct {
	ID        int
	Name      *string
	Firstname *string
	Lastname  *string
	Age       *int
	Height    *float64
	Weight    *float64
	Injured   *bool
	Photo     *string
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
		SetID(input.ID).
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
	update := m.client.Player.UpdateOneID(input.ID)

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

// GetPlayerByID retrieves a player by ID.
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

// GetPlayerTeams retrieves the teams of a player.
func (m *PlayerModel) GetPlayerTeams(ctx context.Context, playerID int) ([]*ent.Team, error) {
	teams, err := m.client.Player.
		Query().
		Where(player.ID(playerID)).
		QueryTeams().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return teams, nil
}

// FilterPlayersByTeams retrieves a list of players who belong to the specified team(s)
func (m *PlayerModel) FilterPlayersByTeams(ctx context.Context, teamIDs []int) ([]*ent.Player, error) {
	players, err := m.client.Player.
		Query().
		Where(player.HasTeamsWith(team.IDIn(teamIDs...))).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return players, nil
}
