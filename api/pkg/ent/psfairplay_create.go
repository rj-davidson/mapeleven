// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psfairplay"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSFairplayCreate is the builder for creating a PSFairplay entity.
type PSFairplayCreate struct {
	config
	mutation *PSFairplayMutation
	hooks    []Hook
}

// SetFoulsCommitted sets the "FoulsCommitted" field.
func (pfc *PSFairplayCreate) SetFoulsCommitted(i int) *PSFairplayCreate {
	pfc.mutation.SetFoulsCommitted(i)
	return pfc
}

// SetNillableFoulsCommitted sets the "FoulsCommitted" field if the given value is not nil.
func (pfc *PSFairplayCreate) SetNillableFoulsCommitted(i *int) *PSFairplayCreate {
	if i != nil {
		pfc.SetFoulsCommitted(*i)
	}
	return pfc
}

// SetYellow sets the "Yellow" field.
func (pfc *PSFairplayCreate) SetYellow(i int) *PSFairplayCreate {
	pfc.mutation.SetYellow(i)
	return pfc
}

// SetNillableYellow sets the "Yellow" field if the given value is not nil.
func (pfc *PSFairplayCreate) SetNillableYellow(i *int) *PSFairplayCreate {
	if i != nil {
		pfc.SetYellow(*i)
	}
	return pfc
}

// SetYellowRed sets the "YellowRed" field.
func (pfc *PSFairplayCreate) SetYellowRed(i int) *PSFairplayCreate {
	pfc.mutation.SetYellowRed(i)
	return pfc
}

// SetNillableYellowRed sets the "YellowRed" field if the given value is not nil.
func (pfc *PSFairplayCreate) SetNillableYellowRed(i *int) *PSFairplayCreate {
	if i != nil {
		pfc.SetYellowRed(*i)
	}
	return pfc
}

// SetRed sets the "Red" field.
func (pfc *PSFairplayCreate) SetRed(i int) *PSFairplayCreate {
	pfc.mutation.SetRed(i)
	return pfc
}

// SetNillableRed sets the "Red" field if the given value is not nil.
func (pfc *PSFairplayCreate) SetNillableRed(i *int) *PSFairplayCreate {
	if i != nil {
		pfc.SetRed(*i)
	}
	return pfc
}

// SetPenaltyConceded sets the "PenaltyConceded" field.
func (pfc *PSFairplayCreate) SetPenaltyConceded(i int) *PSFairplayCreate {
	pfc.mutation.SetPenaltyConceded(i)
	return pfc
}

// SetNillablePenaltyConceded sets the "PenaltyConceded" field if the given value is not nil.
func (pfc *PSFairplayCreate) SetNillablePenaltyConceded(i *int) *PSFairplayCreate {
	if i != nil {
		pfc.SetPenaltyConceded(*i)
	}
	return pfc
}

