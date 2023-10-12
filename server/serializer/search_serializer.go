package serializer

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/models"
	"mapeleven/models/player_models"
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
	leagues, err := leagueModel.ListLeagues(context.Background())
	for _, league := range leagues {
		searchItems = append(searchItems, APISearch{
			Slug:  league.Slug,
			Name:  league.Name,
			Type:  "league",
			Image: league.Logo,
		})
	}

	clubModel := models.NewClubModel(client)
	clubs, err := clubModel.ListAll(context.Background())
	for _, team := range clubs {
		searchItems = append(searchItems, APISearch{
			Slug:  team.Slug,
			Name:  team.Name,
			Type:  "team",
			Image: team.Logo,
		})
	}

	playerModel := player_models.NewPlayerModel(client)
	players, err := playerModel.ListPlayers(context.Background())
	for _, player := range players {
		searchItems = append(searchItems, APISearch{
			Slug:  player.Slug,
			Name:  player.Name,
			Type:  "player",
			Image: player.Photo,
		})
	}

	return searchItems, err
}
