// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/fixture"
	"mapeleven/db/ent/fixtureevents"
	"mapeleven/db/ent/fixturelineups"
	"mapeleven/db/ent/season"
	"mapeleven/db/ent/team"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FixtureCreate is the builder for creating a Fixture entity.
type FixtureCreate struct {
	config
	mutation *FixtureMutation
	hooks    []Hook
}

// SetSlug sets the "slug" field.
func (fc *FixtureCreate) SetSlug(s string) *FixtureCreate {
	fc.mutation.SetSlug(s)
	return fc
}

// SetApiFootballId sets the "apiFootballId" field.
func (fc *FixtureCreate) SetApiFootballId(i int) *FixtureCreate {
	fc.mutation.SetApiFootballId(i)
	return fc
}

// SetReferee sets the "referee" field.
func (fc *FixtureCreate) SetReferee(s string) *FixtureCreate {
	fc.mutation.SetReferee(s)
	return fc
}

// SetNillableReferee sets the "referee" field if the given value is not nil.
func (fc *FixtureCreate) SetNillableReferee(s *string) *FixtureCreate {
	if s != nil {
		fc.SetReferee(*s)
	}
	return fc
}

// SetTimezone sets the "timezone" field.
func (fc *FixtureCreate) SetTimezone(s string) *FixtureCreate {
	fc.mutation.SetTimezone(s)
	return fc
}

// SetNillableTimezone sets the "timezone" field if the given value is not nil.
func (fc *FixtureCreate) SetNillableTimezone(s *string) *FixtureCreate {
	if s != nil {
		fc.SetTimezone(*s)
	}
	return fc
}

// SetDate sets the "date" field.
func (fc *FixtureCreate) SetDate(t time.Time) *FixtureCreate {
	fc.mutation.SetDate(t)
	return fc
}

// SetElapsed sets the "elapsed" field.
func (fc *FixtureCreate) SetElapsed(i int) *FixtureCreate {
	fc.mutation.SetElapsed(i)
	return fc
}

// SetNillableElapsed sets the "elapsed" field if the given value is not nil.
func (fc *FixtureCreate) SetNillableElapsed(i *int) *FixtureCreate {
	if i != nil {
		fc.SetElapsed(*i)
	}
	return fc
}

// SetRound sets the "round" field.
func (fc *FixtureCreate) SetRound(i int) *FixtureCreate {
	fc.mutation.SetRound(i)
	return fc
}

// SetNillableRound sets the "round" field if the given value is not nil.
func (fc *FixtureCreate) SetNillableRound(i *int) *FixtureCreate {
	if i != nil {
		fc.SetRound(*i)
	}
	return fc
}

// SetStatus sets the "status" field.
func (fc *FixtureCreate) SetStatus(s string) *FixtureCreate {
	fc.mutation.SetStatus(s)
	return fc
}

// SetHomeTeamScore sets the "homeTeamScore" field.
func (fc *FixtureCreate) SetHomeTeamScore(i int) *FixtureCreate {
	fc.mutation.SetHomeTeamScore(i)
	return fc
}

// SetNillableHomeTeamScore sets the "homeTeamScore" field if the given value is not nil.
func (fc *FixtureCreate) SetNillableHomeTeamScore(i *int) *FixtureCreate {
	if i != nil {
		fc.SetHomeTeamScore(*i)
	}
	return fc
}

// SetAwayTeamScore sets the "awayTeamScore" field.
func (fc *FixtureCreate) SetAwayTeamScore(i int) *FixtureCreate {
	fc.mutation.SetAwayTeamScore(i)
	return fc
}

// SetNillableAwayTeamScore sets the "awayTeamScore" field if the given value is not nil.
func (fc *FixtureCreate) SetNillableAwayTeamScore(i *int) *FixtureCreate {
	if i != nil {
		fc.SetAwayTeamScore(*i)
	}
	return fc
}

// SetLastUpdated sets the "lastUpdated" field.
func (fc *FixtureCreate) SetLastUpdated(t time.Time) *FixtureCreate {
	fc.mutation.SetLastUpdated(t)
	return fc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (fc *FixtureCreate) SetNillableLastUpdated(t *time.Time) *FixtureCreate {
	if t != nil {
		fc.SetLastUpdated(*t)
	}
	return fc
}

// SetHomeTeamID sets the "homeTeam" edge to the Team entity by ID.
func (fc *FixtureCreate) SetHomeTeamID(id int) *FixtureCreate {
	fc.mutation.SetHomeTeamID(id)
	return fc
}

// SetHomeTeam sets the "homeTeam" edge to the Team entity.
func (fc *FixtureCreate) SetHomeTeam(t *Team) *FixtureCreate {
	return fc.SetHomeTeamID(t.ID)
}

// SetAwayTeamID sets the "awayTeam" edge to the Team entity by ID.
func (fc *FixtureCreate) SetAwayTeamID(id int) *FixtureCreate {
	fc.mutation.SetAwayTeamID(id)
	return fc
}

// SetAwayTeam sets the "awayTeam" edge to the Team entity.
func (fc *FixtureCreate) SetAwayTeam(t *Team) *FixtureCreate {
	return fc.SetAwayTeamID(t.ID)
}

// SetSeasonID sets the "season" edge to the Season entity by ID.
func (fc *FixtureCreate) SetSeasonID(id int) *FixtureCreate {
	fc.mutation.SetSeasonID(id)
	return fc
}

// SetSeason sets the "season" edge to the Season entity.
func (fc *FixtureCreate) SetSeason(s *Season) *FixtureCreate {
	return fc.SetSeasonID(s.ID)
}

// AddLineupIDs adds the "lineups" edge to the FixtureLineups entity by IDs.
func (fc *FixtureCreate) AddLineupIDs(ids ...int) *FixtureCreate {
	fc.mutation.AddLineupIDs(ids...)
	return fc
}

// AddLineups adds the "lineups" edges to the FixtureLineups entity.
func (fc *FixtureCreate) AddLineups(f ...*FixtureLineups) *FixtureCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddLineupIDs(ids...)
}

