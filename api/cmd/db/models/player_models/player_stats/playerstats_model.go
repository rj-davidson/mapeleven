package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psshooting"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pstechnical"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"context"
	_ "context"
)

type PSTeam struct {
	ApiFootballId int    `json:"id"`
	Name          string `json:"name"`
	Logo          string `json:"logo"`
}

type PSLeague struct {
	ApiFootballId int    `json:"id"`
	Name          string `json:"name"`
	Country       string `json:"country"`
	Logo          string `json:"logo"`
	Flag          string `json:"flag"`
	Season        int    `json:"season"`
}

type PSGames struct {
	Appearances int    `json:"appearances"`
	Lineups     int    `json:"lineups"`
	Minutes     int    `json:"minutes"`
	Number      *int   `json:"number"`
	Position    string `json:"position"`
	Rating      string `json:"rating"`
	Captain     bool   `json:"captain"`
}

type PSSubstitutes struct {
	In    int `json:"in"`
	Out   int `json:"out"`
	Bench int `json:"bench"`
}

type PSShots struct {
	Total int `json:"total"`
	On    int `json:"on"`
}

type PSGoals struct {
	Total    int  `json:"total"`
	Conceded int  `json:"conceded"`
	Assists  int  `json:"assists"`
	Saves    *int `json:"saves"`
}

type PSPasses struct {
	Total    int `json:"total"`
	Key      int `json:"key"`
	Accuracy int `json:"accuracy"`
}

type PSTackles struct {
	Total         int  `json:"total"`
	Blocks        *int `json:"blocks"`
	Interceptions int  `json:"interceptions"`
}

type PSDuels struct {
	Total int `json:"total"`
	Won   int `json:"won"`
}

type PSDribbles struct {
	Attempts int  `json:"attempts"`
	Success  int  `json:"success"`
	Past     *int `json:"past"`
}

type PSFouls struct {
	Drawn     int `json:"drawn"`
	Committed int `json:"committed"`
}

type PSCards struct {
	Yellow    int `json:"yellow"`
	YellowRed int `json:"yellowred"`
	Red       int `json:"red"`
}

type PSPenalty struct {
	Won       *int `json:"won"`
	Committed *int `json:"commited"`
	Scored    int  `json:"scored"`
	Missed    int  `json:"missed"`
	Saved     *int `json:"saved"`
}

type PlayerStats struct {
	Team        PSTeam        `json:"team"`
	League      PSLeague      `json:"league"`
	Games       PSGames       `json:"games"`
	Substitutes PSSubstitutes `json:"substitutes"`
	Shots       PSShots       `json:"shots"`
	Goals       PSGoals       `json:"goals"`
	Passes      PSPasses      `json:"passes"`
	Tackles     PSTackles     `json:"tackles"`
	Duels       PSDuels       `json:"duels"`
	Dribbles    PSDribbles    `json:"dribbles"`
	Fouls       PSFouls       `json:"fouls"`
	Cards       PSCards       `json:"cards"`
	Penalty     PSPenalty     `json:"penalty"`
}

type PlayerStatsModel struct {
	client *ent.Client
}

func NewPlayerStatsModel(client *ent.Client) *PlayerStatsModel {
	return &PlayerStatsModel{client: client}
}

