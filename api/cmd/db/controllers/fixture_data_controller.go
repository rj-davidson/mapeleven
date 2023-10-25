package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/fixture_models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type FixtureResponse struct {
	Results  int                             `json:"results"`
	Response []fixture_models.FixtureDetails `json:"response"`
}

type FixtureDataController struct {
	client       *http.Client
	fixtureModel *fixture_models.FixtureModel
}

func NewFixtureDataController(fixtureModel *fixture_models.FixtureModel) *FixtureDataController {
	return &FixtureDataController{
		client:       &http.Client{},
		fixtureModel: fixtureModel,
	}
}

func (fdc *FixtureDataController) FetchFixtures(ctx context.Context) error {
	fs, err := fdc.fixtureModel.ListFixtures(ctx)
	if err != nil {
		return err
	}
	for _, f := range fs {
		fixtureDetails, err := fdc.fetchFixtureByID(ctx, f.ApiFootballId)
		if err != nil {
			fmt.Printf("Error fetching fixture: %v", err)
		}
		err = fdc.writeFixtureData(f.ApiFootballId, *fixtureDetails, ctx)
		if err != nil {
			fmt.Printf("Error writing fixture data: %v", err)
		}
	}
	return nil
}

func (fdc *FixtureDataController) fetchFixtureByID(ctx context.Context, fixtureID int) (*fixture_models.FixtureDetails, error) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?id=%d", fixtureID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	resp, err := fdc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fixtureResponse FixtureResponse
	if err := json.Unmarshal(body, &fixtureResponse); err != nil {
		return nil, err
	}

	if fixtureResponse.Results == 0 {
		return nil, fmt.Errorf("no data available for fixtureID: %d", fixtureID)
	}

	return &fixtureResponse.Response[0], nil
}

func (fdc *FixtureDataController) writeFixtureData(apiFootballId int, data fixture_models.FixtureDetails, ctx context.Context) error {
	err := fdc.fixtureModel.UpsertFixtureData(apiFootballId, data, ctx)
	if err != nil {
		fmt.Printf("Error upserting fixture data: %v", err)
	}
	return nil
}
