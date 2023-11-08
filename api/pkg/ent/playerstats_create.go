// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/fixtureevents"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/matchplayer"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/player"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgoals"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psoffense"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pspenalty"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/team"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlayerStatsCreate is the builder for creating a PlayerStats entity.
type PlayerStatsCreate struct {
	config
	mutation *PlayerStatsMutation
	hooks    []Hook
}

// SetSlug sets the "slug" field.
func (psc *PlayerStatsCreate) SetSlug(s string) *PlayerStatsCreate {
	psc.mutation.SetSlug(s)
	return psc
}

// SetLastUpdated sets the "lastUpdated" field.
func (psc *PlayerStatsCreate) SetLastUpdated(t time.Time) *PlayerStatsCreate {
	psc.mutation.SetLastUpdated(t)
	return psc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (psc *PlayerStatsCreate) SetNillableLastUpdated(t *time.Time) *PlayerStatsCreate {
	if t != nil {
		psc.SetLastUpdated(*t)
	}
	return psc
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (psc *PlayerStatsCreate) SetPlayerID(id int) *PlayerStatsCreate {
	psc.mutation.SetPlayerID(id)
	return psc
}

// SetNillablePlayerID sets the "player" edge to the Player entity by ID if the given value is not nil.
func (psc *PlayerStatsCreate) SetNillablePlayerID(id *int) *PlayerStatsCreate {
	if id != nil {
		psc = psc.SetPlayerID(*id)
	}
	return psc
}

// SetPlayer sets the "player" edge to the Player entity.
func (psc *PlayerStatsCreate) SetPlayer(p *Player) *PlayerStatsCreate {
	return psc.SetPlayerID(p.ID)
}

// AddTeamIDs adds the "team" edge to the Team entity by IDs.
func (psc *PlayerStatsCreate) AddTeamIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddTeamIDs(ids...)
	return psc
}

// AddTeam adds the "team" edges to the Team entity.
func (psc *PlayerStatsCreate) AddTeam(t ...*Team) *PlayerStatsCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return psc.AddTeamIDs(ids...)
}

// AddPlayerEventIDs adds the "playerEvents" edge to the FixtureEvents entity by IDs.
func (psc *PlayerStatsCreate) AddPlayerEventIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddPlayerEventIDs(ids...)
	return psc
}

// AddPlayerEvents adds the "playerEvents" edges to the FixtureEvents entity.
func (psc *PlayerStatsCreate) AddPlayerEvents(f ...*FixtureEvents) *PlayerStatsCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return psc.AddPlayerEventIDs(ids...)
}

// AddMatchPlayerIDs adds the "matchPlayer" edge to the MatchPlayer entity by IDs.
func (psc *PlayerStatsCreate) AddMatchPlayerIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddMatchPlayerIDs(ids...)
	return psc
}

// AddMatchPlayer adds the "matchPlayer" edges to the MatchPlayer entity.
func (psc *PlayerStatsCreate) AddMatchPlayer(m ...*MatchPlayer) *PlayerStatsCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return psc.AddMatchPlayerIDs(ids...)
}

// AddAssistEventIDs adds the "assistEvents" edge to the FixtureEvents entity by IDs.
func (psc *PlayerStatsCreate) AddAssistEventIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddAssistEventIDs(ids...)
	return psc
}

// AddAssistEvents adds the "assistEvents" edges to the FixtureEvents entity.
func (psc *PlayerStatsCreate) AddAssistEvents(f ...*FixtureEvents) *PlayerStatsCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return psc.AddAssistEventIDs(ids...)
}

// AddPsgameIDs adds the "psgames" edge to the PSGames entity by IDs.
func (psc *PlayerStatsCreate) AddPsgameIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddPsgameIDs(ids...)
	return psc
}

// AddPsgames adds the "psgames" edges to the PSGames entity.
func (psc *PlayerStatsCreate) AddPsgames(p ...*PSGames) *PlayerStatsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psc.AddPsgameIDs(ids...)
}

// AddPsgoalIDs adds the "psgoals" edge to the PSGoals entity by IDs.
func (psc *PlayerStatsCreate) AddPsgoalIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddPsgoalIDs(ids...)
	return psc
}

// AddPsgoals adds the "psgoals" edges to the PSGoals entity.
func (psc *PlayerStatsCreate) AddPsgoals(p ...*PSGoals) *PlayerStatsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psc.AddPsgoalIDs(ids...)
}

// AddPsdefenseIDs adds the "psdefense" edge to the PSDefense entity by IDs.
func (psc *PlayerStatsCreate) AddPsdefenseIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddPsdefenseIDs(ids...)
	return psc
}

