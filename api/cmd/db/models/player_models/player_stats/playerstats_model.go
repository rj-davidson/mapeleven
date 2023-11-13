package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psshooting"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pstechnical"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
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

func (m *PlayerStatsModel) StartTransaction(ctx context.Context) (*ent.Tx, error) {
	return m.client.Tx(ctx)
}

// Exists checks if the player statistics already exist for a given player and season.
func (m *PlayerStatsModel) Exists(ctx context.Context, playerID int, szn int) (bool, error) {
	return m.client.PlayerStats.Query().
		Where(playerstats.IDEQ(playerID), playerstats.HasSeasonWith(season.Year(szn))).
		Exist(ctx)
}

// CreatePlayerStats creates a new player statistics record along with related entities like PSDefense and PSFairplay.
func (m *PlayerStatsModel) CreatePlayerStats(ctx context.Context, playerID int, stats PlayerStats) (*ent.PlayerStats, error) {
	// Start transaction
	tx, err := m.StartTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	ps, err := tx.PlayerStats.
		Create().
		SetLastUpdated(time.Now()).
		SetPlayerID(playerID).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating PlayerStats: %w", err)
	}
	if ps == nil {
		fmt.Printf("ps is nil: %v\n", ps)
	}

	// Example of creating a related entity, adjust as needed for other entities
	_, err = tx.PSDefense.
		Create().
		SetBlocks(intValueOrDefault(stats.Tackles.Blocks)).
		SetInterceptions(intValueOrDefault(stats.Tackles.Interceptions)).
		SetTacklesTotal(intValueOrDefault(stats.Tackles.Total)).
		SetDuelsTotal(intValueOrDefault(stats.Duels.Total)).
		SetWonDuels(intValueOrDefault(stats.Duels.Won)).
		SetPlayerStats(ps).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = tx.PSFairplay.
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
		tx.Rollback()
		return nil, err
	}

	_, err = tx.PSGames.
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
		tx.Rollback()
		return nil, err
	}
	_, err = tx.PSSubstitutes.
		Create().
		SetIn(intValueOrDefault(stats.Substitutes.In)).
		SetOut(intValueOrDefault(stats.Substitutes.Out)).
		SetBench(intValueOrDefault(stats.Substitutes.Bench)).
		SetLastUpdated(time.Now()).
		SetPlayerStats(ps).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = tx.PSShooting.
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
		tx.Rollback()
		return nil, err
	}
	_, err = tx.PSTechnical.
		Create().
		SetPassesTotal(intValueOrDefault(stats.Passes.Total)).
		SetPassesKey(intValueOrDefault(stats.Passes.Key)).
		SetPassesAccuracy(intValueOrDefault(stats.Passes.Accuracy)).
		SetDribbleAttempts(intValueOrDefault(stats.Dribbles.Attempts)).
		SetDribbleSuccess(intValueOrDefault(stats.Dribbles.Success)).
		SetDribblePast(intValueOrDefault(stats.Dribbles.Past)).
		SetLastUpdated(time.Now()).
		SetPlayerStats(ps).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.PSPenalty.
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
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	return ps, nil
}

// UpdatePlayerStats updates an existing player statistics record along with related entities like PSDefense and PSFairplay.
func (m *PlayerStatsModel) UpdatePlayerStats(ctx context.Context, playerStatsID int, stats PlayerStats) (*ent.PlayerStats, error) {
	// Start transaction
	tx, err := m.StartTransaction(ctx)
	if err != nil {
		return nil, err
	}

	ps, err := tx.PlayerStats.
		UpdateOneID(playerStatsID).
		SetLastUpdated(time.Now()).
		SetPlayerID(playerStatsID).
		SetTeamID(stats.Team.ApiFootballId).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Example of updating a related entity, adjust as needed for other entities
	_, err = tx.PSDefense.
		Update().
		Where(psdefense.HasPlayerStatsWith(playerstats.IDEQ(playerStatsID))).
		SetBlocks(intValueOrDefault(stats.Tackles.Blocks)).
		SetInterceptions(intValueOrDefault(stats.Tackles.Interceptions)).
		SetTacklesTotal(intValueOrDefault(stats.Tackles.Total)).
		SetDuelsTotal(intValueOrDefault(stats.Duels.Total)).
		SetWonDuels(intValueOrDefault(stats.Duels.Won)).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return ps, nil
}
