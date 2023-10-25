package teamstats_serializer

import "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"

type APITSPenalty struct {
	Scored GoalFixedPercentage `json:"scored"`
	Missed GoalFixedPercentage `json:"missed"`
	Total  int                 `json:"total"`
}

type GoalFixedPercentage struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}

func SerializeTSPenalty(penalty *ent.TSPenalty) APITSPenalty {
	return APITSPenalty{
		Scored: GoalFixedPercentage{
			Total:      penalty.ScoredTotal,
			Percentage: penalty.ScoredPercentage,
		},
		Missed: GoalFixedPercentage{
			Total:      penalty.MissedTotal,
			Percentage: penalty.MissedPercentage,
		},
		Total: penalty.Total,
	}
}
