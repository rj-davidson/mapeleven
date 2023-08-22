package main

import (
	"context"
	"database/sql"
	_ "database/sql"
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
func resetDatabase() {
	// Connect to the default "postgres" database to reset the specific database
	connectionString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASS") + "@" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + "/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to the default postgres database: %v", err)
	}
	defer db.Close()

	// Drop the database
	_, err = db.Exec(`DROP DATABASE IF EXISTS "` + viper.GetString("DB_NAME") + `"`)
	if err != nil {
		log.Fatalf("Failed to drop database: %v", err)
	}

	// Create the database
	_, err = db.Exec(`CREATE DATABASE "` + viper.GetString("DB_NAME") + `"`)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
}

func main() {
	// Set up Viper configuration
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	viper.AutomaticEnv()

	// Reset the database
	resetDatabase()

	// Build Connection String
	connectionString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASS") + "@" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + "/" + viper.GetString("DB_NAME") + "?sslmode=disable"

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

		// Check if data has been transferred successfully
		players, err := client.Player.Query().All(context.Background())
		if err != nil {
			log.Fatalf("Failed to fetch players from database: %v", err)
		}
		if len(players) > 0 {
			fmt.Println("Data has been successfully transferred to the database!")
		} else {
			fmt.Println("No data found in the database!")
		}
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
