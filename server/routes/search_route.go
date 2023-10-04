package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"mapeleven/db/ent"
	"mapeleven/serializer"
)

func SetupSearchRoutes(app *fiber.App, client *ent.Client) {
	app.Get(viper.GetString("ENV_PATH")+"/search", func(c *fiber.Ctx) error {
		searchItems, err := serializer.SearchSerializer(client)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get search items")
		}
		return c.JSON(searchItems)
	})

}
