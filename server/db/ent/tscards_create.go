// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/team"
	"mapeleven/db/ent/tscards"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TSCardsCreate is the builder for creating a TSCards entity.
type TSCardsCreate struct {
	config
	mutation *TSCardsMutation
	hooks    []Hook
}

// SetYellow0To15Total sets the "yellow0To15Total" field.
func (tcc *TSCardsCreate) SetYellow0To15Total(i int) *TSCardsCreate {
	tcc.mutation.SetYellow0To15Total(i)
	return tcc
}

// SetNillableYellow0To15Total sets the "yellow0To15Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow0To15Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetYellow0To15Total(*i)
	}
	return tcc
}

// SetYellow0To15Percentage sets the "yellow0To15Percentage" field.
func (tcc *TSCardsCreate) SetYellow0To15Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetYellow0To15Percentage(s)
	return tcc
}

// SetNillableYellow0To15Percentage sets the "yellow0To15Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow0To15Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetYellow0To15Percentage(*s)
	}
	return tcc
}

// SetYellow16To30Total sets the "yellow16To30Total" field.
func (tcc *TSCardsCreate) SetYellow16To30Total(i int) *TSCardsCreate {
	tcc.mutation.SetYellow16To30Total(i)
	return tcc
}

// SetNillableYellow16To30Total sets the "yellow16To30Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow16To30Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetYellow16To30Total(*i)
	}
	return tcc
}

// SetYellow16To30Percentage sets the "yellow16To30Percentage" field.
func (tcc *TSCardsCreate) SetYellow16To30Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetYellow16To30Percentage(s)
	return tcc
}

// SetNillableYellow16To30Percentage sets the "yellow16To30Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow16To30Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetYellow16To30Percentage(*s)
	}
	return tcc
}

// SetYellow31To45Total sets the "yellow31To45Total" field.
func (tcc *TSCardsCreate) SetYellow31To45Total(i int) *TSCardsCreate {
	tcc.mutation.SetYellow31To45Total(i)
	return tcc
}

// SetNillableYellow31To45Total sets the "yellow31To45Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow31To45Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetYellow31To45Total(*i)
	}
	return tcc
}

// SetYellow31To45Percentage sets the "yellow31To45Percentage" field.
func (tcc *TSCardsCreate) SetYellow31To45Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetYellow31To45Percentage(s)
	return tcc
}

// SetNillableYellow31To45Percentage sets the "yellow31To45Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow31To45Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetYellow31To45Percentage(*s)
	}
	return tcc
}

// SetYellow46To60Total sets the "yellow46To60Total" field.
func (tcc *TSCardsCreate) SetYellow46To60Total(i int) *TSCardsCreate {
	tcc.mutation.SetYellow46To60Total(i)
	return tcc
}

// SetNillableYellow46To60Total sets the "yellow46To60Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow46To60Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetYellow46To60Total(*i)
	}
	return tcc
}

// SetYellow46To60Percentage sets the "yellow46To60Percentage" field.
func (tcc *TSCardsCreate) SetYellow46To60Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetYellow46To60Percentage(s)
	return tcc
}

// SetNillableYellow46To60Percentage sets the "yellow46To60Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow46To60Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetYellow46To60Percentage(*s)
	}
	return tcc
}

// SetYellow61To75Total sets the "yellow61To75Total" field.
func (tcc *TSCardsCreate) SetYellow61To75Total(i int) *TSCardsCreate {
	tcc.mutation.SetYellow61To75Total(i)
	return tcc
}

// SetNillableYellow61To75Total sets the "yellow61To75Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow61To75Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetYellow61To75Total(*i)
	}
	return tcc
}

// SetYellow61To75Percentage sets the "yellow61To75Percentage" field.
func (tcc *TSCardsCreate) SetYellow61To75Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetYellow61To75Percentage(s)
	return tcc
}

// SetNillableYellow61To75Percentage sets the "yellow61To75Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow61To75Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetYellow61To75Percentage(*s)
	}
	return tcc
}

// SetYellow76To90Total sets the "yellow76To90Total" field.
func (tcc *TSCardsCreate) SetYellow76To90Total(i int) *TSCardsCreate {
	tcc.mutation.SetYellow76To90Total(i)
	return tcc
}

