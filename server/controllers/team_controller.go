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
)

type TeamController struct {
	client            *http.Client
	teamModel         *models.TeamModel
	countryController *CountryController
}

type APITeamResponse struct {
	Response []struct {
		Team    models.CreateTeamInput    `json:"team"`
		Country models.CreateCountryInput `json:"country"`
	} `json:"response"`
}

type APITeam struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Founded  int    `json:"founded"`
	National bool   `json:"national"`
	Logo     string `json:"logo"`
}

func NewTeamController(teamModel *models.TeamModel, countryController *CountryController) *TeamController {
	return &TeamController{
		client:            &http.Client{},
		teamModel:         teamModel,
		countryController: countryController,
	}
}

func (tc *TeamController) FetchTeamData(data []byte) (models.CreateTeamInput, models.CreateCountryInput, error) {
	var response struct {
		Response []struct {
			Team    json.RawMessage `json:"team"`
			Country json.RawMessage `json:"country"`
		} `json:"response"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return models.CreateTeamInput{}, models.CreateCountryInput{}, err
	}

	team := APITeam{}
	if err := json.Unmarshal(response.Response[0].Team, &team); err != nil {
		return models.CreateTeamInput{}, models.CreateCountryInput{}, err
	}
	teamInput := models.CreateTeamInput{
		ID:       team.ID,
		Name:     team.Name,
		Code:     team.Code,
		Founded:  team.Founded,
		National: team.National,
		Logo:     team.Logo,
	}

	countryInput := models.CreateCountryInput{}
	if err := json.Unmarshal(response.Response[0].Country, &countryInput); err != nil {
		return models.CreateTeamInput{}, models.CreateCountryInput{}, err
	}

	return teamInput, countryInput, nil
}

func (tc *TeamController) GetTeamData(teamID int) ([]byte, error) {
	// Construct the API URL with the teamID
	url := fmt.Sprintf("https://api-football-v3.p.rapidapi.com/v3/teams?id=%d", teamID)

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := tc.client.Do(req)
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

func (tc *TeamController) CreateOrUpdateTeam(teamID int) (*ent.Team, error) {
	data, err := tc.GetTeamData(teamID)
	if err != nil {
		return nil, err
	}

	inputData, inputCountry, err := tc.FetchTeamData(data)
	if err != nil {
		return nil, err
	}

	upsertedCountry, err := tc.countryController.UpsertCountry(inputCountry)
	if err != nil {
		return nil, err
	}

	// Use the upserted country's code
	inputData.Country = upsertedCountry.Code

	// Download the logo and set the logoLocation
	if inputData.Logo != "" {
		inputData.Logo, _ = utils.DownloadImageIfNeeded(
			inputData.Logo,
			fmt.Sprintf("images/teams/%d%s", inputData.ID, filepath.Ext(inputData.Logo)),
		)
	}

	// Check if the team exists
	_, err = tc.teamModel.GetTeamByID(context.Background(), inputData.ID)
	if ent.IsNotFound(err) {
		// Create the team if it doesn't exist
		newTeam, err := tc.teamModel.CreateTeam(context.Background(), inputData)
		if err != nil {
			return nil, err
		}
		return newTeam, nil
	} else if err == nil {
		// Update the team if it exists
		updateInput := models.UpdateTeamInput{
			ID:       inputData.ID,
			Name:     &inputData.Name,
			Code:     &inputData.Code,
			Founded:  &inputData.Founded,
			National: &inputData.National,
			Logo:     &inputData.Logo,
			Country:  &upsertedCountry.Code,
		}
		updatedTeam, err := tc.teamModel.UpdateTeam(context.Background(), updateInput)
		if err != nil {
			return nil, err
		}
		return updatedTeam, nil
	} else {
		// Return any other error that may occur
		return nil, err
	}
}
