package models

import (
	"context"
	"errors"
	_ "fmt"
	"mapeleven/db/ent"
	"mapeleven/db/ent/season"
	"mapeleven/db/ent/standings"
	"mapeleven/db/ent/team"
)

// CreateStandingsInput holds the required fields to create a new standings entry.
type CreateStandingsInput struct {
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

// UpdateStandingsInput holds the fields to update an existing standings entry.
type UpdateStandingsInput struct {
	Rank         *int
	Description  *string
	SeasonID     *int
	Team         *int
	Points       *int
	GoalsDiff    *int
	Group        *string
	Form         *string
	Status       *string
	Played       *int
	Win          *int
	Draw         *int
	Lose         *int
	GoalsFor     *int
	GoalsAgainst *int

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

// DeleteStandingsInput holds the required fields to delete a standings entry.
type DeleteStandingsInput struct {
	ID int
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
func (m *StandingsModel) CreateStandings(ctx context.Context, input CreateStandingsInput, team *ent.Team) (*ent.Standings, error) {

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

// DeleteStandings deletes a standings entry.
func (m *StandingsModel) DeleteStandings(ctx context.Context, input DeleteStandingsInput) error {
	err := m.client.Standings.
		DeleteOneID(input.ID).
		Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("standings entry not found")
		}
		return err
	}
	return nil
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

// CheckIfStandingsExist checks if a standings entry exists.
func (m *StandingsModel) CheckIfStandingsExist(ctx context.Context, teamID, seasonID int) bool {
	count, _ := m.client.Standings.
		Query().
		Where(
			standings.HasTeamWith(
				team.ID(teamID),
			),
			standings.HasSeasonWith(
				season.ID(seasonID),
			),
		).
		Count(ctx)

	return count > 0
}
