package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"mapeleven/db/ent"
	"mapeleven/models"
	"strconv"
)

// SetupPlayersRoutes sets up the routes for managing players.
func SetupPlayersRoutes(app *fiber.App, client *ent.Client) {
	playerModel := models.NewPlayerModel(client)

	// Get all players
	app.Get(viper.GetString("ENV_PATH")+"/players", func(c *fiber.Ctx) error {
		players, err := playerModel.ListPlayers(context.Background())
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get players")
		}

		return c.JSON(players)
	})

	// Get a player by FootballApiId
	app.Get(viper.GetString("ENV_PATH")+"/players/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, "Invalid player FootballApiId")
		}

		player, err := playerModel.GetPlayerByID(context.Background(), id)
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusNotFound, "Player not found")
		}

		return c.JSON(player)
	})

	// Get a player by Team FootballApiId

}
