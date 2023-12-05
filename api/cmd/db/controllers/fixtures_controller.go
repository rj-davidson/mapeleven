package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/fixture_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/team_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/utils"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type APIFixtureResponse struct {
	Response []struct {
		Fixture struct {
			ID       int    `json:"id"`
			Referee  string `json:"referee"`
			Timezone string `json:"timezone"`
			Date     string `json:"date"`
			Status   struct {
				Long    string `json:"long"`
				Elapsed int    `json:"elapsed"`
			} `json:"status"`
		} `json:"fixture"`
		League struct {
			ID    int    `json:"id"`
			Round string `json:"round"`
		} `json:"league"`
		Teams struct {
			Home struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"home"`
			Away struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"away"`
		} `json:"teams"`
		Score struct {
			Fulltime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"fulltime"`
			Halftime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"halftime"`
			Extratime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"extratime"`
			Penalty struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"penalty"`
		} `json:"score"`
	} `json:"response"`
}

type FixtureController struct {
	httpClient   *http.Client
	fixtureModel *fixture_models.FixtureModel
	leagueModel  *models.LeagueModel
	teamModel    *team_models.TeamModel
	teamCtrl     *TeamController
}

func NewFixtureController(entClient *ent.Client) *FixtureController {
	fm := fixture_models.NewFixtureModel(entClient)
	lm := models.NewLeagueModel(entClient)
	tm := team_models.NewTeamModel(entClient)
	tc := NewTeamController(entClient)
	return &FixtureController{
		httpClient:   http.DefaultClient,
		fixtureModel: fm,
		leagueModel:  lm,
		teamModel:    tm,
		teamCtrl:     tc,
	}
}

func (fc *FixtureController) InitializeFixtures(ctx context.Context) error {
	fmt.Println("Initializing fixtures...")
	// Get all leagues from the database
	leagues, err := fc.leagueModel.ListLeagues(ctx)
	if err != nil {
		return err
	}

	// Fetch and process fixtures for each league
	for _, league := range leagues {
		// Get the current season for the league
		season, err := fc.leagueModel.GetCurrentSeasonForLeague(ctx, league)
		fmt.Println("Fetching fixtures for league: ", league.Name)
		_, err = fc.fetchFixturesByLeague(ctx, league.FootballApiId, season)
		if err != nil {
			log.Printf("Error fetching fixtures for league %d: %v", league.ID, err)
			continue
		}
	}

	fmt.Println("[ Finished initializing fixtures ]")
	return nil
}

// RefreshFixtures fetches fixtures for all leagues in the database
func (fc *FixtureController) RefreshFixtures(ctx context.Context) {
	log.Println("Refreshing fixtures...")

	// Get all leagues from the database
	err := fc.InitializeFixtures(ctx)
	if err != nil {
		log.Printf("Error refreshing fixtures: %v", err)
	}

	log.Println("[ Finished refreshing fixtures ]")
}

// UpdateFixtures updates fixtures given a list of *ent.Fixture
func (fc *FixtureController) UpdateFixtures(fixtures []*ent.Fixture) error {
	ctx := context.Background()
	for _, f := range fixtures {
		_, err := fc.fetchFixtureById(ctx, f.ApiFootballId, f.Edges.Season)
		if err != nil {
			fmt.Printf("Error fetching fixture by id %d: %v", f.ApiFootballId, err)
		}
	}

	return nil
}

func (fc *FixtureController) fetchFixturesByLeague(ctx context.Context, leagueID int, season *ent.Season) ([]*ent.Fixture, error) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?league=%d&season=%d", leagueID, season.Year)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := fc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return fc.parseFixturesResponse(data, season)
}

func (fc *FixtureController) fetchFixturesByIds(ctx context.Context, fixtureIDs []int, season *ent.Season) ([]*ent.Fixture, error) {
	var fixtures []*ent.Fixture
	for _, id := range fixtureIDs {
		url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?id=%d", id)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

		resp, err := fc.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		res, err := fc.parseFixturesResponse(data, season)
		if err != nil {
			return nil, err
		}
		fixtures = append(fixtures, res...)
	}

	return fixtures, nil
}

