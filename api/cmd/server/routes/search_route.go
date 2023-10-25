package routes

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/server/serializer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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
