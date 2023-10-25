package fixture_serializer

import "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"

func SerializeLineup(lu *ent.FixtureLineups) (*APILineupInfo, error) {
	var lineup APILineupInfo
	var matchPlayers []APIPlayerDetail
	for _, mp := range lu.Edges.LineupPlayer {
		matchPlayer, err := serializeMatchPlayer(mp)
		if err != nil {
			return nil, err
		}
		matchPlayers = append(matchPlayers, *matchPlayer)
	}
	lineup = APILineupInfo{
		Team: APITeamInfo{
			Slug: lu.Edges.Team.Edges.Club.Slug,
			Name: lu.Edges.Team.Edges.Club.Name,
			Logo: lu.Edges.Team.Edges.Club.Logo,
		},
		Formation:    lu.Formation,
		MatchPlayers: matchPlayers,
	}
	return &lineup, nil
}

func serializeMatchPlayer(mp *ent.MatchPlayer) (*APIPlayerDetail, error) {
	var playerDetail APIPlayerDetail
	playerDetail = APIPlayerDetail{
		Player: APIPlayerDetailInfo{
			Slug:   mp.Edges.Player.Slug,
			Name:   mp.Edges.Player.Name,
			Number: mp.Number,
			Pos:    mp.Position,
			X:      mp.X,
			Y:      mp.Y,
		},
	}
	return &playerDetail, nil
}
