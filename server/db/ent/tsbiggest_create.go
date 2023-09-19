// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/team"
	"mapeleven/db/ent/tsbiggest"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TSBiggestCreate is the builder for creating a TSBiggest entity.
type TSBiggestCreate struct {
	config
	mutation *TSBiggestMutation
	hooks    []Hook
}

// SetStreakWins sets the "streakWins" field.
func (tbc *TSBiggestCreate) SetStreakWins(i int) *TSBiggestCreate {
	tbc.mutation.SetStreakWins(i)
	return tbc
}

// SetNillableStreakWins sets the "streakWins" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableStreakWins(i *int) *TSBiggestCreate {
	if i != nil {
		tbc.SetStreakWins(*i)
	}
	return tbc
}

// SetStreakLosses sets the "streakLosses" field.
func (tbc *TSBiggestCreate) SetStreakLosses(i int) *TSBiggestCreate {
	tbc.mutation.SetStreakLosses(i)
	return tbc
}

// SetNillableStreakLosses sets the "streakLosses" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableStreakLosses(i *int) *TSBiggestCreate {
	if i != nil {
		tbc.SetStreakLosses(*i)
	}
	return tbc
}

// SetStreakDraws sets the "streakDraws" field.
func (tbc *TSBiggestCreate) SetStreakDraws(i int) *TSBiggestCreate {
	tbc.mutation.SetStreakDraws(i)
	return tbc
}

// SetNillableStreakDraws sets the "streakDraws" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableStreakDraws(i *int) *TSBiggestCreate {
	if i != nil {
		tbc.SetStreakDraws(*i)
	}
	return tbc
}

// SetWinsHome sets the "winsHome" field.
func (tbc *TSBiggestCreate) SetWinsHome(s string) *TSBiggestCreate {
	tbc.mutation.SetWinsHome(s)
	return tbc
}

// SetNillableWinsHome sets the "winsHome" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableWinsHome(s *string) *TSBiggestCreate {
	if s != nil {
		tbc.SetWinsHome(*s)
	}
	return tbc
}

// SetWinsAway sets the "winsAway" field.
func (tbc *TSBiggestCreate) SetWinsAway(s string) *TSBiggestCreate {
	tbc.mutation.SetWinsAway(s)
	return tbc
}

// SetNillableWinsAway sets the "winsAway" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableWinsAway(s *string) *TSBiggestCreate {
	if s != nil {
		tbc.SetWinsAway(*s)
	}
	return tbc
}

// SetLossesHome sets the "lossesHome" field.
func (tbc *TSBiggestCreate) SetLossesHome(s string) *TSBiggestCreate {
	tbc.mutation.SetLossesHome(s)
	return tbc
}

// SetNillableLossesHome sets the "lossesHome" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableLossesHome(s *string) *TSBiggestCreate {
	if s != nil {
		tbc.SetLossesHome(*s)
	}
	return tbc
}

// SetLossesAway sets the "lossesAway" field.
func (tbc *TSBiggestCreate) SetLossesAway(s string) *TSBiggestCreate {
	tbc.mutation.SetLossesAway(s)
	return tbc
}

// SetNillableLossesAway sets the "lossesAway" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableLossesAway(s *string) *TSBiggestCreate {
	if s != nil {
		tbc.SetLossesAway(*s)
	}
	return tbc
}

// SetGoalsForHome sets the "goalsForHome" field.
func (tbc *TSBiggestCreate) SetGoalsForHome(i int) *TSBiggestCreate {
	tbc.mutation.SetGoalsForHome(i)
	return tbc
}

// SetNillableGoalsForHome sets the "goalsForHome" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableGoalsForHome(i *int) *TSBiggestCreate {
	if i != nil {
		tbc.SetGoalsForHome(*i)
	}
	return tbc
}

// SetGoalsForAway sets the "goalsForAway" field.
func (tbc *TSBiggestCreate) SetGoalsForAway(i int) *TSBiggestCreate {
	tbc.mutation.SetGoalsForAway(i)
	return tbc
}

// SetNillableGoalsForAway sets the "goalsForAway" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableGoalsForAway(i *int) *TSBiggestCreate {
	if i != nil {
		tbc.SetGoalsForAway(*i)
	}
	return tbc
}

// SetGoalsAgainstHome sets the "goalsAgainstHome" field.
func (tbc *TSBiggestCreate) SetGoalsAgainstHome(i int) *TSBiggestCreate {
	tbc.mutation.SetGoalsAgainstHome(i)
	return tbc
}

// SetNillableGoalsAgainstHome sets the "goalsAgainstHome" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableGoalsAgainstHome(i *int) *TSBiggestCreate {
	if i != nil {
		tbc.SetGoalsAgainstHome(*i)
	}
	return tbc
}

// SetGoalsAgainstAway sets the "goalsAgainstAway" field.
func (tbc *TSBiggestCreate) SetGoalsAgainstAway(i int) *TSBiggestCreate {
	tbc.mutation.SetGoalsAgainstAway(i)
	return tbc
}

// SetNillableGoalsAgainstAway sets the "goalsAgainstAway" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableGoalsAgainstAway(i *int) *TSBiggestCreate {
	if i != nil {
		tbc.SetGoalsAgainstAway(*i)
	}
	return tbc
}

