// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"mapeleven/db/ent/predicate"
	"mapeleven/db/ent/tsbiggest"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TSBiggestDelete is the builder for deleting a TSBiggest entity.
type TSBiggestDelete struct {
	config
	hooks    []Hook
	mutation *TSBiggestMutation
}

// Where appends a list predicates to the TSBiggestDelete builder.
func (tbd *TSBiggestDelete) Where(ps ...predicate.TSBiggest) *TSBiggestDelete {
	tbd.mutation.Where(ps...)
	return tbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tbd *TSBiggestDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TSBiggestMutation](ctx, tbd.sqlExec, tbd.mutation, tbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tbd *TSBiggestDelete) ExecX(ctx context.Context) int {
	n, err := tbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tbd *TSBiggestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tsbiggest.Table, sqlgraph.NewFieldSpec(tsbiggest.FieldID, field.TypeInt))
	if ps := tbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tbd.mutation.done = true
	return affected, err
}

// TSBiggestDeleteOne is the builder for deleting a single TSBiggest entity.
type TSBiggestDeleteOne struct {
	tbd *TSBiggestDelete
}

// Where appends a list predicates to the TSBiggestDelete builder.
func (tbdo *TSBiggestDeleteOne) Where(ps ...predicate.TSBiggest) *TSBiggestDeleteOne {
	tbdo.tbd.mutation.Where(ps...)
	return tbdo
}

// Exec executes the deletion query.
func (tbdo *TSBiggestDeleteOne) Exec(ctx context.Context) error {
	n, err := tbdo.tbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tsbiggest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tbdo *TSBiggestDeleteOne) ExecX(ctx context.Context) {
	if err := tbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
