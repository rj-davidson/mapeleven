package main

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/controllers"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/fixture_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models/player_stats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/team_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/team_models/team_stats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/migrate"
	"context"
	"entgo.io/ent/dialect"
	"flag"
	"fmt"
	"github.com/go-co-op/gocron"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	initialize := flag.Bool("init", false, "Initialize the database")
	update := flag.Bool("update", false, "Update the database")
	devRun := flag.Bool("dev", false, "Run in development mode")
	flag.Parse()

	envPath := "./.env"
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf("Error: .env file not found at %s", envPath)
	}

	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	viper.AutomaticEnv()

	// Initialize WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Build Connection String
	connectionString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASS") + "@" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + "/" + viper.GetString("DB_NAME") + "?sslmode=disable"

	// Connect to Database
	client, err := ent.Open(dialect.Postgres, connectionString)
	if err != nil {
		log.Fatalf("failed connecting to postgres database: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Run migrations on startup
	err = client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize scheduler
	s := gocron.NewScheduler(time.UTC)

	if *initialize {
		fmt.Println("Initializing database...")
		initializeDatabase(client)
		fmt.Println("Database initialization complete.")
	}

	if *update {
		fmt.Println("Updating database...")
		updateDatabase(client)
		fmt.Println("Database update complete.")
	}

	if *devRun {
		fmt.Println("Running in development mode...")
		runDevelopmentMode(client)
	}

	// Start the scheduler
	if *initialize || *update {
		go s.StartAsync()
		defer s.Stop()
	}

	// Keep the main goroutine alive while scheduler is running
	select {}
}

func initializeDatabase(client *ent.Client) {
	cronScheduler(client, true, false, false)
}

func updateDatabase(client *ent.Client) {
	cronScheduler(client, false, true, false)
}

func runDevelopmentMode(client *ent.Client) {
	cronScheduler(client, false, false, true)
}

func cronScheduler(client *ent.Client, initialize bool, runScheduler bool, devRun bool) {
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
	fixturesModel := fixture_models.NewFixtureModel(client)
	playerModel := player_models.NewPlayerModel(client)
	squadModel := team_models.NewSquadModel(client)
	playerStatsModel := player_stats.NewPlayerStatsModel(client)

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

		// Fetch Squads
		fetchSquads(squadModel, playerModel, teamModel, playerStatsModel)

		// Fetch Player Stats
		//fetchPlayerStats(playerModel, playerStatsModel)

		// Fetch Fixture Data
		fetchFixtureData(fixturesModel)
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
		// Add your development mode code here
		fetchSquads(squadModel, playerModel, teamModel, playerStatsModel)
		//fetchPlayerStats(playerModel, teamModel)

	}
}

func fetchCountries(countryModel *models.CountryModel) {
	countryController := controllers.NewCountryController(countryModel)
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
	//bundesligaID := 78
	//laLigaID := 140
	majorleaguesoccerID := 253
	leagueIDs := []int{premierLeagueID, majorleaguesoccerID, championsLeagueID}
	//leagueIDs := []int{majorleaguesoccerID, championsLeagueID, europaLeagueID}

	leagueController := controllers.NewLeagueController(leagueModel, seasonModel)
	err := leagueController.InitializeLeagues(leagueIDs, context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

// Fetches standings from API and saves them to the database
func fetchStandings(standingsModel *models.StandingsModel, clubModel *models.ClubModel, leagueModel *models.LeagueModel, teamModel *team_models.TeamModel) {
	clubController := controllers.NewClubController(clubModel)
	teamController := controllers.NewTeamController(teamModel)
	standingsController := controllers.NewStandingsController(standingsModel, clubController, teamController)
	err := standingsController.InitializeStandings(leagueModel, context.Background())
	if err != nil {
		log.Printf("error initializing standings: %v", err)
	}
}

// Fetches Fixtures by league from API and saves them to the database
func fetchFixtures(fixturesModel *fixture_models.FixtureModel, leagueModel *models.LeagueModel) {
	fixtureController := controllers.NewFixtureController(fixturesModel)
	err := fixtureController.InitializeFixtures(leagueModel, context.Background())
	if err != nil {
		log.Printf("error initializing fixtures: %v", err)
	}
}

func fetchTeamStats(teamModel *team_models.TeamModel, teamStatsModel *team_stats.TeamStatsModel) {
	teamStatsController := controllers.NewTeamStatsController(teamStatsModel)
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

func fetchSquads(squadModel *team_models.SquadModel, playerModel *player_models.PlayerModel, teamModel *team_models.TeamModel, playerstatsModel *player_stats.PlayerStatsModel) {
	fmt.Println("Fetching Squads...")

	teams, err := teamModel.ListAll(context.Background())
	if err != nil {
		log.Printf("Error fetching teams: %v", err)
		return
	}

	squadController := controllers.NewSquadController(squadModel, playerModel, playerstatsModel)
	err = squadController.InitializeSquads(teams, context.Background())
	if err != nil {
		log.Printf("Error initializing squads: %v", err)
	}
	fmt.Println("Squads successfully loaded.")
}

func fetchFixtureData(fixtureModel *fixture_models.FixtureModel) {
	fmt.Println("Fetching fixture data...")
	fixtureDataController := controllers.NewFixtureDataController(fixtureModel)
	err := fixtureDataController.FetchFixtures(context.Background())
	if err != nil {
		log.Printf("Error fetching fixture data: %v", err)
	}
	fmt.Println("Fixture data successfully loaded.")
}
