package controllers

import (
	"fmt"
	"mapeleven/db/ent"
	"mapeleven/models"
	"mapeleven/utils"
	"path/filepath"
	"strings"
)

type CountryController struct {
	countryModel *models.CountryModel
	entClient    *ent.Client
}

func NewCountryController(countryModel *models.CountryModel, entClient *ent.Client) *CountryController {
	return &CountryController{
		countryModel: countryModel,
		entClient:    entClient,
	}
}

func (cc *CountryController) UpsertCountry(c models.CreateCountryInput) (*ent.Country, error) {
	// Convert the code to uppercase if it is not empty
	c.Code = strings.ToUpper(c.Code)

	// If the code is empty, use the first 3 letters of the name
	if c.Code == "" {
		c.Code = strings.ToUpper(c.Name[:3])
	}

	// Download the flag and set the local path to the flag
	if c.Flag != "" {
		c.Flag, _ = utils.DownloadImageIfNeeded(c.Flag,
			fmt.Sprintf("images/flags/%s%s", c.Code, filepath.Ext(c.Flag)))
	}

	// Search for country by code in Ent DB
	country, _ := cc.countryModel.GetCountryByCode(c.Code)

	if country == nil {
		// Country does not exist, create it
		createInput := &models.CreateCountryInput{
			Name: c.Name,
			Code: c.Code,
			Flag: c.Flag,
		}
		country, _ = cc.countryModel.CreateCountry(*createInput)
	} else {
		// Country exists, update it
		updateInput := &models.UpdateCountryInput{
			Code: c.Code,
			Name: c.Name,
			Flag: c.Flag,
		}
		country, _ = cc.countryModel.UpdateCountry(*updateInput)
	}

	return country, nil
}
