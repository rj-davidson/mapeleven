package controllers

import (
	"context"
	"github.com/go-co-op/gocron"
	"log"
	"mapeleven/db/ent"
	"mapeleven/models"
	"mapeleven/models/team_models"
	"mapeleven/models/team_models/team_stats"
	"time"
)

func CronScheduler(client *ent.Client, initialize bool, runScheduler bool, devRun bool) {
	// Initialize scheduler
	s := gocron.NewScheduler(time.UTC)

	// Model Declaration
	seasonModel := models.NewSeasonModel(client)
	countryModel := models.NewCountryModel(client)
	leagueModel := models.NewLeagueModel(client)
	clubModel := models.NewClubModel(client)
	teamModel := team_models.NewTeamModel(client)
	teamStatsModel := team_stats.NewTeamStatsModel(client)
	standingsModel := models.NewStandingsModel(client)
	fixturesModel := models.NewFixtureModel(client)

	if initialize {
		// Country Initialization
		fetchCountries(countryModel)

		// SeasonID Initialization
		fetchLeagues(leagueModel, seasonModel)

		// Standings Initialization
		fetchStandings(standingsModel, clubModel, leagueModel, teamModel)

		// Fixtures Initialization
		fetchFixtures(fixturesModel, leagueModel)

		// Team Stats Initialization
		fetchTeamStats(teamModel, teamStatsModel)
	}

	if runScheduler {
		// Schedule cron jobs
		s.Every(14).Days().At("00:00").Do(fetchCountries, countryModel)
		s.Every(14).Days().At("00:00").Do(fetchLeagues, leagueModel)
		s.Every(14).Days().At("00:00").Do(fetchStandings, standingsModel, clubModel)
		s.Every(14).Days().At("00:00").Do(fetchFixtures, fixturesModel)
		s.Every(14).Days().At("00:00").Do(fetchTeamStats, teamModel, teamStatsModel)
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
	//europaLeagueID := 3
	premierLeagueID := 39
	//serieAID := 61
	//ligue1ID := 71
	bundesligaID := 78
	laLigaID := 140
	majorleaguesoccerID := 253
	leagueIDs := []int{laLigaID, bundesligaID, premierLeagueID, majorleaguesoccerID, championsLeagueID}
	//leagueIDs := []int{majorleaguesoccerID, championsLeagueID, europaLeagueID}

	leagueController := NewLeagueController(leagueModel, seasonModel)
	err := leagueController.InitializeLeagues(leagueIDs, context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

// Fetches standings from API and saves them to the database
func fetchStandings(standingsModel *models.StandingsModel, clubModel *models.ClubModel, leagueModel *models.LeagueModel, teamModel *team_models.TeamModel) {
	clubController := NewClubController(clubModel)
	teamController := NewTeamController(teamModel)
	standingsController := NewStandingsController(standingsModel, clubController, teamController)
	err := standingsController.InitializeStandings(leagueModel, context.Background())
	if err != nil {
		log.Printf("error initializing standings: %v", err)
	}
}

// Fetches Fixtures by league from API and saves them to the database
func fetchFixtures(fixturesModel *models.FixtureModel, leagueModel *models.LeagueModel) {
	fixtureController := NewFixtureController(fixturesModel)
	err := fixtureController.InitializeFixtures(leagueModel, context.Background())
	if err != nil {
		log.Printf("error initializing fixtures: %v", err)
	}
}

func fetchTeamStats(teamModel *team_models.TeamModel, teamStatsModel *team_stats.TeamStatsModel) {
	teamStatsController := NewTeamStatsController(teamStatsModel)
	teams, err := teamModel.ListAll(context.Background())
	if err != nil {
		log.Printf("error fetching teams: %v", err)
		return
	}
	err = teamStatsController.InitializeFixtures(teams, context.Background())
	if err != nil {
		log.Printf("error initializing team stats: %v", err)
	}
}
