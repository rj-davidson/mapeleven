package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", apitest)
	fmt.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}

func apitest(c *fiber.Ctx) error {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	url := "https://api-football-v1.p.rapidapi.com/v3/fixtures?live=all"

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
	c.Type("application/json", "utf-8")
	return c.Send(body)
}
