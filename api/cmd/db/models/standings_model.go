package models

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/standings"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"context"
	"errors"
)

// StandingsForm holds the required fields to create a new standings entry.
type StandingsForm struct {
	Rank         int
	Description  string
	SeasonID     int
	Team         int
	Points       int
	GoalsDiff    int
	Group        string
	Form         string
	Status       string
	Played       int
	Win          int
	Draw         int
	Lose         int
	GoalsFor     int
	GoalsAgainst int

	// Home
	HomePlayed       int
	HomeWin          int
	HomeDraw         int
	HomeLose         int
	HomeGoalsFor     int
	HomeGoalsAgainst int

	// Away
	AwayPlayed       int
	AwayWin          int
	AwayDraw         int
	AwayLose         int
	AwayGoalsFor     int
	AwayGoalsAgainst int
}

// StandingsModel defines the methods for managing the standings data.
type StandingsModel struct {
	client *ent.Client
}

// NewStandingsModel creates a new StandingsModel.
func NewStandingsModel(client *ent.Client) *StandingsModel {
	return &StandingsModel{client: client}
}

// CreateStandings creates a new standings entry.
func (m *StandingsModel) CreateStandings(ctx context.Context, input StandingsForm, team *ent.Team) (*ent.Standings, error) {

	stands, err := m.client.Standings.
		Create().
		SetRank(input.Rank).
		SetDescription(input.Description).
		SetSeasonID(input.SeasonID).
		SetTeam(team).
		SetPoints(input.Points).
		SetGoalsDiff(input.GoalsDiff).
		SetGroup(input.Group).
		SetForm(input.Form).
		SetStatus(input.Status).

		//All
		SetAllPlayed(input.Played).
		SetAllWin(input.Win).
		SetAllDraw(input.Draw).
		SetAllLose(input.Lose).
		SetAllGoalsFor(input.GoalsFor).
		SetAllGoalsAgainst(input.GoalsAgainst).

		// Home
		SetHomePlayed(input.HomePlayed).
		SetHomeWin(input.HomeWin).
		SetHomeDraw(input.HomeDraw).
		SetHomeLose(input.HomeLose).
		SetHomeGoalsFor(input.HomeGoalsFor).
		SetHomeGoalsAgainst(input.HomeGoalsAgainst).

		// Away
		SetAwayPlayed(input.AwayPlayed).
		SetAwayWin(input.AwayWin).
		SetAwayDraw(input.AwayDraw).
		SetAwayLose(input.AwayLose).
		SetAwayGoalsFor(input.AwayGoalsFor).
		SetAwayGoalsAgainst(input.AwayGoalsAgainst).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return stands, nil
}

// UpdateStandings updates an existing standings entry.
func (m *StandingsModel) UpdateStandings(ctx context.Context, input StandingsForm, id int) (*ent.Standings, error) {
	stands, err := m.client.Standings.
		UpdateOneID(id).
		SetRank(input.Rank).
		SetDescription(input.Description).
		SetSeasonID(input.SeasonID).
		SetPoints(input.Points).
		SetGoalsDiff(input.GoalsDiff).
		SetGroup(input.Group).
		SetForm(input.Form).
		SetStatus(input.Status).

		//All
		SetAllPlayed(input.Played).
		SetAllWin(input.Win).
		SetAllDraw(input.Draw).
		SetAllLose(input.Lose).
		SetAllGoalsFor(input.GoalsFor).
		SetAllGoalsAgainst(input.GoalsAgainst).

		// Home
		SetHomePlayed(input.HomePlayed).
		SetHomeWin(input.HomeWin).
		SetHomeDraw(input.HomeDraw).
		SetHomeLose(input.HomeLose).
		SetHomeGoalsFor(input.HomeGoalsFor).
		SetHomeGoalsAgainst(input.HomeGoalsAgainst).

		// Away
		SetAwayPlayed(input.AwayPlayed).
		SetAwayWin(input.AwayWin).
		SetAwayDraw(input.AwayDraw).
		SetAwayLose(input.AwayLose).
		SetAwayGoalsFor(input.AwayGoalsFor).
		SetAwayGoalsAgainst(input.AwayGoalsAgainst).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return stands, nil
}

// GetStandingsByID retrieves a standings entry by FootballApiId.
func (m *StandingsModel) GetStandingsByID(ctx context.Context, id int) (*ent.Standings, error) {
	stands, err := m.client.Standings.
		Query().
		Where(standings.ID(id)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("stands entry not found")
		}
		return nil, err
	}
	return stands, nil
}

// ListStandings retrieves a list of all standings entries.
func (m *StandingsModel) ListStandings(ctx context.Context) ([]*ent.Standings, error) {
	standingsEntries, err := m.client.Standings.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return standingsEntries, nil
}

// Exists checks if a standings entry exists.
func (m *StandingsModel) Exists(ctx context.Context, teamID, seasonID int) (bool, *ent.Standings) {
	standing, _ := m.client.Standings.
		Query().
		Where(
			standings.HasTeamWith(
				team.HasClubWith(club.ApiFootballIdEQ(teamID)),
			),
			standings.HasSeasonWith(
				season.ID(seasonID),
			),
		).
		Only(ctx)

	if standing == nil {
		return false, nil
	} else {
		return true, standing
	}
}
