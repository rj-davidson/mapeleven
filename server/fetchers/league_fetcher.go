package fetchers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mapeleven/models"
	"mapeleven/models/ent/league"
	"net/http"
	"strings"
)

type LeagueFetcher struct {
	apiKey        string
	client        *http.Client
	leagueModel   *models.LeagueModel
	countryModel  *models.CountryModel
	seasonModel   *models.SeasonModel
	apiBaseURL    string
	apiLeaguePath string
}

type APIResponse struct {
	Results int         `json:"results"`
	Leagues []APILeague `json:"leagues"`
}

type APILeague struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Logo    string `json:"logo"`
	Country struct {
		ID int `json:"id"`
	} `json:"country"`
}

func NewLeagueFetcher(apiKey string, leagueModel *models.LeagueModel, countryModel *models.CountryModel, seasonModel *models.SeasonModel) *LeagueFetcher {
	return &LeagueFetcher{
		apiKey:        apiKey,
		client:        &http.Client{},
		leagueModel:   leagueModel,
		countryModel:  countryModel,
		seasonModel:   seasonModel,
		apiBaseURL:    "https://api-football-v1.p.rapidapi.com",
		apiLeaguePath: "/v3/leagues",
	}
}

func (lf *LeagueFetcher) FetchLeagues(leagueIDs []int) error {
	for _, id := range leagueIDs {
		leag, err := lf.fetchLeague(id)
		if err != nil {
			return err
		}

		input := models.CreateLeagueInput{
			ID:      leag.ID,
			Name:    leag.Name,
			Type:    league.Type(leag.Type),
			Logo:    leag.Logo,
			Country: leag.Country.ID,
		}

		_, err = lf.leagueModel.CreateLeague(input)
		if err != nil {
			return err
		}
	}
	return nil
}

func (lf *LeagueFetcher) fetchLeague(id int) (*APILeague, error) {
	url := fmt.Sprintf("%s%s?id=%d", lf.apiBaseURL, lf.apiLeaguePath, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", lf.apiKey)

	resp, err := lf.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	if len(apiResponse.Leagues) == 0 {
		return nil, fmt.Errorf("league not found")
	}

	l := apiResponse.Leagues[0]
	l.Type = string(leagueType(l.Type))

	return &l, nil
}

func leagueType(typ string) league.Type {
	typ = strings.ToLower(typ)
	switch typ {
	case "league":
		return league.TypeLeague
	case "cup":
		return league.TypeCup
	case "tournament":
		return league.TypeTournament
	case "friendly":
		return league.TypeFriendly
	default:
		return league.TypeLeague
	}
}
