// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/club"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/league"
	"mapeleven/db/ent/player"
	"mapeleven/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CountryUpdate is the builder for updating Country entities.
type CountryUpdate struct {
	config
	hooks    []Hook
	mutation *CountryMutation
}

// Where appends a list predicates to the CountryUpdate builder.
func (cu *CountryUpdate) Where(ps ...predicate.Country) *CountryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCode sets the "code" field.
func (cu *CountryUpdate) SetCode(s string) *CountryUpdate {
	cu.mutation.SetCode(s)
	return cu
}

// SetName sets the "name" field.
func (cu *CountryUpdate) SetName(s string) *CountryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetFlag sets the "flag" field.
func (cu *CountryUpdate) SetFlag(s string) *CountryUpdate {
	cu.mutation.SetFlag(s)
	return cu
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (cu *CountryUpdate) AddPlayerIDs(ids ...int) *CountryUpdate {
	cu.mutation.AddPlayerIDs(ids...)
	return cu
}

// AddPlayers adds the "players" edges to the Player entity.
func (cu *CountryUpdate) AddPlayers(p ...*Player) *CountryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPlayerIDs(ids...)
}

// AddLeagueIDs adds the "leagues" edge to the League entity by IDs.
func (cu *CountryUpdate) AddLeagueIDs(ids ...int) *CountryUpdate {
	cu.mutation.AddLeagueIDs(ids...)
	return cu
}

// AddLeagues adds the "leagues" edges to the League entity.
func (cu *CountryUpdate) AddLeagues(l ...*League) *CountryUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cu.AddLeagueIDs(ids...)
}

// AddClubIDs adds the "clubs" edge to the Club entity by IDs.
func (cu *CountryUpdate) AddClubIDs(ids ...int) *CountryUpdate {
	cu.mutation.AddClubIDs(ids...)
	return cu
}

// AddClubs adds the "clubs" edges to the Club entity.
func (cu *CountryUpdate) AddClubs(c ...*Club) *CountryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddClubIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cu *CountryUpdate) Mutation() *CountryMutation {
	return cu.mutation
}

// ClearPlayers clears all "players" edges to the Player entity.
func (cu *CountryUpdate) ClearPlayers() *CountryUpdate {
	cu.mutation.ClearPlayers()
	return cu
}

// RemovePlayerIDs removes the "players" edge to Player entities by IDs.
func (cu *CountryUpdate) RemovePlayerIDs(ids ...int) *CountryUpdate {
	cu.mutation.RemovePlayerIDs(ids...)
	return cu
}

// RemovePlayers removes "players" edges to Player entities.
func (cu *CountryUpdate) RemovePlayers(p ...*Player) *CountryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePlayerIDs(ids...)
}

// ClearLeagues clears all "leagues" edges to the League entity.
func (cu *CountryUpdate) ClearLeagues() *CountryUpdate {
	cu.mutation.ClearLeagues()
	return cu
}

// RemoveLeagueIDs removes the "leagues" edge to League entities by IDs.
func (cu *CountryUpdate) RemoveLeagueIDs(ids ...int) *CountryUpdate {
	cu.mutation.RemoveLeagueIDs(ids...)
	return cu
}

// RemoveLeagues removes "leagues" edges to League entities.
func (cu *CountryUpdate) RemoveLeagues(l ...*League) *CountryUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cu.RemoveLeagueIDs(ids...)
}

// ClearClubs clears all "clubs" edges to the Club entity.
func (cu *CountryUpdate) ClearClubs() *CountryUpdate {
	cu.mutation.ClearClubs()
	return cu
}

// RemoveClubIDs removes the "clubs" edge to Club entities by IDs.
func (cu *CountryUpdate) RemoveClubIDs(ids ...int) *CountryUpdate {
	cu.mutation.RemoveClubIDs(ids...)
	return cu
}

// RemoveClubs removes "clubs" edges to Club entities.
func (cu *CountryUpdate) RemoveClubs(c ...*Club) *CountryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveClubIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CountryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CountryMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CountryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CountryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CountryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CountryUpdate) check() error {
	if v, ok := cu.mutation.Code(); ok {
		if err := country.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Country.code": %w`, err)}
		}
	}
	return nil
}

