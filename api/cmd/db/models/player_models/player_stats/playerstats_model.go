package player_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgoals"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgoals"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psoffense"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psoffense"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	_ "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	"context"
	_ "context"
)

type PlayerStats struct {
	Form      string    `json:"form"`
	PSPenalty PSPenalty `json:"penalty"`
	PSDefense PSDefense `json:"defense"`
	PSGoals   PSGoals   `json:"goals"`
	PSOffense PSOffense `json:"offense"`
	PSGames   PSGames   `json:"games"`
}
type PSPenalty struct {
	FoulsDrawn      int `json:"fouls_drawn"`
	FoulsCommitted  int `json:"fouls_committed"`
	CardsYellow     int `json:"cards_yellow"`
	CardYellowred   int `json:"card_yellowred"`
	CardsRed        int `json:"cards_red"`
	PenaltyWon      int `json:"penalty_won"`
	PenaltyCommited int `json:"penalty_commited"`
	PenaltyScored   int `json:"penalty_scored"`
	PenaltyMissed   int `json:"penalty_missed"`
	PenaltySaved    int `json:"penalty_saved"`
}
type PSDefense struct {
	TacklesTotal  int `json:"tackles_total"`
	Blocks        int `json:"blocks"`
	Interceptions int `json:"interceptions"`
	TotalDuels    int `json:"total_duels"`
	WonDuels      int `json:"won_duels"`
}

type PSGoals struct {
	TotalGoals    int `json:"total_goals"`
	ConcededGoals int `json:"conceded_goals"`
	AssistGoals   int `json:"assist_goals"`
	SaveGoals     int `json:"save_goals"`
	ShotsTotal    int `json:"shots_total"`
	ShotsOn       int `json:"shots_on"`
}

type PSOffense struct {
	DribbleAttempts int `json:"dribbles_attempts"`
	DribbleSuccess  int `json:"dribbles_success"`
	DribblePast     int `json:"dribbled_past"`
	PassesTotal     int `json:"passes_total"`
	PassesKey       int `json:"passes_key"`
	PassesAccuracy  int `json:"passes_accuracy"`
}

type PlayerStatsModel struct {
	client *ent.Client
}

func NewPlayerStatsModel(client *ent.Client) *PlayerStatsModel {
	return &PlayerStatsModel{client: client}
}

type PlayerStatsPackage struct {
	PSPenalty *ent.PSPenalty
	PSDefense *ent.PSDefense
	PSGoals   *ent.PSGoals
	PSOffense *ent.PSOffense
	PSGames   *ent.PSGames
}

func (m *PlayerStatsModel) UpsertPlayerStats(ctx context.Context, p *ent.Player, playerStats PlayerStats) (*PlayerStatsPackage, error) {
	//Add form to Player object
	_, err := m.client.Player.
		UpdateOne(p).SetForm(playerStats.Form).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var playerStatsPackage PlayerStatsPackage
	playerStatsPackage.PSPenalty, err = m.upsertPSPenalty(ctx, p, playerStats.PSPenalty)
	if err != nil {
		return &playerStatsPackage, err
	}
	playerStatsPackage.PSDefense, err = m.upsertPSDefense(ctx, p, playerStats.PSDefense)
	if err != nil {
		return &playerStatsPackage, err
	}
	playerStatsPackage.PSGoals, err = m.upsertPSGoals(ctx, p, playerStats.PSGoals)
	if err != nil {
		return &playerStatsPackage, err
	}
	playerStatsPackage.PSOffense, err = m.upsertPSOffense(ctx, p, playerStats.PSOffense)
	if err != nil {
		return &playerStatsPackage, err
	}
	playerStatsPackage.PSGames, err = m.upsertPSGames(ctx, p, playerStats.PSGames)
	if err != nil {
		return &playerStatsPackage, err
	}
	return nil, nil
}

func (m *PlayerStatsModel) upsertPSPenalty(ctx context.Context, p *ent.Player, data PSPenalty) (*ent.PSPenalty, error) {
	prev, _ := m.client.PSPenalty.Query().Where(pspenalty.HasPlayerWith(player.IDEQ(p.ID))).First(ctx)
	if prev == nil {
		return CreatePSPenalty(ctx, m.client, data, p)
	} else {
		return UpdatePSPenalty(ctx, m.client, data, prev)
	}
}

func (m *PlayerStatsModel) upsertPSDefense(ctx context.Context, p *ent.Player, data PSDefense) (*ent.PSDefense, error) {
	prev, _ := m.client.PSDefense.Query().Where(psdefense.HasPlayerWith(player.IDEQ(p.ID))).First(ctx)
	if prev == nil {
		return CreatePSDefense(ctx, m.client, data, p)
	} else {
		return UpdatePSDefense(ctx, m.client, data, prev)
	}
}

func (m *PlayerStatsModel) upsertPSOffense(ctx context.Context, p *ent.Player, data PSOffense) (*ent.PSOffense, error) {
	prev, _ := m.client.PSOffense.Query().Where(psoffense.HasPlayerWith(player.IDEQ(p.ID))).First(ctx)
	if prev == nil {
		return CreatePSOffense(ctx, m.client, data, p)
	} else {
		return UpdatePSOffense(ctx, m.client, data, prev)
	}
}

func (m *PlayerStatsModel) upsertPSGoals(ctx context.Context, p *ent.Player, data PSGoals) (*ent.PSGoals, error) {
	// Use the correct predicate for PSGoals, not PSDefense
	prev, _ := m.client.PSGoals.Query().Where(psgoals.HasPlayerWith(player.IDEQ(p.ID))).First(ctx)
	if prev == nil {
		return CreatePSGoals(ctx, m.client, data, p)
	} else {
		return UpdatePSGoals(ctx, m.client, data, prev)
	}
}

func (m *PlayerStatsModel) upsertPSGames(ctx context.Context, p *ent.Player, data PSGames) (*ent.PSGames, error) {
	// Use the correct predicate for PSGames, not PSDefense
	prev, _ := m.client.PSGames.Query().Where(psgames.HasPlayerWith(player.IDEQ(p.ID))).First(ctx)
	if prev == nil {
		return CreatePSGames(ctx, m.client, data, p)
	} else {
		return UpdatePSGames(ctx, m.client, data, prev)
	}
}

func (m *PlayerStatsModel) ListPlayers(ctx context.Context) ([]*ent.Player, error) {
	players, err := m.client.Player.
		Query().
		WithSeason(
			func(q *ent.SeasonQuery) {
				q.WithLeague()
			}).
		WithTeam(
			func(q *ent.TeamQuery) {
				q.WithClub()
			}).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return players, nil
}
