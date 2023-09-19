package team_stats

import (
	"context"
	"mapeleven/db/ent"
)

func CreateTSPenalty(ctx context.Context, c *ent.Client, p Penalty, t *ent.Team) (*ent.TSPenalty, error) {
	return c.TSPenalty.
		Create().
		SetTeam(t).
		SetTotal(p.Total).
		SetScoredTotal(p.Scored.Total).
		SetScoredPercentage(p.Scored.Percentage).
		SetMissedTotal(p.Missed.Total).
		SetMissedPercentage(p.Missed.Percentage).
		Save(ctx)
}

func UpdateTSPenalty(ctx context.Context, c *ent.Client, p Penalty, tsPenalty *ent.TSPenalty) (*ent.TSPenalty, error) {
	return c.TSPenalty.
		UpdateOne(tsPenalty).
		SetTotal(p.Total).
		SetScoredTotal(p.Scored.Total).
		SetScoredPercentage(p.Scored.Percentage).
		SetMissedTotal(p.Missed.Total).
		SetMissedPercentage(p.Missed.Percentage).
		Save(ctx)
}
