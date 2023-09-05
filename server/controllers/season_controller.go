package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"mapeleven/models"
	"net/http"
	"time"
)

type SeasonController struct {
	client      *http.Client
	seasonModel *models.SeasonModel
}

type APISeason struct {
	ID      int    `json:"id"`
	Year    int    `json:"year"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Current bool   `json:"current"`
	League  int    `json:"league"`
}

type APISeasonsResponse struct {
	Response []APISeason `json:"response"`
}

func NewSeasonController(seasonModel *models.SeasonModel) *SeasonController {
	return &SeasonController{
		client:      &http.Client{},
		seasonModel: seasonModel,
	}
}

func (sc *SeasonController) FetchSeasonData(data []byte) ([]models.CreateSeasonInput, error) {
	var response APISeasonsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	seasonInputs := make([]models.CreateSeasonInput, len(response.Response))
	for i, season := range response.Response {
		startTime, _ := time.Parse(time.RFC3339, season.Start)
		endTime, _ := time.Parse(time.RFC3339, season.End)
		seasonInputs[i] = models.CreateSeasonInput{
			Year:    season.Year,
			Start:   startTime,
			End:     endTime,
			Current: season.Current,
			League:  season.League,
		}
	}

	return seasonInputs, nil
}

func (sc *SeasonController) UpsertSeason(ctx context.Context, seasonID int) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/seasons?id=%d", seasonID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	resp, err := sc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse the season data from the API response
	seasonInputs, err := sc.FetchSeasonData(data)
	if err != nil {
		return err
	}

	for _, seasonInput := range seasonInputs {
		// Check if the season exists, and either create or update it accordingly
		existingSeason, err := sc.seasonModel.GetSeasonByID(context.Background(), seasonID)
		if existingSeason == nil {
			_, err = sc.seasonModel.CreateSeason(context.Background(), seasonInput)
			if err != nil {
				return err
			}
		} else {
			_, err = sc.seasonModel.UpdateSeason(context.Background(), models.UpdateSeasonInput{
				Year:    &seasonInput.Year,
				Start:   &seasonInput.Start,
				End:     &seasonInput.End,
				Current: &seasonInput.Current,
				League:  &seasonInput.League,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (sc *SeasonController) InitializeSeasons(seasonIDs []int) error {
	// Fetch and upsert seasons
	for _, seasonID := range seasonIDs {
		err := sc.UpsertSeason(context.Background(), seasonID)
		if err != nil {
			return fmt.Errorf("failed to initialize season - error: %v", err)
		}
	}

	fmt.Println("Seasons loaded")
	return nil
}
