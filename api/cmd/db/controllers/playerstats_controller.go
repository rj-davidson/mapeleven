package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models/player_stats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type APIPlayerStatsResponse struct {
	Response PlayerStatsResponse `json:"response"`
}

type PlayerStatsResponse struct {
	Games   player_stats.PSGames   `json:"games"`
	Goals   player_stats.PSGoals   `json:"goals"`
	Offense player_stats.PSOffense `json:"offense"`
	Defense player_stats.PSDefense `json:"defense"`
	Penalty player_stats.PSPenalty `json:"penalty"`
}

type PlayerStatsController struct {
	client           *http.Client
	playerStatsModel *player_stats.PlayerStatsModel
}

func NewPlayerStatsController(playerStatsModel *player_stats.PlayerStatsModel) *PlayerStatsController {
	return &PlayerStatsController{
		client:           &http.Client{},
		playerStatsModel: playerStatsModel,
	}
}

func (psc *PlayerStatsController) InitializeStats(players []*ent.Player, ctx context.Context) error {
	fmt.Println("Initializing player stats...")
	for _, player := range players {
		err := psc.FetchPlayerStatsByID(ctx, player)
		if err != nil {
			return err
		}
	}

	fmt.Println("[ Finished initializing player stats ]")
	return nil
}

func (psc *PlayerStatsController) FetchPlayerStatsByID(ctx context.Context, player *ent.Player) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players/statistics?league=%d&season=%d&team=%d&player=%d",
		// Replace the placeholders with appropriate values
		player.Edges.Season.Edges.League.FootballApiId,
		player.Edges.Season.Year,
		player.Edges.Team.Edges.Club.ApiFootballId,
		player.ApiFootballId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := psc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return psc.parsePlayerStatsResponse(data, player)
}

func (psc *PlayerStatsController) parsePlayerStatsResponse(data []byte, player *ent.Player) error {
	var playerStatsResponse APIPlayerStatsResponse
	err := json.Unmarshal(data, &playerStatsResponse)
	if err != nil {
		return err
	}

	playerStats := player_stats.PlayerStats{
		PSGames:   playerStatsResponse.Response.Games,
		PSGoals:   playerStatsResponse.Response.Goals,
		PSOffense: playerStatsResponse.Response.Offense,
		PSDefense: playerStatsResponse.Response.Defense,
		PSPenalty: playerStatsResponse.Response.Penalty,
	}

	return psc.upsertPlayerStats(player, &playerStats)
}

func (psc *PlayerStatsController) upsertPlayerStats(player *ent.Player, playerStats *player_stats.PlayerStats) error {
	_, err := psc.playerStatsModel.UpsertPlayerStats(context.Background(), player, *playerStats)
	if err == nil {
		fmt.Printf("Player stats for %s successfully upserted.\n", player.Name)
	} else {
		fmt.Printf("Error upserting player stats for %s: %v\n", player.Name, err)
	}
	return err
}