// AddPsdefense adds the "psdefense" edges to the PSDefense entity.
func (psc *PlayerStatsCreate) AddPsdefense(p ...*PSDefense) *PlayerStatsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psc.AddPsdefenseIDs(ids...)
}

// AddPsoffenseIDs adds the "psoffense" edge to the PSOffense entity by IDs.
func (psc *PlayerStatsCreate) AddPsoffenseIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddPsoffenseIDs(ids...)
	return psc
}

// AddPsoffense adds the "psoffense" edges to the PSOffense entity.
func (psc *PlayerStatsCreate) AddPsoffense(p ...*PSOffense) *PlayerStatsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psc.AddPsoffenseIDs(ids...)
}

// AddPspenaltyIDs adds the "pspenalty" edge to the PSPenalty entity by IDs.
func (psc *PlayerStatsCreate) AddPspenaltyIDs(ids ...int) *PlayerStatsCreate {
	psc.mutation.AddPspenaltyIDs(ids...)
	return psc
}

// AddPspenalty adds the "pspenalty" edges to the PSPenalty entity.
func (psc *PlayerStatsCreate) AddPspenalty(p ...*PSPenalty) *PlayerStatsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psc.AddPspenaltyIDs(ids...)
}

// Mutation returns the PlayerStatsMutation object of the builder.
func (psc *PlayerStatsCreate) Mutation() *PlayerStatsMutation {
	return psc.mutation
}

// Save creates the PlayerStats in the database.
func (psc *PlayerStatsCreate) Save(ctx context.Context) (*PlayerStats, error) {
	psc.defaults()
	return withHooks(ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PlayerStatsCreate) SaveX(ctx context.Context) *PlayerStats {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *PlayerStatsCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *PlayerStatsCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psc *PlayerStatsCreate) defaults() {
	if _, ok := psc.mutation.LastUpdated(); !ok {
		v := playerstats.DefaultLastUpdated()
		psc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PlayerStatsCreate) check() error {
	if _, ok := psc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "PlayerStats.slug"`)}
	}
	return nil
}

func (psc *PlayerStatsCreate) sqlSave(ctx context.Context) (*PlayerStats, error) {
	if err := psc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	psc.mutation.id = &_node.ID
	psc.mutation.done = true
	return _node, nil
}

func (psc *PlayerStatsCreate) createSpec() (*PlayerStats, *sqlgraph.CreateSpec) {
	var (
		_node = &PlayerStats{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(playerstats.Table, sqlgraph.NewFieldSpec(playerstats.FieldID, field.TypeInt))
	)
	if value, ok := psc.mutation.Slug(); ok {
		_spec.SetField(playerstats.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := psc.mutation.LastUpdated(); ok {
		_spec.SetField(playerstats.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := psc.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playerstats.PlayerTable,
			Columns: []string{playerstats.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_player_stats = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playerstats.TeamTable,
			Columns: []string{playerstats.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.PlayerEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playerstats.PlayerEventsTable,
			Columns: []string{playerstats.PlayerEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixtureevents.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.MatchPlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playerstats.MatchPlayerTable,
			Columns: []string{playerstats.MatchPlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(matchplayer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.AssistEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playerstats.AssistEventsTable,
			Columns: []string{playerstats.AssistEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fixtureevents.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.PsgamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playerstats.PsgamesTable,
			Columns: playerstats.PsgamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(psgames.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.PsgoalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playerstats.PsgoalsTable,
			Columns: playerstats.PsgoalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(psgoals.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.PsdefenseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playerstats.PsdefenseTable,
			Columns: playerstats.PsdefensePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(psdefense.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.PsoffenseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playerstats.PsoffenseTable,
			Columns: playerstats.PsoffensePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(psoffense.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.PspenaltyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playerstats.PspenaltyTable,
			Columns: playerstats.PspenaltyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pspenalty.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlayerStatsCreateBulk is the builder for creating many PlayerStats entities in bulk.
type PlayerStatsCreateBulk struct {
	config
	err      error
	builders []*PlayerStatsCreate
}

// Save creates the PlayerStats entities in the database.
func (pscb *PlayerStatsCreateBulk) Save(ctx context.Context) ([]*PlayerStats, error) {
	if pscb.err != nil {
		return nil, pscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PlayerStats, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerStatsMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *PlayerStatsCreateBulk) SaveX(ctx context.Context) []*PlayerStats {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *PlayerStatsCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *PlayerStatsCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}
