package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgoals"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psoffense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	"context"
	_ "context"
)

//type PlayerStats struct {
//	PSPenalty PSPenalty `json:"penalty"`
//	PSDefense PSDefense `json:"defense"`
//	PSGoals   PSGoals   `json:"goals"`
//	PSOffense PSOffense `json:"offense"`
//	PSGames   PSGames   `json:"games"`
//}
//type PSPenalty struct {
//	FoulsDrawn      int `json:"fouls_drawn"`
//	FoulsCommitted  int `json:"fouls_committed"`
//	CardsYellow     int `json:"cards_yellow"`
//	CardYellowred   int `json:"card_yellowred"`
//	CardsRed        int `json:"cards_red"`
//	PenaltyWon      int `json:"penalty_won"`
//	PenaltyCommited int `json:"penalty_commited"`
//	PenaltyScored   int `json:"penalty_scored"`
//	PenaltyMissed   int `json:"penalty_missed"`
//	PenaltySaved    int `json:"penalty_saved"`
//}
//type PSDefense struct {
//	TacklesTotal  int `json:"tackles_total"`
//	Blocks        int `json:"blocks"`
//	Interceptions int `json:"interceptions"`
//	TotalDuels    int `json:"total_duels"`
//	WonDuels      int `json:"won_duels"`
//}
//
//type PSGoals struct {
//	TotalGoals    int `json:"total_goals"`
//	ConcededGoals int `json:"conceded_goals"`
//	AssistGoals   int `json:"assist_goals"`
//	SaveGoals     int `json:"save_goals"`
//	ShotsTotal    int `json:"shots_total"`
//	ShotsOn       int `json:"shots_on"`
//}
//
//type PSOffense struct {
//	DribbleAttempts int `json:"dribbles_attempts"`
//	DribbleSuccess  int `json:"dribbles_success"`
//	DribblePast     int `json:"dribbled_past"`
//	PassesTotal     int `json:"passes_total"`
//	PassesKey       int `json:"passes_key"`
//	PassesAccuracy  int `json:"passes_accuracy"`
//}

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
	Won      *int `json:"won"`
	Commited *int `json:"commited"`
	Scored   int  `json:"scored"`
	Missed   int  `json:"missed"`
	Saved    *int `json:"saved"`
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

