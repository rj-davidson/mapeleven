package fetchers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mapeleven/models"
	"mapeleven/models/ent"
	"mapeleven/models/ent/league"
	"net/http"
)

type LeagueFetcher struct {
	client         *http.Client
	leagueModel    *models.LeagueModel
	countryFetcher *CountryFetcher
}

type APILeagueResponse struct {
	Response []struct {
		League  APILeague                 `json:"league"`
		Country models.CreateCountryInput `json:"country"`
	} `json:"response"`
}

type APILeague struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Type league.Type `json:"type"`
	Logo string      `json:"logo"`
}

func NewLeagueFetcher(leagueModel *models.LeagueModel, countryFetcher *CountryFetcher) *LeagueFetcher {
	return &LeagueFetcher{
		client:         &http.Client{},
		leagueModel:    leagueModel,
		countryFetcher: countryFetcher,
	}
}

func (lf *LeagueFetcher) FetchLeagueData(url string) (models.CreateLeagueInput, models.CreateCountryInput, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.CreateLeagueInput{}, models.CreateCountryInput{}, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.CreateLeagueInput{}, models.CreateCountryInput{}, err
	}

	var apiResponse APILeagueResponse
	err = json.Unmarshal(data, &apiResponse)
	if err != nil {
		return models.CreateLeagueInput{}, models.CreateCountryInput{}, err
	}

	inputData := models.CreateLeagueInput{
		ID:      apiResponse.Response[0].League.ID,
		Name:    apiResponse.Response[0].League.Name,
		Type:    apiResponse.Response[0].League.Type,
		Logo:    apiResponse.Response[0].League.Logo,
		Country: apiResponse.Response[0].Country.Code,
	}

	return inputData, apiResponse.Response[0].Country, nil
}

func (lf *LeagueFetcher) UpsertLeague(url string) (*ent.League, error) {
	ctx := context.Background()

	inputData, inputCountry, err := lf.FetchLeagueData(url)
	if err != nil {
		return nil, err
	}

	// Upsert the foreign key (Country) if necessary
	upsertedCountry, err := lf.countryFetcher.FetchCountry(ctx, inputCountry.Code)
	if err != nil {
		return nil, err
	}

	// Use the upserted country's ID for inputData
	inputData.Country = upsertedCountry.Code

	// Check if the league exists
	_, err = lf.leagueModel.GetLeague(inputData.ID)

	if ent.IsNotFound(errors.Unwrap(err)) {
		// Create the league if it doesn't exist
		newLeague, err := lf.leagueModel.CreateLeague(inputData)
		if err != nil {
			return nil, err
		}
		return newLeague, nil
	} else if err == nil {
		// Update the league if it exists
		updateInput := models.UpdateLeagueInput{
			ID:      inputData.ID,
			Name:    &inputData.Name,
			Type:    &inputData.Type,
			Logo:    &inputData.Logo,
			Country: &upsertedCountry.ID,
		}
		updatedLeague, err := lf.leagueModel.UpdateLeague(updateInput)
		if err != nil {
			return nil, err
		}
		return updatedLeague, nil
	} else {
		// Return any other error that may occur
		return nil, err
	}
}
