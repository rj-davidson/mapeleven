package routes

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/server/serializer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func SetupLeaguesRoutes(app *fiber.App, client *ent.Client) {
	leagueSerializer := serializer.NewLeagueSerializer(client)

	app.Get(viper.GetString("ENV_PATH")+"/leagues", getAllLeagues(leagueSerializer))
	app.Get(viper.GetString("ENV_PATH")+"/leagues/:slug", getLeagueBySlug(leagueSerializer))
}

func getAllLeagues(serializer *serializer.LeagueSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		leagues, err := serializer.GetCurrentLeagues(context.Background())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": fmt.Sprintf("Failed to get leagues: %v", err),
			})
		}
		return c.JSON(leagues)
	}
}

func getLeagueBySlug(serializer *serializer.LeagueSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")

		league, err := serializer.GetLeagueBySlug(context.Background(), slug)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"error": fmt.Sprintf("SeasonID not found: %v", err),
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": fmt.Sprintf("Failed to get league: %v", err),
			})
		}
		return c.JSON(league)
	}
}
