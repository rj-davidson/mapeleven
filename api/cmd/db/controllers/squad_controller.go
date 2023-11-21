package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models/player_stats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/team_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type SquadController struct {
	httpClient  *http.Client
	squadModel  *team_models.SquadModel
	playerModel *player_models.PlayerModel
	psModel     *player_stats.PlayerStatsModel
	playerCtrl  *PlayerController
}

func NewSquadController(entClient *ent.Client) *SquadController {
	sm := team_models.NewSquadModel(entClient)
	pm := player_models.NewPlayerModel(entClient)
	psm := player_stats.NewPlayerStatsModel(entClient)
	pc := NewPlayerController(entClient)
	return &SquadController{
		httpClient:  http.DefaultClient,
		squadModel:  sm,
		playerModel: pm,
		psModel:     psm,
		playerCtrl:  pc,
	}
}

func (sc *SquadController) InitializeSquads(teams []*ent.Team, ctx context.Context) error {
	fmt.Println("Initializing squads...")
	for _, team := range teams {
		err := sc.FetchSquadByTeam(ctx, team)
		if err != nil {
			return err
		}
	}

	fmt.Println("[ Finished initializing squads ]")
	return nil
}

func (sc *SquadController) FetchSquadByTeam(ctx context.Context, team *ent.Team) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players/squads?team=%d", team.Edges.Club.ApiFootballId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := sc.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return sc.parseSquadResponse(data, team)
}

func (sc *SquadController) parseSquadResponse(data []byte, team *ent.Team) error {
	var squadResponse team_models.Squad
	err := json.Unmarshal(data, &squadResponse)
	if err != nil {
		return err
	}

	for _, teamResponse := range squadResponse.Response {
		for _, player := range teamResponse.Players {
			err = sc.playerCtrl.EnsurePlayerExists(context.Background(), player.ApiFootballId)
			err = sc.upsertSquadPlayer(team, player)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (sc *SquadController) upsertSquadPlayer(team *ent.Team, s team_models.Player) error {
	_, err := sc.squadModel.UpsertSquadPlayer(team, s, context.Background())
	if err != nil {
		fmt.Printf("Error upserting squad for %s: %v\n", team.Edges.Club.Name, err)
		return err
	}
	fmt.Printf("Successfully upserted %s (Squad: %s)\n", s.Name, team.Edges.Club.Name)
	return nil
}
