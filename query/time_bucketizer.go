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

import (
	"github.com/uber/aresdb/query/common"
	"github.com/uber/aresdb/query/expr"
)

// used to convert supported time units (string) to single char format
var bucketSizeToNormalized = map[string]string{
	"minutes": "m",
	"minute":  "m",
	"day":     "d",
	"hours":   "h",
	"hour":    "h",
}

// mapping from bucketizer to the functor to get the start of the bucketizer.
var irregularBucketizer2Functor = map[string]expr.Token{
	"month":   expr.GET_MONTH_START,
	"quarter": expr.GET_QUARTER_START,
	"year":    expr.GET_YEAR_START,
	"week":    expr.GET_WEEK_START,
}

// regularRecurringTimeBucketizer is in the format of "x of y" where y is a regular time interval of which number of
// seconds is fixed, e.g (week, day, hour). Y is called the bucket Size and x is called base Unit.
type regularRecurringTimeBucketizer struct {
	baseUnit   int
	bucketSize int
}

// mapping from regular recurring time bucketizer str to base Unit and bucket Size.
var tbStr2regularRecurringTimeBucketizer = map[string]regularRecurringTimeBucketizer{
	"time of day":  {baseUnit: 1, bucketSize: common.SecondsPerDay},
	"hour of day":  {baseUnit: common.SecondsPerHour, bucketSize: common.SecondsPerDay},
	"hour of week": {baseUnit: common.SecondsPerHour, bucketSize: common.SecondsPerWeek},
	"day of week":  {baseUnit: common.SecondsPerDay, bucketSize: common.SecondsPerWeek},
}

var irregularRecurringBucketizer2Functor = map[string]expr.Token{
	"day of month":    expr.GET_DAY_OF_MONTH,
	"day of year":     expr.GET_DAY_OF_YEAR,
	"month of year":   expr.GET_MONTH_OF_YEAR,
	"quarter of year": expr.GET_QUARTER_OF_YEAR,
}

// buildTimeDimensionExpr constructs sub ast based on several query params:
// the time bucketizer string and the timezone string.
// we parse time bucketizer into bucketInSeconds, for timezone string:
// if fixed (non-UTC) timezone is passed in, we extend the ast to `(timeColumn CONVERT_TZ fixed_timezone_offset) FLOOR bucketInSeconds`
// if timezoneColumn exists, we extend the ast to `(timeColumn CONVERT_TZ timezoneColumn) FLOOR bucketInSeconds`
func (qc *AQLQueryContext) buildTimeDimensionExpr(timeBucketizerString string, timeColumn expr.Expr) (expr.Expr, error) {
	_ = "STUB: not implemented"
	return *new(expr.Expr), nil
}

// construct TimeSeriesBucketizer expr

// expand ast by offsetting timezone column

// simulate IF statement. sub ast: timeCol + fromOffset + (timeCol > switchTs) * offsetDiff
// where (timeCol > switchTs) will return 1 or 0

// getRegularRecurringTimeBucketizer converts a time bucketizer string to a regularRecurringTimeBucketizer struct.
// Nil means it does not match.
func getRegularRecurringTimeBucketizer(tbStr string) (*regularRecurringTimeBucketizer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseRecurringTimeBucketizer parses the time bucketizer string into a composite expression tree if it's a recurring
// time bucketizer. The tree will be like floor((timeColumn % bucketSize), unitSize) if it's a regular recurring
// time bucketizer, e.g.(each individual time duration contains the same amount of seconds). Otherwise it will
// return a AST of a special function call. E.g. getDayOfMonth.
func parseRecurringTimeBucketizer(timeBucketizerString string, timeColumnExpr expr.Expr) (expr.Expr, error) {
	_ = "STUB: not implemented"
	return *new(expr.Expr), nil
}

// if bucket size is equal to week, we need to adjust it to Monday by subtracting number of seconds per
// 4 days (since 1970-01-01 is a Thursday).

// if base unit >= day, we need to divide it by the base unit.

// For division, everything is converted to float.

// parseIrregularTimeBucketizer parses the time bucketizer into a UnaryExpr with the corresponding functor as the OP
// node and original time column expression as the call argument. Return nil if it is not a irregular time series
// bucketizer.
func parseIrregularTimeBucketizer(timeBucketizerString string, timeColumn expr.Expr) expr.Expr {
	_ = "STUB: not implemented"
	return *new(expr.Expr)
}
