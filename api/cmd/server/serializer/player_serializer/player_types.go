package player_serializer

type APIPlayer struct {
	Slug       string           `json:"slug"`
	Name       string           `json:"name"`
	Firstname  string           `json:"firstname"`
	Lastname   string           `json:"lastname"`
	Age        int              `json:"age"`
	Height     string           `json:"height"`
	Weight     string           `json:"weight"`
	Injured    bool             `json:"injured"`
	Photo      string           `json:"photo"`
	Statistics []APIPlayerStats `json:"statistics,omitempty"`
	Popularity int              `json:"popularity,omitempty"`
}

type APIPlayerStats struct {
	Team        APIPSTeam        `json:"team"`
	League      APIPSLeague      `json:"league"`
	Games       APIPSGames       `json:"games"`
	Substitutes APIPSSubstitutes `json:"substitutes"`
	Shooting    APIPSShooting    `json:"shooting"`
	Technical   APIPSTechnical   `json:"technical"`
	Penalty     APIPSPenalty     `json:"penalty"`
	Defense     APIPSDefense     `json:"defense"`
	Fairplay    APIPSFairplay    `json:"fairplay"`
}

type APIPSTeam struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type APIPSLeague struct {
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
}

type APIPSDefense struct {
	TacklesTotal  int `json:"tackles_total"`
	Blocks        int `json:"blocks"`
	Interceptions int `json:"interceptions"`
	DuelsTotal    int `json:"duels_total"`
	DuelsWon      int `json:"duels_won"`
	DribblePast   int `json:"dribble_past"`
}

type APIPSGames struct {
	Appearances int    `json:"appearances"`
	Lineups     int    `json:"lineups"`
	Minutes     int    `json:"minutes"`
	Number      int    `json:"number"`
	Position    string `json:"position"`
	Rating      string `json:"rating"`
	Captain     bool   `json:"captain"`
}

type APIPSSubstitutes struct {
	In    int `json:"in"`
	Out   int `json:"out"`
	Bench int `json:"bench"`
}

type APIPSShooting struct {
	Goals    int `json:"goals"`
	Conceded int `json:"conceded"`
	Assists  int `json:"assists"`
	Saves    int `json:"saves"`
	Shots    int `json:"shots"`
	OnTarget int `json:"shots_on_target"`
}

type APIPSTechnical struct {
	FoulsDrawn      int `json:"fouls_drawn"`
	DribbleAttempts int `json:"dribble_attempts"`
	DribblesSuccess int `json:"dribbles_success"`
	Passes          int `json:"passes"`
	KeyPasses       int `json:"key_passes"`
	PassAccuracy    int `json:"pass_accuracy"`
}

type APIPSFairplay struct {
	Yellow            int `json:"yellow"`
	YellowRed         int `json:"yellow_red"`
	Red               int `json:"red"`
	FoulsCommitted    int `json:"fouls_committed"`
	PenaltiesConceded int `json:"penalties_conceded"`
}

type APIPSPenalty struct {
	Won       int `json:"won"`
	Scored    int `json:"scored"`
	Missed    int `json:"missed"`
	Saved     int `json:"saved"`
	Committed int `json:"committed"`
}
