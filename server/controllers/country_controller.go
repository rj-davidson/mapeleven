package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"mapeleven/db/ent"
	"mapeleven/models"
	"mapeleven/utils"
	"net/http"
	"path/filepath"
	"strings"
)

type CountryController struct {
	client       *http.Client
	countryModel *models.CountryModel
	entClient    *ent.Client
}

func NewCountryController(countryModel *models.CountryModel, client *ent.Client) *CountryController {
	return &CountryController{
		client:       &http.Client{},
		countryModel: countryModel,
		entClient:    client,
	}
}

func (cc *CountryController) UpsertCountry(c models.CreateCountryInput) (*ent.Country, error) {
	// Convert the code to uppercase if it is not empty
	c.Code = strings.ToUpper(c.Code)

	// If the code is empty, use the first 3 letters of the name
	if c.Code == "" {
		c.Code = strings.ToUpper(c.Name[:3])
	}

	// Download the flag and set the local path to the flag
	if c.Flag != "" {
		c.Flag, _ = utils.DownloadImageIfNeeded(c.Flag,
			fmt.Sprintf("images/flags/%s%s", c.Code, filepath.Ext(c.Flag)))
	}

	// Search for country by code in Ent DB
	country, _ := cc.countryModel.GetCountryByName(c.Name)

	if country == nil {
		// Country does not exist, create it
		createInput := &models.CreateCountryInput{
			Name: c.Name,
			Code: c.Code,
			Flag: c.Flag,
		}
		country, _ = cc.countryModel.CreateCountry(*createInput)
	} else {
		// Country exists, update it
		updateInput := &models.UpdateCountryInput{
			Code: c.Code,
			Name: c.Name,
			Flag: c.Flag,
		}
		country, _ = cc.countryModel.UpdateCountry(*updateInput)
	}

	return country, nil
}

func (cc *CountryController) FetchCountryByName(countryName string) (models.CreateCountryInput, error) {
	// Fetch the country data from the API
	data, err := cc.GetCountryData(countryName)
	if err != nil {
		fmt.Printf("Error fetching country data: %s     Error: %s", countryName, err.Error())
	}

	// Decode the API response into a map
	var resp map[string]interface{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("Error decoding country data: %s     Error: %s", countryName, err.Error())
	}

	// Extract the country
	c := resp["response"].([]interface{})[0].(map[string]interface{})

	// Decode the first object into a struct
	apiCountry := models.CreateCountryInput{
		Name: c["name"].(string),
		Code: c["code"].(string),
		Flag: c["flag"].(string),
	}

	// Check if country exists in the Ent database
	entCountry, err := cc.countryModel.GetCountryByName(apiCountry.Name)
	if entCountry != nil {
		return apiCountry, nil
	} else {
		// Country does not exist in the Ent database, create it
		_, err = cc.UpsertCountry(apiCountry)
	}

	return apiCountry, nil
}

func (cc *CountryController) GetCountryData(countryName string) ([]byte, error) {
	// Construct the API URL with the country name
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/countries?name=%s", strings.ToLower(countryName))

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := cc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
