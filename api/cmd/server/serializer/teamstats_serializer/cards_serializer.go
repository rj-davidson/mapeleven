package teamstats_serializer

import "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"

type APITSCards struct {
	Yellow MinuteDetail `json:"yellow"`
	Red    MinuteDetail `json:"red"`
}

func SerializeTSCards(cards *ent.TSCards) APITSCards {
	return APITSCards{
		Yellow: MinuteDetail{
			Min0to15: MinuteStats{
				Total:      cards.Yellow0To15Total,
				Percentage: cards.Yellow0To15Percentage,
			},
			Min16to30: MinuteStats{
				Total:      cards.Yellow16To30Total,
				Percentage: cards.Yellow16To30Percentage,
			},
			Min31to45: MinuteStats{
				Total:      cards.Yellow31To45Total,
				Percentage: cards.Yellow31To45Percentage,
			},
			Min46to60: MinuteStats{
				Total:      cards.Yellow46To60Total,
				Percentage: cards.Yellow46To60Percentage,
			},
			Min61to75: MinuteStats{
				Total:      cards.Yellow61To75Total,
				Percentage: cards.Yellow61To75Percentage,
			},
			Min76to90: MinuteStats{
				Total:      cards.Yellow76To90Total,
				Percentage: cards.Yellow76To90Percentage,
			},
			Min91to105: MinuteStats{
				Total:      cards.Yellow91to105Total,
				Percentage: cards.Yellow91to105Percentage,
			},
			Min106to120: MinuteStats{
				Total:      cards.Yellow106To120Total,
				Percentage: cards.Yellow106To120Percentage,
			},
		},
		Red: MinuteDetail{
			Min0to15: MinuteStats{
				Total:      cards.Red0To15Total,
				Percentage: cards.Red0To15Percentage,
			},
			Min16to30: MinuteStats{
				Total:      cards.Red16To30Total,
				Percentage: cards.Red16To30Percentage,
			},
			Min31to45: MinuteStats{
				Total:      cards.Red31To45Total,
				Percentage: cards.Red31To45Percentage,
			},
			Min46to60: MinuteStats{
				Total:      cards.Red46To60Total,
				Percentage: cards.Red46To60Percentage,
			},
			Min61to75: MinuteStats{
				Total:      cards.Red61To75Total,
				Percentage: cards.Red61To75Percentage,
			},
			Min76to90: MinuteStats{
				Total:      cards.Red76To90Total,
				Percentage: cards.Red76To90Percentage,
			},
			Min91to105: MinuteStats{
				Total:      cards.Red91to105Total,
				Percentage: cards.Red91to105Percentage,
			},
			Min106to120: MinuteStats{
				Total:      cards.Red106To120Total,
				Percentage: cards.Red106To120Percentage,
			},
		},
	}
}
