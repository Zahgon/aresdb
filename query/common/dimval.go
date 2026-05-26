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
	"unsafe"

	memCom "github.com/uber/aresdb/memstore/common"
)

const (
	// string representing null dimension values
	NULLString = "NULL"
)

// DimCountsPerDimWidth defines dimension counts per dimension width
// 16-byte 8-byte 4-byte 2-byte 1-byte
type DimCountsPerDimWidth [5]uint8

// ReadDimension reads a dimension value given the index and corresponding data type of node.
// tzRemedy is used to remedy the timezone offset
func ReadDimension(valueStart, nullStart unsafe.Pointer,
	index int, dataType memCom.DataType, enumReverseDict []string, meta *TimeDimensionMeta, cache map[TimeDimensionMeta]map[int64]string) *string {
	_ = "STUB: not implemented"
	return nil
}

// check for nulls

// determine value width in bytes

// read intValue; handle float and signed types

// in case time dimension value was converted to float for division

// Should never happen.

// translate enum case back to string for unsigned types

// formatWithDataValue formats value with given type
func formatWithDataValue(valuePtr unsafe.Pointer, dataType memCom.DataType) *string {
	_ = "STUB: not implemented"
	return nil
}

// GetDimensionStartOffsets calculates the value and null starting position for given dimension inside dimension vector
// dimIndex is the ordered index of given dimension inside the dimension vector
func GetDimensionStartOffsets(numDimsPerDimWidth DimCountsPerDimWidth, dimIndex int, length int) (valueOffset, nullOffset int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// found which range this dimension vector belongs to

// dimBytes /= 2

func formatTimeDimension(val int64, meta TimeDimensionMeta, cache map[TimeDimensionMeta]map[int64]string) (result string) {
	_ = "STUB: not implemented"
	// We will not process timeUnit for application/hll because if application/hll holds the raw uint32
	// value. If we convert it to milliseconds, it will overflow.
	return ""
}

// skip timezone table dims
// TODO(shz): support timezone table dims

// 1970-01-01 was a Thursday
