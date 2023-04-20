package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"net/http"
)

func SetupFixtureRoutes(app *fiber.App) {
	// Get Fixtures of a specific date
	app.Get("/fixtures/:date", func(c *fiber.Ctx) error {
		date := c.Params("date")
		url := "https://api-football-v1.p.rapidapi.com/v3/fixtures?date=" + date
		req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code %d", resp.StatusCode)
		}

		var result interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}

		return c.JSON(result)
	})
}
