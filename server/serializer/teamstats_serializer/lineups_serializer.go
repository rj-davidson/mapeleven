package teamstats_serializer

import "mapeleven/db/ent"

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
