package fixture_serializer

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/fixture"
)

type FixtureSerializer struct {
	client *ent.Client
}

func NewFixtureSerializer(client *ent.Client) *FixtureSerializer {
	return &FixtureSerializer{client: client}
}

func (fs *FixtureSerializer) GetFixtureBySlug(ctx context.Context, slug string) (*APIFixtureSet, error) {
	f, err := fs.client.Fixture.
		Query().
		Where(
			fixture.Slug(slug),
		).
		WithSeason().
		WithLineups().
		WithHomeTeam(
			func(q *ent.TeamQuery) {
				q.WithClub()
			},
		).
		WithAwayTeam(
			func(q *ent.TeamQuery) {
				q.WithClub()
			},
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	sfxt, std, err := fs.Serialize(f)
	if err != nil {
		return nil, err
	}
	return &APIFixtureSet{
		Fixture: sfxt,
		Teams:   std,
	}, nil
}

func (fs *FixtureSerializer) Serialize(f *ent.Fixture) (*APIFixture, *APITeamDetails, error) {
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
