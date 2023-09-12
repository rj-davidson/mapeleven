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

var countryModel *models.CountryModel
var leagueModel *models.LeagueModel
var teamModel *models.TeamModel
var standingsModel *models.StandingsModel
var fixturesModel *models.FixtureModel

//var seasonModel *models.SeasonModel
//var teamSeasonsModel *models.TeamSeasonsModel

func CronScheduler(client *ent.Client, initialize bool, runScheduler bool, devRun bool) {
	// Initialize scheduler
	s := gocron.NewScheduler(time.UTC)

	// Model Declaration
	countryModel = models.NewCountryModel(client)
	leagueModel = models.NewLeagueModel(client)
	teamModel = models.NewTeamModel(client)
	standingsModel = models.NewStandingsModel(client)
	fixturesModel = models.NewFixtureModel(client)
	//seasonModel = models.NewSeasonModel(client)
	//teamSeasonsModel = models.NewTeamSeasonsModel(client)

	if initialize {
		// Country Initialization
		fetchCountries(countryModel)

		// League Initialization
		fetchLeagues(leagueModel)

		// Team Initialization
		fetchTeams(teamModel)

		// Standings Initialization
		fetchStandings(standingsModel)

		// Fixtures Initialization
		fetchFixtures(fixturesModel)

		// Season Initialization
		//fetchSeason(seasonModel)

		// Team Seasons Initialization
		//fetchTeamSeasons(teamSeasonsModel)
	}

	if runScheduler {
		// Schedule cron jobs
		s.Every(14).Days().At("00:00").Do(fetchCountries, countryModel)
		s.Every(14).Days().At("00:00").Do(fetchLeagues, leagueModel)
		s.Every(14).Days().At("00:00").Do(fetchTeams, teamModel)
		s.Every(14).Days().At("00:00").Do(fetchStandings, standingsModel)
		s.Every(14).Days().At("00:00").Do(fetchFixtures, fixturesModel)
		//s.Every(14).Days().At("00:00").Do(fetchSeason, seasonModel)
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

func fetchLeagues(leagueModel *models.LeagueModel) {
	laLigaID := 140
	serieAID := 39
	bundesligaID := 78
	premierLeagueID := 61
	ligue1ID := 71
	majorleaguesoccerID := 253
	championsLeagueID := 2
	europaLeagueID := 3
	//clubIDs := []int{42, 50, 34, 33, 47, 66, 40, 51, 36, 55, 49, 52, 48, 39, 35, 63, 46, 45, 65, 41}
	leagueIDs := []int{laLigaID, serieAID, bundesligaID, premierLeagueID, ligue1ID, majorleaguesoccerID, championsLeagueID, europaLeagueID}

	leagueController := NewLeagueController(leagueModel)
	err := leagueController.InitializeLeagues(leagueIDs)
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
		leagueID := league.ID
		url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/standings?season=2022&league=%d", leagueID)

		req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
		if err != nil {
			fmt.Println("Error getting teams for league: ", leagueID, " ", err)
		}
		req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("Error getting teams for league: ", leagueID, " ", err)
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error getting teams for league: ", leagueID, " ", err)
		}

		// Get Team ID from json response 'data'
		var response StandingsResponse
		err = json.Unmarshal(data, &response)
		if err != nil {
			fmt.Println("Error getting teams for league: ", leagueID, " ", err)
		}

		// Get team IDs from response
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
		fmt.Errorf("error initializing standings: %w", err)
	}
}

// Fetches season from API and saves them to the database
func fetchSeason(seasonModel *models.SeasonModel) {
	// Fetch seasonIDs from the database (this is just a placeholder)
	seasonIDs, err := seasonModel.ListSeasonIDs(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	seasonController := NewSeasonController(seasonModel)
	err = seasonController.InitializeSeasons(seasonIDs)
	if err != nil {
		log.Fatal(err)
	}
}

// Fetches team seasons from API and saves them to the database
func fetchTeamSeasons(teamSeasonsModel *models.TeamSeasonsModel) {
	// Initialize the TeamSeasonsController
	teamSeasonsController := NewTeamSeasonsController(teamSeasonsModel)

	// Initialize an HTTP client
	httpClient := &http.Client{}

	// Fetch team IDs based on your requirements.
	teamIDs := fetchTeamIDs()

	// Loop through each team ID to fetch and store data
	for _, teamID := range teamIDs {
		// Construct the API URL based on the current team ID
		url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams/seasons?team=%d", teamID)

		// Create a new request
		req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
		if err != nil {
			log.Printf("Error fetching team seasons for team %d: %v", teamID, err)
			continue
		}
		req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

		// Make the request
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Printf("Error fetching team seasons for team %d: %v", teamID, err)
			continue
		}
		defer resp.Body.Close()

		// Read the response data
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading team seasons data for team %d: %v", teamID, err)
			continue
		}

		// Define a structure to unmarshal the response data
		type ApiResponse struct {
			Response []int `json:"response"`
		}

		var apiResponse ApiResponse
		err = json.Unmarshal(data, &apiResponse)
		if err != nil {
			log.Printf("Error unmarshaling team seasons data for team %d: %v", teamID, err)
			continue
		}

		// Initialize team seasons for each season in the response
		for _, seasonID := range apiResponse.Response {
			input := models.CreateTeamSeasonInput{
				TeamSeasonID: teamID*1000 + seasonID, // Create a unique ID
				TeamID:       teamID,
				SeasonID:     seasonID,
			}
			_, err := teamSeasonsController.teamSeasonsModel.CreateTeamSeason(context.Background(), input)
			if err != nil {
				log.Printf("Error inserting team season for team %d and season %d: %v", teamID, seasonID, err)
			}
		}
	}
}

// Fetches Fixtures by league from API and saves them to the database
func fetchFixtures(fixturesModel *models.FixtureModel) {
	fixtureController := NewFixtureController(fixturesModel)
	err := fixtureController.InitializeFixtures(leagueModel, context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
