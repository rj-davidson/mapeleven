package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"net/http"
)

func SetupStatsRoutes(app *fiber.App) {

	// Get Top Scorers
	app.Get("/stats/topscorers", func(c *fiber.Ctx) error {
		req, err := http.NewRequestWithContext(context.Background(), "GET",
			"https://api-football-v1.p.rapidapi.com/v3/players/topscorers?league=39&season=2022", nil)
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

	// Get Team Stats
	app.Get("/stats/team", func(c *fiber.Ctx) error {
		req, err := http.NewRequestWithContext(context.Background(), "GET",
			"https://api-football-v1.p.rapidapi.com/v3/teams/statistics?league=39&season=2022&team=33", nil)
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

	// Get Player Stats
	app.Get("/stats/player", func(c *fiber.Ctx) error {
		req, err := http.NewRequestWithContext(context.Background(), "GET",
			"https://api-football-v1.p.rapidapi.com/v3/players?id=276&season=2022", nil)
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

	// Get Fixture Stats
	app.Get("/stats/fixture", func(c *fiber.Ctx) error {
		req, err := http.NewRequestWithContext(context.Background(), "GET",
			"https://api-football-v1.p.rapidapi.com/v3/fixtures/statistics?fixture=215662", nil)
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
