package fetchers

import (
	"fmt"
	"mapeleven/models"
	"mapeleven/models/ent"
	"mapeleven/utils"
	"path/filepath"
	"strings"
)

type CountryFetcher struct {
	countryModel *models.CountryModel
	entClient    *ent.Client
}

func NewCountryFetcher(countryModel *models.CountryModel, entClient *ent.Client) *CountryFetcher {
	return &CountryFetcher{
		countryModel: countryModel,
		entClient:    entClient,
	}
}

func (cf *CountryFetcher) UpsertCountry(c models.CreateCountryInput) (*ent.Country, error) {
	// Convert the code to uppercase if it is not empty
	c.Code = strings.ToUpper(c.Code)
	if c.Code == "" {
		c.Code = strings.ToUpper(c.Name[:3])
	}

	// Download the flag and set the local path to the flag
	if c.Flag != "" {
		c.Flag, _ = utils.DownloadImageIfNeeded(c.Flag,
			fmt.Sprintf("images/flags/%s%s", c.Code, filepath.Ext(c.Flag)))
	}

	// Search for country by code in Ent DB
	country, _ := cf.countryModel.GetCountryByCode(c.Code)

	if country == nil {
		// Country does not exist, create it
		createInput := &models.CreateCountryInput{
			Name: c.Name,
			Code: c.Code,
			Flag: c.Flag,
		}
		country, _ = cf.countryModel.CreateCountry(*createInput)
	} else {
		// Country exists, update it
		updateInput := &models.UpdateCountryInput{
			Code: c.Code,
			Name: c.Name,
			Flag: c.Flag,
		}
		country, _ = cf.countryModel.UpdateCountry(*updateInput)
	}

	return country, nil
}
