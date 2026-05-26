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

package vectors

// #include <stdlib.h>
// #include <string.h>

import (
	"unsafe"

	"github.com/uber/aresdb/memstore/common"
)

// Vector stores a batch of columnar data (values, nulls, or counts) for a column.
type Vector struct {
	// The data type of the value stored in the vector.
	DataType common.DataType
	CmpFunc  common.CompareFunc

	// Max number of values that can be stored in the vector.
	Size int
	// Allocated size of the vector in bytes.
	Bytes int

	// Number of bits occupied per unit, possible values: 1, 8, 16, 32, 64, 128.
	unitBits int
	// Pointer to the vector buffer.
	buffer uintptr

	// **All following fields only works for live batch's vectors.**

	// Min and Max values seen, only used for time columns of fact tables.
	minValue uint32
	maxValue uint32
	// Number of trues in a bool typed vector.
	numTrues int
}

// NewVector creates a vector with the specified bits per unit and size(capacity).
// The majority of its storage space is managed in C.
func NewVector(dataType common.DataType, size int) *Vector { _ = "STUB: not implemented"; return nil }

// CalculateVectorBytes calculates bytes the vector will occupy given data type and size without actual allocation.
func CalculateVectorBytes(dataType common.DataType, size int) int {
	_ = "STUB: not implemented"
	return 0
}

// Round up to 512 bits (64 bytes).

// CalculateVectorPartyBytes calculates bytes the vector party will occupy.
// Note: data type supported in go memory will report memory usage when value is actually set, therefore report 0 here
func CalculateVectorPartyBytes(dataType common.DataType, size int, hasNulls bool, hasCounts bool) int {
	_ = "STUB: not implemented"
	return 0
}

// batchSize * size of golang pointer

// this only calculates the offset and caps for list live vector party, value vector is controlled inside vp
// list archive vector party can not use this either

// SafeDestruct destructs this vector's storage space managed in C.
func (v *Vector) SafeDestruct() { _ = "STUB: not implemented"; return }

// Buffer returns the pointer to the underlying buffer.
func (v *Vector) Buffer() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

// SetBool sets the bool value for the specified index.
func (v *Vector) SetBool(index int, value bool) { _ = "STUB: not implemented"; return }

// SetValue sets the data value for the specified index.
// index bound is not checked!
// data points to a buffer (in UpsertBatch for instance) that contains the value to be set.
func (v *Vector) SetValue(index int, data unsafe.Pointer) { _ = "STUB: not implemented"; return }

// GetBool returns the bool value for the specified index.
// index bound is not checked!
func (v *Vector) GetBool(index int) bool { _ = "STUB: not implemented"; return false }

// GetValue returns the data value for the specified index.
// index bound is not checked!
// The return value points to the internal buffer location that stores the value.
func (v *Vector) GetValue(index int) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// GetMinValue return the min value of the Vector Party
func (v *Vector) GetMinValue() uint32 {
	_ = "STUB: not implemented"

	// GetMaxValue return the max value of the Vector Party
	return 0
}

func (v *Vector) GetMaxValue() uint32 {
	_ = "STUB: not implemented"

	// LowerBound returns the index of the first element in vector[first, last) that is greater or equal
	// to the given value. The result is only valid if vector[first, last) is fully sorted in ascendant
	// order. If all values in the given range is less than the given value, LowerBound
	// returns last.
	// Note that first/last is not checked against vector bound.
	return 0
}

func (v *Vector) LowerBound(first int, last int, value unsafe.Pointer) int {
	_ = "STUB: not implemented"
	return 0
}

// UpperBound returns the index of the first element in vector[first, last) that is greater than the
// given value. The result is only valid if vector[first, last) is fully sorted in ascendant
// order. If all values in the given range is less than the given value, LowerBound returns last.
// Note that first/last is not checked against vector bound.
func (v *Vector) UpperBound(first int, last int, value unsafe.Pointer) int {
	_ = "STUB: not implemented"
	return 0
}

// GetSliceBytesAligned calculate the number of bytes of a slice of the vector,
// represented by [lowerBound, upperBound),
// aligned to 64-byte
// return the buffer pointer, new start index (start entry in vector), and length in bytes
func (v *Vector) GetSliceBytesAligned(lowerBound int, upperBound int) (buffer unsafe.Pointer, startIndex int, bytes int) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), 0, 0
}

// find the latest 64-byte aligned boundary

// SetAllValid set all bits to be 1 in a bool typed vector.
func (v *Vector) SetAllValid() {
	_ = "STUB: not implemented"
	// buffer are 64 bytes aligned so we can assign word by word.
	return
}

// CheckAllValid checks whether all bits are 1 in a bool typed vector.
func (v *Vector) CheckAllValid() bool { _ = "STUB: not implemented"; return false }

// First check word by word.

// then we check byte by byte.

// check remaining bits.
