//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package query

// #include "time_series_aggregate.h"
import "C"

import (
	"github.com/uber/aresdb/cluster/topology"

	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/query/common"
	"github.com/uber/aresdb/query/expr"
)

const (
	unsupportedInputType      = "unsupported input type for %s: %s"
	defaultTimezoneTableAlias = "__timezone_lookup"
	geoShapeLimit             = 100
	nonAggregationQueryLimit  = 1000
)

// Compile compiles AQLQueryContext for data feeding and query
// execution. Caller should check for AQLQueryContext.Error.
func (qc *AQLQueryContext) Compile(tableSchemaReader memCom.TableSchemaReader, shardOwner topology.ShardOwner) {
	_ = "STUB: not implemented"
	// processTimezone might append additional joins
	return
}

// Read schema for every table used.

// Parse all other SQL expressions to ASTs.

// Resolve data types in the ASTs against schema, also translate enum values.

// Process join conditions first to collect information about geo join.

// Identify prefilters.

// Process filters.

// Process measure and dimensions.

// TODO: VM instruction generation

// adjustFilterToTimeFilter try to find one rowfilter to be time filter if there is no timefilter for fact table query
func (qc *AQLQueryContext) adjustFilterToTimeFilter() { _ = "STUB: not implemented"; return }

// check if this filter on main table event time column

// only support number literal or string literal

// only convert first LT

// only convert first GTE

// processTimeFilter will handle the from is nil case

// remove from original query filter

func (qc *AQLQueryContext) processJoinConditions() { _ = "STUB: not implemented"; return }

// we will extract the geo join out of the join conditions since we are going to handle geo intersects
// as filter instead of an equal join.

// matchGeoJoin initializes the GeoIntersection struct for later query process use. For now only one geo join is
// allowed per query. If users want to intersect with multiple geo join conditions, they should specify multiple geo
// shapeLatLongs in the geo filter.
// There are following constrictions:
// 1. At most one geo join condition.
// 2. Geo table must be dimension table.
// 3. The join condition must include exactly one shape column and one point column.
// 4. Exactly one geo filter should be specified.
// 5. Geo filter column must be the primary key of the geo table.
// 6. Geo UUIDs must be string in query.
// 7. Geo filter operator must be EQ or IN
// 8. Geo table's fields are not allowed in measures.
// 9. Only one geo dimension allowed.
func (qc *AQLQueryContext) matchGeoJoin(joinTableID int, mainTableSchema *memCom.TableSchema,
	joinSchema *memCom.TableSchema, conditions []expr.Expr) {
	_ = "STUB: not implemented"
	return
}

// one foreign table primary key columns only.

// guaranteed by query rewrite.

// Set column usage for geo points.

func isGeoJoin(j common.Join) bool { _ = "STUB: not implemented"; return false }

// list of join conditions enforced for now
// 1. equi-join only
// 2. many-to-one join only
// 3. foreign table must be a dimension table
// 4. one foreign table primary key columns only
// 5. foreign table primary key can have only one column
// 6. every foreign table must be joined directly to the main table, i.e. no bridges?
// 7. up to 8 foreign tables
func (qc *AQLQueryContext) matchEqualJoin(joinTableID int, joinSchema *memCom.TableSchema, conditions []expr.Expr) {
	_ = "STUB: not implemented"
	return
}

// foreign table must be a dimension table

// one foreign table primary key columns only

// equi-join only

// main table at left and foreign table at right

// every foreign table must be joined directly to the main table

// many-to-one join only (join with foreign table's primary key)

// set column usage for join column in main table
// no need to set usage for remote join column in foreign table since
// we only use primary key of foreign table to join

func (qc *AQLQueryContext) parseExprs() {
	_ = "STUB: not implemented"

	// Join conditions.
	return
}

// Filters.

// Dimensions.

// make sure time column is defined

// dimension is defined as sqlExpression

// Measures.

func (qc *AQLQueryContext) processTimezone() { _ = "STUB: not implemented"; return }

// append timezone table to joins

func (qc *AQLQueryContext) readSchema(tableSchemaReader memCom.TableSchemaReader, shardOwner topology.ShardOwner) {
	_ = "STUB: not implemented"
	return
}