// SetNillableYellow76To90Total sets the "yellow76To90Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow76To90Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetYellow76To90Total(*i)
	}
	return tcc
}

// SetYellow76To90Percentage sets the "yellow76To90Percentage" field.
func (tcc *TSCardsCreate) SetYellow76To90Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetYellow76To90Percentage(s)
	return tcc
}

// SetNillableYellow76To90Percentage sets the "yellow76To90Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow76To90Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetYellow76To90Percentage(*s)
	}
	return tcc
}

// SetYellow91to105Total sets the "yellow91to105Total" field.
func (tcc *TSCardsCreate) SetYellow91to105Total(i int) *TSCardsCreate {
	tcc.mutation.SetYellow91to105Total(i)
	return tcc
}

// SetNillableYellow91to105Total sets the "yellow91to105Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow91to105Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetYellow91to105Total(*i)
	}
	return tcc
}

// SetYellow91to105Percentage sets the "yellow91to105Percentage" field.
func (tcc *TSCardsCreate) SetYellow91to105Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetYellow91to105Percentage(s)
	return tcc
}

// SetNillableYellow91to105Percentage sets the "yellow91to105Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow91to105Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetYellow91to105Percentage(*s)
	}
	return tcc
}

// SetYellow106To120Total sets the "yellow106To120Total" field.
func (tcc *TSCardsCreate) SetYellow106To120Total(i int) *TSCardsCreate {
	tcc.mutation.SetYellow106To120Total(i)
	return tcc
}

// SetNillableYellow106To120Total sets the "yellow106To120Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow106To120Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetYellow106To120Total(*i)
	}
	return tcc
}

// SetYellow106To120Percentage sets the "yellow106To120Percentage" field.
func (tcc *TSCardsCreate) SetYellow106To120Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetYellow106To120Percentage(s)
	return tcc
}

// SetNillableYellow106To120Percentage sets the "yellow106To120Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableYellow106To120Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetYellow106To120Percentage(*s)
	}
	return tcc
}

// SetRed0To15Total sets the "red0To15Total" field.
func (tcc *TSCardsCreate) SetRed0To15Total(i int) *TSCardsCreate {
	tcc.mutation.SetRed0To15Total(i)
	return tcc
}

// SetNillableRed0To15Total sets the "red0To15Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed0To15Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetRed0To15Total(*i)
	}
	return tcc
}

// SetRed0To15Percentage sets the "red0To15Percentage" field.
func (tcc *TSCardsCreate) SetRed0To15Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetRed0To15Percentage(s)
	return tcc
}

// SetNillableRed0To15Percentage sets the "red0To15Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed0To15Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetRed0To15Percentage(*s)
	}
	return tcc
}

// SetRed16To30Total sets the "red16To30Total" field.
func (tcc *TSCardsCreate) SetRed16To30Total(i int) *TSCardsCreate {
	tcc.mutation.SetRed16To30Total(i)
	return tcc
}

// SetNillableRed16To30Total sets the "red16To30Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed16To30Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetRed16To30Total(*i)
	}
	return tcc
}

// SetRed16To30Percentage sets the "red16To30Percentage" field.
func (tcc *TSCardsCreate) SetRed16To30Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetRed16To30Percentage(s)
	return tcc
}

// SetNillableRed16To30Percentage sets the "red16To30Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed16To30Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetRed16To30Percentage(*s)
	}
	return tcc
}

// SetRed31To45Total sets the "red31To45Total" field.
func (tcc *TSCardsCreate) SetRed31To45Total(i int) *TSCardsCreate {
	tcc.mutation.SetRed31To45Total(i)
	return tcc
}

// SetNillableRed31To45Total sets the "red31To45Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed31To45Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetRed31To45Total(*i)
	}
	return tcc
}

// SetRed31To45Percentage sets the "red31To45Percentage" field.
func (tcc *TSCardsCreate) SetRed31To45Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetRed31To45Percentage(s)
	return tcc
}

// SetNillableRed31To45Percentage sets the "red31To45Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed31To45Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetRed31To45Percentage(*s)
	}
	return tcc
}

