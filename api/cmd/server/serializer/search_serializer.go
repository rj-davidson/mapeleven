package serializer

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
)

type APISearch struct {
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Image string `json:"image"`
}

func SearchSerializer(client *ent.Client) ([]APISearch, error) {
	var searchItems []APISearch

	leagueModel := models.NewLeagueModel(client)
	ls, err := leagueModel.ListLeagues(context.Background())
	for _, l := range ls {
		searchItems = append(searchItems, APISearch{
			Slug:  l.Slug,
			Name:  l.Name,
			Type:  "league",
			Image: l.Logo,
		})
	}

	clubModel := models.NewClubModel(client)
	cs, err := clubModel.ListAll(context.Background())
	for _, t := range cs {
		searchItems = append(searchItems, APISearch{
			Slug:  t.Slug,
			Name:  t.Name,
			Type:  "team",
			Image: t.Logo,
		})
	}

	playerModel := player_models.NewPlayerModel(client)
	ps, err := playerModel.ListPlayers(context.Background())
	for _, p := range ps {
		searchItems = append(searchItems, APISearch{
			Slug:  p.Slug,
			Name:  p.Name,
			Type:  "player",
			Image: p.Photo,
		})
	}

	return searchItems, err
}
