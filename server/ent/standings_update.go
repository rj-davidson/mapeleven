// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/league"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/standings"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/team"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StandingsUpdate is the builder for updating Standings entities.
type StandingsUpdate struct {
	config
	hooks    []Hook
	mutation *StandingsMutation
}

// Where appends a list predicates to the StandingsUpdate builder.
func (su *StandingsUpdate) Where(ps ...predicate.Standings) *StandingsUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetRank sets the "rank" field.
func (su *StandingsUpdate) SetRank(i int) *StandingsUpdate {
	su.mutation.ResetRank()
	su.mutation.SetRank(i)
	return su
}

// AddRank adds i to the "rank" field.
func (su *StandingsUpdate) AddRank(i int) *StandingsUpdate {
	su.mutation.AddRank(i)
	return su
}

// SetDescription sets the "description" field.
func (su *StandingsUpdate) SetDescription(s string) *StandingsUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetLeagueID sets the "league" edge to the League entity by ID.
func (su *StandingsUpdate) SetLeagueID(id int) *StandingsUpdate {
	su.mutation.SetLeagueID(id)
	return su
}

// SetNillableLeagueID sets the "league" edge to the League entity by ID if the given value is not nil.
func (su *StandingsUpdate) SetNillableLeagueID(id *int) *StandingsUpdate {
	if id != nil {
		su = su.SetLeagueID(*id)
	}
	return su
}

// SetLeague sets the "league" edge to the League entity.
func (su *StandingsUpdate) SetLeague(l *League) *StandingsUpdate {
	return su.SetLeagueID(l.ID)
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (su *StandingsUpdate) SetTeamID(id int) *StandingsUpdate {
	su.mutation.SetTeamID(id)
	return su
}

// SetNillableTeamID sets the "team" edge to the Team entity by ID if the given value is not nil.
func (su *StandingsUpdate) SetNillableTeamID(id *int) *StandingsUpdate {
	if id != nil {
		su = su.SetTeamID(*id)
	}
	return su
}

// SetTeam sets the "team" edge to the Team entity.
func (su *StandingsUpdate) SetTeam(t *Team) *StandingsUpdate {
	return su.SetTeamID(t.ID)
}

// Mutation returns the StandingsMutation object of the builder.
func (su *StandingsUpdate) Mutation() *StandingsMutation {
	return su.mutation
}

// ClearLeague clears the "league" edge to the League entity.
func (su *StandingsUpdate) ClearLeague() *StandingsUpdate {
	su.mutation.ClearLeague()
	return su
}

// ClearTeam clears the "team" edge to the Team entity.
func (su *StandingsUpdate) ClearTeam() *StandingsUpdate {
	su.mutation.ClearTeam()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StandingsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, StandingsMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StandingsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StandingsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StandingsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StandingsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(standings.Table, standings.Columns, sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Rank(); ok {
		_spec.SetField(standings.FieldRank, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedRank(); ok {
		_spec.AddField(standings.FieldRank, field.TypeInt, value)
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.SetField(standings.FieldDescription, field.TypeString, value)
	}
	if su.mutation.LeagueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   standings.LeagueTable,
			Columns: []string{standings.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.LeagueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   standings.LeagueTable,
			Columns: []string{standings.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   standings.TeamTable,
			Columns: []string{standings.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   standings.TeamTable,
			Columns: []string{standings.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{standings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StandingsUpdateOne is the builder for updating a single Standings entity.
type StandingsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StandingsMutation
}

// SetRank sets the "rank" field.
func (suo *StandingsUpdateOne) SetRank(i int) *StandingsUpdateOne {
	suo.mutation.ResetRank()
	suo.mutation.SetRank(i)
	return suo
}

// AddRank adds i to the "rank" field.
func (suo *StandingsUpdateOne) AddRank(i int) *StandingsUpdateOne {
	suo.mutation.AddRank(i)
	return suo
}

// SetDescription sets the "description" field.
func (suo *StandingsUpdateOne) SetDescription(s string) *StandingsUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetLeagueID sets the "league" edge to the League entity by ID.
func (suo *StandingsUpdateOne) SetLeagueID(id int) *StandingsUpdateOne {
	suo.mutation.SetLeagueID(id)
	return suo
}

// SetNillableLeagueID sets the "league" edge to the League entity by ID if the given value is not nil.
func (suo *StandingsUpdateOne) SetNillableLeagueID(id *int) *StandingsUpdateOne {
	if id != nil {
		suo = suo.SetLeagueID(*id)
	}
	return suo
}

// SetLeague sets the "league" edge to the League entity.
func (suo *StandingsUpdateOne) SetLeague(l *League) *StandingsUpdateOne {
	return suo.SetLeagueID(l.ID)
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (suo *StandingsUpdateOne) SetTeamID(id int) *StandingsUpdateOne {
	suo.mutation.SetTeamID(id)
	return suo
}

// SetNillableTeamID sets the "team" edge to the Team entity by ID if the given value is not nil.
func (suo *StandingsUpdateOne) SetNillableTeamID(id *int) *StandingsUpdateOne {
	if id != nil {
		suo = suo.SetTeamID(*id)
	}
	return suo
}

// SetTeam sets the "team" edge to the Team entity.
func (suo *StandingsUpdateOne) SetTeam(t *Team) *StandingsUpdateOne {
	return suo.SetTeamID(t.ID)
}

// Mutation returns the StandingsMutation object of the builder.
func (suo *StandingsUpdateOne) Mutation() *StandingsMutation {
	return suo.mutation
}

// ClearLeague clears the "league" edge to the League entity.
func (suo *StandingsUpdateOne) ClearLeague() *StandingsUpdateOne {
	suo.mutation.ClearLeague()
	return suo
}

// ClearTeam clears the "team" edge to the Team entity.
func (suo *StandingsUpdateOne) ClearTeam() *StandingsUpdateOne {
	suo.mutation.ClearTeam()
	return suo
}

// Where appends a list predicates to the StandingsUpdate builder.
func (suo *StandingsUpdateOne) Where(ps ...predicate.Standings) *StandingsUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StandingsUpdateOne) Select(field string, fields ...string) *StandingsUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Standings entity.
func (suo *StandingsUpdateOne) Save(ctx context.Context) (*Standings, error) {
	return withHooks[*Standings, StandingsMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StandingsUpdateOne) SaveX(ctx context.Context) *Standings {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StandingsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StandingsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StandingsUpdateOne) sqlSave(ctx context.Context) (_node *Standings, err error) {
	_spec := sqlgraph.NewUpdateSpec(standings.Table, standings.Columns, sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Standings.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, standings.FieldID)
		for _, f := range fields {
			if !standings.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != standings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Rank(); ok {
		_spec.SetField(standings.FieldRank, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedRank(); ok {
		_spec.AddField(standings.FieldRank, field.TypeInt, value)
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.SetField(standings.FieldDescription, field.TypeString, value)
	}
	if suo.mutation.LeagueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   standings.LeagueTable,
			Columns: []string{standings.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.LeagueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   standings.LeagueTable,
			Columns: []string{standings.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   standings.TeamTable,
			Columns: []string{standings.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   standings.TeamTable,
			Columns: []string{standings.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Standings{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{standings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
