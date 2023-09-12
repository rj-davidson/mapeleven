package serializer

import (
	"context"
	"fmt"
	"mapeleven/db/ent"
	"mapeleven/db/ent/league"
)

type LeagueItem struct {
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Logo      string     `json:"logo"`
	Country   Country    `json:"country,omitempty"`
	Standings []Standing `json:"standings,omitempty"`
}

type LeagueSerializer struct {
	client *ent.Client
}

func NewLeagueSerializer(client *ent.Client) *LeagueSerializer {
	return &LeagueSerializer{client: client}
}

func SerializeLeague(league *ent.League) (*LeagueItem, error) {
	countryItem, err := SerializeCountry(league.Edges.Country)
	standings := make([]Standing, len(league.Edges.Standings))
	for _, standing := range league.Edges.Standings {
		s := SerializeStanding(standing)
		standings = append(standings, *s)
	}
	if err != nil {
		return nil, fmt.Errorf("serialize country for league %s: %w", league.Name, err)
	}

	return &LeagueItem{
		Slug:      league.Slug,
		Name:      league.Name,
		Type:      league.Type.String(),
		Logo:      league.Logo,
		Country:   *countryItem,
		Standings: standings,
	}, nil
}

func (ls *LeagueSerializer) GetLeagueBySlug(ctx context.Context, slug string) (*LeagueItem, error) {
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

func (ls *LeagueSerializer) GetLeagues(ctx context.Context) ([]*LeagueItem, error) {
	leagues, err := ls.client.League.Query().
		WithCountry().
		WithStandings(func(q *ent.StandingsQuery) {
			q.WithLeague().WithTeam()
		}).
		All(ctx)

	if err != nil {
		return nil, err
	}

	leagueItems := make([]*LeagueItem, len(leagues))
	for i, l := range leagues {
		leagueItems[i], err = SerializeLeague(l)
		if err != nil {
			return nil, err
		}
	}

	return leagueItems, nil
}
