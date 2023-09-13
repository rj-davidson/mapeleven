// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/fixture"
	"mapeleven/db/ent/predicate"
	"mapeleven/db/ent/standings"
	"mapeleven/db/ent/team"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeamUpdate is the builder for updating Team entities.
type TeamUpdate struct {
	config
	hooks    []Hook
	mutation *TeamMutation
}

// Where appends a list predicates to the TeamUpdate builder.
func (tu *TeamUpdate) Where(ps ...predicate.Team) *TeamUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TeamUpdate) SetName(s string) *TeamUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetCode sets the "code" field.
func (tu *TeamUpdate) SetCode(s string) *TeamUpdate {
	tu.mutation.SetCode(s)
	return tu
}

// SetFounded sets the "founded" field.
func (tu *TeamUpdate) SetFounded(i int) *TeamUpdate {
	tu.mutation.ResetFounded()
	tu.mutation.SetFounded(i)
	return tu
}

// AddFounded adds i to the "founded" field.
func (tu *TeamUpdate) AddFounded(i int) *TeamUpdate {
	tu.mutation.AddFounded(i)
	return tu
}

// SetNational sets the "national" field.
func (tu *TeamUpdate) SetNational(b bool) *TeamUpdate {
	tu.mutation.SetNational(b)
	return tu
}

// SetLogo sets the "logo" field.
func (tu *TeamUpdate) SetLogo(s string) *TeamUpdate {
	tu.mutation.SetLogo(s)
	return tu
}

// AddStandingIDs adds the "standings" edge to the Standings entity by IDs.
func (tu *TeamUpdate) AddStandingIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddStandingIDs(ids...)
	return tu
}

// AddStandings adds the "standings" edges to the Standings entity.
func (tu *TeamUpdate) AddStandings(s ...*Standings) *TeamUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.AddStandingIDs(ids...)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (tu *TeamUpdate) SetCountryID(id int) *TeamUpdate {
	tu.mutation.SetCountryID(id)
	return tu
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (tu *TeamUpdate) SetNillableCountryID(id *int) *TeamUpdate {
	if id != nil {
		tu = tu.SetCountryID(*id)
	}
	return tu
}

// SetCountry sets the "country" edge to the Country entity.
func (tu *TeamUpdate) SetCountry(c *Country) *TeamUpdate {
	return tu.SetCountryID(c.ID)
}

// AddHomeFixtureIDs adds the "homeFixtures" edge to the Fixture entity by IDs.
func (tu *TeamUpdate) AddHomeFixtureIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddHomeFixtureIDs(ids...)
	return tu
}

// AddHomeFixtures adds the "homeFixtures" edges to the Fixture entity.
func (tu *TeamUpdate) AddHomeFixtures(f ...*Fixture) *TeamUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.AddHomeFixtureIDs(ids...)
}

// AddAwayFixtureIDs adds the "awayFixtures" edge to the Fixture entity by IDs.
func (tu *TeamUpdate) AddAwayFixtureIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddAwayFixtureIDs(ids...)
	return tu
}

// AddAwayFixtures adds the "awayFixtures" edges to the Fixture entity.
func (tu *TeamUpdate) AddAwayFixtures(f ...*Fixture) *TeamUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.AddAwayFixtureIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tu *TeamUpdate) Mutation() *TeamMutation {
	return tu.mutation
}

// ClearStandings clears all "standings" edges to the Standings entity.
func (tu *TeamUpdate) ClearStandings() *TeamUpdate {
	tu.mutation.ClearStandings()
	return tu
}

// RemoveStandingIDs removes the "standings" edge to Standings entities by IDs.
func (tu *TeamUpdate) RemoveStandingIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveStandingIDs(ids...)
	return tu
}

// RemoveStandings removes "standings" edges to Standings entities.
func (tu *TeamUpdate) RemoveStandings(s ...*Standings) *TeamUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.RemoveStandingIDs(ids...)
}

// ClearCountry clears the "country" edge to the Country entity.
func (tu *TeamUpdate) ClearCountry() *TeamUpdate {
	tu.mutation.ClearCountry()
	return tu
}

// ClearHomeFixtures clears all "homeFixtures" edges to the Fixture entity.
func (tu *TeamUpdate) ClearHomeFixtures() *TeamUpdate {
	tu.mutation.ClearHomeFixtures()
	return tu
}

// RemoveHomeFixtureIDs removes the "homeFixtures" edge to Fixture entities by IDs.
func (tu *TeamUpdate) RemoveHomeFixtureIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveHomeFixtureIDs(ids...)
	return tu
}

