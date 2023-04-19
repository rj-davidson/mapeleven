package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"mapeleven/db/ent"
	"mapeleven/models"
	"strconv"
)

func SetupLeaguesRoutes(app *fiber.App, client *ent.Client) {
	leagueModel := models.NewLeagueModel(client)
	// Get all leagues
	app.Get("/leagues", func(c *fiber.Ctx) error {
		leagues, err := leagueModel.ListLeagues(context.Background())
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get leagues")
		}
		return c.JSON(leagues)
	})

	// Get a specific league by ID
	app.Get("/leagues/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid ID parameter")
		}

		league, err := leagueModel.GetLeague(context.Background(), id)
		if err != nil {
			if ent.IsNotFound(err) {
				return fiber.NewError(fiber.StatusNotFound, "League not found")
			}
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get league")
		}

		return c.JSON(league)
	})
}
