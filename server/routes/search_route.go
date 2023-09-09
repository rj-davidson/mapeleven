package routes

import (
	"github.com/gofiber/fiber/v2"
	"mapeleven/db/ent"
	"mapeleven/serializer"
)

func SetupSearchRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/search", func(c *fiber.Ctx) error {
		searchItems, err := serializer.SearchSerializer(client)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get search items")
		}
		return c.JSON(searchItems)
	})

}
