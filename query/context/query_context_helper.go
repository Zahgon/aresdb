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

package context

import (
	"github.com/uber/aresdb/query/expr"
)

// QueryContextHelper is a helper Class to group common code
// for broker query compiler and datanode query compiler
type QueryContextHelper struct {
	QCOptions QueryContextOptions
}

// NormalizeAndFilters AND filter
func (qc *QueryContextHelper) NormalizeAndFilters(filters []expr.Expr) []expr.Expr {
	_ = "STUB: not implemented"
	return nil
}

// resolveColumn resolves the VarRef identifier against the schema,
// and returns the matched tableID (query scoped) and columnID (schema scoped).
func (qc *QueryContextHelper) ResolveColumn(identifier string) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func blockNumericOpsForColumnOverFourBytes(token expr.Token, expressions ...expr.Expr) error {
	_ = "STUB: not implemented"
	return nil
}

func blockInt64(expressions ...expr.Expr) error { _ = "STUB: not implemented"; return nil }

func (qc *QueryContextHelper) expandINop(e *expr.BinaryExpr) (expandedExpr expr.Expr) {
	_ = "STUB: not implemented"
	return *new(expr.Expr)
}

// Rewrite walks the expresison AST and resolves data types bottom up.
// In addition it also translates enum strings and rewrites their predicates.
func (qc *QueryContextHelper) Rewrite(expression expr.Expr) expr.Expr {
	_ = "STUB: not implemented"
	return *new(expr.Expr)
}

// Strip parenthesis from the input

// Normalize the operator.

// Upgrade to signed.

// Strip IS_TRUE if child is already boolean.

// Rewrite to NOT(NOT(child)).

// Cast child to unsigned.

// Cast child to unsigned.

// Cast child to unsigned.

// TODO: @shz support int64 binary transform

// Upgrade and cast to highestType.

// For lhs - rhs, upgrade to signed at least.

// Upgrade and cast to highestType.

// Upgrade and cast to float.

// Cast to unsigned.

// Cast to boolean.

// Cast to boolean.

// swap lhs and rhs if rhs is VarRef but lhs is not.

// Match enum = 'case' and enum != 'case'.

// rhs is bool

// rhs is string enum

// Enum dictionary translation

// Combination of nullable data with not/and/or operators on top makes
// short circuiting hard.
// To play it safe we match against an invalid value.

// Cast to highestType.

// dayofweek from ts: (ts / secondsInDay + 4) % 7 + 1
// ref: https://dev.mysql.com/doc/refman/5.5/en/date-and-time-functions.html#function_dayofweek

// offset for

// no-op, this will be over written

// for now, only the following format is allowed for backward compatibility
// from_unixtime(time_col / 1000)

// hour(ts) = (ts % secondsInDay) / secondsInHour

// list of literals, no need to cast it for now.

// Switch geo point so that lhs is geo shape and rhs is geo point

// 1. noop when column itself is hll column
// 2. compute hll on the fly when column is not hll column

// For avg, the expression type should always be float.

// validate first argument

// build rhs literal

// if the request is from broker, it should be already a number literal

// Enum dictionary translation

// Combination of nullable data with not/and/or operators on top makes
// short circuiting hard.
// To play it safe we match against an invalid value.

// Cast else and thens to highestType, cast whens to boolean.
