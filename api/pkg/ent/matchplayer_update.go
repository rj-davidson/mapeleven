// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixturelineups"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/matchplayer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MatchPlayerUpdate is the builder for updating MatchPlayer entities.
type MatchPlayerUpdate struct {
	config
	hooks    []Hook
	mutation *MatchPlayerMutation
}

// Where appends a list predicates to the MatchPlayerUpdate builder.
func (mpu *MatchPlayerUpdate) Where(ps ...predicate.MatchPlayer) *MatchPlayerUpdate {
	mpu.mutation.Where(ps...)
	return mpu
}

// SetNumber sets the "number" field.
func (mpu *MatchPlayerUpdate) SetNumber(i int) *MatchPlayerUpdate {
	mpu.mutation.ResetNumber()
	mpu.mutation.SetNumber(i)
	return mpu
}

// AddNumber adds i to the "number" field.
func (mpu *MatchPlayerUpdate) AddNumber(i int) *MatchPlayerUpdate {
	mpu.mutation.AddNumber(i)
	return mpu
}

// SetPosition sets the "position" field.
func (mpu *MatchPlayerUpdate) SetPosition(s string) *MatchPlayerUpdate {
	mpu.mutation.SetPosition(s)
	return mpu
}

// SetNillablePosition sets the "position" field if the given value is not nil.
func (mpu *MatchPlayerUpdate) SetNillablePosition(s *string) *MatchPlayerUpdate {
	if s != nil {
		mpu.SetPosition(*s)
	}
	return mpu
}

// ClearPosition clears the value of the "position" field.
func (mpu *MatchPlayerUpdate) ClearPosition() *MatchPlayerUpdate {
	mpu.mutation.ClearPosition()
	return mpu
}

// SetX sets the "x" field.
func (mpu *MatchPlayerUpdate) SetX(i int) *MatchPlayerUpdate {
	mpu.mutation.ResetX()
	mpu.mutation.SetX(i)
	return mpu
}

// SetNillableX sets the "x" field if the given value is not nil.
func (mpu *MatchPlayerUpdate) SetNillableX(i *int) *MatchPlayerUpdate {
	if i != nil {
		mpu.SetX(*i)
	}
	return mpu
}

// AddX adds i to the "x" field.
func (mpu *MatchPlayerUpdate) AddX(i int) *MatchPlayerUpdate {
	mpu.mutation.AddX(i)
	return mpu
}

// ClearX clears the value of the "x" field.
func (mpu *MatchPlayerUpdate) ClearX() *MatchPlayerUpdate {
	mpu.mutation.ClearX()
	return mpu
}

// SetY sets the "y" field.
func (mpu *MatchPlayerUpdate) SetY(i int) *MatchPlayerUpdate {
	mpu.mutation.ResetY()
	mpu.mutation.SetY(i)
	return mpu
}

// SetNillableY sets the "y" field if the given value is not nil.
func (mpu *MatchPlayerUpdate) SetNillableY(i *int) *MatchPlayerUpdate {
	if i != nil {
		mpu.SetY(*i)
	}
	return mpu
}

// AddY adds i to the "y" field.
func (mpu *MatchPlayerUpdate) AddY(i int) *MatchPlayerUpdate {
	mpu.mutation.AddY(i)
	return mpu
}

// ClearY clears the value of the "y" field.
func (mpu *MatchPlayerUpdate) ClearY() *MatchPlayerUpdate {
	mpu.mutation.ClearY()
	return mpu
}

// SetLastUpdated sets the "lastUpdated" field.
func (mpu *MatchPlayerUpdate) SetLastUpdated(t time.Time) *MatchPlayerUpdate {
	mpu.mutation.SetLastUpdated(t)
	return mpu
}

