package db

import (
	"log"
	"mapeleven/controllers"
	"mapeleven/db/ent"
	"mapeleven/models"
)

func InitializeData(client *ent.Client) {
	// Initialize data
	laLigaID := 140
	serieAID := 39
	bundesligaID := 78
	premierLeagueID := 61
	ligue1ID := 71
	majorleaguesoccerID := 253
	leagueIDs := []int{laLigaID, serieAID, bundesligaID, premierLeagueID, ligue1ID, majorleaguesoccerID}

	// Country Initialization
	{
		countryModel := models.NewCountryModel(client)
		countryController := controllers.NewCountryController(countryModel)
		err := countryController.InitializeCountries()
		if err != nil {
			log.Fatal(err)
		}
	}

	// League Initialization
	{
		leagueModel := models.NewLeagueModel(client)
		leagueController := controllers.NewLeagueController(leagueModel)
		err := leagueController.InitializeLeagues(leagueIDs)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Team Initialization
	{
		teamModel := models.NewTeamModel(client)
		teamController := controllers.NewTeamController(teamModel)
		err := teamController.InitializeTeams(leagueIDs)
		if err != nil {
			log.Fatal(err)
		}
	}
}
