// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/pstechnical"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSTechnicalQuery is the builder for querying PSTechnical entities.
type PSTechnicalQuery struct {
	config
	ctx             *QueryContext
	order           []pstechnical.OrderOption
	inters          []Interceptor
	predicates      []predicate.PSTechnical
	withPlayerStats *PlayerStatsQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PSTechnicalQuery builder.
func (ptq *PSTechnicalQuery) Where(ps ...predicate.PSTechnical) *PSTechnicalQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit the number of records to be returned by this query.
func (ptq *PSTechnicalQuery) Limit(limit int) *PSTechnicalQuery {
	ptq.ctx.Limit = &limit
	return ptq
}

// Offset to start from.
func (ptq *PSTechnicalQuery) Offset(offset int) *PSTechnicalQuery {
	ptq.ctx.Offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *PSTechnicalQuery) Unique(unique bool) *PSTechnicalQuery {
	ptq.ctx.Unique = &unique
	return ptq
}

// Order specifies how the records should be ordered.
func (ptq *PSTechnicalQuery) Order(o ...pstechnical.OrderOption) *PSTechnicalQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryPlayerStats chains the current query on the "playerStats" edge.
func (ptq *PSTechnicalQuery) QueryPlayerStats() *PlayerStatsQuery {
	query := (&PlayerStatsClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pstechnical.Table, pstechnical.FieldID, selector),
			sqlgraph.To(playerstats.Table, playerstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, pstechnical.PlayerStatsTable, pstechnical.PlayerStatsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PSTechnical entity from the query.
// Returns a *NotFoundError when no PSTechnical was found.
func (ptq *PSTechnicalQuery) First(ctx context.Context) (*PSTechnical, error) {
	nodes, err := ptq.Limit(1).All(setContextOp(ctx, ptq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pstechnical.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PSTechnicalQuery) FirstX(ctx context.Context) *PSTechnical {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PSTechnical ID from the query.
// Returns a *NotFoundError when no PSTechnical ID was found.
func (ptq *PSTechnicalQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(1).IDs(setContextOp(ctx, ptq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pstechnical.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *PSTechnicalQuery) FirstIDX(ctx context.Context) int {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PSTechnical entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PSTechnical entity is found.
// Returns a *NotFoundError when no PSTechnical entities are found.
func (ptq *PSTechnicalQuery) Only(ctx context.Context) (*PSTechnical, error) {
	nodes, err := ptq.Limit(2).All(setContextOp(ctx, ptq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pstechnical.Label}
	default:
		return nil, &NotSingularError{pstechnical.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PSTechnicalQuery) OnlyX(ctx context.Context) *PSTechnical {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PSTechnical ID in the query.
// Returns a *NotSingularError when more than one PSTechnical ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptq *PSTechnicalQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(2).IDs(setContextOp(ctx, ptq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pstechnical.Label}
	default:
		err = &NotSingularError{pstechnical.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PSTechnicalQuery) OnlyIDX(ctx context.Context) int {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PSTechnicals.
func (ptq *PSTechnicalQuery) All(ctx context.Context) ([]*PSTechnical, error) {
	ctx = setContextOp(ctx, ptq.ctx, "All")
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PSTechnical, *PSTechnicalQuery]()
	return withInterceptors[[]*PSTechnical](ctx, ptq, qr, ptq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PSTechnicalQuery) AllX(ctx context.Context) []*PSTechnical {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PSTechnical IDs.
func (ptq *PSTechnicalQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ptq.ctx.Unique == nil && ptq.path != nil {
		ptq.Unique(true)
	}
	ctx = setContextOp(ctx, ptq.ctx, "IDs")
	if err = ptq.Select(pstechnical.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PSTechnicalQuery) IDsX(ctx context.Context) []int {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PSTechnicalQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Count")
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptq, querierCount[*PSTechnicalQuery](), ptq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PSTechnicalQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PSTechnicalQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Exist")
	switch _, err := ptq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PSTechnicalQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PSTechnicalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PSTechnicalQuery) Clone() *PSTechnicalQuery {
	if ptq == nil {
		return nil
	}
	return &PSTechnicalQuery{
		config:          ptq.config,
		ctx:             ptq.ctx.Clone(),
		order:           append([]pstechnical.OrderOption{}, ptq.order...),
		inters:          append([]Interceptor{}, ptq.inters...),
		predicates:      append([]predicate.PSTechnical{}, ptq.predicates...),
		withPlayerStats: ptq.withPlayerStats.Clone(),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

// WithPlayerStats tells the query-builder to eager-load the nodes that are connected to
// the "playerStats" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *PSTechnicalQuery) WithPlayerStats(opts ...func(*PlayerStatsQuery)) *PSTechnicalQuery {
	query := (&PlayerStatsClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withPlayerStats = query
	return ptq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FoulsDrawn int `json:"FoulsDrawn,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PSTechnical.Query().
//		GroupBy(pstechnical.FieldFoulsDrawn).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptq *PSTechnicalQuery) GroupBy(field string, fields ...string) *PSTechnicalGroupBy {
	ptq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PSTechnicalGroupBy{build: ptq}
	grbuild.flds = &ptq.ctx.Fields
	grbuild.label = pstechnical.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FoulsDrawn int `json:"FoulsDrawn,omitempty"`
//	}
//
//	client.PSTechnical.Query().
//		Select(pstechnical.FieldFoulsDrawn).
//		Scan(ctx, &v)
func (ptq *PSTechnicalQuery) Select(fields ...string) *PSTechnicalSelect {
	ptq.ctx.Fields = append(ptq.ctx.Fields, fields...)
	sbuild := &PSTechnicalSelect{PSTechnicalQuery: ptq}
	sbuild.label = pstechnical.Label
	sbuild.flds, sbuild.scan = &ptq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PSTechnicalSelect configured with the given aggregations.
func (ptq *PSTechnicalQuery) Aggregate(fns ...AggregateFunc) *PSTechnicalSelect {
	return ptq.Select().Aggregate(fns...)
}

func (ptq *PSTechnicalQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptq.ctx.Fields {
		if !pstechnical.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PSTechnicalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PSTechnical, error) {
	var (
		nodes       = []*PSTechnical{}
		withFKs     = ptq.withFKs
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withPlayerStats != nil,
		}
	)
	if ptq.withPlayerStats != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, pstechnical.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PSTechnical).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PSTechnical{config: ptq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptq.withPlayerStats; query != nil {
		if err := ptq.loadPlayerStats(ctx, query, nodes, nil,
			func(n *PSTechnical, e *PlayerStats) { n.Edges.PlayerStats = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptq *PSTechnicalQuery) loadPlayerStats(ctx context.Context, query *PlayerStatsQuery, nodes []*PSTechnical, init func(*PSTechnical), assign func(*PSTechnical, *PlayerStats)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PSTechnical)
	for i := range nodes {
		if nodes[i].player_stats_ps_technical == nil {
			continue
		}
		fk := *nodes[i].player_stats_ps_technical
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
			return fmt.Errorf(`unexpected foreign-key "player_stats_ps_technical" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ptq *PSTechnicalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	_spec.Node.Columns = ptq.ctx.Fields
	if len(ptq.ctx.Fields) > 0 {
		_spec.Unique = ptq.ctx.Unique != nil && *ptq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PSTechnicalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pstechnical.Table, pstechnical.Columns, sqlgraph.NewFieldSpec(pstechnical.FieldID, field.TypeInt))
	_spec.From = ptq.sql
	if unique := ptq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptq.path != nil {
		_spec.Unique = true
	}
	if fields := ptq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pstechnical.FieldID)
		for i := range fields {
			if fields[i] != pstechnical.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *PSTechnicalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(pstechnical.Table)
	columns := ptq.ctx.Fields
	if len(columns) == 0 {
		columns = pstechnical.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptq.ctx.Unique != nil && *ptq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PSTechnicalGroupBy is the group-by builder for PSTechnical entities.
type PSTechnicalGroupBy struct {
	selector
	build *PSTechnicalQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PSTechnicalGroupBy) Aggregate(fns ...AggregateFunc) *PSTechnicalGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptgb *PSTechnicalGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptgb.build.ctx, "GroupBy")
	if err := ptgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PSTechnicalQuery, *PSTechnicalGroupBy](ctx, ptgb.build, ptgb, ptgb.build.inters, v)
}

func (ptgb *PSTechnicalGroupBy) sqlScan(ctx context.Context, root *PSTechnicalQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptgb.flds)+len(ptgb.fns))
		for _, f := range *ptgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PSTechnicalSelect is the builder for selecting fields of PSTechnical entities.
type PSTechnicalSelect struct {
	*PSTechnicalQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pts *PSTechnicalSelect) Aggregate(fns ...AggregateFunc) *PSTechnicalSelect {
	pts.fns = append(pts.fns, fns...)
	return pts
}

// Scan applies the selector query and scans the result into the given value.
func (pts *PSTechnicalSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pts.ctx, "Select")
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PSTechnicalQuery, *PSTechnicalSelect](ctx, pts.PSTechnicalQuery, pts, pts.inters, v)
}

func (pts *PSTechnicalSelect) sqlScan(ctx context.Context, root *PSTechnicalQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pts.fns))
	for _, fn := range pts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