// AddFixtureEventIDs adds the "fixtureEvents" edge to the FixtureEvents entity by IDs.
func (fc *FixtureCreate) AddFixtureEventIDs(ids ...int) *FixtureCreate {
	fc.mutation.AddFixtureEventIDs(ids...)
	return fc
}

// AddFixtureEvents adds the "fixtureEvents" edges to the FixtureEvents entity.
func (fc *FixtureCreate) AddFixtureEvents(f ...*FixtureEvents) *FixtureCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFixtureEventIDs(ids...)
}

// Mutation returns the FixtureMutation object of the builder.
func (fc *FixtureCreate) Mutation() *FixtureMutation {
	return fc.mutation
}

// Save creates the Fixture in the database.
func (fc *FixtureCreate) Save(ctx context.Context) (*Fixture, error) {
	fc.defaults()
	return withHooks[*Fixture, FixtureMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FixtureCreate) SaveX(ctx context.Context) *Fixture {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FixtureCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FixtureCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FixtureCreate) defaults() {
	if _, ok := fc.mutation.LastUpdated(); !ok {
		v := fixture.DefaultLastUpdated()
		fc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FixtureCreate) check() error {
	if _, ok := fc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Fixture.slug"`)}
	}
	if _, ok := fc.mutation.ApiFootballId(); !ok {
		return &ValidationError{Name: "apiFootballId", err: errors.New(`ent: missing required field "Fixture.apiFootballId"`)}
	}
	if _, ok := fc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Fixture.date"`)}
	}
	if _, ok := fc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Fixture.status"`)}
	}
	if _, ok := fc.mutation.HomeTeamID(); !ok {
		return &ValidationError{Name: "homeTeam", err: errors.New(`ent: missing required edge "Fixture.homeTeam"`)}
	}
	if _, ok := fc.mutation.AwayTeamID(); !ok {
		return &ValidationError{Name: "awayTeam", err: errors.New(`ent: missing required edge "Fixture.awayTeam"`)}
	}
	if _, ok := fc.mutation.SeasonID(); !ok {
		return &ValidationError{Name: "season", err: errors.New(`ent: missing required edge "Fixture.season"`)}
	}
	return nil
}

func (fc *FixtureCreate) sqlSave(ctx context.Context) (*Fixture, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FixtureCreate) createSpec() (*Fixture, *sqlgraph.CreateSpec) {
	var (
		_node = &Fixture{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(fixture.Table, sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt))
	)
	if value, ok := fc.mutation.Slug(); ok {
		_spec.SetField(fixture.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := fc.mutation.ApiFootballId(); ok {
		_spec.SetField(fixture.FieldApiFootballId, field.TypeInt, value)
		_node.ApiFootballId = value
	}
	if value, ok := fc.mutation.Referee(); ok {
		_spec.SetField(fixture.FieldReferee, field.TypeString, value)
		_node.Referee = value
	}
	if value, ok := fc.mutation.Timezone(); ok {
		_spec.SetField(fixture.FieldTimezone, field.TypeString, value)
		_node.Timezone = value
	}
	if value, ok := fc.mutation.Date(); ok {
		_spec.SetField(fixture.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := fc.mutation.Elapsed(); ok {
		_spec.SetField(fixture.FieldElapsed, field.TypeInt, value)
		_node.Elapsed = value
	}
	if value, ok := fc.mutation.Round(); ok {
		_spec.SetField(fixture.FieldRound, field.TypeInt, value)
		_node.Round = value
	}
	if value, ok := fc.mutation.Status(); ok {
		_spec.SetField(fixture.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := fc.mutation.HomeTeamScore(); ok {
		_spec.SetField(fixture.FieldHomeTeamScore, field.TypeInt, value)
		_node.HomeTeamScore = value
	}
	if value, ok := fc.mutation.AwayTeamScore(); ok {
		_spec.SetField(fixture.FieldAwayTeamScore, field.TypeInt, value)
		_node.AwayTeamScore = value
	}
	if value, ok := fc.mutation.LastUpdated(); ok {
		_spec.SetField(fixture.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := fc.mutation.HomeTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixture.HomeTeamTable,
			Columns: []string{fixture.HomeTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.team_home_fixtures = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.AwayTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixture.AwayTeamTable,
			Columns: []string{fixture.AwayTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.team_away_fixtures = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.SeasonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixture.SeasonTable,
			Columns: []string{fixture.SeasonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.season_fixtures = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.LineupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixture.LineupsTable,
			Columns: []string{fixture.LineupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixturelineups.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FixtureEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixture.FixtureEventsTable,
			Columns: []string{fixture.FixtureEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixtureevents.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FixtureCreateBulk is the builder for creating many Fixture entities in bulk.
type FixtureCreateBulk struct {
	config
	builders []*FixtureCreate
}

// Save creates the Fixture entities in the database.
func (fcb *FixtureCreateBulk) Save(ctx context.Context) ([]*Fixture, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fixture, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FixtureMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FixtureCreateBulk) SaveX(ctx context.Context) []*Fixture {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FixtureCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FixtureCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
