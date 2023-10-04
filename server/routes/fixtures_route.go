package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"mapeleven/db/ent"
	"mapeleven/serializer"
	"net/http"
)

func SetupFixtureRoutes(app *fiber.App, client *ent.Client) {
	ss := serializer.NewScoreboardSerializer(client)

	app.Get(viper.GetString("ENV_PATH")+"/fixtures/:slug", getFixtureBySlug())
	app.Get(viper.GetString("ENV_PATH")+"/fixtures/league/:slug", getFixturesByLeague())
	app.Get(viper.GetString("ENV_PATH")+"/scoreboard", getScoreboard(ss))

	// TODO: DEPRECATED
	app.Get("/fixtures", func(c *fiber.Ctx) error {
		date := c.Query("date")
		live := c.Query("live")
		url := "https://api-football-v1.p.rapidapi.com/v3/fixtures"

		// If neither are requested, return Error
		if date != "" && live != "" {
			return fmt.Errorf("invalid fixture query parameters for end-point [%s]: date=%s, live=%s",
				"/fixtures", date, live)
		}

		// Format URL to fetch live fixtures
		if live == "all" {
			url += "?live=all"
		}

		// Format URL to fetch fixtures by date
		if date != "" {
			url += "?date=" + date
		}

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

// TODO: Implement
func getFixtureBySlug() fiber.Handler {
	return nil
}

// TODO: Implement
func getFixturesByLeague() fiber.Handler {
	return nil
}

func getScoreboard(serializer *serializer.ScoreboardSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		scoreboard := serializer.SerializeScoreboard()
		return c.JSON(scoreboard)
	}
}
