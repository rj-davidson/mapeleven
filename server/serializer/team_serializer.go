package serializer

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/club"
	"mapeleven/db/ent/season"
	"mapeleven/serializer/teamstats_serializer"
)

type APITeamName struct {
	Long  string `json:"long"`
	Short string `json:"short,omitempty"`
}

type APICompetitions struct {
	LeagueItem APILeague `json:"league"`
	Current    bool      `json:"current"`
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
	Stats        APITeamStats      `json:"stats,omitempty"`
}

type TeamSerializer struct {
	client *ent.Client
}

func NewTeamSerializer(client *ent.Client) *TeamSerializer {
	return &TeamSerializer{client: client}
}

func SerializeTeam(club *ent.Club) *APITeam {
	var tsList = make([]APITeamStats, 0)
	for _, t := range club.Edges.Team {
		tsList = append(tsList, APITeamStats{
			Form:          t.Form,
			Biggest:       teamstats_serializer.SerializeTSBiggest(t.Edges.BiggestStats),
			Cards:         teamstats_serializer.SerializeTSCards(t.Edges.CardsStats),
			CleanSheet:    teamstats_serializer.SerializeTSCleanSheet(t.Edges.CleanSheetStats),
			FailedToScore: teamstats_serializer.SerializeTSFailedToScore(t.Edges.FailedToScoreStats),
			Fixtures:      teamstats_serializer.SerializeTSFixtures(t.Edges.FixturesStats),
			Goals:         teamstats_serializer.SerializeTSGoals(t.Edges.GoalsStats),
			Lineups:       teamstats_serializer.SerializeTSLineups(t.Edges.Lineups),
			Penalty:       teamstats_serializer.SerializeTSPenalty(t.Edges.PenaltyStats),
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
		Stats: tsList[len(tsList)-1],
	}
}

func (ts *TeamSerializer) GetTeamBySlug(ctx context.Context, slug string) (*APITeam, error) {
	t, err := ts.client.Club.Query().
		Where(club.Slug(slug)).
		WithCountry().
		WithTeam(
			func(q *ent.TeamQuery) {
				q.WithSeason(
					func(q *ent.SeasonQuery) {
						q.Where(season.Current(true))
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

	return SerializeTeam(t), nil
}

func (ts *TeamSerializer) GetTeams(ctx context.Context) ([]*APITeam, error) {
	teams, err := ts.client.Club.Query().
		WithCountry().
		WithTeam(
			func(q *ent.TeamQuery) {
				q.WithSeason(
					func(q *ent.SeasonQuery) {
						q.Where(season.Current(true))
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
		teamItems = append(teamItems, SerializeTeam(t))
	}

	return teamItems, nil
}
