// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/birth"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BirthUpdate is the builder for updating Birth entities.
type BirthUpdate struct {
	config
	hooks    []Hook
	mutation *BirthMutation
}

// Where appends a list predicates to the BirthUpdate builder.
func (bu *BirthUpdate) Where(ps ...predicate.Birth) *BirthUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetDate sets the "date" field.
func (bu *BirthUpdate) SetDate(t time.Time) *BirthUpdate {
	bu.mutation.SetDate(t)
	return bu
}

// SetPlace sets the "place" field.
func (bu *BirthUpdate) SetPlace(s string) *BirthUpdate {
	bu.mutation.SetPlace(s)
	return bu
}

// SetCountry sets the "country" field.
func (bu *BirthUpdate) SetCountry(s string) *BirthUpdate {
	bu.mutation.SetCountry(s)
	return bu
}

// AddPlayerIDs adds the "player" edge to the Player entity by IDs.
func (bu *BirthUpdate) AddPlayerIDs(ids ...int) *BirthUpdate {
	bu.mutation.AddPlayerIDs(ids...)
	return bu
}

// AddPlayer adds the "player" edges to the Player entity.
func (bu *BirthUpdate) AddPlayer(p ...*Player) *BirthUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddPlayerIDs(ids...)
}

// Mutation returns the BirthMutation object of the builder.
func (bu *BirthUpdate) Mutation() *BirthMutation {
	return bu.mutation
}

// ClearPlayer clears all "player" edges to the Player entity.
func (bu *BirthUpdate) ClearPlayer() *BirthUpdate {
	bu.mutation.ClearPlayer()
	return bu
}

// RemovePlayerIDs removes the "player" edge to Player entities by IDs.
func (bu *BirthUpdate) RemovePlayerIDs(ids ...int) *BirthUpdate {
	bu.mutation.RemovePlayerIDs(ids...)
	return bu
}

// RemovePlayer removes "player" edges to Player entities.
func (bu *BirthUpdate) RemovePlayer(p ...*Player) *BirthUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemovePlayerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BirthUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BirthUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BirthUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BirthUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BirthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(birth.Table, birth.Columns, sqlgraph.NewFieldSpec(birth.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Date(); ok {
		_spec.SetField(birth.FieldDate, field.TypeTime, value)
	}
	if value, ok := bu.mutation.Place(); ok {
		_spec.SetField(birth.FieldPlace, field.TypeString, value)
	}
	if value, ok := bu.mutation.Country(); ok {
		_spec.SetField(birth.FieldCountry, field.TypeString, value)
	}
	if bu.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   birth.PlayerTable,
			Columns: []string{birth.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedPlayerIDs(); len(nodes) > 0 && !bu.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   birth.PlayerTable,
			Columns: []string{birth.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   birth.PlayerTable,
			Columns: []string{birth.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{birth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BirthUpdateOne is the builder for updating a single Birth entity.
type BirthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BirthMutation
}

// SetDate sets the "date" field.
func (buo *BirthUpdateOne) SetDate(t time.Time) *BirthUpdateOne {
	buo.mutation.SetDate(t)
	return buo
}

// SetPlace sets the "place" field.
func (buo *BirthUpdateOne) SetPlace(s string) *BirthUpdateOne {
	buo.mutation.SetPlace(s)
	return buo
}

// SetCountry sets the "country" field.
func (buo *BirthUpdateOne) SetCountry(s string) *BirthUpdateOne {
	buo.mutation.SetCountry(s)
	return buo
}

// AddPlayerIDs adds the "player" edge to the Player entity by IDs.
func (buo *BirthUpdateOne) AddPlayerIDs(ids ...int) *BirthUpdateOne {
	buo.mutation.AddPlayerIDs(ids...)
	return buo
}

// AddPlayer adds the "player" edges to the Player entity.
func (buo *BirthUpdateOne) AddPlayer(p ...*Player) *BirthUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddPlayerIDs(ids...)
}

// Mutation returns the BirthMutation object of the builder.
func (buo *BirthUpdateOne) Mutation() *BirthMutation {
	return buo.mutation
}

// ClearPlayer clears all "player" edges to the Player entity.
func (buo *BirthUpdateOne) ClearPlayer() *BirthUpdateOne {
	buo.mutation.ClearPlayer()
	return buo
}

// RemovePlayerIDs removes the "player" edge to Player entities by IDs.
func (buo *BirthUpdateOne) RemovePlayerIDs(ids ...int) *BirthUpdateOne {
	buo.mutation.RemovePlayerIDs(ids...)
	return buo
}

// RemovePlayer removes "player" edges to Player entities.
func (buo *BirthUpdateOne) RemovePlayer(p ...*Player) *BirthUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemovePlayerIDs(ids...)
}

// Where appends a list predicates to the BirthUpdate builder.
func (buo *BirthUpdateOne) Where(ps ...predicate.Birth) *BirthUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BirthUpdateOne) Select(field string, fields ...string) *BirthUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Birth entity.
func (buo *BirthUpdateOne) Save(ctx context.Context) (*Birth, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BirthUpdateOne) SaveX(ctx context.Context) *Birth {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BirthUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BirthUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BirthUpdateOne) sqlSave(ctx context.Context) (_node *Birth, err error) {
	_spec := sqlgraph.NewUpdateSpec(birth.Table, birth.Columns, sqlgraph.NewFieldSpec(birth.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Birth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, birth.FieldID)
		for _, f := range fields {
			if !birth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != birth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Date(); ok {
		_spec.SetField(birth.FieldDate, field.TypeTime, value)
	}
	if value, ok := buo.mutation.Place(); ok {
		_spec.SetField(birth.FieldPlace, field.TypeString, value)
	}
	if value, ok := buo.mutation.Country(); ok {
		_spec.SetField(birth.FieldCountry, field.TypeString, value)
	}
	if buo.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   birth.PlayerTable,
			Columns: []string{birth.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedPlayerIDs(); len(nodes) > 0 && !buo.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   birth.PlayerTable,
			Columns: []string{birth.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   birth.PlayerTable,
			Columns: []string{birth.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Birth{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{birth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
