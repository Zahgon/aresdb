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

package memstore

import (
	"io"
	"os"
	"unsafe"

	"github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/memstore/vectors"
)

// TransferableVectorParty is vector party that can be transferred to gpu for processing
type TransferableVectorParty interface {
	// GetHostVectorPartySlice slice vector party between [startIndex, startIndex+length) before transfer to gpu
	GetHostVectorPartySlice(startIndex, length int) common.HostVectorPartySlice
}

// baseVectorParty is the base vector party type
type baseVectorParty struct {
	// DataType of values. We need it since for mode 0 vector party, we cannot
	// get data type from values vector. Also we store it on disk anyway.
	dataType common.DataType
	// Number of non-default values stored, not always the same as Nulls.numTrues.
	nonDefaultValueCount int
	// Length/Size of each vector (not necessarily number of records).
	length int

	// Following fields are initialized during struct initialization.
	// DefaultValue for this column. For convenience.
	defaultValue common.DataValue
}

// cVectorParty combines the value, null, and count vector of a column in a batch.
// Some of the vectors can be nil, following the same modes described at
// https://github.com/uber/aresdb/wiki/VectorStore#vector-party
type cVectorParty struct {
	baseVectorParty

	// Set during archiving/backfill and stored on disk.
	columnMode common.ColumnMode

	values *vectors.Vector
	// Stores the validity bitmap (0 means null) for each value in values.
	nulls *vectors.Vector
	// Stores the accumulative count from the beginning of the vector
	// to the current position. Its length is vp.Length + 1 with first value to
	// be 0 and last value to be vp.Length. We can get a count of current value
	// by Counts[i+1] - Counts[i] for Values[i]
	counts *vectors.Vector
}

// IsList tells whether it's a list vector party or not.
func (vp *baseVectorParty) IsList() bool {
	_ = "STUB: not implemented"

	// AsList returns ListVectorParty representation of this vector party.
	// Caller should always call IsList before conversion, otherwise panic may happens
	// for incompatible vps.
	return false
}

func (vp *baseVectorParty) AsList() common.ListVectorParty {
	_ = "STUB: not implemented"
	return *new(common.ListVectorParty)
}

// GetLength returns the length this vector party
func (vp *baseVectorParty) GetLength() int {
	_ = "STUB: not implemented"

	// GetDataType returns the min and max value of this vector party
	return 0
}

func (vp *baseVectorParty) GetDataType() common.DataType {
	_ = "STUB: not implemented"
	return *

	// GetNonDefaultValueCount get Number of non-default values stored
	new(common.DataType)
}

func (vp *baseVectorParty) GetNonDefaultValueCount() int { _ = "STUB: not implemented"; return 0 }

// GetMode returns the stored column mode of this vector party.
func (vp *cVectorParty) GetMode() common.ColumnMode {
	_ = "STUB: not implemented"
	return *

	// fillWithDefaultValue fills the values vector and nulls vector with default value
	// if it's valid. **It should be only called when both values vector and nulls vector
	// presents, otherwise it will panic. Also it requires both value vector and null vector
	// is initialized with zeros and have never been touched yet**
	new(common.ColumnMode)
}

func (vp *cVectorParty) fillWithDefaultValue() { _ = "STUB: not implemented"; return }

// SafeDestruct destructs all vectors of this vector party. Corresponding pointer should be set
// as nil after destruction.
func (vp *cVectorParty) SafeDestruct() { _ = "STUB: not implemented"; return }

// GetBytes returns space occupied by this vector party.
func (vp *cVectorParty) GetBytes() int64 { _ = "STUB: not implemented"; return 0 }

// setValidity set the validity of given offset and update NonDefaultValueCount.
// Third parameter count should only be passed for compressed columns. If
// not passed, the default value is 1.
func (vp *cVectorParty) setValidity(offset int, valid bool) { _ = "STUB: not implemented"; return }

// GetValidity implements GetValidity in cVectorParty
func (vp *cVectorParty) GetValidity(offset int) bool { _ = "STUB: not implemented"; return false }

// GetDataValue returns the DataValue for the specified index.
// It first check validity of the value, then it check whether it's a
// boolean column to decide whether to load bool value or other value
// type. Index bound is not checked!
func (vp *cVectorParty) GetDataValue(offset int) common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

