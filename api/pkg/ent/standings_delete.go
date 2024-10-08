// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/standings"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StandingsDelete is the builder for deleting a Standings entity.
type StandingsDelete struct {
	config
	hooks    []Hook
	mutation *StandingsMutation
}

// Where appends a list predicates to the StandingsDelete builder.
func (sd *StandingsDelete) Where(ps ...predicate.Standings) *StandingsDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *StandingsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *StandingsDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *StandingsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(standings.Table, sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt))
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

// StandingsDeleteOne is the builder for deleting a single Standings entity.
type StandingsDeleteOne struct {
	sd *StandingsDelete
}

// Where appends a list predicates to the StandingsDelete builder.
func (sdo *StandingsDeleteOne) Where(ps ...predicate.Standings) *StandingsDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *StandingsDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{standings.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *StandingsDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
