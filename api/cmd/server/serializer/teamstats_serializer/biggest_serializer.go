package teamstats_serializer

import "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"

type APITSBiggest struct {
	Streak APITSStreak  `json:"streak"`
	Win    Win          `json:"wins"`
	Lose   Lose         `json:"loses"`
	Goals  BiggestGoals `json:"goals"`
}

type APITSStreak struct {
	Wins   int `json:"wins"`
	Draws  int `json:"draws"`
	Losses int `json:"loses"`
}

type BiggestGoals struct {
	For     HomeAwayTotalInt `json:"for"`
	Against HomeAwayTotalInt `json:"against"`
}

func SerializeTSBiggest(biggest *ent.TSBiggest) APITSBiggest {
	return APITSBiggest{
		Streak: APITSStreak{
			Wins:   biggest.StreakWins,
			Draws:  biggest.StreakDraws,
			Losses: biggest.StreakLosses,
		},
		Win: Win{
			Home: biggest.WinsHome,
			Away: biggest.WinsAway,
		},
		Lose: Lose{
			Home: biggest.LossesHome,
			Away: biggest.LossesAway,
		},
		Goals: BiggestGoals{
			For: HomeAwayTotalInt{
				Home: biggest.GoalsForHome,
				Away: biggest.GoalsForAway,
			},
			Against: HomeAwayTotalInt{
				Home: biggest.GoalsAgainstHome,
				Away: biggest.GoalsAgainstAway,
			},
		},
	}
}