// Main table.

// use user query specified shards
// or all shards it owns when user did not specify

// Archiving cutoff filter usage for fact table.

// Foreign tables.

// Prevent double locking.

// we will only support fact to fact join within same shard

// Archiving cutoff filter usage for fact table.

// for fact to dimension table join
// we can assume shard zero for dimension table
// since dimension table is not sharded

func (qc *AQLQueryContext) releaseSchema() { _ = "STUB: not implemented"; return }

// Rewrite walks the expresison AST and resolves data types bottom up.
// In addition it also translates enum strings and rewrites their predicates.
func (qc *AQLQueryContext) Rewrite(expression expr.Expr) expr.Expr {
	_ = "STUB: not implemented"
	return *new(expr.Expr)
}

// resolveTypes walks all expresison ASTs and resolves data types bottom up.
// In addition it also translates enum strings and rewrites their predicates.
func (qc *AQLQueryContext) resolveTypes() {
	_ = "STUB: not implemented"
	// Join conditions.
	return
}

// Dimensions.

// Measures.

// Filters.

// extractFitler processes the specified query level filter and matches it
// against the following formats:
//
//	column = value
//	column > value
//	column >= value
//	column < value
//	column <= value
//	column
//	not column
//
// It returns the numeric constant value associated with the filter in a uint32
// space (for all types including float32).
// In addition it also returns the boundaryType for >, >=, <, <= operators.
// Note that since the candidate filters have already been preselected against
// some criterias, this function does not perform full format validation.
func (qc *AQLQueryContext) extractFilter(filterID int) (
	value uint32, boundary boundaryType, success bool) {
	_ = "STUB: not implemented"
	return 0, *new(boundaryType), false
}

// Match `column` format

// Match `not column` format

// Match `column op value` format

// matchPrefilters identifies all prefilters from query level filters,
// stores them in AQLQueryContext.Prefilters,
// and stores their values in TableScanner for future prefilter vector slicing.
func (qc *AQLQueryContext) matchPrefilters() {
	_ = "STUB: not implemented"
	// Format of candidateFilters:
	// [tableID]map[columnID]{filterIDs for lower bound, upper bound, equality}
	// tableID is query scoped, while columnID is schema scoped.
	return
}

// Index candidate filters by table/column

// Match `column` format

// Match `not column` format

// TODO: IS_NULL can be matched as an equality filter.
// TODO: IS_NOT_NULL can be matched as the final range filter.

// Match `column op value` format, where op can be =, <, <=, >, >=.

// Prefilter matching

// Match in archiving sort column order

// Stop on first missing column

// Equality

// Stop if the value fails to be extracted

// Continue matching the next column

// Lower bound

// Upper bound

// Stop after the first range filter

// columnUsageCollector is the visitor used to traverses an AST, finds VarRef columns
// and sets the usage bits in tableScanners. The VarRef nodes must have already
// been resolved and annotated with TableID and ColumnID.
type columnUsageCollector struct {
	tableScanners []*TableScanner
	usages        columnUsage
}

func (c columnUsageCollector) Visit(expression expr.Expr) expr.Visitor {
	_ = "STUB: not implemented"
	return *new(expr.Visitor)
}

// foreignTableColumnDetector detects foreign table columns involved in AST
type foreignTableColumnDetector struct {
	hasForeignTableColumn bool
}

func (c *foreignTableColumnDetector) Visit(expression expr.Expr) expr.Visitor {
	_ = "STUB: not implemented"
	return *new(expr.Visitor)
}

// processFilters processes all filters and categorize them into common filters,
// prefilters, and time filters. It also collect column usages from the filters.
func (qc *AQLQueryContext) processFilters() {
	_ = "STUB: not implemented"
	// OOPK engine only supports one measure per query.
	return
}

// Categorize common filters and prefilters based on matched prefilters.

// common filters

// Process time filter.

// Collect column usages from the filters.

func getStrFromNumericalOrStrLiteral(e expr.Expr) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// matchGeoFilter tries to match the filter as a geo filter and prepare shapeUUIDs for aql processor. It returns whether
// the filterExpr is a geo filter.
func (qc *AQLQueryContext) matchGeoFilter(filterExpr expr.Expr, joinTableID int,
	joinSchema *memCom.TableSchema, geoFilterFound bool) (geoFilterFoundInCurrentExpr bool) {
	_ = "STUB: not implemented"
	return false
}

