package team_stats

import (
	"context"
	"mapeleven/db/ent"
)

func CreateTSGoals(ctx context.Context, c *ent.Client, g Goals, t *ent.Team) (*ent.TSGoals, error) {
	return c.TSGoals.
		Create().
		SetTeam(t).
		SetGoalsForTotalHome(g.For.Total.Home).
		SetGoalsForTotalAway(g.For.Total.Away).
		SetGoalsForTotal(g.For.Total.Total).
		SetGoalsAgainstTotalHome(g.Against.Total.Home).
		SetGoalsAgainstTotalAway(g.Against.Total.Away).
		SetGoalsAgainstTotal(g.Against.Total.Total).
		SetGoalsForAverageHome(g.For.Average.Home).
		SetGoalsForAverageAway(g.For.Average.Away).
		SetGoalsForAverageTotal(g.For.Average.Total).
		SetGoalsAgainstAverageHome(g.Against.Average.Home).
		SetGoalsAgainstAverageAway(g.Against.Average.Away).
		SetGoalsAgainstAverageTotal(g.Against.Average.Total).
		SetGoalsForMinute0To15Total(g.For.Minute.Min0to15.Total).
		SetGoalsForMinute0To15Percentage(g.For.Minute.Min0to15.Percentage).
		SetGoalsForMinute16To30Total(g.For.Minute.Min16to30.Total).
		SetGoalsForMinute16To30Percentage(g.For.Minute.Min16to30.Percentage).
		SetGoalsForMinute31To45Total(g.For.Minute.Min31to45.Total).
		SetGoalsForMinute31To45Percentage(g.For.Minute.Min31to45.Percentage).
		SetGoalsForMinute46To60Total(g.For.Minute.Min46to60.Total).
		SetGoalsForMinute46To60Percentage(g.For.Minute.Min46to60.Percentage).
		SetGoalsForMinute61To75Total(g.For.Minute.Min61to75.Total).
		SetGoalsForMinute61To75Percentage(g.For.Minute.Min61to75.Percentage).
		SetGoalsForMinute76To90Total(g.For.Minute.Min76to90.Total).
		SetGoalsForMinute76To90Percentage(g.For.Minute.Min76to90.Percentage).
		SetGoalsForMinute91To105Total(g.For.Minute.Min91to105.Total).
		SetGoalsForMinute91To105Percentage(g.For.Minute.Min91to105.Percentage).
		SetGoalsForMinute106To120Total(g.For.Minute.Min106to120.Total).
		SetGoalsForMinute106To120Percentage(g.For.Minute.Min106to120.Percentage).
		SetGoalsAgainstMinute0To15Total(g.Against.Minute.Min0to15.Total).
		SetGoalsAgainstMinute0To15Percentage(g.Against.Minute.Min0to15.Percentage).
		SetGoalsAgainstMinute16To30Total(g.Against.Minute.Min16to30.Total).
		SetGoalsAgainstMinute16To30Percentage(g.Against.Minute.Min16to30.Percentage).
		SetGoalsAgainstMinute31To45Total(g.Against.Minute.Min31to45.Total).
		SetGoalsAgainstMinute31To45Percentage(g.Against.Minute.Min31to45.Percentage).
		SetGoalsAgainstMinute46To60Total(g.Against.Minute.Min46to60.Total).
		SetGoalsAgainstMinute46To60Percentage(g.Against.Minute.Min46to60.Percentage).
		SetGoalsAgainstMinute61To75Total(g.Against.Minute.Min61to75.Total).
		SetGoalsAgainstMinute61To75Percentage(g.Against.Minute.Min61to75.Percentage).
		SetGoalsAgainstMinute76To90Total(g.Against.Minute.Min76to90.Total).
		SetGoalsAgainstMinute76To90Percentage(g.Against.Minute.Min76to90.Percentage).
		SetGoalsAgainstMinute91To105Total(g.Against.Minute.Min91to105.Total).
		SetGoalsAgainstMinute91To105Percentage(g.Against.Minute.Min91to105.Percentage).
		SetGoalsAgainstMinute106To120Total(g.Against.Minute.Min106to120.Total).
		SetGoalsAgainstMinute106To120Percentage(g.Against.Minute.Min106to120.Percentage).
		Save(ctx)
}

