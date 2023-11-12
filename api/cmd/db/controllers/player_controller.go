package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/player_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/utils"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type PSTeam struct {
	ApiFootballId int    `json:"id"`
	Name          string `json:"name"`
	Logo          string `json:"logo"`
}

type PSLeague struct {
	ApiFootballId int    `json:"id"`
	Name          string `json:"name"`
	Country       string `json:"country"`
	Logo          string `json:"logo"`
	Flag          string `json:"flag"`
	Season        int    `json:"season"`
}

type PSGames struct {
	Appearances int    `json:"appearances"`
	Lineups     int    `json:"lineups"`
	Minutes     int    `json:"minutes"`
	Number      int    `json:"number"`
	Position    string `json:"position"`
	Rating      string `json:"rating"`
	Captain     bool   `json:"captain"`
}

type PSSubstitutes struct {
	In    int `json:"in"`
	Out   int `json:"out"`
	Bench int `json:"bench"`
}

type PSShots struct {
	Total int `json:"total"`
	On    int `json:"on"`
}

type PSGoals struct {
	Total    int `json:"total"`
	Conceded int `json:"conceded"`
	Assists  int `json:"assists"`
	Saves    int `json:"saves"`
}

type PSPasses struct {
	Total    int `json:"total"`
	Key      int `json:"key"`
	Accuracy int `json:"accuracy"`
}

type PSTackles struct {
	Total         int `json:"total"`
	Blocks        int `json:"blocks"`
	Interceptions int `json:"interceptions"`
}

type PSDuels struct {
	Total int `json:"total"`
	Won   int `json:"won"`
}

type PSDribbles struct {
	Attempts int `json:"attempts"`
	Success  int `json:"success"`
	Past     int `json:"past"`
}

type PSFouls struct {
	Drawn     int `json:"drawn"`
	Committed int `json:"committed"`
}

type PSCards struct {
	Yellow    int `json:"yellow"`
	YellowRed int `json:"yellowred"`
	Red       int `json:"red"`
}

type PSPenalty struct {
	Won       int `json:"won"`
	Committed int `json:"commited"`
	Scored    int `json:"scored"`
	Missed    int `json:"missed"`
	Saved     int `json:"saved"`
}

type PlayerStats struct {
	Team        PSTeam        `json:"team"`
	League      PSLeague      `json:"league"`
	Games       PSGames       `json:"games"`
	Substitutes PSSubstitutes `json:"substitutes"`
	Shots       PSShots       `json:"shots"`
	Goals       PSGoals       `json:"goals"`
	Passes      PSPasses      `json:"passes"`
	Tackles     PSTackles     `json:"tackles"`
	Duels       PSDuels       `json:"duels"`
	Dribbles    PSDribbles    `json:"dribbles"`
	Fouls       PSFouls       `json:"fouls"`
	Cards       PSCards       `json:"cards"`
	Penalty     PSPenalty     `json:"penalty"`
	Position    string
}

type LeaguePlayerInfo struct {
	ID int `json:"id"`
}

type LeaguePlayer struct {
	Player LeaguePlayerInfo `json:"player"`
}

type PlayerInfo struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Firstname  string      `json:"firstname"`
	Lastname   string      `json:"lastname"`
	Age        int         `json:"age"`
	Height     string      `json:"height"`
	Weight     string      `json:"weight"`
	Injured    bool        `json:"injured"`
	Photo      string      `json:"photo"`
	Statistics PlayerStats `json:"statistics"`
}
type LeaguePlayerResponse struct {
	Results  int `json:"results"`
	Response []struct {
		Player struct {
			ID int `json:"id"`
		} `json:"player"`
	} `json:"response"`
}

type Player struct {
	Player PlayerInfo `json:"player"`
}

type PlayerResponse struct {
	Results  int      `json:"results"`
	Response []Player `json:"response"`
}

type PlayerController struct {
	client      *http.Client
	playerModel *player_models.PlayerModel
}

