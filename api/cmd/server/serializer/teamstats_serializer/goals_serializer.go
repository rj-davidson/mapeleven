package teamstats_serializer

import "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"

type APITSGoals struct {
	For     ScoreDetail `json:"for"`
	Against ScoreDetail `json:"against"`
}

func SerializeTSGoals(goals *ent.TSGoals) APITSGoals {
	return APITSGoals{
		For: ScoreDetail{
			Total: HomeAwayTotalInt{
				Home:  goals.GoalsForTotalHome,
				Away:  goals.GoalsForTotalAway,
				Total: goals.GoalsForTotal,
			},
			Average: HomeAwayTotalString{
				Home:  goals.GoalsForAverageHome,
				Away:  goals.GoalsForAverageAway,
				Total: goals.GoalsForAverageTotal,
			},
			Minute: MinuteDetail{
				Min0to15: MinuteStats{
					Total:      goals.GoalsForMinute0To15Total,
					Percentage: goals.GoalsForMinute0To15Percentage,
				},
				Min16to30: MinuteStats{
					Total:      goals.GoalsForMinute16To30Total,
					Percentage: goals.GoalsForMinute16To30Percentage,
				},
				Min31to45: MinuteStats{
					Total:      goals.GoalsForMinute31To45Total,
					Percentage: goals.GoalsForMinute31To45Percentage,
				},
				Min46to60: MinuteStats{
					Total:      goals.GoalsForMinute46To60Total,
					Percentage: goals.GoalsForMinute46To60Percentage,
				},
				Min61to75: MinuteStats{
					Total:      goals.GoalsForMinute61To75Total,
					Percentage: goals.GoalsForMinute61To75Percentage,
				},
				Min76to90: MinuteStats{
					Total:      goals.GoalsForMinute76To90Total,
					Percentage: goals.GoalsForMinute76To90Percentage,
				},
				Min91to105: MinuteStats{
					Total:      goals.GoalsForMinute91To105Total,
					Percentage: goals.GoalsForMinute91To105Percentage,
				},
				Min106to120: MinuteStats{
					Total:      goals.GoalsForMinute106To120Total,
					Percentage: goals.GoalsForMinute106To120Percentage,
				},
			},
		},
		Against: ScoreDetail{
			Total: HomeAwayTotalInt{
				Home:  goals.GoalsAgainstTotalHome,
				Away:  goals.GoalsAgainstTotalAway,
				Total: goals.GoalsAgainstTotal,
			},
			Average: HomeAwayTotalString{
				Home:  goals.GoalsAgainstAverageHome,
				Away:  goals.GoalsAgainstAverageAway,
				Total: goals.GoalsAgainstAverageTotal,
			},
			Minute: MinuteDetail{
				Min0to15: MinuteStats{
					Total:      goals.GoalsAgainstMinute0To15Total,
					Percentage: goals.GoalsAgainstMinute0To15Percentage,
				},
				Min16to30: MinuteStats{
					Total:      goals.GoalsAgainstMinute16To30Total,
					Percentage: goals.GoalsAgainstMinute16To30Percentage,
				},
				Min31to45: MinuteStats{
					Total:      goals.GoalsAgainstMinute31To45Total,
					Percentage: goals.GoalsAgainstMinute31To45Percentage,
				},
				Min46to60: MinuteStats{
					Total:      goals.GoalsAgainstMinute46To60Total,
					Percentage: goals.GoalsAgainstMinute46To60Percentage,
				},
				Min61to75: MinuteStats{
					Total:      goals.GoalsAgainstMinute61To75Total,
					Percentage: goals.GoalsAgainstMinute61To75Percentage,
				},
				Min76to90: MinuteStats{
					Total:      goals.GoalsAgainstMinute76To90Total,
					Percentage: goals.GoalsAgainstMinute76To90Percentage,
				},
				Min91to105: MinuteStats{
					Total:      goals.GoalsAgainstMinute91To105Total,
					Percentage: goals.GoalsAgainstMinute91To105Percentage,
				},
				Min106to120: MinuteStats{
					Total:      goals.GoalsAgainstMinute106To120Total,
					Percentage: goals.GoalsAgainstMinute106To120Percentage,
				},
			},
		},
	}
}
