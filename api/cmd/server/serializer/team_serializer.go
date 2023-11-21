package serializer

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/server/serializer/teamstats_serializer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"context"
	"sort"
	"time"
)

type APITeamName struct {
	Long  string `json:"long"`
	Short string `json:"short,omitempty"`
}

type APICompetitions struct {
	LeagueItem APILeague    `json:"league"`
	Current    bool         `json:"current"`
	Stats      APITeamStats `json:"stats,omitempty"`
	Squad      APISquad     `json:"squad,omitempty"`
	Fixtures   []APIFixture `json:"fixtures,omitempty"`
}

type APITeamStats struct {
	Form          string                                  `json:"form,omitempty"`
	Biggest       teamstats_serializer.APITSBiggest       `json:"biggest,omitempty"`
	Cards         teamstats_serializer.APITSCards         `json:"cards,omitempty"`
	CleanSheet    teamstats_serializer.APITSCleanSheet    `json:"clean_sheet,omitempty"`
	FailedToScore teamstats_serializer.APITSFailedToScore `json:"failed_to_score,omitempty"`
	Fixtures      teamstats_serializer.APITSFixtures      `json:"fixtures,omitempty"`
	Goals         teamstats_serializer.APITSGoals         `json:"goals,omitempty"`
	Lineups       []teamstats_serializer.APITSLineup      `json:"lineups,omitempty"`
	Penalty       teamstats_serializer.APITSPenalty       `json:"penalty,omitempty"`
}

type APITeam struct {
	Slug         string            `json:"slug"`
	Name         APITeamName       `json:"name,omitempty"`
	Badge        string            `json:"badge"`
	Founded      int               `json:"founded,omitempty"`
	NationalTeam bool              `json:"nationalTeam,omitempty"`
	Country      APICountry        `json:"country,omitempty"`
	Competitions []APICompetitions `json:"competitions,omitempty"`
	Popularity   int               `json:"popularity,omitempty"`
}

type APISquad struct {
	Players []APIPlayer `json:"players"`
}

type APIPlayer struct {
	Slug        string `json:"slug"`
	Name        string `json:"name,omitempty"`
	Photo       string `json:"photo,omitempty"`
	Position    string `json:"position,omitempty"`
	SquadNumber int    `json:"number,omitempty"`
}

