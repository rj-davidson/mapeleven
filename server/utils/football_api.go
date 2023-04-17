package utils

import (
	"encoding/json"
	"errors"
	"mapeleven/models/ent/league"
	"strings"

	"mapeleven/models"
)

type APILeague struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Type league.Type `json:"type"`
	Logo string      `json:"logo"`
}

// Returns a CreateCountryInput struct, accepts a JSON country object
func GetCountryData(countryJSON []byte) (models.CreateCountryInput, error) {
	var countryData map[string]interface{}
	if err := json.Unmarshal(countryJSON, &countryData); err != nil {
		return models.CreateCountryInput{}, err
	}

	name, ok := countryData["name"].(string)
	if !ok {
		return models.CreateCountryInput{}, errors.New("missing or invalid name field")
	}

	code, ok := countryData["code"].(string)
	if !ok {
		// Set code to the first three characters of the name in uppercase
		code = strings.ToUpper(name[:3])
	}

	flag, ok := countryData["flag"].(string)
	if !ok {
		return models.CreateCountryInput{}, errors.New("missing or invalid flag field")
	}

	return models.CreateCountryInput{
		Name: name,
		Code: code,
		Flag: flag,
	}, nil
}

// Returns a CreateLeagueInput struct, accepts a JSON league object
func GetLeagueData(leagueJSON []byte) (models.CreateLeagueInput, error) {
	var leagueData map[string]interface{}
	if err := json.Unmarshal(leagueJSON, &leagueData); err != nil {
		return models.CreateLeagueInput{}, err
	}

	name, ok := leagueData["name"].(string)
	if !ok {
		return models.CreateLeagueInput{}, errors.New("missing or invalid name field")
	}

	leagueType, ok := leagueData["type"].(string)
	if !ok {
		return models.CreateLeagueInput{}, errors.New("missing or invalid type field")
	}

	logo, ok := leagueData["logo"].(string)
	if !ok {
		return models.CreateLeagueInput{}, errors.New("missing or invalid logo field")
	}

	return models.CreateLeagueInput{
		Name: name,
		Type: league.Type(leagueType),
		Logo: logo,
	}, nil
}
