package main

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/controllers"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/fixture_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/team_models"
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

	if initialize {
		// Country Initialization
		fetchCountries(client)

		// SeasonID Initialization
		fetchLeagues(client)

		// Standings Initialization
		fetchStandings(client)

		// Fixtures Initialization
		fetchFixtures(client)

		// Team Stats Initialization
		fetchTeamStats(client)

		// Fetch Squads
		fetchSquads(client)

		// Fetch Fixture Data
		fetchFixtureData(client)

		if runScheduler {
			// Schedule cron jobs
			_, err := s.Every(12).Hours().At("00:00").Do(fetchCountries, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
			_, err = s.Every(12).Hours().At("00:00").Do(fetchLeagues, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
			_, err = s.Every(12).Hours().At("00:00").Do(fetchStandings, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
			_, err = s.Every(12).Hours().At("00:00").Do(fetchFixtures, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
			_, err = s.Every(12).Hours().At("00:00").Do(fetchTeamStats, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
			_, err = s.Every(12).Hours().At("00:00").Do(fetchSquads, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
			_, err = s.Every(12).Hours().At("00:00").Do(fetchFixtureData, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
			_, err = s.Every(150).Seconds().At("00:00").Do(fetchTodaysFixtures, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
			_, err = s.Every(45).Seconds().At("00:00").Do(fetchLiveFixtures, client)
			if err != nil {
				fmt.Printf("error scheduling job: %v", err)
			}
		}
	}

	if devRun {
		fetchSquads(client)
	}
}

func fetchCountries(client *ent.Client) {
	countryController := controllers.NewCountryController(client)
	err := countryController.InitializeCountries()
	if err != nil {
		log.Fatal(err)
	}
}

func fetchLeagues(client *ent.Client) {
	championsLeagueID := 2
	premierLeagueID := 39
	ligue1ID := 71
	bundesligaID := 78
	laLigaID := 140
	mlsID := 253
	leagueIDs := []int{premierLeagueID, mlsID, championsLeagueID, ligue1ID, bundesligaID, laLigaID}

	leagueController := controllers.NewLeagueController(client)
	err := leagueController.InitializeLeagues(leagueIDs, context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

// Fetches standings from API and saves them to the database
func fetchStandings(client *ent.Client) {
	standingsController := controllers.NewStandingsController(client)
	err := standingsController.InitializeStandings(context.Background())
	if err != nil {
		log.Printf("error initializing standings: %v", err)
	}
}

// Fetches Fixtures by league from API and saves them to the database
func fetchFixtures(client *ent.Client) {
	fixtureController := controllers.NewFixtureController(client)
	err := fixtureController.InitializeFixtures(context.Background())
	if err != nil {
		log.Printf("error initializing fixtures: %v", err)
	}
}

func fetchTeamStats(client *ent.Client) {
	teamStatsController := controllers.NewTeamStatsController(client)
	tm := team_models.NewTeamModel(client)
	teams, err := tm.ListAll(context.Background())
	if err != nil {
		log.Printf("error fetching teams: %v", err)
		return
	}
	err = teamStatsController.InitializeFixtures(teams, context.Background())
	if err != nil {
		log.Printf("error initializing team stats: %v", err)
	}
}

func fetchSquads(client *ent.Client) {
	fmt.Println("Fetching Squads...")

	tm := team_models.NewTeamModel(client)
	teams, err := tm.ListAll(context.Background())
	if err != nil {
		log.Printf("Error fetching teams: %v", err)
		return
	}

	squadController := controllers.NewSquadController(client)
	err = squadController.InitializeSquads(teams, context.Background())
	if err != nil {
		log.Printf("Error initializing squads: %v", err)
	}
	fmt.Println("Squads successfully loaded.")
}

func fetchFixtureData(client *ent.Client) {
	fmt.Println("Fetching fixture data...")
	fixtureDataController := controllers.NewFixtureDataController(client)
	err := fixtureDataController.FetchFixtures(context.Background())
	if err != nil {
		log.Printf("Error fetching fixture data: %v", err)
	}
	fmt.Println("Fixture data successfully loaded.")
}

func fetchTodaysFixtures(client *ent.Client) error {
	fmt.Println("Fetching today's fixtures...")
	fc := controllers.NewFixtureController(client)
	fdc := controllers.NewFixtureDataController(client)
	fm := fixture_models.NewFixtureModel(client)
	todaysFixtures, err := fm.ListTodaysFixtures(context.Background())
	if err != nil {
		fmt.Printf("Error fetching today's fixtures: %v", err)
	}
	err = fc.UpdateFixtures(todaysFixtures)
	if err != nil {
		fmt.Printf("Error updating today's fixtures: %v", err)
	}
	err = fdc.UpdateFixtures(context.Background(), todaysFixtures)
	if err != nil {
		fmt.Printf("Error updating today's fixtures: %v", err)
	}
	fmt.Println("Today's fixtures successfully loaded.")
	return nil
}

func fetchLiveFixtures(client *ent.Client) error {
	fmt.Println("Fetching live fixtures...")
	fc := controllers.NewFixtureController(client)
	fdc := controllers.NewFixtureDataController(client)
	fm := fixture_models.NewFixtureModel(client)
	liveFixtures, err := fm.ListLiveFixtures(context.Background())
	if err != nil {
		fmt.Printf("Error fetching live fixtures: %v", err)
	}
	err = fc.UpdateFixtures(liveFixtures)
	if err != nil {
		fmt.Printf("Error updating live fixtures: %v", err)
	}
	err = fdc.UpdateFixtures(context.Background(), liveFixtures)
	if err != nil {
		fmt.Printf("Error updating live fixtures: %v", err)
	}
	fmt.Println("Live fixtures successfully loaded.")
	return nil
}
