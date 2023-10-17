package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"mapeleven/db/ent"
	"mapeleven/models"
	"mapeleven/models/fixture_models"
	"mapeleven/utils"
	"net/http"
	"time"
)

type FixtureController struct {
	client       *http.Client
	fixtureModel *fixture_models.FixtureModel
}

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
		} `json:"score"`
	} `json:"response"`
}

func NewFixtureController(fixtureModel *fixture_models.FixtureModel) *FixtureController {
	return &FixtureController{
		client:       &http.Client{},
		fixtureModel: fixtureModel,
	}
}

func (fc *FixtureController) InitializeFixtures(lm *models.LeagueModel, ctx context.Context) error {
	fmt.Println("Initializing fixtures...")
	// Get all leagues from the database
	leagues, err := lm.ListLeagues(ctx)
	if err != nil {
		return err
	}

	// Fetch and process fixtures for each league
	for _, league := range leagues {
		// Get the current season for the league
		season, err := lm.GetCurrentSeasonForLeague(ctx, league)
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

func (fc *FixtureController) fetchFixturesByLeague(ctx context.Context, leagueID int, season *ent.Season) ([]*ent.Fixture, error) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?league=%d&season=%d", leagueID, season.Year)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := fc.client.Do(req)
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

		resp, err := fc.client.Do(req)
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

		fixtureInput := fixture_models.CreateFixtureInput{
			ApiFootballId: res.Fixture.ID,
			Referee:       &res.Fixture.Referee,
			Timezone:      &res.Fixture.Timezone,
			Date:          parsedDate,
			Elapsed:       &res.Fixture.Status.Elapsed,
			Status:        res.Fixture.Status.Long,
			HomeTeamScore: &res.Score.Fulltime.Home,
			AwayTeamScore: &res.Score.Fulltime.Away,
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
		// If the fixture exists, we update it
		updatedFixtureInput := fixture_models.UpdateFixtureInput{
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
