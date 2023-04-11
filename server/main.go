package main

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"mapeleven-server/ent"
	"net/http"
	"time"
)

func main() {
	// Ent -> Database
	cfg := &Config{
		Dialect:  dialect.Postgres,
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASS"),
		Database: viper.GetString("DB_NAME"),
	}

	client, err := ent.Open(cfg.Dialect.String(), fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database,
	))
	if err != nil {
		// handle error
	}
	defer client.Close()
	// Web App
	app := fiber.New()
	app.Get("/", apitest)
	fmt.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
func apitest(c *fiber.Ctx) error {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	url := "https://api-football-v1.p.rapidapi.com/v3/players/topscorers?league=39&season=2022"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", viper.GetString("API_HOST"))
	fmt.Println(viper.GetString("API_KEY"))
	fmt.Println(viper.GetString("API_HOST"))
	fmt.Println("Viper Data Should Be Above")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	fmt.Println("Requested Data", viper.GetString("USER"))
	return c.SendString(string(body))
}

//func mainold() {
//	// Setup Viper Environment Variables
//	viper.SetConfigFile(".env")
//	viper.ReadInConfig()
//
//	connStr := "user=" + viper.GetString("USER") + " password=" + viper.GetString("PASSWORD") + " dbname=" + viper.GetString("DB_NAME") + " sslmode=verify-full"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	db.Stats()
//	app := fiber.New()
//
//	app.Use(cors.New())
//
//	app.Use(func(c *fiber.Ctx) error {
//		return c.SendStatus(404) // => 404 "Not Found"
//	})
//
//	basicApiCall()
//	log.Fatal(app.Listen(":8080"))
//}
