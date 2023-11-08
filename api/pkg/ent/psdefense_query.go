// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/playerstats"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/predicate"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent/psdefense"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PSDefenseQuery is the builder for querying PSDefense entities.
type PSDefenseQuery struct {
	config
	ctx             *QueryContext
	order           []psdefense.OrderOption
	inters          []Interceptor
	predicates      []predicate.PSDefense
	withPlayerStats *PlayerStatsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PSDefenseQuery builder.
func (pdq *PSDefenseQuery) Where(ps ...predicate.PSDefense) *PSDefenseQuery {
	pdq.predicates = append(pdq.predicates, ps...)
	return pdq
}

// Limit the number of records to be returned by this query.
func (pdq *PSDefenseQuery) Limit(limit int) *PSDefenseQuery {
	pdq.ctx.Limit = &limit
	return pdq
}

// Offset to start from.
func (pdq *PSDefenseQuery) Offset(offset int) *PSDefenseQuery {
	pdq.ctx.Offset = &offset
	return pdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pdq *PSDefenseQuery) Unique(unique bool) *PSDefenseQuery {
	pdq.ctx.Unique = &unique
	return pdq
}

// Order specifies how the records should be ordered.
func (pdq *PSDefenseQuery) Order(o ...psdefense.OrderOption) *PSDefenseQuery {
	pdq.order = append(pdq.order, o...)
	return pdq
}

// QueryPlayerStats chains the current query on the "playerStats" edge.
func (pdq *PSDefenseQuery) QueryPlayerStats() *PlayerStatsQuery {
	query := (&PlayerStatsClient{config: pdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(psdefense.Table, psdefense.FieldID, selector),
			sqlgraph.To(playerstats.Table, playerstats.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, psdefense.PlayerStatsTable, psdefense.PlayerStatsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PSDefense entity from the query.
// Returns a *NotFoundError when no PSDefense was found.
func (pdq *PSDefenseQuery) First(ctx context.Context) (*PSDefense, error) {
	nodes, err := pdq.Limit(1).All(setContextOp(ctx, pdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{psdefense.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pdq *PSDefenseQuery) FirstX(ctx context.Context) *PSDefense {
	node, err := pdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PSDefense ID from the query.
// Returns a *NotFoundError when no PSDefense ID was found.
func (pdq *PSDefenseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pdq.Limit(1).IDs(setContextOp(ctx, pdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{psdefense.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pdq *PSDefenseQuery) FirstIDX(ctx context.Context) int {
	id, err := pdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PSDefense entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PSDefense entity is found.
// Returns a *NotFoundError when no PSDefense entities are found.
func (pdq *PSDefenseQuery) Only(ctx context.Context) (*PSDefense, error) {
	nodes, err := pdq.Limit(2).All(setContextOp(ctx, pdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{psdefense.Label}
	default:
		return nil, &NotSingularError{psdefense.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pdq *PSDefenseQuery) OnlyX(ctx context.Context) *PSDefense {
	node, err := pdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PSDefense ID in the query.
// Returns a *NotSingularError when more than one PSDefense ID is found.
// Returns a *NotFoundError when no entities are found.
func (pdq *PSDefenseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pdq.Limit(2).IDs(setContextOp(ctx, pdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{psdefense.Label}
	default:
		err = &NotSingularError{psdefense.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pdq *PSDefenseQuery) OnlyIDX(ctx context.Context) int {
	id, err := pdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PSDefenses.
func (pdq *PSDefenseQuery) All(ctx context.Context) ([]*PSDefense, error) {
	ctx = setContextOp(ctx, pdq.ctx, "All")
	if err := pdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PSDefense, *PSDefenseQuery]()
	return withInterceptors[[]*PSDefense](ctx, pdq, qr, pdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pdq *PSDefenseQuery) AllX(ctx context.Context) []*PSDefense {
	nodes, err := pdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PSDefense IDs.
func (pdq *PSDefenseQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pdq.ctx.Unique == nil && pdq.path != nil {
		pdq.Unique(true)
	}
	ctx = setContextOp(ctx, pdq.ctx, "IDs")
	if err = pdq.Select(psdefense.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pdq *PSDefenseQuery) IDsX(ctx context.Context) []int {
	ids, err := pdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pdq *PSDefenseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pdq.ctx, "Count")
	if err := pdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pdq, querierCount[*PSDefenseQuery](), pdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pdq *PSDefenseQuery) CountX(ctx context.Context) int {
	count, err := pdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pdq *PSDefenseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pdq.ctx, "Exist")
	switch _, err := pdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pdq *PSDefenseQuery) ExistX(ctx context.Context) bool {
	exist, err := pdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PSDefenseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pdq *PSDefenseQuery) Clone() *PSDefenseQuery {
	if pdq == nil {
		return nil
	}
	return &PSDefenseQuery{
		config:          pdq.config,
		ctx:             pdq.ctx.Clone(),
		order:           append([]psdefense.OrderOption{}, pdq.order...),
		inters:          append([]Interceptor{}, pdq.inters...),
		predicates:      append([]predicate.PSDefense{}, pdq.predicates...),
		withPlayerStats: pdq.withPlayerStats.Clone(),
		// clone intermediate query.
		sql:  pdq.sql.Clone(),
		path: pdq.path,
	}
}

// WithPlayerStats tells the query-builder to eager-load the nodes that are connected to
// the "playerStats" edge. The optional arguments are used to configure the query builder of the edge.
func (pdq *PSDefenseQuery) WithPlayerStats(opts ...func(*PlayerStatsQuery)) *PSDefenseQuery {
	query := (&PlayerStatsClient{config: pdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pdq.withPlayerStats = query
	return pdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TacklesTotal int `json:"tacklesTotal,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PSDefense.Query().
//		GroupBy(psdefense.FieldTacklesTotal).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pdq *PSDefenseQuery) GroupBy(field string, fields ...string) *PSDefenseGroupBy {
	pdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PSDefenseGroupBy{build: pdq}
	grbuild.flds = &pdq.ctx.Fields
	grbuild.label = psdefense.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TacklesTotal int `json:"tacklesTotal,omitempty"`
//	}
//
//	client.PSDefense.Query().
//		Select(psdefense.FieldTacklesTotal).
//		Scan(ctx, &v)
func (pdq *PSDefenseQuery) Select(fields ...string) *PSDefenseSelect {
	pdq.ctx.Fields = append(pdq.ctx.Fields, fields...)
	sbuild := &PSDefenseSelect{PSDefenseQuery: pdq}
	sbuild.label = psdefense.Label
	sbuild.flds, sbuild.scan = &pdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PSDefenseSelect configured with the given aggregations.
func (pdq *PSDefenseQuery) Aggregate(fns ...AggregateFunc) *PSDefenseSelect {
	return pdq.Select().Aggregate(fns...)
}

func (pdq *PSDefenseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pdq); err != nil {
				return err
			}
		}
	}
	for _, f := range pdq.ctx.Fields {
		if !psdefense.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pdq.path != nil {
		prev, err := pdq.path(ctx)
		if err != nil {
			return err
		}
		pdq.sql = prev
	}
	return nil
}

func (pdq *PSDefenseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PSDefense, error) {
	var (
		nodes       = []*PSDefense{}
		_spec       = pdq.querySpec()
		loadedTypes = [1]bool{
			pdq.withPlayerStats != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PSDefense).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PSDefense{config: pdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pdq.withPlayerStats; query != nil {
		if err := pdq.loadPlayerStats(ctx, query, nodes,
			func(n *PSDefense) { n.Edges.PlayerStats = []*PlayerStats{} },
			func(n *PSDefense, e *PlayerStats) { n.Edges.PlayerStats = append(n.Edges.PlayerStats, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pdq *PSDefenseQuery) loadPlayerStats(ctx context.Context, query *PlayerStatsQuery, nodes []*PSDefense, init func(*PSDefense), assign func(*PSDefense, *PlayerStats)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*PSDefense)
	nids := make(map[int]map[*PSDefense]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(psdefense.PlayerStatsTable)
		s.Join(joinT).On(s.C(playerstats.FieldID), joinT.C(psdefense.PlayerStatsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(psdefense.PlayerStatsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(psdefense.PlayerStatsPrimaryKey[1]))
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
					nids[inValue] = map[*PSDefense]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*PlayerStats](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "playerStats" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pdq *PSDefenseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pdq.querySpec()
	_spec.Node.Columns = pdq.ctx.Fields
	if len(pdq.ctx.Fields) > 0 {
		_spec.Unique = pdq.ctx.Unique != nil && *pdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pdq.driver, _spec)
}

func (pdq *PSDefenseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(psdefense.Table, psdefense.Columns, sqlgraph.NewFieldSpec(psdefense.FieldID, field.TypeInt))
	_spec.From = pdq.sql
	if unique := pdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pdq.path != nil {
		_spec.Unique = true
	}
	if fields := pdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, psdefense.FieldID)
		for i := range fields {
			if fields[i] != psdefense.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pdq *PSDefenseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pdq.driver.Dialect())
	t1 := builder.Table(psdefense.Table)
	columns := pdq.ctx.Fields
	if len(columns) == 0 {
		columns = psdefense.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pdq.sql != nil {
		selector = pdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pdq.ctx.Unique != nil && *pdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pdq.predicates {
		p(selector)
	}
	for _, p := range pdq.order {
		p(selector)
	}
	if offset := pdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PSDefenseGroupBy is the group-by builder for PSDefense entities.
type PSDefenseGroupBy struct {
	selector
	build *PSDefenseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pdgb *PSDefenseGroupBy) Aggregate(fns ...AggregateFunc) *PSDefenseGroupBy {
	pdgb.fns = append(pdgb.fns, fns...)
	return pdgb
}

// Scan applies the selector query and scans the result into the given value.
func (pdgb *PSDefenseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pdgb.build.ctx, "GroupBy")
	if err := pdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PSDefenseQuery, *PSDefenseGroupBy](ctx, pdgb.build, pdgb, pdgb.build.inters, v)
}

func (pdgb *PSDefenseGroupBy) sqlScan(ctx context.Context, root *PSDefenseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pdgb.fns))
	for _, fn := range pdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pdgb.flds)+len(pdgb.fns))
		for _, f := range *pdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PSDefenseSelect is the builder for selecting fields of PSDefense entities.
type PSDefenseSelect struct {
	*PSDefenseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pds *PSDefenseSelect) Aggregate(fns ...AggregateFunc) *PSDefenseSelect {
	pds.fns = append(pds.fns, fns...)
	return pds
}

// Scan applies the selector query and scans the result into the given value.
func (pds *PSDefenseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pds.ctx, "Select")
	if err := pds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PSDefenseQuery, *PSDefenseSelect](ctx, pds.PSDefenseQuery, pds, pds.inters, v)
}

func (pds *PSDefenseSelect) sqlScan(ctx context.Context, root *PSDefenseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pds.fns))
	for _, fn := range pds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
