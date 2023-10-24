package teamstats_serializer

import "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"

type APITSLineup struct {
	Formation string `json:"formation"`
	Played    int    `json:"played"`
}

func SerializeTSLineups(lus []*ent.TSLineups) []APITSLineup {
	var lineups []APITSLineup
	for _, lu := range lus {
		lineups = append(lineups, APITSLineup{
			Formation: lu.Formation,
			Played:    lu.Played,
		})
	}
	return lineups
}
