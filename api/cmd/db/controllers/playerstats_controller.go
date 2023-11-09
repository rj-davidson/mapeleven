package controllers

//
//import (
//	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models/player_stats"
//	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/team_models"
//	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
//	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/spf13/viper"
//	"io/ioutil"
//	"net/http"
//)
//
//type APIPlayerStatsResponse struct {
//	Response PlayerStatsResponse `json:"response"`
//}
//
//type PlayerStatsResponse struct {
//	Slug        string                 `json:"slug"`
//	LastUpdated int                    `json:"lastUpdated"`
//	Season      string                 `json:"season"`
//	Games       player_stats.PSGames   `json:"games"`
//	Goals       player_stats.PSGoals   `json:"goals"`
//	Offense     player_stats.PSOffense `json:"offense"`
//	Defense     player_stats.PSDefense `json:"defense"`
//	Penalty     player_stats.PSPenalty `json:"penalty"`
//}
//
//type PlayerStatsController struct {
//	client *http.Client
//	tm     *team_models.TeamModel
//}
//
//func NewPlayerStatsController(tm *team_models.TeamModel) *PlayerStatsController {
//	if tm == nil {
//		panic("teamModel is nil")
//	}
//
//	return &PlayerStatsController{
//		client: &http.Client{},
//		tm:     tm,
//	}
//}
//
//func (psc *PlayerStatsController) InitializeStats(ctx context.Context) error {
//	if psc == nil {
//		return fmt.Errorf("PlayerStatsController is nil")
//	}
//
//	fmt.Println("Initializing player stats...")
//	teams, err := psc.tm.ListAll(ctx) // This method should list all teams
//	if err != nil {
//		return err
//	}
//	for _, player := range teams {
//		// Check if the PlayerStats edge is loaded
//		if player.Edges.PlayerStats == nil {
//			return fmt.Errorf("player stats edge not loaded for player ID %d", player.ID)
//		}
//		// Iterate over all PlayerStats records for the player
//		for _, ps := range player.Edges.PlayerStats {
//			if err := psc.FetchPlayerStatsByID(ctx, ps); err != nil {
//				return err
//			}
//		}
//	}
//
//	fmt.Println("[ Finished initializing player stats ]")
//	return nil
//}
//
//func (psc *PlayerStatsController) FetchPlayerStatsByID(ctx context.Context, ps *ent.Player) error {
//	if psc == nil {
//		return fmt.Errorf("PlayerStatsController is nil")
//	}
//
//	// Check for nil edges and preload them if necessary
//	if ps.Edges.Player == nil {
//		return fmt.Errorf("player edge is not loaded for PlayerStats ID %d", ps.ID)
//	}
//	if ps.Edges.Team == nil || ps.Edges.Team.Edges.Season == nil {
//		return fmt.Errorf("team or season edge is not loaded for PlayerStats ID %d", ps.ID)
//	}
//
//	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?id=%d&season=%d",
//		ps.Edges.Player.ApiFootballId,
//		ps.Edges.Team.Edges.Season.Year)
//
//	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
//	if err != nil {
//		return err
//	}
//	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
//	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))
//
//	resp, err := psc.client.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	// Check the response status code
//	if resp.StatusCode != http.StatusOK {
//		return fmt.Errorf("bad status: %s", resp.Status)
//	}
//
//	data, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return err
//	}
//
//	return psc.parsePlayerStatsResponse(ctx, data, ps)
//}
//
//func (psc *PlayerStatsController) parsePlayerStatsResponse(ctx context.Context, data []byte, player *ent.PlayerStats) error {
//	fmt.Printf("Raw JSON response: %s\n", string(data))
//	var playerStatsResponse APIPlayerStatsResponse
//	err := json.Unmarshal(data, &playerStatsResponse)
//	if err != nil {
//		return err
//	}
//
//	playerStats := player_stats.PlayerStats{
//		PSGames:   playerStatsResponse.Response.Games,
//		PSGoals:   playerStatsResponse.Response.Goals,
//		PSOffense: playerStatsResponse.Response.Offense,
//		PSDefense: playerStatsResponse.Response.Defense,
//		PSPenalty: playerStatsResponse.Response.Penalty,
//	}
//
//	return psc.upsertPlayerStats(ctx, player, &playerStats)
//}
//
//func (psc *PlayerStatsController) upsertPlayerStats(ctx context.Context, player *ent.PlayerStats, playerStats *player_stats.PlayerStats) error {
//	_, err := psc.playerStatsModel.UpsertPlayerStats(ctx, player, *playerStats)
//	if err != nil {
//		return fmt.Errorf("error upserting player stats for %s: %v", player.Edges.Player.Name, err)
//	}
//
//	fmt.Printf("Player stats for %s successfully upserted.\n", player.Edges.Player.Name)
//	return nil
//}