// SetLastUpdated sets the "lastUpdated" field.
func (pfc *PSFairplayCreate) SetLastUpdated(t time.Time) *PSFairplayCreate {
	pfc.mutation.SetLastUpdated(t)
	return pfc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (pfc *PSFairplayCreate) SetNillableLastUpdated(t *time.Time) *PSFairplayCreate {
	if t != nil {
		pfc.SetLastUpdated(*t)
	}
	return pfc
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (pfc *PSFairplayCreate) SetPlayerStatsID(id int) *PSFairplayCreate {
	pfc.mutation.SetPlayerStatsID(id)
	return pfc
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (pfc *PSFairplayCreate) SetNillablePlayerStatsID(id *int) *PSFairplayCreate {
	if id != nil {
		pfc = pfc.SetPlayerStatsID(*id)
	}
	return pfc
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (pfc *PSFairplayCreate) SetPlayerStats(p *PlayerStats) *PSFairplayCreate {
	return pfc.SetPlayerStatsID(p.ID)
}

// Mutation returns the PSFairplayMutation object of the builder.
func (pfc *PSFairplayCreate) Mutation() *PSFairplayMutation {
	return pfc.mutation
}

// Save creates the PSFairplay in the database.
func (pfc *PSFairplayCreate) Save(ctx context.Context) (*PSFairplay, error) {
	pfc.defaults()
	return withHooks(ctx, pfc.sqlSave, pfc.mutation, pfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pfc *PSFairplayCreate) SaveX(ctx context.Context) *PSFairplay {
	v, err := pfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pfc *PSFairplayCreate) Exec(ctx context.Context) error {
	_, err := pfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfc *PSFairplayCreate) ExecX(ctx context.Context) {
	if err := pfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pfc *PSFairplayCreate) defaults() {
	if _, ok := pfc.mutation.FoulsCommitted(); !ok {
		v := psfairplay.DefaultFoulsCommitted
		pfc.mutation.SetFoulsCommitted(v)
	}
	if _, ok := pfc.mutation.Yellow(); !ok {
		v := psfairplay.DefaultYellow
		pfc.mutation.SetYellow(v)
	}
	if _, ok := pfc.mutation.YellowRed(); !ok {
		v := psfairplay.DefaultYellowRed
		pfc.mutation.SetYellowRed(v)
	}
	if _, ok := pfc.mutation.Red(); !ok {
		v := psfairplay.DefaultRed
		pfc.mutation.SetRed(v)
	}
	if _, ok := pfc.mutation.PenaltyConceded(); !ok {
		v := psfairplay.DefaultPenaltyConceded
		pfc.mutation.SetPenaltyConceded(v)
	}
	if _, ok := pfc.mutation.LastUpdated(); !ok {
		v := psfairplay.DefaultLastUpdated()
		pfc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pfc *PSFairplayCreate) check() error {
	if _, ok := pfc.mutation.FoulsCommitted(); !ok {
		return &ValidationError{Name: "FoulsCommitted", err: errors.New(`ent: missing required field "PSFairplay.FoulsCommitted"`)}
	}
	if _, ok := pfc.mutation.Yellow(); !ok {
		return &ValidationError{Name: "Yellow", err: errors.New(`ent: missing required field "PSFairplay.Yellow"`)}
	}
	if _, ok := pfc.mutation.YellowRed(); !ok {
		return &ValidationError{Name: "YellowRed", err: errors.New(`ent: missing required field "PSFairplay.YellowRed"`)}
	}
	if _, ok := pfc.mutation.Red(); !ok {
		return &ValidationError{Name: "Red", err: errors.New(`ent: missing required field "PSFairplay.Red"`)}
	}
	if _, ok := pfc.mutation.PenaltyConceded(); !ok {
		return &ValidationError{Name: "PenaltyConceded", err: errors.New(`ent: missing required field "PSFairplay.PenaltyConceded"`)}
	}
	return nil
}

func (pfc *PSFairplayCreate) sqlSave(ctx context.Context) (*PSFairplay, error) {
	if err := pfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pfc.mutation.id = &_node.ID
	pfc.mutation.done = true
	return _node, nil
}

func (pfc *PSFairplayCreate) createSpec() (*PSFairplay, *sqlgraph.CreateSpec) {
	var (
		_node = &PSFairplay{config: pfc.config}
		_spec = sqlgraph.NewCreateSpec(psfairplay.Table, sqlgraph.NewFieldSpec(psfairplay.FieldID, field.TypeInt))
	)
	if value, ok := pfc.mutation.FoulsCommitted(); ok {
		_spec.SetField(psfairplay.FieldFoulsCommitted, field.TypeInt, value)
		_node.FoulsCommitted = value
	}
	if value, ok := pfc.mutation.Yellow(); ok {
		_spec.SetField(psfairplay.FieldYellow, field.TypeInt, value)
		_node.Yellow = value
	}
	if value, ok := pfc.mutation.YellowRed(); ok {
		_spec.SetField(psfairplay.FieldYellowRed, field.TypeInt, value)
		_node.YellowRed = value
	}
	if value, ok := pfc.mutation.Red(); ok {
		_spec.SetField(psfairplay.FieldRed, field.TypeInt, value)
		_node.Red = value
	}
	if value, ok := pfc.mutation.PenaltyConceded(); ok {
		_spec.SetField(psfairplay.FieldPenaltyConceded, field.TypeInt, value)
		_node.PenaltyConceded = value
	}
	if value, ok := pfc.mutation.LastUpdated(); ok {
		_spec.SetField(psfairplay.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := pfc.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   psfairplay.PlayerStatsTable,
			Columns: []string{psfairplay.PlayerStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_stats_ps_fairplay = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PSFairplayCreateBulk is the builder for creating many PSFairplay entities in bulk.
type PSFairplayCreateBulk struct {
	config
	err      error
	builders []*PSFairplayCreate
}

// Save creates the PSFairplay entities in the database.
func (pfcb *PSFairplayCreateBulk) Save(ctx context.Context) ([]*PSFairplay, error) {
	if pfcb.err != nil {
		return nil, pfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pfcb.builders))
	nodes := make([]*PSFairplay, len(pfcb.builders))
	mutators := make([]Mutator, len(pfcb.builders))
	for i := range pfcb.builders {
		func(i int, root context.Context) {
			builder := pfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PSFairplayMutation)
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
					_, err = mutators[i+1].Mutate(root, pfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pfcb *PSFairplayCreateBulk) SaveX(ctx context.Context) []*PSFairplay {
	v, err := pfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pfcb *PSFairplayCreateBulk) Exec(ctx context.Context) error {
	_, err := pfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfcb *PSFairplayCreateBulk) ExecX(ctx context.Context) {
	if err := pfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