// SetLastUpdated sets the "lastUpdated" field.
func (tbc *TSBiggestCreate) SetLastUpdated(t time.Time) *TSBiggestCreate {
	tbc.mutation.SetLastUpdated(t)
	return tbc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (tbc *TSBiggestCreate) SetNillableLastUpdated(t *time.Time) *TSBiggestCreate {
	if t != nil {
		tbc.SetLastUpdated(*t)
	}
	return tbc
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (tbc *TSBiggestCreate) SetTeamID(id int) *TSBiggestCreate {
	tbc.mutation.SetTeamID(id)
	return tbc
}

// SetTeam sets the "team" edge to the Team entity.
func (tbc *TSBiggestCreate) SetTeam(t *Team) *TSBiggestCreate {
	return tbc.SetTeamID(t.ID)
}

// Mutation returns the TSBiggestMutation object of the builder.
func (tbc *TSBiggestCreate) Mutation() *TSBiggestMutation {
	return tbc.mutation
}

// Save creates the TSBiggest in the database.
func (tbc *TSBiggestCreate) Save(ctx context.Context) (*TSBiggest, error) {
	tbc.defaults()
	return withHooks[*TSBiggest, TSBiggestMutation](ctx, tbc.sqlSave, tbc.mutation, tbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tbc *TSBiggestCreate) SaveX(ctx context.Context) *TSBiggest {
	v, err := tbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tbc *TSBiggestCreate) Exec(ctx context.Context) error {
	_, err := tbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tbc *TSBiggestCreate) ExecX(ctx context.Context) {
	if err := tbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tbc *TSBiggestCreate) defaults() {
	if _, ok := tbc.mutation.LastUpdated(); !ok {
		v := tsbiggest.DefaultLastUpdated()
		tbc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tbc *TSBiggestCreate) check() error {
	if _, ok := tbc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`ent: missing required edge "TSBiggest.team"`)}
	}
	return nil
}

func (tbc *TSBiggestCreate) sqlSave(ctx context.Context) (*TSBiggest, error) {
	if err := tbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tbc.mutation.id = &_node.ID
	tbc.mutation.done = true
	return _node, nil
}

func (tbc *TSBiggestCreate) createSpec() (*TSBiggest, *sqlgraph.CreateSpec) {
	var (
		_node = &TSBiggest{config: tbc.config}
		_spec = sqlgraph.NewCreateSpec(tsbiggest.Table, sqlgraph.NewFieldSpec(tsbiggest.FieldID, field.TypeInt))
	)
	if value, ok := tbc.mutation.StreakWins(); ok {
		_spec.SetField(tsbiggest.FieldStreakWins, field.TypeInt, value)
		_node.StreakWins = value
	}
	if value, ok := tbc.mutation.StreakLosses(); ok {
		_spec.SetField(tsbiggest.FieldStreakLosses, field.TypeInt, value)
		_node.StreakLosses = value
	}
	if value, ok := tbc.mutation.StreakDraws(); ok {
		_spec.SetField(tsbiggest.FieldStreakDraws, field.TypeInt, value)
		_node.StreakDraws = value
	}
	if value, ok := tbc.mutation.WinsHome(); ok {
		_spec.SetField(tsbiggest.FieldWinsHome, field.TypeString, value)
		_node.WinsHome = value
	}
	if value, ok := tbc.mutation.WinsAway(); ok {
		_spec.SetField(tsbiggest.FieldWinsAway, field.TypeString, value)
		_node.WinsAway = value
	}
	if value, ok := tbc.mutation.LossesHome(); ok {
		_spec.SetField(tsbiggest.FieldLossesHome, field.TypeString, value)
		_node.LossesHome = value
	}
	if value, ok := tbc.mutation.LossesAway(); ok {
		_spec.SetField(tsbiggest.FieldLossesAway, field.TypeString, value)
		_node.LossesAway = value
	}
	if value, ok := tbc.mutation.GoalsForHome(); ok {
		_spec.SetField(tsbiggest.FieldGoalsForHome, field.TypeInt, value)
		_node.GoalsForHome = value
	}
	if value, ok := tbc.mutation.GoalsForAway(); ok {
		_spec.SetField(tsbiggest.FieldGoalsForAway, field.TypeInt, value)
		_node.GoalsForAway = value
	}
	if value, ok := tbc.mutation.GoalsAgainstHome(); ok {
		_spec.SetField(tsbiggest.FieldGoalsAgainstHome, field.TypeInt, value)
		_node.GoalsAgainstHome = value
	}
	if value, ok := tbc.mutation.GoalsAgainstAway(); ok {
		_spec.SetField(tsbiggest.FieldGoalsAgainstAway, field.TypeInt, value)
		_node.GoalsAgainstAway = value
	}
	if value, ok := tbc.mutation.LastUpdated(); ok {
		_spec.SetField(tsbiggest.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := tbc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   tsbiggest.TeamTable,
			Columns: []string{tsbiggest.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.team_biggest_stats = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TSBiggestCreateBulk is the builder for creating many TSBiggest entities in bulk.
type TSBiggestCreateBulk struct {
	config
	builders []*TSBiggestCreate
}

// Save creates the TSBiggest entities in the database.
func (tbcb *TSBiggestCreateBulk) Save(ctx context.Context) ([]*TSBiggest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tbcb.builders))
	nodes := make([]*TSBiggest, len(tbcb.builders))
	mutators := make([]Mutator, len(tbcb.builders))
	for i := range tbcb.builders {
		func(i int, root context.Context) {
			builder := tbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TSBiggestMutation)
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
					_, err = mutators[i+1].Mutate(root, tbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tbcb *TSBiggestCreateBulk) SaveX(ctx context.Context) []*TSBiggest {
	v, err := tbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tbcb *TSBiggestCreateBulk) Exec(ctx context.Context) error {
	_, err := tbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tbcb *TSBiggestCreateBulk) ExecX(ctx context.Context) {
	if err := tbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
