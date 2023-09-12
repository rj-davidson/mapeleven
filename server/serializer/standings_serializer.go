package serializer

import (
	"mapeleven/db/ent"
)

type APIRecord struct {
	Played       int `json:"played"`
	Won          int `json:"won"`
	Draw         int `json:"draw"`
	Lost         int `json:"lost"`
	GoalsFor     int `json:"goalsFor"`
	GoalsAgainst int `json:"goalsAgainst"`
}

type APIStanding struct {
	Rank        int       `json:"rank"`
	Description string    `json:"description"`
	League      APILeague `json:"league"`
	Team        APITeam   `json:"team"`

	Points    int    `json:"points"`
	GoalsDiff int    `json:"goalsDiff"`
	Group     string `json:"group"`
	Form      string `json:"form"`
	Status    string `json:"status"`

	Home    APIRecord `json:"home"`
	Away    APIRecord `json:"away"`
	Overall APIRecord `json:"overall"`
}

func SerializeStanding(standings *ent.Standings) *APIStanding {
	return &APIStanding{
		Rank:        standings.Rank,
		Description: standings.Description,
		League: APILeague{
			Slug: standings.Edges.League.Slug,
			Name: standings.Edges.League.Name,
			Type: standings.Edges.League.Type.String(),
			Logo: standings.Edges.League.Logo,
		},
		Team: APITeam{
			Slug: standings.Edges.Team.Slug,
			Name: APITeamName{
				Long:  standings.Edges.Team.Name,
				Short: standings.Edges.Team.Code,
			},
			Badge: standings.Edges.Team.Logo,
		},
		Points:    standings.Points,
		GoalsDiff: standings.GoalsDiff,
		Group:     standings.Group,
		Form:      standings.Form,
		Status:    standings.Status,
		Home: APIRecord{
			Played:       standings.HomePlayed,
			Won:          standings.HomeWin,
			Draw:         standings.HomeDraw,
			Lost:         standings.HomeLose,
			GoalsFor:     standings.HomeGoalsFor,
			GoalsAgainst: standings.HomeGoalsAgainst,
		},
		Away: APIRecord{
			Played:       standings.AwayPlayed,
			Won:          standings.AwayWin,
			Draw:         standings.AwayDraw,
			Lost:         standings.AwayLose,
			GoalsFor:     standings.AwayGoalsFor,
			GoalsAgainst: standings.AwayGoalsAgainst,
		},
		Overall: APIRecord{
			Played:       standings.AllPlayed,
			Won:          standings.AllWin,
			Draw:         standings.AllDraw,
			Lost:         standings.AllLose,
			GoalsFor:     standings.AllGoalsFor,
			GoalsAgainst: standings.AllGoalsAgainst,
		},
	}
}
