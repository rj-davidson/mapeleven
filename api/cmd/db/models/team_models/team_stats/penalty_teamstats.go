package team_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
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
