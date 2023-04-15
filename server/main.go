package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"mapeleven/ent"
	"mapeleven/ent/migrate"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {
	// Set up Viper configuration
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
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

	// Set up routes
	setupRoutes(app, client)

	fmt.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}

func setupRoutes(app *fiber.App, client *ent.Client) {
	// Add your routes here, e.g., creating a new user
	app.Get("/", func(c *fiber.Ctx) error {
		return api(c, client)
	})
}

func api(c *fiber.Ctx, client *ent.Client) error {
	url := "https://api-football-v1.p.rapidapi.com/v3/teams?id=33"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", viper.GetString("API_HOST"))
	fmt.Println("API Key: ", viper.GetString("API_KEY"))
	fmt.Println("API Host: ", viper.GetString("API_HOST"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(res)
	fmt.Println(string(body))

	c.Type("application/json", "utf-8")
	return c.Send(body)
}
