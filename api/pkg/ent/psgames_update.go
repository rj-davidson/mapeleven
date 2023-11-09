// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSGamesUpdate is the builder for updating PSGames entities.
type PSGamesUpdate struct {
	config
	hooks    []Hook
	mutation *PSGamesMutation
}

// Where appends a list predicates to the PSGamesUpdate builder.
func (pgu *PSGamesUpdate) Where(ps ...predicate.PSGames) *PSGamesUpdate {
	pgu.mutation.Where(ps...)
	return pgu
}

// SetAppearances sets the "Appearances" field.
func (pgu *PSGamesUpdate) SetAppearances(i int) *PSGamesUpdate {
	pgu.mutation.ResetAppearances()
	pgu.mutation.SetAppearances(i)
	return pgu
}

// AddAppearances adds i to the "Appearances" field.
func (pgu *PSGamesUpdate) AddAppearances(i int) *PSGamesUpdate {
	pgu.mutation.AddAppearances(i)
	return pgu
}

// SetLineups sets the "Lineups" field.
func (pgu *PSGamesUpdate) SetLineups(i int) *PSGamesUpdate {
	pgu.mutation.ResetLineups()
	pgu.mutation.SetLineups(i)
	return pgu
}

// AddLineups adds i to the "Lineups" field.
func (pgu *PSGamesUpdate) AddLineups(i int) *PSGamesUpdate {
	pgu.mutation.AddLineups(i)
	return pgu
}

// SetMinutes sets the "Minutes" field.
func (pgu *PSGamesUpdate) SetMinutes(i int) *PSGamesUpdate {
	pgu.mutation.ResetMinutes()
	pgu.mutation.SetMinutes(i)
	return pgu
}

// AddMinutes adds i to the "Minutes" field.
func (pgu *PSGamesUpdate) AddMinutes(i int) *PSGamesUpdate {
	pgu.mutation.AddMinutes(i)
	return pgu
}

// SetNumber sets the "Number" field.
func (pgu *PSGamesUpdate) SetNumber(i int) *PSGamesUpdate {
	pgu.mutation.ResetNumber()
	pgu.mutation.SetNumber(i)
	return pgu
}

// SetNillableNumber sets the "Number" field if the given value is not nil.
func (pgu *PSGamesUpdate) SetNillableNumber(i *int) *PSGamesUpdate {
	if i != nil {
		pgu.SetNumber(*i)
	}
	return pgu
}

// AddNumber adds i to the "Number" field.
func (pgu *PSGamesUpdate) AddNumber(i int) *PSGamesUpdate {
	pgu.mutation.AddNumber(i)
	return pgu
}

// SetPosition sets the "Position" field.
func (pgu *PSGamesUpdate) SetPosition(s string) *PSGamesUpdate {
	pgu.mutation.SetPosition(s)
	return pgu
}

// SetNillablePosition sets the "Position" field if the given value is not nil.
func (pgu *PSGamesUpdate) SetNillablePosition(s *string) *PSGamesUpdate {
	if s != nil {
		pgu.SetPosition(*s)
	}
	return pgu
}

// SetRating sets the "Rating" field.
func (pgu *PSGamesUpdate) SetRating(s string) *PSGamesUpdate {
	pgu.mutation.SetRating(s)
	return pgu
}

// SetNillableRating sets the "Rating" field if the given value is not nil.
func (pgu *PSGamesUpdate) SetNillableRating(s *string) *PSGamesUpdate {
	if s != nil {
		pgu.SetRating(*s)
	}
	return pgu
}

// SetCaptain sets the "Captain" field.
func (pgu *PSGamesUpdate) SetCaptain(b bool) *PSGamesUpdate {
	pgu.mutation.SetCaptain(b)
	return pgu
}

