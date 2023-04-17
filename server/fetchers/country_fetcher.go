package fetchers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mapeleven/models"
	"mapeleven/models/ent"
	"mapeleven/utils"
	"net/http"
	"os"
)

type CountryFetcher struct {
	apiKey         string
	client         *http.Client
	countryModel   *models.CountryModel
	entClient      *ent.Client
	apiBaseURL     string
	apiCountryPath string
}

type APICountryResponse struct {
	Response []APICountry `json:"response"`
}

type APICountry struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

func NewCountryFetcher(apiKey string, countryModel *models.CountryModel, entClient *ent.Client) *CountryFetcher {
	return &CountryFetcher{
		apiKey:         apiKey,
		client:         &http.Client{},
		countryModel:   countryModel,
		entClient:      entClient,
		apiBaseURL:     "https://api-football-v1.p.rapidapi.com",
		apiCountryPath: "/v3/countries",
	}
}

func (cf *CountryFetcher) FetchCountry(ctx context.Context, countryCode string) (*ent.Country, error) {
	// Search for country by code in Ent DB
	country, err := cf.countryModel.GetCountryByCode(countryCode)

	if err != nil {
		// Country does not exist, fetch from API
		apiCountry, err := cf.fetchCountryByCode(ctx, countryCode)
		if err != nil {
			return nil, err
		}

		filePath := fmt.Sprintf("public/images/flags/%s.svg", apiCountry.Code)
		err = utils.DownloadFile(apiCountry.Flag, filePath)
		if err != nil {
			return nil, err
		}

		// Create country in Ent DB
		createInput := &models.CreateCountryInput{
			Name: apiCountry.Name,
			Code: apiCountry.Code,
			Flag: filePath,
		}
		country, err = cf.countryModel.CreateCountry(*createInput)
		if err != nil {
			return nil, err
		}
	} else {
		// Country exists, update flag if necessary
		flagPath := fmt.Sprintf("public/images/flags/%s.svg", country.Code)
		if _, err := os.Stat(flagPath); os.IsNotExist(err) {
			// Flag file does not exist, download from API
			err = utils.DownloadFile(country.Flag, flagPath)
			if err != nil {
				return nil, err
			}

			// Update country in Ent DB with new flag path
			updateInput := &models.UpdateCountryInput{
				Code: country.Code,
				Flag: string(flagPath),
			}

			_, err = cf.countryModel.UpdateCountry(*updateInput)
			if err != nil {
				return nil, err
			}

			// Reload country from Ent DB to get updated flag path
			country, err = cf.countryModel.GetCountryByCode(countryCode)
			if err != nil {
				return nil, err
			}
		}
	}

	return country, nil
}

func (cf *CountryFetcher) fetchCountryByCode(ctx context.Context, code string) (*APICountry, error) {
	url := fmt.Sprintf("%s%s?code=%s", cf.apiBaseURL, cf.apiCountryPath, code)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", cf.apiKey)

	resp, err := cf.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse APICountryResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	if len(apiResponse.Response) == 0 {
		return nil, fmt.Errorf("country not found")
	}

	return &apiResponse.Response[0], nil
}
