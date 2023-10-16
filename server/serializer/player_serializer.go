package serializer

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/player"
)

type PlayerInfo struct {
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	Injured   bool   `json:"injured"`
	Photo     string `json:"photo"`
}

type PlayerSerializer struct {
	client *ent.Client
}

func NewPlayerSerializer(client *ent.Client) *PlayerSerializer {
	return &PlayerSerializer{client: client}
}

func (ps *PlayerSerializer) getPlayerData(slug string) (*ent.Player, error) {
	playerData, err := ps.client.Player.
		Query().
		Where(player.SlugEQ(slug)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return playerData, nil
}

func (ps *PlayerSerializer) SerializePlayer(slug string) *PlayerInfo {
	playerData, err := ps.getPlayerData(slug)
	if err != nil {
		return nil
	}

	return &PlayerInfo{
		Name:      playerData.Name,
		Firstname: playerData.Firstname,
		Lastname:  playerData.Lastname,
		Age:       playerData.Age,
		Height:    playerData.Height,
		Weight:    playerData.Weight,
		Injured:   playerData.Injured,
		Photo:     playerData.Photo,
	}
}
