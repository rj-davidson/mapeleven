// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgoals"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSGoalsUpdate is the builder for updating PSGoals entities.
type PSGoalsUpdate struct {
	config
	hooks    []Hook
	mutation *PSGoalsMutation
}

// Where appends a list predicates to the PSGoalsUpdate builder.
func (pgu *PSGoalsUpdate) Where(ps ...predicate.PSGoals) *PSGoalsUpdate {
	pgu.mutation.Where(ps...)
	return pgu
}

// SetTotalGoals sets the "totalGoals" field.
func (pgu *PSGoalsUpdate) SetTotalGoals(i int) *PSGoalsUpdate {
	pgu.mutation.ResetTotalGoals()
	pgu.mutation.SetTotalGoals(i)
	return pgu
}

// AddTotalGoals adds i to the "totalGoals" field.
func (pgu *PSGoalsUpdate) AddTotalGoals(i int) *PSGoalsUpdate {
	pgu.mutation.AddTotalGoals(i)
	return pgu
}

// SetConcededGoals sets the "concededGoals" field.
func (pgu *PSGoalsUpdate) SetConcededGoals(i int) *PSGoalsUpdate {
	pgu.mutation.ResetConcededGoals()
	pgu.mutation.SetConcededGoals(i)
	return pgu
}

// AddConcededGoals adds i to the "concededGoals" field.
func (pgu *PSGoalsUpdate) AddConcededGoals(i int) *PSGoalsUpdate {
	pgu.mutation.AddConcededGoals(i)
	return pgu
}

// SetAssistGoals sets the "assistGoals" field.
func (pgu *PSGoalsUpdate) SetAssistGoals(i int) *PSGoalsUpdate {
	pgu.mutation.ResetAssistGoals()
	pgu.mutation.SetAssistGoals(i)
	return pgu
}

// AddAssistGoals adds i to the "assistGoals" field.
func (pgu *PSGoalsUpdate) AddAssistGoals(i int) *PSGoalsUpdate {
	pgu.mutation.AddAssistGoals(i)
	return pgu
}

// SetSaveGoals sets the "saveGoals" field.
func (pgu *PSGoalsUpdate) SetSaveGoals(i int) *PSGoalsUpdate {
	pgu.mutation.ResetSaveGoals()
	pgu.mutation.SetSaveGoals(i)
	return pgu
}

// SetNillableSaveGoals sets the "saveGoals" field if the given value is not nil.
func (pgu *PSGoalsUpdate) SetNillableSaveGoals(i *int) *PSGoalsUpdate {
	if i != nil {
		pgu.SetSaveGoals(*i)
	}
	return pgu
}

// AddSaveGoals adds i to the "saveGoals" field.
func (pgu *PSGoalsUpdate) AddSaveGoals(i int) *PSGoalsUpdate {
	pgu.mutation.AddSaveGoals(i)
	return pgu
}

// SetShotsTotal sets the "shotsTotal" field.
func (pgu *PSGoalsUpdate) SetShotsTotal(i int) *PSGoalsUpdate {
	pgu.mutation.ResetShotsTotal()
	pgu.mutation.SetShotsTotal(i)
	return pgu
}

// AddShotsTotal adds i to the "shotsTotal" field.
func (pgu *PSGoalsUpdate) AddShotsTotal(i int) *PSGoalsUpdate {
	pgu.mutation.AddShotsTotal(i)
	return pgu
}

// SetShotsOn sets the "shotsOn" field.
func (pgu *PSGoalsUpdate) SetShotsOn(i int) *PSGoalsUpdate {
	pgu.mutation.ResetShotsOn()
	pgu.mutation.SetShotsOn(i)
	return pgu
}

// AddShotsOn adds i to the "shotsOn" field.
func (pgu *PSGoalsUpdate) AddShotsOn(i int) *PSGoalsUpdate {
	pgu.mutation.AddShotsOn(i)
	return pgu
}

