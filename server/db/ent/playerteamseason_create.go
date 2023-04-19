// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/player"
	"mapeleven/db/ent/playerteamseason"
	"mapeleven/db/ent/teamseason"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlayerTeamSeasonCreate is the builder for creating a PlayerTeamSeason entity.
type PlayerTeamSeasonCreate struct {
	config
	mutation *PlayerTeamSeasonMutation
	hooks    []Hook
}

// SetPlayerTeamSeasonID sets the "player_team_season_id" field.
func (ptsc *PlayerTeamSeasonCreate) SetPlayerTeamSeasonID(i int) *PlayerTeamSeasonCreate {
	ptsc.mutation.SetPlayerTeamSeasonID(i)
	return ptsc
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (ptsc *PlayerTeamSeasonCreate) SetPlayerID(id int) *PlayerTeamSeasonCreate {
	ptsc.mutation.SetPlayerID(id)
	return ptsc
}

// SetPlayer sets the "player" edge to the Player entity.
func (ptsc *PlayerTeamSeasonCreate) SetPlayer(p *Player) *PlayerTeamSeasonCreate {
	return ptsc.SetPlayerID(p.ID)
}

// SetTeamSeasonID sets the "teamSeason" edge to the TeamSeason entity by ID.
func (ptsc *PlayerTeamSeasonCreate) SetTeamSeasonID(id int) *PlayerTeamSeasonCreate {
	ptsc.mutation.SetTeamSeasonID(id)
	return ptsc
}

// SetTeamSeason sets the "teamSeason" edge to the TeamSeason entity.
func (ptsc *PlayerTeamSeasonCreate) SetTeamSeason(t *TeamSeason) *PlayerTeamSeasonCreate {
	return ptsc.SetTeamSeasonID(t.ID)
}

// Mutation returns the PlayerTeamSeasonMutation object of the builder.
func (ptsc *PlayerTeamSeasonCreate) Mutation() *PlayerTeamSeasonMutation {
	return ptsc.mutation
}

// Save creates the PlayerTeamSeason in the database.
func (ptsc *PlayerTeamSeasonCreate) Save(ctx context.Context) (*PlayerTeamSeason, error) {
	return withHooks[*PlayerTeamSeason, PlayerTeamSeasonMutation](ctx, ptsc.sqlSave, ptsc.mutation, ptsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ptsc *PlayerTeamSeasonCreate) SaveX(ctx context.Context) *PlayerTeamSeason {
	v, err := ptsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptsc *PlayerTeamSeasonCreate) Exec(ctx context.Context) error {
	_, err := ptsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptsc *PlayerTeamSeasonCreate) ExecX(ctx context.Context) {
	if err := ptsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptsc *PlayerTeamSeasonCreate) check() error {
	if _, ok := ptsc.mutation.PlayerTeamSeasonID(); !ok {
		return &ValidationError{Name: "player_team_season_id", err: errors.New(`ent: missing required field "PlayerTeamSeason.player_team_season_id"`)}
	}
	if _, ok := ptsc.mutation.PlayerID(); !ok {
		return &ValidationError{Name: "player", err: errors.New(`ent: missing required edge "PlayerTeamSeason.player"`)}
	}
	if _, ok := ptsc.mutation.TeamSeasonID(); !ok {
		return &ValidationError{Name: "teamSeason", err: errors.New(`ent: missing required edge "PlayerTeamSeason.teamSeason"`)}
	}
	return nil
}

func (ptsc *PlayerTeamSeasonCreate) sqlSave(ctx context.Context) (*PlayerTeamSeason, error) {
	if err := ptsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ptsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ptsc.mutation.id = &_node.ID
	ptsc.mutation.done = true
	return _node, nil
}

func (ptsc *PlayerTeamSeasonCreate) createSpec() (*PlayerTeamSeason, *sqlgraph.CreateSpec) {
	var (
		_node = &PlayerTeamSeason{config: ptsc.config}
		_spec = sqlgraph.NewCreateSpec(playerteamseason.Table, sqlgraph.NewFieldSpec(playerteamseason.FieldID, field.TypeInt))
	)
	if value, ok := ptsc.mutation.PlayerTeamSeasonID(); ok {
		_spec.SetField(playerteamseason.FieldPlayerTeamSeasonID, field.TypeInt, value)
		_node.PlayerTeamSeasonID = value
	}
	if nodes := ptsc.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playerteamseason.PlayerTable,
			Columns: []string{playerteamseason.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_player_team_seasons = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ptsc.mutation.TeamSeasonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playerteamseason.TeamSeasonTable,
			Columns: []string{playerteamseason.TeamSeasonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamseason.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.team_season_player_team_seasons = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlayerTeamSeasonCreateBulk is the builder for creating many PlayerTeamSeason entities in bulk.
type PlayerTeamSeasonCreateBulk struct {
	config
	builders []*PlayerTeamSeasonCreate
}

// Save creates the PlayerTeamSeason entities in the database.
func (ptscb *PlayerTeamSeasonCreateBulk) Save(ctx context.Context) ([]*PlayerTeamSeason, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ptscb.builders))
	nodes := make([]*PlayerTeamSeason, len(ptscb.builders))
	mutators := make([]Mutator, len(ptscb.builders))
	for i := range ptscb.builders {
		func(i int, root context.Context) {
			builder := ptscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerTeamSeasonMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ptscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ptscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptscb *PlayerTeamSeasonCreateBulk) SaveX(ctx context.Context) []*PlayerTeamSeason {
	v, err := ptscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptscb *PlayerTeamSeasonCreateBulk) Exec(ctx context.Context) error {
	_, err := ptscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptscb *PlayerTeamSeasonCreateBulk) ExecX(ctx context.Context) {
	if err := ptscb.Exec(ctx); err != nil {
		panic(err)
	}
}
