package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
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
	httpClient     *http.Client
	standingsModel *models.StandingsModel
	leagueModel    *models.LeagueModel
	clubController *ClubController
	teamController *TeamController
}

func NewStandingsController(entClient *ent.Client) *StandingsController {
	sm := models.NewStandingsModel(entClient)
	cc := NewClubController(entClient)
	lm := models.NewLeagueModel(entClient)
	tc := NewTeamController(entClient)
	return &StandingsController{
		httpClient:     http.DefaultClient,
		leagueModel:    lm,
		standingsModel: sm,
		clubController: cc,
		teamController: tc,
	}
}

func (sc *StandingsController) InitializeStandings(ctx context.Context) error {
	fmt.Println("Initializing standings...")

	leagues, err := sc.leagueModel.ListLeagues(ctx)
	if err != nil {
		return err
	}

	// Fetch and process standings for each league
	for _, l := range leagues {
		s, err := sc.leagueModel.GetCurrentSeasonForLeague(ctx, l)
		fmt.Println("Fetching standings for league: ", l.Name)
		_, err = sc.FetchStandingsByLeague(ctx, l, s)
		if err != nil {
			fmt.Printf("fetch standings for league %s: %w", league.Name, err)
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

	resp, err := sc.httpClient.Do(req)
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
				standingsInput := models.StandingsForm{
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

func (sc *StandingsController) upsertStandings(ctx context.Context, standingsInput *models.StandingsForm, year, apiFootballLeagueId, apiFootballClubId int) (*ent.Standings, error) {
	standingExists, standing := sc.standingsModel.Exists(ctx, standingsInput.Team, standingsInput.SeasonID)
	t, err := sc.teamController.GetTeam(ctx, year, apiFootballLeagueId, apiFootballClubId)

	if !standingExists {
		if err != nil {
			return nil, err
		}
		return sc.standingsModel.CreateStandings(ctx, *standingsInput, t)
	} else {
		return sc.standingsModel.UpdateStandings(ctx, *standingsInput, standing.ID)

	}
}
