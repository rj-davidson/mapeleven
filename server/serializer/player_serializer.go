package serializer

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/player"
)

type APIPlayer struct {
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

// GetPlayerBySlug returns a player by slug.
func (ps *PlayerSerializer) GetPlayerBySlug(slug string, ctx context.Context) (*APIPlayer, error) {
	p, err := ps.client.Player.
		Query().
		Where(player.SlugEQ(slug)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return ps.SerializePlayer(p), nil
}

// GetPlayers returns all players.
func (ps *PlayerSerializer) GetPlayers(ctx context.Context) ([]*APIPlayer, error) {
	players, err := ps.client.Player.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var playerItems []*APIPlayer
	for _, p := range players {
		playerItems = append(playerItems, ps.SerializePlayer(p))
	}

	return playerItems, nil
}

// SerializePlayer serializes a player entity into an APIPlayer.
func (ps *PlayerSerializer) SerializePlayer(p *ent.Player) *APIPlayer {

	return &APIPlayer{
		Name:      p.Name,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Age:       p.Age,
		Height:    p.Height,
		Weight:    p.Weight,
		Injured:   p.Injured,
		Photo:     p.Photo,
	}
}
