package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/utils"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type CountryController struct {
	httpClient   *http.Client
	countryModel *models.CountryModel
}

func NewCountryController(entClient *ent.Client) *CountryController {
	cm := models.NewCountryModel(entClient)
	return &CountryController{
		httpClient:   http.DefaultClient,
		countryModel: cm,
	}
}

func (cc *CountryController) UpsertCountry(c models.CreateCountryInput) (ent.Country, error) {
	c.Code = strings.ToUpper(c.Code)
	if c.Code == "" {
		c.Code = strings.ToUpper(c.Name[:3])
	}
	if c.Flag != "" {
		c.Flag, _ = utils.DownloadImageIfNeeded(c.Flag,
			fmt.Sprintf("images/flags/%s%s", c.Code, filepath.Ext(c.Flag)))
	}
	country, err := cc.countryModel.GetCountryByName(c.Name)
	if country == nil {
		fmt.Printf("DB[countries] -> Adding <   %s   > \n", c.Name)
		createInput := &models.CreateCountryInput{
			Name: c.Name,
			Code: c.Code,
			Flag: c.Flag,
		}
		country, err = cc.countryModel.CreateCountry(*createInput)
		if err != nil {
			return ent.Country{}, err
		}
	} else {
		updateInput := &models.UpdateCountryInput{
			Code: c.Code,
			Name: c.Name,
			Flag: c.Flag,
		}
		country, err = cc.countryModel.UpdateCountry(*updateInput)
		if err != nil {
			return ent.Country{}, err
		}
	}

	return *country, nil
}

func (cc *CountryController) InitializeCountries() error {
	// Fetch all country data from the API
	data, err := cc.GetAllCountryData(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching country data: %s", err.Error())
	}

	// Decode the API response into a map
	var resp map[string]interface{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return fmt.Errorf("error decoding country data: %s", err.Error())
	}

	// Extract the countries
	response := resp["response"].([]interface{})
	if len(response) == 0 {
		return fmt.Errorf("no countries found")
	}

	for _, countryData := range response {
		// Decode each country object into a struct
		c := countryData.(map[string]interface{})
		apiCountry := models.CreateCountryInput{
			Name: c["name"].(string),
			Code: "",
			Flag: "",
		}
		if code, ok := c["code"].(string); ok && c["code"] != nil {
			apiCountry.Code = code
		}
		if flag, ok := c["flag"].(string); ok && c["flag"] != nil {
			apiCountry.Flag = flag
		}

		// Upsert Country
		_, err := cc.UpsertCountry(apiCountry)
		if err != nil {
			return fmt.Errorf("error upserting country: %s, error: %s", apiCountry.Name, err.Error())
		}
	}
	return nil
}

func (cc *CountryController) GetAllCountryData(ctx context.Context) ([]byte, error) {
	// Construct the API URL to fetch all countries
	url := "https://api-football-v1.p.rapidapi.com/v3/countries"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := cc.httpClient.Do(req)
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