// SetNillableCaptain sets the "Captain" field if the given value is not nil.
func (pgu *PSGamesUpdate) SetNillableCaptain(b *bool) *PSGamesUpdate {
	if b != nil {
		pgu.SetCaptain(*b)
	}
	return pgu
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (pgu *PSGamesUpdate) SetPlayerStatsID(id int) *PSGamesUpdate {
	pgu.mutation.SetPlayerStatsID(id)
	return pgu
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (pgu *PSGamesUpdate) SetNillablePlayerStatsID(id *int) *PSGamesUpdate {
	if id != nil {
		pgu = pgu.SetPlayerStatsID(*id)
	}
	return pgu
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (pgu *PSGamesUpdate) SetPlayerStats(p *PlayerStats) *PSGamesUpdate {
	return pgu.SetPlayerStatsID(p.ID)
}

// Mutation returns the PSGamesMutation object of the builder.
func (pgu *PSGamesUpdate) Mutation() *PSGamesMutation {
	return pgu.mutation
}

// ClearPlayerStats clears the "playerStats" edge to the PlayerStats entity.
func (pgu *PSGamesUpdate) ClearPlayerStats() *PSGamesUpdate {
	pgu.mutation.ClearPlayerStats()
	return pgu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pgu *PSGamesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pgu.sqlSave, pgu.mutation, pgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pgu *PSGamesUpdate) SaveX(ctx context.Context) int {
	affected, err := pgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pgu *PSGamesUpdate) Exec(ctx context.Context) error {
	_, err := pgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pgu *PSGamesUpdate) ExecX(ctx context.Context) {
	if err := pgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pgu *PSGamesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(psgames.Table, psgames.Columns, sqlgraph.NewFieldSpec(psgames.FieldID, field.TypeInt))
	if ps := pgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pgu.mutation.Appearances(); ok {
		_spec.SetField(psgames.FieldAppearances, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedAppearances(); ok {
		_spec.AddField(psgames.FieldAppearances, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.Lineups(); ok {
		_spec.SetField(psgames.FieldLineups, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedLineups(); ok {
		_spec.AddField(psgames.FieldLineups, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.Minutes(); ok {
		_spec.SetField(psgames.FieldMinutes, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedMinutes(); ok {
		_spec.AddField(psgames.FieldMinutes, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.Number(); ok {
		_spec.SetField(psgames.FieldNumber, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedNumber(); ok {
		_spec.AddField(psgames.FieldNumber, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.Position(); ok {
		_spec.SetField(psgames.FieldPosition, field.TypeString, value)
	}
	if value, ok := pgu.mutation.Rating(); ok {
		_spec.SetField(psgames.FieldRating, field.TypeString, value)
	}
	if value, ok := pgu.mutation.Captain(); ok {
		_spec.SetField(psgames.FieldCaptain, field.TypeBool, value)
	}
	if pgu.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   psgames.PlayerStatsTable,
			Columns: []string{psgames.PlayerStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   psgames.PlayerStatsTable,
			Columns: []string{psgames.PlayerStatsColumn},
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
			err = &NotFoundError{psgames.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pgu.mutation.done = true
	return n, nil
}

// PSGamesUpdateOne is the builder for updating a single PSGames entity.
type PSGamesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PSGamesMutation
}

// SetAppearances sets the "Appearances" field.
func (pguo *PSGamesUpdateOne) SetAppearances(i int) *PSGamesUpdateOne {
	pguo.mutation.ResetAppearances()
	pguo.mutation.SetAppearances(i)
	return pguo
}

// AddAppearances adds i to the "Appearances" field.
func (pguo *PSGamesUpdateOne) AddAppearances(i int) *PSGamesUpdateOne {
	pguo.mutation.AddAppearances(i)
	return pguo
}

// SetLineups sets the "Lineups" field.
func (pguo *PSGamesUpdateOne) SetLineups(i int) *PSGamesUpdateOne {
	pguo.mutation.ResetLineups()
	pguo.mutation.SetLineups(i)
	return pguo
}

// AddLineups adds i to the "Lineups" field.
func (pguo *PSGamesUpdateOne) AddLineups(i int) *PSGamesUpdateOne {
	pguo.mutation.AddLineups(i)
	return pguo
}

// SetMinutes sets the "Minutes" field.
func (pguo *PSGamesUpdateOne) SetMinutes(i int) *PSGamesUpdateOne {
	pguo.mutation.ResetMinutes()
	pguo.mutation.SetMinutes(i)
	return pguo
}

// AddMinutes adds i to the "Minutes" field.
func (pguo *PSGamesUpdateOne) AddMinutes(i int) *PSGamesUpdateOne {
	pguo.mutation.AddMinutes(i)
	return pguo
}

// SetNumber sets the "Number" field.
func (pguo *PSGamesUpdateOne) SetNumber(i int) *PSGamesUpdateOne {
	pguo.mutation.ResetNumber()
	pguo.mutation.SetNumber(i)
	return pguo
}

// SetNillableNumber sets the "Number" field if the given value is not nil.
func (pguo *PSGamesUpdateOne) SetNillableNumber(i *int) *PSGamesUpdateOne {
	if i != nil {
		pguo.SetNumber(*i)
	}
	return pguo
}

// AddNumber adds i to the "Number" field.
func (pguo *PSGamesUpdateOne) AddNumber(i int) *PSGamesUpdateOne {
	pguo.mutation.AddNumber(i)
	return pguo
}

// SetPosition sets the "Position" field.
func (pguo *PSGamesUpdateOne) SetPosition(s string) *PSGamesUpdateOne {
	pguo.mutation.SetPosition(s)
	return pguo
}

// SetNillablePosition sets the "Position" field if the given value is not nil.
func (pguo *PSGamesUpdateOne) SetNillablePosition(s *string) *PSGamesUpdateOne {
	if s != nil {
		pguo.SetPosition(*s)
	}
	return pguo
}

// SetRating sets the "Rating" field.
func (pguo *PSGamesUpdateOne) SetRating(s string) *PSGamesUpdateOne {
	pguo.mutation.SetRating(s)
	return pguo
}

// SetNillableRating sets the "Rating" field if the given value is not nil.
func (pguo *PSGamesUpdateOne) SetNillableRating(s *string) *PSGamesUpdateOne {
	if s != nil {
		pguo.SetRating(*s)
	}
	return pguo
}

// SetCaptain sets the "Captain" field.
func (pguo *PSGamesUpdateOne) SetCaptain(b bool) *PSGamesUpdateOne {
	pguo.mutation.SetCaptain(b)
	return pguo
}

// SetNillableCaptain sets the "Captain" field if the given value is not nil.
func (pguo *PSGamesUpdateOne) SetNillableCaptain(b *bool) *PSGamesUpdateOne {
	if b != nil {
		pguo.SetCaptain(*b)
	}
	return pguo
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (pguo *PSGamesUpdateOne) SetPlayerStatsID(id int) *PSGamesUpdateOne {
	pguo.mutation.SetPlayerStatsID(id)
	return pguo
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (pguo *PSGamesUpdateOne) SetNillablePlayerStatsID(id *int) *PSGamesUpdateOne {
	if id != nil {
		pguo = pguo.SetPlayerStatsID(*id)
	}
	return pguo
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (pguo *PSGamesUpdateOne) SetPlayerStats(p *PlayerStats) *PSGamesUpdateOne {
	return pguo.SetPlayerStatsID(p.ID)
}

// Mutation returns the PSGamesMutation object of the builder.
func (pguo *PSGamesUpdateOne) Mutation() *PSGamesMutation {
	return pguo.mutation
}

// ClearPlayerStats clears the "playerStats" edge to the PlayerStats entity.
func (pguo *PSGamesUpdateOne) ClearPlayerStats() *PSGamesUpdateOne {
	pguo.mutation.ClearPlayerStats()
	return pguo
}

// Where appends a list predicates to the PSGamesUpdate builder.
func (pguo *PSGamesUpdateOne) Where(ps ...predicate.PSGames) *PSGamesUpdateOne {
	pguo.mutation.Where(ps...)
	return pguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pguo *PSGamesUpdateOne) Select(field string, fields ...string) *PSGamesUpdateOne {
	pguo.fields = append([]string{field}, fields...)
	return pguo
}

// Save executes the query and returns the updated PSGames entity.
func (pguo *PSGamesUpdateOne) Save(ctx context.Context) (*PSGames, error) {
	return withHooks(ctx, pguo.sqlSave, pguo.mutation, pguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pguo *PSGamesUpdateOne) SaveX(ctx context.Context) *PSGames {
	node, err := pguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pguo *PSGamesUpdateOne) Exec(ctx context.Context) error {
	_, err := pguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pguo *PSGamesUpdateOne) ExecX(ctx context.Context) {
	if err := pguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pguo *PSGamesUpdateOne) sqlSave(ctx context.Context) (_node *PSGames, err error) {
	_spec := sqlgraph.NewUpdateSpec(psgames.Table, psgames.Columns, sqlgraph.NewFieldSpec(psgames.FieldID, field.TypeInt))
	id, ok := pguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PSGames.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, psgames.FieldID)
		for _, f := range fields {
			if !psgames.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != psgames.FieldID {
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
	if value, ok := pguo.mutation.Appearances(); ok {
		_spec.SetField(psgames.FieldAppearances, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedAppearances(); ok {
		_spec.AddField(psgames.FieldAppearances, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.Lineups(); ok {
		_spec.SetField(psgames.FieldLineups, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedLineups(); ok {
		_spec.AddField(psgames.FieldLineups, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.Minutes(); ok {
		_spec.SetField(psgames.FieldMinutes, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedMinutes(); ok {
		_spec.AddField(psgames.FieldMinutes, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.Number(); ok {
		_spec.SetField(psgames.FieldNumber, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedNumber(); ok {
		_spec.AddField(psgames.FieldNumber, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.Position(); ok {
		_spec.SetField(psgames.FieldPosition, field.TypeString, value)
	}
	if value, ok := pguo.mutation.Rating(); ok {
		_spec.SetField(psgames.FieldRating, field.TypeString, value)
	}
	if value, ok := pguo.mutation.Captain(); ok {
		_spec.SetField(psgames.FieldCaptain, field.TypeBool, value)
	}
	if pguo.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   psgames.PlayerStatsTable,
			Columns: []string{psgames.PlayerStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   psgames.PlayerStatsTable,
			Columns: []string{psgames.PlayerStatsColumn},
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
	_node = &PSGames{config: pguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{psgames.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pguo.mutation.done = true
	return _node, nil
}