func (qc *AQLQueryContext) matchGeoFilterHelper(filterExpr expr.Expr, joinTableID int,
	joinSchema *memCom.TableSchema, shapeUUIDs *[]string) (inValidOpFound, foundGeoFilter bool) {
	_ = "STUB: not implemented"
	return false, false
}

// geo filter's column must be primary key.

// keep traversing to find geo fields

// processTimeFilter processes the time filter by matching it against the time
// column of the main fact table. The time filter will be identified as common
// filter if it does not match with the designated time column.
func (qc *AQLQueryContext) processTimeFilter() { _ = "STUB: not implemented"; return }

// Match against time column of the main fact table.

// TODO: resolve time filter column against foreign tables.

// Validate column existence and type.

// matchAndRewriteGeoDimension tells whether a dimension matches geo join and whether it's a valid
// geo join. It returns the rewritten geo dimension and error. If the err is non nil, it means it's a invalid geo join.
// A valid geo dimension can only in one of the following format:
//  1. UUID
//  2. hex(UUID)
func (qc *AQLQueryContext) matchAndRewriteGeoDimension(dimExpr expr.Expr) (expr.Expr, error) {
	_ = "STUB: not implemented"
	return *new(expr.Expr), nil
}

// geoTableUsageCollector traverses an AST expression tree, finds VarRef columns
// and check whether it uses any geo table columns.
type geoTableUsageCollector struct {
	geoIntersection geoIntersection
	useGeoTable     bool
}

func (g *geoTableUsageCollector) Visit(expression expr.Expr) expr.Visitor {
	_ = "STUB: not implemented"
	return *new(expr.Visitor)
}

// arrayColumnUsageCollector traverses an AST expression tree, finds VarRef columns
// and check whether it uses any array column
type arrayColumnUsageCollector struct {
	useArrayColumn bool
}

func (ac *arrayColumnUsageCollector) Visit(expression expr.Expr) expr.Visitor {
	_ = "STUB: not implemented"
	return *new(expr.Visitor)
}

func (qc *AQLQueryContext) processMeasure() {
	_ = "STUB: not implemented"
	// OOPK engine only supports one measure per query.
	return
}

// in case user forgot to provide limit

// Match and strip the aggregate function.

// check if any array column is used in measure

// default is 4 bytes

// 4 bytes for storing average result and another 4 byte for count

// for average, we should always use float type as the agg type.

func (qc *AQLQueryContext) getAllColumnsDimension() (columns []common.Dimension) {
	_ = "STUB: not implemented"
	// only main table columns wildcard match supported
	return nil
}

// no geoshape and array type directly supported as dimension

func (qc *AQLQueryContext) processDimensions() {
	_ = "STUB: not implemented"
	// Copy dimension ASTs.
	return
}

// TODO: support numeric bucketizer.

// array column can not be used as dimension directly

// Check whether measure and dimensions are referencing any geo table columns.

// Collect column usage from measure and dimensions

// Sort dimension columns based on the data width in bytes
// dimension columns in OOPK will not be reordered, but a mapping
// from original id to ordered offsets (value and validity) in
// dimension vector will be stored.
// GeoUUID dimension will be 1 bytes. VarRef expression will use column data length,
// others will be default to 4 bytes.
func (qc *AQLQueryContext) sortDimensionColumns() { _ = "STUB: not implemented"; return }

// record value offset, null offset pair
// null offsets will have to add total dim bytes later

// plus one byte per dimension column for validity

// no dimension size checking for non-aggregation query

func (qc *AQLQueryContext) sortUsedColumns() { _ = "STUB: not implemented"; return }

// Unsorted/uncompressed columns

// Sorted/compressed columns

func parseTimezoneColumnString(timezoneColumnString string) (column, joinKey string, success bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (qc *AQLQueryContext) expandINop(e *expr.BinaryExpr) (expandedExpr expr.Expr) {
	_ = "STUB: not implemented"
	return *new(expr.Expr)
}
