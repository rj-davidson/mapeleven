package fixture_serializer

import (
	"time"
)

type APIFixture struct {
	Slug          string    `json:"slug"`
	Referee       string    `json:"referee"`
	Timezone      string    `json:"timezone"`
	Date          time.Time `json:"date"`
	Elapsed       int       `json:"elapsed"`
	Round         int       `json:"round"`
	Status        string    `json:"status"`
	HomeTeamScore int       `json:"homeTeamScore"`
	AwayTeamScore int       `json:"awayTeamScore"`
	Season        APISeason `json:"season"`
}

type APISeason struct {
	Year    int       `json:"year"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Current bool      `json:"current"`
}

type APITeam struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type APIFixtureSet struct {
	Fixture *APIFixture     `json:"fixture,omitempty"`
	Teams   *APITeamDetails `json:"teams,omitempty"`
	Events  *[]APIEvent     `json:"events,omitempty"`
	Lineups *APIHomeAway    `json:"lineups,omitempty"`
}

type APIHomeAway struct {
	Home *APILineupInfo `json:"home,omitempty"`
	Away *APILineupInfo `json:"away,omitempty"`
}

type APITeamDetails struct {
	Home APITeamInfo `json:"home"`
	Away APITeamInfo `json:"away"`
}

type APITeamInfo struct {
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner *bool  `json:"winner,omitempty"`
}

type APIEvent struct {
	Time     APITimeDetail `json:"time"`
	Team     APITeamInfo   `json:"team"`
	Player   APIPlayer     `json:"player"`
	Assist   *APIPlayer    `json:"assist"`
	Type     string        `json:"type"`
	Detail   string        `json:"detail"`
	Comments *string       `json:"comments"`
}

type APITimeDetail struct {
	Elapsed int  `json:"elapsed"`
	Extra   *int `json:"extra"`
}

type APIPlayer struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type APILineupInfo struct {
	Team         APITeamInfo       `json:"team"`
	Coach        APICoach          `json:"coach"`
	Formation    string            `json:"formation"`
	MatchPlayers []APIPlayerDetail `json:"matchPlayers"`
}

type APICoach struct {
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type APIPlayerDetail struct {
	Player APIPlayerDetailInfo `json:"player"`
}

type APIPlayerDetailInfo struct {
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Number int    `json:"number"`
	Pos    string `json:"pos"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}
