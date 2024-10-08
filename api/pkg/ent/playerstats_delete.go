// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlayerStatsDelete is the builder for deleting a PlayerStats entity.
type PlayerStatsDelete struct {
	config
	hooks    []Hook
	mutation *PlayerStatsMutation
}

// Where appends a list predicates to the PlayerStatsDelete builder.
func (psd *PlayerStatsDelete) Where(ps ...predicate.PlayerStats) *PlayerStatsDelete {
	psd.mutation.Where(ps...)
	return psd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (psd *PlayerStatsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, psd.sqlExec, psd.mutation, psd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (psd *PlayerStatsDelete) ExecX(ctx context.Context) int {
	n, err := psd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (psd *PlayerStatsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(playerstats.Table, sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt))
	if ps := psd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, psd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	psd.mutation.done = true
	return affected, err
}

// PlayerStatsDeleteOne is the builder for deleting a single PlayerStats entity.
type PlayerStatsDeleteOne struct {
	psd *PlayerStatsDelete
}

// Where appends a list predicates to the PlayerStatsDelete builder.
func (psdo *PlayerStatsDeleteOne) Where(ps ...predicate.PlayerStats) *PlayerStatsDeleteOne {
	psdo.psd.mutation.Where(ps...)
	return psdo
}

// Exec executes the deletion query.
func (psdo *PlayerStatsDeleteOne) Exec(ctx context.Context) error {
	n, err := psdo.psd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{playerstats.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (psdo *PlayerStatsDeleteOne) ExecX(ctx context.Context) {
	if err := psdo.Exec(ctx); err != nil {
		panic(err)
	}
}
