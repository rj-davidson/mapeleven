package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"mapeleven/db/ent"
	"mapeleven/models"
	"net/http"
	"time"
)

var seasonModel *models.SeasonModel
var countryModel *models.CountryModel
var leagueModel *models.LeagueModel
var teamModel *models.TeamModel
var standingsModel *models.StandingsModel
var fixturesModel *models.FixtureModel

//var teamSeasonsModel *models.TeamSeasonsModel

func CronScheduler(client *ent.Client, initialize bool, runScheduler bool, devRun bool) {
	// Initialize scheduler
	s := gocron.NewScheduler(time.UTC)

	// Model Declaration
	seasonModel = models.NewSeasonModel(client)
	countryModel = models.NewCountryModel(client)
	leagueModel = models.NewLeagueModel(client)
	teamModel = models.NewTeamModel(client)
	standingsModel = models.NewStandingsModel(client)
	fixturesModel = models.NewFixtureModel(client)

	if initialize {
		// Country Initialization
		fetchCountries(countryModel)

		// SeasonID Initialization
		fetchLeagues(leagueModel, seasonModel)

		// Team Initialization
		fetchTeams(teamModel)

		// Standings Initialization
		fetchStandings(standingsModel)

		// Fixtures Initialization
		fetchFixtures(fixturesModel)
	}

	if runScheduler {
		// Schedule cron jobs
		s.Every(14).Days().At("00:00").Do(fetchCountries, countryModel)
		s.Every(14).Days().At("00:00").Do(fetchLeagues, leagueModel)
		s.Every(14).Days().At("00:00").Do(fetchTeams, teamModel)
		s.Every(14).Days().At("00:00").Do(fetchStandings, standingsModel)
		s.Every(14).Days().At("00:00").Do(fetchFixtures, fixturesModel)
		//s.Every(14).Days().At("00:00").Do(fetchTeamSeasons, teamSeasonsModel)
	}

	if devRun {

	}
}

func fetchCountries(countryModel *models.CountryModel) {
	countryController := NewCountryController(countryModel)
	err := countryController.InitializeCountries()
	if err != nil {
		log.Fatal(err)
	}
}

func fetchLeagues(leagueModel *models.LeagueModel, seasonModel *models.SeasonModel) {
	championsLeagueID := 2
	europaLeagueID := 3
	premierLeagueID := 39
	serieAID := 61
	ligue1ID := 71
	bundesligaID := 78
	laLigaID := 140
	majorleaguesoccerID := 253
	leagueIDs := []int{laLigaID, serieAID, bundesligaID, premierLeagueID, ligue1ID, majorleaguesoccerID, championsLeagueID, europaLeagueID}
	//leagueIDs := []int{majorleaguesoccerID, championsLeagueID, europaLeagueID}

	leagueController := NewLeagueController(leagueModel, seasonModel)
	err := leagueController.InitializeLeagues(leagueIDs, context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

// Fetches teams from API and saves them to the database
func fetchTeams(teamModel *models.TeamModel) {
	// Request Standings for each league
	// Initialize empty teams array of type int
	teamIDs := fetchTeamIDs()

	// Save teams to database
	teamController := NewTeamController(teamModel)
	err := teamController.InitializeTeams(teamIDs)
	if err != nil {
		log.Fatal(err)
	}
}

// Fetches team IDs from API
func fetchTeamIDs() []int {
	type StandingsResponse struct {
		Response []struct {
			League struct {
				Standings [][]struct {
					Team struct {
						ID int `json:"id"`
					} `json:"team"`
				} `json:"standings"`
			} `json:"league"`
		} `json:"response"`
	}

	// Get league IDs from database
	leagues, err := leagueModel.ListLeagues(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	httpClient := &http.Client{}

	// Initialize an empty list of team IDs
	teamIDs := []int{}

	for _, league := range leagues {
		footballApiId := league.FootballApiId
		s, err := leagueModel.GetCurrentSeasonForLeague(context.Background(), league)
		url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/standings?season=%d&league=%d", s.Year, footballApiId)

		req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
		if err != nil {
			fmt.Println("Error getting teams for league: ", footballApiId, " ", err)
		}
		req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("Error getting teams for league: ", footballApiId, " ", err)
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error getting teams for league: ", footballApiId, " ", err)
		}

		// Get Team FootballApiId from json response 'data'
		var response StandingsResponse
		err = json.Unmarshal(data, &response)
		if err != nil {
			fmt.Println("Error getting teams for league: ", footballApiId, " ", err)
			continue
		}

		// Get team IDs from response
		if len(response.Response) == 0 {
			continue
		}
		standings := response.Response[0].League.Standings[0]

		for _, teamStanding := range standings {
			teamIDs = append(teamIDs, teamStanding.Team.ID)
		}
	}

	return teamIDs
}

// Fetches standings from API and saves them to the database
func fetchStandings(standingsModel *models.StandingsModel) {
	standingsController := NewStandingsController(standingsModel)
	err := standingsController.InitializeStandings(leagueModel, context.Background())
	if err != nil {
		log.Printf("error initializing standings: %v", err)
	}
}

// Fetches Fixtures by league from API and saves them to the database
func fetchFixtures(fixturesModel *models.FixtureModel) {
	fixtureController := NewFixtureController(fixturesModel)
	err := fixtureController.InitializeFixtures(leagueModel, context.Background())
	if err != nil {
		log.Printf("error initializing fixtures: %v", err)
	}
}
