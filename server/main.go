package main

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

func initEntClient() (*ent.Client, error) {
	viper.AutomaticEnv()
	dbUser := viper.GetString("DB_USER")
	dbPass := viper.GetString("DB_PASS")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	connStr := "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"
	connStr = fmt.Sprintf(connStr, dbHost, dbPort, dbUser, dbName, dbPass)

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}

func main() {
	// Web App
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	entClient, errorEnt := initEntClient()
	if errorEnt != nil {
		log.Fatalf("failed to connect to ent: %v", errorEnt)
	}
	defer entClient.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return api(c, entClient)
	})

	fmt.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}

func api(c *fiber.Ctx, client *ent.Client) error {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	url := "https://api-football-v1.p.rapidapi.com/v3/players/topscorers?league=39&season=2022"

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
