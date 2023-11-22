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
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			fmt.Printf("Error closing client: %v", err)
		}
	}(client)
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

	dataManager(client, *initialize, *update, *devRun)

	select {}
}

func dataManager(client *ent.Client, initialize bool, runScheduler bool, devRun bool) {
	if initialize || runScheduler {
		fetchCountries(client)
		fetchLeagues(client)
		fetchStandings(client)
		fetchFixtures(client)
		fetchTeamStats(client)
		fetchSquads(client)
		fetchFixtureData(client)
	}

	if runScheduler {
		updaterTicker(client)
	}

	if devRun {
		/* EMPTY */
	}
}

func updaterTicker(client *ent.Client) {
	counTicker := time.NewTicker(12 * time.Hour)
	leagTicker := time.NewTicker(12 * time.Hour)
	standTicker := time.NewTicker(12 * time.Hour)
	fixtTicker := time.NewTicker(12 * time.Hour)
	teamTicker := time.NewTicker(12 * time.Hour)
	squaTicker := time.NewTicker(12 * time.Hour)
	fixDTicker := time.NewTicker(12 * time.Hour)
	fetchTodayTicker := time.NewTicker(150 * time.Second)
	fetchLiveTicker := time.NewTicker(45 * time.Second)

	go func() {
		for {
			select {
			case <-counTicker.C:
				go fetchCountries(client)
			case <-leagTicker.C:
				go fetchLeagues(client)
			case <-standTicker.C:
				go fetchStandings(client)
			case <-fixtTicker.C:
				go fetchFixtures(client)
			case <-teamTicker.C:
				go fetchTeamStats(client)
			case <-squaTicker.C:
				go fetchSquads(client)
			case <-fixDTicker.C:
				go fetchFixtureData(client)
			case <-fetchTodayTicker.C:
				go fetchTodaysFixtures(client)
			case <-fetchLiveTicker.C:
				go fetchLiveFixtures(client)
			}
		}
	}()
}

func fetchCountries(client *ent.Client) {
	countryController := controllers.NewCountryController(client)
	err := countryController.InitializeCountries()
	if err != nil {
		fmt.Printf("error initializing countries: %v", err)
	}
}

func fetchLeagues(client *ent.Client) {
	championsLeagueID := 2
	premierLeagueID := 39
	ligue1ID := 61
	bundesligaID := 78
	laLigaID := 140
	mlsID := 253
	leagueIDs := []int{premierLeagueID, mlsID, championsLeagueID, ligue1ID, bundesligaID, laLigaID}

	leagueController := controllers.NewLeagueController(client)
	err := leagueController.InitializeLeagues(leagueIDs, context.Background())
	if err != nil {
		fmt.Printf("error initializing leagues: %v", err)
	}
}

// Fetches standings from API and saves them to the database
func fetchStandings(client *ent.Client) {
	standingsController := controllers.NewStandingsController(client)
	err := standingsController.InitializeStandings(context.Background())
	if err != nil {
		fmt.Printf("error initializing standings: %v", err)
	}
}

// Fetches Fixtures by league from API and saves them to the database
func fetchFixtures(client *ent.Client) {
	fixtureController := controllers.NewFixtureController(client)
	err := fixtureController.InitializeFixtures(context.Background())
	if err != nil {
		fmt.Printf("error initializing fixtures: %v", err)
	}
}

func fetchTeamStats(client *ent.Client) {
	teamStatsController := controllers.NewTeamStatsController(client)
	tm := team_models.NewTeamModel(client)
	teams, err := tm.ListAll(context.Background())
	if err != nil {
		fmt.Printf("error fetching teams: %v", err)
		return
	}
	err = teamStatsController.InitializeFixtures(teams, context.Background())
	if err != nil {
		fmt.Printf("error initializing team stats: %v", err)
	}
}

func fetchSquads(client *ent.Client) {
	fmt.Println("Fetching Squads...")

	tm := team_models.NewTeamModel(client)
	teams, err := tm.ListAll(context.Background())
	if err != nil {
		fmt.Printf("error fetching teams: %v", err)
		return
	}

	squadController := controllers.NewSquadController(client)
	err = squadController.InitializeSquads(teams, context.Background())
	if err != nil {
		fmt.Printf("error initializing squads: %v", err)
	}
	fmt.Println("Squads successfully loaded.")
}

func fetchFixtureData(client *ent.Client) {
	fmt.Println("Fetching fixture data...")
	fixtureDataController := controllers.NewFixtureDataController(client)
	err := fixtureDataController.FetchFixtures(context.Background())
	if err != nil {
		fmt.Printf("error fetching fixture data: %v", err)
	}
	fmt.Println("Fixture data successfully loaded.")
}

func fetchTodaysFixtures(client *ent.Client) {
	fc := controllers.NewFixtureController(client)
	fdc := controllers.NewFixtureDataController(client)
	fm := fixture_models.NewFixtureModel(client)
	todaysFixtures, err := fm.ListTodaysFixtures(context.Background())
	if len(todaysFixtures) == 0 {
		return
	}
	fmt.Println("Fetching today's fixtures...")
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
}

func fetchLiveFixtures(client *ent.Client) {
	fc := controllers.NewFixtureController(client)
	fdc := controllers.NewFixtureDataController(client)
	fm := fixture_models.NewFixtureModel(client)
	liveFixtures, err := fm.ListLiveFixtures(context.Background())
	if len(liveFixtures) == 0 {
		return
	}
	fmt.Println("Fetching live fixtures...")
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
}
