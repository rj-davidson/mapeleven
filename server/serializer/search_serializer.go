package serializer

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/models"
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

	teamModel := models.NewTeamModel(client)
	teams, err := teamModel.ListTeams(context.Background())
	for _, team := range teams {
		searchItems = append(searchItems, APISearch{
			Slug:  team.Slug,
			Name:  team.Name,
			Type:  "team",
			Image: team.Logo,
		})
	}

	playerModel := models.NewPlayerModel(client)
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