func (cu *CountryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Code(); ok {
		_spec.SetField(country.FieldCode, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Flag(); ok {
		_spec.SetField(country.FieldFlag, field.TypeString, value)
	}
	if cu.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.PlayersTable,
			Columns: []string{country.PlayersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPlayersIDs(); len(nodes) > 0 && !cu.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.PlayersTable,
			Columns: []string{country.PlayersColumn},
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
	if nodes := cu.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.PlayersTable,
			Columns: []string{country.PlayersColumn},
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
	if cu.mutation.LeaguesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LeaguesTable,
			Columns: []string{country.LeaguesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedLeaguesIDs(); len(nodes) > 0 && !cu.mutation.LeaguesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LeaguesTable,
			Columns: []string{country.LeaguesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LeaguesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LeaguesTable,
			Columns: []string{country.LeaguesColumn},
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
	if cu.mutation.ClubsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.ClubsTable,
			Columns: []string{country.ClubsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(club.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedClubsIDs(); len(nodes) > 0 && !cu.mutation.ClubsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.ClubsTable,
			Columns: []string{country.ClubsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(club.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClubsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.ClubsTable,
			Columns: []string{country.ClubsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(club.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CountryUpdateOne is the builder for updating a single Country entity.
type CountryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CountryMutation
}

// SetCode sets the "code" field.
func (cuo *CountryUpdateOne) SetCode(s string) *CountryUpdateOne {
	cuo.mutation.SetCode(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CountryUpdateOne) SetName(s string) *CountryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetFlag sets the "flag" field.
func (cuo *CountryUpdateOne) SetFlag(s string) *CountryUpdateOne {
	cuo.mutation.SetFlag(s)
	return cuo
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (cuo *CountryUpdateOne) AddPlayerIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.AddPlayerIDs(ids...)
	return cuo
}

// AddPlayers adds the "players" edges to the Player entity.
func (cuo *CountryUpdateOne) AddPlayers(p ...*Player) *CountryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPlayerIDs(ids...)
}

// AddLeagueIDs adds the "leagues" edge to the League entity by IDs.
func (cuo *CountryUpdateOne) AddLeagueIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.AddLeagueIDs(ids...)
	return cuo
}

// AddLeagues adds the "leagues" edges to the League entity.
func (cuo *CountryUpdateOne) AddLeagues(l ...*League) *CountryUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cuo.AddLeagueIDs(ids...)
}

// AddClubIDs adds the "clubs" edge to the Club entity by IDs.
func (cuo *CountryUpdateOne) AddClubIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.AddClubIDs(ids...)
	return cuo
}

// AddClubs adds the "clubs" edges to the Club entity.
func (cuo *CountryUpdateOne) AddClubs(c ...*Club) *CountryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddClubIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cuo *CountryUpdateOne) Mutation() *CountryMutation {
	return cuo.mutation
}

// ClearPlayers clears all "players" edges to the Player entity.
func (cuo *CountryUpdateOne) ClearPlayers() *CountryUpdateOne {
	cuo.mutation.ClearPlayers()
	return cuo
}

// RemovePlayerIDs removes the "players" edge to Player entities by IDs.
func (cuo *CountryUpdateOne) RemovePlayerIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.RemovePlayerIDs(ids...)
	return cuo
}

// RemovePlayers removes "players" edges to Player entities.
func (cuo *CountryUpdateOne) RemovePlayers(p ...*Player) *CountryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePlayerIDs(ids...)
}

// ClearLeagues clears all "leagues" edges to the League entity.
func (cuo *CountryUpdateOne) ClearLeagues() *CountryUpdateOne {
	cuo.mutation.ClearLeagues()
	return cuo
}

// RemoveLeagueIDs removes the "leagues" edge to League entities by IDs.
func (cuo *CountryUpdateOne) RemoveLeagueIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.RemoveLeagueIDs(ids...)
	return cuo
}

// RemoveLeagues removes "leagues" edges to League entities.
func (cuo *CountryUpdateOne) RemoveLeagues(l ...*League) *CountryUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cuo.RemoveLeagueIDs(ids...)
}

// ClearClubs clears all "clubs" edges to the Club entity.
func (cuo *CountryUpdateOne) ClearClubs() *CountryUpdateOne {
	cuo.mutation.ClearClubs()
	return cuo
}

// RemoveClubIDs removes the "clubs" edge to Club entities by IDs.
func (cuo *CountryUpdateOne) RemoveClubIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.RemoveClubIDs(ids...)
	return cuo
}

// RemoveClubs removes "clubs" edges to Club entities.
func (cuo *CountryUpdateOne) RemoveClubs(c ...*Club) *CountryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveClubIDs(ids...)
}

// Where appends a list predicates to the CountryUpdate builder.
func (cuo *CountryUpdateOne) Where(ps ...predicate.Country) *CountryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CountryUpdateOne) Select(field string, fields ...string) *CountryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Country entity.
func (cuo *CountryUpdateOne) Save(ctx context.Context) (*Country, error) {
	return withHooks[*Country, CountryMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CountryUpdateOne) SaveX(ctx context.Context) *Country {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CountryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CountryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CountryUpdateOne) check() error {
	if v, ok := cuo.mutation.Code(); ok {
		if err := country.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Country.code": %w`, err)}
		}
	}
	return nil
}

func (cuo *CountryUpdateOne) sqlSave(ctx context.Context) (_node *Country, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Country.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, country.FieldID)
		for _, f := range fields {
			if !country.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != country.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Code(); ok {
		_spec.SetField(country.FieldCode, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Flag(); ok {
		_spec.SetField(country.FieldFlag, field.TypeString, value)
	}
	if cuo.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.PlayersTable,
			Columns: []string{country.PlayersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPlayersIDs(); len(nodes) > 0 && !cuo.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.PlayersTable,
			Columns: []string{country.PlayersColumn},
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
	if nodes := cuo.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.PlayersTable,
			Columns: []string{country.PlayersColumn},
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
	if cuo.mutation.LeaguesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LeaguesTable,
			Columns: []string{country.LeaguesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedLeaguesIDs(); len(nodes) > 0 && !cuo.mutation.LeaguesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LeaguesTable,
			Columns: []string{country.LeaguesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LeaguesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.LeaguesTable,
			Columns: []string{country.LeaguesColumn},
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
	if cuo.mutation.ClubsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.ClubsTable,
			Columns: []string{country.ClubsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(club.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedClubsIDs(); len(nodes) > 0 && !cuo.mutation.ClubsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.ClubsTable,
			Columns: []string{country.ClubsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(club.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClubsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.ClubsTable,
			Columns: []string{country.ClubsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(club.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Country{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
