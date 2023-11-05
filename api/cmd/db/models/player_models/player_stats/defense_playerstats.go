package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
)

func CreatePSDefense(ctx context.Context, c *ent.Client, psdefense PSDefense, p *ent.Player) (*ent.PSDefense, error) {
	return c.PSDefense.
		Create().
		SetPlayer(p).
		SetTacklesTotal(psdefense.TacklesTotal).
		SetBlocks(psdefense.Blocks).
		SetInterceptions(psdefense.Interceptions).
		SetTotalDuels(psdefense.TotalDuels).
		SetWonDuels(psdefense.WonDuels).
		Save(ctx)
}

func UpdatePSDefense(ctx context.Context, c *ent.Client, psdefense PSDefense, ps *ent.PSDefense) (*ent.PSDefense, error) {
	return c.PSDefense.
		UpdateOne(ps).
		SetTacklesTotal(psdefense.TacklesTotal).
		SetBlocks(psdefense.Blocks).
		SetInterceptions(psdefense.Interceptions).
		SetTotalDuels(psdefense.TotalDuels).
		SetWonDuels(psdefense.WonDuels).
		Save(ctx)
}
