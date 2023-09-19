package team_stats

import (
	"context"
	"mapeleven/db/ent"
)

func CreateTSCards(ctx context.Context, c *ent.Client, cards Cards, t *ent.Team) (*ent.TSCards, error) {
	return c.TSCards.
		Create().
		SetTeam(t).SetYellow0To15Total(cards.Yellow.Min0to15.Total).
		SetYellow0To15Percentage(cards.Yellow.Min0to15.Percentage).
		SetYellow16To30Total(cards.Yellow.Min16to30.Total).
		SetYellow16To30Percentage(cards.Yellow.Min16to30.Percentage).
		SetYellow31To45Total(cards.Yellow.Min31to45.Total).
		SetYellow31To45Percentage(cards.Yellow.Min31to45.Percentage).
		SetYellow46To60Total(cards.Yellow.Min46to60.Total).
		SetYellow46To60Percentage(cards.Yellow.Min46to60.Percentage).
		SetYellow61To75Total(cards.Yellow.Min61to75.Total).
		SetYellow61To75Percentage(cards.Yellow.Min61to75.Percentage).
		SetYellow76To90Total(cards.Yellow.Min76to90.Total).
		SetYellow76To90Percentage(cards.Yellow.Min76to90.Percentage).
		SetYellow91to105Total(cards.Yellow.Min91to105.Total).
		SetYellow91to105Percentage(cards.Yellow.Min91to105.Percentage).
		SetRed0To15Total(cards.Red.Min0to15.Total).
		SetRed0To15Percentage(cards.Red.Min0to15.Percentage).
		SetRed16To30Total(cards.Red.Min16to30.Total).
		SetRed16To30Percentage(cards.Red.Min16to30.Percentage).
		SetRed31To45Total(cards.Red.Min31to45.Total).
		SetRed31To45Percentage(cards.Red.Min31to45.Percentage).
		SetRed46To60Total(cards.Red.Min46to60.Total).
		SetRed46To60Percentage(cards.Red.Min46to60.Percentage).
		SetRed61To75Total(cards.Red.Min61to75.Total).
		SetRed61To75Percentage(cards.Red.Min61to75.Percentage).
		SetRed76To90Total(cards.Red.Min76to90.Total).
		SetRed76To90Percentage(cards.Red.Min76to90.Percentage).
		SetRed91to105Total(cards.Red.Min91to105.Total).
		SetRed91to105Percentage(cards.Red.Min91to105.Percentage).
		SetRed106To120Total(cards.Red.Min106to120.Total).
		SetRed106To120Percentage(cards.Red.Min106to120.Percentage).
		Save(ctx)
}

func UpdateTSCards(ctx context.Context, c *ent.Client, cards Cards, cardsObj *ent.TSCards) (*ent.TSCards, error) {
	return c.TSCards.UpdateOne(cardsObj).
		SetYellow0To15Total(cards.Yellow.Min0to15.Total).
		SetYellow0To15Percentage(cards.Yellow.Min0to15.Percentage).
		SetYellow16To30Total(cards.Yellow.Min16to30.Total).
		SetYellow16To30Percentage(cards.Yellow.Min16to30.Percentage).
		SetYellow31To45Total(cards.Yellow.Min31to45.Total).
		SetYellow31To45Percentage(cards.Yellow.Min31to45.Percentage).
		SetYellow46To60Total(cards.Yellow.Min46to60.Total).
		SetYellow46To60Percentage(cards.Yellow.Min46to60.Percentage).
		SetYellow61To75Total(cards.Yellow.Min61to75.Total).
		SetYellow61To75Percentage(cards.Yellow.Min61to75.Percentage).
		SetYellow76To90Total(cards.Yellow.Min76to90.Total).
		SetYellow76To90Percentage(cards.Yellow.Min76to90.Percentage).
		SetYellow91to105Total(cards.Yellow.Min91to105.Total).
		SetYellow91to105Percentage(cards.Yellow.Min91to105.Percentage).
		SetYellow106To120Total(cards.Yellow.Min106to120.Total).
		SetYellow106To120Percentage(cards.Yellow.Min106to120.Percentage).
		SetRed0To15Total(cards.Red.Min0to15.Total).
		SetRed0To15Percentage(cards.Red.Min0to15.Percentage).
		SetRed16To30Total(cards.Red.Min16to30.Total).
		SetRed16To30Percentage(cards.Red.Min16to30.Percentage).
		SetRed31To45Total(cards.Red.Min31to45.Total).
		SetRed31To45Percentage(cards.Red.Min31to45.Percentage).
		SetRed46To60Total(cards.Red.Min46to60.Total).
		SetRed46To60Percentage(cards.Red.Min46to60.Percentage).
		SetRed61To75Total(cards.Red.Min61to75.Total).
		SetRed61To75Percentage(cards.Red.Min61to75.Percentage).
		SetRed76To90Total(cards.Red.Min76to90.Total).
		SetRed76To90Percentage(cards.Red.Min76to90.Percentage).
		SetRed91to105Total(cards.Red.Min91to105.Total).
		SetRed91to105Percentage(cards.Red.Min91to105.Percentage).
		SetRed106To120Total(cards.Red.Min106to120.Total).
		SetRed106To120Percentage(cards.Red.Min106to120.Percentage).
		Save(ctx)
}