func (fc *FixtureController) fetchFixtureById(ctx context.Context, fixtureID int, season *ent.Season) ([]*ent.Fixture, error) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?id=%d", fixtureID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := fc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return fc.parseFixturesResponse(data, season)
}

func (fc *FixtureController) parseFixturesResponse(data []byte, season *ent.Season) ([]*ent.Fixture, error) {
	var response APIFixtureResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	var fixtures []*ent.Fixture
	for _, res := range response.Response {
		parsedDate, err := time.Parse(time.RFC3339, res.Fixture.Date)
		if err != nil {
			return nil, err
		}

		var homeScore *int
		var awayScore *int

		if res.Fixture.Status.Long == "Match Finished" || res.Fixture.Status.Long == "Full Time" {
			homeScore = &res.Score.Fulltime.Home
			awayScore = &res.Score.Fulltime.Away
		} else {
			homeScore = nil
			awayScore = nil
		}

		fixtureInput := fixture_models.CreateFixtureInput{
			ApiFootballId: res.Fixture.ID,
			Referee:       &res.Fixture.Referee,
			Timezone:      &res.Fixture.Timezone,
			Date:          parsedDate,
			Elapsed:       &res.Fixture.Status.Elapsed,
			Status:        res.Fixture.Status.Long,
			HomeTeamScore: homeScore,
			AwayTeamScore: awayScore,
			Season:        season,
			HomeTeamID:    res.Teams.Home.ID,
			AwayTeamID:    res.Teams.Away.ID,
			Round:         nil,
			Slug:          utils.Slugify(res.Teams.Home.Name + "-" + parsedDate.Format("2006-01-02") + "-" + res.Teams.Away.Name),
		}

		fixture, err := fc.upsertFixture(context.Background(), &fixtureInput)
		if err != nil {
			fmt.Println("Error upserting fixture: ", err)
		}
		fixtures = append(fixtures, fixture)
	}

	return fixtures, nil
}

func (fc *FixtureController) upsertFixture(ctx context.Context, fixtureInput *fixture_models.CreateFixtureInput) (*ent.Fixture, error) {
	// Check if the fixture exists, and either create or update it accordingly
	existingFixture := fc.fixtureModel.FixtureExistsByApiFootballId(ctx, fixtureInput.ApiFootballId)

	// If the fixture does not exist, we create it
	if !existingFixture {
		return fc.fixtureModel.CreateBaseFixture(ctx, *fixtureInput)
	} else {
		fc.handleFixtureTeams(ctx, fixtureInput)
		updatedFixtureInput := fixture_models.UpdateFixtureInput{
			ApiFootballId: fixtureInput.ApiFootballId,
			Referee:       fixtureInput.Referee,
			Timezone:      fixtureInput.Timezone,
			Date:          fixtureInput.Date,
			Elapsed:       fixtureInput.Elapsed,
			Round:         fixtureInput.Round,
			Status:        &fixtureInput.Status,
			HomeTeamScore: fixtureInput.HomeTeamScore,
			AwayTeamScore: fixtureInput.AwayTeamScore,
			HomeTeamID:    &fixtureInput.HomeTeamID,
			AwayTeamID:    &fixtureInput.AwayTeamID,
		}
		return fc.fixtureModel.UpdateBaseFixture(ctx, updatedFixtureInput)
	}
}

func (fc *FixtureController) handleFixtureTeams(ctx context.Context, input *fixture_models.CreateFixtureInput) {
	// Ensure the home team exists
	err := fc.teamCtrl.EnsureTeamExists(ctx, input.Season.Year, input.Season.Edges.League.FootballApiId, input.HomeTeamID)
	if err != nil {
		fmt.Printf("[Fixture Controller] Error ensuring team exists: %v\n", err)
	}

	// Ensure the away team exists
	err = fc.teamCtrl.EnsureTeamExists(ctx, input.Season.Year, input.Season.Edges.League.FootballApiId, input.AwayTeamID)
	if err != nil {
		fmt.Printf("[Fixture Controller] Error ensuring team exists: %v\n", err)
	}
}
