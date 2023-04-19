package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"mapeleven/db/ent"
	"mapeleven/db/ent/league"
	"mapeleven/models"
	"mapeleven/utils"
	"net/http"
	"path/filepath"
)

type LeagueController struct {
	client            *http.Client
	leagueModel       *models.LeagueModel
	countryController *CountryController
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

func NewLeagueController(leagueModel *models.LeagueModel) *LeagueController {
	return &LeagueController{
		client:      &http.Client{},
		leagueModel: leagueModel,
	}
}

func (lc *LeagueController) FetchLeagueData(data []byte) (models.CreateLeagueInput, models.CreateCountryInput, error) {
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

func (lc *LeagueController) UpsertLeague(ctx context.Context, leagueID int) (*ent.League, error) {
	// Construct the API URL with the leagueID
	url := fmt.Sprintf("https://api-football-v3.p.rapidapi.com/v3/leagues?id=%d", leagueID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := lc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the league and country data from the API response
	leagueInput, inputCountry, err := lc.parseLeagueData(data)
	if err != nil {
		return nil, err
	}

	// Use the country's name for inputData
	leagueInput.Country = inputCountry.Name

	// Download the logo and set the logoLocation
	err = lc.downloadLeagueLogoIfNeeded(&leagueInput)
	if err != nil {
		return nil, err
	}

	// Check if the league exists, and either create or update it accordingly
	existingLeague, err := lc.leagueModel.GetLeagueByID(context.Background(), leagueInput.ID)
	if existingLeague == nil {
		return lc.createLeague(&leagueInput)
	} else {
		return lc.updateLeague(&leagueInput)
	}
}

func (lc *LeagueController) parseLeagueData(data []byte) (models.CreateLeagueInput, models.CreateCountryInput, error) {
	var response APILeagueResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return models.CreateLeagueInput{}, models.CreateCountryInput{}, err
	}

	leag := response.Response[0].League
	country := response.Response[0].Country
	return leag, country, nil
}

func (lc *LeagueController) downloadLeagueLogoIfNeeded(leagueInput *models.CreateLeagueInput) error {
	if leagueInput.Logo != "" {
		logoLocation, err := utils.DownloadImageIfNeeded(
			leagueInput.Logo,
			fmt.Sprintf("images/leagues/%d%s", leagueInput.ID, filepath.Ext(leagueInput.Logo)),
		)
		if err != nil {
			return err
		}
		leagueInput.Logo = logoLocation
	}
	return nil
}

func (lc *LeagueController) createLeague(leagueInput *models.CreateLeagueInput) (*ent.League, error) {
	newLeague, err := lc.leagueModel.CreateLeague(context.Background(), *leagueInput)
	if err != nil {
		return nil, err
	}
	return newLeague, nil
}

func (lc *LeagueController) updateLeague(leagueInput *models.CreateLeagueInput) (*ent.League, error) {
	updateInput := models.UpdateLeagueInput{
		ID:      leagueInput.ID,
		Name:    &leagueInput.Name,
		Type:    &leagueInput.Type,
		Logo:    &leagueInput.Logo,
		Country: &leagueInput.Country,
	}
	updatedLeague, err := lc.leagueModel.UpdateLeague(context.Background(), updateInput)
	if err != nil {
		return nil, err
	}
	return updatedLeague, nil
}

func (lc *LeagueController) InitializeLeagues(leagueIDs []int) error {
	// Fetch and upsert leagues
	for _, id := range leagueIDs {
		leag, err := lc.UpsertLeague(context.Background(), id)
		if err != nil {
			return fmt.Errorf("failed to upsert league with ID %d: %v", id, err)
		}
		fmt.Printf("Upserted league with ID %d: %+v\n", id, leag)
	}

	fmt.Println("Leagues loaded")
	return nil
}
