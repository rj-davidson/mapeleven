package fetchers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mapeleven/models"
	"mapeleven/models/ent"
	"net/http"
)

type TeamFetcher struct {
	client         *http.Client
	teamModel      *models.TeamModel
	countryFetcher *CountryFetcher
}

type APITeamResponse struct {
	Response []struct {
		Team models.CreateTeamInput `json:"team"`
	} `json:"response"`
}

func NewTeamFetcher(teamModel *models.TeamModel, countryFetcher *CountryFetcher) *TeamFetcher {
	return &TeamFetcher{
		client:         &http.Client{},
		teamModel:      teamModel,
		countryFetcher: countryFetcher,
	}
}

func (tf *TeamFetcher) FetchTeamData(teamID int) (models.CreateTeamInput, error) {
	// Fetch data from API using teamID
	data, err := tf.getAPITeamData(teamID)
	if err != nil {
		return models.CreateTeamInput{}, err
	}

	var response APITeamResponse

	if err := json.Unmarshal(data, &response); err != nil {
		return models.CreateTeamInput{}, err
	}

	teamInput := response.Response[0].Team

	return teamInput, nil
}

func (tf *TeamFetcher) getAPITeamData(teamID int) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.example.com/teams/%d", teamID) // Replace with your actual API URL
	req, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer fda98bbd4fmshf7ac8b80f84ef2cp1d3f5bjsnfc890892cfe4")

	resp, err := tf.client.Do(req)
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

func (tf *TeamFetcher) UpsertTeam(ctx context.Context, teamID int) (*ent.Team, error) {
	// Fetch team data from API
	teamInput, err := tf.FetchTeamData(teamID)
	if err != nil {
		return nil, err
	}

	// Upsert Country
	upsertedCountry, err := tf.countryFetcher.UpsertCountry(teamInput.Country)
	if err != nil {
		return nil, err
	}

	// Set the upserted country's ID for teamInput
	teamInput.Country = upsertedCountry.ID

	// Check if the team exists
	_, err = tf.teamModel.GetTeamByID(ctx, teamInput.ID)
	if ent.IsNotFound(err) {
		// Create the team if it doesn't exist
		newTeam, err := tf.teamModel.CreateTeam(ctx, teamInput)
		if err != nil {
			return nil, err
		}
		return newTeam, nil
	} else if err == nil {
		// Update the team if it exists
		updateInput := models.UpdateTeamInput{
			ID:       teamInput.ID,
			Name:     &teamInput.Name,
			Code:     &teamInput.Code,
			Founded:  &teamInput.Founded,
			National: &teamInput.National,
			Logo:     &teamInput.Logo,
			Country:  &teamInput.Country,
		}
		updatedTeam, err := tf.teamModel.UpdateTeam(ctx, updateInput)
		if err != nil {
			return nil, err
		}
		return updatedTeam, nil
	} else {
		// Return any other error that may occur
		return nil, err
	}
}
