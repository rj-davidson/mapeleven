package team_stats

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tsbiggest"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tscards"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tscleansheet"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tsfailedtoscore"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tsfixtures"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tsgoals"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tslineups"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/tspenalty"
	"context"
)

type TeamStats struct {
	Form          string        `json:"form"`
	Fixtures      Fixtures      `json:"fixtures"`
	Goals         Goals         `json:"goals"`
	Biggest       Biggest       `json:"biggest"`
	CleanSheet    CleanSheet    `json:"clean_sheet"`
	FailedToScore FailedToScore `json:"failed_to_score"`
	Penalty       Penalty       `json:"penalty"`
	Lineups       []Lineup      `json:"lineups"`
	Cards         Cards         `json:"cards"`
}

type Fixtures struct {
	Played HomeAwayTotalInt `json:"played"`
	Wins   HomeAwayTotalInt `json:"wins"`
	Draws  HomeAwayTotalInt `json:"draws"`
	Losses HomeAwayTotalInt `json:"loses"`
}

type HomeAwayTotalInt struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total,omitempty"`
}

type HomeAwayTotalString struct {
	Home  string `json:"home"`
	Away  string `json:"away"`
	Total string `json:"total,omitempty"`
}

type BiggestGoals struct {
	For     HomeAwayTotalInt `json:"for"`
	Against HomeAwayTotalInt `json:"against"`
}

type Goals struct {
	For     ScoreDetail `json:"for"`
	Against ScoreDetail `json:"against"`
}

type ScoreDetail struct {
	Total   HomeAwayTotalInt    `json:"total"`
	Average HomeAwayTotalString `json:"average"`
	Minute  MinuteDetail        `json:"minute"`
}

type AverageStats struct {
	Home  string `json:"home"`
	Away  string `json:"away"`
	Total string `json:"total"`
}

type MinuteDetail struct {
	Min0to15    MinuteStats `json:"0-15"`
	Min16to30   MinuteStats `json:"16-30"`
	Min31to45   MinuteStats `json:"31-45"`
	Min46to60   MinuteStats `json:"46-60"`
	Min61to75   MinuteStats `json:"61-75"`
	Min76to90   MinuteStats `json:"76-90"`
	Min91to105  MinuteStats `json:"91-105"`
	Min106to120 MinuteStats `json:"106-120"`
}

type MinuteStats struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}

type Biggest struct {
	Streak Streak       `json:"streak"`
	Win    Win          `json:"wins"`
	Lose   Lose         `json:"loses"`
	Goals  BiggestGoals `json:"goals"`
}

type Streak struct {
	Wins   int `json:"wins"`
	Draws  int `json:"draws"`
	Losses int `json:"loses"`
}

type Win struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type Lose struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type CleanSheet struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type FailedToScore struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type Penalty struct {
	Scored GoalFixedPercentage `json:"scored"`
	Missed GoalFixedPercentage `json:"missed"`
	Total  int                 `json:"total"`
}

type GoalFixedPercentage struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}

type Lineup struct {
	Formation string `json:"formation"`
	Played    int    `json:"played"`
}

type Cards struct {
	Yellow MinuteDetail `json:"yellow"`
	Red    MinuteDetail `json:"red"`
}

type TeamStatsModel struct {
	client *ent.Client
}

func NewTeamStatsModel(client *ent.Client) *TeamStatsModel {
	return &TeamStatsModel{client: client}
}

type TeamStatsPackage struct {
	Biggest       *ent.TSBiggest
	Cards         *ent.TSCards
	CleanSheets   *ent.TSCleanSheet
	FailedToScore *ent.TSFailedToScore
	Fixture       *ent.TSFixtures
	Goals         *ent.TSGoals
	Lineups       []*ent.TSLineups
	Penalty       *ent.TSPenalty
}

func (m *TeamStatsModel) UpsertTeamStats(ctx context.Context, t *ent.Team, teamStats TeamStats) (*TeamStatsPackage, error) {
	// Add Form To Team Object
	_, err := m.client.Team.
		UpdateOne(t).
		SetForm(teamStats.Form).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var teamStatsPackage TeamStatsPackage
	teamStatsPackage.Biggest, err = m.upsertBiggest(ctx, t, teamStats.Biggest)
	if err != nil {
		return &teamStatsPackage, err
	}
	teamStatsPackage.Cards, err = m.upsertCards(ctx, t, teamStats.Cards)
	if err != nil {
		return &teamStatsPackage, err
	}
	teamStatsPackage.CleanSheets, err = m.upsertCleanSheets(ctx, t, teamStats.CleanSheet)
	if err != nil {
		return &teamStatsPackage, err
	}
	teamStatsPackage.FailedToScore, err = m.upsertFailedToScore(ctx, t, teamStats.FailedToScore)
	if err != nil {
		return &teamStatsPackage, err
	}
	teamStatsPackage.Fixture, err = m.upsertFixtures(ctx, t, teamStats.Fixtures)
	if err != nil {
		return &teamStatsPackage, err
	}
	teamStatsPackage.Goals, err = m.upsertGoals(ctx, t, teamStats.Goals)
	if err != nil {
		return &teamStatsPackage, err
	}
	teamStatsPackage.Lineups, err = m.upsertLineups(ctx, t, teamStats.Lineups)
	if err != nil {
		return &teamStatsPackage, err
	}
	teamStatsPackage.Penalty, err = m.upsertPenalty(ctx, t, teamStats.Penalty)
	if err != nil {
		return &teamStatsPackage, err
	}

	return &teamStatsPackage, nil
}

