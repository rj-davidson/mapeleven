package team_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
)

func CreateTSCleanSheet(ctx context.Context, c *ent.Client, cleanSheet CleanSheet, t *ent.Team) (*ent.TSCleanSheet, error) {
	return c.TSCleanSheet.
		Create().
		SetTeam(t).
		SetHome(cleanSheet.Home).
		SetAway(cleanSheet.Away).
		SetTotal(cleanSheet.Total).
		Save(ctx)
}

func UpdateTSCleanSheet(ctx context.Context, c *ent.Client, cleanSheet CleanSheet, cs *ent.TSCleanSheet) (*ent.TSCleanSheet, error) {
	return c.TSCleanSheet.
		UpdateOne(cs).
		SetHome(cleanSheet.Home).
		SetAway(cleanSheet.Away).
		SetTotal(cleanSheet.Total).
		Save(ctx)
}
