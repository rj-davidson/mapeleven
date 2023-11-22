package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psfairplay"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psshooting"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psshooting"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pssubstitutes"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pstechnical"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pstechnical"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"context"
	_ "context"
	"fmt"
	"time"
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

// Exists checks if the player statistics already exist for a given player and season.
func (m *PlayerStatsModel) Exists(ctx context.Context, p *ent.Player, szn int) (bool, *ent.PlayerStats) {
	ps, _ := m.client.PlayerStats.Query().
		Where(playerstats.IDEQ(p.ID), playerstats.HasSeasonWith(season.Year(szn))).
		Only(ctx)
	if ps == nil {
		return false, nil
	}
	return true, ps
}

// UpsertPlayerStats creates a new player statistics record along with related entities like PSDefense and PSFairplay.
func (m *PlayerStatsModel) UpsertPlayerStats(ctx context.Context, p *ent.Player, stats PlayerStats) (*ent.PlayerStats, error) {
	// Check if league exists
	l, err := m.client.League.
		Query().
		Where(league.FootballApiIdEQ(stats.League.ApiFootballId)).
		Only(ctx)
	if l == nil { // Skipping leagues that don't exist in our database
		return nil, nil
	}

	t, err := m.client.Team.
		Query().
		Where(
			team.HasClubWith(
				club.ApiFootballIdEQ(stats.Team.ApiFootballId),
			),
			team.HasSeasonWith(
				season.YearEQ(stats.League.Season),
				season.HasLeagueWith(league.FootballApiIdEQ(stats.League.ApiFootballId)),
			),
		).
		Only(ctx)
	if err != nil {
		fmt.Printf("Error fetching team for playerStats: %v\n", err)
		return nil, nil
	}

	// Check if player stats already exist
	exists, ps := m.Exists(ctx, p, stats.League.Season)
	if !exists {
		ps, err = m.client.PlayerStats.
			Create().
			SetPlayer(p).
			SetTeam(t).
			Save(ctx)
		if err != nil {
			fmt.Printf("Error creating playerStats: %v\n", err)
			return nil, err
		}
	}
	err = m.upsertDefense(ctx, ps, stats)
	if err != nil {
		fmt.Printf("Error upserting defense: %v\n", err)
	}
	err = m.upsertFairplay(ctx, ps, stats)
	if err != nil {
		fmt.Printf("Error upserting fairplay: %v\n", err)
	}
	err = m.upsertGames(ctx, ps, stats)
	if err != nil {
		fmt.Printf("Error upserting games: %v\n", err)
	}
	err = m.upsertPenalty(ctx, ps, stats)
	if err != nil {
		fmt.Printf("Error upserting penalty: %v\n", err)
	}
	err = m.upsertShooting(ctx, ps, stats)
	if err != nil {
		fmt.Printf("Error upserting shooting: %v\n", err)
	}
	err = m.upsertSubstitutes(ctx, ps, stats)
	if err != nil {
		fmt.Printf("Error upserting substitutes: %v\n", err)
	}
	err = m.upsertTechnical(ctx, ps, stats)
	if err != nil {
		fmt.Printf("Error upserting technical: %v\n", err)
	}
	return ps, nil
}

