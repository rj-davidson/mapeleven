package fixture_serializer

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixture"
	"context"
)

type FixtureSerializer struct {
	client *ent.Client
}

func NewFixtureSerializer(client *ent.Client) *FixtureSerializer {
	return &FixtureSerializer{client: client}
}

func (dfs *FixtureSerializer) GetFixtureBySlug(ctx context.Context, slug string, getFixture, getEvents, getLineups bool) (*APIFixtureSet, error) {
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
		sFixture, sTeams, err = dfs.serializeFixture(f)
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

	// If fixture status is not "Match Finished", then the score should be calculated using "Goal" events that are not "Missed Penalty" events
	if f.Status != "Match Finished" {
		var homeScore int
		var awayScore int
		for _, e := range *sEvents {
			if e.Type == "Goal" && e.Assist == nil {
				if e.Team.Slug == sTeams.Home.Slug {
					homeScore++
				} else {
					awayScore++
				}
			}
		}
		sFixture.HomeTeamScore = homeScore
		sFixture.AwayTeamScore = awayScore
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

func (dfs *FixtureSerializer) serializeFixture(f *ent.Fixture) (*APIFixture, *APITeamDetails, error) {
	var fxt APIFixture
	var td APITeamDetails

	fxt = APIFixture{
		Slug:          f.Slug,
		Referee:       f.Referee,
		Timezone:      f.Timezone,
		Date:          f.Date,
		Elapsed:       f.Elapsed,
		Round:         f.Round,
		Status:        f.Status,
		HomeTeamScore: f.HomeTeamScore,
		AwayTeamScore: f.AwayTeamScore,
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