// RemoveHomeFixtures removes "homeFixtures" edges to Fixture entities.
func (tu *TeamUpdate) RemoveHomeFixtures(f ...*Fixture) *TeamUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.RemoveHomeFixtureIDs(ids...)
}

// ClearAwayFixtures clears all "awayFixtures" edges to the Fixture entity.
func (tu *TeamUpdate) ClearAwayFixtures() *TeamUpdate {
	tu.mutation.ClearAwayFixtures()
	return tu
}

// RemoveAwayFixtureIDs removes the "awayFixtures" edge to Fixture entities by IDs.
func (tu *TeamUpdate) RemoveAwayFixtureIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveAwayFixtureIDs(ids...)
	return tu
}

// RemoveAwayFixtures removes "awayFixtures" edges to Fixture entities.
func (tu *TeamUpdate) RemoveAwayFixtures(f ...*Fixture) *TeamUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.RemoveAwayFixtureIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeamUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TeamMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeamUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeamUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeamUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TeamUpdate) check() error {
	if v, ok := tu.mutation.Code(); ok {
		if err := team.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Team.code": %w`, err)}
		}
	}
	return nil
}

func (tu *TeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Code(); ok {
		_spec.SetField(team.FieldCode, field.TypeString, value)
	}
	if value, ok := tu.mutation.Founded(); ok {
		_spec.SetField(team.FieldFounded, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedFounded(); ok {
		_spec.AddField(team.FieldFounded, field.TypeInt, value)
	}
	if value, ok := tu.mutation.National(); ok {
		_spec.SetField(team.FieldNational, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Logo(); ok {
		_spec.SetField(team.FieldLogo, field.TypeString, value)
	}
	if tu.mutation.StandingsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedStandingsIDs(); len(nodes) > 0 && !tu.mutation.StandingsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.StandingsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.CountryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CountryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.HomeFixturesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedHomeFixturesIDs(); len(nodes) > 0 && !tu.mutation.HomeFixturesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.HomeFixturesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.AwayFixturesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedAwayFixturesIDs(); len(nodes) > 0 && !tu.mutation.AwayFixturesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AwayFixturesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeamUpdateOne is the builder for updating a single Team entity.
type TeamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamMutation
}

// SetName sets the "name" field.
func (tuo *TeamUpdateOne) SetName(s string) *TeamUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetCode sets the "code" field.
func (tuo *TeamUpdateOne) SetCode(s string) *TeamUpdateOne {
	tuo.mutation.SetCode(s)
	return tuo
}

// SetFounded sets the "founded" field.
func (tuo *TeamUpdateOne) SetFounded(i int) *TeamUpdateOne {
	tuo.mutation.ResetFounded()
	tuo.mutation.SetFounded(i)
	return tuo
}

// AddFounded adds i to the "founded" field.
func (tuo *TeamUpdateOne) AddFounded(i int) *TeamUpdateOne {
	tuo.mutation.AddFounded(i)
	return tuo
}

// SetNational sets the "national" field.
func (tuo *TeamUpdateOne) SetNational(b bool) *TeamUpdateOne {
	tuo.mutation.SetNational(b)
	return tuo
}

// SetLogo sets the "logo" field.
func (tuo *TeamUpdateOne) SetLogo(s string) *TeamUpdateOne {
	tuo.mutation.SetLogo(s)
	return tuo
}

// AddStandingIDs adds the "standings" edge to the Standings entity by IDs.
func (tuo *TeamUpdateOne) AddStandingIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddStandingIDs(ids...)
	return tuo
}

// AddStandings adds the "standings" edges to the Standings entity.
func (tuo *TeamUpdateOne) AddStandings(s ...*Standings) *TeamUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.AddStandingIDs(ids...)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (tuo *TeamUpdateOne) SetCountryID(id int) *TeamUpdateOne {
	tuo.mutation.SetCountryID(id)
	return tuo
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableCountryID(id *int) *TeamUpdateOne {
	if id != nil {
		tuo = tuo.SetCountryID(*id)
	}
	return tuo
}

// SetCountry sets the "country" edge to the Country entity.
func (tuo *TeamUpdateOne) SetCountry(c *Country) *TeamUpdateOne {
	return tuo.SetCountryID(c.ID)
}

// AddHomeFixtureIDs adds the "homeFixtures" edge to the Fixture entity by IDs.
func (tuo *TeamUpdateOne) AddHomeFixtureIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddHomeFixtureIDs(ids...)
	return tuo
}

// AddHomeFixtures adds the "homeFixtures" edges to the Fixture entity.
func (tuo *TeamUpdateOne) AddHomeFixtures(f ...*Fixture) *TeamUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.AddHomeFixtureIDs(ids...)
}

// AddAwayFixtureIDs adds the "awayFixtures" edge to the Fixture entity by IDs.
func (tuo *TeamUpdateOne) AddAwayFixtureIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddAwayFixtureIDs(ids...)
	return tuo
}

// AddAwayFixtures adds the "awayFixtures" edges to the Fixture entity.
func (tuo *TeamUpdateOne) AddAwayFixtures(f ...*Fixture) *TeamUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.AddAwayFixtureIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tuo *TeamUpdateOne) Mutation() *TeamMutation {
	return tuo.mutation
}

