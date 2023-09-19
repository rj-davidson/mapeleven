package teamstats_serializer

import "mapeleven/db/ent"

type APITSFixtures struct {
	Played HomeAwayTotalInt `json:"played"`
	Wins   HomeAwayTotalInt `json:"wins"`
	Draws  HomeAwayTotalInt `json:"draws"`
	Losses HomeAwayTotalInt `json:"loses"`
}

func SerializeTSFixtures(fixtures *ent.TSFixtures) APITSFixtures {
	return APITSFixtures{
		Played: HomeAwayTotalInt{
			Home: fixtures.PlayedHome,
			Away: fixtures.PlayedAway,
		},
		Wins: HomeAwayTotalInt{
			Home: fixtures.WinsHome,
			Away: fixtures.WinsAway,
		},
		Draws: HomeAwayTotalInt{
			Home: fixtures.DrawsHome,
			Away: fixtures.DrawsAway,
		},
		Losses: HomeAwayTotalInt{
			Home: fixtures.LossesHome,
			Away: fixtures.LossesAway,
		},
	}
}
