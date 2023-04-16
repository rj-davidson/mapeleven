package initialization

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"mapeleven/models"
	"mapeleven/models/ent"
)

func InitializeData(c *fiber.Ctx, client *ent.Client) error {
	fmt.Println("- Loading data...")

	// Check if any data already exists in the database
	count, err := client.Player.Query().Count(context.Background())
	if err != nil {
		return err
	}

	if count == 0 {
		// Load initial data from the external API
		if err := models.LoadCountriesFromAPI(c, client); err != nil {
			return err
		}
		fmt.Println("    Countries loaded")

		// Load other initial data from the external API...
	} else {
		fmt.Println("- Data already exists in database")
	}

	fmt.Println("- Data loaded")
	return nil
}
