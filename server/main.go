package main

import (
	"context"
	_ "database/sql"
	"entgo.io/ent/dialect"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"mapeleven/controllers"
	"mapeleven/db/ent"
	"mapeleven/db/ent/migrate"
	"mapeleven/routes"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var updateData bool

var initialize bool

var devRun bool

func init() {
	flag.BoolVar(&initialize, "initialize", false, "Force initialization")
	flag.BoolVar(&updateData, "update", false, "Enable cron job scheduler")
	flag.BoolVar(&devRun, "dev", false, "Enable dev mode")
	flag.Parse()
}

//func resetDatabase() {
//	// Connect to the default "postgres" database to reset the specific database
//	connectionString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASS") + "@" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + "/postgres?sslmode=disable"
//	db, err := sql.Open("postgres", connectionString)
//	if err != nil {
//		log.Fatalf("Failed to connect to the default postgres database: %v", err)
//	}
//	defer db.Close()
//
//	// Drop the database
//	_, err = db.Exec(`DROP DATABASE IF EXISTS "` + viper.GetString("DB_NAME") + `"`)
//	if err != nil {
//		log.Fatalf("Failed to drop database: %v", err)
//	}
//
//	// Create the database
//	_, err = db.Exec(`CREATE DATABASE "` + viper.GetString("DB_NAME") + `"`)
//	if err != nil {
//		log.Fatalf("Failed to create database: %v", err)
//	}
//}

func main() {
	// Set up Viper configuration
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
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

	// If initialize flag is true, run initialization logic
	fmt.Println("-- Initialize flag: ", initialize)
	fmt.Println("-- Update flag: ", updateData)
	fmt.Println("-- Dev flag: ", devRun)
	go func() {
		controllers.CronScheduler(client, initialize, updateData, devRun)
		wg.Done()
	}()

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run migrations on startup
	err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Web App Initialization
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Set up routes
	SetupRoutes(app, client)

	// Print Server Running Message
	fmt.Println("Server is running on port 8080")

	// Start the server
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutdown signal received. Shutting down...")

	// Wait for all goroutines (like the CronScheduler) to finish
	wg.Wait()

	fmt.Println("Server stopped.")
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

	// Setup routes for stats
	routes.SetupFixtureRoutes(app)

	// Setup routes for search
	routes.SetupSearchRoutes(app, client)

	// Serve images from public directory
	app.Static("/images", "./public/images")
}
