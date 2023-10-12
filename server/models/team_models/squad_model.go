package team_models

import (
	"context"
	"mapeleven/db/ent"
	"mapeleven/db/ent/player"
	"mapeleven/db/ent/squad"
	"mapeleven/db/ent/team"
)

type Squad struct {
	Results  int            `json:"results"`
	Response []TeamResponse `json:"response"`
}

type TeamResponse struct {
	Team    Team     `json:"team"`
	Players []Player `json:"players"`
}

type Team struct {
	ApiFootballId int    `json:"id"`
	Name          string `json:"name"`
	Logo          string `json:"logo"`
}

type Player struct {
	ApiFootballId int    `json:"id"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Number        int    `json:"number"`
	Position      string `json:"position"`
	Photo         string `json:"photo"`
}

type SquadModel struct {
	client *ent.Client
}

func NewSquadModel(client *ent.Client) *SquadModel {
	return &SquadModel{client: client}
}

func (m *SquadModel) UpsertSquadPlayer(t *ent.Team, s Player, ctx context.Context) (*ent.Squad, error) {
	sq, err := m.client.Squad.
		Query().
		Where(squad.HasPlayerWith(player.ApiFootballIdEQ(s.ApiFootballId)), squad.HasTeamWith(team.ID(t.ID))).
		Only(ctx)

	p, err := m.client.Player.
		Query().
		Where(player.ApiFootballIdEQ(s.ApiFootballId)).
		Only(ctx)
	if p == nil {
		// TODO: Better error logging here when the player doesn't exist
		return nil, nil
	}

	if sq == nil {
		sq, err = m.client.Squad.
			Create().
			SetPosition(s.Position).
			SetNumber(s.Number).
			SetPlayer(p).
			SetTeam(t).
			Save(ctx)

		if err != nil {
			return nil, err
		}
	} else {
		sq, err = sq.Update().
			SetPosition(s.Position).
			SetNumber(s.Number).
			Save(ctx)

		if err != nil {
			return nil, err
		}
	}

	return sq, nil
}
