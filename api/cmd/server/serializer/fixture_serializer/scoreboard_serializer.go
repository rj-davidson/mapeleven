package fixture_serializer

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/server/serializer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixture"
	"context"
	"time"
)

type team struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type fixtureSimple struct {
	Slug     string `json:"slug"`
	Referee  string `json:"referee,omitempty"`
	Elapsed  int    `json:"elapsed,omitempty"`
	Round    int    `json:"round,omitempty"`
	Status   string `json:"status,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	HomeTeam team   `json:"homeTeam"`
	AwayTeam team   `json:"awayTeam"`
	Score    score  `json:"score,omitempty"`
}

type lg struct {
	Slug     string                `json:"slug"`
	Name     string                `json:"name"`
	Type     string                `json:"type"`
	Logo     string                `json:"logo"`
	Country  serializer.APICountry `json:"country,omitempty"`
	Fixtures []fixtureSimple       `json:"fixtures,omitempty"`
}

type fixtureDate struct {
	Date    string `json:"date"`
	Leagues []lg   `json:"leagues"`
}

type APIFixturesByDate struct {
	Dates []fixtureDate `json:"date"`
}

func mapEntTeamToTeam(entTeam *ent.Team) team {
	return team{
		Slug: entTeam.Edges.Club.Slug,
		Name: entTeam.Edges.Club.Name,
		Logo: entTeam.Edges.Club.Logo,
	}
}

func mapEntFixtureToAPIFixture(fixture *ent.Fixture) fixtureSimple {
	return fixtureSimple{
		Slug:     fixture.Slug,
		Referee:  fixture.Referee,
		Elapsed:  fixture.Elapsed,
		Round:    fixture.Round,
		Status:   fixture.Status,
		Timezone: fixture.Timezone,
		HomeTeam: mapEntTeamToTeam(fixture.Edges.HomeTeam),
		AwayTeam: mapEntTeamToTeam(fixture.Edges.AwayTeam),
		Score: score{
			Home: fixture.HomeTeamScore,
			Away: fixture.AwayTeamScore,
		},
	}
}

type ScoreboardSerializer struct {
	client *ent.Client
}

func NewScoreboardSerializer(client *ent.Client) *ScoreboardSerializer {
	return &ScoreboardSerializer{client: client}
}

func (ss *ScoreboardSerializer) getScoreboardFixtures() ([]*ent.Fixture, error) {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -21)
	threeWeeksFromNow := now.AddDate(0, 0, 21)

	fixtures, err := ss.client.Fixture.
		Query().
		Where(
			fixture.DateGTE(oneWeekAgo),
			fixture.DateLTE(threeWeeksFromNow),
		).
		WithHomeTeam(func(q *ent.TeamQuery) {
			q.WithClub()
		}).
		WithAwayTeam(func(q *ent.TeamQuery) {
			q.WithClub()
		}).
		WithSeason(func(q *ent.SeasonQuery) {
			q.WithLeague(func(q *ent.LeagueQuery) {
				q.WithCountry()
			})
		}).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return fixtures, nil
}

func (ss *ScoreboardSerializer) SerializeScoreboard() *APIFixturesByDate {
	fixtures, err := ss.getScoreboardFixtures()
	if err != nil {
		return nil
	}
	fixturesByDate := groupFixturesByDateAndLeague(fixtures)

	apiFixturesByDate := &APIFixturesByDate{}

	for date, leagues := range fixturesByDate {
		fixtureDate := fixtureDate{
			Date:    date,
			Leagues: mapLeaguesToAPILeagues(leagues),
		}

		apiFixturesByDate.Dates = append(apiFixturesByDate.Dates, fixtureDate)
	}

	return apiFixturesByDate
}

func groupFixturesByDateAndLeague(fixtures []*ent.Fixture) map[string]map[*ent.League][]*ent.Fixture {
	m := make(map[string]map[*ent.League][]*ent.Fixture)

	for _, f := range fixtures {
		dateStr := f.Date.Format("01-02-2006")
		if _, exists := m[dateStr]; !exists {
			m[dateStr] = make(map[*ent.League][]*ent.Fixture)
		}

		m[dateStr][f.Edges.Season.Edges.League] = append(m[dateStr][f.Edges.Season.Edges.League], f)
	}

	return m
}

func mapLeaguesToAPILeagues(leagues map[*ent.League][]*ent.Fixture) []lg {
	var apiLeagues []lg

	for l, fixtures := range leagues {
		apiLeagues = append(apiLeagues, lg{
			Slug: l.Slug,
			Name: l.Name,
			Type: string(l.Type),
			Logo: l.Logo,
			Country: serializer.APICountry{
				Code: l.Edges.Country.Code,
				Name: l.Edges.Country.Name,
				Flag: l.Edges.Country.Flag,
			},
			Fixtures: mapFixturesToAPIFixtures(fixtures),
		})
	}

	return apiLeagues
}

func mapFixturesToAPIFixtures(fixtures []*ent.Fixture) []fixtureSimple {
	var apiFixtures []fixtureSimple

	for _, f := range fixtures {
		apiFixtures = append(apiFixtures, mapEntFixtureToAPIFixture(f))
	}

	return apiFixtures
}
