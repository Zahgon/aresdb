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

package list

import (
	"unsafe"

	"github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/memstore/vectors"
)

// baseVectorParty is the shared struct for live store list vp and archive store
// list vp. Some difference between normal base vp and list base vp:
//  1. There is no mode concept for list vp. therefore nonDefaultValueCount does not make sense for
//     list vp.
//  2. There is no default value concept for list at least for now, so we will not store default value.
type baseVectorParty struct {
	// offset is a pair of uint32 [offset, length]. therefore its length is 2 * length of vp.
	offsets *vectors.Vector
	// length of vp.
	length int
	// DataType of values. We need it since for mode 0 vector party, we cannot
	// get data type from values vector. Also we store it on disk anyway.
	dataType common.DataType

	reporter HostMemoryChangeReporter
	//convenient function to call child class to retrieve data
	getDataValueFn func(int) common.DataValue
}

// GetOffsetLength returns the <offset, length> pair at ith row.
func (vp *baseVectorParty) GetOffsetLength(row int) (offset uint32, length uint32, valid bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

// SetOffsetLength update offset/length for nth fow
func (vp *baseVectorParty) SetOffsetLength(row int, offset, length unsafe.Pointer) {
	_ = "STUB: not implemented"
	return
}

// GetElemCount return the number of element for value in n-th row
func (vp *baseVectorParty) GetElemCount(row int) uint32 { _ = "STUB: not implemented"; return 0 }

// GetValidity get validity of given offset.
func (vp *baseVectorParty) GetValidity(row int) bool { _ = "STUB: not implemented"; return false }

// IsList tells whether this vp is list vp. And user can later on cast it to proper interface.
func (vp *baseVectorParty) IsList() bool {
	_ = "STUB: not implemented"

	// GetDataType returns the element date type of this vp.
	return false
}

func (vp *baseVectorParty) GetDataType() common.DataType {
	_ = "STUB: not implemented"
	return *

	// GetLength returns the length of the vp.
	new(common.DataType)
}

func (vp *baseVectorParty) GetLength() int {
	_ = "STUB: not implemented"

	// Slice vector party into human readable SlicedVector format. For now just return an
	// empty slice.
	return 0
}

func (vp *baseVectorParty) Slice(startRow, numRows int) common.SlicedVector {
	_ = "STUB: not implemented"
	return *new(common.SlicedVector)
}

// Check whether two vector parties are equal (used only in unit tests)
// Check common properties like data type and length.
// leftVP is vp underneath
func (vp *baseVectorParty) equals(other common.VectorParty, leftVP common.ListVectorParty) bool {
	_ = "STUB: not implemented"
	return false
}

func arrayValueCompare(dataType common.DataType, left, right unsafe.Pointer) bool {
	_ = "STUB: not implemented"
	return false
}

// GetNonDefaultValueCount get Number of non-default values stored. Since we
// count all list values as valid values, it should be equal to the length of
// the vp. If in future we want to get a count of non default element value
// count, we may need to scan all the old element values when overwriting.
func (vp *baseVectorParty) GetNonDefaultValueCount() int { _ = "STUB: not implemented"; return 0 }
