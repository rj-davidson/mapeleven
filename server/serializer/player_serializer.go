package serializer

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/player"
)

type PlayerInfo struct {
	ApiFootballID int    `json:"apiFootballID"`
	Name          string `json:"name"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Age           int    `json:"age"`
	Height        string `json:"height"`
	Weight        string `json:"weight"`
	Injured       bool   `json:"injured"`
	Photo         string `json:"photo"`
}

type PlayerSerializer struct {
	client *ent.Client
}

func NewPlayerSerializer(client *ent.Client) *PlayerSerializer {
	return &PlayerSerializer{client: client}
}

func (ps *PlayerSerializer) getPlayerData(apiFootballID int) (*ent.Player, error) {
	playerData, err := ps.client.Player.
		Query().
		Where(player.ApiFootballIdEQ(apiFootballID)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return playerData, nil
}

func (ps *PlayerSerializer) SerializePlayer(apiFootballID int) *PlayerInfo {
	playerData, err := ps.getPlayerData(apiFootballID)
	if err != nil {
		return nil
	}

	return &PlayerInfo{
		ApiFootballID: playerData.ApiFootballId,
		Name:          playerData.Name,
		Firstname:     playerData.Firstname,
		Lastname:      playerData.Lastname,
		Age:           playerData.Age,
		Height:        playerData.Height,
		Weight:        playerData.Weight,
		Injured:       playerData.Injured,
		Photo:         playerData.Photo,
	}
}
