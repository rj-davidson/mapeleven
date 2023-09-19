package team_stats

import (
	"context"
	"mapeleven/db/ent"
)

func CreateTSBiggest(ctx context.Context, c *ent.Client, biggest Biggest, t *ent.Team) (*ent.TSBiggest, error) {
	return c.TSBiggest.
		Create().
		SetTeam(t).
		SetStreakWins(biggest.Streak.Wins).
		SetStreakDraws(biggest.Streak.Draws).
		SetStreakLosses(biggest.Streak.Losses).
		SetWinsHome(biggest.Win.Home).
		SetWinsAway(biggest.Win.Away).
		SetLossesHome(biggest.Lose.Home).
		SetLossesAway(biggest.Lose.Away).
		SetGoalsForHome(biggest.Goals.For.Home).
		SetGoalsForAway(biggest.Goals.For.Away).
		SetGoalsAgainstHome(biggest.Goals.Against.Home).
		SetGoalsAgainstAway(biggest.Goals.Against.Away).
		Save(ctx)
}

func UpdateTSBiggest(ctx context.Context, c *ent.Client, biggest Biggest, b *ent.TSBiggest) (*ent.TSBiggest, error) {
	return c.TSBiggest.
		UpdateOne(b).
		SetStreakWins(biggest.Streak.Wins).
		SetStreakDraws(biggest.Streak.Draws).
		SetStreakLosses(biggest.Streak.Losses).
		SetWinsHome(biggest.Win.Home).
		SetWinsAway(biggest.Win.Away).
		SetLossesHome(biggest.Lose.Home).
		SetLossesAway(biggest.Lose.Away).
		SetGoalsForHome(biggest.Goals.For.Home).
		SetGoalsForAway(biggest.Goals.For.Away).
		SetGoalsAgainstHome(biggest.Goals.Against.Home).
		SetGoalsAgainstAway(biggest.Goals.Against.Away).
		Save(ctx)
}