type APIFixture struct {
	Slug      string    `json:"slug"`
	HomeTeam  APITeam   `json:"homeTeam"`
	AwayTeam  APITeam   `json:"awayTeam"`
	HomeScore int       `json:"homeScore,omitempty"`
	AwayScore int       `json:"awayScore,omitempty"`
	League    APILeague `json:"league"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
}

type TeamSerializer struct {
	client *ent.Client
}

func NewTeamSerializer(client *ent.Client) *TeamSerializer {
	return &TeamSerializer{client: client}
}

func (ts *TeamSerializer) SerializeTeam(club *ent.Club) *APITeam {
	ls := NewLeagueSerializer(ts.client)

	var compList = make([]APICompetitions, 0)
	for _, t := range club.Edges.Team {
		tsList := APITeamStats{
			Form:          t.Form,
			Biggest:       teamstats_serializer.SerializeTSBiggest(t.Edges.BiggestStats),
			Cards:         teamstats_serializer.SerializeTSCards(t.Edges.CardsStats),
			CleanSheet:    teamstats_serializer.SerializeTSCleanSheet(t.Edges.CleanSheetStats),
			FailedToScore: teamstats_serializer.SerializeTSFailedToScore(t.Edges.FailedToScoreStats),
			Fixtures:      teamstats_serializer.SerializeTSFixtures(t.Edges.FixturesStats),
			Goals:         teamstats_serializer.SerializeTSGoals(t.Edges.GoalsStats),
			Lineups:       teamstats_serializer.SerializeTSLineups(t.Edges.Lineups),
			Penalty:       teamstats_serializer.SerializeTSPenalty(t.Edges.PenaltyStats),
		}
		li, err := ls.Serialize(t.Edges.Season.Edges.League)
		if err != nil {
			continue
		}
		compList = append(compList, APICompetitions{
			LeagueItem: *li,
			Current:    t.Edges.Season.Current,
			Stats:      tsList,
			Squad:      ts.serializeCompetitionSquad(t.Edges.Squad),
			Fixtures:   ts.serializeFixtures(t.Edges.AwayFixtures, t.Edges.HomeFixtures),
		})
	}

	return &APITeam{
		Slug:         club.Slug,
		Name:         APITeamName{Long: club.Name, Short: club.Code},
		Badge:        club.Logo,
		Founded:      club.Founded,
		NationalTeam: club.National,
		Country: APICountry{
			Code: club.Edges.Country.Code,
			Name: club.Edges.Country.Name,
			Flag: club.Edges.Country.Flag,
		},
		Competitions: compList,
		Popularity:   club.Popularity,
	}
}

func (ts *TeamSerializer) GetTeamBySlug(ctx context.Context, slug string, seasonStr string) (*APITeam, error) {
	year, err := time.Parse("2006", seasonStr)
	if err != nil {
		year = time.Now()
	}
	_, err = ts.client.Club.Update().
		Where(club.SlugEQ(slug)).
		AddPopularity(1).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	t, err := ts.client.Club.Query().
		Where(club.Slug(slug)).
		WithCountry().
		WithTeam(
			func(q *ent.TeamQuery) {
				q.WithSquad(func(q *ent.SquadQuery) {
					q.WithPlayer()
				})
				q.WithSeason(
					func(q *ent.SeasonQuery) {
						q.Where(season.YearEQ(year.Year())).
							WithLeague(
								func(q *ent.LeagueQuery) {
									q.WithCountry()
								},
							)
					},
				)
				q.WithBiggestStats()
				q.WithCardsStats()
				q.WithCleanSheetStats()
				q.WithFailedToScoreStats()
				q.WithFixturesStats()
				q.WithGoalsStats()
				q.WithLineups()
				q.WithPenaltyStats()
				q.WithAwayFixtures(
					func(q *ent.FixtureQuery) {
						q.WithAwayTeam(func(tq *ent.TeamQuery) {
							tq.WithClub()
						})
						q.WithHomeTeam(func(tq *ent.TeamQuery) {
							tq.WithClub()
						})
						q.WithSeason(func(sq *ent.SeasonQuery) {
							sq.WithLeague()
						})
					})
				q.WithHomeFixtures(
					func(q *ent.FixtureQuery) {
						q.WithAwayTeam(func(tq *ent.TeamQuery) {
							tq.WithClub()
						})
						q.WithHomeTeam(func(tq *ent.TeamQuery) {
							tq.WithClub()
						})
						q.WithSeason(func(sq *ent.SeasonQuery) {
							sq.WithLeague()
						})
					})
			},
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return ts.SerializeTeam(t), nil
}

func (ts *TeamSerializer) GetTeams(ctx context.Context, seasonStr string) ([]*APITeam, error) {
	year, err := time.Parse("2006", seasonStr)
	if err != nil {
		year = time.Now()
	}

	teams, err := ts.client.Club.Query().
		WithCountry().
		WithTeam(
			func(q *ent.TeamQuery) {
				q.WithSquad(func(q *ent.SquadQuery) {
					q.WithPlayer()
				})
				q.WithSeason(
					func(q *ent.SeasonQuery) {
						q.Where(season.YearEQ(year.Year())).
							WithLeague(
								func(q *ent.LeagueQuery) {
									q.WithCountry()
								},
							)
					},
				)
				q.WithBiggestStats()
				q.WithCardsStats()
				q.WithCleanSheetStats()
				q.WithFailedToScoreStats()
				q.WithFixturesStats()
				q.WithGoalsStats()
				q.WithLineups()
				q.WithPenaltyStats()
				q.WithAwayFixtures(
					func(q *ent.FixtureQuery) {
						q.WithAwayTeam(func(tq *ent.TeamQuery) {
							tq.WithClub()
						})
						q.WithHomeTeam(func(tq *ent.TeamQuery) {
							tq.WithClub()
						})
						q.WithSeason(func(sq *ent.SeasonQuery) {
							sq.WithLeague()
						})
					})
				q.WithHomeFixtures(
					func(q *ent.FixtureQuery) {
						q.WithAwayTeam(func(tq *ent.TeamQuery) {
							tq.WithClub()
						})
						q.WithHomeTeam(func(tq *ent.TeamQuery) {
							tq.WithClub()
						})
						q.WithSeason(func(sq *ent.SeasonQuery) {
							sq.WithLeague()
						})
					})
			}).
		All(ctx)

	if err != nil {
		return nil, err
	}

	var teamItems []*APITeam
	for _, t := range teams {
		serializedTeam := ts.SerializeTeam(t)
		teamItems = append(teamItems, serializedTeam)
	}

	// Sort teamItems based on Popularity
	sort.Slice(teamItems, func(i, j int) bool {
		return teamItems[i].Popularity > teamItems[j].Popularity
	})

	return teamItems, nil
}

func (ts *TeamSerializer) serializeCompetitionSquad(squads []*ent.Squad) APISquad {
	var players []APIPlayer
	for _, s := range squads {
		players = append(players, ts.serializeSquadPlayer(s))
	}

	return APISquad{Players: players}
}

func (ts *TeamSerializer) serializeSquadPlayer(s *ent.Squad) APIPlayer {
	return APIPlayer{
		Name:        s.Edges.Player.Name,
		Slug:        s.Edges.Player.Slug,
		Photo:       s.Edges.Player.Photo,
		Position:    s.Position,
		SquadNumber: s.Number,
	}
}

func (ts *TeamSerializer) serializeFixtures(home []*ent.Fixture, away []*ent.Fixture) []APIFixture {
	var fixtures []APIFixture
	for _, f := range home {
		fixtures = append(fixtures, ts.serializeFixture(f))
	}
	for _, f := range away {
		fixtures = append(fixtures, ts.serializeFixture(f))
	}

	return fixtures
}

func (ts *TeamSerializer) serializeFixture(f *ent.Fixture) APIFixture {
	// If the fixture status is not FT AET or PEN, then no need to get the score
	var homeScore int
	var awayScore int
	if f.Status == "FT" || f.Status == "AET" || f.Status == "PEN" {
		homeScore = f.HomeTeamScore
		awayScore = f.AwayTeamScore
	}
	return APIFixture{
		Slug: f.Slug,
		HomeTeam: APITeam{
			Slug:  f.Edges.HomeTeam.Edges.Club.Slug,
			Name:  APITeamName{Long: f.Edges.HomeTeam.Edges.Club.Name, Short: f.Edges.HomeTeam.Edges.Club.Code},
			Badge: f.Edges.HomeTeam.Edges.Club.Logo,
		},
		AwayTeam: APITeam{
			Slug:  f.Edges.HomeTeam.Edges.Club.Slug,
			Name:  APITeamName{Long: f.Edges.HomeTeam.Edges.Club.Name, Short: f.Edges.HomeTeam.Edges.Club.Code},
			Badge: f.Edges.HomeTeam.Edges.Club.Logo,
		},
		HomeScore: homeScore,
		AwayScore: awayScore,
		League: APILeague{
			Slug: f.Edges.Season.Edges.League.Slug,
			Name: f.Edges.Season.Edges.League.Name,
			Type: string(f.Edges.Season.Edges.League.Type),
			Logo: f.Edges.Season.Edges.League.Logo,
		},
		Date:   f.Date,
		Status: f.Status,
	}
}
