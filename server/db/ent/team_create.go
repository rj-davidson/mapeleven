// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/fixture"
	"mapeleven/db/ent/league"
	"mapeleven/db/ent/player"
	"mapeleven/db/ent/standings"
	"mapeleven/db/ent/team"
	"mapeleven/db/ent/teamseason"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeamCreate is the builder for creating a Team entity.
type TeamCreate struct {
	config
	mutation *TeamMutation
	hooks    []Hook
}

// SetSlug sets the "slug" field.
func (tc *TeamCreate) SetSlug(s string) *TeamCreate {
	tc.mutation.SetSlug(s)
	return tc
}

// SetName sets the "name" field.
func (tc *TeamCreate) SetName(s string) *TeamCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetCode sets the "code" field.
func (tc *TeamCreate) SetCode(s string) *TeamCreate {
	tc.mutation.SetCode(s)
	return tc
}

// SetFounded sets the "founded" field.
func (tc *TeamCreate) SetFounded(i int) *TeamCreate {
	tc.mutation.SetFounded(i)
	return tc
}

// SetNational sets the "national" field.
func (tc *TeamCreate) SetNational(b bool) *TeamCreate {
	tc.mutation.SetNational(b)
	return tc
}

// SetLogo sets the "logo" field.
func (tc *TeamCreate) SetLogo(s string) *TeamCreate {
	tc.mutation.SetLogo(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TeamCreate) SetID(i int) *TeamCreate {
	tc.mutation.SetID(i)
	return tc
}

// AddStandingIDs adds the "standings" edge to the Standings entity by IDs.
func (tc *TeamCreate) AddStandingIDs(ids ...int) *TeamCreate {
	tc.mutation.AddStandingIDs(ids...)
	return tc
}

// AddStandings adds the "standings" edges to the Standings entity.
func (tc *TeamCreate) AddStandings(s ...*Standings) *TeamCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddStandingIDs(ids...)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (tc *TeamCreate) SetCountryID(id int) *TeamCreate {
	tc.mutation.SetCountryID(id)
	return tc
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (tc *TeamCreate) SetNillableCountryID(id *int) *TeamCreate {
	if id != nil {
		tc = tc.SetCountryID(*id)
	}
	return tc
}

// SetCountry sets the "country" edge to the Country entity.
func (tc *TeamCreate) SetCountry(c *Country) *TeamCreate {
	return tc.SetCountryID(c.ID)
}

// AddLeagueIDs adds the "leagues" edge to the League entity by IDs.
func (tc *TeamCreate) AddLeagueIDs(ids ...int) *TeamCreate {
	tc.mutation.AddLeagueIDs(ids...)
	return tc
}

// AddLeagues adds the "leagues" edges to the League entity.
func (tc *TeamCreate) AddLeagues(l ...*League) *TeamCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return tc.AddLeagueIDs(ids...)
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (tc *TeamCreate) AddPlayerIDs(ids ...int) *TeamCreate {
	tc.mutation.AddPlayerIDs(ids...)
	return tc
}

// AddPlayers adds the "players" edges to the Player entity.
func (tc *TeamCreate) AddPlayers(p ...*Player) *TeamCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddPlayerIDs(ids...)
}

// AddTeamSeasonIDs adds the "teamSeasons" edge to the TeamSeason entity by IDs.
func (tc *TeamCreate) AddTeamSeasonIDs(ids ...int) *TeamCreate {
	tc.mutation.AddTeamSeasonIDs(ids...)
	return tc
}

// AddTeamSeasons adds the "teamSeasons" edges to the TeamSeason entity.
func (tc *TeamCreate) AddTeamSeasons(t ...*TeamSeason) *TeamCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTeamSeasonIDs(ids...)
}

// AddHomeFixtureIDs adds the "homeFixtures" edge to the Fixture entity by IDs.
func (tc *TeamCreate) AddHomeFixtureIDs(ids ...int) *TeamCreate {
	tc.mutation.AddHomeFixtureIDs(ids...)
	return tc
}

// AddHomeFixtures adds the "homeFixtures" edges to the Fixture entity.
func (tc *TeamCreate) AddHomeFixtures(f ...*Fixture) *TeamCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tc.AddHomeFixtureIDs(ids...)
}

// AddAwayFixtureIDs adds the "awayFixtures" edge to the Fixture entity by IDs.
func (tc *TeamCreate) AddAwayFixtureIDs(ids ...int) *TeamCreate {
	tc.mutation.AddAwayFixtureIDs(ids...)
	return tc
}

// AddAwayFixtures adds the "awayFixtures" edges to the Fixture entity.
func (tc *TeamCreate) AddAwayFixtures(f ...*Fixture) *TeamCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tc.AddAwayFixtureIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tc *TeamCreate) Mutation() *TeamMutation {
	return tc.mutation
}

// Save creates the Team in the database.
func (tc *TeamCreate) Save(ctx context.Context) (*Team, error) {
	return withHooks[*Team, TeamMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeamCreate) SaveX(ctx context.Context) *Team {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeamCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeamCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeamCreate) check() error {
	if _, ok := tc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Team.slug"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Team.name"`)}
	}
	if _, ok := tc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Team.code"`)}
	}
	if v, ok := tc.mutation.Code(); ok {
		if err := team.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Team.code": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Founded(); !ok {
		return &ValidationError{Name: "founded", err: errors.New(`ent: missing required field "Team.founded"`)}
	}
	if _, ok := tc.mutation.National(); !ok {
		return &ValidationError{Name: "national", err: errors.New(`ent: missing required field "Team.national"`)}
	}
	if _, ok := tc.mutation.Logo(); !ok {
		return &ValidationError{Name: "logo", err: errors.New(`ent: missing required field "Team.logo"`)}
	}
	return nil
}

func (tc *TeamCreate) sqlSave(ctx context.Context) (*Team, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TeamCreate) createSpec() (*Team, *sqlgraph.CreateSpec) {
	var (
		_node = &Team{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(team.Table, sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Slug(); ok {
		_spec.SetField(team.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Code(); ok {
		_spec.SetField(team.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := tc.mutation.Founded(); ok {
		_spec.SetField(team.FieldFounded, field.TypeInt, value)
		_node.Founded = value
	}
	if value, ok := tc.mutation.National(); ok {
		_spec.SetField(team.FieldNational, field.TypeBool, value)
		_node.National = value
	}
	if value, ok := tc.mutation.Logo(); ok {
		_spec.SetField(team.FieldLogo, field.TypeString, value)
		_node.Logo = value
	}
	if nodes := tc.mutation.StandingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.StandingsTable,
			Columns: []string{team.StandingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   team.CountryTable,
			Columns: []string{team.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.country_teams = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.LeaguesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   team.LeaguesTable,
			Columns: team.LeaguesPrimaryKey,
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
	if nodes := tc.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.PlayersTable,
			Columns: team.PlayersPrimaryKey,
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
	if nodes := tc.mutation.TeamSeasonsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamSeasonsTable,
			Columns: []string{team.TeamSeasonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamseason.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.HomeFixturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.HomeFixturesTable,
			Columns: []string{team.HomeFixturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.AwayFixturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.AwayFixturesTable,
			Columns: []string{team.AwayFixturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TeamCreateBulk is the builder for creating many Team entities in bulk.
type TeamCreateBulk struct {
	config
	builders []*TeamCreate
}

// Save creates the Team entities in the database.
func (tcb *TeamCreateBulk) Save(ctx context.Context) ([]*Team, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Team, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeamMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeamCreateBulk) SaveX(ctx context.Context) []*Team {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeamCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeamCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
