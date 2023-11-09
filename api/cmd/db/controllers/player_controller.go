package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models/player_stats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/utils"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// PlayerInfo holds basic information about a player.
type PlayerInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	Injured   bool   `json:"injured"`
	Photo     string `json:"photo"`
}

// Statistics wraps all statistics related to a player.
type Statistics struct {
	Team        player_stats.PSTeam        `json:"team"`
	League      player_stats.PSLeague      `json:"league"`
	Games       player_stats.PSGames       `json:"games"`
	Substitutes player_stats.PSSubstitutes `json:"substitutes"`
	Shots       player_stats.PSShots       `json:"shots"`
	Goals       player_stats.PSGoals       `json:"goals"`
	Passes      player_stats.PSPasses      `json:"passes"`
	Tackles     player_stats.PSTackles     `json:"tackles"`
	Duels       player_stats.PSDuels       `json:"duels"`
	Dribbles    player_stats.PSDribbles    `json:"dribbles"`
	Fouls       player_stats.PSFouls       `json:"fouls"`
	Cards       player_stats.PSCards       `json:"cards"`
	Penalty     player_stats.PSPenalty     `json:"penalty"`
}

// Player wraps PlayerInfo and includes a slice of Statistics.
type Player struct {
	Player     PlayerInfo   `json:"player"`
	Statistics []Statistics `json:"statistics"`
}

// PlayerResponse is the expected structure of the JSON response.
type PlayerResponse struct {
	Results  int      `json:"results"`
	Response []Player `json:"response"`
}

// PlayerController manages HTTP requests for player data.
type PlayerController struct {
	client      *http.Client
	playerModel *player_models.PlayerModel
}

// NewPlayerController creates a new PlayerController.
func NewPlayerController(playerModel *player_models.PlayerModel) *PlayerController {
	return &PlayerController{
		client:      &http.Client{},
		playerModel: playerModel,
	}
}

// EnsurePlayerExists checks if a player exists and fetches their data if not.
func (pc *PlayerController) EnsurePlayerExists(ctx context.Context, apiFootballId int) error {
	exists := pc.playerModel.Exists(ctx, apiFootballId)
	if !exists {
		err := pc.fetchPlayerByID(ctx, apiFootballId)
		if err != nil {
			return err
		}
	}
	return nil
}

// fetchPlayerByID makes an HTTP request to fetch player data by ID.
func (pc *PlayerController) fetchPlayerByID(ctx context.Context, playerID int) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?id=%d&season=2022", playerID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	resp, err := pc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return pc.parsePlayerResponse(ctx, data, playerID)
}

// parsePlayerResponse parses the JSON response into the PlayerResponse struct.
func (pc *PlayerController) parsePlayerResponse(ctx context.Context, data []byte, playerID int) error {
	var response PlayerResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	if response.Results == 0 {
		fmt.Printf("No data available for playerID: %d for season 2022.\n", playerID)
		return nil
	}

	playerData := response.Response[0]

	if err := pc.downloadPhotoIfNeeded(&playerData.Player); err != nil {
		return err
	}

	input := player_models.CreatePlayerInput{
		ApiFootballId: playerData.Player.ID,
		Name:          playerData.Player.Name,
		Firstname:     playerData.Player.Firstname,
		Lastname:      playerData.Player.Lastname,
		Age:           playerData.Player.Age,
		Height:        playerData.Player.Height,
		Weight:        playerData.Player.Weight,
		Injured:       playerData.Player.Injured,
		Photo:         playerData.Player.Photo,
	}

	err := pc.upsertPlayer(ctx, input)
	if err != nil {
		return err
	}

	// Here you would handle the playerData.Statistics which is a slice of Statistics.
	// This could involve iterating over the slice and updating/inserting each statistic in your database.
	// This part of the logic is omitted for brevity, but it would follow a similar pattern to the upsertPlayer function.

	return nil
}

// upsertPlayer inserts or updates a player in the database.
func (pc *PlayerController) upsertPlayer(ctx context.Context, input player_models.CreatePlayerInput) error {
	return pc.playerModel.WithTransaction(ctx, func(tx *ent.Tx) error {
		exists := tx.Player.Query().Where(player.ApiFootballIdEQ(input.ApiFootballId)).ExistX(ctx)
		if exists {
			updateInput := player_models.UpdatePlayerInput{
				ApiFootballId: input.ApiFootballId,
				Name:          &input.Name,
				Firstname:     &input.Firstname,
				Lastname:      &input.Lastname,
				Age:           &input.Age,
				Height:        &input.Height,
				Weight:        &input.Weight,
				Injured:       &input.Injured,
				Photo:         &input.Photo,
			}
			_, err := pc.playerModel.UpdatePlayer(ctx, updateInput)
			if err != nil {
				return err
			}
		} else {
			_, err := pc.playerModel.CreatePlayer(ctx, input)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// downloadPhotoIfNeeded downloads a player's photo if it is not already present.
func (pc *PlayerController) downloadPhotoIfNeeded(p *PlayerInfo) error {
	if p.Photo != "" {
		photoLocation, err := utils.DownloadImageIfNeeded(
			p.Photo,
			fmt.Sprintf("images/players/%d%s", p.ID, filepath.Ext(p.Photo)),
		)
		if err != nil {
			return err
		}
		p.Photo = photoLocation
	}
	return nil
}
