package player_stats

//import (
//	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
//	"context"
//)
//
//func CreatePSOffense(ctx context.Context, c *ent.Client, psoffense PSOffense, p *ent.PlayerStats) (*ent.PSOffense, error) {
//	return c.PSOffense.
//		Create().
//		SetPlayerStats(p).
//		SetDribbleAttempts(psoffense.DribbleAttempts).
//		SetDribbleSuccess(psoffense.DribbleSuccess).
//		SetDribblePast(psoffense.DribblePast).
//		SetPassesKey(psoffense.PassesKey).
//		SetPassesTotal(psoffense.PassesTotal).
//		SetPassesAccuracy(psoffense.PassesAccuracy).
//		Save(ctx)
//}
//
//func UpdatePSOffense(ctx context.Context, c *ent.Client, psoffense PSOffense, ps *ent.PSOffense) (*ent.PSOffense, error) {
//	return c.PSOffense.
//		UpdateOne(ps).
//		SetDribbleAttempts(psoffense.DribbleAttempts).
//		SetDribbleSuccess(psoffense.DribbleSuccess).
//		SetDribblePast(psoffense.DribblePast).
//		SetPassesKey(psoffense.PassesKey).
//		SetPassesTotal(psoffense.PassesTotal).
//		SetPassesAccuracy(psoffense.PassesAccuracy).
//		Save(ctx)
//}