// ClearStandings clears all "standings" edges to the Standings entity.
func (tuo *TeamUpdateOne) ClearStandings() *TeamUpdateOne {
	tuo.mutation.ClearStandings()
	return tuo
}

// RemoveStandingIDs removes the "standings" edge to Standings entities by IDs.
func (tuo *TeamUpdateOne) RemoveStandingIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveStandingIDs(ids...)
	return tuo
}

// RemoveStandings removes "standings" edges to Standings entities.
func (tuo *TeamUpdateOne) RemoveStandings(s ...*Standings) *TeamUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.RemoveStandingIDs(ids...)
}

// ClearCountry clears the "country" edge to the Country entity.
func (tuo *TeamUpdateOne) ClearCountry() *TeamUpdateOne {
	tuo.mutation.ClearCountry()
	return tuo
}

// ClearHomeFixtures clears all "homeFixtures" edges to the Fixture entity.
func (tuo *TeamUpdateOne) ClearHomeFixtures() *TeamUpdateOne {
	tuo.mutation.ClearHomeFixtures()
	return tuo
}

// RemoveHomeFixtureIDs removes the "homeFixtures" edge to Fixture entities by IDs.
func (tuo *TeamUpdateOne) RemoveHomeFixtureIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveHomeFixtureIDs(ids...)
	return tuo
}

// RemoveHomeFixtures removes "homeFixtures" edges to Fixture entities.
func (tuo *TeamUpdateOne) RemoveHomeFixtures(f ...*Fixture) *TeamUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.RemoveHomeFixtureIDs(ids...)
}

// ClearAwayFixtures clears all "awayFixtures" edges to the Fixture entity.
func (tuo *TeamUpdateOne) ClearAwayFixtures() *TeamUpdateOne {
	tuo.mutation.ClearAwayFixtures()
	return tuo
}

// RemoveAwayFixtureIDs removes the "awayFixtures" edge to Fixture entities by IDs.
func (tuo *TeamUpdateOne) RemoveAwayFixtureIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveAwayFixtureIDs(ids...)
	return tuo
}

// RemoveAwayFixtures removes "awayFixtures" edges to Fixture entities.
func (tuo *TeamUpdateOne) RemoveAwayFixtures(f ...*Fixture) *TeamUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.RemoveAwayFixtureIDs(ids...)
}

// Where appends a list predicates to the TeamUpdate builder.
func (tuo *TeamUpdateOne) Where(ps ...predicate.Team) *TeamUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeamUpdateOne) Select(field string, fields ...string) *TeamUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Team entity.
func (tuo *TeamUpdateOne) Save(ctx context.Context) (*Team, error) {
	return withHooks[*Team, TeamMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeamUpdateOne) SaveX(ctx context.Context) *Team {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeamUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeamUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TeamUpdateOne) check() error {
	if v, ok := tuo.mutation.Code(); ok {
		if err := team.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Team.code": %w`, err)}
		}
	}
	return nil
}

func (tuo *TeamUpdateOne) sqlSave(ctx context.Context) (_node *Team, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Team.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, team.FieldID)
		for _, f := range fields {
			if !team.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != team.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Code(); ok {
		_spec.SetField(team.FieldCode, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Founded(); ok {
		_spec.SetField(team.FieldFounded, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedFounded(); ok {
		_spec.AddField(team.FieldFounded, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.National(); ok {
		_spec.SetField(team.FieldNational, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Logo(); ok {
		_spec.SetField(team.FieldLogo, field.TypeString, value)
	}
	if tuo.mutation.StandingsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedStandingsIDs(); len(nodes) > 0 && !tuo.mutation.StandingsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.StandingsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.CountryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CountryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.HomeFixturesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedHomeFixturesIDs(); len(nodes) > 0 && !tuo.mutation.HomeFixturesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.HomeFixturesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.AwayFixturesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedAwayFixturesIDs(); len(nodes) > 0 && !tuo.mutation.AwayFixturesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AwayFixturesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Team{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
