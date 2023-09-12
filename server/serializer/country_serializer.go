package serializer

import (
	"mapeleven/db/ent"
)

type APICountry struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

func SerializeCountry(country *ent.Country) (*APICountry, error) {
	return &APICountry{
		Code: country.Code,
		Name: country.Name,
		Flag: country.Flag,
	}, nil
}
