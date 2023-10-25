package main

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/server/routes"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/migrate"
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

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
		AllowOrigins: "http://" + viper.GetString("IP_HOST") +
			":3000, http://localhost:3000, https://" + viper.GetString("DOMAIN") + ", http://" +
			viper.GetString("IP HOST") + ":80, http://" + viper.GetString("IP_HOST"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Set up routes
	SetupRoutes(app, client)

	// Print Server Running Message
	fmt.Println("Server is running on port 8080")

	// Start the api
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
	routes.SetupFixtureRoutes(app, client)

	// Setup routes for search
	routes.SetupSearchRoutes(app, client)

	// Serve images from public directory
	app.Static(viper.GetString("ENV_PATH")+"/images", "./public/images")
}
