package fetchers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mapeleven/models"
	"mapeleven/utils"
	"net/http"
)

type CountryFetcher struct {
	apiKey         string
	client         *http.Client
	countryModel   *models.CountryModel
	apiBaseURL     string
	apiCountryPath string
}

type APICountryResponse struct {
	Response []APICountry `json:"response"`
}

type APICountry struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

func NewCountryFetcher(apiKey string, countryModel *models.CountryModel) *CountryFetcher {
	return &CountryFetcher{
		apiKey:         apiKey,
		client:         &http.Client{},
		countryModel:   countryModel,
		apiBaseURL:     "https://api-football-v1.p.rapidapi.com",
		apiCountryPath: "/v3/countries",
	}
}

func (cf *CountryFetcher) FetchCountry(countryCode string) (*models.CreateCountryInput, error) {
	apiCountry, err := cf.fetchCountryByCode(countryCode)
	if err != nil {
		return nil, err
	}

	filePath := fmt.Sprintf("public/images/flags/%s.svg", apiCountry.Code)
	err = utils.DownloadFile(apiCountry.Flag, filePath)
	if err != nil {
		return nil, err
	}

	input := &models.CreateCountryInput{
		Name: apiCountry.Name,
		Code: apiCountry.Code,
		Flag: filePath,
	}

	return input, nil
}

func (cf *CountryFetcher) fetchCountryByCode(code string) (*APICountry, error) {
	url := fmt.Sprintf("%s%s?code=%s", cf.apiBaseURL, cf.apiCountryPath, code)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", cf.apiKey)

	resp, err := cf.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse APICountryResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	if len(apiResponse.Response) == 0 {
		return nil, fmt.Errorf("country not found")
	}

	return &apiResponse.Response[0], nil
}