// GetDataValueByRow implements GetDataValueByRow in cVectorParty
func (vp *cVectorParty) GetDataValueByRow(row int) common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

// SetDataValue implements SetDataValue in cVectorParty
func (vp *cVectorParty) SetDataValue(offset int, value common.DataValue, countsUpdateMode common.ValueCountsUpdateMode, counts ...uint32) {
	_ = "STUB: not implemented"
	return
}

// JudgeMode judges column mode of current vector party according to value count fields.
func (vp *cVectorParty) JudgeMode() common.ColumnMode {
	_ = "STUB: not implemented"
	return *new(common.ColumnMode)
}

// both

// compressed columns.

// no null vector.

// uncompressed columns

// Equals checks whether two vector parties are the same. **Only for unit test use.**
func (vp *cVectorParty) Equals(other common.VectorParty) bool {
	_ = "STUB: not implemented"
	return false
}

// check vector elements

// compare first count
// usually this is not needed since first count should always be 0

// only compare next count

// Slice slice the vector party into the interval of [startRow, startRow+numRows)
func (vp *cVectorParty) Slice(startRow int, numRows int) (vector common.SlicedVector) {
	_ = "STUB: not implemented"
	return *

	// size is the number of entries in the vector,
	// size != numRows when compressed,
	// although here size is initialized as if vector is uncompressed.
	new(common.SlicedVector)
}

// find the indexes [beginIndex, endIndex) based on [startRow, startRow + numRows)

// subtract endIndex by 1 when endIndex points to vp.length+1

// compressed

// uncompressed

// SliceByValue returns a subrange withing [lowerBoundRow, upperBoundRow) that matches the specified value
func (vp *cVectorParty) SliceByValue(lowerBoundRow, upperBoundRow int, value unsafe.Pointer) (
	startRow int, endRow int, startIndex int, endIndex int) {
	_ = "STUB: not implemented"
	// has counts
	return 0, 0, 0, 0
}

// TODO: check whether value itself is null, for IS_NULL
// return as if the slice is empty [upperBound, upperBound)

// If the default value is equal to the value, we return the whole slice.

// subtract endIndex by 1 when endIndex points to vp.length+1

// SliceIndex returns the startIndex and endIndex of the vector party slice given startRow and endRow of original vector
// party.
func (vp *cVectorParty) SliceIndex(lowerBoundRow, upperBoundRow int) (startIndex, endIndex int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// subtract endIndex by 1 when endIndex points to vp.length+1

// Write writes a vector party to underlying writer. It first writes header and then writes vectors
// based on vector party mode. **This vector party should be from archive batch and already pruned.**
func (vp *cVectorParty) Write(writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// Write 6 bytes padding.

// Starting writing vectors.
// Stop writing since there are no vectors in this vp.

// Write value vector.
// Here we directly move data from c allocated memory into writer.

// Stop writing since there are no more vectors in this vp.

// Write null vector.
// Here we directly move data from c allocated memory into writer.

// Stop writing since there are no more vectors in this vp.

// Write count vector.
// Here we directly move data from c allocated memory into writer.

// Read reads a vector party from underlying reader. It first reads header from the reader and does
// several sanity checks. Then it reads vectors based on vector party mode.
func (vp *cVectorParty) Read(reader io.Reader, s common.VectorPartySerializer) error {
	_ = "STUB: not implemented"
	return nil
}

// Read unused bytes

// Stop reading since there are no vectors in this vp.

// Read value vector.

// Here we directly read from reader into the c allocated bytes.

// Stop reading since there are no more vectors in this vp.

// Read null vector.

// Here we directly read from reader into the c allocated bytes.

// Stop reading since there are no more vectors in this vp.

// Read count vector.

// Here we directly read from reader into the c allocated bytes.

// GetHostVectorPartySlice implements GetHostVectorPartySlice in cVectorParty
func (vp *cVectorParty) GetHostVectorPartySlice(startIndex, length int) common.HostVectorPartySlice {
	_ = "STUB: not implemented"
	return *new(common.HostVectorPartySlice)
}

// Allocates implements Allocate in cVectorParty
func (vp *cVectorParty) Allocate(hasCount bool) { _ = "STUB: not implemented"; return }

// Dump is for testing purpose
func (vp *cVectorParty) Dump(file *os.File) { _ = "STUB: not implemented"; return }