// SetRed46To60Total sets the "red46To60Total" field.
func (tcc *TSCardsCreate) SetRed46To60Total(i int) *TSCardsCreate {
	tcc.mutation.SetRed46To60Total(i)
	return tcc
}

// SetNillableRed46To60Total sets the "red46To60Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed46To60Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetRed46To60Total(*i)
	}
	return tcc
}

// SetRed46To60Percentage sets the "red46To60Percentage" field.
func (tcc *TSCardsCreate) SetRed46To60Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetRed46To60Percentage(s)
	return tcc
}

// SetNillableRed46To60Percentage sets the "red46To60Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed46To60Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetRed46To60Percentage(*s)
	}
	return tcc
}

// SetRed61To75Total sets the "red61To75Total" field.
func (tcc *TSCardsCreate) SetRed61To75Total(i int) *TSCardsCreate {
	tcc.mutation.SetRed61To75Total(i)
	return tcc
}

// SetNillableRed61To75Total sets the "red61To75Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed61To75Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetRed61To75Total(*i)
	}
	return tcc
}

// SetRed61To75Percentage sets the "red61To75Percentage" field.
func (tcc *TSCardsCreate) SetRed61To75Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetRed61To75Percentage(s)
	return tcc
}

// SetNillableRed61To75Percentage sets the "red61To75Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed61To75Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetRed61To75Percentage(*s)
	}
	return tcc
}

// SetRed76To90Total sets the "red76To90Total" field.
func (tcc *TSCardsCreate) SetRed76To90Total(i int) *TSCardsCreate {
	tcc.mutation.SetRed76To90Total(i)
	return tcc
}

// SetNillableRed76To90Total sets the "red76To90Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed76To90Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetRed76To90Total(*i)
	}
	return tcc
}

// SetRed76To90Percentage sets the "red76To90Percentage" field.
func (tcc *TSCardsCreate) SetRed76To90Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetRed76To90Percentage(s)
	return tcc
}

// SetNillableRed76To90Percentage sets the "red76To90Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed76To90Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetRed76To90Percentage(*s)
	}
	return tcc
}

// SetRed91to105Total sets the "red91to105Total" field.
func (tcc *TSCardsCreate) SetRed91to105Total(i int) *TSCardsCreate {
	tcc.mutation.SetRed91to105Total(i)
	return tcc
}

// SetNillableRed91to105Total sets the "red91to105Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed91to105Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetRed91to105Total(*i)
	}
	return tcc
}

// SetRed91to105Percentage sets the "red91to105Percentage" field.
func (tcc *TSCardsCreate) SetRed91to105Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetRed91to105Percentage(s)
	return tcc
}

// SetNillableRed91to105Percentage sets the "red91to105Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed91to105Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetRed91to105Percentage(*s)
	}
	return tcc
}

// SetRed106To120Total sets the "red106To120Total" field.
func (tcc *TSCardsCreate) SetRed106To120Total(i int) *TSCardsCreate {
	tcc.mutation.SetRed106To120Total(i)
	return tcc
}

// SetNillableRed106To120Total sets the "red106To120Total" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed106To120Total(i *int) *TSCardsCreate {
	if i != nil {
		tcc.SetRed106To120Total(*i)
	}
	return tcc
}

// SetRed106To120Percentage sets the "red106To120Percentage" field.
func (tcc *TSCardsCreate) SetRed106To120Percentage(s string) *TSCardsCreate {
	tcc.mutation.SetRed106To120Percentage(s)
	return tcc
}

// SetNillableRed106To120Percentage sets the "red106To120Percentage" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableRed106To120Percentage(s *string) *TSCardsCreate {
	if s != nil {
		tcc.SetRed106To120Percentage(*s)
	}
	return tcc
}

