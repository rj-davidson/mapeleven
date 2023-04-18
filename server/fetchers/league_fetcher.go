package fetchers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"mapeleven/models"
	"mapeleven/models/ent"
	"mapeleven/models/ent/league"
	"mapeleven/utils"
	"net/http"
	"path/filepath"
)

type LeagueFetcher struct {
	client         *http.Client
	leagueModel    *models.LeagueModel
	countryFetcher *CountryFetcher
}

type APILeagueResponse struct {
	Response []struct {
		League  models.CreateLeagueInput  `json:"league"`
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

func (lf *LeagueFetcher) FetchLeagueData(data []byte) (models.CreateLeagueInput, models.CreateCountryInput, error) {
	var response struct {
		Response []struct {
			League  json.RawMessage `json:"league"`
			Country json.RawMessage `json:"country"`
		} `json:"response"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return models.CreateLeagueInput{}, models.CreateCountryInput{}, err
	}

	lg := APILeague{}
	if err := json.Unmarshal(response.Response[0].League, &lg); err != nil {
		return models.CreateLeagueInput{}, models.CreateCountryInput{}, err
	}
	leagueInput := models.CreateLeagueInput{
		ID:   lg.ID,
		Name: lg.Name,
		Type: lg.Type,
		Logo: lg.Logo,
	}

	co := models.CreateCountryInput{}
	if err := json.Unmarshal(response.Response[0].Country, &co); err != nil {
		return models.CreateLeagueInput{}, models.CreateCountryInput{}, err
	}

	return leagueInput, co, nil
}

func (lf *LeagueFetcher) UpsertLeague(leagueID int) (*ent.League, error) {
	// Construct the API URL with the leagueID
	url := fmt.Sprintf("https://api-football-v3.p.rapidapi.com/v3/leagues?id=%d", leagueID)

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := lf.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	inputData, inputCountry, err := lf.FetchLeagueData(data)
	if err != nil {
		return nil, err
	}

	upsertedCountry, err := lf.countryFetcher.UpsertCountry(inputCountry)
	if err != nil {
		return nil, err
	}

	// Use the upserted country's ID for inputData
	inputData.Country = upsertedCountry.Code

	// Download the logo and set the logoLocation
	if inputData.Logo != "" {
		inputData.Logo, _ = utils.DownloadImageIfNeeded(
			inputData.Logo,
			fmt.Sprintf("images/leagues/%d%s", inputData.ID, filepath.Ext(inputData.Logo)),
		)
	}

	// Check if the league exists
	_, err = lf.leagueModel.GetLeague(context.Background(), inputData.ID)
	if ent.IsNotFound(err) {
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