func (m *PlayerStatsModel) UpsertPlayerStats(ctx context.Context, p *ent.Player, stats *ent.PlayerStats) (*ent.PlayerStats, error) {
	_, err := m.client.PlayerStats.
		UpdateOne(stats).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

// WithTransaction method to execute a transaction
func (m *PlayerStatsModel) WithTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error {
	// Start a new transaction
	tx, err := m.client.Tx(ctx)
	if err != nil {
		return err
	}
	// Execute logic
	if err := fn(tx); err != nil {
		_ = tx.Rollback() // rollback the transaction in case of errors
		return err
	}
	return tx.Commit() // commit the transaction
}

// write the method to create the playerstats package
//func (m *PlayerStatsModel) CreatePlayerStats(ctx context.Context, client *ent.Client, data *ent.PlayerStats, p *ent.Player) (*ent.PlayerStats, error) {
//	//create the playerstats
//	psp := &PlayerStats{
//		Team: PSTeam{
//			ApiFootballId: data.Edges.Team.ID,
//			Name:          data.Edges.Team.Name,
//			Logo:          data.Edges.Team.Logo,
//		},
//		League: PSLeague{
//			ApiFootballId: data.Edges..ApiFootballId,
//			Name:          data.League.Name,
//			Country:       data.League.Country,
//			Logo:          data.League.Logo,
//			Flag:          data.League.Flag,
//			Season:        data.League.Season,
//		},
//		Games: PSGames{
//			Appearances: data.Games.Appearances,
//			Lineups:     data.Games.Lineups,
//			Minutes:     data.Games.Minutes,
//			Number:      data.Games.Number,
//			Position:    data.Games.Position,
//			Rating:      data.Games.Rating,
//			Captain:     data.Games.Captain,
//		},
//		Substitutes: PSSubstitutes{
//			In:    data.Substitutes.In,
//			Out:   data.Substitutes.Out,
//			Bench: data.Substitutes.Bench,
//		},
//		Shots: PSShots{
//			Total: data.Shots.Total,
//			On:    data.Shots.On,
//		},
//		Goals: PSGoals{
//			Total:    data.Goals.Total,
//			Conceded: data.Goals.Conceded,
//			Assists:  data.Goals.Assists,
//			Saves:    data.Goals.Saves,
//		},
//		Passes: PSPasses{
//			Total:    data.Passes.Total,
//			Key:      data.Passes.Key,
//			Accuracy: data.Passes.Accuracy,
//		},
//		Tackles: PSTackles{
//			Total:         data.Tackles.Total,
//			Blocks:        data.Tackles.Blocks,
//			Interceptions: data.Tackles.Interceptions,
//		},
//		Duels: PSDuels{
//			Total: data.Duels.Total,
//			Won:   data.Duels.Won,
//		},
//		Dribbles: PSDribbles{
//			Attempts: data.Dribbles.Attempts,
//			Success:  data.Dribbles.Success,
//			Past:     data.Dribbles.Past,
//		},
//		Fouls: PSFouls{
//			Drawn:     data.Fouls.Drawn,
//			Committed: data.Fouls.Committed,
//		},
//		Cards: PSCards{
//			Yellow:    data.Cards.Yellow,
//			YellowRed: data.Cards.YellowRed,
//			Red:       data.Cards.Red,
//		},
//		Penalty: PSPenalty{
//			Won:      data.Penalty.Won,
//			Commited: data.Penalty.Commited,
//			Scored:   data.Penalty.Scored,
//			Missed:   data.Penalty.Missed,
//			Saved:    data.Penalty.Saved,
//		},
//	}
//	return psp, nil
//}

//func (m *PlayerStatsModel) upsertPSPenalty(ctx context.Context, p *ent.PlayerStats, data PSPenalty) (*ent.PSPenalty, error) {
//	prev, err := m.client.PSPenalty.Query().Where(pspenalty.HasPlayerStatsWith(playerstats.IDEQ(p.ID))).First(ctx)
//	if err != nil && !ent.IsNotFound(err) {
//		return nil, fmt.Errorf("error fetching existing PSPenalty: %v", err)
//	}
//	if prev == nil {
//		return CreatePSPenalty(ctx, m.client, data, p)
//	} else {
//		return UpdatePSPenalty(ctx, m.client, data, prev)
//	}
//}
//
//func (m *PlayerStatsModel) upsertPSDefense(ctx context.Context, p *ent.PlayerStats, data PSDefense) (*ent.PSDefense, error) {
//	prev, err := m.client.PSDefense.Query().Where(psdefense.HasPlayerStatsWith(playerstats.IDEQ(p.ID))).First(ctx)
//	if err != nil && !ent.IsNotFound(err) {
//		return nil, fmt.Errorf("error fetching existing PSDefense: %v", err)
//	}
//	if prev == nil {
//		return CreatePSDefense(ctx, m.client, data, p)
//	} else {
//		return UpdatePSDefense(ctx, m.client, data, prev)
//	}
//}

//func (m *PlayerStatsModel) upsertPSOffense(ctx context.Context, p *ent.PlayerStats, data PSOffense) (*ent.PSOffense, error) {
//	prev, err := m.client.PSOffense.Query().Where(psoffense.HasPlayerStatsWith(playerstats.IDEQ(p.ID))).First(ctx)
//	if err != nil && !ent.IsNotFound(err) {
//		return nil, fmt.Errorf("error fetching existing PSOffense: %v", err)
//	}
//	if prev == nil {
//		return CreatePSOffense(ctx, m.client, data, p)
//	} else {
//		return UpdatePSOffense(ctx, m.client, data, prev)
//	}
//}

//func (m *PlayerStatsModel) upsertPSGoals(ctx context.Context, p *ent.PlayerStats, data PSGoals) (*ent.PSGoals, error) {
//	prev, err := m.client.PSGoals.Query().Where(psgoals.HasPlayerStatsWith(playerstats.IDEQ(p.ID))).First(ctx)
//	if err != nil && !ent.IsNotFound(err) {
//		return nil, fmt.Errorf("error fetching existing PSGoals: %v", err)
//	}
//	if prev == nil {
//		return CreatePSGoals(ctx, m.client, data, p)
//	} else {
//		return UpdatePSGoals(ctx, m.client, data, prev)
//	}
//}

//func (m *PlayerStatsModel) upsertPSGames(ctx context.Context, p *ent.PlayerStats, data PSGames) (*ent.PSGames, error) {
//	prev, err := m.client.PSGames.Query().Where(psgames.HasPlayerStatsWith(playerstats.IDEQ(p.ID))).First(ctx)
//	if err != nil && !ent.IsNotFound(err) {
//		return nil, fmt.Errorf("error fetching existing PSGames: %v", err)
//	}
//	if prev == nil {
//		return CreatePSGames(ctx, m.client, data, p)
//	} else {
//		return UpdatePSGames(ctx, m.client, data, prev)
//	}
//}

//func (m *PlayerStatsModel) ListPlayers(ctx context.Context) ([]*ent.Player, error) {
//	players, err := m.client.Player.
//		Query().
//		WithPlayerStats(func(query *ent.PlayerStatsQuery) {
//			query.WithPlayer().WithTeam(func(tq *ent.TeamQuery) {
//				tq.WithSeason()
//			})
//		}).
//		All(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("error listing players: %v", err)
//	}
//	return players, nil
//}