// ClearLastUpdated clears the value of the "lastUpdated" field.
func (mpu *MatchPlayerUpdate) ClearLastUpdated() *MatchPlayerUpdate {
	mpu.mutation.ClearLastUpdated()
	return mpu
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (mpu *MatchPlayerUpdate) SetPlayerID(id int) *MatchPlayerUpdate {
	mpu.mutation.SetPlayerID(id)
	return mpu
}

// SetPlayer sets the "player" edge to the Player entity.
func (mpu *MatchPlayerUpdate) SetPlayer(p *Player) *MatchPlayerUpdate {
	return mpu.SetPlayerID(p.ID)
}

// SetLineupID sets the "lineup" edge to the FixtureLineups entity by ID.
func (mpu *MatchPlayerUpdate) SetLineupID(id int) *MatchPlayerUpdate {
	mpu.mutation.SetLineupID(id)
	return mpu
}

// SetLineup sets the "lineup" edge to the FixtureLineups entity.
func (mpu *MatchPlayerUpdate) SetLineup(f *FixtureLineups) *MatchPlayerUpdate {
	return mpu.SetLineupID(f.ID)
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (mpu *MatchPlayerUpdate) SetPlayerStatsID(id int) *MatchPlayerUpdate {
	mpu.mutation.SetPlayerStatsID(id)
	return mpu
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (mpu *MatchPlayerUpdate) SetNillablePlayerStatsID(id *int) *MatchPlayerUpdate {
	if id != nil {
		mpu = mpu.SetPlayerStatsID(*id)
	}
	return mpu
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (mpu *MatchPlayerUpdate) SetPlayerStats(p *PlayerStats) *MatchPlayerUpdate {
	return mpu.SetPlayerStatsID(p.ID)
}

// Mutation returns the MatchPlayerMutation object of the builder.
func (mpu *MatchPlayerUpdate) Mutation() *MatchPlayerMutation {
	return mpu.mutation
}

// ClearPlayer clears the "player" edge to the Player entity.
func (mpu *MatchPlayerUpdate) ClearPlayer() *MatchPlayerUpdate {
	mpu.mutation.ClearPlayer()
	return mpu
}

// ClearLineup clears the "lineup" edge to the FixtureLineups entity.
func (mpu *MatchPlayerUpdate) ClearLineup() *MatchPlayerUpdate {
	mpu.mutation.ClearLineup()
	return mpu
}

// ClearPlayerStats clears the "playerStats" edge to the PlayerStats entity.
func (mpu *MatchPlayerUpdate) ClearPlayerStats() *MatchPlayerUpdate {
	mpu.mutation.ClearPlayerStats()
	return mpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mpu *MatchPlayerUpdate) Save(ctx context.Context) (int, error) {
	mpu.defaults()
	return withHooks(ctx, mpu.sqlSave, mpu.mutation, mpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpu *MatchPlayerUpdate) SaveX(ctx context.Context) int {
	affected, err := mpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mpu *MatchPlayerUpdate) Exec(ctx context.Context) error {
	_, err := mpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpu *MatchPlayerUpdate) ExecX(ctx context.Context) {
	if err := mpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpu *MatchPlayerUpdate) defaults() {
	if _, ok := mpu.mutation.LastUpdated(); !ok && !mpu.mutation.LastUpdatedCleared() {
		v := matchplayer.UpdateDefaultLastUpdated()
		mpu.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpu *MatchPlayerUpdate) check() error {
	if _, ok := mpu.mutation.PlayerID(); mpu.mutation.PlayerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MatchPlayer.player"`)
	}
	if _, ok := mpu.mutation.LineupID(); mpu.mutation.LineupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MatchPlayer.lineup"`)
	}
	return nil
}

func (mpu *MatchPlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(matchplayer.Table, matchplayer.Columns, sqlgraph.NewFieldSpec(matchplayer.FieldID, field.TypeInt))
	if ps := mpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpu.mutation.Number(); ok {
		_spec.SetField(matchplayer.FieldNumber, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.AddedNumber(); ok {
		_spec.AddField(matchplayer.FieldNumber, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.Position(); ok {
		_spec.SetField(matchplayer.FieldPosition, field.TypeString, value)
	}
	if mpu.mutation.PositionCleared() {
		_spec.ClearField(matchplayer.FieldPosition, field.TypeString)
	}
	if value, ok := mpu.mutation.X(); ok {
		_spec.SetField(matchplayer.FieldX, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.AddedX(); ok {
		_spec.AddField(matchplayer.FieldX, field.TypeInt, value)
	}
	if mpu.mutation.XCleared() {
		_spec.ClearField(matchplayer.FieldX, field.TypeInt)
	}
	if value, ok := mpu.mutation.Y(); ok {
		_spec.SetField(matchplayer.FieldY, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.AddedY(); ok {
		_spec.AddField(matchplayer.FieldY, field.TypeInt, value)
	}
	if mpu.mutation.YCleared() {
		_spec.ClearField(matchplayer.FieldY, field.TypeInt)
	}
	if value, ok := mpu.mutation.LastUpdated(); ok {
		_spec.SetField(matchplayer.FieldLastUpdated, field.TypeTime, value)
	}
	if mpu.mutation.LastUpdatedCleared() {
		_spec.ClearField(matchplayer.FieldLastUpdated, field.TypeTime)
	}
	if mpu.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.PlayerTable,
			Columns: []string{matchplayer.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpu.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.PlayerTable,
			Columns: []string{matchplayer.PlayerColumn},
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
	if mpu.mutation.LineupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.LineupTable,
			Columns: []string{matchplayer.LineupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixturelineups.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpu.mutation.LineupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.LineupTable,
			Columns: []string{matchplayer.LineupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixturelineups.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mpu.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.PlayerStatsTable,
			Columns: []string{matchplayer.PlayerStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpu.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.PlayerStatsTable,
			Columns: []string{matchplayer.PlayerStatsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{matchplayer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mpu.mutation.done = true
	return n, nil
}

// MatchPlayerUpdateOne is the builder for updating a single MatchPlayer entity.
type MatchPlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MatchPlayerMutation
}

// SetNumber sets the "number" field.
func (mpuo *MatchPlayerUpdateOne) SetNumber(i int) *MatchPlayerUpdateOne {
	mpuo.mutation.ResetNumber()
	mpuo.mutation.SetNumber(i)
	return mpuo
}

// AddNumber adds i to the "number" field.
func (mpuo *MatchPlayerUpdateOne) AddNumber(i int) *MatchPlayerUpdateOne {
	mpuo.mutation.AddNumber(i)
	return mpuo
}

// SetPosition sets the "position" field.
func (mpuo *MatchPlayerUpdateOne) SetPosition(s string) *MatchPlayerUpdateOne {
	mpuo.mutation.SetPosition(s)
	return mpuo
}

// SetNillablePosition sets the "position" field if the given value is not nil.
func (mpuo *MatchPlayerUpdateOne) SetNillablePosition(s *string) *MatchPlayerUpdateOne {
	if s != nil {
		mpuo.SetPosition(*s)
	}
	return mpuo
}

// ClearPosition clears the value of the "position" field.
func (mpuo *MatchPlayerUpdateOne) ClearPosition() *MatchPlayerUpdateOne {
	mpuo.mutation.ClearPosition()
	return mpuo
}

// SetX sets the "x" field.
func (mpuo *MatchPlayerUpdateOne) SetX(i int) *MatchPlayerUpdateOne {
	mpuo.mutation.ResetX()
	mpuo.mutation.SetX(i)
	return mpuo
}

// SetNillableX sets the "x" field if the given value is not nil.
func (mpuo *MatchPlayerUpdateOne) SetNillableX(i *int) *MatchPlayerUpdateOne {
	if i != nil {
		mpuo.SetX(*i)
	}
	return mpuo
}

// AddX adds i to the "x" field.
func (mpuo *MatchPlayerUpdateOne) AddX(i int) *MatchPlayerUpdateOne {
	mpuo.mutation.AddX(i)
	return mpuo
}

// ClearX clears the value of the "x" field.
func (mpuo *MatchPlayerUpdateOne) ClearX() *MatchPlayerUpdateOne {
	mpuo.mutation.ClearX()
	return mpuo
}

// SetY sets the "y" field.
func (mpuo *MatchPlayerUpdateOne) SetY(i int) *MatchPlayerUpdateOne {
	mpuo.mutation.ResetY()
	mpuo.mutation.SetY(i)
	return mpuo
}

// SetNillableY sets the "y" field if the given value is not nil.
func (mpuo *MatchPlayerUpdateOne) SetNillableY(i *int) *MatchPlayerUpdateOne {
	if i != nil {
		mpuo.SetY(*i)
	}
	return mpuo
}

// AddY adds i to the "y" field.
func (mpuo *MatchPlayerUpdateOne) AddY(i int) *MatchPlayerUpdateOne {
	mpuo.mutation.AddY(i)
	return mpuo
}

// ClearY clears the value of the "y" field.
func (mpuo *MatchPlayerUpdateOne) ClearY() *MatchPlayerUpdateOne {
	mpuo.mutation.ClearY()
	return mpuo
}

// SetLastUpdated sets the "lastUpdated" field.
func (mpuo *MatchPlayerUpdateOne) SetLastUpdated(t time.Time) *MatchPlayerUpdateOne {
	mpuo.mutation.SetLastUpdated(t)
	return mpuo
}

// ClearLastUpdated clears the value of the "lastUpdated" field.
func (mpuo *MatchPlayerUpdateOne) ClearLastUpdated() *MatchPlayerUpdateOne {
	mpuo.mutation.ClearLastUpdated()
	return mpuo
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (mpuo *MatchPlayerUpdateOne) SetPlayerID(id int) *MatchPlayerUpdateOne {
	mpuo.mutation.SetPlayerID(id)
	return mpuo
}

// SetPlayer sets the "player" edge to the Player entity.
func (mpuo *MatchPlayerUpdateOne) SetPlayer(p *Player) *MatchPlayerUpdateOne {
	return mpuo.SetPlayerID(p.ID)
}

// SetLineupID sets the "lineup" edge to the FixtureLineups entity by ID.
func (mpuo *MatchPlayerUpdateOne) SetLineupID(id int) *MatchPlayerUpdateOne {
	mpuo.mutation.SetLineupID(id)
	return mpuo
}

// SetLineup sets the "lineup" edge to the FixtureLineups entity.
func (mpuo *MatchPlayerUpdateOne) SetLineup(f *FixtureLineups) *MatchPlayerUpdateOne {
	return mpuo.SetLineupID(f.ID)
}

// SetPlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID.
func (mpuo *MatchPlayerUpdateOne) SetPlayerStatsID(id int) *MatchPlayerUpdateOne {
	mpuo.mutation.SetPlayerStatsID(id)
	return mpuo
}

// SetNillablePlayerStatsID sets the "playerStats" edge to the PlayerStats entity by ID if the given value is not nil.
func (mpuo *MatchPlayerUpdateOne) SetNillablePlayerStatsID(id *int) *MatchPlayerUpdateOne {
	if id != nil {
		mpuo = mpuo.SetPlayerStatsID(*id)
	}
	return mpuo
}

// SetPlayerStats sets the "playerStats" edge to the PlayerStats entity.
func (mpuo *MatchPlayerUpdateOne) SetPlayerStats(p *PlayerStats) *MatchPlayerUpdateOne {
	return mpuo.SetPlayerStatsID(p.ID)
}

// Mutation returns the MatchPlayerMutation object of the builder.
func (mpuo *MatchPlayerUpdateOne) Mutation() *MatchPlayerMutation {
	return mpuo.mutation
}

// ClearPlayer clears the "player" edge to the Player entity.
func (mpuo *MatchPlayerUpdateOne) ClearPlayer() *MatchPlayerUpdateOne {
	mpuo.mutation.ClearPlayer()
	return mpuo
}

// ClearLineup clears the "lineup" edge to the FixtureLineups entity.
func (mpuo *MatchPlayerUpdateOne) ClearLineup() *MatchPlayerUpdateOne {
	mpuo.mutation.ClearLineup()
	return mpuo
}

// ClearPlayerStats clears the "playerStats" edge to the PlayerStats entity.
func (mpuo *MatchPlayerUpdateOne) ClearPlayerStats() *MatchPlayerUpdateOne {
	mpuo.mutation.ClearPlayerStats()
	return mpuo
}

// Where appends a list predicates to the MatchPlayerUpdate builder.
func (mpuo *MatchPlayerUpdateOne) Where(ps ...predicate.MatchPlayer) *MatchPlayerUpdateOne {
	mpuo.mutation.Where(ps...)
	return mpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mpuo *MatchPlayerUpdateOne) Select(field string, fields ...string) *MatchPlayerUpdateOne {
	mpuo.fields = append([]string{field}, fields...)
	return mpuo
}

// Save executes the query and returns the updated MatchPlayer entity.
func (mpuo *MatchPlayerUpdateOne) Save(ctx context.Context) (*MatchPlayer, error) {
	mpuo.defaults()
	return withHooks(ctx, mpuo.sqlSave, mpuo.mutation, mpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpuo *MatchPlayerUpdateOne) SaveX(ctx context.Context) *MatchPlayer {
	node, err := mpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mpuo *MatchPlayerUpdateOne) Exec(ctx context.Context) error {
	_, err := mpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpuo *MatchPlayerUpdateOne) ExecX(ctx context.Context) {
	if err := mpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpuo *MatchPlayerUpdateOne) defaults() {
	if _, ok := mpuo.mutation.LastUpdated(); !ok && !mpuo.mutation.LastUpdatedCleared() {
		v := matchplayer.UpdateDefaultLastUpdated()
		mpuo.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpuo *MatchPlayerUpdateOne) check() error {
	if _, ok := mpuo.mutation.PlayerID(); mpuo.mutation.PlayerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MatchPlayer.player"`)
	}
	if _, ok := mpuo.mutation.LineupID(); mpuo.mutation.LineupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MatchPlayer.lineup"`)
	}
	return nil
}

func (mpuo *MatchPlayerUpdateOne) sqlSave(ctx context.Context) (_node *MatchPlayer, err error) {
	if err := mpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(matchplayer.Table, matchplayer.Columns, sqlgraph.NewFieldSpec(matchplayer.FieldID, field.TypeInt))
	id, ok := mpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MatchPlayer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, matchplayer.FieldID)
		for _, f := range fields {
			if !matchplayer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != matchplayer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpuo.mutation.Number(); ok {
		_spec.SetField(matchplayer.FieldNumber, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.AddedNumber(); ok {
		_spec.AddField(matchplayer.FieldNumber, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.Position(); ok {
		_spec.SetField(matchplayer.FieldPosition, field.TypeString, value)
	}
	if mpuo.mutation.PositionCleared() {
		_spec.ClearField(matchplayer.FieldPosition, field.TypeString)
	}
	if value, ok := mpuo.mutation.X(); ok {
		_spec.SetField(matchplayer.FieldX, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.AddedX(); ok {
		_spec.AddField(matchplayer.FieldX, field.TypeInt, value)
	}
	if mpuo.mutation.XCleared() {
		_spec.ClearField(matchplayer.FieldX, field.TypeInt)
	}
	if value, ok := mpuo.mutation.Y(); ok {
		_spec.SetField(matchplayer.FieldY, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.AddedY(); ok {
		_spec.AddField(matchplayer.FieldY, field.TypeInt, value)
	}
	if mpuo.mutation.YCleared() {
		_spec.ClearField(matchplayer.FieldY, field.TypeInt)
	}
	if value, ok := mpuo.mutation.LastUpdated(); ok {
		_spec.SetField(matchplayer.FieldLastUpdated, field.TypeTime, value)
	}
	if mpuo.mutation.LastUpdatedCleared() {
		_spec.ClearField(matchplayer.FieldLastUpdated, field.TypeTime)
	}
	if mpuo.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.PlayerTable,
			Columns: []string{matchplayer.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpuo.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.PlayerTable,
			Columns: []string{matchplayer.PlayerColumn},
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
	if mpuo.mutation.LineupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.LineupTable,
			Columns: []string{matchplayer.LineupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixturelineups.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpuo.mutation.LineupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.LineupTable,
			Columns: []string{matchplayer.LineupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixturelineups.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mpuo.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.PlayerStatsTable,
			Columns: []string{matchplayer.PlayerStatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpuo.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchplayer.PlayerStatsTable,
			Columns: []string{matchplayer.PlayerStatsColumn},
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
	_node = &MatchPlayer{config: mpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{matchplayer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mpuo.mutation.done = true
	return _node, nil
}
