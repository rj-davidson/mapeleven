package models

import (
	"context"
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"io"
	"mapeleven/models/ent"
	"mapeleven/models/ent/country"
	"net/http"
	"os"
)

func LoadCountriesFromAPI(c *fiber.Ctx, client *ent.Client) error {
	url := "https://api-football-v1.p.rapidapi.com/v3/countries"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", viper.GetString("API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Unmarshal response body into CountryResponse struct
	var countryResponse CountryResponse
	err = json.Unmarshal(body, &countryResponse)
	if err != nil {
		return err
	}

	for _, apiCountry := range countryResponse.Response {
		// Check if country already exists in the database
		existingCountry, err := client.Country.Query().Where(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(country.FieldCode), apiCountry.Code))
		}).Only(context.Background())

		if err != nil && !ent.IsNotFound(err) {
			return err
		}

		if existingCountry != nil {
			continue
		}

		// Check if flag image already exists
		fileName := fmt.Sprintf("%s.svg", apiCountry.Code)
		filePath := fmt.Sprintf("public/images/flags/%s", fileName)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// Download flag image
			resp, err := http.Get(apiCountry.Flag)
			println(apiCountry.Flag)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// Save flag image to public/images directory
			file, err := os.Create(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, resp.Body)
			if err != nil {
				return err
			}
		}

		// Create new Country entity with flag field set to file path
		_, err = client.Country.
			Create().
			SetName(apiCountry.Name).
			SetCode(apiCountry.Code).
			SetFlag(filePath).
			Save(context.Background())
		if err != nil {
			return err
		}
	}

	c.Type("application/json", "utf-8")
	return c.Send(body)
}

func LoadLeaguesFromAPI(c *fiber.Ctx, client *ent.Client) error {
	url := "https://api-football-v1.p.rapidapi.com/v3/teams?id=33"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", viper.GetString("API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// TODO: Process the response data and load it into the database using the `client` instance

	c.Type("application/json", "utf-8")
	return c.Send(body)
}

func LoadTeamsFromAPI() error {
	// TODO: Make a request to the Teams API endpoint and load the data into the database
	return nil
}

func LoadPlayersFromAPI() error {
	// TODO: Make a request to the Players API endpoint and load the data into the database
	return nil
}

// CountryResponse holds the schema definition for the API response containing country data.
type CountryResponse struct {
	Response []struct {
		Name string `json:"name"`
		Code string `json:"code"`
		Flag string `json:"flag"`
	} `json:"response"`
}

func LoadStandingsFromAPI() error {
	// TODO: Make a request to the Standings API endpoint and load the data into the database
	return nil
}
