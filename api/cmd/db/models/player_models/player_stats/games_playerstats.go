package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
)

type PSGames struct {
	Appearences int
	Lineups     int
	Minutes     int
	Number      int
	Position    string
	Rating      string
	Captain     bool
}

func CreatePSGames(ctx context.Context, c *ent.Client, psgames PSGames, p *ent.Player) (*ent.PSGames, error) {
	return c.PSGames.
		Create().
		SetPlayer(p).
		SetAppearences(psgames.Appearences).
		SetLineups(psgames.Lineups).
		SetMinutes(psgames.Minutes).
		SetNumber(psgames.Number).
		SetPosition(psgames.Position).
		SetRating(psgames.Rating).
		SetCaptain(psgames.Captain).
		Save(ctx)
}

func UpdatePSGames(ctx context.Context, c *ent.Client, psgames PSGames, ps *ent.PSGames) (*ent.PSGames, error) {
	return c.PSGames.
		UpdateOne(ps).
		SetAppearences(psgames.Appearences).
		SetLineups(psgames.Lineups).
		SetMinutes(psgames.Minutes).
		SetNumber(psgames.Number).
		SetPosition(psgames.Position).
		SetRating(psgames.Rating).
		SetCaptain(psgames.Captain).
		Save(ctx)
}
