package fixture_serializer

import "mapeleven/db/ent"

func SerializeEvents(es []*ent.FixtureEvents) (*[]APIEvent, error) {
	var events []APIEvent
	for _, e := range es {
		event, err := serializeEvent(e)
		if err != nil {
			return nil, err
		}
		events = append(events, *event)
	}
	return &events, nil
}

func serializeEvent(e *ent.FixtureEvents) (*APIEvent, error) {
	var event APIEvent
	event = APIEvent{
		Time: APITimeDetail{
			Elapsed: e.ElapsedTime,
			Extra:   &e.ExtraTime,
		},
		Team: APITeamInfo{
			Slug: e.Edges.Team.Edges.Club.Slug,
			Name: e.Edges.Team.Edges.Club.Name,
			Logo: e.Edges.Team.Edges.Club.Logo,
		},
		Player: APIPlayer{
			Slug: e.Edges.Player.Slug,
			Name: e.Edges.Player.Name,
		},
		Type:     e.Type,
		Detail:   e.Detail,
		Comments: &e.Comments,
	}

	if e.Edges.Assist != nil {
		event.Assist = &APIPlayer{
			Slug: e.Edges.Assist.Slug,
			Name: e.Edges.Assist.Name,
		}
	}

	return &event, nil
}
