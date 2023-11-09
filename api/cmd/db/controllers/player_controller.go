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
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type PlayerStatistics struct {
	Team struct {
		League struct {
			ID int `json:"id"`
		} `json:"league"`
	} `json:"team"`
	Season int `json:"season"`
}

type LeaguePlayerInfo struct {
	ID int `json:"id"`
}

type LeaguePlayer struct {
	Player LeaguePlayerInfo `json:"player"`
}

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
type LeaguePlayerResponse struct {
	Results  int `json:"results"`
	Response []struct {
		Player struct {
			ID int `json:"id"`
		} `json:"player"`
		Statistics []player_stats.PlayerStats `json:"statistics"`
	} `json:"response"`
}

type Player struct {
	Player PlayerInfo `json:"player"`
}

type PlayerResponse struct {
	Results  int      `json:"results"`
	Response []Player `json:"response"`
}

type PlayerController struct {
	client      *http.Client
	playerModel *player_models.PlayerModel
}

func NewPlayerController(playerModel *player_models.PlayerModel) *PlayerController {
	return &PlayerController{
		client:      &http.Client{},
		playerModel: playerModel,
	}
}

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

func (pc *PlayerController) fetchPlayerByID(ctx context.Context, playerID int) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?id=%d&season=2022", playerID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	fmt.Printf("Fetching player by ID: %d\n", playerID)
	resp, err := pc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received non-OK HTTP status %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return pc.parsePlayerResponse(ctx, data, playerID)
}

func (pc *PlayerController) parsePlayerResponse(ctx context.Context, data []byte, playerID int) error {
	var response PlayerResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	// Check if the response is empty
	if response.Results == 0 {
		// Log the scenario and return without an error
		fmt.Printf("No data available for playerID: %d for season 2022.\n", playerID)
		return nil
	}

	p := response.Response[0].Player

	if err := pc.downloadPhotoIfNeeded(&p); err != nil {
		return err
	}

	input := player_models.CreatePlayerInput{
		ApiFootballId: p.ID,
		Name:          p.Name,
		Firstname:     p.Firstname,
		Lastname:      p.Lastname,
		Age:           p.Age,
		Height:        p.Height,
		Weight:        p.Weight,
		Injured:       p.Injured,
		Photo:         p.Photo,
	}

	err := pc.upsertPlayer(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

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

func (pc *PlayerController) GetPlayerIDsForLeague(ctx context.Context, leagueID int) ([]int, error) {
	// Create a URL for fetching players by their league ID
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?league=%d&season=2022", leagueID)
	// Create an HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers for the API
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal data to get player IDs
	var leaguePlayerResponse LeaguePlayerResponse
	if err := json.Unmarshal(body, &leaguePlayerResponse); err != nil {
		return nil, err
	}

	var playerIDs []int
	for _, p := range leaguePlayerResponse.Response {
		playerIDs = append(playerIDs, p.Player.ID)
	}

	return playerIDs, nil
}

func (pc *PlayerController) FetchPlayersByLeague(ctx context.Context, leagueID int) error {
	playerIDs, err := pc.GetPlayerIDsForLeague(ctx, leagueID)
	fmt.Println("Initializing Player... IN CONTROLLER")
	fmt.Println(playerIDs)
	if err != nil {
		return err
	}

	// Ensure each player exists in the database
	for _, playerID := range playerIDs {
		err := pc.EnsurePlayerExists(ctx, playerID)
		if err != nil {
			return err
		}
	}

	return nil
}

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
