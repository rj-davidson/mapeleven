// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/fixture"
	"mapeleven/db/ent/predicate"
	"mapeleven/db/ent/standings"
	"mapeleven/db/ent/team"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeamQuery is the builder for querying Team entities.
type TeamQuery struct {
	config
	ctx              *QueryContext
	order            []team.Order
	inters           []Interceptor
	predicates       []predicate.Team
	withStandings    *StandingsQuery
	withCountry      *CountryQuery
	withHomeFixtures *FixtureQuery
	withAwayFixtures *FixtureQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TeamQuery builder.
func (tq *TeamQuery) Where(ps ...predicate.Team) *TeamQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TeamQuery) Limit(limit int) *TeamQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TeamQuery) Offset(offset int) *TeamQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TeamQuery) Unique(unique bool) *TeamQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TeamQuery) Order(o ...team.Order) *TeamQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryStandings chains the current query on the "standings" edge.
func (tq *TeamQuery) QueryStandings() *StandingsQuery {
	query := (&StandingsClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, selector),
			sqlgraph.To(standings.Table, standings.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, team.StandingsTable, team.StandingsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCountry chains the current query on the "country" edge.
func (tq *TeamQuery) QueryCountry() *CountryQuery {
	query := (&CountryClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, selector),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, team.CountryTable, team.CountryColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHomeFixtures chains the current query on the "homeFixtures" edge.
func (tq *TeamQuery) QueryHomeFixtures() *FixtureQuery {
	query := (&FixtureClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, selector),
			sqlgraph.To(fixture.Table, fixture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, team.HomeFixturesTable, team.HomeFixturesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAwayFixtures chains the current query on the "awayFixtures" edge.
func (tq *TeamQuery) QueryAwayFixtures() *FixtureQuery {
	query := (&FixtureClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, selector),
			sqlgraph.To(fixture.Table, fixture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, team.AwayFixturesTable, team.AwayFixturesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Team entity from the query.
// Returns a *NotFoundError when no Team was found.
func (tq *TeamQuery) First(ctx context.Context) (*Team, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{team.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TeamQuery) FirstX(ctx context.Context) *Team {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Team ID from the query.
// Returns a *NotFoundError when no Team ID was found.
func (tq *TeamQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{team.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TeamQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Team entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Team entity is found.
// Returns a *NotFoundError when no Team entities are found.
func (tq *TeamQuery) Only(ctx context.Context) (*Team, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{team.Label}
	default:
		return nil, &NotSingularError{team.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TeamQuery) OnlyX(ctx context.Context) *Team {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Team ID in the query.
// Returns a *NotSingularError when more than one Team ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TeamQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{team.Label}
	default:
		err = &NotSingularError{team.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TeamQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Teams.
func (tq *TeamQuery) All(ctx context.Context) ([]*Team, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Team, *TeamQuery]()
	return withInterceptors[[]*Team](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TeamQuery) AllX(ctx context.Context) []*Team {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Team IDs.
func (tq *TeamQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(team.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TeamQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TeamQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TeamQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TeamQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TeamQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TeamQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TeamQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TeamQuery) Clone() *TeamQuery {
	if tq == nil {
		return nil
	}
	return &TeamQuery{
		config:           tq.config,
		ctx:              tq.ctx.Clone(),
		order:            append([]team.Order{}, tq.order...),
		inters:           append([]Interceptor{}, tq.inters...),
		predicates:       append([]predicate.Team{}, tq.predicates...),
		withStandings:    tq.withStandings.Clone(),
		withCountry:      tq.withCountry.Clone(),
		withHomeFixtures: tq.withHomeFixtures.Clone(),
		withAwayFixtures: tq.withAwayFixtures.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithStandings tells the query-builder to eager-load the nodes that are connected to
// the "standings" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeamQuery) WithStandings(opts ...func(*StandingsQuery)) *TeamQuery {
	query := (&StandingsClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withStandings = query
	return tq
}

// WithCountry tells the query-builder to eager-load the nodes that are connected to
// the "country" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeamQuery) WithCountry(opts ...func(*CountryQuery)) *TeamQuery {
	query := (&CountryClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withCountry = query
	return tq
}

// WithHomeFixtures tells the query-builder to eager-load the nodes that are connected to
// the "homeFixtures" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeamQuery) WithHomeFixtures(opts ...func(*FixtureQuery)) *TeamQuery {
	query := (&FixtureClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withHomeFixtures = query
	return tq
}

// WithAwayFixtures tells the query-builder to eager-load the nodes that are connected to
// the "awayFixtures" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeamQuery) WithAwayFixtures(opts ...func(*FixtureQuery)) *TeamQuery {
	query := (&FixtureClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withAwayFixtures = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Slug string `json:"slug,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Team.Query().
//		GroupBy(team.FieldSlug).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TeamQuery) GroupBy(field string, fields ...string) *TeamGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TeamGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = team.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Slug string `json:"slug,omitempty"`
//	}
//
//	client.Team.Query().
//		Select(team.FieldSlug).
//		Scan(ctx, &v)
func (tq *TeamQuery) Select(fields ...string) *TeamSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TeamSelect{TeamQuery: tq}
	sbuild.label = team.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TeamSelect configured with the given aggregations.
func (tq *TeamQuery) Aggregate(fns ...AggregateFunc) *TeamSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TeamQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !team.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TeamQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Team, error) {
	var (
		nodes       = []*Team{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [4]bool{
			tq.withStandings != nil,
			tq.withCountry != nil,
			tq.withHomeFixtures != nil,
			tq.withAwayFixtures != nil,
		}
	)
	if tq.withCountry != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, team.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Team).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Team{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withStandings; query != nil {
		if err := tq.loadStandings(ctx, query, nodes,
			func(n *Team) { n.Edges.Standings = []*Standings{} },
			func(n *Team, e *Standings) { n.Edges.Standings = append(n.Edges.Standings, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withCountry; query != nil {
		if err := tq.loadCountry(ctx, query, nodes, nil,
			func(n *Team, e *Country) { n.Edges.Country = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withHomeFixtures; query != nil {
		if err := tq.loadHomeFixtures(ctx, query, nodes,
			func(n *Team) { n.Edges.HomeFixtures = []*Fixture{} },
			func(n *Team, e *Fixture) { n.Edges.HomeFixtures = append(n.Edges.HomeFixtures, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withAwayFixtures; query != nil {
		if err := tq.loadAwayFixtures(ctx, query, nodes,
			func(n *Team) { n.Edges.AwayFixtures = []*Fixture{} },
			func(n *Team, e *Fixture) { n.Edges.AwayFixtures = append(n.Edges.AwayFixtures, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TeamQuery) loadStandings(ctx context.Context, query *StandingsQuery, nodes []*Team, init func(*Team), assign func(*Team, *Standings)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Team)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Standings(func(s *sql.Selector) {
		s.Where(sql.InValues(team.StandingsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.team_standings
		if fk == nil {
			return fmt.Errorf(`foreign-key "team_standings" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_standings" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TeamQuery) loadCountry(ctx context.Context, query *CountryQuery, nodes []*Team, init func(*Team), assign func(*Team, *Country)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Team)
	for i := range nodes {
		if nodes[i].country_teams == nil {
			continue
		}
		fk := *nodes[i].country_teams
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(country.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "country_teams" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TeamQuery) loadHomeFixtures(ctx context.Context, query *FixtureQuery, nodes []*Team, init func(*Team), assign func(*Team, *Fixture)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Team)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Fixture(func(s *sql.Selector) {
		s.Where(sql.InValues(team.HomeFixturesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.team_home_fixtures
		if fk == nil {
			return fmt.Errorf(`foreign-key "team_home_fixtures" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_home_fixtures" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TeamQuery) loadAwayFixtures(ctx context.Context, query *FixtureQuery, nodes []*Team, init func(*Team), assign func(*Team, *Fixture)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Team)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Fixture(func(s *sql.Selector) {
		s.Where(sql.InValues(team.AwayFixturesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.team_away_fixtures
		if fk == nil {
			return fmt.Errorf(`foreign-key "team_away_fixtures" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_away_fixtures" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TeamQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TeamQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, team.FieldID)
		for i := range fields {
			if fields[i] != team.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TeamQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(team.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = team.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TeamGroupBy is the group-by builder for Team entities.
type TeamGroupBy struct {
	selector
	build *TeamQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TeamGroupBy) Aggregate(fns ...AggregateFunc) *TeamGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TeamGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeamQuery, *TeamGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TeamGroupBy) sqlScan(ctx context.Context, root *TeamQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TeamSelect is the builder for selecting fields of Team entities.
type TeamSelect struct {
	*TeamQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TeamSelect) Aggregate(fns ...AggregateFunc) *TeamSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TeamSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeamQuery, *TeamSelect](ctx, ts.TeamQuery, ts, ts.inters, v)
}

func (ts *TeamSelect) sqlScan(ctx context.Context, root *TeamQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
