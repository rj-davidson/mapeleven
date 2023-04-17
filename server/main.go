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
	"mapeleven/fetchers"
	"mapeleven/handlers"
	"mapeleven/models"

	"mapeleven/models/ent"
	"mapeleven/models/ent/migrate"
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
	err = initializeLeagues(client)
	if err != nil {
		log.Fatal(err)
	}

	// Set up routes
	setupAPIRoutes(app)

	fmt.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}

func setupAPIRoutes(app *fiber.App) {
	// Players endpoints
	app.Get("/players", func(c *fiber.Ctx) error {
		return handlers.GetAllPlayers(c)
	})
	app.Get("/players/:id", func(c *fiber.Ctx) error {
		return handlers.GetPlayerByID(c)
	})

	// Other endpoints
	// ...
}

func initializeLeagues(client *ent.Client) error {
	ids := []int{1, 39}

	// Initialize models
	leagueModel := models.NewLeagueModel(client)
	countryModel := models.NewCountryModel(client)

	// Create a new CountryFetcher instance
	countryFetcher := fetchers.NewCountryFetcher(countryModel, client)

	// Create a new LeagueFetcher instance
	leagueFetcher := fetchers.NewLeagueFetcher(leagueModel, countryFetcher)

	// Fetch and upsert leagues
	for _, id := range ids {
		league, err := leagueFetcher.UpsertLeague(id)
		if err != nil {
			return fmt.Errorf("failed to upsert league with ID %d: %v", id, err)
		}
		fmt.Printf("Upserted league with ID %d: %+v\n", id, league)
	}

	fmt.Println("Leagues loaded")

	return nil
}