// AddPlayerStatIDs adds the "playerStats" edge to the PlayerStats entity by IDs.
func (pgu *PSGoalsUpdate) AddPlayerStatIDs(ids ...int) *PSGoalsUpdate {
	pgu.mutation.AddPlayerStatIDs(ids...)
	return pgu
}

// AddPlayerStats adds the "playerStats" edges to the PlayerStats entity.
func (pgu *PSGoalsUpdate) AddPlayerStats(p ...*PlayerStats) *PSGoalsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pgu.AddPlayerStatIDs(ids...)
}

// Mutation returns the PSGoalsMutation object of the builder.
func (pgu *PSGoalsUpdate) Mutation() *PSGoalsMutation {
	return pgu.mutation
}

// ClearPlayerStats clears all "playerStats" edges to the PlayerStats entity.
func (pgu *PSGoalsUpdate) ClearPlayerStats() *PSGoalsUpdate {
	pgu.mutation.ClearPlayerStats()
	return pgu
}

// RemovePlayerStatIDs removes the "playerStats" edge to PlayerStats entities by IDs.
func (pgu *PSGoalsUpdate) RemovePlayerStatIDs(ids ...int) *PSGoalsUpdate {
	pgu.mutation.RemovePlayerStatIDs(ids...)
	return pgu
}

