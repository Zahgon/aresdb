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

package common

import (
	"time"

	"github.com/uber/aresdb/query/expr"
)

var timeUnitMap = map[string]string{
	"year":         "y",
	"quarter":      "q",
	"month":        "M",
	"week":         "w",
	"day":          "d",
	"hour":         "h",
	"quarter-hour": "15m",
	"minute":       "m",
	"second":       "s",
}

// AlignedTime  is time that is calendar aligned to the unit.
type AlignedTime struct {
	Time time.Time `json:"time"`
	// Values for unit: y, q, M, w, d, {12, 8, 6, 4, 3, 2}h, h, {30, 20, 15, 12, 10, 6, 5, 4, 3, 2}m, m
	Unit string `json:"unit"`
}

// adjustMidnight adjusts daylight saving anomalies in a few timezones
// that return day boundary/midnight as either 23:00 (of the previous day) or 01:00.
//
// Fix for America/Sao_Paulo daylight saving starts (2016-10-16):
// The midnight of 2016-10-16 does not exist and time.Date returns 23:00 of the previous day.
//
// Must check whether the 1 hour rewind still gives the same day:
// For Asia/Beirut, this is not true for 2017-03-26, and true for 2017-03-27 and beyond.
func adjustMidnight(t time.Time) time.Time {
	_ = "STUB: not implemented"

	// Add one hour from 23:00 to 01:00 on the transition day;
	// and from 23:00 to 00:00 on non-transition days.
	return *new(time.Time)
}

// Must check whether the 1 hour rewind still gives the same day:
// For Asia/Beirut, this is false for 2017-03-26 (transition day), and true for 2017-03-27.

// ParseTimezone parses timezone
func ParseTimezone(timezone string) (*time.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCurrentCalendarUnit returns the start and end of the calendar unit for base.
func GetCurrentCalendarUnit(base time.Time, unit string) (start, end time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), nil
}

// Returns the start and end of the calendar `unit` that is `amount` `unit`s later from `base`.
func applyTimeOffset(base time.Time, amount int, unit string) (start, end time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), nil
}

// Round to hour.

// Apply the offset.

// Round to quarter-hour.

// Apply the offset.

// Round to minute.

// Apply the offset.

// Returns the start and end of the absolute calendar unit specified in `dateExpr` and `timeExpr`.
func parseAbsoluteTime(dateExpr, timeExpr string, location *time.Location) (start, end time.Time, unit string, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), "", nil
}

// Temporary hack until summary switch to use relative time expression.

// Returns the start and end of the calendar unit specified in `expression`.
func parseTimeFilterExpression(expression string, now time.Time) (start, end time.Time, unit string, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), "", nil
}

//we will assume data over 99999999999 will be timestamp in ms, and convert it to be in seconds

// Numbers above 9999999 are treated as timestamps, otherwise the corresponding Time object (of year 10000 and beyond)
// will fail JSON marshaling, and criples debugz.

// ParseTimeFilter parses time filter
func ParseTimeFilter(filter TimeFilter, loc *time.Location, now time.Time) (from, to *AlignedTime, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Populate to with now if from is present.

// CreateTimeFilterExpr creates time filter expr
func CreateTimeFilterExpr(expression expr.Expr, from, to *AlignedTime) (fromExpr, toExpr expr.Expr) {
	_ = "STUB: not implemented"
	return *new(expr.Expr), *new(expr.Expr)
}
