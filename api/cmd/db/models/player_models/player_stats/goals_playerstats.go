package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
)

func CreatePSGoals(ctx context.Context, c *ent.Client, psgoals PSGoals, p *ent.PlayerStats) (*ent.PSGoals, error) {
	return c.PSGoals.
		Create().
		SetPlayerStats(p).
		SetTotalGoals(psgoals.TotalGoals).
		SetConcededGoals(psgoals.ConcededGoals).
		SetAssistGoals(psgoals.AssistGoals).
		SetSaveGoals(psgoals.SaveGoals).
		SetShotsTotal(psgoals.ShotsTotal).
		SetShotsOn(psgoals.ShotsOn).
		Save(ctx)
}

func UpdatePSGoals(ctx context.Context, c *ent.Client, psgoals PSGoals, ps *ent.PSGoals) (*ent.PSGoals, error) {
	return c.PSGoals.
		UpdateOne(ps).
		SetTotalGoals(psgoals.TotalGoals).
		SetConcededGoals(psgoals.ConcededGoals).
		SetAssistGoals(psgoals.AssistGoals).
		SetSaveGoals(psgoals.SaveGoals).
		SetShotsTotal(psgoals.ShotsTotal).
		SetShotsOn(psgoals.ShotsOn).
		Save(ctx)
}
