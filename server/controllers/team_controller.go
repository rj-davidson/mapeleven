package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"mapeleven/models"
	"mapeleven/utils"
	"net/http"
	"path/filepath"
)

type TeamController struct {
	client    *http.Client
	teamModel *models.TeamModel
}

type APITeam struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Founded  int    `json:"founded"`
	National bool   `json:"national"`
	Logo     string `json:"logo"`
	Country  struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"country"`
}

type APITeamsResponse struct {
	Response []APITeam `json:"response"`
}

func NewTeamController(teamModel *models.TeamModel) *TeamController {
	return &TeamController{
		client:    &http.Client{},
		teamModel: teamModel,
	}
}

func (tc *TeamController) FetchTeamData(data []byte) ([]models.CreateTeamInput, error) {
	var response APITeamsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	teamInputs := make([]models.CreateTeamInput, len(response.Response))
	for i, team := range response.Response {
		teamInputs[i] = models.CreateTeamInput{
			ID:       team.ID,
			Name:     team.Name,
			Code:     team.Code,
			Founded:  team.Founded,
			National: team.National,
			Logo:     team.Logo,
			Country:  team.Country.Name,
		}
	}

	return teamInputs, nil
}

func (tc *TeamController) UpsertTeam(ctx context.Context, teamID int) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams?id=%d", teamID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	resp, err := tc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse the team data from the API response
	teamInput, err := tc.parseTeamData(data)
	if err != nil {
		return err
	}

	// Download the logo and set the logoLocation
	err = tc.downloadTeamLogoIfNeeded(teamInput)
	if err != nil {
		return err
	}

	// Check if the team exists, and either create or update it accordingly
	existingTeam, err := tc.teamModel.GetTeamByID(context.Background(), teamInput.ID)
	if existingTeam == nil {
		_, err = tc.teamModel.CreateTeam(context.Background(), *teamInput)
		if err != nil {
			return err
		}
	} else {
		_, err = tc.teamModel.UpdateTeam(context.Background(), models.UpdateTeamInput{
			ID:       teamInput.ID,
			Name:     &teamInput.Name,
			Code:     &teamInput.Code,
			Founded:  &teamInput.Founded,
			National: &teamInput.National,
			Logo:     &teamInput.Logo,
			Country:  &teamInput.Country,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (tc *TeamController) parseTeamData(data []byte) (*models.CreateTeamInput, error) {
	var response struct {
		Response []struct {
			Team models.CreateTeamInput `json:"team"`
		} `json:"response"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	// Extract the first team from the response
	r := response.Response[0]
	teamInput := r.Team

	return &teamInput, nil
}

func (tc *TeamController) downloadTeamLogoIfNeeded(teamInput *models.CreateTeamInput) error {
	if teamInput.Logo != "" {
		logoLocation, err := utils.DownloadImageIfNeeded(
			teamInput.Logo,
			fmt.Sprintf("images/teams/%d%s", teamInput.ID, filepath.Ext(teamInput.Logo)),
		)
		if err != nil {
			return err
		}
		teamInput.Logo = logoLocation
	}
	return nil
}

func (tc *TeamController) InitializeTeams(teamIDs []int) error {
	// Fetch and upsert teams
	for _, teamID := range teamIDs {
		err := tc.UpsertTeam(context.Background(), teamID)
		if err != nil {
			return fmt.Errorf("failed to initialize team - error: %v", err)
		}
	}

	fmt.Println("Teams loaded")
	return nil
}
