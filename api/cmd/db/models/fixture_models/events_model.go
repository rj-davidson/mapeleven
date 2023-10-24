package fixture_models

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/club"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixture"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixtureevents"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"context"
	"fmt"
)

// RepopulateEvents Remove all events for the fixture, then add the new ones
func RepopulateEvents(f *ent.Fixture, client *ent.Client, events []Event, ctx context.Context) error {
	if err := deleteFixtureEvents(f, client, ctx); err != nil {
		return err
	}

	teamCache := make(map[int]*ent.Team)
	playerCache := make(map[int]*ent.Player)

	for _, event := range events {
		if err := createFixtureEvent(f, client, event, ctx, teamCache, playerCache); err != nil {
			return err
		}
	}

	return nil
}

func deleteFixtureEvents(f *ent.Fixture, client *ent.Client, ctx context.Context) error {
	_, err := client.FixtureEvents.
		Delete().
		Where(
			fixtureevents.HasFixtureWith(fixture.IDEQ(f.ID)),
		).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("error deleting fixture events: %w", err)
	}
	return nil
}

func getOrCreateTeam(client *ent.Client, event Event, ctx context.Context, teamCache map[int]*ent.Team) (*ent.Team, error) {
	if t, ok := teamCache[event.Team.ID]; ok {
		return t, nil
	}
	t, err := client.Team.
		Query().
		Where(
			team.HasClubWith(club.ApiFootballIdEQ(event.Team.ID)),
		).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting team: %w", err)
	}
	teamCache[event.Team.ID] = t
	return t, nil
}

func getOrCreatePlayer(client *ent.Client, playerId int, ctx context.Context, playerCache map[int]*ent.Player) (*ent.Player, error) {
	if p, ok := playerCache[playerId]; ok {
		return p, nil
	}
	p, err := client.Player.
		Query().
		Where(
			player.ApiFootballIdEQ(playerId),
		).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting player: %w", err)
	}
	playerCache[playerId] = p
	return p, nil
}

func createFixtureEvent(f *ent.Fixture, client *ent.Client, event Event, ctx context.Context, teamCache map[int]*ent.Team, playerCache map[int]*ent.Player) error {
	t, err := getOrCreateTeam(client, event, ctx, teamCache)
	if err != nil {
		return err
	}

	p, err := getOrCreatePlayer(client, event.Player.ID, ctx, playerCache)
	if err != nil {
		return err
	}

	var ap *ent.Player
	if event.Assist != nil && event.Assist.ID != 0 {
		ap, err = getOrCreatePlayer(client, event.Assist.ID, ctx, playerCache)
		if err != nil {
			return err
		}
	}

	builder := client.FixtureEvents.
		Create().
		SetFixture(f).
		SetTeam(t).
		SetPlayer(p).
		SetType(event.Type).
		SetDetail(event.Detail).
		SetElapsedTime(event.Time.Elapsed)

	// Handle potential nil values
	if event.Comments != nil {
		builder.SetComments(*event.Comments)
	}
	if event.Time.Extra != nil {
		builder.SetExtraTime(*event.Time.Extra)
	}
	if ap != nil {
		builder.SetAssist(ap)
	}

	_, err = builder.Save(ctx)
	if err != nil {
		return fmt.Errorf("error saving fixture event: %w", err)
	}

	return nil
}
