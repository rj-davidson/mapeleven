// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psgames"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSGamesQuery is the builder for querying PSGames entities.
type PSGamesQuery struct {
	config
	ctx             *QueryContext
	order           []psgames.OrderOption
	inters          []Interceptor
	predicates      []predicate.PSGames
	withPlayerStats *PlayerStatsQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PSGamesQuery builder.
func (pgq *PSGamesQuery) Where(ps ...predicate.PSGames) *PSGamesQuery {
	pgq.predicates = append(pgq.predicates, ps...)
	return pgq
}

// Limit the number of records to be returned by this query.
func (pgq *PSGamesQuery) Limit(limit int) *PSGamesQuery {
	pgq.ctx.Limit = &limit
	return pgq
}

// Offset to start from.
func (pgq *PSGamesQuery) Offset(offset int) *PSGamesQuery {
	pgq.ctx.Offset = &offset
	return pgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pgq *PSGamesQuery) Unique(unique bool) *PSGamesQuery {
	pgq.ctx.Unique = &unique
	return pgq
}

// Order specifies how the records should be ordered.
func (pgq *PSGamesQuery) Order(o ...psgames.OrderOption) *PSGamesQuery {
	pgq.order = append(pgq.order, o...)
	return pgq
}

// QueryPlayerStats chains the current query on the "playerStats" edge.
func (pgq *PSGamesQuery) QueryPlayerStats() *PlayerStatsQuery {
	query := (&PlayerStatsClient{config: pgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(psgames.Table, psgames.FieldID, selector),
			sqlgraph.To(playerstats.Table, playerstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, psgames.PlayerStatsTable, psgames.PlayerStatsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PSGames entity from the query.
// Returns a *NotFoundError when no PSGames was found.
func (pgq *PSGamesQuery) First(ctx context.Context) (*PSGames, error) {
	nodes, err := pgq.Limit(1).All(setContextOp(ctx, pgq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{psgames.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pgq *PSGamesQuery) FirstX(ctx context.Context) *PSGames {
	node, err := pgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PSGames ID from the query.
// Returns a *NotFoundError when no PSGames ID was found.
func (pgq *PSGamesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pgq.Limit(1).IDs(setContextOp(ctx, pgq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{psgames.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pgq *PSGamesQuery) FirstIDX(ctx context.Context) int {
	id, err := pgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PSGames entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PSGames entity is found.
// Returns a *NotFoundError when no PSGames entities are found.
func (pgq *PSGamesQuery) Only(ctx context.Context) (*PSGames, error) {
	nodes, err := pgq.Limit(2).All(setContextOp(ctx, pgq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{psgames.Label}
	default:
		return nil, &NotSingularError{psgames.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pgq *PSGamesQuery) OnlyX(ctx context.Context) *PSGames {
	node, err := pgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PSGames ID in the query.
// Returns a *NotSingularError when more than one PSGames ID is found.
// Returns a *NotFoundError when no entities are found.
func (pgq *PSGamesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pgq.Limit(2).IDs(setContextOp(ctx, pgq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{psgames.Label}
	default:
		err = &NotSingularError{psgames.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pgq *PSGamesQuery) OnlyIDX(ctx context.Context) int {
	id, err := pgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PSGamesSlice.
func (pgq *PSGamesQuery) All(ctx context.Context) ([]*PSGames, error) {
	ctx = setContextOp(ctx, pgq.ctx, "All")
	if err := pgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PSGames, *PSGamesQuery]()
	return withInterceptors[[]*PSGames](ctx, pgq, qr, pgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pgq *PSGamesQuery) AllX(ctx context.Context) []*PSGames {
	nodes, err := pgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PSGames IDs.
func (pgq *PSGamesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pgq.ctx.Unique == nil && pgq.path != nil {
		pgq.Unique(true)
	}
	ctx = setContextOp(ctx, pgq.ctx, "IDs")
	if err = pgq.Select(psgames.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pgq *PSGamesQuery) IDsX(ctx context.Context) []int {
	ids, err := pgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pgq *PSGamesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pgq.ctx, "Count")
	if err := pgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pgq, querierCount[*PSGamesQuery](), pgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pgq *PSGamesQuery) CountX(ctx context.Context) int {
	count, err := pgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pgq *PSGamesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pgq.ctx, "Exist")
	switch _, err := pgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pgq *PSGamesQuery) ExistX(ctx context.Context) bool {
	exist, err := pgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PSGamesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pgq *PSGamesQuery) Clone() *PSGamesQuery {
	if pgq == nil {
		return nil
	}
	return &PSGamesQuery{
		config:          pgq.config,
		ctx:             pgq.ctx.Clone(),
		order:           append([]psgames.OrderOption{}, pgq.order...),
		inters:          append([]Interceptor{}, pgq.inters...),
		predicates:      append([]predicate.PSGames{}, pgq.predicates...),
		withPlayerStats: pgq.withPlayerStats.Clone(),
		// clone intermediate query.
		sql:  pgq.sql.Clone(),
		path: pgq.path,
	}
}

// WithPlayerStats tells the query-builder to eager-load the nodes that are connected to
// the "playerStats" edge. The optional arguments are used to configure the query builder of the edge.
func (pgq *PSGamesQuery) WithPlayerStats(opts ...func(*PlayerStatsQuery)) *PSGamesQuery {
	query := (&PlayerStatsClient{config: pgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pgq.withPlayerStats = query
	return pgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Appearances int `json:"Appearances,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PSGames.Query().
//		GroupBy(psgames.FieldAppearances).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pgq *PSGamesQuery) GroupBy(field string, fields ...string) *PSGamesGroupBy {
	pgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PSGamesGroupBy{build: pgq}
	grbuild.flds = &pgq.ctx.Fields
	grbuild.label = psgames.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Appearances int `json:"Appearances,omitempty"`
//	}
//
//	client.PSGames.Query().
//		Select(psgames.FieldAppearances).
//		Scan(ctx, &v)
func (pgq *PSGamesQuery) Select(fields ...string) *PSGamesSelect {
	pgq.ctx.Fields = append(pgq.ctx.Fields, fields...)
	sbuild := &PSGamesSelect{PSGamesQuery: pgq}
	sbuild.label = psgames.Label
	sbuild.flds, sbuild.scan = &pgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PSGamesSelect configured with the given aggregations.
func (pgq *PSGamesQuery) Aggregate(fns ...AggregateFunc) *PSGamesSelect {
	return pgq.Select().Aggregate(fns...)
}

func (pgq *PSGamesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pgq); err != nil {
				return err
			}
		}
	}
	for _, f := range pgq.ctx.Fields {
		if !psgames.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pgq.path != nil {
		prev, err := pgq.path(ctx)
		if err != nil {
			return err
		}
		pgq.sql = prev
	}
	return nil
}

func (pgq *PSGamesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PSGames, error) {
	var (
		nodes       = []*PSGames{}
		withFKs     = pgq.withFKs
		_spec       = pgq.querySpec()
		loadedTypes = [1]bool{
			pgq.withPlayerStats != nil,
		}
	)
	if pgq.withPlayerStats != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, psgames.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PSGames).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PSGames{config: pgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pgq.withPlayerStats; query != nil {
		if err := pgq.loadPlayerStats(ctx, query, nodes, nil,
			func(n *PSGames, e *PlayerStats) { n.Edges.PlayerStats = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pgq *PSGamesQuery) loadPlayerStats(ctx context.Context, query *PlayerStatsQuery, nodes []*PSGames, init func(*PSGames), assign func(*PSGames, *PlayerStats)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PSGames)
	for i := range nodes {
		if nodes[i].player_stats_ps_games == nil {
			continue
		}
		fk := *nodes[i].player_stats_ps_games
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(playerstats.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "player_stats_ps_games" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pgq *PSGamesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pgq.querySpec()
	_spec.Node.Columns = pgq.ctx.Fields
	if len(pgq.ctx.Fields) > 0 {
		_spec.Unique = pgq.ctx.Unique != nil && *pgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pgq.driver, _spec)
}

func (pgq *PSGamesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(psgames.Table, psgames.Columns, sqlgraph.NewFieldSpec(psgames.FieldID, field.TypeInt))
	_spec.From = pgq.sql
	if unique := pgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pgq.path != nil {
		_spec.Unique = true
	}
	if fields := pgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, psgames.FieldID)
		for i := range fields {
			if fields[i] != psgames.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pgq *PSGamesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pgq.driver.Dialect())
	t1 := builder.Table(psgames.Table)
	columns := pgq.ctx.Fields
	if len(columns) == 0 {
		columns = psgames.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pgq.sql != nil {
		selector = pgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pgq.ctx.Unique != nil && *pgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pgq.predicates {
		p(selector)
	}
	for _, p := range pgq.order {
		p(selector)
	}
	if offset := pgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PSGamesGroupBy is the group-by builder for PSGames entities.
type PSGamesGroupBy struct {
	selector
	build *PSGamesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pggb *PSGamesGroupBy) Aggregate(fns ...AggregateFunc) *PSGamesGroupBy {
	pggb.fns = append(pggb.fns, fns...)
	return pggb
}

// Scan applies the selector query and scans the result into the given value.
func (pggb *PSGamesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pggb.build.ctx, "GroupBy")
	if err := pggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PSGamesQuery, *PSGamesGroupBy](ctx, pggb.build, pggb, pggb.build.inters, v)
}

func (pggb *PSGamesGroupBy) sqlScan(ctx context.Context, root *PSGamesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pggb.fns))
	for _, fn := range pggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pggb.flds)+len(pggb.fns))
		for _, f := range *pggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PSGamesSelect is the builder for selecting fields of PSGames entities.
type PSGamesSelect struct {
	*PSGamesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pgs *PSGamesSelect) Aggregate(fns ...AggregateFunc) *PSGamesSelect {
	pgs.fns = append(pgs.fns, fns...)
	return pgs
}

// Scan applies the selector query and scans the result into the given value.
func (pgs *PSGamesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgs.ctx, "Select")
	if err := pgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PSGamesQuery, *PSGamesSelect](ctx, pgs.PSGamesQuery, pgs, pgs.inters, v)
}

func (pgs *PSGamesSelect) sqlScan(ctx context.Context, root *PSGamesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pgs.fns))
	for _, fn := range pgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
