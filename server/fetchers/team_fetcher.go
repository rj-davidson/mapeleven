package fetchers

import (
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

// Remove VenueModel from the constructor
func NewTeamFetcher(teamModel *models.TeamModel, countryFetcher *CountryFetcher) *TeamFetcher {
	return &TeamFetcher{
		client:         &http.Client{},
		teamModel:      teamModel,
		countryFetcher: countryFetcher,
	}
}

func (tf *TeamFetcher) FetchTeamData(data []byte) (models.CreateTeamInput, error) {
	var response APITeamResponse

	if err := json.Unmarshal(data, &response); err != nil {
		return models.CreateTeamInput{}, err
	}

	teamInput := response.Response[0].Team

	return teamInput, nil
}

// Other functions ...

func (tf *TeamFetcher) UpsertTeam(teamID int) (*ent.Team, error) {
	// Fetch team data from API
	teamInput, err := tf.fetchTeamData(teamID)
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
