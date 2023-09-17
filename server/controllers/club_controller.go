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

type ClubController struct {
	client    *http.Client
	clubModel *models.ClubModel
}

type APIClub struct {
	Response []struct {
		Team models.Club `json:"team"`
	} `json:"response"`
}

func NewClubController(clubModel *models.ClubModel) *ClubController {
	return &ClubController{
		client:    &http.Client{},
		clubModel: clubModel,
	}
}

func (cc *ClubController) EnsureClubExists(ctx context.Context, apiFootballId int) error {
	exists := cc.clubModel.Exists(ctx, apiFootballId)
	if !exists {
		err := cc.fetchClubByID(ctx, apiFootballId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cc *ClubController) fetchClubByID(ctx context.Context, clubID int) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams?id=%d", clubID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	resp, err := cc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return cc.parseClubResponse(ctx, data)
}

func (cc *ClubController) parseClubResponse(ctx context.Context, data []byte) error {
	var response APIClub
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	c := response.Response[0].Team
	if err := cc.downloadLogoIfNeeded(&c); err != nil {
		return err
	}

	err := cc.upsertClub(ctx, &c)
	if err != nil {
		return err
	}
	return nil
}

func (cc *ClubController) upsertClub(ctx context.Context, c *models.Club) error {
	if !cc.clubModel.Exists(ctx, c.ApiFootballId) {
		_, err := cc.clubModel.CreateClub(ctx, *c)
		if err != nil {
			return err
		}
	} else {
		err := cc.clubModel.UpdateClub(ctx, *c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cc *ClubController) downloadLogoIfNeeded(c *models.Club) error {
	if c.Logo != "" {
		logoLocation, err := utils.DownloadImageIfNeeded(
			c.Logo,
			fmt.Sprintf("images/teams/%d%s", c.ApiFootballId, filepath.Ext(c.Logo)),
		)
		if err != nil {
			return err
		}
		c.Logo = logoLocation
	}
	return nil
}