// RemovePlayerStats removes "playerStats" edges to PlayerStats entities.
func (pgu *PSGoalsUpdate) RemovePlayerStats(p ...*PlayerStats) *PSGoalsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pgu.RemovePlayerStatIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pgu *PSGoalsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pgu.sqlSave, pgu.mutation, pgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pgu *PSGoalsUpdate) SaveX(ctx context.Context) int {
	affected, err := pgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pgu *PSGoalsUpdate) Exec(ctx context.Context) error {
	_, err := pgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pgu *PSGoalsUpdate) ExecX(ctx context.Context) {
	if err := pgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pgu *PSGoalsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(psgoals.Table, psgoals.Columns, sqlgraph.NewFieldSpec(psgoals.FieldID, field.TypeInt))
	if ps := pgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pgu.mutation.TotalGoals(); ok {
		_spec.SetField(psgoals.FieldTotalGoals, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedTotalGoals(); ok {
		_spec.AddField(psgoals.FieldTotalGoals, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.ConcededGoals(); ok {
		_spec.SetField(psgoals.FieldConcededGoals, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedConcededGoals(); ok {
		_spec.AddField(psgoals.FieldConcededGoals, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AssistGoals(); ok {
		_spec.SetField(psgoals.FieldAssistGoals, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedAssistGoals(); ok {
		_spec.AddField(psgoals.FieldAssistGoals, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.SaveGoals(); ok {
		_spec.SetField(psgoals.FieldSaveGoals, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedSaveGoals(); ok {
		_spec.AddField(psgoals.FieldSaveGoals, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.ShotsTotal(); ok {
		_spec.SetField(psgoals.FieldShotsTotal, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedShotsTotal(); ok {
		_spec.AddField(psgoals.FieldShotsTotal, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.ShotsOn(); ok {
		_spec.SetField(psgoals.FieldShotsOn, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedShotsOn(); ok {
		_spec.AddField(psgoals.FieldShotsOn, field.TypeInt, value)
	}
	if pgu.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   psgoals.PlayerStatsTable,
			Columns: psgoals.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.RemovedPlayerStatsIDs(); len(nodes) > 0 && !pgu.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   psgoals.PlayerStatsTable,
			Columns: psgoals.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   psgoals.PlayerStatsTable,
			Columns: psgoals.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{psgoals.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pgu.mutation.done = true
	return n, nil
}

// PSGoalsUpdateOne is the builder for updating a single PSGoals entity.
type PSGoalsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PSGoalsMutation
}

// SetTotalGoals sets the "totalGoals" field.
func (pguo *PSGoalsUpdateOne) SetTotalGoals(i int) *PSGoalsUpdateOne {
	pguo.mutation.ResetTotalGoals()
	pguo.mutation.SetTotalGoals(i)
	return pguo
}

// AddTotalGoals adds i to the "totalGoals" field.
func (pguo *PSGoalsUpdateOne) AddTotalGoals(i int) *PSGoalsUpdateOne {
	pguo.mutation.AddTotalGoals(i)
	return pguo
}

// SetConcededGoals sets the "concededGoals" field.
func (pguo *PSGoalsUpdateOne) SetConcededGoals(i int) *PSGoalsUpdateOne {
	pguo.mutation.ResetConcededGoals()
	pguo.mutation.SetConcededGoals(i)
	return pguo
}

// AddConcededGoals adds i to the "concededGoals" field.
func (pguo *PSGoalsUpdateOne) AddConcededGoals(i int) *PSGoalsUpdateOne {
	pguo.mutation.AddConcededGoals(i)
	return pguo
}

// SetAssistGoals sets the "assistGoals" field.
func (pguo *PSGoalsUpdateOne) SetAssistGoals(i int) *PSGoalsUpdateOne {
	pguo.mutation.ResetAssistGoals()
	pguo.mutation.SetAssistGoals(i)
	return pguo
}

// AddAssistGoals adds i to the "assistGoals" field.
func (pguo *PSGoalsUpdateOne) AddAssistGoals(i int) *PSGoalsUpdateOne {
	pguo.mutation.AddAssistGoals(i)
	return pguo
}

// SetSaveGoals sets the "saveGoals" field.
func (pguo *PSGoalsUpdateOne) SetSaveGoals(i int) *PSGoalsUpdateOne {
	pguo.mutation.ResetSaveGoals()
	pguo.mutation.SetSaveGoals(i)
	return pguo
}

// SetNillableSaveGoals sets the "saveGoals" field if the given value is not nil.
func (pguo *PSGoalsUpdateOne) SetNillableSaveGoals(i *int) *PSGoalsUpdateOne {
	if i != nil {
		pguo.SetSaveGoals(*i)
	}
	return pguo
}

// AddSaveGoals adds i to the "saveGoals" field.
func (pguo *PSGoalsUpdateOne) AddSaveGoals(i int) *PSGoalsUpdateOne {
	pguo.mutation.AddSaveGoals(i)
	return pguo
}

// SetShotsTotal sets the "shotsTotal" field.
func (pguo *PSGoalsUpdateOne) SetShotsTotal(i int) *PSGoalsUpdateOne {
	pguo.mutation.ResetShotsTotal()
	pguo.mutation.SetShotsTotal(i)
	return pguo
}

// AddShotsTotal adds i to the "shotsTotal" field.
func (pguo *PSGoalsUpdateOne) AddShotsTotal(i int) *PSGoalsUpdateOne {
	pguo.mutation.AddShotsTotal(i)
	return pguo
}

// SetShotsOn sets the "shotsOn" field.
func (pguo *PSGoalsUpdateOne) SetShotsOn(i int) *PSGoalsUpdateOne {
	pguo.mutation.ResetShotsOn()
	pguo.mutation.SetShotsOn(i)
	return pguo
}

// AddShotsOn adds i to the "shotsOn" field.
func (pguo *PSGoalsUpdateOne) AddShotsOn(i int) *PSGoalsUpdateOne {
	pguo.mutation.AddShotsOn(i)
	return pguo
}

// AddPlayerStatIDs adds the "playerStats" edge to the PlayerStats entity by IDs.
func (pguo *PSGoalsUpdateOne) AddPlayerStatIDs(ids ...int) *PSGoalsUpdateOne {
	pguo.mutation.AddPlayerStatIDs(ids...)
	return pguo
}

// AddPlayerStats adds the "playerStats" edges to the PlayerStats entity.
func (pguo *PSGoalsUpdateOne) AddPlayerStats(p ...*PlayerStats) *PSGoalsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pguo.AddPlayerStatIDs(ids...)
}

// Mutation returns the PSGoalsMutation object of the builder.
func (pguo *PSGoalsUpdateOne) Mutation() *PSGoalsMutation {
	return pguo.mutation
}

// ClearPlayerStats clears all "playerStats" edges to the PlayerStats entity.
func (pguo *PSGoalsUpdateOne) ClearPlayerStats() *PSGoalsUpdateOne {
	pguo.mutation.ClearPlayerStats()
	return pguo
}

// RemovePlayerStatIDs removes the "playerStats" edge to PlayerStats entities by IDs.
func (pguo *PSGoalsUpdateOne) RemovePlayerStatIDs(ids ...int) *PSGoalsUpdateOne {
	pguo.mutation.RemovePlayerStatIDs(ids...)
	return pguo
}

// RemovePlayerStats removes "playerStats" edges to PlayerStats entities.
func (pguo *PSGoalsUpdateOne) RemovePlayerStats(p ...*PlayerStats) *PSGoalsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pguo.RemovePlayerStatIDs(ids...)
}

// Where appends a list predicates to the PSGoalsUpdate builder.
func (pguo *PSGoalsUpdateOne) Where(ps ...predicate.PSGoals) *PSGoalsUpdateOne {
	pguo.mutation.Where(ps...)
	return pguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pguo *PSGoalsUpdateOne) Select(field string, fields ...string) *PSGoalsUpdateOne {
	pguo.fields = append([]string{field}, fields...)
	return pguo
}

// Save executes the query and returns the updated PSGoals entity.
func (pguo *PSGoalsUpdateOne) Save(ctx context.Context) (*PSGoals, error) {
	return withHooks(ctx, pguo.sqlSave, pguo.mutation, pguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pguo *PSGoalsUpdateOne) SaveX(ctx context.Context) *PSGoals {
	node, err := pguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pguo *PSGoalsUpdateOne) Exec(ctx context.Context) error {
	_, err := pguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pguo *PSGoalsUpdateOne) ExecX(ctx context.Context) {
	if err := pguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pguo *PSGoalsUpdateOne) sqlSave(ctx context.Context) (_node *PSGoals, err error) {
	_spec := sqlgraph.NewUpdateSpec(psgoals.Table, psgoals.Columns, sqlgraph.NewFieldSpec(psgoals.FieldID, field.TypeInt))
	id, ok := pguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PSGoals.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, psgoals.FieldID)
		for _, f := range fields {
			if !psgoals.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != psgoals.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pguo.mutation.TotalGoals(); ok {
		_spec.SetField(psgoals.FieldTotalGoals, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedTotalGoals(); ok {
		_spec.AddField(psgoals.FieldTotalGoals, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.ConcededGoals(); ok {
		_spec.SetField(psgoals.FieldConcededGoals, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedConcededGoals(); ok {
		_spec.AddField(psgoals.FieldConcededGoals, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AssistGoals(); ok {
		_spec.SetField(psgoals.FieldAssistGoals, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedAssistGoals(); ok {
		_spec.AddField(psgoals.FieldAssistGoals, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.SaveGoals(); ok {
		_spec.SetField(psgoals.FieldSaveGoals, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedSaveGoals(); ok {
		_spec.AddField(psgoals.FieldSaveGoals, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.ShotsTotal(); ok {
		_spec.SetField(psgoals.FieldShotsTotal, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedShotsTotal(); ok {
		_spec.AddField(psgoals.FieldShotsTotal, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.ShotsOn(); ok {
		_spec.SetField(psgoals.FieldShotsOn, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedShotsOn(); ok {
		_spec.AddField(psgoals.FieldShotsOn, field.TypeInt, value)
	}
	if pguo.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   psgoals.PlayerStatsTable,
			Columns: psgoals.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.RemovedPlayerStatsIDs(); len(nodes) > 0 && !pguo.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   psgoals.PlayerStatsTable,
			Columns: psgoals.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   psgoals.PlayerStatsTable,
			Columns: psgoals.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PSGoals{config: pguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{psgoals.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pguo.mutation.done = true
	return _node, nil
}
