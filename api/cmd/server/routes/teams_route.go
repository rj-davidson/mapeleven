package routes

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/server/serializer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// SetupTeamsRoutes sets up the routes for managing teams.
func SetupTeamsRoutes(app *fiber.App, client *ent.Client) {
	teamSerializer := serializer.NewTeamSerializer(client)

	app.Get(viper.GetString("ENV_PATH")+"/teams", getAllTeams(teamSerializer))
	app.Get(viper.GetString("ENV_PATH")+"/teams/:slug", getTeamBySlug(teamSerializer))
}

func getAllTeams(serializer *serializer.TeamSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		season := c.Query("season", "") // Default to empty if not provided

		teams, err := serializer.GetTeams(context.Background(), season)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": fmt.Sprintf("Failed to get teams: %v", err),
			})
		}
		return c.JSON(teams)
	}
}

func getTeamBySlug(serializer *serializer.TeamSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		season := c.Query("season", "")

		team, err := serializer.GetTeamBySlug(context.Background(), slug, season)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"error": fmt.Sprintf("Team not found: %v", err),
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": fmt.Sprintf("Failed to get team: %v", err),
			})
		}
		return c.JSON(team)
	}
}
