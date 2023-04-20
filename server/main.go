package main

import (
	"context"
	"entgo.io/ent/dialect"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"mapeleven/db"
	"mapeleven/db/ent"
	"mapeleven/db/ent/migrate"
	"mapeleven/routes"
)

var updateData bool

func init() {
	flag.BoolVar(&updateData, "update-data", false, "Update data on startup")
	flag.Parse()
}

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
	if updateData {
		db.InitializeData(client) // Use the exported function
	}

	// Set up routes
	SetupRoutes(app, client)

	fmt.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
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

	// Setup Routes for stats
	routes.SetupFixtureRoutes(app)

	// Serve images from public directory
	app.Static("/images", "./public/images")
}
