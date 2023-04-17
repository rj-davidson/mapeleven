package fetchers

import (
	"errors"
	"mapeleven/models"
	"mapeleven/models/ent"
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

	// Search for country by code in Ent DB
	country, err := cf.countryModel.GetCountryByCode(c.Code)

	if ent.IsNotFound(errors.Unwrap(err)) {
		// Country does not exist, create it
		createInput := &models.CreateCountryInput{
			Name: c.Name,
			Code: c.Code,
			Flag: c.Flag,
		}
		country, err = cf.countryModel.CreateCountry(*createInput)
		if err != nil {
			return nil, err
		}
	} else if err == nil {
		// Country exists, update it
		updateInput := &models.UpdateCountryInput{
			Code: c.Code,
			Name: c.Name,
			Flag: c.Flag,
		}

		country, err = cf.countryModel.UpdateCountry(*updateInput)
		if err != nil {
			return nil, err
		}
	} else {
		// Return any other error that may occur
		return nil, err
	}

	return country, nil
}
