package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/utils"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

type LeagueController struct {
	client      *http.Client
	leagueModel *models.LeagueModel
	seasonModel *models.SeasonModel
}

type APILeagueResponse struct {
	Response []struct {
		League  models.CreateLeagueInput   `json:"league"`
		Country models.CreateCountryInput  `json:"country"`
		Seasons []models.CreateSeasonInput `json:"seasons"`
	} `json:"response"`
}

type APILeague struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Type league.Type `json:"type"`
	Logo string      `json:"logo"`
}

func NewLeagueController(leagueModel *models.LeagueModel, seasonModel *models.SeasonModel) *LeagueController {
	return &LeagueController{
		client:      &http.Client{},
		leagueModel: leagueModel,
		seasonModel: seasonModel,
	}
}

func (lc *LeagueController) InitializeLeagues(leagueIDs []int, ctx context.Context) error {
	fmt.Println("Initializing leagues...")
	for _, leagueID := range leagueIDs {
		fmt.Println("Initializing league with FootballApiId: ", leagueID)
		err := lc.fetchLeagueByID(ctx, leagueID)
		if err != nil {
			fmt.Errorf("fetch league with FootballApiId %d: %w", leagueID, err)
		}
	}
	fmt.Println("[ Finished initializing leagues ]")
	return nil
}

func (lc *LeagueController) fetchLeagueByID(ctx context.Context, leagueID int) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/leagues?id=%d", leagueID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", viper.GetString("API_KEY"))

	resp, err := lc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return lc.parseLeagueResponse(ctx, data)
}

func (lc *LeagueController) parseLeagueResponse(ctx context.Context, data []byte) error {
	var response APILeagueResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	leagueInput := response.Response[0].League
	leagueInput.Country = response.Response[0].Country.Name
	var seasonInputs []models.CreateSeasonInput
	for _, season := range response.Response[0].Seasons {
		seasonInputs = append(seasonInputs, season)
	}

	// Download the logo and set the logoLocation
	err := lc.downloadLeagueLogoIfNeeded(&leagueInput)
	if err != nil {
		return err
	}

	l, err := lc.upsertLeague(ctx, &leagueInput)
	if err != nil {
		return err
	}

	for _, seasonInput := range seasonInputs {
		_, err := lc.upsertSeason(ctx, &seasonInput,
			utils.Slugify(leagueInput.Name+"-"+strconv.Itoa(seasonInput.Year)),
			l.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (lc *LeagueController) upsertLeague(ctx context.Context, leagueInput *models.CreateLeagueInput) (*ent.League, error) {
	existingLeague := lc.leagueModel.LeagueExistsByFootballApiId(ctx, leagueInput.FootballApiId)
	var l *ent.League
	err := error(nil)
	if existingLeague == nil {
		l, err = lc.leagueModel.CreateLeague(ctx, *leagueInput)
		if err != nil {
			return nil, err
		}
	} else {
		l, err = lc.leagueModel.UpdateLeague(ctx, models.UpdateLeagueInput{
			ID:   existingLeague.ID,
			Name: &leagueInput.Name,
			Type: &leagueInput.Type,
			Logo: &leagueInput.Logo,
		})
		if err != nil {
			return nil, err
		}
	}
	return l, nil
}

func (lc *LeagueController) upsertSeason(ctx context.Context, seasonInput *models.CreateSeasonInput, slug string, leagueID int) (*ent.Season, error) {
	existingSeason := lc.seasonModel.SeasonExistsByLeagueIdAndYear(ctx, leagueID, seasonInput.Year)
	if existingSeason == nil {
		return lc.seasonModel.CreateSeason(ctx, *seasonInput, slug, leagueID)
	} else {
		return lc.seasonModel.UpdateSeason(ctx, models.UpdateSeasonInput{
			ID:      existingSeason.ID,
			Start:   seasonInput.Start,
			End:     seasonInput.End,
			Current: seasonInput.Current,
		})
	}
}

func (lc *LeagueController) downloadLeagueLogoIfNeeded(leagueInput *models.CreateLeagueInput) error {
	if leagueInput.Logo != "" {
		logoLocation, err := utils.DownloadImageIfNeeded(
			leagueInput.Logo,
			fmt.Sprintf("images/leagues/%d%s", leagueInput.FootballApiId, filepath.Ext(leagueInput.Logo)),
		)
		if err != nil {
			return err
		}
		leagueInput.Logo = logoLocation
	}
	return nil
}