type Statistics struct {
	Games struct {
		Appearances int `json:"appearances"`
		Lineups     int `json:"lineups"`
		Minutes     int `json:"minutes"`
	} `json:"games"`
	Shots struct {
		Total int `json:"total"`
		On    int `json:"on"`
	} `json:"shots"`
	Goals struct {
		Total    int `json:"total"`
		Conceded int `json:"conceded"`
		Assists  int `json:"assists"`
		Saves    int `json:"saves"`
	} `json:"goals"`
	Passes struct {
		Total    int `json:"total"`
		Key      int `json:"key"`
		Accuracy int `json:"accuracy"`
	} `json:"passes"`
	Tackles struct {
		Total         int `json:"total"`
		Blocks        int `json:"blocks"`
		Interceptions int `json:"interceptions"`
	} `json:"tackles"`
	Duels struct {
		Total int `json:"total"`
		Won   int `json:"won"`
	} `json:"duels"`
	Dribbles struct {
		Attempts int `json:"attempts"`
		Success  int `json:"success"`
		Past     int `json:"past"`
	} `json:"dribbles"`
	Fouls struct {
		Drawn     int `json:"drawn"`
		Committed int `json:"committed"`
	} `json:"fouls"`
	Cards struct {
		Yellow    int `json:"yellow"`
		YellowRed int `json:"yellowred"`
		Red       int `json:"red"`
	} `json:"cards"`
	Penalty struct {
		Won      int `json:"won"`
		Commited int `json:"commited"`
		Scored   int `json:"scored"`
		Missed   int `json:"missed"`
		Saved    int `json:"saved"`
	} `json:"penalty"`
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	League struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country"`
		Logo    string `json:"logo"`
		Flag    string `json:"flag"`
	} `json:"league"`
	Position   string `json:"position"`
	Rating     string `json:"rating"`
	TeamName   string `json:"team_name"`
	TeamID     int    `json:"team_id"`
	LeagueID   int    `json:"league_id"`
	LeagueName string `json:"league_name"`
}

// PlayerStatisticsResponse is the struct to hold the API response for player statistics
type PlayerStatisticsResponse struct {
	Results  int `json:"results"`
	Response []struct {
		Player     PlayerInfo   `json:"player"`
		Statistics []Statistics `json:"statistics"`
	} `json:"response"`
}

func NewPlayerController(playerModel *player_models.PlayerModel) *PlayerController {
	return &PlayerController{
		client:      &http.Client{},
		playerModel: playerModel,
	}
}

func (pc *PlayerController) EnsurePlayerExists(ctx context.Context, apiFootballId int) error {
	exists := pc.playerModel.Exists(ctx, apiFootballId)
	if !exists {
		err1 := pc.fetchPlayerByID(ctx, apiFootballId)
		err2 := pc.fetchPlayerStatistics(ctx, apiFootballId, 2022) // Fetch stats for the year 2022
		if err1 != nil || err2 != nil {
			return fmt.Errorf("Error fetching player or statistics: %v %v", err1, err2)
		}
	}
	return nil
}

func (pc *PlayerController) fetchPlayerByID(ctx context.Context, playerID int) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?id=%d&season=2022", playerID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	fmt.Printf("Fetching player by ID: %d\n", playerID)
	resp, err := pc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received non-OK HTTP status %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return pc.parsePlayerResponse(ctx, data, playerID)
}

