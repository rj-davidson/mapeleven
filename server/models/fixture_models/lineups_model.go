package fixture_models

import (
	"context"
	"fmt"
	"mapeleven/db/ent"
	"mapeleven/db/ent/fixture"
	"mapeleven/db/ent/fixturelineups"
	"mapeleven/db/ent/matchplayer"
	"mapeleven/db/ent/player"
	"mapeleven/db/ent/team"
	"strconv"
	"strings"
)

func UpsertLineup(f *ent.Fixture, t *ent.Team, lineupInfo LineupInfo, client *ent.Client, ctx context.Context) (*ent.FixtureLineups, error) {
	lu, err := client.FixtureLineups.
		Query().
		Where(
			fixturelineups.HasTeamWith(
				team.IDEQ(t.ID),
			),
			fixturelineups.HasFixtureWith(
				fixture.IDEQ(f.ID),
			),
		).
		Only(ctx)

	if lu != nil {
		// Update lineup
		lu, err = client.FixtureLineups.
			UpdateOne(lu).
			SetFormation(lineupInfo.Formation).
			Save(ctx)
	} else {
		// Create lineup
		lu, err = client.FixtureLineups.
			Create().
			SetFormation(lineupInfo.Formation).
			SetTeam(t).
			SetFixture(f).
			Save(ctx)
	}

	for _, pdi := range lineupInfo.StartXI {
		_, err = upsertMatchPlayer(lu, pdi.Player, client, ctx)
		if err != nil {
			return nil, err
		}
	}

	for _, pdi := range lineupInfo.Substitutes {
		_, err = upsertMatchPlayer(lu, pdi.Player, client, ctx)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return lu, nil
}

func upsertMatchPlayer(lu *ent.FixtureLineups, pdi PlayerDetailInfo, client *ent.Client, ctx context.Context) (*ent.MatchPlayer, error) {
	p, err := client.Player.Query().Where(player.ApiFootballIdEQ(pdi.ID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	mp, err := client.MatchPlayer.
		Query().
		Where(
			matchplayer.HasLineupWith(
				fixturelineups.IDEQ(lu.ID),
			),
			matchplayer.HasPlayerWith(
				player.IDEQ(p.ID),
			),
		).
		Only(ctx)

	var x, y *int
	if pdi.Grid != "" {
		parts := strings.Split(pdi.Grid, ":")
		if len(parts) != 2 {
			fmt.Println("Unexpected format")
			return nil, fmt.Errorf("lineups_model.go: unexpected grid format: %s", pdi.Grid)
		}

		valX, errX := strconv.Atoi(parts[0])
		valY, errY := strconv.Atoi(parts[1])
		if errX == nil && errY == nil {
			x = &valX
			y = &valY
		} else {
			fmt.Println("lineups_model.go: Error converting to integer")
			return nil, fmt.Errorf("lineups_model.go: error converting grid values to integers")
		}
	}

	if mp != nil {
		upd := client.MatchPlayer.UpdateOne(mp).SetPosition(pdi.Pos).SetNumber(pdi.Number)
		if x != nil {
			upd = upd.SetX(*x)
		}
		if y != nil {
			upd = upd.SetY(*y)
		}
		mp, err = upd.Save(ctx)
	} else {
		cre := client.MatchPlayer.Create().SetPosition(pdi.Pos).SetNumber(pdi.Number).SetPlayer(p).SetLineup(lu)
		if x != nil {
			cre = cre.SetX(*x)
		}
		if y != nil {
			cre = cre.SetY(*y)
		}
		mp, err = cre.Save(ctx)
	}

	if err != nil {
		return nil, err
	}

	return mp, nil
}
