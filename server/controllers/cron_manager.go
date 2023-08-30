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

func CronScheduler(client *ent.Client, initialize bool, runScheduler bool) {
	// Initialize scheduler
	s := gocron.NewScheduler(time.UTC)

	// Model Declaration
	countryModel = models.NewCountryModel(client)
	leagueModel = models.NewLeagueModel(client)
	teamModel = models.NewTeamModel(client)

	if initialize {
		// Country Initialization
		fetchCountries(countryModel)

		// League Initialization
		fetchLeagues(leagueModel)

		// Team Initialization
		fetchTeams(teamModel)
	}

	if runScheduler {
		// Schedule cron jobs
		s.Every(14).Days().At("00:00").Do(fetchCountries, countryModel)
		s.Every(14).Days().At("00:00").Do(fetchLeagues, leagueModel)
		s.Every(14).Days().At("00:00").Do(fetchTeams, teamModel)
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
