package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"mapeleven/db/ent"
	"mapeleven/models"
	"strconv"
)

// SetupTeamsRoutes sets up the routes for managing teams.
func SetupTeamsRoutes(app *fiber.App, client *ent.Client) {
	teamModel := models.NewTeamModel(client)

	// Get all teams
	app.Get("/teams", func(c *fiber.Ctx) error {
		teams, err := teamModel.ListTeams(context.Background())
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get teams")
		}

		return c.JSON(teams)
	})

	// Get a team by ID
	app.Get("/teams/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, "Invalid team ID")
		}

		team, err := teamModel.GetTeamByID(context.Background(), id)
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusNotFound, "Team not found")
		}

		return c.JSON(team)
	})
}
