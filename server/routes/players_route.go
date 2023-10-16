package routes

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"mapeleven/db/ent"
	"mapeleven/serializer"
)

func SetupPlayersRoutes(app *fiber.App, client *ent.Client) {
	playerSerializer := serializer.NewPlayerSerializer(client)

	app.Get(viper.GetString("ENV_PATH")+"/players", getAllPlayers(playerSerializer))
	app.Get(viper.GetString("ENV_PATH")+"/players/:slug", getPlayerBySlug(playerSerializer))
}

func getAllPlayers(serializer *serializer.PlayerSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		players, err := serializer.GetPlayers(context.Background())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": fmt.Sprintf("Failed to get players: %v", err),
			})
		}
		return c.JSON(players)
	}
}

func getPlayerBySlug(serializer *serializer.PlayerSerializer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")

		player, err := serializer.GetPlayerBySlug(slug, context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"error": fmt.Sprintf("Player not found: %v", err),
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": fmt.Sprintf("Failed to get player: %v", err),
			})
		}
		return c.JSON(player)
	}
}