func (m *PlayerStatsModel) upsertDefense(ctx context.Context, ps *ent.PlayerStats, stats PlayerStats) error {
	defense, _ := m.client.PSDefense.
		Query().
		Where(psdefense.HasPlayerStatsWith(playerstats.IDEQ(ps.ID))).
		Only(ctx)
	if defense == nil {
		_, err := m.client.PSDefense.
			Create().
			SetBlocks(intValueOrDefault(stats.Tackles.Blocks)).
			SetInterceptions(intValueOrDefault(stats.Tackles.Interceptions)).
			SetTacklesTotal(intValueOrDefault(stats.Tackles.Total)).
			SetDribblePast(intValueOrDefault(stats.Dribbles.Past)).
			SetDuelsTotal(intValueOrDefault(stats.Duels.Total)).
			SetWonDuels(intValueOrDefault(stats.Duels.Won)).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := m.client.PSDefense.
			UpdateOne(defense).
			SetBlocks(intValueOrDefault(stats.Tackles.Blocks)).
			SetInterceptions(intValueOrDefault(stats.Tackles.Interceptions)).
			SetTacklesTotal(intValueOrDefault(stats.Tackles.Total)).
			SetDribblePast(intValueOrDefault(stats.Dribbles.Past)).
			SetDuelsTotal(intValueOrDefault(stats.Duels.Total)).
			SetWonDuels(intValueOrDefault(stats.Duels.Won)).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *PlayerStatsModel) upsertFairplay(ctx context.Context, ps *ent.PlayerStats, stats PlayerStats) error {
	fairplay, _ := m.client.PSFairplay.
		Query().
		Where(psfairplay.HasPlayerStatsWith(playerstats.IDEQ(ps.ID))).
		Only(ctx)
	if fairplay == nil {
		_, err := m.client.PSFairplay.
			Create().
			SetYellow(intValueOrDefault(stats.Cards.Yellow)).
			SetYellowRed(intValueOrDefault(stats.Cards.YellowRed)).
			SetRed(intValueOrDefault(stats.Cards.Red)).
			SetFoulsCommitted(intValueOrDefault(stats.Fouls.Committed)).
			SetPenaltyConceded(intValueOrDefault(stats.Penalty.Committed)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := m.client.PSFairplay.
			UpdateOne(fairplay).
			SetYellow(intValueOrDefault(stats.Cards.Yellow)).
			SetYellowRed(intValueOrDefault(stats.Cards.YellowRed)).
			SetRed(intValueOrDefault(stats.Cards.Red)).
			SetFoulsCommitted(intValueOrDefault(stats.Fouls.Committed)).
			SetPenaltyConceded(intValueOrDefault(stats.Penalty.Committed)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *PlayerStatsModel) upsertGames(ctx context.Context, ps *ent.PlayerStats, stats PlayerStats) error {
	games, _ := m.client.PSGames.
		Query().
		Where(psgames.HasPlayerStatsWith(playerstats.IDEQ(ps.ID))).
		Only(ctx)
	if games == nil {
		_, err := m.client.PSGames.
			Create().
			SetAppearances(intValueOrDefault(stats.Games.Appearances)).
			SetLineups(intValueOrDefault(stats.Games.Lineups)).
			SetMinutes(intValueOrDefault(stats.Games.Minutes)).
			SetNumber(intValueOrDefault(stats.Games.Number)).
			SetPosition(stats.Games.Position).
			SetRating(stats.Games.Rating).
			SetCaptain(stats.Games.Captain).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := m.client.PSGames.
			UpdateOne(games).
			SetAppearances(intValueOrDefault(stats.Games.Appearances)).
			SetLineups(intValueOrDefault(stats.Games.Lineups)).
			SetMinutes(intValueOrDefault(stats.Games.Minutes)).
			SetNumber(intValueOrDefault(stats.Games.Number)).
			SetPosition(stats.Games.Position).
			SetRating(stats.Games.Rating).
			SetCaptain(stats.Games.Captain).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *PlayerStatsModel) upsertPenalty(ctx context.Context, ps *ent.PlayerStats, stats PlayerStats) error {
	penalty, _ := m.client.PSPenalty.
		Query().
		Where(pspenalty.HasPlayerStatsWith(playerstats.IDEQ(ps.ID))).
		Only(ctx)
	if penalty == nil {
		_, err := m.client.PSPenalty.
			Create().
			SetWon(intValueOrDefault(stats.Penalty.Won)).
			SetCommitted(intValueOrDefault(stats.Penalty.Committed)).
			SetScored(intValueOrDefault(stats.Penalty.Scored)).
			SetMissed(intValueOrDefault(stats.Penalty.Missed)).
			SetSaved(intValueOrDefault(stats.Penalty.Saved)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := m.client.PSPenalty.
			UpdateOne(penalty).
			SetWon(intValueOrDefault(stats.Penalty.Won)).
			SetCommitted(intValueOrDefault(stats.Penalty.Committed)).
			SetScored(intValueOrDefault(stats.Penalty.Scored)).
			SetMissed(intValueOrDefault(stats.Penalty.Missed)).
			SetSaved(intValueOrDefault(stats.Penalty.Saved)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *PlayerStatsModel) upsertShooting(ctx context.Context, ps *ent.PlayerStats, stats PlayerStats) error {
	shooting, _ := m.client.PSShooting.
		Query().
		Where(psshooting.HasPlayerStatsWith(playerstats.IDEQ(ps.ID))).
		Only(ctx)
	if shooting == nil {
		_, err := m.client.PSShooting.
			Create().
			SetShots(intValueOrDefault(stats.Shots.Total)).
			SetOnTarget(intValueOrDefault(stats.Shots.On)).
			SetAssists(intValueOrDefault(stats.Goals.Assists)).
			SetGoals(intValueOrDefault(stats.Goals.Total)).
			SetSaves(intValueOrDefault(stats.Goals.Saves)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := m.client.PSShooting.
			UpdateOne(shooting).
			SetShots(intValueOrDefault(stats.Shots.Total)).
			SetOnTarget(intValueOrDefault(stats.Shots.On)).
			SetAssists(intValueOrDefault(stats.Goals.Assists)).
			SetGoals(intValueOrDefault(stats.Goals.Total)).
			SetSaves(intValueOrDefault(stats.Goals.Saves)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *PlayerStatsModel) upsertSubstitutes(ctx context.Context, ps *ent.PlayerStats, stats PlayerStats) error {
	substitutes, _ := m.client.PSSubstitutes.
		Query().
		Where(pssubstitutes.HasPlayerStatsWith(playerstats.IDEQ(ps.ID))).
		Only(ctx)
	if substitutes == nil {
		_, err := m.client.PSSubstitutes.
			Create().
			SetIn(intValueOrDefault(stats.Substitutes.In)).
			SetOut(intValueOrDefault(stats.Substitutes.Out)).
			SetBench(intValueOrDefault(stats.Substitutes.Bench)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := m.client.PSSubstitutes.
			UpdateOne(substitutes).
			SetIn(intValueOrDefault(stats.Substitutes.In)).
			SetOut(intValueOrDefault(stats.Substitutes.Out)).
			SetBench(intValueOrDefault(stats.Substitutes.Bench)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *PlayerStatsModel) upsertTechnical(ctx context.Context, ps *ent.PlayerStats, stats PlayerStats) error {
	technical, _ := m.client.PSTechnical.
		Query().
		Where(pstechnical.HasPlayerStatsWith(playerstats.IDEQ(ps.ID))).
		Only(ctx)
	if technical == nil {
		_, err := m.client.PSTechnical.
			Create().
			SetPassesTotal(intValueOrDefault(stats.Passes.Total)).
			SetPassesKey(intValueOrDefault(stats.Passes.Key)).
			SetPassesAccuracy(intValueOrDefault(stats.Passes.Accuracy)).
			SetDribbleAttempts(intValueOrDefault(stats.Dribbles.Attempts)).
			SetDribbleSuccess(intValueOrDefault(stats.Dribbles.Success)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := m.client.PSTechnical.
			UpdateOne(technical).
			SetPassesTotal(intValueOrDefault(stats.Passes.Total)).
			SetPassesKey(intValueOrDefault(stats.Passes.Key)).
			SetPassesAccuracy(intValueOrDefault(stats.Passes.Accuracy)).
			SetDribbleAttempts(intValueOrDefault(stats.Dribbles.Attempts)).
			SetDribbleSuccess(intValueOrDefault(stats.Dribbles.Success)).
			SetLastUpdated(time.Now()).
			SetPlayerStats(ps).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
