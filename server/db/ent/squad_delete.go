// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"mapeleven/db/ent/predicate"
	"mapeleven/db/ent/squad"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SquadDelete is the builder for deleting a Squad entity.
type SquadDelete struct {
	config
	hooks    []Hook
	mutation *SquadMutation
}

// Where appends a list predicates to the SquadDelete builder.
func (sd *SquadDelete) Where(ps ...predicate.Squad) *SquadDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SquadDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, SquadMutation](ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SquadDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SquadDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(squad.Table, sqlgraph.NewFieldSpec(squad.FieldID, field.TypeInt))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SquadDeleteOne is the builder for deleting a single Squad entity.
type SquadDeleteOne struct {
	sd *SquadDelete
}

// Where appends a list predicates to the SquadDelete builder.
func (sdo *SquadDeleteOne) Where(ps ...predicate.Squad) *SquadDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SquadDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{squad.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SquadDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