func (pc *PlayerController) fetchPlayerStatistics(ctx context.Context, playerID int, season int) error {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?id=%d&season=%d", playerID, season)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	fmt.Printf("Fetching statistics for player ID: %d for season %d\n", playerID, season)
	resp, err := pc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received non-OK HTTP status: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var statsResponse PlayerStatisticsResponse
	if err := json.Unmarshal(data, &statsResponse); err != nil {
		return err
	}

	if statsResponse.Results == 0 {
		fmt.Printf("No statistics available for playerID: %d for season %d.\n", playerID, season)
		return nil
	}

	// Assume the first response contains the statistics for the latest season
	stats := statsResponse.Response[0].Statistics[0]
	// Prepare a new PlayerStatistics struct with the data from the API response
	newStats := &player_models.PlayerStats{
		Position: stats.Position,
	}

	// Prepare the UpdatePlayerInput with the new statistics
	updateInput := player_models.UpdatePlayerInput{
		ApiFootballId: playerID,
		Statistics:    newStats, // use the pointer here

	}

	// Call UpdatePlayer to update the statistics in the database
	_, err = pc.playerModel.UpdatePlayer(ctx, updateInput)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PlayerController) parsePlayerResponse(ctx context.Context, data []byte, playerID int) error {
	var response PlayerResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	// Check if the response is empty
	if response.Results == 0 {
		// Log the scenario and return without an error
		fmt.Printf("No data available for playerID: %d for season 2022.\n", playerID)
		return nil
	}

	p := response.Response[0].Player

	if err := pc.downloadPhotoIfNeeded(&p); err != nil {
		return err
	}

	input := player_models.CreatePlayerInput{
		ApiFootballId: p.ID,
		Name:          p.Name,
		Firstname:     p.Firstname,
		Lastname:      p.Lastname,
		Age:           p.Age,
		Height:        p.Height,
		Weight:        p.Weight,
		Injured:       p.Injured,
		Photo:         p.Photo,
		Statistics: player_models.PlayerStats{
			Team: player_models.PSTeam{
				ApiFootballId: p.Statistics.Team.ApiFootballId,
				Name:          p.Statistics.Team.Name,
				Logo:          p.Statistics.Team.Logo,
			},
			League: player_models.PSLeague{
				ApiFootballId: p.Statistics.League.ApiFootballId,
				Name:          p.Statistics.League.Name,
				Country:       p.Statistics.League.Country,
				Logo:          p.Statistics.League.Logo,
				Flag:          p.Statistics.League.Flag,
				Season:        p.Statistics.League.Season,
			},
			Games: player_models.PSGames{
				Appearances: p.Statistics.Games.Appearances,
				Lineups:     p.Statistics.Games.Lineups,
				Minutes:     p.Statistics.Games.Minutes,
				Number:      p.Statistics.Games.Number,
				Position:    p.Statistics.Games.Position,
				Rating:      p.Statistics.Games.Rating,
				Captain:     p.Statistics.Games.Captain,
			},
			Substitutes: player_models.PSSubstitutes{
				In:    p.Statistics.Substitutes.In,
				Out:   p.Statistics.Substitutes.Out,
				Bench: p.Statistics.Substitutes.Bench,
			},
			Shots: player_models.PSShots{
				Total: p.Statistics.Shots.Total,
				On:    p.Statistics.Shots.On,
			},
			Goals: player_models.PSGoals{
				Total:    p.Statistics.Goals.Total,
				Conceded: p.Statistics.Goals.Conceded,
			},
			Passes: player_models.PSPasses{
				Total:    p.Statistics.Passes.Total,
				Key:      p.Statistics.Passes.Key,
				Accuracy: p.Statistics.Passes.Accuracy,
			},
			Tackles: player_models.PSTackles{
				Total:         p.Statistics.Tackles.Total,
				Blocks:        p.Statistics.Tackles.Blocks,
				Interceptions: p.Statistics.Tackles.Interceptions,
			},
			Duels: player_models.PSDuels{
				Total: p.Statistics.Duels.Total,
				Won:   p.Statistics.Duels.Won,
			},
			Dribbles: player_models.PSDribbles{
				Attempts: p.Statistics.Dribbles.Attempts,
				Success:  p.Statistics.Dribbles.Success,
				Past:     p.Statistics.Dribbles.Past,
			},
			Fouls: player_models.PSFouls{
				Drawn:     p.Statistics.Fouls.Drawn,
				Committed: p.Statistics.Fouls.Committed,
			},
			Cards: player_models.PSCards{
				Yellow:    p.Statistics.Cards.Yellow,
				YellowRed: p.Statistics.Cards.YellowRed,
				Red:       p.Statistics.Cards.Red,
			},
			Penalty: player_models.PSPenalty{
				Won:      p.Statistics.Penalty.Won,
				Commited: p.Statistics.Penalty.Committed,
				Scored:   p.Statistics.Penalty.Scored,
				Missed:   p.Statistics.Penalty.Missed,
				Saved:    p.Statistics.Penalty.Saved,
			},
		},
	}

	err := pc.upsertPlayer(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PlayerController) upsertPlayer(ctx context.Context, input player_models.CreatePlayerInput) error {
	return pc.playerModel.WithTransaction(ctx, func(tx *ent.Tx) error {
		exists := tx.Player.Query().Where(player.ApiFootballIdEQ(input.ApiFootballId)).ExistX(ctx)
		if exists {
			updateInput := player_models.UpdatePlayerInput{
				ApiFootballId: input.ApiFootballId,
				Name:          &input.Name,
				Firstname:     &input.Firstname,
				Lastname:      &input.Lastname,
				Age:           &input.Age,
				Height:        &input.Height,
				Weight:        &input.Weight,
				Injured:       &input.Injured,
				Photo:         &input.Photo,
				Statistics: &player_models.PlayerStats{
					Team: player_models.PSTeam{
						ApiFootballId: input.Statistics.Team.ApiFootballId,
						Name:          input.Statistics.Team.Name,
						Logo:          input.Statistics.Team.Logo,
					},
					League: player_models.PSLeague{
						ApiFootballId: input.Statistics.League.ApiFootballId,
						Name:          input.Statistics.League.Name,
						Country:       input.Statistics.League.Country,
						Logo:          input.Statistics.League.Logo,
						Flag:          input.Statistics.League.Flag,
						Season:        input.Statistics.League.Season,
					},
					Games: player_models.PSGames{
						Appearances: input.Statistics.Games.Appearances,
						Lineups:     input.Statistics.Games.Lineups,
						Minutes:     input.Statistics.Games.Minutes,
						Number:      input.Statistics.Games.Number,
						Position:    input.Statistics.Games.Position,
						Rating:      input.Statistics.Games.Rating,
						Captain:     input.Statistics.Games.Captain,
					},
					Substitutes: player_models.PSSubstitutes{
						In:    input.Statistics.Substitutes.In,
						Out:   input.Statistics.Substitutes.Out,
						Bench: input.Statistics.Substitutes.Bench,
					},
					Shots: player_models.PSShots{
						Total: input.Statistics.Shots.Total,
						On:    input.Statistics.Shots.On,
					},
					Goals: player_models.PSGoals{
						Total:    input.Statistics.Goals.Total,
						Conceded: input.Statistics.Goals.Conceded,
					},
					Passes: player_models.PSPasses{
						Total:    input.Statistics.Passes.Total,
						Key:      input.Statistics.Passes.Key,
						Accuracy: input.Statistics.Passes.Accuracy,
					},
					Tackles: player_models.PSTackles{
						Total:         input.Statistics.Tackles.Total,
						Blocks:        input.Statistics.Tackles.Blocks,
						Interceptions: input.Statistics.Tackles.Interceptions,
					},
					Duels: player_models.PSDuels{
						Total: input.Statistics.Duels.Total,
						Won:   input.Statistics.Duels.Won,
					},
					Dribbles: player_models.PSDribbles{
						Attempts: input.Statistics.Dribbles.Attempts,
						Success:  input.Statistics.Dribbles.Success,
						Past:     input.Statistics.Dribbles.Past,
					},
					Fouls: player_models.PSFouls{
						Drawn:     input.Statistics.Fouls.Drawn,
						Committed: input.Statistics.Fouls.Committed,
					},
					Cards: player_models.PSCards{
						Yellow:    input.Statistics.Cards.Yellow,
						YellowRed: input.Statistics.Cards.YellowRed,
						Red:       input.Statistics.Cards.Red,
					},
					Penalty: player_models.PSPenalty{
						Won:      input.Statistics.Penalty.Won,
						Commited: input.Statistics.Penalty.Commited,
						Scored:   input.Statistics.Penalty.Scored,
						Missed:   input.Statistics.Penalty.Missed,
						Saved:    input.Statistics.Penalty.Saved,
					},
				},
			}
			_, err := pc.playerModel.UpdatePlayer(ctx, updateInput)
			if err != nil {
				return err
			}
		} else {
			_, err := pc.playerModel.CreatePlayer(ctx, input)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (pc *PlayerController) GetPlayerIDsForLeague(ctx context.Context, leagueID int) ([]int, error) {
	// Create a URL for fetching players by their league ID
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?league=%d&season=2022", leagueID)
	// Create an HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers for the API
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", viper.GetString("API_KEY"))

	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal data to get player IDs
	var leaguePlayerResponse LeaguePlayerResponse
	if err := json.Unmarshal(body, &leaguePlayerResponse); err != nil {
		return nil, err
	}

	var playerIDs []int
	for _, p := range leaguePlayerResponse.Response {
		playerIDs = append(playerIDs, p.Player.ID)
	}

	return playerIDs, nil
}

func (pc *PlayerController) FetchPlayersByLeague(ctx context.Context, leagueID int) error {
	playerIDs, err := pc.GetPlayerIDsForLeague(ctx, leagueID)
	fmt.Println("Initializing Player... IN CONTROLLER")
	fmt.Println(playerIDs)
	if err != nil {
		return err
	}

	// Ensure each player exists in the database
	for _, playerID := range playerIDs {
		err := pc.EnsurePlayerExists(ctx, playerID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pc *PlayerController) downloadPhotoIfNeeded(p *PlayerInfo) error {
	if p.Photo != "" {
		photoLocation, err := utils.DownloadImageIfNeeded(
			p.Photo,
			fmt.Sprintf("images/players/%d%s", p.ID, filepath.Ext(p.Photo)),
		)
		if err != nil {
			return err
		}
		p.Photo = photoLocation
	}
	return nil
}
