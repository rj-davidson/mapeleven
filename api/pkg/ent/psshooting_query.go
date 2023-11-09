// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psshooting"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSShootingQuery is the builder for querying PSShooting entities.
type PSShootingQuery struct {
	config
	ctx             *QueryContext
	order           []psshooting.OrderOption
	inters          []Interceptor
	predicates      []predicate.PSShooting
	withPlayerStats *PlayerStatsQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PSShootingQuery builder.
func (psq *PSShootingQuery) Where(ps ...predicate.PSShooting) *PSShootingQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit the number of records to be returned by this query.
func (psq *PSShootingQuery) Limit(limit int) *PSShootingQuery {
	psq.ctx.Limit = &limit
	return psq
}

// Offset to start from.
func (psq *PSShootingQuery) Offset(offset int) *PSShootingQuery {
	psq.ctx.Offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *PSShootingQuery) Unique(unique bool) *PSShootingQuery {
	psq.ctx.Unique = &unique
	return psq
}

// Order specifies how the records should be ordered.
func (psq *PSShootingQuery) Order(o ...psshooting.OrderOption) *PSShootingQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// QueryPlayerStats chains the current query on the "playerStats" edge.
func (psq *PSShootingQuery) QueryPlayerStats() *PlayerStatsQuery {
	query := (&PlayerStatsClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(psshooting.Table, psshooting.FieldID, selector),
			sqlgraph.To(playerstats.Table, playerstats.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, psshooting.PlayerStatsTable, psshooting.PlayerStatsColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PSShooting entity from the query.
// Returns a *NotFoundError when no PSShooting was found.
func (psq *PSShootingQuery) First(ctx context.Context) (*PSShooting, error) {
	nodes, err := psq.Limit(1).All(setContextOp(ctx, psq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{psshooting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *PSShootingQuery) FirstX(ctx context.Context) *PSShooting {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PSShooting ID from the query.
// Returns a *NotFoundError when no PSShooting ID was found.
func (psq *PSShootingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(1).IDs(setContextOp(ctx, psq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{psshooting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *PSShootingQuery) FirstIDX(ctx context.Context) int {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PSShooting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PSShooting entity is found.
// Returns a *NotFoundError when no PSShooting entities are found.
func (psq *PSShootingQuery) Only(ctx context.Context) (*PSShooting, error) {
	nodes, err := psq.Limit(2).All(setContextOp(ctx, psq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{psshooting.Label}
	default:
		return nil, &NotSingularError{psshooting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *PSShootingQuery) OnlyX(ctx context.Context) *PSShooting {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PSShooting ID in the query.
// Returns a *NotSingularError when more than one PSShooting ID is found.
// Returns a *NotFoundError when no entities are found.
func (psq *PSShootingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(2).IDs(setContextOp(ctx, psq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{psshooting.Label}
	default:
		err = &NotSingularError{psshooting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *PSShootingQuery) OnlyIDX(ctx context.Context) int {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PSShootings.
func (psq *PSShootingQuery) All(ctx context.Context) ([]*PSShooting, error) {
	ctx = setContextOp(ctx, psq.ctx, "All")
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PSShooting, *PSShootingQuery]()
	return withInterceptors[[]*PSShooting](ctx, psq, qr, psq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psq *PSShootingQuery) AllX(ctx context.Context) []*PSShooting {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PSShooting IDs.
func (psq *PSShootingQuery) IDs(ctx context.Context) (ids []int, err error) {
	if psq.ctx.Unique == nil && psq.path != nil {
		psq.Unique(true)
	}
	ctx = setContextOp(ctx, psq.ctx, "IDs")
	if err = psq.Select(psshooting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *PSShootingQuery) IDsX(ctx context.Context) []int {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *PSShootingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psq.ctx, "Count")
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psq, querierCount[*PSShootingQuery](), psq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psq *PSShootingQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *PSShootingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psq.ctx, "Exist")
	switch _, err := psq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *PSShootingQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PSShootingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *PSShootingQuery) Clone() *PSShootingQuery {
	if psq == nil {
		return nil
	}
	return &PSShootingQuery{
		config:          psq.config,
		ctx:             psq.ctx.Clone(),
		order:           append([]psshooting.OrderOption{}, psq.order...),
		inters:          append([]Interceptor{}, psq.inters...),
		predicates:      append([]predicate.PSShooting{}, psq.predicates...),
		withPlayerStats: psq.withPlayerStats.Clone(),
		// clone intermediate query.
		sql:  psq.sql.Clone(),
		path: psq.path,
	}
}

// WithPlayerStats tells the query-builder to eager-load the nodes that are connected to
// the "playerStats" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *PSShootingQuery) WithPlayerStats(opts ...func(*PlayerStatsQuery)) *PSShootingQuery {
	query := (&PlayerStatsClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withPlayerStats = query
	return psq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Goals int `json:"Goals,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PSShooting.Query().
//		GroupBy(psshooting.FieldGoals).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (psq *PSShootingQuery) GroupBy(field string, fields ...string) *PSShootingGroupBy {
	psq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PSShootingGroupBy{build: psq}
	grbuild.flds = &psq.ctx.Fields
	grbuild.label = psshooting.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Goals int `json:"Goals,omitempty"`
//	}
//
//	client.PSShooting.Query().
//		Select(psshooting.FieldGoals).
//		Scan(ctx, &v)
func (psq *PSShootingQuery) Select(fields ...string) *PSShootingSelect {
	psq.ctx.Fields = append(psq.ctx.Fields, fields...)
	sbuild := &PSShootingSelect{PSShootingQuery: psq}
	sbuild.label = psshooting.Label
	sbuild.flds, sbuild.scan = &psq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PSShootingSelect configured with the given aggregations.
func (psq *PSShootingQuery) Aggregate(fns ...AggregateFunc) *PSShootingSelect {
	return psq.Select().Aggregate(fns...)
}

func (psq *PSShootingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psq); err != nil {
				return err
			}
		}
	}
	for _, f := range psq.ctx.Fields {
		if !psshooting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *PSShootingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PSShooting, error) {
	var (
		nodes       = []*PSShooting{}
		withFKs     = psq.withFKs
		_spec       = psq.querySpec()
		loadedTypes = [1]bool{
			psq.withPlayerStats != nil,
		}
	)
	if psq.withPlayerStats != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, psshooting.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PSShooting).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PSShooting{config: psq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := psq.withPlayerStats; query != nil {
		if err := psq.loadPlayerStats(ctx, query, nodes, nil,
			func(n *PSShooting, e *PlayerStats) { n.Edges.PlayerStats = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (psq *PSShootingQuery) loadPlayerStats(ctx context.Context, query *PlayerStatsQuery, nodes []*PSShooting, init func(*PSShooting), assign func(*PSShooting, *PlayerStats)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PSShooting)
	for i := range nodes {
		if nodes[i].player_stats_ps_shooting == nil {
			continue
		}
		fk := *nodes[i].player_stats_ps_shooting
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
			return fmt.Errorf(`unexpected foreign-key "player_stats_ps_shooting" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (psq *PSShootingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	_spec.Node.Columns = psq.ctx.Fields
	if len(psq.ctx.Fields) > 0 {
		_spec.Unique = psq.ctx.Unique != nil && *psq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *PSShootingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(psshooting.Table, psshooting.Columns, sqlgraph.NewFieldSpec(psshooting.FieldID, field.TypeInt))
	_spec.From = psq.sql
	if unique := psq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psq.path != nil {
		_spec.Unique = true
	}
	if fields := psq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, psshooting.FieldID)
		for i := range fields {
			if fields[i] != psshooting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *PSShootingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(psshooting.Table)
	columns := psq.ctx.Fields
	if len(columns) == 0 {
		columns = psshooting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psq.ctx.Unique != nil && *psq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PSShootingGroupBy is the group-by builder for PSShooting entities.
type PSShootingGroupBy struct {
	selector
	build *PSShootingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *PSShootingGroupBy) Aggregate(fns ...AggregateFunc) *PSShootingGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the selector query and scans the result into the given value.
func (psgb *PSShootingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psgb.build.ctx, "GroupBy")
	if err := psgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PSShootingQuery, *PSShootingGroupBy](ctx, psgb.build, psgb, psgb.build.inters, v)
}

func (psgb *PSShootingGroupBy) sqlScan(ctx context.Context, root *PSShootingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psgb.flds)+len(psgb.fns))
		for _, f := range *psgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PSShootingSelect is the builder for selecting fields of PSShooting entities.
type PSShootingSelect struct {
	*PSShootingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pss *PSShootingSelect) Aggregate(fns ...AggregateFunc) *PSShootingSelect {
	pss.fns = append(pss.fns, fns...)
	return pss
}

// Scan applies the selector query and scans the result into the given value.
func (pss *PSShootingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pss.ctx, "Select")
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PSShootingQuery, *PSShootingSelect](ctx, pss.PSShootingQuery, pss, pss.inters, v)
}

func (pss *PSShootingSelect) sqlScan(ctx context.Context, root *PSShootingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pss.fns))
	for _, fn := range pss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
