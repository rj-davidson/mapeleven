package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
)

func CreatePSPenalty(ctx context.Context, c *ent.Client, pspenalty PSPenalty, p *ent.Player) (*ent.PSPenalty, error) {
	return c.PSPenalty.
		Create().
		SetPlayer(p).
		SetFoulsDrawn(pspenalty.FoulsDrawn).
		SetFoulsCommitted(pspenalty.FoulsCommitted).
		SetCardsYellow(pspenalty.CardsYellow).
		SetCardYellowred(pspenalty.CardYellowred).
		SetCardsRed(pspenalty.CardsRed).
		SetPenaltyWon(pspenalty.PenaltyWon).
		SetPenaltyCommited(pspenalty.PenaltyCommited).
		SetPenaltyScored(pspenalty.PenaltyScored).
		SetPenaltyMissed(pspenalty.PenaltyMissed).
		SetPenaltySaved(pspenalty.PenaltySaved).
		Save(ctx)
}

func UpdatePSPenalty(ctx context.Context, c *ent.Client, pspenalty PSPenalty, ps *ent.PSPenalty) (*ent.PSPenalty, error) {
	return c.PSPenalty.
		UpdateOne(ps).
		SetFoulsDrawn(pspenalty.FoulsDrawn).
		SetFoulsCommitted(pspenalty.FoulsCommitted).
		SetCardsYellow(pspenalty.CardsYellow).
		SetCardYellowred(pspenalty.CardYellowred).
		SetCardsRed(pspenalty.CardsRed).
		SetPenaltyWon(pspenalty.PenaltyWon).
		SetPenaltyCommited(pspenalty.PenaltyCommited).
		SetPenaltyScored(pspenalty.PenaltyScored).
		SetPenaltyMissed(pspenalty.PenaltyMissed).
		SetPenaltySaved(pspenalty.PenaltySaved).
		Save(ctx)
}
