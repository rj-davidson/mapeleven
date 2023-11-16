package fixture_serializer

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixture"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixtureevents"
	"context"
	"time"
)

type DemoFixtureSerializer struct {
	client *ent.Client
}

func NewDemoFixtureSerializer(client *ent.Client) *DemoFixtureSerializer {
	return &DemoFixtureSerializer{client: client}
}

func (dfs *DemoFixtureSerializer) GetDemoFixtureBySlug(ctx context.Context, slug string, getFixture, getEvents, getLineups bool) (*APIFixtureSet, error) {
	currentCycleTime := time.Now().Unix() % (5 * 60) / 60
	maxElapsedForCycle := int(currentCycleTime * 24)

	// Existing query code...
	f, err := dfs.client.Fixture.
		Query().
		Where(fixture.Slug(slug)).
		WithSeason().
		WithLineups(func(q *ent.FixtureLineupsQuery) {
			q.WithTeam(func(q *ent.TeamQuery) {
				q.WithClub()
			}).WithLineupPlayer(func(q *ent.MatchPlayerQuery) {
				q.WithPlayer()
			})
		}).
		WithHomeTeam(func(q *ent.TeamQuery) {
			q.WithClub()
		}).
		WithAwayTeam(func(q *ent.TeamQuery) {
			q.WithClub()
		}).
		WithFixtureEvents(func(q *ent.FixtureEventsQuery) {
			q.Where(fixtureevents.ElapsedTimeLT(maxElapsedForCycle))
			q.WithPlayer()
			q.WithAssist()
			q.WithTeam(
				func(q *ent.TeamQuery) {
					q.WithClub()
				},
			)
		},
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	var sFixture *APIFixture
	var sTeams *APITeamDetails
	var sEvents *[]APIEvent
	sLineups := &APIHomeAway{}

	if getFixture {
		sFixture, sTeams, err = dfs.serializeDemoFixture(f, maxElapsedForCycle)
		if err != nil {
			return nil, err
		}
	}

	if getEvents {
		sEvents, err = SerializeEvents(f.Edges.FixtureEvents)
		if err != nil {
			return nil, err
		}
	}

	if getLineups {
		for _, l := range f.Edges.Lineups {
			sLineup, err := SerializeLineup(l)
			if err != nil {
				return nil, err
			}
			if l.Edges.Team.Edges.Club.Slug == f.Edges.HomeTeam.Edges.Club.Slug {
				sLineups.Home = sLineup
			} else {
				sLineups.Away = sLineup
			}
		}
	}

	return &APIFixtureSet{
		Fixture: sFixture,
		Teams:   sTeams,
		Events:  sEvents,
		Lineups: sLineups,
	}, nil
}

func (dfs *DemoFixtureSerializer) serializeDemoFixture(f *ent.Fixture, maxElapsedForCycle int) (*APIFixture, *APITeamDetails, error) {
	var fxt APIFixture
	var td APITeamDetails
	homeTeamScore := 0
	awayTeamScore := 0

	for _, e := range f.Edges.FixtureEvents {
		if e.Type == "Goal" && e.ElapsedTime <= maxElapsedForCycle {
			if e.Edges.Team.Edges.Club.Slug == f.Edges.HomeTeam.Edges.Club.Slug {
				homeTeamScore++
			} else if e.Edges.Team.Edges.Club.Slug == f.Edges.AwayTeam.Edges.Club.Slug {
				awayTeamScore++
			}
		}
	}

	fxt = APIFixture{
		Slug:          f.Slug,
		Referee:       f.Referee,
		Timezone:      f.Timezone,
		Date:          f.Date,
		Elapsed:       f.Elapsed,
		Round:         f.Round,
		Status:        f.Status,
		HomeTeamScore: homeTeamScore,
		AwayTeamScore: awayTeamScore,
		Season: APISeason{
			Year:    f.Edges.Season.Year,
			Start:   f.Edges.Season.StartDate,
			End:     f.Edges.Season.EndDate,
			Current: f.Edges.Season.Current,
		},
	}

	var homeWinner *bool
	var awayWinner *bool
	if f.Status == "Match Finished" {
		homeWin := f.HomeTeamScore > f.AwayTeamScore
		awayWin := f.AwayTeamScore > f.HomeTeamScore
		homeWinner = &homeWin
		awayWinner = &awayWin
	}

	td = APITeamDetails{
		Home: APITeamInfo{
			Slug:   f.Edges.HomeTeam.Edges.Club.Slug,
			Name:   f.Edges.HomeTeam.Edges.Club.Name,
			Logo:   f.Edges.HomeTeam.Edges.Club.Logo,
			Winner: homeWinner,
		},
		Away: APITeamInfo{
			Slug:   f.Edges.AwayTeam.Edges.Club.Slug,
			Name:   f.Edges.AwayTeam.Edges.Club.Name,
			Logo:   f.Edges.AwayTeam.Edges.Club.Logo,
			Winner: awayWinner,
		},
	}
	return &fxt, &td, nil
}
