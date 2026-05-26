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

package utils

import (
	"time"
)

const (
	secondsInHour = 3600
)

// NowFunc type for function of getting current time
type NowFunc func() time.Time

// TimeIncrementer increment current time by configurable incremental
type TimeIncrementer struct {
	IncBySecond int64
	currentSec  int64
}

var nowFunc NowFunc

func init() {
	ResetClockImplementation()
}

// ResetClockImplementation resets implementation to use time.Now
func ResetClockImplementation() {
	_ = "STUB: not implemented"

	// SetClockImplementation sets implementation to use passed in nowFunc
	return
}

func SetClockImplementation(f NowFunc) {
	_ = "STUB: not implemented"

	// SetCurrentTime sets the clock implementation to the specified time,
	return
}

func SetCurrentTime(t time.Time) { _ = "STUB: not implemented"; return }

// Now returns current time using nowFunc
func Now() time.Time {
	_ = "STUB: not implemented"

	// FormatTimeStampToUTC formats a epoch timestamp to a time string in UTC time zone.
	return *new(time.Time)
}

func FormatTimeStampToUTC(ts int64) string { _ = "STUB: not implemented"; return "" }

// TimeStampToUTC converts a timestamp to a Time struct in UTC time zone.
func TimeStampToUTC(ts int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Now increment current time by one second at a time
func (r *TimeIncrementer) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// CrossDST tells whether a time range crosses a DST switch time for given zone
func CrossDST(fromTs, toTs int64, loc *time.Location) bool { _ = "STUB: not implemented"; return false }

// CalculateDSTSwitchTs calculates DST switch timestamp given a time range and a timezone
// it returns 0 if given range doesn't contain a switch for that zone
// it assumes the range won't contain more than 1 switch timestamp, otherwise it will
// return one of them (which one to return is not determined)
func CalculateDSTSwitchTs(fromTs, toTs int64, loc *time.Location) (switchTs int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// AdjustOffset adjusts timestamp value with time range start or end ts basing on DST switch ts
func AdjustOffset(fromOffset, toOffset int, switchTs, ts int64) int64 {
	_ = "STUB: not implemented"
	return 0
}
