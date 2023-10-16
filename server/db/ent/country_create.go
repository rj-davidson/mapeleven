// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/club"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/league"
	"mapeleven/db/ent/player"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CountryCreate is the builder for creating a Country entity.
type CountryCreate struct {
	config
	mutation *CountryMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (cc *CountryCreate) SetCode(s string) *CountryCreate {
	cc.mutation.SetCode(s)
	return cc
}

// SetName sets the "name" field.
func (cc *CountryCreate) SetName(s string) *CountryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetFlag sets the "flag" field.
func (cc *CountryCreate) SetFlag(s string) *CountryCreate {
	cc.mutation.SetFlag(s)
	return cc
}

// SetLastUpdated sets the "lastUpdated" field.
func (cc *CountryCreate) SetLastUpdated(t time.Time) *CountryCreate {
	cc.mutation.SetLastUpdated(t)
	return cc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (cc *CountryCreate) SetNillableLastUpdated(t *time.Time) *CountryCreate {
	if t != nil {
		cc.SetLastUpdated(*t)
	}
	return cc
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (cc *CountryCreate) AddPlayerIDs(ids ...int) *CountryCreate {
	cc.mutation.AddPlayerIDs(ids...)
	return cc
}

// AddPlayers adds the "players" edges to the Player entity.
func (cc *CountryCreate) AddPlayers(p ...*Player) *CountryCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPlayerIDs(ids...)
}

// AddLeagueIDs adds the "leagues" edge to the League entity by IDs.
func (cc *CountryCreate) AddLeagueIDs(ids ...int) *CountryCreate {
	cc.mutation.AddLeagueIDs(ids...)
	return cc
}

// AddLeagues adds the "leagues" edges to the League entity.
func (cc *CountryCreate) AddLeagues(l ...*League) *CountryCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cc.AddLeagueIDs(ids...)
}

// AddClubIDs adds the "clubs" edge to the Club entity by IDs.
func (cc *CountryCreate) AddClubIDs(ids ...int) *CountryCreate {
	cc.mutation.AddClubIDs(ids...)
	return cc
}

// AddClubs adds the "clubs" edges to the Club entity.
func (cc *CountryCreate) AddClubs(c ...*Club) *CountryCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddClubIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cc *CountryCreate) Mutation() *CountryMutation {
	return cc.mutation
}

// Save creates the Country in the database.
func (cc *CountryCreate) Save(ctx context.Context) (*Country, error) {
	cc.defaults()
	return withHooks[*Country, CountryMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CountryCreate) SaveX(ctx context.Context) *Country {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CountryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CountryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CountryCreate) defaults() {
	if _, ok := cc.mutation.LastUpdated(); !ok {
		v := country.DefaultLastUpdated()
		cc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CountryCreate) check() error {
	if _, ok := cc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Country.code"`)}
	}
	if v, ok := cc.mutation.Code(); ok {
		if err := country.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Country.code": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Country.name"`)}
	}
	if _, ok := cc.mutation.Flag(); !ok {
		return &ValidationError{Name: "flag", err: errors.New(`ent: missing required field "Country.flag"`)}
	}
	return nil
}

func (cc *CountryCreate) sqlSave(ctx context.Context) (*Country, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CountryCreate) createSpec() (*Country, *sqlgraph.CreateSpec) {
	var (
		_node = &Country{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(country.Table, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Code(); ok {
		_spec.SetField(country.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Flag(); ok {
		_spec.SetField(country.FieldFlag, field.TypeString, value)
		_node.Flag = value
	}
	if value, ok := cc.mutation.LastUpdated(); ok {
		_spec.SetField(country.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := cc.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.PlayersTable,
			Columns: []string{country.PlayersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.LeaguesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LeaguesTable,
			Columns: []string{country.LeaguesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ClubsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.ClubsTable,
			Columns: []string{country.ClubsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(club.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CountryCreateBulk is the builder for creating many Country entities in bulk.
type CountryCreateBulk struct {
	config
	builders []*CountryCreate
}

// Save creates the Country entities in the database.
func (ccb *CountryCreateBulk) Save(ctx context.Context) ([]*Country, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Country, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CountryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CountryCreateBulk) SaveX(ctx context.Context) []*Country {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CountryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CountryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
