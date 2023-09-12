package serializer

import (
	"mapeleven/db/ent"
)

type Standing struct {
	Rank        int            `json:"rank"`
	Description string         `json:"description"`
	League      LeagueItem     `json:"league"`
	Team        TeamSerializer `json:"team"`

	Points       int    `json:"points"`
	GoalsDiff    int    `json:"goalsDiff"`
	Group        string `json:"group"`
	Form         string `json:"form"`
	Status       string `json:"status"`
	Played       int    `json:"played"`
	Win          int    `json:"win"`
	Draw         int    `json:"draw"`
	Lose         int    `json:"lose"`
	GoalsFor     int    `json:"goalsFor"`
	GoalsAgainst int    `json:"goalsAgainst"`

	// Home
	HomePlayed       int `json:"homePlayed"`
	HomeWin          int `json:"homeWin"`
	HomeDraw         int `json:"homeDraw"`
	HomeLose         int `json:"homeLose"`
	HomeGoalsFor     int `json:"homeGoalsFor"`
	HomeGoalsAgainst int `json:"homeGoalsAgainst"`

	// Away
	AwayPlayed       int `json:"awayPlayed"`
	AwayWin          int `json:"awayWin"`
	AwayDraw         int `json:"awayDraw"`
	AwayLose         int `json:"awayLose"`
	AwayGoalsFor     int `json:"awayGoalsFor"`
	AwayGoalsAgainst int `json:"awayGoalsAgainst"`
}

func SerializeStanding(standings *ent.Standings) *Standing {
	return &Standing{
		Rank:        standings.Rank,
		Description: standings.Description,
		League: LeagueItem{
			Slug: standings.Edges.League.Slug,
			Name: standings.Edges.League.Name,
			Type: standings.Edges.League.Type.String(),
			Logo: standings.Edges.League.Logo,
		},
		Team: TeamSerializer{
			Slug:  standings.Edges.Team.Slug,
			Name:  standings.Edges.Team.Name,
			Badge: standings.Edges.Team.Logo,
		},
		Points:           standings.Points,
		GoalsDiff:        standings.GoalsDiff,
		Group:            standings.Group,
		Form:             standings.Form,
		Status:           standings.Status,
		Played:           standings.AllPlayed,
		Win:              standings.AllWin,
		Draw:             standings.AllDraw,
		Lose:             standings.AllLose,
		GoalsFor:         standings.AllGoalsFor,
		GoalsAgainst:     standings.AllGoalsAgainst,
		HomePlayed:       standings.HomePlayed,
		HomeWin:          standings.HomeWin,
		HomeDraw:         standings.HomeDraw,
		HomeLose:         standings.HomeLose,
		HomeGoalsFor:     standings.HomeGoalsFor,
		HomeGoalsAgainst: standings.HomeGoalsAgainst,
		AwayPlayed:       standings.AwayPlayed,
		AwayWin:          standings.AwayWin,
		AwayDraw:         standings.AwayDraw,
		AwayLose:         standings.AwayLose,
		AwayGoalsFor:     standings.AwayGoalsFor,
		AwayGoalsAgainst: standings.AwayGoalsAgainst,
	}
}
