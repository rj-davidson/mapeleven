// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSDefenseCreate is the builder for creating a PSDefense entity.
type PSDefenseCreate struct {
	config
	mutation *PSDefenseMutation
	hooks    []Hook
}

// SetTacklesTotal sets the "TacklesTotal" field.
func (pdc *PSDefenseCreate) SetTacklesTotal(i int) *PSDefenseCreate {
	pdc.mutation.SetTacklesTotal(i)
	return pdc
}

// SetBlocks sets the "Blocks" field.
func (pdc *PSDefenseCreate) SetBlocks(i int) *PSDefenseCreate {
	pdc.mutation.SetBlocks(i)
	return pdc
}

// SetNillableBlocks sets the "Blocks" field if the given value is not nil.
func (pdc *PSDefenseCreate) SetNillableBlocks(i *int) *PSDefenseCreate {
	if i != nil {
		pdc.SetBlocks(*i)
	}
	return pdc
}

// SetInterceptions sets the "Interceptions" field.
func (pdc *PSDefenseCreate) SetInterceptions(i int) *PSDefenseCreate {
	pdc.mutation.SetInterceptions(i)
	return pdc
}

// SetDuelsTotal sets the "DuelsTotal" field.
func (pdc *PSDefenseCreate) SetDuelsTotal(i int) *PSDefenseCreate {
	pdc.mutation.SetDuelsTotal(i)
	return pdc
}

// SetWonDuels sets the "WonDuels" field.
func (pdc *PSDefenseCreate) SetWonDuels(i int) *PSDefenseCreate {
	pdc.mutation.SetWonDuels(i)
	return pdc
}

// SetLastUpdated sets the "lastUpdated" field.
func (pdc *PSDefenseCreate) SetLastUpdated(t time.Time) *PSDefenseCreate {
	pdc.mutation.SetLastUpdated(t)
	return pdc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (pdc *PSDefenseCreate) SetNillableLastUpdated(t *time.Time) *PSDefenseCreate {
	if t != nil {
		pdc.SetLastUpdated(*t)
	}
	return pdc
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (pdc *PSDefenseCreate) SetPlayerStatsID(id int) *PSDefenseCreate {
	pdc.mutation.SetPlayerStatsID(id)
	return pdc
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (pdc *PSDefenseCreate) SetNillablePlayerStatsID(id *int) *PSDefenseCreate {
	if id != nil {
		pdc = pdc.SetPlayerStatsID(*id)
	}
	return pdc
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (pdc *PSDefenseCreate) SetPlayerStats(p *PlayerStats) *PSDefenseCreate {
	return pdc.SetPlayerStatsID(p.ID)
}

// Mutation returns the PSDefenseMutation object of the builder.
func (pdc *PSDefenseCreate) Mutation() *PSDefenseMutation {
	return pdc.mutation
}

// Save creates the PSDefense in the database.
func (pdc *PSDefenseCreate) Save(ctx context.Context) (*PSDefense, error) {
	pdc.defaults()
	return withHooks(ctx, pdc.sqlSave, pdc.mutation, pdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pdc *PSDefenseCreate) SaveX(ctx context.Context) *PSDefense {
	v, err := pdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdc *PSDefenseCreate) Exec(ctx context.Context) error {
	_, err := pdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdc *PSDefenseCreate) ExecX(ctx context.Context) {
	if err := pdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdc *PSDefenseCreate) defaults() {
	if _, ok := pdc.mutation.Blocks(); !ok {
		v := psdefense.DefaultBlocks
		pdc.mutation.SetBlocks(v)
	}
	if _, ok := pdc.mutation.LastUpdated(); !ok {
		v := psdefense.DefaultLastUpdated()
		pdc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdc *PSDefenseCreate) check() error {
	if _, ok := pdc.mutation.TacklesTotal(); !ok {
		return &ValidationError{Name: "TacklesTotal", err: errors.New(`ent: missing required field "PSDefense.TacklesTotal"`)}
	}
	if _, ok := pdc.mutation.Blocks(); !ok {
		return &ValidationError{Name: "Blocks", err: errors.New(`ent: missing required field "PSDefense.Blocks"`)}
	}
	if _, ok := pdc.mutation.Interceptions(); !ok {
		return &ValidationError{Name: "Interceptions", err: errors.New(`ent: missing required field "PSDefense.Interceptions"`)}
	}
	if _, ok := pdc.mutation.DuelsTotal(); !ok {
		return &ValidationError{Name: "DuelsTotal", err: errors.New(`ent: missing required field "PSDefense.DuelsTotal"`)}
	}
	if _, ok := pdc.mutation.WonDuels(); !ok {
		return &ValidationError{Name: "WonDuels", err: errors.New(`ent: missing required field "PSDefense.WonDuels"`)}
	}
	return nil
}

func (pdc *PSDefenseCreate) sqlSave(ctx context.Context) (*PSDefense, error) {
	if err := pdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pdc.mutation.id = &_node.ID
	pdc.mutation.done = true
	return _node, nil
}

func (pdc *PSDefenseCreate) createSpec() (*PSDefense, *sqlgraph.CreateSpec) {
	var (
		_node = &PSDefense{config: pdc.config}
		_spec = sqlgraph.NewCreateSpec(psdefense.Table, sqlgraph.NewFieldSpec(psdefense.FieldID, field.TypeInt))
	)
	if value, ok := pdc.mutation.TacklesTotal(); ok {
		_spec.SetField(psdefense.FieldTacklesTotal, field.TypeInt, value)
		_node.TacklesTotal = value
	}
	if value, ok := pdc.mutation.Blocks(); ok {
		_spec.SetField(psdefense.FieldBlocks, field.TypeInt, value)
		_node.Blocks = value
	}
	if value, ok := pdc.mutation.Interceptions(); ok {
		_spec.SetField(psdefense.FieldInterceptions, field.TypeInt, value)
		_node.Interceptions = value
	}
	if value, ok := pdc.mutation.DuelsTotal(); ok {
		_spec.SetField(psdefense.FieldDuelsTotal, field.TypeInt, value)
		_node.DuelsTotal = value
	}
	if value, ok := pdc.mutation.WonDuels(); ok {
		_spec.SetField(psdefense.FieldWonDuels, field.TypeInt, value)
		_node.WonDuels = value
	}
	if value, ok := pdc.mutation.LastUpdated(); ok {
		_spec.SetField(psdefense.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := pdc.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   psdefense.PlayerStatsTable,
			Columns: []string{psdefense.PlayerStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_stats_ps_defense = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PSDefenseCreateBulk is the builder for creating many PSDefense entities in bulk.
type PSDefenseCreateBulk struct {
	config
	err      error
	builders []*PSDefenseCreate
}

// Save creates the PSDefense entities in the database.
func (pdcb *PSDefenseCreateBulk) Save(ctx context.Context) ([]*PSDefense, error) {
	if pdcb.err != nil {
		return nil, pdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pdcb.builders))
	nodes := make([]*PSDefense, len(pdcb.builders))
	mutators := make([]Mutator, len(pdcb.builders))
	for i := range pdcb.builders {
		func(i int, root context.Context) {
			builder := pdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PSDefenseMutation)
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
					_, err = mutators[i+1].Mutate(root, pdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pdcb *PSDefenseCreateBulk) SaveX(ctx context.Context) []*PSDefense {
	v, err := pdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdcb *PSDefenseCreateBulk) Exec(ctx context.Context) error {
	_, err := pdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdcb *PSDefenseCreateBulk) ExecX(ctx context.Context) {
	if err := pdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
