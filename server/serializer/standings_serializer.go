package serializer

import (
	"mapeleven/db/ent"
)

type Record struct {
	Played       int `json:"played"`
	Won          int `json:"won"`
	Draw         int `json:"draw"`
	Lost         int `json:"lost"`
	GoalsFor     int `json:"goalsFor"`
	GoalsAgainst int `json:"goalsAgainst"`
}

type Standing struct {
	Rank        int            `json:"rank"`
	Description string         `json:"description"`
	League      LeagueItem     `json:"league"`
	Team        TeamSerializer `json:"team"`

	Points    int    `json:"points"`
	GoalsDiff int    `json:"goalsDiff"`
	Group     string `json:"group"`
	Form      string `json:"form"`
	Status    string `json:"status"`

	Home    Record `json:"home"`
	Away    Record `json:"away"`
	Overall Record `json:"overall"`
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
		Points:    standings.Points,
		GoalsDiff: standings.GoalsDiff,
		Group:     standings.Group,
		Form:      standings.Form,
		Status:    standings.Status,
		Home: Record{
			Played:       standings.HomePlayed,
			Won:          standings.HomeWin,
			Draw:         standings.HomeDraw,
			Lost:         standings.HomeLose,
			GoalsFor:     standings.HomeGoalsFor,
			GoalsAgainst: standings.HomeGoalsAgainst,
		},
		Away: Record{
			Played:       standings.AwayPlayed,
			Won:          standings.AwayWin,
			Draw:         standings.AwayDraw,
			Lost:         standings.AwayLose,
			GoalsFor:     standings.AwayGoalsFor,
			GoalsAgainst: standings.AwayGoalsAgainst,
		},
		Overall: Record{
			Played:       standings.AllPlayed,
			Won:          standings.AllWin,
			Draw:         standings.AllDraw,
			Lost:         standings.AllLose,
			GoalsFor:     standings.AllGoalsFor,
			GoalsAgainst: standings.AllGoalsAgainst,
		},
	}
}
