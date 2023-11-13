// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSDefenseUpdate is the builder for updating PSDefense entities.
type PSDefenseUpdate struct {
	config
	hooks    []Hook
	mutation *PSDefenseMutation
}

// Where appends a list predicates to the PSDefenseUpdate builder.
func (pdu *PSDefenseUpdate) Where(ps ...predicate.PSDefense) *PSDefenseUpdate {
	pdu.mutation.Where(ps...)
	return pdu
}

// SetTacklesTotal sets the "TacklesTotal" field.
func (pdu *PSDefenseUpdate) SetTacklesTotal(i int) *PSDefenseUpdate {
	pdu.mutation.ResetTacklesTotal()
	pdu.mutation.SetTacklesTotal(i)
	return pdu
}

// SetNillableTacklesTotal sets the "TacklesTotal" field if the given value is not nil.
func (pdu *PSDefenseUpdate) SetNillableTacklesTotal(i *int) *PSDefenseUpdate {
	if i != nil {
		pdu.SetTacklesTotal(*i)
	}
	return pdu
}

// AddTacklesTotal adds i to the "TacklesTotal" field.
func (pdu *PSDefenseUpdate) AddTacklesTotal(i int) *PSDefenseUpdate {
	pdu.mutation.AddTacklesTotal(i)
	return pdu
}

// SetBlocks sets the "Blocks" field.
func (pdu *PSDefenseUpdate) SetBlocks(i int) *PSDefenseUpdate {
	pdu.mutation.ResetBlocks()
	pdu.mutation.SetBlocks(i)
	return pdu
}

// SetNillableBlocks sets the "Blocks" field if the given value is not nil.
func (pdu *PSDefenseUpdate) SetNillableBlocks(i *int) *PSDefenseUpdate {
	if i != nil {
		pdu.SetBlocks(*i)
	}
	return pdu
}

// AddBlocks adds i to the "Blocks" field.
func (pdu *PSDefenseUpdate) AddBlocks(i int) *PSDefenseUpdate {
	pdu.mutation.AddBlocks(i)
	return pdu
}

// SetInterceptions sets the "Interceptions" field.
func (pdu *PSDefenseUpdate) SetInterceptions(i int) *PSDefenseUpdate {
	pdu.mutation.ResetInterceptions()
	pdu.mutation.SetInterceptions(i)
	return pdu
}

// SetNillableInterceptions sets the "Interceptions" field if the given value is not nil.
func (pdu *PSDefenseUpdate) SetNillableInterceptions(i *int) *PSDefenseUpdate {
	if i != nil {
		pdu.SetInterceptions(*i)
	}
	return pdu
}

// AddInterceptions adds i to the "Interceptions" field.
func (pdu *PSDefenseUpdate) AddInterceptions(i int) *PSDefenseUpdate {
	pdu.mutation.AddInterceptions(i)
	return pdu
}

// SetDuelsTotal sets the "DuelsTotal" field.
func (pdu *PSDefenseUpdate) SetDuelsTotal(i int) *PSDefenseUpdate {
	pdu.mutation.ResetDuelsTotal()
	pdu.mutation.SetDuelsTotal(i)
	return pdu
}

// SetNillableDuelsTotal sets the "DuelsTotal" field if the given value is not nil.
func (pdu *PSDefenseUpdate) SetNillableDuelsTotal(i *int) *PSDefenseUpdate {
	if i != nil {
		pdu.SetDuelsTotal(*i)
	}
	return pdu
}

// AddDuelsTotal adds i to the "DuelsTotal" field.
func (pdu *PSDefenseUpdate) AddDuelsTotal(i int) *PSDefenseUpdate {
	pdu.mutation.AddDuelsTotal(i)
	return pdu
}

// SetWonDuels sets the "WonDuels" field.
func (pdu *PSDefenseUpdate) SetWonDuels(i int) *PSDefenseUpdate {
	pdu.mutation.ResetWonDuels()
	pdu.mutation.SetWonDuels(i)
	return pdu
}

// SetNillableWonDuels sets the "WonDuels" field if the given value is not nil.
func (pdu *PSDefenseUpdate) SetNillableWonDuels(i *int) *PSDefenseUpdate {
	if i != nil {
		pdu.SetWonDuels(*i)
	}
	return pdu
}

// AddWonDuels adds i to the "WonDuels" field.
func (pdu *PSDefenseUpdate) AddWonDuels(i int) *PSDefenseUpdate {
	pdu.mutation.AddWonDuels(i)
	return pdu
}

