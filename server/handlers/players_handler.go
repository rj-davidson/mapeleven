package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"mapeleven/models/ent"
	"mapeleven/models/ent/player"
	"strconv"
)

// GetAllPlayers retrieves all players
func GetAllPlayers(c *fiber.Ctx) error {
	// Get a client to interact with the database
	client := c.Locals("client").(*ent.Client)

	// Get all players from the database
	players, err := client.Player.Query().All(context.Background())
	if err != nil {
		return err
	}

	// Return the players as JSON
	return c.JSON(players)
}

// GetPlayerByID retrieves a player by ID
func GetPlayerByID(c *fiber.Ctx) error {
	// Get a client to interact with the database
	client := c.Locals("client").(*ent.Client)

	// Get the ID parameter from the request
	id := c.Params("id")

	// Convert the ID parameter to an integer
	playerID, err := strconv.Atoi(id)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get the pl with the given ID from the database
	pl, err := client.Player.Query().Where(player.IDEQ(playerID)).Only(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Return the pl as JSON
	return c.JSON(pl)
}
