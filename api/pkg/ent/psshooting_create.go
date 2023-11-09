// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psshooting"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSShootingCreate is the builder for creating a PSShooting entity.
type PSShootingCreate struct {
	config
	mutation *PSShootingMutation
	hooks    []Hook
}

// SetGoals sets the "Goals" field.
func (psc *PSShootingCreate) SetGoals(i int) *PSShootingCreate {
	psc.mutation.SetGoals(i)
	return psc
}

// SetConceded sets the "Conceded" field.
func (psc *PSShootingCreate) SetConceded(i int) *PSShootingCreate {
	psc.mutation.SetConceded(i)
	return psc
}

// SetAssists sets the "Assists" field.
func (psc *PSShootingCreate) SetAssists(i int) *PSShootingCreate {
	psc.mutation.SetAssists(i)
	return psc
}

// SetSaves sets the "Saves" field.
func (psc *PSShootingCreate) SetSaves(i int) *PSShootingCreate {
	psc.mutation.SetSaves(i)
	return psc
}

// SetNillableSaves sets the "Saves" field if the given value is not nil.
func (psc *PSShootingCreate) SetNillableSaves(i *int) *PSShootingCreate {
	if i != nil {
		psc.SetSaves(*i)
	}
	return psc
}

// SetShots sets the "Shots" field.
func (psc *PSShootingCreate) SetShots(i int) *PSShootingCreate {
	psc.mutation.SetShots(i)
	return psc
}

// SetOnTarget sets the "OnTarget" field.
func (psc *PSShootingCreate) SetOnTarget(i int) *PSShootingCreate {
	psc.mutation.SetOnTarget(i)
	return psc
}

// SetLastUpdated sets the "lastUpdated" field.
func (psc *PSShootingCreate) SetLastUpdated(t time.Time) *PSShootingCreate {
	psc.mutation.SetLastUpdated(t)
	return psc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (psc *PSShootingCreate) SetNillableLastUpdated(t *time.Time) *PSShootingCreate {
	if t != nil {
		psc.SetLastUpdated(*t)
	}
	return psc
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (psc *PSShootingCreate) SetPlayerStatsID(id int) *PSShootingCreate {
	psc.mutation.SetPlayerStatsID(id)
	return psc
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (psc *PSShootingCreate) SetNillablePlayerStatsID(id *int) *PSShootingCreate {
	if id != nil {
		psc = psc.SetPlayerStatsID(*id)
	}
	return psc
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (psc *PSShootingCreate) SetPlayerStats(p *PlayerStats) *PSShootingCreate {
	return psc.SetPlayerStatsID(p.ID)
}

// Mutation returns the PSShootingMutation object of the builder.
func (psc *PSShootingCreate) Mutation() *PSShootingMutation {
	return psc.mutation
}

// Save creates the PSShooting in the database.
func (psc *PSShootingCreate) Save(ctx context.Context) (*PSShooting, error) {
	psc.defaults()
	return withHooks(ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PSShootingCreate) SaveX(ctx context.Context) *PSShooting {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *PSShootingCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *PSShootingCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psc *PSShootingCreate) defaults() {
	if _, ok := psc.mutation.Saves(); !ok {
		v := psshooting.DefaultSaves
		psc.mutation.SetSaves(v)
	}
	if _, ok := psc.mutation.LastUpdated(); !ok {
		v := psshooting.DefaultLastUpdated()
		psc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PSShootingCreate) check() error {
	if _, ok := psc.mutation.Goals(); !ok {
		return &ValidationError{Name: "Goals", err: errors.New(`ent: missing required field "PSShooting.Goals"`)}
	}
	if _, ok := psc.mutation.Conceded(); !ok {
		return &ValidationError{Name: "Conceded", err: errors.New(`ent: missing required field "PSShooting.Conceded"`)}
	}
	if _, ok := psc.mutation.Assists(); !ok {
		return &ValidationError{Name: "Assists", err: errors.New(`ent: missing required field "PSShooting.Assists"`)}
	}
	if _, ok := psc.mutation.Saves(); !ok {
		return &ValidationError{Name: "Saves", err: errors.New(`ent: missing required field "PSShooting.Saves"`)}
	}
	if _, ok := psc.mutation.Shots(); !ok {
		return &ValidationError{Name: "Shots", err: errors.New(`ent: missing required field "PSShooting.Shots"`)}
	}
	if _, ok := psc.mutation.OnTarget(); !ok {
		return &ValidationError{Name: "OnTarget", err: errors.New(`ent: missing required field "PSShooting.OnTarget"`)}
	}
	return nil
}

func (psc *PSShootingCreate) sqlSave(ctx context.Context) (*PSShooting, error) {
	if err := psc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	psc.mutation.id = &_node.ID
	psc.mutation.done = true
	return _node, nil
}

func (psc *PSShootingCreate) createSpec() (*PSShooting, *sqlgraph.CreateSpec) {
	var (
		_node = &PSShooting{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(psshooting.Table, sqlgraph.NewFieldSpec(psshooting.FieldID, field.TypeInt))
	)
	if value, ok := psc.mutation.Goals(); ok {
		_spec.SetField(psshooting.FieldGoals, field.TypeInt, value)
		_node.Goals = value
	}
	if value, ok := psc.mutation.Conceded(); ok {
		_spec.SetField(psshooting.FieldConceded, field.TypeInt, value)
		_node.Conceded = value
	}
	if value, ok := psc.mutation.Assists(); ok {
		_spec.SetField(psshooting.FieldAssists, field.TypeInt, value)
		_node.Assists = value
	}
	if value, ok := psc.mutation.Saves(); ok {
		_spec.SetField(psshooting.FieldSaves, field.TypeInt, value)
		_node.Saves = value
	}
	if value, ok := psc.mutation.Shots(); ok {
		_spec.SetField(psshooting.FieldShots, field.TypeInt, value)
		_node.Shots = value
	}
	if value, ok := psc.mutation.OnTarget(); ok {
		_spec.SetField(psshooting.FieldOnTarget, field.TypeInt, value)
		_node.OnTarget = value
	}
	if value, ok := psc.mutation.LastUpdated(); ok {
		_spec.SetField(psshooting.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := psc.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   psshooting.PlayerStatsTable,
			Columns: []string{psshooting.PlayerStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_stats_ps_shooting = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PSShootingCreateBulk is the builder for creating many PSShooting entities in bulk.
type PSShootingCreateBulk struct {
	config
	err      error
	builders []*PSShootingCreate
}

// Save creates the PSShooting entities in the database.
func (pscb *PSShootingCreateBulk) Save(ctx context.Context) ([]*PSShooting, error) {
	if pscb.err != nil {
		return nil, pscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PSShooting, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PSShootingMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *PSShootingCreateBulk) SaveX(ctx context.Context) []*PSShooting {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *PSShootingCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *PSShootingCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}
