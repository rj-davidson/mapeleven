package routes

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mapeleven/db/ent"
	"mapeleven/serializer"
)

func SetupLeaguesRoutes(app *fiber.App, client *ent.Client) {
	leagueSerializer := serializer.NewLeagueSerializer(client)

	app.Get("/leagues", getAllLeagues(leagueSerializer))
	app.Get("/leagues/:slug", getLeagueBySlug(leagueSerializer))
}

func getAllLeagues(serializer *serializer.LeagueSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		leagues, err := serializer.GetLeagues(context.Background())
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
					"error": fmt.Sprintf("League not found: %v", err),
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": fmt.Sprintf("Failed to get league: %v", err),
			})
		}
		return c.JSON(league)
	}
}
