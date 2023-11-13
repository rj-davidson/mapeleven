package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psshooting"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pstechnical"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"context"
	_ "context"
	"fmt"
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
	Appearances *int   `json:"appearances"`
	Lineups     *int   `json:"lineups"`
	Minutes     *int   `json:"minutes"`
	Number      *int   `json:"number"`
	Position    string `json:"position"`
	Rating      string `json:"rating"`
	Captain     bool   `json:"captain"`
}

type PSSubstitutes struct {
	In    *int `json:"in"`
	Out   *int `json:"out"`
	Bench *int `json:"bench"`
}

type PSShots struct {
	Total *int `json:"total"`
	On    *int `json:"on"`
}

type PSGoals struct {
	Total    *int `json:"total"`
	Conceded *int `json:"conceded"`
	Assists  *int `json:"assists"`
	Saves    *int `json:"saves"`
}

type PSPasses struct {
	Total    *int `json:"total"`
	Key      *int `json:"key"`
	Accuracy *int `json:"accuracy"`
}

type PSTackles struct {
	Total         *int `json:"total"`
	Blocks        *int `json:"blocks"`
	Interceptions *int `json:"interceptions"`
}

type PSDuels struct {
	Total *int `json:"total"`
	Won   *int `json:"won"`
}

type PSDribbles struct {
	Attempts *int `json:"attempts"`
	Success  *int `json:"success"`
	Past     *int `json:"past"`
}

type PSFouls struct {
	Drawn     *int `json:"drawn"`
	Committed *int `json:"committed"`
}

type PSCards struct {
	Yellow    *int `json:"yellow"`
	YellowRed *int `json:"yellowred"`
	Red       *int `json:"red"`
}

