package serializer

import (
	"context"
	"fmt"
	"mapeleven/db/ent"
	"mapeleven/db/ent/league"
)

type APILeague struct {
	Slug      string        `json:"slug"`
	Name      string        `json:"name"`
	Type      string        `json:"type"`
	Logo      string        `json:"logo"`
	Country   APICountry    `json:"country,omitempty"`
	Standings []APIStanding `json:"standings,omitempty"`
}

type LeagueSerializer struct {
	client *ent.Client
}

func NewLeagueSerializer(client *ent.Client) *LeagueSerializer {
	return &LeagueSerializer{client: client}
}

func SerializeLeague(league *ent.League) (*APILeague, error) {
	countryItem, err := SerializeCountry(league.Edges.Country)
	APIStandings := make([]APIStanding, len(league.Edges.Standings))
	for _, standing := range league.Edges.Standings {
		s := SerializeStanding(standing)
		APIStandings = append(APIStandings, *s)
	}
	if err != nil {
		return nil, fmt.Errorf("serialize country for league %s: %w", league.Name, err)
	}

	return &APILeague{
		Slug:      league.Slug,
		Name:      league.Name,
		Type:      league.Type.String(),
		Logo:      league.Logo,
		Country:   *countryItem,
		Standings: APIStandings,
	}, nil
}

func (ls *LeagueSerializer) GetLeagueBySlug(ctx context.Context, slug string) (*APILeague, error) {
	l, err := ls.client.League.Query().
		Where(league.Slug(slug)).
		WithCountry().
		WithStandings(func(q *ent.StandingsQuery) {
			q.WithLeague().WithTeam()
		}).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return SerializeLeague(l)
}

func (ls *LeagueSerializer) GetLeagues(ctx context.Context) ([]*APILeague, error) {
	leagues, err := ls.client.League.Query().
		WithCountry().
		WithStandings(func(q *ent.StandingsQuery) {
			q.WithLeague().WithTeam()
		}).
		All(ctx)

	if err != nil {
		return nil, err
	}

	leagueItems := make([]*APILeague, len(leagues))
	for i, l := range leagues {
		leagueItems[i], err = SerializeLeague(l)
		if err != nil {
			return nil, err
		}
	}

	return leagueItems, nil
}
