package team_stats

import (
	"context"
	"mapeleven/db/ent"
)

func CreateTSFailedToScore(ctx context.Context, c *ent.Client, failedToScore FailedToScore, t *ent.Team) (*ent.TSFailedToScore, error) {
	return c.TSFailedToScore.
		Create().
		SetTeam(t).
		SetHome(failedToScore.Home).
		SetAway(failedToScore.Away).
		SetTotal(failedToScore.Total).
		Save(ctx)
}

func UpdateTSFailedToScore(ctx context.Context, c *ent.Client, failedToScore FailedToScore, fts *ent.TSFailedToScore) (*ent.TSFailedToScore, error) {
	return c.TSFailedToScore.
		UpdateOne(fts).
		SetHome(failedToScore.Home).
		SetAway(failedToScore.Away).
		SetTotal(failedToScore.Total).
		Save(ctx)
}
