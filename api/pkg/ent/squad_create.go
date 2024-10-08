// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/squad"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SquadCreate is the builder for creating a Squad entity.
type SquadCreate struct {
	config
	mutation *SquadMutation
	hooks    []Hook
}

// SetPosition sets the "position" field.
func (sc *SquadCreate) SetPosition(s string) *SquadCreate {
	sc.mutation.SetPosition(s)
	return sc
}

// SetNumber sets the "number" field.
func (sc *SquadCreate) SetNumber(i int) *SquadCreate {
	sc.mutation.SetNumber(i)
	return sc
}

// SetLastUpdated sets the "lastUpdated" field.
func (sc *SquadCreate) SetLastUpdated(t time.Time) *SquadCreate {
	sc.mutation.SetLastUpdated(t)
	return sc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (sc *SquadCreate) SetNillableLastUpdated(t *time.Time) *SquadCreate {
	if t != nil {
		sc.SetLastUpdated(*t)
	}
	return sc
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (sc *SquadCreate) SetPlayerID(id int) *SquadCreate {
	sc.mutation.SetPlayerID(id)
	return sc
}

// SetNillablePlayerID sets the "player" edge to the Player entity by ID if the given value is not nil.
func (sc *SquadCreate) SetNillablePlayerID(id *int) *SquadCreate {
	if id != nil {
		sc = sc.SetPlayerID(*id)
	}
	return sc
}

// SetPlayer sets the "player" edge to the Player entity.
func (sc *SquadCreate) SetPlayer(p *Player) *SquadCreate {
	return sc.SetPlayerID(p.ID)
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (sc *SquadCreate) SetTeamID(id int) *SquadCreate {
	sc.mutation.SetTeamID(id)
	return sc
}

// SetNillableTeamID sets the "team" edge to the Team entity by ID if the given value is not nil.
func (sc *SquadCreate) SetNillableTeamID(id *int) *SquadCreate {
	if id != nil {
		sc = sc.SetTeamID(*id)
	}
	return sc
}

// SetTeam sets the "team" edge to the Team entity.
func (sc *SquadCreate) SetTeam(t *Team) *SquadCreate {
	return sc.SetTeamID(t.ID)
}

// SetSeasonID sets the "season" edge to the Season entity by ID.
func (sc *SquadCreate) SetSeasonID(id int) *SquadCreate {
	sc.mutation.SetSeasonID(id)
	return sc
}

// SetNillableSeasonID sets the "season" edge to the Season entity by ID if the given value is not nil.
func (sc *SquadCreate) SetNillableSeasonID(id *int) *SquadCreate {
	if id != nil {
		sc = sc.SetSeasonID(*id)
	}
	return sc
}

// SetSeason sets the "season" edge to the Season entity.
func (sc *SquadCreate) SetSeason(s *Season) *SquadCreate {
	return sc.SetSeasonID(s.ID)
}

// Mutation returns the SquadMutation object of the builder.
func (sc *SquadCreate) Mutation() *SquadMutation {
	return sc.mutation
}

// Save creates the Squad in the database.
func (sc *SquadCreate) Save(ctx context.Context) (*Squad, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SquadCreate) SaveX(ctx context.Context) *Squad {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SquadCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SquadCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SquadCreate) defaults() {
	if _, ok := sc.mutation.LastUpdated(); !ok {
		v := squad.DefaultLastUpdated()
		sc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SquadCreate) check() error {
	if _, ok := sc.mutation.Position(); !ok {
		return &ValidationError{Name: "position", err: errors.New(`ent: missing required field "Squad.position"`)}
	}
	if _, ok := sc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Squad.number"`)}
	}
	return nil
}

func (sc *SquadCreate) sqlSave(ctx context.Context) (*Squad, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SquadCreate) createSpec() (*Squad, *sqlgraph.CreateSpec) {
	var (
		_node = &Squad{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(squad.Table, sqlgraph.NewFieldSpec(squad.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Position(); ok {
		_spec.SetField(squad.FieldPosition, field.TypeString, value)
		_node.Position = value
	}
	if value, ok := sc.mutation.Number(); ok {
		_spec.SetField(squad.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if value, ok := sc.mutation.LastUpdated(); ok {
		_spec.SetField(squad.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := sc.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   squad.PlayerTable,
			Columns: []string{squad.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_squad = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   squad.TeamTable,
			Columns: []string{squad.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.team_squad = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.SeasonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   squad.SeasonTable,
			Columns: []string{squad.SeasonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.season_squad = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SquadCreateBulk is the builder for creating many Squad entities in bulk.
type SquadCreateBulk struct {
	config
	err      error
	builders []*SquadCreate
}

// Save creates the Squad entities in the database.
func (scb *SquadCreateBulk) Save(ctx context.Context) ([]*Squad, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Squad, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SquadMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SquadCreateBulk) SaveX(ctx context.Context) []*Squad {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SquadCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SquadCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
