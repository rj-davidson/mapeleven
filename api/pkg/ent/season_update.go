// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixture"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/league"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/squad"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/standings"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SeasonUpdate is the builder for updating Season entities.
type SeasonUpdate struct {
	config
	hooks    []Hook
	mutation *SeasonMutation
}

// Where appends a list predicates to the SeasonUpdate builder.
func (su *SeasonUpdate) Where(ps ...predicate.Season) *SeasonUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetStartDate sets the "start_date" field.
func (su *SeasonUpdate) SetStartDate(t time.Time) *SeasonUpdate {
	su.mutation.SetStartDate(t)
	return su
}

// SetEndDate sets the "end_date" field.
func (su *SeasonUpdate) SetEndDate(t time.Time) *SeasonUpdate {
	su.mutation.SetEndDate(t)
	return su
}

// SetCurrent sets the "current" field.
func (su *SeasonUpdate) SetCurrent(b bool) *SeasonUpdate {
	su.mutation.SetCurrent(b)
	return su
}

// SetLastUpdated sets the "lastUpdated" field.
func (su *SeasonUpdate) SetLastUpdated(t time.Time) *SeasonUpdate {
	su.mutation.SetLastUpdated(t)
	return su
}

// ClearLastUpdated clears the value of the "lastUpdated" field.
func (su *SeasonUpdate) ClearLastUpdated() *SeasonUpdate {
	su.mutation.ClearLastUpdated()
	return su
}

// SetLeagueID sets the "league" edge to the League entity by ID.
func (su *SeasonUpdate) SetLeagueID(id int) *SeasonUpdate {
	su.mutation.SetLeagueID(id)
	return su
}

// SetNillableLeagueID sets the "league" edge to the League entity by ID if the given value is not nil.
func (su *SeasonUpdate) SetNillableLeagueID(id *int) *SeasonUpdate {
	if id != nil {
		su = su.SetLeagueID(*id)
	}
	return su
}

// SetLeague sets the "league" edge to the League entity.
func (su *SeasonUpdate) SetLeague(l *League) *SeasonUpdate {
	return su.SetLeagueID(l.ID)
}

// AddFixtureIDs adds the "fixtures" edge to the Fixture entity by IDs.
func (su *SeasonUpdate) AddFixtureIDs(ids ...int) *SeasonUpdate {
	su.mutation.AddFixtureIDs(ids...)
	return su
}

// AddFixtures adds the "fixtures" edges to the Fixture entity.
func (su *SeasonUpdate) AddFixtures(f ...*Fixture) *SeasonUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.AddFixtureIDs(ids...)
}

// AddStandingIDs adds the "standings" edge to the Standings entity by IDs.
func (su *SeasonUpdate) AddStandingIDs(ids ...int) *SeasonUpdate {
	su.mutation.AddStandingIDs(ids...)
	return su
}

// AddStandings adds the "standings" edges to the Standings entity.
func (su *SeasonUpdate) AddStandings(s ...*Standings) *SeasonUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddStandingIDs(ids...)
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (su *SeasonUpdate) AddTeamIDs(ids ...int) *SeasonUpdate {
	su.mutation.AddTeamIDs(ids...)
	return su
}

// AddTeams adds the "teams" edges to the Team entity.
func (su *SeasonUpdate) AddTeams(t ...*Team) *SeasonUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTeamIDs(ids...)
}

// AddSquadIDs adds the "squad" edge to the Squad entity by IDs.
func (su *SeasonUpdate) AddSquadIDs(ids ...int) *SeasonUpdate {
	su.mutation.AddSquadIDs(ids...)
	return su
}

// AddSquad adds the "squad" edges to the Squad entity.
func (su *SeasonUpdate) AddSquad(s ...*Squad) *SeasonUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSquadIDs(ids...)
}

// AddPlayerStatIDs adds the "playerStats" edge to the PlayerStats entity by IDs.
func (su *SeasonUpdate) AddPlayerStatIDs(ids ...int) *SeasonUpdate {
	su.mutation.AddPlayerStatIDs(ids...)
	return su
}