// SetLastUpdated sets the "lastUpdated" field.
func (tcc *TSCardsCreate) SetLastUpdated(t time.Time) *TSCardsCreate {
	tcc.mutation.SetLastUpdated(t)
	return tcc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (tcc *TSCardsCreate) SetNillableLastUpdated(t *time.Time) *TSCardsCreate {
	if t != nil {
		tcc.SetLastUpdated(*t)
	}
	return tcc
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (tcc *TSCardsCreate) SetTeamID(id int) *TSCardsCreate {
	tcc.mutation.SetTeamID(id)
	return tcc
}

// SetTeam sets the "team" edge to the Team entity.
func (tcc *TSCardsCreate) SetTeam(t *Team) *TSCardsCreate {
	return tcc.SetTeamID(t.ID)
}

// Mutation returns the TSCardsMutation object of the builder.
func (tcc *TSCardsCreate) Mutation() *TSCardsMutation {
	return tcc.mutation
}

// Save creates the TSCards in the database.
func (tcc *TSCardsCreate) Save(ctx context.Context) (*TSCards, error) {
	tcc.defaults()
	return withHooks[*TSCards, TSCardsMutation](ctx, tcc.sqlSave, tcc.mutation, tcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tcc *TSCardsCreate) SaveX(ctx context.Context) *TSCards {
	v, err := tcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcc *TSCardsCreate) Exec(ctx context.Context) error {
	_, err := tcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcc *TSCardsCreate) ExecX(ctx context.Context) {
	if err := tcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcc *TSCardsCreate) defaults() {
	if _, ok := tcc.mutation.Yellow0To15Total(); !ok {
		v := tscards.DefaultYellow0To15Total
		tcc.mutation.SetYellow0To15Total(v)
	}
	if _, ok := tcc.mutation.Yellow0To15Percentage(); !ok {
		v := tscards.DefaultYellow0To15Percentage
		tcc.mutation.SetYellow0To15Percentage(v)
	}
	if _, ok := tcc.mutation.Yellow16To30Total(); !ok {
		v := tscards.DefaultYellow16To30Total
		tcc.mutation.SetYellow16To30Total(v)
	}
	if _, ok := tcc.mutation.Yellow16To30Percentage(); !ok {
		v := tscards.DefaultYellow16To30Percentage
		tcc.mutation.SetYellow16To30Percentage(v)
	}
	if _, ok := tcc.mutation.Yellow31To45Total(); !ok {
		v := tscards.DefaultYellow31To45Total
		tcc.mutation.SetYellow31To45Total(v)
	}
	if _, ok := tcc.mutation.Yellow31To45Percentage(); !ok {
		v := tscards.DefaultYellow31To45Percentage
		tcc.mutation.SetYellow31To45Percentage(v)
	}
	if _, ok := tcc.mutation.Yellow46To60Total(); !ok {
		v := tscards.DefaultYellow46To60Total
		tcc.mutation.SetYellow46To60Total(v)
	}
	if _, ok := tcc.mutation.Yellow46To60Percentage(); !ok {
		v := tscards.DefaultYellow46To60Percentage
		tcc.mutation.SetYellow46To60Percentage(v)
	}
	if _, ok := tcc.mutation.Yellow61To75Total(); !ok {
		v := tscards.DefaultYellow61To75Total
		tcc.mutation.SetYellow61To75Total(v)
	}
	if _, ok := tcc.mutation.Yellow61To75Percentage(); !ok {
		v := tscards.DefaultYellow61To75Percentage
		tcc.mutation.SetYellow61To75Percentage(v)
	}
	if _, ok := tcc.mutation.Yellow76To90Total(); !ok {
		v := tscards.DefaultYellow76To90Total
		tcc.mutation.SetYellow76To90Total(v)
	}
	if _, ok := tcc.mutation.Yellow76To90Percentage(); !ok {
		v := tscards.DefaultYellow76To90Percentage
		tcc.mutation.SetYellow76To90Percentage(v)
	}
	if _, ok := tcc.mutation.Yellow91to105Total(); !ok {
		v := tscards.DefaultYellow91to105Total
		tcc.mutation.SetYellow91to105Total(v)
	}
	if _, ok := tcc.mutation.Yellow91to105Percentage(); !ok {
		v := tscards.DefaultYellow91to105Percentage
		tcc.mutation.SetYellow91to105Percentage(v)
	}
	if _, ok := tcc.mutation.Yellow106To120Total(); !ok {
		v := tscards.DefaultYellow106To120Total
		tcc.mutation.SetYellow106To120Total(v)
	}
	if _, ok := tcc.mutation.Yellow106To120Percentage(); !ok {
		v := tscards.DefaultYellow106To120Percentage
		tcc.mutation.SetYellow106To120Percentage(v)
	}
	if _, ok := tcc.mutation.Red0To15Total(); !ok {
		v := tscards.DefaultRed0To15Total
		tcc.mutation.SetRed0To15Total(v)
	}
	if _, ok := tcc.mutation.Red0To15Percentage(); !ok {
		v := tscards.DefaultRed0To15Percentage
		tcc.mutation.SetRed0To15Percentage(v)
	}
	if _, ok := tcc.mutation.Red16To30Total(); !ok {
		v := tscards.DefaultRed16To30Total
		tcc.mutation.SetRed16To30Total(v)
	}
	if _, ok := tcc.mutation.Red16To30Percentage(); !ok {
		v := tscards.DefaultRed16To30Percentage
		tcc.mutation.SetRed16To30Percentage(v)
	}
	if _, ok := tcc.mutation.Red31To45Total(); !ok {
		v := tscards.DefaultRed31To45Total
		tcc.mutation.SetRed31To45Total(v)
	}
	if _, ok := tcc.mutation.Red31To45Percentage(); !ok {
		v := tscards.DefaultRed31To45Percentage
		tcc.mutation.SetRed31To45Percentage(v)
	}
	if _, ok := tcc.mutation.Red46To60Total(); !ok {
		v := tscards.DefaultRed46To60Total
		tcc.mutation.SetRed46To60Total(v)
	}
	if _, ok := tcc.mutation.Red46To60Percentage(); !ok {
		v := tscards.DefaultRed46To60Percentage
		tcc.mutation.SetRed46To60Percentage(v)
	}
	if _, ok := tcc.mutation.Red61To75Total(); !ok {
		v := tscards.DefaultRed61To75Total
		tcc.mutation.SetRed61To75Total(v)
	}
	if _, ok := tcc.mutation.Red61To75Percentage(); !ok {
		v := tscards.DefaultRed61To75Percentage
		tcc.mutation.SetRed61To75Percentage(v)
	}
	if _, ok := tcc.mutation.Red76To90Total(); !ok {
		v := tscards.DefaultRed76To90Total
		tcc.mutation.SetRed76To90Total(v)
	}
	if _, ok := tcc.mutation.Red76To90Percentage(); !ok {
		v := tscards.DefaultRed76To90Percentage
		tcc.mutation.SetRed76To90Percentage(v)
	}
	if _, ok := tcc.mutation.Red91to105Total(); !ok {
		v := tscards.DefaultRed91to105Total
		tcc.mutation.SetRed91to105Total(v)
	}
	if _, ok := tcc.mutation.Red91to105Percentage(); !ok {
		v := tscards.DefaultRed91to105Percentage
		tcc.mutation.SetRed91to105Percentage(v)
	}
	if _, ok := tcc.mutation.Red106To120Total(); !ok {
		v := tscards.DefaultRed106To120Total
		tcc.mutation.SetRed106To120Total(v)
	}
	if _, ok := tcc.mutation.Red106To120Percentage(); !ok {
		v := tscards.DefaultRed106To120Percentage
		tcc.mutation.SetRed106To120Percentage(v)
	}
	if _, ok := tcc.mutation.LastUpdated(); !ok {
		v := tscards.DefaultLastUpdated()
		tcc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcc *TSCardsCreate) check() error {
	if _, ok := tcc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`ent: missing required edge "TSCards.team"`)}
	}
	return nil
}

func (tcc *TSCardsCreate) sqlSave(ctx context.Context) (*TSCards, error) {
	if err := tcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tcc.mutation.id = &_node.ID
	tcc.mutation.done = true
	return _node, nil
}

func (tcc *TSCardsCreate) createSpec() (*TSCards, *sqlgraph.CreateSpec) {
	var (
		_node = &TSCards{config: tcc.config}
		_spec = sqlgraph.NewCreateSpec(tscards.Table, sqlgraph.NewFieldSpec(tscards.FieldID, field.TypeInt))
	)
	if value, ok := tcc.mutation.Yellow0To15Total(); ok {
		_spec.SetField(tscards.FieldYellow0To15Total, field.TypeInt, value)
		_node.Yellow0To15Total = value
	}
	if value, ok := tcc.mutation.Yellow0To15Percentage(); ok {
		_spec.SetField(tscards.FieldYellow0To15Percentage, field.TypeString, value)
		_node.Yellow0To15Percentage = value
	}
	if value, ok := tcc.mutation.Yellow16To30Total(); ok {
		_spec.SetField(tscards.FieldYellow16To30Total, field.TypeInt, value)
		_node.Yellow16To30Total = value
	}
	if value, ok := tcc.mutation.Yellow16To30Percentage(); ok {
		_spec.SetField(tscards.FieldYellow16To30Percentage, field.TypeString, value)
		_node.Yellow16To30Percentage = value
	}
	if value, ok := tcc.mutation.Yellow31To45Total(); ok {
		_spec.SetField(tscards.FieldYellow31To45Total, field.TypeInt, value)
		_node.Yellow31To45Total = value
	}
	if value, ok := tcc.mutation.Yellow31To45Percentage(); ok {
		_spec.SetField(tscards.FieldYellow31To45Percentage, field.TypeString, value)
		_node.Yellow31To45Percentage = value
	}
	if value, ok := tcc.mutation.Yellow46To60Total(); ok {
		_spec.SetField(tscards.FieldYellow46To60Total, field.TypeInt, value)
		_node.Yellow46To60Total = value
	}
	if value, ok := tcc.mutation.Yellow46To60Percentage(); ok {
		_spec.SetField(tscards.FieldYellow46To60Percentage, field.TypeString, value)
		_node.Yellow46To60Percentage = value
	}
	if value, ok := tcc.mutation.Yellow61To75Total(); ok {
		_spec.SetField(tscards.FieldYellow61To75Total, field.TypeInt, value)
		_node.Yellow61To75Total = value
	}
	if value, ok := tcc.mutation.Yellow61To75Percentage(); ok {
		_spec.SetField(tscards.FieldYellow61To75Percentage, field.TypeString, value)
		_node.Yellow61To75Percentage = value
	}
	if value, ok := tcc.mutation.Yellow76To90Total(); ok {
		_spec.SetField(tscards.FieldYellow76To90Total, field.TypeInt, value)
		_node.Yellow76To90Total = value
	}
	if value, ok := tcc.mutation.Yellow76To90Percentage(); ok {
		_spec.SetField(tscards.FieldYellow76To90Percentage, field.TypeString, value)
		_node.Yellow76To90Percentage = value
	}
	if value, ok := tcc.mutation.Yellow91to105Total(); ok {
		_spec.SetField(tscards.FieldYellow91to105Total, field.TypeInt, value)
		_node.Yellow91to105Total = value
	}
	if value, ok := tcc.mutation.Yellow91to105Percentage(); ok {
		_spec.SetField(tscards.FieldYellow91to105Percentage, field.TypeString, value)
		_node.Yellow91to105Percentage = value
	}
	if value, ok := tcc.mutation.Yellow106To120Total(); ok {
		_spec.SetField(tscards.FieldYellow106To120Total, field.TypeInt, value)
		_node.Yellow106To120Total = value
	}
	if value, ok := tcc.mutation.Yellow106To120Percentage(); ok {
		_spec.SetField(tscards.FieldYellow106To120Percentage, field.TypeString, value)
		_node.Yellow106To120Percentage = value
	}
	if value, ok := tcc.mutation.Red0To15Total(); ok {
		_spec.SetField(tscards.FieldRed0To15Total, field.TypeInt, value)
		_node.Red0To15Total = value
	}
	if value, ok := tcc.mutation.Red0To15Percentage(); ok {
		_spec.SetField(tscards.FieldRed0To15Percentage, field.TypeString, value)
		_node.Red0To15Percentage = value
	}
	if value, ok := tcc.mutation.Red16To30Total(); ok {
		_spec.SetField(tscards.FieldRed16To30Total, field.TypeInt, value)
		_node.Red16To30Total = value
	}
	if value, ok := tcc.mutation.Red16To30Percentage(); ok {
		_spec.SetField(tscards.FieldRed16To30Percentage, field.TypeString, value)
		_node.Red16To30Percentage = value
	}
	if value, ok := tcc.mutation.Red31To45Total(); ok {
		_spec.SetField(tscards.FieldRed31To45Total, field.TypeInt, value)
		_node.Red31To45Total = value
	}
	if value, ok := tcc.mutation.Red31To45Percentage(); ok {
		_spec.SetField(tscards.FieldRed31To45Percentage, field.TypeString, value)
		_node.Red31To45Percentage = value
	}
	if value, ok := tcc.mutation.Red46To60Total(); ok {
		_spec.SetField(tscards.FieldRed46To60Total, field.TypeInt, value)
		_node.Red46To60Total = value
	}
	if value, ok := tcc.mutation.Red46To60Percentage(); ok {
		_spec.SetField(tscards.FieldRed46To60Percentage, field.TypeString, value)
		_node.Red46To60Percentage = value
	}
	if value, ok := tcc.mutation.Red61To75Total(); ok {
		_spec.SetField(tscards.FieldRed61To75Total, field.TypeInt, value)
		_node.Red61To75Total = value
	}
	if value, ok := tcc.mutation.Red61To75Percentage(); ok {
		_spec.SetField(tscards.FieldRed61To75Percentage, field.TypeString, value)
		_node.Red61To75Percentage = value
	}
	if value, ok := tcc.mutation.Red76To90Total(); ok {
		_spec.SetField(tscards.FieldRed76To90Total, field.TypeInt, value)
		_node.Red76To90Total = value
	}
	if value, ok := tcc.mutation.Red76To90Percentage(); ok {
		_spec.SetField(tscards.FieldRed76To90Percentage, field.TypeString, value)
		_node.Red76To90Percentage = value
	}
	if value, ok := tcc.mutation.Red91to105Total(); ok {
		_spec.SetField(tscards.FieldRed91to105Total, field.TypeInt, value)
		_node.Red91to105Total = value
	}
	if value, ok := tcc.mutation.Red91to105Percentage(); ok {
		_spec.SetField(tscards.FieldRed91to105Percentage, field.TypeString, value)
		_node.Red91to105Percentage = value
	}
	if value, ok := tcc.mutation.Red106To120Total(); ok {
		_spec.SetField(tscards.FieldRed106To120Total, field.TypeInt, value)
		_node.Red106To120Total = value
	}
	if value, ok := tcc.mutation.Red106To120Percentage(); ok {
		_spec.SetField(tscards.FieldRed106To120Percentage, field.TypeString, value)
		_node.Red106To120Percentage = value
	}
	if value, ok := tcc.mutation.LastUpdated(); ok {
		_spec.SetField(tscards.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := tcc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   tscards.TeamTable,
			Columns: []string{tscards.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.team_cards_stats = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TSCardsCreateBulk is the builder for creating many TSCards entities in bulk.
type TSCardsCreateBulk struct {
	config
	builders []*TSCardsCreate
}

// Save creates the TSCards entities in the database.
func (tccb *TSCardsCreateBulk) Save(ctx context.Context) ([]*TSCards, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tccb.builders))
	nodes := make([]*TSCards, len(tccb.builders))
	mutators := make([]Mutator, len(tccb.builders))
	for i := range tccb.builders {
		func(i int, root context.Context) {
			builder := tccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TSCardsMutation)
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
					_, err = mutators[i+1].Mutate(root, tccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, tccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tccb *TSCardsCreateBulk) SaveX(ctx context.Context) []*TSCards {
	v, err := tccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tccb *TSCardsCreateBulk) Exec(ctx context.Context) error {
	_, err := tccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tccb *TSCardsCreateBulk) ExecX(ctx context.Context) {
	if err := tccb.Exec(ctx); err != nil {
		panic(err)
	}
}
