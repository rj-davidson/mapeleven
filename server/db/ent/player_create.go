// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mapeleven/db/ent/birth"
	"mapeleven/db/ent/player"
	"mapeleven/db/ent/playerteamseason"
	"mapeleven/db/ent/team"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlayerCreate is the builder for creating a Player entity.
type PlayerCreate struct {
	config
	mutation *PlayerMutation
	hooks    []Hook
}

// SetSlug sets the "slug" field.
func (pc *PlayerCreate) SetSlug(s string) *PlayerCreate {
	pc.mutation.SetSlug(s)
	return pc
}

// SetName sets the "name" field.
func (pc *PlayerCreate) SetName(s string) *PlayerCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetFirstname sets the "firstname" field.
func (pc *PlayerCreate) SetFirstname(s string) *PlayerCreate {
	pc.mutation.SetFirstname(s)
	return pc
}

// SetLastname sets the "lastname" field.
func (pc *PlayerCreate) SetLastname(s string) *PlayerCreate {
	pc.mutation.SetLastname(s)
	return pc
}

// SetAge sets the "age" field.
func (pc *PlayerCreate) SetAge(i int) *PlayerCreate {
	pc.mutation.SetAge(i)
	return pc
}

// SetHeight sets the "height" field.
func (pc *PlayerCreate) SetHeight(f float64) *PlayerCreate {
	pc.mutation.SetHeight(f)
	return pc
}

// SetWeight sets the "weight" field.
func (pc *PlayerCreate) SetWeight(f float64) *PlayerCreate {
	pc.mutation.SetWeight(f)
	return pc
}

// SetInjured sets the "injured" field.
func (pc *PlayerCreate) SetInjured(b bool) *PlayerCreate {
	pc.mutation.SetInjured(b)
	return pc
}

// SetPhoto sets the "photo" field.
func (pc *PlayerCreate) SetPhoto(s string) *PlayerCreate {
	pc.mutation.SetPhoto(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PlayerCreate) SetID(i int) *PlayerCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetBirthID sets the "birth" edge to the Birth entity by ID.
func (pc *PlayerCreate) SetBirthID(id int) *PlayerCreate {
	pc.mutation.SetBirthID(id)
	return pc
}

// SetNillableBirthID sets the "birth" edge to the Birth entity by ID if the given value is not nil.
func (pc *PlayerCreate) SetNillableBirthID(id *int) *PlayerCreate {
	if id != nil {
		pc = pc.SetBirthID(*id)
	}
	return pc
}

// SetBirth sets the "birth" edge to the Birth entity.
func (pc *PlayerCreate) SetBirth(b *Birth) *PlayerCreate {
	return pc.SetBirthID(b.ID)
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (pc *PlayerCreate) AddTeamIDs(ids ...int) *PlayerCreate {
	pc.mutation.AddTeamIDs(ids...)
	return pc
}

// AddTeams adds the "teams" edges to the Team entity.
func (pc *PlayerCreate) AddTeams(t ...*Team) *PlayerCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTeamIDs(ids...)
}

// AddPlayerTeamSeasonIDs adds the "playerTeamSeasons" edge to the PlayerTeamSeason entity by IDs.
func (pc *PlayerCreate) AddPlayerTeamSeasonIDs(ids ...int) *PlayerCreate {
	pc.mutation.AddPlayerTeamSeasonIDs(ids...)
	return pc
}

// AddPlayerTeamSeasons adds the "playerTeamSeasons" edges to the PlayerTeamSeason entity.
func (pc *PlayerCreate) AddPlayerTeamSeasons(p ...*PlayerTeamSeason) *PlayerCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPlayerTeamSeasonIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (pc *PlayerCreate) Mutation() *PlayerMutation {
	return pc.mutation
}

// Save creates the Player in the database.
func (pc *PlayerCreate) Save(ctx context.Context) (*Player, error) {
	return withHooks[*Player, PlayerMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlayerCreate) SaveX(ctx context.Context) *Player {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlayerCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlayerCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlayerCreate) check() error {
	if _, ok := pc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Player.slug"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Player.name"`)}
	}
	if _, ok := pc.mutation.Firstname(); !ok {
		return &ValidationError{Name: "firstname", err: errors.New(`ent: missing required field "Player.firstname"`)}
	}
	if _, ok := pc.mutation.Lastname(); !ok {
		return &ValidationError{Name: "lastname", err: errors.New(`ent: missing required field "Player.lastname"`)}
	}
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Player.age"`)}
	}
	if _, ok := pc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Player.height"`)}
	}
	if _, ok := pc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Player.weight"`)}
	}
	if _, ok := pc.mutation.Injured(); !ok {
		return &ValidationError{Name: "injured", err: errors.New(`ent: missing required field "Player.injured"`)}
	}
	if _, ok := pc.mutation.Photo(); !ok {
		return &ValidationError{Name: "photo", err: errors.New(`ent: missing required field "Player.photo"`)}
	}
	return nil
}

func (pc *PlayerCreate) sqlSave(ctx context.Context) (*Player, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PlayerCreate) createSpec() (*Player, *sqlgraph.CreateSpec) {
	var (
		_node = &Player{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(player.Table, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Slug(); ok {
		_spec.SetField(player.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(player.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Firstname(); ok {
		_spec.SetField(player.FieldFirstname, field.TypeString, value)
		_node.Firstname = value
	}
	if value, ok := pc.mutation.Lastname(); ok {
		_spec.SetField(player.FieldLastname, field.TypeString, value)
		_node.Lastname = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.SetField(player.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if value, ok := pc.mutation.Height(); ok {
		_spec.SetField(player.FieldHeight, field.TypeFloat64, value)
		_node.Height = value
	}
	if value, ok := pc.mutation.Weight(); ok {
		_spec.SetField(player.FieldWeight, field.TypeFloat64, value)
		_node.Weight = value
	}
	if value, ok := pc.mutation.Injured(); ok {
		_spec.SetField(player.FieldInjured, field.TypeBool, value)
		_node.Injured = value
	}
	if value, ok := pc.mutation.Photo(); ok {
		_spec.SetField(player.FieldPhoto, field.TypeString, value)
		_node.Photo = value
	}
	if nodes := pc.mutation.BirthIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   player.BirthTable,
			Columns: []string{player.BirthColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(birth.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.TeamsTable,
			Columns: player.TeamsPrimaryKey,
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
	if nodes := pc.mutation.PlayerTeamSeasonsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.PlayerTeamSeasonsTable,
			Columns: []string{player.PlayerTeamSeasonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playerteamseason.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlayerCreateBulk is the builder for creating many Player entities in bulk.
type PlayerCreateBulk struct {
	config
	builders []*PlayerCreate
}

// Save creates the Player entities in the database.
func (pcb *PlayerCreateBulk) Save(ctx context.Context) ([]*Player, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Player, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlayerCreateBulk) SaveX(ctx context.Context) []*Player {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlayerCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlayerCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
