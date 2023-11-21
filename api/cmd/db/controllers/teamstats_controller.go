package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/team_models/team_stats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type APITeamStatsResponse struct {
	Response TeamStatsResponse `json:"response"`
}

type TeamStatsResponse struct {
	Form          string                   `json:"form"`
	Fixtures      team_stats.Fixtures      `json:"fixtures"`
	Goals         team_stats.Goals         `json:"goals"`
	Biggest       team_stats.Biggest       `json:"biggest"`
	CleanSheet    team_stats.CleanSheet    `json:"clean_sheet"`
	FailedToScore team_stats.FailedToScore `json:"failed_to_score"`
	Penalty       team_stats.Penalty       `json:"penalty"`
	Lineups       []team_stats.Lineup      `json:"lineups"`
	Cards         team_stats.Cards         `json:"cards"`
}

type TeamStatsController struct {
	httpClient     *http.Client
	teamStatsModel *team_stats.TeamStatsModel
}

func NewTeamStatsController(entClient *ent.Client) *TeamStatsController {
	tsm := team_stats.NewTeamStatsModel(entClient)
	return &TeamStatsController{
		httpClient:     http.DefaultClient,
		teamStatsModel: tsm,
	}
}

func (tsc *TeamStatsController) InitializeFixtures(teams []*ent.Team, ctx context.Context) error {
	fmt.Println("Initializing team stats...")
	for _, team := range teams {
		err := tsc.FetchTeamStatsByID(ctx, team)
		if err != nil {
			return err
		}
	}

	fmt.Println("[ Finished initializing team stats ]")
	return nil
}

func (tsc *TeamStatsController) FetchTeamStatsByID(ctx context.Context, team *ent.Team) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams/statistics?league=%d&season=%d&team=%d",
		team.Edges.Season.Edges.League.FootballApiId,
		team.Edges.Season.Year,
		team.Edges.Club.ApiFootballId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := tsc.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return tsc.parseTeamStatsResponse(data, team)
}

func (tsc *TeamStatsController) parseTeamStatsResponse(data []byte, team *ent.Team) error {
	var teamStatsResponse APITeamStatsResponse
	err := json.Unmarshal(data, &teamStatsResponse)
	if err != nil {
		return err
	}

	teamStats := team_stats.TeamStats{
		Form:          teamStatsResponse.Response.Form,
		Fixtures:      teamStatsResponse.Response.Fixtures,
		Goals:         teamStatsResponse.Response.Goals,
		Biggest:       teamStatsResponse.Response.Biggest,
		CleanSheet:    teamStatsResponse.Response.CleanSheet,
		FailedToScore: teamStatsResponse.Response.FailedToScore,
		Penalty:       teamStatsResponse.Response.Penalty,
		Lineups:       teamStatsResponse.Response.Lineups,
		Cards:         teamStatsResponse.Response.Cards,
	}

	return tsc.upsertTeamStats(team, &teamStats)
}

func (tsc *TeamStatsController) upsertTeamStats(team *ent.Team, teamStats *team_stats.TeamStats) error {
	_, err := tsc.teamStatsModel.UpsertTeamStats(context.Background(), team, *teamStats)
	if err == nil {
		fmt.Printf("Team stats for %s successfully upserted.\n", team.Edges.Club.Name)
	} else {
		fmt.Printf("Error upserting team stats for %s: %v\n", team.Edges.Club.Name, err)
	}
	return err
}