// AddPlayerStats adds the "playerStats" edges to the PlayerStats entity.
func (su *SeasonUpdate) AddPlayerStats(p ...*PlayerStats) *SeasonUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddPlayerStatIDs(ids...)
}

// Mutation returns the SeasonMutation object of the builder.
func (su *SeasonUpdate) Mutation() *SeasonMutation {
	return su.mutation
}

// ClearLeague clears the "league" edge to the League entity.
func (su *SeasonUpdate) ClearLeague() *SeasonUpdate {
	su.mutation.ClearLeague()
	return su
}

// ClearFixtures clears all "fixtures" edges to the Fixture entity.
func (su *SeasonUpdate) ClearFixtures() *SeasonUpdate {
	su.mutation.ClearFixtures()
	return su
}

// RemoveFixtureIDs removes the "fixtures" edge to Fixture entities by IDs.
func (su *SeasonUpdate) RemoveFixtureIDs(ids ...int) *SeasonUpdate {
	su.mutation.RemoveFixtureIDs(ids...)
	return su
}

// RemoveFixtures removes "fixtures" edges to Fixture entities.
func (su *SeasonUpdate) RemoveFixtures(f ...*Fixture) *SeasonUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.RemoveFixtureIDs(ids...)
}

// ClearStandings clears all "standings" edges to the Standings entity.
func (su *SeasonUpdate) ClearStandings() *SeasonUpdate {
	su.mutation.ClearStandings()
	return su
}

// RemoveStandingIDs removes the "standings" edge to Standings entities by IDs.
func (su *SeasonUpdate) RemoveStandingIDs(ids ...int) *SeasonUpdate {
	su.mutation.RemoveStandingIDs(ids...)
	return su
}

// RemoveStandings removes "standings" edges to Standings entities.
func (su *SeasonUpdate) RemoveStandings(s ...*Standings) *SeasonUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveStandingIDs(ids...)
}

// ClearTeams clears all "teams" edges to the Team entity.
func (su *SeasonUpdate) ClearTeams() *SeasonUpdate {
	su.mutation.ClearTeams()
	return su
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (su *SeasonUpdate) RemoveTeamIDs(ids ...int) *SeasonUpdate {
	su.mutation.RemoveTeamIDs(ids...)
	return su
}

// RemoveTeams removes "teams" edges to Team entities.
func (su *SeasonUpdate) RemoveTeams(t ...*Team) *SeasonUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTeamIDs(ids...)
}

// ClearSquad clears all "squad" edges to the Squad entity.
func (su *SeasonUpdate) ClearSquad() *SeasonUpdate {
	su.mutation.ClearSquad()
	return su
}

// RemoveSquadIDs removes the "squad" edge to Squad entities by IDs.
func (su *SeasonUpdate) RemoveSquadIDs(ids ...int) *SeasonUpdate {
	su.mutation.RemoveSquadIDs(ids...)
	return su
}

// RemoveSquad removes "squad" edges to Squad entities.
func (su *SeasonUpdate) RemoveSquad(s ...*Squad) *SeasonUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSquadIDs(ids...)
}

// ClearPlayerStats clears all "playerStats" edges to the PlayerStats entity.
func (su *SeasonUpdate) ClearPlayerStats() *SeasonUpdate {
	su.mutation.ClearPlayerStats()
	return su
}

// RemovePlayerStatIDs removes the "playerStats" edge to PlayerStats entities by IDs.
func (su *SeasonUpdate) RemovePlayerStatIDs(ids ...int) *SeasonUpdate {
	su.mutation.RemovePlayerStatIDs(ids...)
	return su
}

// RemovePlayerStats removes "playerStats" edges to PlayerStats entities.
func (su *SeasonUpdate) RemovePlayerStats(p ...*PlayerStats) *SeasonUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemovePlayerStatIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SeasonUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeasonUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeasonUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeasonUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SeasonUpdate) defaults() {
	if _, ok := su.mutation.LastUpdated(); !ok && !su.mutation.LastUpdatedCleared() {
		v := season.UpdateDefaultLastUpdated()
		su.mutation.SetLastUpdated(v)
	}
}

