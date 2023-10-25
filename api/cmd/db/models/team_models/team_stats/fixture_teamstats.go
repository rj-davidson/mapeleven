package team_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
)

func CreateTSFixtures(ctx context.Context, c *ent.Client, f Fixtures, t *ent.Team) (*ent.TSFixtures, error) {
	return c.TSFixtures.
		Create().
		SetTeam(t).
		SetPlayedHome(f.Played.Home).
		SetPlayedAway(f.Played.Away).
		SetPlayedTotal(f.Played.Total).
		SetWinsHome(f.Wins.Home).
		SetWinsAway(f.Wins.Away).
		SetWinsTotal(f.Wins.Total).
		SetDrawsHome(f.Draws.Home).
		SetDrawsAway(f.Draws.Away).
		SetDrawsTotal(f.Draws.Total).
		SetLossesHome(f.Losses.Home).
		SetLossesAway(f.Losses.Away).
		SetLossesTotal(f.Losses.Total).
		Save(ctx)
}

func UpdateTSFixtures(ctx context.Context, c *ent.Client, f Fixtures, tsFixtures *ent.TSFixtures) (*ent.TSFixtures, error) {
	return c.TSFixtures.
		UpdateOne(tsFixtures).
		SetPlayedHome(f.Played.Home).
		SetPlayedAway(f.Played.Away).
		SetPlayedTotal(f.Played.Total).
		SetWinsHome(f.Wins.Home).
		SetWinsAway(f.Wins.Away).
		SetWinsTotal(f.Wins.Total).
		SetDrawsHome(f.Draws.Home).
		SetDrawsAway(f.Draws.Away).
		SetDrawsTotal(f.Draws.Total).
		SetLossesHome(f.Losses.Home).
		SetLossesAway(f.Losses.Away).
		SetLossesTotal(f.Losses.Total).
		Save(ctx)
}
