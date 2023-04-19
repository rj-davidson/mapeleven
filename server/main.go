package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"mapeleven/controllers"
	"mapeleven/db/ent"
	"mapeleven/db/ent/migrate"
	"mapeleven/models"
	"mapeleven/routes"
)

func main() {
	// Set up Viper configuration
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	viper.AutomaticEnv()

	// Build Connection String
	connectionString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASS") + "@" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + "/" + viper.GetString("DB_NAME") + "?sslmode=disable"

	// Create a new Ent client instance
	client, err := ent.Open(dialect.Postgres, connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run migrations on startup
	err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Web App
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Initialize data
	initializeData(client)

	// Set up routes
	SetupRoutes(app, client)

	fmt.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}

func initializeData(client *ent.Client) {
	// Initialize data
	laLigaID := 140
	serieAID := 39
	bundesligaID := 78
	premierLeagueID := 61
	ligue1ID := 71
	majorleaguesoccerID := 253
	leagueIDs := []int{laLigaID, serieAID, bundesligaID, premierLeagueID, ligue1ID, majorleaguesoccerID}

	// Country Initialization
	//{
	//	countryModel := models.NewCountryModel(client)
	//	countryController := controllers.NewCountryController(countryModel)
	//	err := countryController.InitializeCountries()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

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

func SetupRoutes(app *fiber.App, client *ent.Client) {
	// Setup routes for players
	routes.SetupPlayersRoutes(app, client)

	// Setup routes for teams
	routes.SetupTeamsRoutes(app, client)

	// Setup routes for leagues
	routes.SetupLeaguesRoutes(app, client)

	// Setup routes for stats
	routes.SetupStatsRoutes(app)

	// Serve images from public directory
	app.Static("/images", "./public/images")
}