func (su *SeasonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(season.Table, season.Columns, sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.StartDate(); ok {
		_spec.SetField(season.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := su.mutation.EndDate(); ok {
		_spec.SetField(season.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := su.mutation.Current(); ok {
		_spec.SetField(season.FieldCurrent, field.TypeBool, value)
	}
	if value, ok := su.mutation.LastUpdated(); ok {
		_spec.SetField(season.FieldLastUpdated, field.TypeTime, value)
	}
	if su.mutation.LastUpdatedCleared() {
		_spec.ClearField(season.FieldLastUpdated, field.TypeTime)
	}
	if su.mutation.LeagueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   season.LeagueTable,
			Columns: []string{season.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.LeagueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   season.LeagueTable,
			Columns: []string{season.LeagueColumn},
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
	if su.mutation.FixturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.FixturesTable,
			Columns: []string{season.FixturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedFixturesIDs(); len(nodes) > 0 && !su.mutation.FixturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.FixturesTable,
			Columns: []string{season.FixturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.FixturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.FixturesTable,
			Columns: []string{season.FixturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.StandingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.StandingsTable,
			Columns: []string{season.StandingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedStandingsIDs(); len(nodes) > 0 && !su.mutation.StandingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.StandingsTable,
			Columns: []string{season.StandingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StandingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.StandingsTable,
			Columns: []string{season.StandingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.TeamsTable,
			Columns: []string{season.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !su.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.TeamsTable,
			Columns: []string{season.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.TeamsTable,
			Columns: []string{season.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.SquadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.SquadTable,
			Columns: []string{season.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSquadIDs(); len(nodes) > 0 && !su.mutation.SquadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.SquadTable,
			Columns: []string{season.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SquadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.SquadTable,
			Columns: []string{season.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   season.PlayerStatsTable,
			Columns: season.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedPlayerStatsIDs(); len(nodes) > 0 && !su.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   season.PlayerStatsTable,
			Columns: season.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   season.PlayerStatsTable,
			Columns: season.PlayerStatsPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{season.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SeasonUpdateOne is the builder for updating a single Season entity.
type SeasonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeasonMutation
}

// SetStartDate sets the "start_date" field.
func (suo *SeasonUpdateOne) SetStartDate(t time.Time) *SeasonUpdateOne {
	suo.mutation.SetStartDate(t)
	return suo
}

// SetEndDate sets the "end_date" field.
func (suo *SeasonUpdateOne) SetEndDate(t time.Time) *SeasonUpdateOne {
	suo.mutation.SetEndDate(t)
	return suo
}

// SetCurrent sets the "current" field.
func (suo *SeasonUpdateOne) SetCurrent(b bool) *SeasonUpdateOne {
	suo.mutation.SetCurrent(b)
	return suo
}

// SetLastUpdated sets the "lastUpdated" field.
func (suo *SeasonUpdateOne) SetLastUpdated(t time.Time) *SeasonUpdateOne {
	suo.mutation.SetLastUpdated(t)
	return suo
}

// ClearLastUpdated clears the value of the "lastUpdated" field.
func (suo *SeasonUpdateOne) ClearLastUpdated() *SeasonUpdateOne {
	suo.mutation.ClearLastUpdated()
	return suo
}

// SetLeagueID sets the "league" edge to the League entity by ID.
func (suo *SeasonUpdateOne) SetLeagueID(id int) *SeasonUpdateOne {
	suo.mutation.SetLeagueID(id)
	return suo
}

// SetNillableLeagueID sets the "league" edge to the League entity by ID if the given value is not nil.
func (suo *SeasonUpdateOne) SetNillableLeagueID(id *int) *SeasonUpdateOne {
	if id != nil {
		suo = suo.SetLeagueID(*id)
	}
	return suo
}

// SetLeague sets the "league" edge to the League entity.
func (suo *SeasonUpdateOne) SetLeague(l *League) *SeasonUpdateOne {
	return suo.SetLeagueID(l.ID)
}

// AddFixtureIDs adds the "fixtures" edge to the Fixture entity by IDs.
func (suo *SeasonUpdateOne) AddFixtureIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.AddFixtureIDs(ids...)
	return suo
}

// AddFixtures adds the "fixtures" edges to the Fixture entity.
func (suo *SeasonUpdateOne) AddFixtures(f ...*Fixture) *SeasonUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.AddFixtureIDs(ids...)
}

// AddStandingIDs adds the "standings" edge to the Standings entity by IDs.
func (suo *SeasonUpdateOne) AddStandingIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.AddStandingIDs(ids...)
	return suo
}

// AddStandings adds the "standings" edges to the Standings entity.
func (suo *SeasonUpdateOne) AddStandings(s ...*Standings) *SeasonUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddStandingIDs(ids...)
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (suo *SeasonUpdateOne) AddTeamIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.AddTeamIDs(ids...)
	return suo
}

// AddTeams adds the "teams" edges to the Team entity.
func (suo *SeasonUpdateOne) AddTeams(t ...*Team) *SeasonUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTeamIDs(ids...)
}

// AddSquadIDs adds the "squad" edge to the Squad entity by IDs.
func (suo *SeasonUpdateOne) AddSquadIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.AddSquadIDs(ids...)
	return suo
}

// AddSquad adds the "squad" edges to the Squad entity.
func (suo *SeasonUpdateOne) AddSquad(s ...*Squad) *SeasonUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSquadIDs(ids...)
}

// AddPlayerStatIDs adds the "playerStats" edge to the PlayerStats entity by IDs.
func (suo *SeasonUpdateOne) AddPlayerStatIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.AddPlayerStatIDs(ids...)
	return suo
}

// AddPlayerStats adds the "playerStats" edges to the PlayerStats entity.
func (suo *SeasonUpdateOne) AddPlayerStats(p ...*PlayerStats) *SeasonUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddPlayerStatIDs(ids...)
}

// Mutation returns the SeasonMutation object of the builder.
func (suo *SeasonUpdateOne) Mutation() *SeasonMutation {
	return suo.mutation
}

// ClearLeague clears the "league" edge to the League entity.
func (suo *SeasonUpdateOne) ClearLeague() *SeasonUpdateOne {
	suo.mutation.ClearLeague()
	return suo
}

// ClearFixtures clears all "fixtures" edges to the Fixture entity.
func (suo *SeasonUpdateOne) ClearFixtures() *SeasonUpdateOne {
	suo.mutation.ClearFixtures()
	return suo
}

// RemoveFixtureIDs removes the "fixtures" edge to Fixture entities by IDs.
func (suo *SeasonUpdateOne) RemoveFixtureIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.RemoveFixtureIDs(ids...)
	return suo
}

// RemoveFixtures removes "fixtures" edges to Fixture entities.
func (suo *SeasonUpdateOne) RemoveFixtures(f ...*Fixture) *SeasonUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.RemoveFixtureIDs(ids...)
}

// ClearStandings clears all "standings" edges to the Standings entity.
func (suo *SeasonUpdateOne) ClearStandings() *SeasonUpdateOne {
	suo.mutation.ClearStandings()
	return suo
}

// RemoveStandingIDs removes the "standings" edge to Standings entities by IDs.
func (suo *SeasonUpdateOne) RemoveStandingIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.RemoveStandingIDs(ids...)
	return suo
}

// RemoveStandings removes "standings" edges to Standings entities.
func (suo *SeasonUpdateOne) RemoveStandings(s ...*Standings) *SeasonUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveStandingIDs(ids...)
}

// ClearTeams clears all "teams" edges to the Team entity.
func (suo *SeasonUpdateOne) ClearTeams() *SeasonUpdateOne {
	suo.mutation.ClearTeams()
	return suo
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (suo *SeasonUpdateOne) RemoveTeamIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.RemoveTeamIDs(ids...)
	return suo
}

// RemoveTeams removes "teams" edges to Team entities.
func (suo *SeasonUpdateOne) RemoveTeams(t ...*Team) *SeasonUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTeamIDs(ids...)
}

// ClearSquad clears all "squad" edges to the Squad entity.
func (suo *SeasonUpdateOne) ClearSquad() *SeasonUpdateOne {
	suo.mutation.ClearSquad()
	return suo
}

// RemoveSquadIDs removes the "squad" edge to Squad entities by IDs.
func (suo *SeasonUpdateOne) RemoveSquadIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.RemoveSquadIDs(ids...)
	return suo
}

// RemoveSquad removes "squad" edges to Squad entities.
func (suo *SeasonUpdateOne) RemoveSquad(s ...*Squad) *SeasonUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSquadIDs(ids...)
}

// ClearPlayerStats clears all "playerStats" edges to the PlayerStats entity.
func (suo *SeasonUpdateOne) ClearPlayerStats() *SeasonUpdateOne {
	suo.mutation.ClearPlayerStats()
	return suo
}

// RemovePlayerStatIDs removes the "playerStats" edge to PlayerStats entities by IDs.
func (suo *SeasonUpdateOne) RemovePlayerStatIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.RemovePlayerStatIDs(ids...)
	return suo
}

// RemovePlayerStats removes "playerStats" edges to PlayerStats entities.
func (suo *SeasonUpdateOne) RemovePlayerStats(p ...*PlayerStats) *SeasonUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemovePlayerStatIDs(ids...)
}

// Where appends a list predicates to the SeasonUpdate builder.
func (suo *SeasonUpdateOne) Where(ps ...predicate.Season) *SeasonUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SeasonUpdateOne) Select(field string, fields ...string) *SeasonUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Season entity.
func (suo *SeasonUpdateOne) Save(ctx context.Context) (*Season, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeasonUpdateOne) SaveX(ctx context.Context) *Season {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SeasonUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeasonUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SeasonUpdateOne) defaults() {
	if _, ok := suo.mutation.LastUpdated(); !ok && !suo.mutation.LastUpdatedCleared() {
		v := season.UpdateDefaultLastUpdated()
		suo.mutation.SetLastUpdated(v)
	}
}

