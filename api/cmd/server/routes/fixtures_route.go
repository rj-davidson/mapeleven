package routes

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/server/serializer/fixture_serializer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func SetupFixtureRoutes(app *fiber.App, client *ent.Client) {
	ss := fixture_serializer.NewScoreboardSerializer(client)
	fs := fixture_serializer.NewFixtureSerializer(client)
	dfs := fixture_serializer.NewDemoFixtureSerializer(client)

	app.Get(viper.GetString("ENV_PATH")+"/fixtures/:slug", getFixtureBySlug(fs))
	app.Get(viper.GetString("ENV_PATH")+"/fixtures-demo/:slug", getDemoFixtureBySlug(dfs))
	app.Get(viper.GetString("ENV_PATH")+"/scoreboard", getScoreboard(ss))
}

func getFixtureBySlug(serializer *fixture_serializer.FixtureSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		hasAnyParam := c.Query("fixture") != "" || c.Query("events") != "" || c.Query("lineups") != ""
		defaultVal := !hasAnyParam

		getFixture := getParamOrDefault(c, "fixture", defaultVal)
		getEvents := getParamOrDefault(c, "events", defaultVal)
		getLineups := getParamOrDefault(c, "lineups", defaultVal)

		f, err := serializer.GetFixtureBySlug(context.Background(), slug, getFixture, getEvents, getLineups)
		if handleError(c, err) {
			return nil
		}
		return c.JSON(f)
	}
}

func getDemoFixtureBySlug(dfs *fixture_serializer.DemoFixtureSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		hasAnyParam := c.Query("fixture") != "" || c.Query("events") != "" || c.Query("lineups") != ""
		defaultVal := !hasAnyParam

		getFixture := getParamOrDefault(c, "fixture", defaultVal)
		getEvents := getParamOrDefault(c, "events", defaultVal)
		getLineups := getParamOrDefault(c, "lineups", defaultVal)

		f, err := dfs.GetDemoFixtureBySlug(context.Background(), slug, getFixture, getEvents, getLineups)
		if handleError(c, err) {
			return nil
		}
		return c.JSON(f)
	}
}

func getParamOrDefault(c *fiber.Ctx, param string, defaultVal bool) bool {
	value := c.Query(param)
	if value == "" {
		return defaultVal
	}
	return value == "true"
}

// TODO: Implement
func getFixturesByLeague() fiber.Handler {
	return nil
}

func getScoreboard(serializer *fixture_serializer.ScoreboardSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		scoreboard := serializer.SerializeScoreboard()
		return c.JSON(scoreboard)
	}
}

func handleError(c *fiber.Ctx, err error) bool {
	if err != nil {
		if ent.IsNotFound(err) {
			c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"error": fmt.Sprintf("Entity not found: %v", err),
			})
			return true
		}
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": fmt.Sprintf("Internal Server Error: %v", err),
		})
		return true
	}
	return false
}
