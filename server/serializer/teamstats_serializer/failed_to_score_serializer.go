package teamstats_serializer

import "mapeleven/db/ent"

type APITSFailedToScore struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

func SerializeTSFailedToScore(failedToScore *ent.TSFailedToScore) APITSFailedToScore {
	return APITSFailedToScore{
		Home:  failedToScore.Home,
		Away:  failedToScore.Away,
		Total: failedToScore.Total,
	}
}