func (suo *SeasonUpdateOne) sqlSave(ctx context.Context) (_node *Season, err error) {
	_spec := sqlgraph.NewUpdateSpec(season.Table, season.Columns, sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Season.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, season.FieldID)
		for _, f := range fields {
			if !season.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != season.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.StartDate(); ok {
		_spec.SetField(season.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := suo.mutation.EndDate(); ok {
		_spec.SetField(season.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Current(); ok {
		_spec.SetField(season.FieldCurrent, field.TypeBool, value)
	}
	if value, ok := suo.mutation.LastUpdated(); ok {
		_spec.SetField(season.FieldLastUpdated, field.TypeTime, value)
	}
	if suo.mutation.LastUpdatedCleared() {
		_spec.ClearField(season.FieldLastUpdated, field.TypeTime)
	}
	if suo.mutation.LeagueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   season.LeagueTable,
			Columns: []string{season.LeagueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.LeagueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   season.LeagueTable,
			Columns: []string{season.LeagueColumn},
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
	if suo.mutation.FixturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.FixturesTable,
			Columns: []string{season.FixturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedFixturesIDs(); len(nodes) > 0 && !suo.mutation.FixturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.FixturesTable,
			Columns: []string{season.FixturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.FixturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.FixturesTable,
			Columns: []string{season.FixturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.StandingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.StandingsTable,
			Columns: []string{season.StandingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedStandingsIDs(); len(nodes) > 0 && !suo.mutation.StandingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.StandingsTable,
			Columns: []string{season.StandingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StandingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.StandingsTable,
			Columns: []string{season.StandingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(standings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.TeamsTable,
			Columns: []string{season.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !suo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.TeamsTable,
			Columns: []string{season.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.TeamsTable,
			Columns: []string{season.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.SquadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.SquadTable,
			Columns: []string{season.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSquadIDs(); len(nodes) > 0 && !suo.mutation.SquadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.SquadTable,
			Columns: []string{season.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SquadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.SquadTable,
			Columns: []string{season.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   season.PlayerStatsTable,
			Columns: season.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedPlayerStatsIDs(); len(nodes) > 0 && !suo.mutation.PlayerStatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   season.PlayerStatsTable,
			Columns: season.PlayerStatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PlayerStatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   season.PlayerStatsTable,
			Columns: season.PlayerStatsPrimaryKey,
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
	_node = &Season{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{season.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
