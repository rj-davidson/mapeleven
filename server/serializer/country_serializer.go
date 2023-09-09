package serializer

import (
	"mapeleven/db/ent"
)

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

func SerializeCountry(country *ent.Country) (*Country, error) {
	return &Country{
		Code: country.Code,
		Name: country.Name,
		Flag: country.Flag,
	}, nil
}
