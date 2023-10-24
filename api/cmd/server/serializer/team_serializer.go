package serializer

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/server/serializer/teamstats_serializer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"context"
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
	}
}

func (ts *TeamSerializer) GetTeamBySlug(ctx context.Context, slug string, seasonStr string) (*APITeam, error) {
	year, err := time.Parse("2006", seasonStr)
	if err != nil {
		year = time.Now()
	}

	t, err := ts.client.Club.Query().
		Where(club.Slug(slug)).
		WithCountry().
		WithTeam(
			func(q *ent.TeamQuery) {
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
			}).
		All(ctx)

	if err != nil {
		return nil, err
	}

	var teamItems []*APITeam
	for _, t := range teams {
		teamItems = append(teamItems, ts.SerializeTeam(t))
	}

	return teamItems, nil
}
