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
	Statistics    PlayerStats
	Info          PlayerInfo
}

// UpdatePlayerInput holds the fields to update an existing player.
type UpdatePlayerInput struct {
	ApiFootballId int
	Name          *string
	Firstname     *string
	Lastname      *string
	Age           *int
	Height        *string
	Weight        *string
	Injured       *bool
	Photo         *string
	LeagueID      *int
	Season        *int
	Statistics    *PlayerStats
	Info          *PlayerInfo
}

type PSTeam struct {
	ApiFootballId int
	Name          string
	Logo          string
}

type PSLeague struct {
	ApiFootballId int
	Name          string
	Country       string
	Logo          string
	Flag          string
	Season        int
}

type PSGames struct {
	Appearances int
	Lineups     int
	Minutes     int
	Number      int
	Position    string
	Rating      string
	Captain     bool
}

type PSSubstitutes struct {
	In    int
	Out   int
	Bench int
}

type PSShots struct {
	Total int
	On    int
}

type PSGoals struct {
	Total    int
	Conceded int
	Assists  int
	Saves    int
}
type PSPasses struct {
	Total    int
	Key      int
	Accuracy int
}

type PSTackles struct {
	Total         int
	Blocks        int
	Interceptions int
}

type PSDuels struct {
	Total int
	Won   int
}

type PSDribbles struct {
	Attempts int
	Success  int
	Past     int
}

type PSFouls struct {
	Drawn     int
	Committed int
}
type PSCards struct {
	Yellow    int
	YellowRed int
	Red       int
}

type PSPenalty struct {
	Won      int
	Commited int
	Scored   int
	Saved    int
	Missed   int
}

type PlayerStats struct {
	Team        PSTeam
	League      PSLeague
	Games       PSGames
	Substitutes PSSubstitutes
	Shots       PSShots
	Goals       PSGoals
	Passes      PSPasses
	Tackles     PSTackles
	Duels       PSDuels
	Dribbles    PSDribbles
	Fouls       PSFouls
	Cards       PSCards
	Penalty     PSPenalty
	Position    string
}

type LeaguePlayerInfo struct {
	ApiFootballId int
}

type PlayerInfo struct {
	ApiFootballId int
	Name          string
	Firstname     string
	Lastname      string
	Age           int
	Height        string
	Weight        string
	Injured       bool
	Photo         string
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
	if input.Statistics.Games.Appearances != 0 {
		update.SetGameAppearances(input.Statistics.Games.Appearances)
	}
	if input.Statistics.Games.Lineups != 0 {
		update.SetGameLineups(input.Statistics.Games.Lineups)
	}
	if input.Statistics.Games.Minutes != 0 {
		update.SetGameMinutes(input.Statistics.Games.Minutes)
	}
	if input.Statistics.Shots.Total != 0 {
		update.SetTotalShots(input.Statistics.Shots.Total)
	}
	if input.Statistics.Shots.On != 0 {
		update.SetShotsOnTarget(input.Statistics.Shots.On)
	}
	if input.Statistics.Goals.Total != 0 {
		update.SetTotalGoals(input.Statistics.Goals.Total)
	}
	if input.Statistics.Goals.Conceded != 0 {
		update.SetGoalsConceded(input.Statistics.Goals.Conceded)
	}
	if input.Statistics.Goals.Assists != 0 {
		update.SetAssists(input.Statistics.Goals.Assists)
	}
	if input.Statistics.Goals.Saves != 0 {
		update.SetSaves(input.Statistics.Goals.Saves)
	}
	if input.Statistics.Position != "" {
		update.SetPosition(input.Statistics.Position)
	}
	if input.Statistics.Games.Rating != "" {
		update.SetRating(input.Statistics.Games.Rating)
	}
	if input.Statistics.Team.Name != "" {
		update.SetTeamName(input.Statistics.Team.Name)
	}
	if input.Statistics.Team.ApiFootballId != 0 {
		update.SetTeamID(input.Statistics.Team.ApiFootballId)
	}
	if input.Statistics.League.ApiFootballId != 0 {
		update.SetLeagueID(input.Statistics.League.ApiFootballId)
	}
	if input.Statistics.League.Name != "" {
		update.SetLeagueName(input.Statistics.League.Name)
	}
	if input.Statistics.Passes.Total != 0 {
		update.SetPassTotal(input.Statistics.Passes.Total)
	}
	if input.Statistics.Passes.Accuracy != 0 {
		update.SetPassAccuracy(input.Statistics.Passes.Accuracy)
	}
	if input.Statistics.Passes.Key != 0 {
		update.SetPassKey(input.Statistics.Passes.Key)
	}
	if input.Statistics.Tackles.Total != 0 {
		update.SetTotalTackle(input.Statistics.Tackles.Total)
	}
	if input.Statistics.Tackles.Blocks != 0 {
		update.SetBlocks(input.Statistics.Tackles.Blocks)
	}
	if input.Statistics.Tackles.Interceptions != 0 {
		update.SetInterceptions(input.Statistics.Tackles.Interceptions)
	}
	if input.Statistics.Duels.Total != 0 {
		update.SetDuelsTotal(input.Statistics.Duels.Total)
	}
	if input.Statistics.Duels.Won != 0 {
		update.SetDuelsWon(input.Statistics.Duels.Won)
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
func (m *PlayerModel) Exists(ctx context.Context, apiFootballId int) bool {
	count, _ := m.client.Player.
		Query().
		Where(player.ApiFootballIdEQ(apiFootballId)).
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
		_ = tx.Rollback() // rollback the transaction in case of errors
		return err
	}
	return tx.Commit() // commit the transaction
}