func UpdateTSGoals(ctx context.Context, c *ent.Client, g Goals, gs *ent.TSGoals) (*ent.TSGoals, error) {
	return c.TSGoals.
		UpdateOne(gs).
		SetGoalsForTotalHome(g.For.Total.Home).
		SetGoalsForTotalAway(g.For.Total.Away).
		SetGoalsForTotal(g.For.Total.Total).
		SetGoalsAgainstTotalHome(g.Against.Total.Home).
		SetGoalsAgainstTotalAway(g.Against.Total.Away).
		SetGoalsAgainstTotal(g.Against.Total.Total).
		SetGoalsForAverageHome(g.For.Average.Home).
		SetGoalsForAverageAway(g.For.Average.Away).
		SetGoalsForAverageTotal(g.For.Average.Total).
		SetGoalsAgainstAverageHome(g.Against.Average.Home).
		SetGoalsAgainstAverageAway(g.Against.Average.Away).
		SetGoalsAgainstAverageTotal(g.Against.Average.Total).
		SetGoalsForMinute0To15Total(g.For.Minute.Min0to15.Total).
		SetGoalsForMinute0To15Percentage(g.For.Minute.Min0to15.Percentage).
		SetGoalsForMinute16To30Total(g.For.Minute.Min16to30.Total).
		SetGoalsForMinute16To30Percentage(g.For.Minute.Min16to30.Percentage).
		SetGoalsForMinute31To45Total(g.For.Minute.Min31to45.Total).
		SetGoalsForMinute31To45Percentage(g.For.Minute.Min31to45.Percentage).
		SetGoalsForMinute46To60Total(g.For.Minute.Min46to60.Total).
		SetGoalsForMinute46To60Percentage(g.For.Minute.Min46to60.Percentage).
		SetGoalsForMinute61To75Total(g.For.Minute.Min61to75.Total).
		SetGoalsForMinute61To75Percentage(g.For.Minute.Min61to75.Percentage).
		SetGoalsForMinute76To90Total(g.For.Minute.Min76to90.Total).
		SetGoalsForMinute76To90Percentage(g.For.Minute.Min76to90.Percentage).
		SetGoalsForMinute91To105Total(g.For.Minute.Min91to105.Total).
		SetGoalsForMinute91To105Percentage(g.For.Minute.Min91to105.Percentage).
		SetGoalsForMinute106To120Total(g.For.Minute.Min106to120.Total).
		SetGoalsForMinute106To120Percentage(g.For.Minute.Min106to120.Percentage).
		SetGoalsAgainstMinute0To15Total(g.Against.Minute.Min0to15.Total).
		SetGoalsAgainstMinute0To15Percentage(g.Against.Minute.Min0to15.Percentage).
		SetGoalsAgainstMinute16To30Total(g.Against.Minute.Min16to30.Total).
		SetGoalsAgainstMinute16To30Percentage(g.Against.Minute.Min16to30.Percentage).
		SetGoalsAgainstMinute31To45Total(g.Against.Minute.Min31to45.Total).
		SetGoalsAgainstMinute31To45Percentage(g.Against.Minute.Min31to45.Percentage).
		SetGoalsAgainstMinute46To60Total(g.Against.Minute.Min46to60.Total).
		SetGoalsAgainstMinute46To60Percentage(g.Against.Minute.Min46to60.Percentage).
		SetGoalsAgainstMinute61To75Total(g.Against.Minute.Min61to75.Total).
		SetGoalsAgainstMinute61To75Percentage(g.Against.Minute.Min61to75.Percentage).
		SetGoalsAgainstMinute76To90Total(g.Against.Minute.Min76to90.Total).
		SetGoalsAgainstMinute76To90Percentage(g.Against.Minute.Min76to90.Percentage).
		SetGoalsAgainstMinute91To105Total(g.Against.Minute.Min91to105.Total).
		SetGoalsAgainstMinute91To105Percentage(g.Against.Minute.Min91to105.Percentage).
		SetGoalsAgainstMinute106To120Total(g.Against.Minute.Min106to120.Total).
		SetGoalsAgainstMinute106To120Percentage(g.Against.Minute.Min106to120.Percentage).
		Save(ctx)
}
