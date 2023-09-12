package serializer

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/team"
)

type APITeamName struct {
	Long  string `json:"long"`
	Short string `json:"short,omitempty"`
}

type APICompetitions struct {
	LeagueItem APILeague `json:"league"`
	Current    bool      `json:"current"`
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

func SerializeTeam(team *ent.Team) *APITeam {
	return &APITeam{
		Slug:         team.Slug,
		Name:         APITeamName{Long: team.Name, Short: team.Code},
		Badge:        team.Logo,
		Founded:      team.Founded,
		NationalTeam: team.National,
		Country: APICountry{
			Code: team.Edges.Country.Code,
			Name: team.Edges.Country.Name,
			Flag: team.Edges.Country.Flag,
		},
	}
}

func (ts *TeamSerializer) GetTeamBySlug(ctx context.Context, slug string) (*APITeam, error) {
	t, err := ts.client.Team.Query().
		Where(team.Slug(slug)).
		WithCountry().
		First(ctx)

	if err != nil {
		return nil, err
	}

	return SerializeTeam(t), nil
}

func (ts *TeamSerializer) GetTeams(ctx context.Context) ([]*APITeam, error) {
	teams, err := ts.client.Team.Query().
		WithCountry().
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
