package player_serializer

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"context"
	"sort"
)

type PlayerSerializer struct {
	client *ent.Client
}

func NewPlayerSerializer(client *ent.Client) *PlayerSerializer {
	return &PlayerSerializer{
		client: client,
	}

}

// GetPlayerBySlug returns a player by slug.
func (ps *PlayerSerializer) GetPlayerBySlug(ctx context.Context, slug string) (*APIPlayer, error) {
	_, err := ps.client.Player.Update().
		Where(player.SlugEQ(slug)).
		AddPopularity(1).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	p, err := ps.client.Player.
		Query().
		Where(player.SlugEQ(slug)).
		WithPlayerStats(func(q *ent.PlayerStatsQuery) {
			q.WithPsShooting().
				WithPsDefense().
				WithPsTechnical().
				WithPsFairplay().
				WithPsGames().
				WithPsPenalty().
				WithPsSubstitutes().
				WithTeam(func(tq *ent.TeamQuery) {
					tq.WithClub().
						WithSeason(func(sq *ent.SeasonQuery) {
							sq.WithLeague(func(lq *ent.LeagueQuery) {
								lq.WithCountry()
							})
						})
				})
		}).
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

	// Sort playerItems based on Popularity
	sort.Slice(playerItems, func(i, j int) bool {
		return playerItems[i].Popularity > playerItems[j].Popularity
	})

	return playerItems, nil
}

// SerializePlayer serializes a player entity into an APIPlayer.
func (ps *PlayerSerializer) SerializePlayer(p *ent.Player) *APIPlayer {
	var playerStats []APIPlayerStats
	if p.Edges.PlayerStats != nil {
		for _, stat := range p.Edges.PlayerStats {
			playerStats = append(playerStats, ps.serializePlayerStats(stat))
		}
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
		Popularity: p.Popularity,
	}
}

func (ps *PlayerSerializer) serializePlayerStats(s *ent.PlayerStats) APIPlayerStats {
	var team APIPSTeam
	var league APIPSLeague

	if s.Edges.Team != nil {
		if s.Edges.Team.Edges.Club != nil {
			team = APIPSTeam{
				Name: s.Edges.Team.Edges.Club.Name,
				Slug: s.Edges.Team.Edges.Club.Slug,
				Logo: s.Edges.Team.Edges.Club.Logo,
			}
		}
		if s.Edges.Team.Edges.Season != nil && s.Edges.Team.Edges.Season.Edges.League != nil {
			league = APIPSLeague{
				Slug:    s.Edges.Team.Edges.Season.Edges.League.Slug,
				Name:    s.Edges.Team.Edges.Season.Edges.League.Name,
				Country: s.Edges.Team.Edges.Season.Edges.League.Edges.Country.Name,
				Logo:    s.Edges.Team.Edges.Season.Edges.League.Logo,
				Flag:    s.Edges.Team.Edges.Season.Edges.League.Edges.Country.Flag,
				Season:  s.Edges.Team.Edges.Season.Year,
			}
		}
	}

	var defense APIPSDefense
	if s.Edges.PsDefense != nil {
		defense = APIPSDefense{
			TacklesTotal:  s.Edges.PsDefense.TacklesTotal,
			Blocks:        s.Edges.PsDefense.Blocks,
			Interceptions: s.Edges.PsDefense.Interceptions,
			DuelsTotal:    s.Edges.PsDefense.DuelsTotal,
			DuelsWon:      s.Edges.PsDefense.WonDuels,
			DribblePast:   s.Edges.PsDefense.DribblePast,
		}
	}

	var games APIPSGames
	if s.Edges.PsGames != nil {
		games = APIPSGames{
			Appearances: s.Edges.PsGames.Appearances,
			Lineups:     s.Edges.PsGames.Lineups,
			Minutes:     s.Edges.PsGames.Minutes,
			Number:      s.Edges.PsGames.Number,
			Position:    s.Edges.PsGames.Position,
			Rating:      s.Edges.PsGames.Rating,
			Captain:     s.Edges.PsGames.Captain,
		}
	}

	var shooting APIPSShooting
	if s.Edges.PsShooting != nil {
		shooting = APIPSShooting{
			Goals:    s.Edges.PsShooting.Goals,
			Conceded: s.Edges.PsShooting.Conceded,
			Assists:  s.Edges.PsShooting.Assists,
			Saves:    s.Edges.PsShooting.Saves,
			Shots:    s.Edges.PsShooting.Shots,
			OnTarget: s.Edges.PsShooting.OnTarget,
		}
	}

	var substitutes APIPSSubstitutes
	if s.Edges.PsSubstitutes != nil {
		substitutes = APIPSSubstitutes{
			In:    s.Edges.PsSubstitutes.In,
			Out:   s.Edges.PsSubstitutes.Out,
			Bench: s.Edges.PsSubstitutes.Bench,
		}
	}

	var technical APIPSTechnical
	if s.Edges.PsTechnical != nil {
		technical = APIPSTechnical{
			FoulsDrawn:      s.Edges.PsTechnical.FoulsDrawn,
			DribbleAttempts: s.Edges.PsTechnical.DribbleAttempts,
			DribblesSuccess: s.Edges.PsTechnical.DribbleSuccess,
			Passes:          s.Edges.PsTechnical.PassesTotal,
			KeyPasses:       s.Edges.PsTechnical.PassesKey,
			PassAccuracy:    s.Edges.PsTechnical.PassesAccuracy,
		}
	}

	var penalty APIPSPenalty
	if s.Edges.PsPenalty != nil {
		penalty = APIPSPenalty{
			Won:       s.Edges.PsPenalty.Won,
			Committed: s.Edges.PsPenalty.Committed,
			Scored:    s.Edges.PsPenalty.Scored,
			Saved:     s.Edges.PsPenalty.Saved,
			Missed:    s.Edges.PsPenalty.Missed,
		}
	}

	var fairplay APIPSFairplay
	if s.Edges.PsFairplay != nil {
		fairplay = APIPSFairplay{
			Yellow:            s.Edges.PsFairplay.Yellow,
			YellowRed:         s.Edges.PsFairplay.YellowRed,
			Red:               s.Edges.PsFairplay.Red,
			FoulsCommitted:    s.Edges.PsFairplay.FoulsCommitted,
			PenaltiesConceded: s.Edges.PsFairplay.PenaltyConceded,
		}
	}

	return APIPlayerStats{
		Team:        team,
		League:      league,
		Defense:     defense,
		Games:       games,
		Shooting:    shooting,
		Substitutes: substitutes,
		Technical:   technical,
		Penalty:     penalty,
		Fairplay:    fairplay,
	}
}
