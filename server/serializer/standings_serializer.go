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
	Rank        int     `json:"rank"`
	Description string  `json:"description"`
	Team        APITeam `json:"team"`

	Points    int    `json:"points"`
	GoalsDiff int    `json:"goalsDiff"`
	Group     string `json:"group"`
	Form      string `json:"form"`
	Status    string `json:"status"`

	Home    APIRecord `json:"home"`
	Away    APIRecord `json:"away"`
	Overall APIRecord `json:"overall"`
}

func SerializeStanding(standing *ent.Standings) *APIStanding {
	return &APIStanding{
		Rank:        standing.Rank,
		Description: standing.Description,
		Team: APITeam{
			Slug: standing.Edges.Team.Edges.Club.Slug,
			Name: APITeamName{
				Long:  standing.Edges.Team.Edges.Club.Name,
				Short: standing.Edges.Team.Edges.Club.Code,
			},
			Badge: standing.Edges.Team.Edges.Club.Logo,
		},
		Points:    standing.Points,
		GoalsDiff: standing.GoalsDiff,
		Group:     standing.Group,
		Form:      standing.Form,
		Status:    standing.Status,
		Home: APIRecord{
			Played:       standing.HomePlayed,
			Won:          standing.HomeWin,
			Draw:         standing.HomeDraw,
			Lost:         standing.HomeLose,
			GoalsFor:     standing.HomeGoalsFor,
			GoalsAgainst: standing.HomeGoalsAgainst,
		},
		Away: APIRecord{
			Played:       standing.AwayPlayed,
			Won:          standing.AwayWin,
			Draw:         standing.AwayDraw,
			Lost:         standing.AwayLose,
			GoalsFor:     standing.AwayGoalsFor,
			GoalsAgainst: standing.AwayGoalsAgainst,
		},
		Overall: APIRecord{
			Played:       standing.AllPlayed,
			Won:          standing.AllWin,
			Draw:         standing.AllDraw,
			Lost:         standing.AllLose,
			GoalsFor:     standing.AllGoalsFor,
			GoalsAgainst: standing.AllGoalsAgainst,
		},
	}
}