func (m *TeamStatsModel) upsertBiggest(ctx context.Context, t *ent.Team, data Biggest) (*ent.TSBiggest, error) {
	prev, _ := m.client.TSBiggest.Query().Where(tsbiggest.HasTeamWith(team.IDEQ(t.ID))).Only(ctx)
	if prev == nil {
		return CreateTSBiggest(ctx, m.client, data, t)
	} else {
		return UpdateTSBiggest(ctx, m.client, data, prev)
	}
}

func (m *TeamStatsModel) upsertCards(ctx context.Context, t *ent.Team, data Cards) (*ent.TSCards, error) {
	prev, _ := m.client.TSCards.Query().Where(tscards.HasTeamWith(team.IDEQ(t.ID))).Only(ctx)
	if prev == nil {
		return CreateTSCards(ctx, m.client, data, t)
	} else {
		return UpdateTSCards(ctx, m.client, data, prev)
	}
}

func (m *TeamStatsModel) upsertCleanSheets(ctx context.Context, t *ent.Team, data CleanSheet) (*ent.TSCleanSheet, error) {
	prev, _ := m.client.TSCleanSheet.Query().Where(tscleansheet.HasTeamWith(team.IDEQ(t.ID))).Only(ctx)
	if prev == nil {
		return CreateTSCleanSheet(ctx, m.client, data, t)
	} else {
		return UpdateTSCleanSheet(ctx, m.client, data, prev)
	}
}

func (m *TeamStatsModel) upsertFailedToScore(ctx context.Context, t *ent.Team, data FailedToScore) (*ent.TSFailedToScore, error) {
	prev, _ := m.client.TSFailedToScore.Query().Where(tsfailedtoscore.HasTeamWith(team.IDEQ(t.ID))).Only(ctx)
	if prev == nil {
		return CreateTSFailedToScore(ctx, m.client, data, t)
	} else {
		return UpdateTSFailedToScore(ctx, m.client, data, prev)
	}
}

func (m *TeamStatsModel) upsertFixtures(ctx context.Context, t *ent.Team, data Fixtures) (*ent.TSFixtures, error) {
	prev, _ := m.client.TSFixtures.Query().Where(tsfixtures.HasTeamWith(team.IDEQ(t.ID))).Only(ctx)
	if prev == nil {
		return CreateTSFixtures(ctx, m.client, data, t)
	} else {
		return UpdateTSFixtures(ctx, m.client, data, prev)
	}
}

func (m *TeamStatsModel) upsertGoals(ctx context.Context, t *ent.Team, data Goals) (*ent.TSGoals, error) {
	prev, _ := m.client.TSGoals.Query().Where(tsgoals.HasTeamWith(team.IDEQ(t.ID))).Only(ctx)
	if prev == nil {
		return CreateTSGoals(ctx, m.client, data, t)
	} else {
		return UpdateTSGoals(ctx, m.client, data, prev)
	}
}

func (m *TeamStatsModel) upsertLineups(ctx context.Context, t *ent.Team, data []Lineup) ([]*ent.TSLineups, error) {
	var lineups []*ent.TSLineups
	for _, lineup := range data {
		prev, _ := m.client.TSLineups.Query().Where(tslineups.HasTeamWith(team.IDEQ(t.ID))).Where(tslineups.FormationEQ(lineup.Formation)).Only(ctx)
		if prev == nil {
			l, err := CreateTSLineup(ctx, m.client, lineup, t)
			if err != nil {
				return nil, err
			}
			lineups = append(lineups, l)
		} else {
			l, err := UpdateTSLineup(ctx, m.client, lineup, prev)
			if err != nil {
				return nil, err
			}
			lineups = append(lineups, l)
		}
	}
	return lineups, nil
}

func (m *TeamStatsModel) upsertPenalty(ctx context.Context, t *ent.Team, data Penalty) (*ent.TSPenalty, error) {
	prev, _ := m.client.TSPenalty.Query().Where(tspenalty.HasTeamWith(team.IDEQ(t.ID))).Only(ctx)
	if prev == nil {
		return CreateTSPenalty(ctx, m.client, data, t)
	} else {
		return UpdateTSPenalty(ctx, m.client, data, prev)
	}
}
