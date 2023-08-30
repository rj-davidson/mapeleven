// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/league"
	"mapeleven/db/ent/predicate"
	"mapeleven/db/ent/season"
	"mapeleven/db/ent/standings"
	"mapeleven/db/ent/team"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LeagueQuery is the builder for querying League entities.
type LeagueQuery struct {
	config
	ctx           *QueryContext
	order         []league.Order
	inters        []Interceptor
	predicates    []predicate.League
	withSeason    *SeasonQuery
	withStandings *StandingsQuery
	withTeams     *TeamQuery
	withCountry   *CountryQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LeagueQuery builder.
func (lq *LeagueQuery) Where(ps ...predicate.League) *LeagueQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LeagueQuery) Limit(limit int) *LeagueQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LeagueQuery) Offset(offset int) *LeagueQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LeagueQuery) Unique(unique bool) *LeagueQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LeagueQuery) Order(o ...league.Order) *LeagueQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QuerySeason chains the current query on the "season" edge.
func (lq *LeagueQuery) QuerySeason() *SeasonQuery {
	query := (&SeasonClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(league.Table, league.FieldID, selector),
			sqlgraph.To(season.Table, season.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, league.SeasonTable, league.SeasonColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStandings chains the current query on the "standings" edge.
func (lq *LeagueQuery) QueryStandings() *StandingsQuery {
	query := (&StandingsClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(league.Table, league.FieldID, selector),
			sqlgraph.To(standings.Table, standings.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, league.StandingsTable, league.StandingsColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeams chains the current query on the "teams" edge.
func (lq *LeagueQuery) QueryTeams() *TeamQuery {
	query := (&TeamClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(league.Table, league.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, league.TeamsTable, league.TeamsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCountry chains the current query on the "country" edge.
func (lq *LeagueQuery) QueryCountry() *CountryQuery {
	query := (&CountryClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(league.Table, league.FieldID, selector),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, league.CountryTable, league.CountryColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first League entity from the query.
// Returns a *NotFoundError when no League was found.
func (lq *LeagueQuery) First(ctx context.Context) (*League, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{league.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LeagueQuery) FirstX(ctx context.Context) *League {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first League ID from the query.
// Returns a *NotFoundError when no League ID was found.
func (lq *LeagueQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{league.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LeagueQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single League entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one League entity is found.
// Returns a *NotFoundError when no League entities are found.
func (lq *LeagueQuery) Only(ctx context.Context) (*League, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{league.Label}
	default:
		return nil, &NotSingularError{league.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LeagueQuery) OnlyX(ctx context.Context) *League {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only League ID in the query.
// Returns a *NotSingularError when more than one League ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LeagueQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{league.Label}
	default:
		err = &NotSingularError{league.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LeagueQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Leagues.
func (lq *LeagueQuery) All(ctx context.Context) ([]*League, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*League, *LeagueQuery]()
	return withInterceptors[[]*League](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LeagueQuery) AllX(ctx context.Context) []*League {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of League IDs.
func (lq *LeagueQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(league.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LeagueQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LeagueQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LeagueQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LeagueQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LeagueQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LeagueQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LeagueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LeagueQuery) Clone() *LeagueQuery {
	if lq == nil {
		return nil
	}
	return &LeagueQuery{
		config:        lq.config,
		ctx:           lq.ctx.Clone(),
		order:         append([]league.Order{}, lq.order...),
		inters:        append([]Interceptor{}, lq.inters...),
		predicates:    append([]predicate.League{}, lq.predicates...),
		withSeason:    lq.withSeason.Clone(),
		withStandings: lq.withStandings.Clone(),
		withTeams:     lq.withTeams.Clone(),
		withCountry:   lq.withCountry.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithSeason tells the query-builder to eager-load the nodes that are connected to
// the "season" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LeagueQuery) WithSeason(opts ...func(*SeasonQuery)) *LeagueQuery {
	query := (&SeasonClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withSeason = query
	return lq
}

// WithStandings tells the query-builder to eager-load the nodes that are connected to
// the "standings" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LeagueQuery) WithStandings(opts ...func(*StandingsQuery)) *LeagueQuery {
	query := (&StandingsClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withStandings = query
	return lq
}

// WithTeams tells the query-builder to eager-load the nodes that are connected to
// the "teams" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LeagueQuery) WithTeams(opts ...func(*TeamQuery)) *LeagueQuery {
	query := (&TeamClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withTeams = query
	return lq
}

// WithCountry tells the query-builder to eager-load the nodes that are connected to
// the "country" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LeagueQuery) WithCountry(opts ...func(*CountryQuery)) *LeagueQuery {
	query := (&CountryClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withCountry = query
	return lq
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
//	client.League.Query().
//		GroupBy(league.FieldSlug).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LeagueQuery) GroupBy(field string, fields ...string) *LeagueGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LeagueGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = league.Label
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
//	client.League.Query().
//		Select(league.FieldSlug).
//		Scan(ctx, &v)
func (lq *LeagueQuery) Select(fields ...string) *LeagueSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LeagueSelect{LeagueQuery: lq}
	sbuild.label = league.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LeagueSelect configured with the given aggregations.
func (lq *LeagueQuery) Aggregate(fns ...AggregateFunc) *LeagueSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LeagueQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !league.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LeagueQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*League, error) {
	var (
		nodes       = []*League{}
		withFKs     = lq.withFKs
		_spec       = lq.querySpec()
		loadedTypes = [4]bool{
			lq.withSeason != nil,
			lq.withStandings != nil,
			lq.withTeams != nil,
			lq.withCountry != nil,
		}
	)
	if lq.withSeason != nil || lq.withCountry != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, league.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*League).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &League{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withSeason; query != nil {
		if err := lq.loadSeason(ctx, query, nodes, nil,
			func(n *League, e *Season) { n.Edges.Season = e }); err != nil {
			return nil, err
		}
	}
	if query := lq.withStandings; query != nil {
		if err := lq.loadStandings(ctx, query, nodes,
			func(n *League) { n.Edges.Standings = []*Standings{} },
			func(n *League, e *Standings) { n.Edges.Standings = append(n.Edges.Standings, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withTeams; query != nil {
		if err := lq.loadTeams(ctx, query, nodes,
			func(n *League) { n.Edges.Teams = []*Team{} },
			func(n *League, e *Team) { n.Edges.Teams = append(n.Edges.Teams, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withCountry; query != nil {
		if err := lq.loadCountry(ctx, query, nodes, nil,
			func(n *League, e *Country) { n.Edges.Country = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LeagueQuery) loadSeason(ctx context.Context, query *SeasonQuery, nodes []*League, init func(*League), assign func(*League, *Season)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*League)
	for i := range nodes {
		if nodes[i].season_league == nil {
			continue
		}
		fk := *nodes[i].season_league
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(season.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "season_league" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lq *LeagueQuery) loadStandings(ctx context.Context, query *StandingsQuery, nodes []*League, init func(*League), assign func(*League, *Standings)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*League)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Standings(func(s *sql.Selector) {
		s.Where(sql.InValues(league.StandingsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.league_standings
		if fk == nil {
			return fmt.Errorf(`foreign-key "league_standings" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "league_standings" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (lq *LeagueQuery) loadTeams(ctx context.Context, query *TeamQuery, nodes []*League, init func(*League), assign func(*League, *Team)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*League)
	nids := make(map[int]map[*League]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(league.TeamsTable)
		s.Join(joinT).On(s.C(team.FieldID), joinT.C(league.TeamsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(league.TeamsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(league.TeamsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*League]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Team](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "teams" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (lq *LeagueQuery) loadCountry(ctx context.Context, query *CountryQuery, nodes []*League, init func(*League), assign func(*League, *Country)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*League)
	for i := range nodes {
		if nodes[i].country_leagues == nil {
			continue
		}
		fk := *nodes[i].country_leagues
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
			return fmt.Errorf(`unexpected foreign-key "country_leagues" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lq *LeagueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LeagueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(league.Table, league.Columns, sqlgraph.NewFieldSpec(league.FieldID, field.TypeInt))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, league.FieldID)
		for i := range fields {
			if fields[i] != league.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LeagueQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(league.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = league.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LeagueGroupBy is the group-by builder for League entities.
type LeagueGroupBy struct {
	selector
	build *LeagueQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LeagueGroupBy) Aggregate(fns ...AggregateFunc) *LeagueGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LeagueGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LeagueQuery, *LeagueGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LeagueGroupBy) sqlScan(ctx context.Context, root *LeagueQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LeagueSelect is the builder for selecting fields of League entities.
type LeagueSelect struct {
	*LeagueQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LeagueSelect) Aggregate(fns ...AggregateFunc) *LeagueSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LeagueSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LeagueQuery, *LeagueSelect](ctx, ls.LeagueQuery, ls, ls.inters, v)
}

func (ls *LeagueSelect) sqlScan(ctx context.Context, root *LeagueQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