func (m *PlayerStatsModel) UpsertPlayerStats(ctx context.Context, p *ent.Player, s PlayerStats) error {
	t, err := m.client.Team.
		Query().
		WithClub(
			func(q *ent.ClubQuery) {
				q.Where(club.ApiFootballId(s.Team.ApiFootballId))
			}).
		WithSeason(
			func(q *ent.SeasonQuery) {
				q.Where(season.Year(s.League.Season))
			}).Only(ctx)

	ps, err := m.client.PlayerStats.
		Query().
		WithPlayer(
			func(q *ent.PlayerQuery) {
				q.Where(player.IDEQ(p.ID))
			},
		).
		WithTeam(
			func(q *ent.TeamQuery) {
				q.Where(team.IDEQ(t.ID))
			},
		).
		Only(ctx)
	if ps != nil {
		_, err = ps.Update().
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		ps, err = m.client.PlayerStats.
			Create().
			SetPlayer(p).
			SetTeam(t).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	_, err = m.upsertDefense(ctx, ps, s.Duels, s.Tackles)
	if err != nil {
		return err
	}

	_, err = m.upsertGames(ctx, ps, s.Games)
	if err != nil {
		return err
	}

	_, err = m.upsertShooting(ctx, ps, s.Goals, s.Shots)
	if err != nil {
		return err
	}

	_, err = m.upsertPenalty(ctx, ps, s.Penalty, s.Fouls)
	if err != nil {
		return err
	}

	_, err = m.upsertFairplay(ctx, ps, s.Cards, s.Fouls, s.Penalty)
	if err != nil {
		return err
	}

	_, err = m.upsertSubstitutes(ctx, ps, s.Substitutes)
	if err != nil {
		return err
	}

	_, err = m.upsertTechnical(ctx, ps, s.Passes, s.Dribbles, s.Fouls)
	if err != nil {
		return err
	}
	return nil
}

// upsertDefense inserts or updates a player's defense stats in the database.
func (m *PlayerStatsModel) upsertDefense(ctx context.Context, ps *ent.PlayerStats, duels PSDuels, tackles PSTackles) (*ent.PSDefense, error) {
	def, err := m.client.PSDefense.
		Query().
		WithPlayerStats(
			func(q *ent.PlayerStatsQuery) {
				q.Where(playerstats.IDEQ(ps.ID))
			},
		).
		Only(ctx)

	if def != nil {
		_, err = def.Update().
			SetTacklesTotal(tackles.Total).
			SetBlocks(*tackles.Blocks).
			SetInterceptions(tackles.Interceptions).
			SetDuelsTotal(duels.Total).
			SetWonDuels(duels.Won).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return def, nil
	} else {
		def, err := m.client.PSDefense.
			Create().
			SetTacklesTotal(tackles.Total).
			SetBlocks(*tackles.Blocks).
			SetInterceptions(tackles.Interceptions).
			SetDuelsTotal(duels.Total).
			SetWonDuels(duels.Won).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return def, nil
	}
}

// upsertGames inserts or updates a player's games stats in the database.
func (m *PlayerStatsModel) upsertGames(ctx context.Context, ps *ent.PlayerStats, games PSGames) (*ent.PSGames, error) {
	g, err := m.client.PSGames.
		Query().
		WithPlayerStats(
			func(q *ent.PlayerStatsQuery) {
				q.Where(playerstats.IDEQ(ps.ID))
			},
		).
		Only(ctx)

	if g != nil {
		_, err = g.Update().
			SetAppearances(games.Appearances).
			SetLineups(games.Lineups).
			SetMinutes(games.Minutes).
			SetNumber(*games.Number).
			SetPosition(games.Position).
			SetRating(games.Rating).
			SetCaptain(games.Captain).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return g, nil
	} else {
		g, err := m.client.PSGames.
			Create().
			SetAppearances(games.Appearances).
			SetLineups(games.Lineups).
			SetMinutes(games.Minutes).
			SetNumber(*games.Number).
			SetPosition(games.Position).
			SetRating(games.Rating).
			SetCaptain(games.Captain).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return g, nil
	}
}

// upsertShooting inserts or updates a player's shooting stats in the database.
func (m *PlayerStatsModel) upsertShooting(ctx context.Context, ps *ent.PlayerStats, goals PSGoals, shots PSShots) (*ent.PSShooting, error) {
	g, err := m.client.PSShooting.
		Query().
		WithPlayerStats(
			func(q *ent.PlayerStatsQuery) {
				q.Where(playerstats.IDEQ(ps.ID))
			},
		).
		Only(ctx)

	if g != nil {
		_, err = g.Update().
			SetGoals(goals.Total).
			SetConceded(goals.Conceded).
			SetAssists(goals.Assists).
			SetSaves(*goals.Saves).
			SetShots(shots.Total).
			SetOnTarget(shots.On).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return g, nil
	} else {
		g, err := m.client.PSShooting.
			Create().
			SetGoals(goals.Total).
			SetConceded(goals.Conceded).
			SetAssists(goals.Assists).
			SetSaves(*goals.Saves).
			SetShots(shots.Total).
			SetOnTarget(shots.On).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return g, nil
	}
}

// upsertPenalty inserts or updates a player's penalty and foul stats in the database.
func (m *PlayerStatsModel) upsertPenalty(ctx context.Context, ps *ent.PlayerStats, penalty PSPenalty, fouls PSFouls) (*ent.PSPenalty, error) {
	p, err := m.client.PSPenalty.
		Query().
		WithPlayerStats(
			func(q *ent.PlayerStatsQuery) {
				q.Where(playerstats.IDEQ(ps.ID))
			},
		).
		Only(ctx)

	if p != nil {
		_, err = p.Update().
			SetWon(*penalty.Won).
			SetScored(penalty.Scored).
			SetMissed(penalty.Missed).
			SetSaved(*penalty.Saved).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		p, err := m.client.PSPenalty.
			Create().
			SetWon(*penalty.Won).
			SetScored(penalty.Scored).
			SetMissed(penalty.Missed).
			SetSaved(*penalty.Saved).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return p, nil
	}
}

// upsertFairplay inserts or updates a player's fairplay stats in the database.
func (m *PlayerStatsModel) upsertFairplay(ctx context.Context, ps *ent.PlayerStats, cards PSCards, fouls PSFouls, penalty PSPenalty) (*ent.PSFairplay, error) {
	f, err := m.client.PSFairplay.
		Query().
		WithPlayerStats(
			func(q *ent.PlayerStatsQuery) {
				q.Where(playerstats.IDEQ(ps.ID))
			},
		).
		Only(ctx)

	if f != nil {
		_, err = f.Update().
			SetFoulsCommitted(fouls.Committed).
			SetYellow(cards.Yellow).
			SetYellowRed(cards.YellowRed).
			SetRed(cards.Red).
			SetPenaltyConceded(*penalty.Committed).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return f, nil
	} else {
		f, err := m.client.PSFairplay.
			Create().
			SetFoulsCommitted(fouls.Committed).
			SetYellow(cards.Yellow).
			SetYellowRed(cards.YellowRed).
			SetRed(cards.Red).
			SetPenaltyConceded(*penalty.Committed).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
}

// upsertSubstitutes inserts or updates a player's substitutes stats in the database.
func (m *PlayerStatsModel) upsertSubstitutes(ctx context.Context, ps *ent.PlayerStats, subs PSSubstitutes) (*ent.PSSubstitutes, error) {
	s, err := m.client.PSSubstitutes.
		Query().
		WithPlayerStats(
			func(q *ent.PlayerStatsQuery) {
				q.Where(playerstats.IDEQ(ps.ID))
			},
		).
		Only(ctx)

	if s != nil {
		_, err = s.Update().
			SetIn(subs.In).
			SetOut(subs.Out).
			SetBench(subs.Bench).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return s, nil
	} else {
		s, err := m.client.PSSubstitutes.
			Create().
			SetIn(subs.In).
			SetOut(subs.Out).
			SetBench(subs.Bench).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return s, nil
	}
}

// upsertTechnical inserts or updates a player's technical stats in the database.
func (m *PlayerStatsModel) upsertTechnical(ctx context.Context, ps *ent.PlayerStats, passes PSPasses, dribbles PSDribbles, fouls PSFouls) (*ent.PSTechnical, error) {
	t, err := m.client.PSTechnical.
		Query().
		WithPlayerStats(
			func(q *ent.PlayerStatsQuery) {
				q.Where(playerstats.IDEQ(ps.ID))
			},
		).
		Only(ctx)

	if t != nil {
		_, err = t.Update().
			SetFoulsDrawn(fouls.Drawn).
			SetDribbleAttempts(dribbles.Attempts).
			SetDribbleSuccess(dribbles.Success).
			SetDribblePast(*dribbles.Past).
			SetPassesTotal(passes.Total).
			SetPassesKey(passes.Key).
			SetPassesAccuracy(passes.Accuracy).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return t, nil
	} else {
		t, err := m.client.PSTechnical.
			Create().
			SetFoulsDrawn(fouls.Drawn).
			SetDribbleAttempts(dribbles.Attempts).
			SetDribbleSuccess(dribbles.Success).
			SetDribblePast(*dribbles.Past).
			SetPassesTotal(passes.Total).
			SetPassesKey(passes.Key).
			SetPassesAccuracy(passes.Accuracy).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
}