type PSPenalty struct {
	Won       *int `json:"won"`
	Committed *int `json:"commited"`
	Scored    *int `json:"scored"`
	Missed    *int `json:"missed"`
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
func intValueOrDefault(val *int) int {
	if val != nil {
		return *val
	}
	return 0
}

func (m *PlayerStatsModel) StartTransaction(ctx context.Context) (*ent.Tx, error) {
	return m.client.Tx(ctx)
}

// UpsertPlayerStats updates existing player stats or creates new ones.
func (m *PlayerStatsModel) UpsertPlayerStats(ctx context.Context, tx *ent.Tx, p *ent.Player, stats PlayerStats) error {
	t, err := m.client.Team.
		Query().
		WithClub(func(q *ent.ClubQuery) {
			q.Where(club.ApiFootballIdEQ(stats.Team.ApiFootballId))
		},
		).
		WithSeason(func(q *ent.SeasonQuery) {
			q.Where(season.YearEQ(stats.League.Season))
			q.Where(season.HasLeagueWith(league.FootballApiIdEQ(stats.League.ApiFootballId)))
		},
		).
		All(ctx)
	fmt.Printf("Team: %v\n", t)
	if err != nil {
		return fmt.Errorf("failed to query team: %w", err)
	}
	// Query for existing player stats.
	ps, err := tx.PlayerStats.
		Query().
		Where(playerstats.HasPlayerWith(player.IDEQ(p.ID))).
		WithPlayer().
		Only(ctx)
	fmt.Printf("Player stats: %v\n", stats)
	// Handle not found error by creating a new player stats record.
	if ps == nil {
		ps, err = tx.PlayerStats.
			Create().
			SetPlayer(p).
			//SetTeam(t).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create player stats: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to query player stats: %w", err)
	}

	// Upsert sub-statistics using the individual wrapper functions.
	err = m.upsertSubStats(ctx, ps, stats)
	if err != nil {
		return fmt.Errorf("failed to upsert sub-stats: %w", err)
	}

	return nil
}

func (m *PlayerStatsModel) upsertDefenseWrapper(ctx context.Context, ps *ent.PlayerStats, s PlayerStats) error {
	_, err := m.upsertDefense(ctx, ps, s.Duels, s.Tackles)
	return err
}
func (m *PlayerStatsModel) upsertGamesWrapper(ctx context.Context, ps *ent.PlayerStats, s PlayerStats) error {
	_, err := m.upsertGames(ctx, ps, s.Games)
	return err
}
func (m *PlayerStatsModel) upsertShootingWrapper(ctx context.Context, ps *ent.PlayerStats, s PlayerStats) error {
	_, err := m.upsertShooting(ctx, ps, s.Goals, s.Shots)
	return err
}
func (m *PlayerStatsModel) upsertPenaltyWrapper(ctx context.Context, ps *ent.PlayerStats, s PlayerStats) error {
	_, err := m.upsertPenalty(ctx, ps, s.Penalty, s.Fouls)
	return err
}
func (m *PlayerStatsModel) upsertFairplayWrapper(ctx context.Context, ps *ent.PlayerStats, s PlayerStats) error {
	_, err := m.upsertFairplay(ctx, ps, s.Cards, s.Fouls, s.Penalty)
	return err
}
func (m *PlayerStatsModel) upsertSubstitutesWrapper(ctx context.Context, ps *ent.PlayerStats, s PlayerStats) error {
	_, err := m.upsertSubstitutes(ctx, ps, s.Substitutes)
	return err
}
func (m *PlayerStatsModel) upsertTechnicalWrapper(ctx context.Context, ps *ent.PlayerStats, s PlayerStats) error {
	_, err := m.upsertTechnical(ctx, ps, s.Passes, s.Dribbles, s.Fouls)
	return err
}

// upsertSubStats upserts all of the sub-statistics for a player.
func (m *PlayerStatsModel) upsertSubStats(ctx context.Context, ps *ent.PlayerStats, s PlayerStats) error {
	err := m.upsertDefenseWrapper(ctx, ps, s)
	if err != nil {
		return fmt.Errorf("failed to upsert defense: %w", err)
	}

	err = m.upsertGamesWrapper(ctx, ps, s)
	if err != nil {
		return fmt.Errorf("failed to upsert games: %w", err)
	}

	err = m.upsertShootingWrapper(ctx, ps, s)
	if err != nil {
		return fmt.Errorf("failed to upsert shooting: %w", err)
	}

	err = m.upsertPenaltyWrapper(ctx, ps, s)
	if err != nil {
		return fmt.Errorf("failed to upsert penalty: %w", err)
	}

	err = m.upsertFairplayWrapper(ctx, ps, s)
	if err != nil {
		return fmt.Errorf("failed to upsert fairplay: %w", err)
	}

	err = m.upsertSubstitutesWrapper(ctx, ps, s)
	if err != nil {
		return fmt.Errorf("failed to upsert substitutes: %w", err)
	}
	err = m.upsertTechnicalWrapper(ctx, ps, s)
	if err != nil {
		return fmt.Errorf("failed to upsert technical: %w", err)
	}

	fmt.Printf("Finished upserting sub-stats for %s\n", ps.Edges.Player.Name)
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

	tacklesTotal := intValueOrDefault(tackles.Total)
	blocks := intValueOrDefault(tackles.Blocks)
	interceptions := intValueOrDefault(tackles.Interceptions)
	duelsTotal := intValueOrDefault(duels.Total)
	duelsWon := intValueOrDefault(duels.Won)
	if def != nil {
		_, err = def.Update().
			SetTacklesTotal(tacklesTotal).
			SetBlocks(blocks).
			SetInterceptions(interceptions).
			SetDuelsTotal(duelsTotal).
			SetWonDuels(duelsWon).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return def, nil
	} else {
		def, err := m.client.PSDefense.
			Create().
			SetTacklesTotal(tacklesTotal).
			SetBlocks(blocks).
			SetInterceptions(interceptions).
			SetDuelsTotal(duelsTotal).
			SetWonDuels(duelsWon).
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
	appearances := intValueOrDefault(games.Appearances)
	lineups := intValueOrDefault(games.Lineups)
	minutes := intValueOrDefault(games.Minutes)
	number := intValueOrDefault(games.Number)

	if g != nil {
		_, err = g.Update().
			SetAppearances(appearances).
			SetLineups(lineups).
			SetMinutes(minutes).
			SetNumber(number).
			SetPosition(games.Position).
			SetRating(games.Rating).
			SetCaptain(games.Captain).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return g, nil
	} else {
		g, err := m.client.PSGames.
			Create().
			SetAppearances(appearances).
			SetLineups(lineups).
			SetMinutes(minutes).
			SetNumber(number).
			SetPosition(games.Position).
			SetRating(games.Rating).
			SetCaptain(games.Captain).
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

	goalsTotal := intValueOrDefault(goals.Total)
	conceded := intValueOrDefault(goals.Conceded)
	assists := intValueOrDefault(goals.Assists)
	saves := intValueOrDefault(goals.Saves)
	shotsTotal := intValueOrDefault(shots.Total)
	shotsOn := intValueOrDefault(shots.On)

	if g != nil {
		_, err = g.Update().
			SetGoals(goalsTotal).
			SetConceded(conceded).
			SetAssists(assists).
			SetSaves(saves).
			SetShots(shotsTotal).
			SetOnTarget(shotsOn).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return g, nil
	} else {
		g, err := m.client.PSShooting.
			Create().
			SetGoals(goalsTotal).
			SetConceded(conceded).
			SetAssists(assists).
			SetSaves(saves).
			SetShots(shotsTotal).
			SetOnTarget(shotsOn).
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
	won := intValueOrDefault(penalty.Won)
	scored := intValueOrDefault(penalty.Scored)
	missed := intValueOrDefault(penalty.Missed)
	saved := intValueOrDefault(penalty.Saved)
	committed := intValueOrDefault(penalty.Committed)

	if p != nil {
		_, err = p.Update().
			SetWon(won).
			SetScored(scored).
			SetMissed(missed).
			SetSaved(saved).
			SetCommitted(committed).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		p, err := m.client.PSPenalty.
			Create().
			SetWon(won).
			SetScored(scored).
			SetMissed(missed).
			SetSaved(saved).
			SetCommitted(committed).
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
	foulsCommitted := intValueOrDefault(fouls.Committed)
	yellow := intValueOrDefault(cards.Yellow)
	yellowRed := intValueOrDefault(cards.YellowRed)
	red := intValueOrDefault(cards.Red)
	penaltyCommitted := intValueOrDefault(penalty.Committed)

	if f != nil {
		_, err = f.Update().
			SetFoulsCommitted(foulsCommitted).
			SetYellow(yellow).
			SetYellowRed(yellowRed).
			SetRed(red).
			SetPenaltyConceded(penaltyCommitted).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return f, nil
	} else {
		f, err := m.client.PSFairplay.
			Create().
			SetFoulsCommitted(foulsCommitted).
			SetYellow(yellow).
			SetYellowRed(yellowRed).
			SetRed(red).
			SetPenaltyConceded(penaltyCommitted).
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

	in := intValueOrDefault(subs.In)
	out := intValueOrDefault(subs.Out)
	bench := intValueOrDefault(subs.Bench)

	if s != nil {
		_, err = s.Update().
			SetIn(in).
			SetOut(out).
			SetBench(bench).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return s, nil
	} else {
		s, err := m.client.PSSubstitutes.
			Create().
			SetIn(in).
			SetOut(out).
			SetBench(bench).
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
	foulsDrawn := intValueOrDefault(fouls.Drawn)
	dribblesAttempts := intValueOrDefault(dribbles.Attempts)
	dribblesSuccess := intValueOrDefault(dribbles.Success)
	dribblesPast := intValueOrDefault(dribbles.Past)
	passesTotal := intValueOrDefault(passes.Total)
	passesKey := intValueOrDefault(passes.Key)
	passesAccuracy := intValueOrDefault(passes.Accuracy)

	if t != nil {
		_, err = t.Update().
			SetFoulsDrawn(foulsDrawn).
			SetDribbleAttempts(dribblesAttempts).
			SetDribbleSuccess(dribblesSuccess).
			SetDribblePast(dribblesPast).
			SetPassesTotal(passesTotal).
			SetPassesKey(passesKey).
			SetPassesAccuracy(passesAccuracy).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return t, nil
	} else {
		t, err := m.client.PSTechnical.
			Create().
			SetFoulsDrawn(foulsDrawn).
			SetDribbleAttempts(dribblesAttempts).
			SetDribbleSuccess(dribblesSuccess).
			SetDribblePast(dribblesPast).
			SetPassesTotal(passesTotal).
			SetPassesKey(passesKey).
			SetPassesAccuracy(passesAccuracy).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
}
