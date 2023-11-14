package player_serializer

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"context"
)

type PlayerSerializer struct {
	client *ent.Client
}

func NewPlayerSerializer(client *ent.Client) *PlayerSerializer {
	return &PlayerSerializer{client: client}
}

// GetPlayerBySlug returns a player by slug.
func (ps *PlayerSerializer) GetPlayerBySlug(slug string, ctx context.Context) (*APIPlayer, error) {
	p, err := ps.client.Player.
		Query().
		Where(player.SlugEQ(slug)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return ps.SerializePlayer(p), nil
}

// GetPlayers returns all players.
func (ps *PlayerSerializer) GetPlayers(ctx context.Context) ([]*APIPlayer, error) {
	players, err := ps.client.Player.
		Query().
		WithPlayerStats(func(q *ent.PlayerStatsQuery) {
			q.WithPsShooting()
			q.WithPsDefense()
			q.WithPsTechnical()
			q.WithPsFairplay()
			q.WithPsGames()
			q.WithPsPenalty()
			q.WithPsSubstitutes()
			q.WithTeam(func(q *ent.TeamQuery) {
				q.WithClub()
				q.WithSeason(func(q *ent.SeasonQuery) {
					q.WithLeague(func(q *ent.LeagueQuery) {
						q.WithCountry()
					})
				})
			},
			)
		},
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var playerItems []*APIPlayer
	for _, p := range players {
		playerItems = append(playerItems, ps.SerializePlayer(p))
	}

	return playerItems, nil
}

// SerializePlayer serializes a player entity into an APIPlayer.
func (ps *PlayerSerializer) SerializePlayer(p *ent.Player) *APIPlayer {
	var playerStats []APIPlayerStats
	for _, stat := range p.Edges.PlayerStats {
		playerStats = append(playerStats, ps.serializePlayerStats(stat))
	}
	return &APIPlayer{
		Slug:       p.Slug,
		Name:       p.Name,
		Firstname:  p.Firstname,
		Lastname:   p.Lastname,
		Age:        p.Age,
		Height:     p.Height,
		Weight:     p.Weight,
		Injured:    p.Injured,
		Photo:      p.Photo,
		Statistics: playerStats,
	}
}

func (ps *PlayerSerializer) serializePlayerStats(s *ent.PlayerStats) APIPlayerStats {
	return APIPlayerStats{
		Team: APIPSTeam{
			Slug: s.Edges.Team.Edges.Club.Slug,
			Name: s.Edges.Team.Edges.Club.Name,
			Logo: s.Edges.Team.Edges.Club.Logo,
		},
		League: APIPSLeague{
			Slug:    s.Edges.Team.Edges.Season.Edges.League.Slug,
			Name:    s.Edges.Team.Edges.Season.Edges.League.Name,
			Country: s.Edges.Team.Edges.Season.Edges.League.Edges.Country.Name,
			Logo:    s.Edges.Team.Edges.Season.Edges.League.Logo,
			Flag:    s.Edges.Team.Edges.Season.Edges.League.Edges.Country.Flag,
			Season:  s.Edges.Team.Edges.Season.Year,
		},
		Defense: APIPSDefense{
			TacklesTotal:  s.Edges.PsDefense.TacklesTotal,
			Blocks:        s.Edges.PsDefense.Blocks,
			Interceptions: s.Edges.PsDefense.Interceptions,
			DuelsTotal:    s.Edges.PsDefense.DuelsTotal,
			DuelsWon:      s.Edges.PsDefense.WonDuels,
			DribblePast:   s.Edges.PsDefense.DribblePast,
		},
		Games: APIPSGames{
			Appearances: s.Edges.PsGames.Appearances,
			Lineups:     s.Edges.PsGames.Lineups,
			Minutes:     s.Edges.PsGames.Minutes,
			Number:      s.Edges.PsGames.Number,
			Position:    s.Edges.PsGames.Position,
			Rating:      s.Edges.PsGames.Rating,
			Captain:     s.Edges.PsGames.Captain,
		},
		Shooting: APIPSShooting{
			Goals:    s.Edges.PsShooting.Goals,
			Conceded: s.Edges.PsShooting.Conceded,
			Assists:  s.Edges.PsShooting.Assists,
			Saves:    s.Edges.PsShooting.Saves,
			Shots:    s.Edges.PsShooting.Shots,
			OnTarget: s.Edges.PsShooting.OnTarget,
		},
		Substitutes: APIPSSubstitutes{
			In:    s.Edges.PsSubstitutes.In,
			Out:   s.Edges.PsSubstitutes.Out,
			Bench: s.Edges.PsSubstitutes.Bench,
		},
		Technical: APIPSTechnical{
			FoulsDrawn:      s.Edges.PsTechnical.FoulsDrawn,
			DribbleAttempts: s.Edges.PsTechnical.DribbleAttempts,
			DribblesSuccess: s.Edges.PsTechnical.DribbleSuccess,
			Passes:          s.Edges.PsTechnical.PassesTotal,
			KeyPasses:       s.Edges.PsTechnical.PassesKey,
			PassAccuracy:    s.Edges.PsTechnical.PassesAccuracy,
		},
		Penalty: APIPSPenalty{
			Won:       s.Edges.PsPenalty.Won,
			Committed: s.Edges.PsPenalty.Committed,
			Scored:    s.Edges.PsPenalty.Scored,
			Saved:     s.Edges.PsPenalty.Saved,
			Missed:    s.Edges.PsPenalty.Missed,
		},
		Fairplay: APIPSFairplay{
			Yellow:            s.Edges.PsFairplay.Yellow,
			YellowRed:         s.Edges.PsFairplay.YellowRed,
			Red:               s.Edges.PsFairplay.Red,
			FoulsCommitted:    s.Edges.PsFairplay.FoulsCommitted,
			PenaltiesConceded: s.Edges.PsFairplay.PenaltyConceded,
		},
	}
}
