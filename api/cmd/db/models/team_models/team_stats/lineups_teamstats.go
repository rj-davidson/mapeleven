package team_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
)

func CreateTSLineup(ctx context.Context, c *ent.Client, lineup Lineup, t *ent.Team) (*ent.TSLineups, error) {
	return c.TSLineups.
		Create().
		SetTeam(t).
		SetFormation(lineup.Formation).
		SetPlayed(lineup.Played).
		Save(ctx)
}

func UpdateTSLineup(ctx context.Context, c *ent.Client, lineup Lineup, tsLineup *ent.TSLineups) (*ent.TSLineups, error) {
	return c.TSLineups.
		UpdateOne(tsLineup).
		SetPlayed(lineup.Played).
		Save(ctx)
}