// SetLastUpdated sets the "lastUpdated" field.
func (pdu *PSDefenseUpdate) SetLastUpdated(t time.Time) *PSDefenseUpdate {
	pdu.mutation.SetLastUpdated(t)
	return pdu
}

// ClearLastUpdated clears the value of the "lastUpdated" field.
func (pdu *PSDefenseUpdate) ClearLastUpdated() *PSDefenseUpdate {
	pdu.mutation.ClearLastUpdated()
	return pdu
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (pdu *PSDefenseUpdate) SetPlayerStatsID(id int) *PSDefenseUpdate {
	pdu.mutation.SetPlayerStatsID(id)
	return pdu
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (pdu *PSDefenseUpdate) SetNillablePlayerStatsID(id *int) *PSDefenseUpdate {
	if id != nil {
		pdu = pdu.SetPlayerStatsID(*id)
	}
	return pdu
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (pdu *PSDefenseUpdate) SetPlayerStats(p *PlayerStats) *PSDefenseUpdate {
	return pdu.SetPlayerStatsID(p.ID)
}

// Mutation returns the PSDefenseMutation object of the builder.
func (pdu *PSDefenseUpdate) Mutation() *PSDefenseMutation {
	return pdu.mutation
}

// ClearPlayerStats clears the "playerStats" edge to the PlayerStats entity.
func (pdu *PSDefenseUpdate) ClearPlayerStats() *PSDefenseUpdate {
	pdu.mutation.ClearPlayerStats()
	return pdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pdu *PSDefenseUpdate) Save(ctx context.Context) (int, error) {
	pdu.defaults()
	return withHooks(ctx, pdu.sqlSave, pdu.mutation, pdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pdu *PSDefenseUpdate) SaveX(ctx context.Context) int {
	affected, err := pdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pdu *PSDefenseUpdate) Exec(ctx context.Context) error {
	_, err := pdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdu *PSDefenseUpdate) ExecX(ctx context.Context) {
	if err := pdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdu *PSDefenseUpdate) defaults() {
	if _, ok := pdu.mutation.LastUpdated(); !ok && !pdu.mutation.LastUpdatedCleared() {
		v := psdefense.UpdateDefaultLastUpdated()
		pdu.mutation.SetLastUpdated(v)
	}
}

func (pdu *PSDefenseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(psdefense.Table, psdefense.Columns, sqlgraph.NewFieldSpec(psdefense.FieldID, field.TypeInt))
	if ps := pdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pdu.mutation.TacklesTotal(); ok {
		_spec.SetField(psdefense.FieldTacklesTotal, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.AddedTacklesTotal(); ok {
		_spec.AddField(psdefense.FieldTacklesTotal, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.Blocks(); ok {
		_spec.SetField(psdefense.FieldBlocks, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.AddedBlocks(); ok {
		_spec.AddField(psdefense.FieldBlocks, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.Interceptions(); ok {
		_spec.SetField(psdefense.FieldInterceptions, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.AddedInterceptions(); ok {
		_spec.AddField(psdefense.FieldInterceptions, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.DuelsTotal(); ok {
		_spec.SetField(psdefense.FieldDuelsTotal, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.AddedDuelsTotal(); ok {
		_spec.AddField(psdefense.FieldDuelsTotal, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.WonDuels(); ok {
		_spec.SetField(psdefense.FieldWonDuels, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.AddedWonDuels(); ok {
		_spec.AddField(psdefense.FieldWonDuels, field.TypeInt, value)
	}
	if value, ok := pdu.mutation.LastUpdated(); ok {
		_spec.SetField(psdefense.FieldLastUpdated, field.TypeTime, value)
	}
	if pdu.mutation.LastUpdatedCleared() {
		_spec.ClearField(psdefense.FieldLastUpdated, field.TypeTime)
	}
	if pdu.mutation.PlayerStatsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pdu.mutation.PlayerStatsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{psdefense.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pdu.mutation.done = true
	return n, nil
}

// PSDefenseUpdateOne is the builder for updating a single PSDefense entity.
type PSDefenseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PSDefenseMutation
}

// SetTacklesTotal sets the "TacklesTotal" field.
func (pduo *PSDefenseUpdateOne) SetTacklesTotal(i int) *PSDefenseUpdateOne {
	pduo.mutation.ResetTacklesTotal()
	pduo.mutation.SetTacklesTotal(i)
	return pduo
}

// SetNillableTacklesTotal sets the "TacklesTotal" field if the given value is not nil.
func (pduo *PSDefenseUpdateOne) SetNillableTacklesTotal(i *int) *PSDefenseUpdateOne {
	if i != nil {
		pduo.SetTacklesTotal(*i)
	}
	return pduo
}

// AddTacklesTotal adds i to the "TacklesTotal" field.
func (pduo *PSDefenseUpdateOne) AddTacklesTotal(i int) *PSDefenseUpdateOne {
	pduo.mutation.AddTacklesTotal(i)
	return pduo
}

// SetBlocks sets the "Blocks" field.
func (pduo *PSDefenseUpdateOne) SetBlocks(i int) *PSDefenseUpdateOne {
	pduo.mutation.ResetBlocks()
	pduo.mutation.SetBlocks(i)
	return pduo
}

// SetNillableBlocks sets the "Blocks" field if the given value is not nil.
func (pduo *PSDefenseUpdateOne) SetNillableBlocks(i *int) *PSDefenseUpdateOne {
	if i != nil {
		pduo.SetBlocks(*i)
	}
	return pduo
}

// AddBlocks adds i to the "Blocks" field.
func (pduo *PSDefenseUpdateOne) AddBlocks(i int) *PSDefenseUpdateOne {
	pduo.mutation.AddBlocks(i)
	return pduo
}

// SetInterceptions sets the "Interceptions" field.
func (pduo *PSDefenseUpdateOne) SetInterceptions(i int) *PSDefenseUpdateOne {
	pduo.mutation.ResetInterceptions()
	pduo.mutation.SetInterceptions(i)
	return pduo
}

// SetNillableInterceptions sets the "Interceptions" field if the given value is not nil.
func (pduo *PSDefenseUpdateOne) SetNillableInterceptions(i *int) *PSDefenseUpdateOne {
	if i != nil {
		pduo.SetInterceptions(*i)
	}
	return pduo
}

// AddInterceptions adds i to the "Interceptions" field.
func (pduo *PSDefenseUpdateOne) AddInterceptions(i int) *PSDefenseUpdateOne {
	pduo.mutation.AddInterceptions(i)
	return pduo
}

// SetDuelsTotal sets the "DuelsTotal" field.
func (pduo *PSDefenseUpdateOne) SetDuelsTotal(i int) *PSDefenseUpdateOne {
	pduo.mutation.ResetDuelsTotal()
	pduo.mutation.SetDuelsTotal(i)
	return pduo
}

// SetNillableDuelsTotal sets the "DuelsTotal" field if the given value is not nil.
func (pduo *PSDefenseUpdateOne) SetNillableDuelsTotal(i *int) *PSDefenseUpdateOne {
	if i != nil {
		pduo.SetDuelsTotal(*i)
	}
	return pduo
}

// AddDuelsTotal adds i to the "DuelsTotal" field.
func (pduo *PSDefenseUpdateOne) AddDuelsTotal(i int) *PSDefenseUpdateOne {
	pduo.mutation.AddDuelsTotal(i)
	return pduo
}

// SetWonDuels sets the "WonDuels" field.
func (pduo *PSDefenseUpdateOne) SetWonDuels(i int) *PSDefenseUpdateOne {
	pduo.mutation.ResetWonDuels()
	pduo.mutation.SetWonDuels(i)
	return pduo
}

// SetNillableWonDuels sets the "WonDuels" field if the given value is not nil.
func (pduo *PSDefenseUpdateOne) SetNillableWonDuels(i *int) *PSDefenseUpdateOne {
	if i != nil {
		pduo.SetWonDuels(*i)
	}
	return pduo
}

// AddWonDuels adds i to the "WonDuels" field.
func (pduo *PSDefenseUpdateOne) AddWonDuels(i int) *PSDefenseUpdateOne {
	pduo.mutation.AddWonDuels(i)
	return pduo
}

// SetLastUpdated sets the "lastUpdated" field.
func (pduo *PSDefenseUpdateOne) SetLastUpdated(t time.Time) *PSDefenseUpdateOne {
	pduo.mutation.SetLastUpdated(t)
	return pduo
}

// ClearLastUpdated clears the value of the "lastUpdated" field.
func (pduo *PSDefenseUpdateOne) ClearLastUpdated() *PSDefenseUpdateOne {
	pduo.mutation.ClearLastUpdated()
	return pduo
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (pduo *PSDefenseUpdateOne) SetPlayerStatsID(id int) *PSDefenseUpdateOne {
	pduo.mutation.SetPlayerStatsID(id)
	return pduo
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (pduo *PSDefenseUpdateOne) SetNillablePlayerStatsID(id *int) *PSDefenseUpdateOne {
	if id != nil {
		pduo = pduo.SetPlayerStatsID(*id)
	}
	return pduo
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (pduo *PSDefenseUpdateOne) SetPlayerStats(p *PlayerStats) *PSDefenseUpdateOne {
	return pduo.SetPlayerStatsID(p.ID)
}

// Mutation returns the PSDefenseMutation object of the builder.
func (pduo *PSDefenseUpdateOne) Mutation() *PSDefenseMutation {
	return pduo.mutation
}

// ClearPlayerStats clears the "playerStats" edge to the PlayerStats entity.
func (pduo *PSDefenseUpdateOne) ClearPlayerStats() *PSDefenseUpdateOne {
	pduo.mutation.ClearPlayerStats()
	return pduo
}

// Where appends a list predicates to the PSDefenseUpdate builder.
func (pduo *PSDefenseUpdateOne) Where(ps ...predicate.PSDefense) *PSDefenseUpdateOne {
	pduo.mutation.Where(ps...)
	return pduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pduo *PSDefenseUpdateOne) Select(field string, fields ...string) *PSDefenseUpdateOne {
	pduo.fields = append([]string{field}, fields...)
	return pduo
}

// Save executes the query and returns the updated PSDefense entity.
func (pduo *PSDefenseUpdateOne) Save(ctx context.Context) (*PSDefense, error) {
	pduo.defaults()
	return withHooks(ctx, pduo.sqlSave, pduo.mutation, pduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pduo *PSDefenseUpdateOne) SaveX(ctx context.Context) *PSDefense {
	node, err := pduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pduo *PSDefenseUpdateOne) Exec(ctx context.Context) error {
	_, err := pduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pduo *PSDefenseUpdateOne) ExecX(ctx context.Context) {
	if err := pduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pduo *PSDefenseUpdateOne) defaults() {
	if _, ok := pduo.mutation.LastUpdated(); !ok && !pduo.mutation.LastUpdatedCleared() {
		v := psdefense.UpdateDefaultLastUpdated()
		pduo.mutation.SetLastUpdated(v)
	}
}

func (pduo *PSDefenseUpdateOne) sqlSave(ctx context.Context) (_node *PSDefense, err error) {
	_spec := sqlgraph.NewUpdateSpec(psdefense.Table, psdefense.Columns, sqlgraph.NewFieldSpec(psdefense.FieldID, field.TypeInt))
	id, ok := pduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PSDefense.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, psdefense.FieldID)
		for _, f := range fields {
			if !psdefense.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != psdefense.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pduo.mutation.TacklesTotal(); ok {
		_spec.SetField(psdefense.FieldTacklesTotal, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.AddedTacklesTotal(); ok {
		_spec.AddField(psdefense.FieldTacklesTotal, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.Blocks(); ok {
		_spec.SetField(psdefense.FieldBlocks, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.AddedBlocks(); ok {
		_spec.AddField(psdefense.FieldBlocks, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.Interceptions(); ok {
		_spec.SetField(psdefense.FieldInterceptions, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.AddedInterceptions(); ok {
		_spec.AddField(psdefense.FieldInterceptions, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.DuelsTotal(); ok {
		_spec.SetField(psdefense.FieldDuelsTotal, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.AddedDuelsTotal(); ok {
		_spec.AddField(psdefense.FieldDuelsTotal, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.WonDuels(); ok {
		_spec.SetField(psdefense.FieldWonDuels, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.AddedWonDuels(); ok {
		_spec.AddField(psdefense.FieldWonDuels, field.TypeInt, value)
	}
	if value, ok := pduo.mutation.LastUpdated(); ok {
		_spec.SetField(psdefense.FieldLastUpdated, field.TypeTime, value)
	}
	if pduo.mutation.LastUpdatedCleared() {
		_spec.ClearField(psdefense.FieldLastUpdated, field.TypeTime)
	}
	if pduo.mutation.PlayerStatsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pduo.mutation.PlayerStatsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PSDefense{config: pduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{psdefense.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pduo.mutation.done = true
	return _node, nil
}
