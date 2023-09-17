package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"mapeleven/db/ent"
	"mapeleven/db/ent/league"
	"mapeleven/models"
	"net/http"
)

type APIStandingsResponse struct {
	Response []struct {
		League struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Country   string `json:"country"`
			Logo      string `json:"logo"`
			Flag      string `json:"flag"`
			Season    int    `json:"season"`
			Standings [][]struct {
				Rank int `json:"rank"`
				Team struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Logo string `json:"logo"`
				} `json:"team"`
				Points      int    `json:"points"`
				GoalsDiff   int    `json:"goalsDiff"`
				Group       string `json:"group"`
				Form        string `json:"form"`
				Status      string `json:"status"`
				Description string `json:"description"`
				All         struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"all"`
				Home struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"home"`
				Away struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"away"`
				Update string `json:"update"`
			} `json:"standings"`
		} `json:"league"`
	} `json:"response"`
}

type StandingsController struct {
	client         *http.Client
	standingsModel *models.StandingsModel
	clubController *ClubController
	teamController *TeamController
}

func NewStandingsController(standingsModel *models.StandingsModel, clubController *ClubController, teamController *TeamController) *StandingsController {
	return &StandingsController{
		client:         &http.Client{},
		standingsModel: standingsModel,
		clubController: clubController,
		teamController: teamController,
	}
}

func (sc *StandingsController) InitializeStandings(lm *models.LeagueModel, ctx context.Context) error {
	fmt.Println("Initializing standings...")

	leagues, err := lm.ListLeagues(ctx)
	if err != nil {
		return err
	}

	// Fetch and process standings for each league
	for _, l := range leagues {
		s, err := lm.GetCurrentSeasonForLeague(ctx, l)
		fmt.Println("Fetching standings for league: ", l.Name)
		_, err = sc.FetchStandingsByLeague(ctx, l, s)
		if err != nil {
			fmt.Errorf("fetch standings for league %s: %w", league.Name, err)
		}
	}

	fmt.Println("[ Finished initializing standings ]")
	return nil
}

func (sc *StandingsController) FetchStandingsByLeague(ctx context.Context, l *ent.League, s *ent.Season) ([]*ent.Standings, error) {
	url := fmt.Sprintf("https://api-football-v3.p.rapidapi.com/v3/standings?season=%d&league=%d", s.Year, l.FootballApiId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := sc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return sc.parseStandingsResponse(data, s)
}

func (sc *StandingsController) parseStandingsResponse(data []byte, s *ent.Season) ([]*ent.Standings, error) {
	var response APIStandingsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	// TODO: Add Club Controller to fetch club data

	var leagueStandings []*ent.Standings
	for _, res := range response.Response {
		for _, standings := range res.League.Standings {
			for _, standing := range standings {
				err := sc.clubController.EnsureClubExists(context.Background(), standing.Team.ID)
				if err != nil {
					fmt.Println("Error ensuring club exists: ", err)
				}
				err = sc.teamController.EnsureTeamExists(context.Background(), s.Year, res.League.ID, standing.Team.ID)
				if err != nil {
					fmt.Println("Error ensuring team exists: ", err)
				}
				standingsInput := models.CreateStandingsInput{
					Rank:             standing.Rank,
					Description:      standing.Description,
					SeasonID:         s.ID,
					Team:             standing.Team.ID,
					Points:           standing.Points,
					GoalsDiff:        standing.GoalsDiff,
					Group:            standing.Group,
					Form:             standing.Form,
					Status:           standing.Status,
					Played:           standing.All.Played,
					Win:              standing.All.Win,
					Draw:             standing.All.Draw,
					Lose:             standing.All.Lose,
					GoalsFor:         standing.All.Goals.For,
					GoalsAgainst:     standing.All.Goals.Against,
					HomePlayed:       standing.Home.Played,
					HomeWin:          standing.Home.Win,
					HomeDraw:         standing.Home.Draw,
					HomeLose:         standing.Home.Lose,
					HomeGoalsFor:     standing.Home.Goals.For,
					HomeGoalsAgainst: standing.Home.Goals.Against,
					AwayPlayed:       standing.Away.Played,
					AwayWin:          standing.Away.Win,
					AwayDraw:         standing.Away.Draw,
					AwayLose:         standing.Away.Lose,
					AwayGoalsFor:     standing.Away.Goals.For,
					AwayGoalsAgainst: standing.Away.Goals.Against,
				}

				standing, err := sc.upsertStandings(context.Background(), &standingsInput, s.Year, res.League.ID, standing.Team.ID)
				if err != nil {
					fmt.Println("Error upserting standings: ", standingsInput.Team, " - ", err)
				}
				leagueStandings = append(leagueStandings, standing)
			}
		}
	}
	return leagueStandings, nil
}

func (sc *StandingsController) upsertStandings(ctx context.Context, standingsInput *models.CreateStandingsInput, year, apiFootballLeagueId, apiFootballClubId int) (*ent.Standings, error) {
	existingStanding := sc.standingsModel.CheckIfStandingsExist(ctx, standingsInput.Team, standingsInput.SeasonID)

	if !existingStanding {
		t, err := sc.teamController.GetTeam(ctx, year, apiFootballLeagueId, apiFootballClubId)
		if err != nil {
			return nil, err
		}
		return sc.standingsModel.CreateStandings(ctx, *standingsInput, t)
	} else {
		// TODO Implement update

		// Return  nil and an error statying that update standings has not been implemented yet
		return nil, fmt.Errorf("update standings has not been implemented yet")

	}
}
