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
	"io"
	"os"
	"sync"
	"unsafe"

	"github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/memstore/vectors"
)

// LiveVectorParty is the representation of list data type vector party in live store.
// It supports random access read and write. However, it does not support serialization into disk.
// It underlying uses a high level memory pool to store the list data. Therefore when this vector
// party is destructed, the underlying memory pool needs to be destroyed as well.
type LiveVectorParty struct {
	baseVectorParty
	// storing the offset to slab footer offset for each row.
	caps       *vectors.Vector
	memoryPool HighLevelMemoryPool
	sync.RWMutex
}

// GetBytes returns the bytes this vp occupies except memory pool
func (vp *LiveVectorParty) GetBytes() int64 { _ = "STUB: not implemented"; return 0 }

// GetTotalBytes return the bytes this vp occupies including memory pool
func (vp *LiveVectorParty) GetTotalBytes() int64 { _ = "STUB: not implemented"; return 0 }

// SafeDestruct destructs vector party memory.
func (vp *LiveVectorParty) SafeDestruct() { _ = "STUB: not implemented"; return }

// Write serialize vector party.
func (vp *LiveVectorParty) Write(writer io.Writer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// length

// data type

// nonDefaultValue count, 0 for List VectorParty

// columnMode, AllValuesPresent for now

// Write 6 bytes padding.

// write offsets

// to compatible with archive vp, align to 64 bytes alignment

// value bytes, to compatible with archive vp, align to 64 bytes alignment

// write values

// Read deserialize vector party
func (vp *LiveVectorParty) Read(reader io.Reader, serializer common.VectorPartySerializer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// non default value count

// column mode

// Read unused bytes

// Read value bytes

// update offset.

// Set footer offset.

// GetCap returns the cap at ith row. Only used for free a list element in live store.
func (vp *LiveVectorParty) GetCap(row int) uint32 { _ = "STUB: not implemented"; return 0 }

// SetBool is not supported by list vector party.
func (vp *LiveVectorParty) SetBool(offset int, val bool, valid bool) {
	_ = "STUB: not implemented"
	return
}

// SetValue is the implementation of common.LiveVectorParty
func (vp *LiveVectorParty) SetValue(row int, val unsafe.Pointer, valid bool) {
	_ = "STUB: not implemented"
	return
}

// Set footer offset.

// AsList is the implementation from common.VectorParty
func (vp *LiveVectorParty) AsList() common.ListVectorParty {
	_ = "STUB: not implemented"

	// Equals is the implementation from common.VectorParty
	return *new(common.ListVectorParty)
}

func (vp *LiveVectorParty) Equals(other common.VectorParty) bool {
	_ = "STUB: not implemented"
	return false
}

// SetListValue is the implentation of common.ListVecotrParty
func (vp *LiveVectorParty) GetListValue(row int) (unsafe.Pointer, bool) {
	_ = "STUB: not implemented"
	return *

	// SetListValue is the implentation of common.ListVecotrParty
	new(unsafe.Pointer), false
}

func (vp *LiveVectorParty) SetListValue(row int, val unsafe.Pointer, valid bool) {
	_ = "STUB: not implemented"
	return
}

// SetGoValue is not supported by list vector party.
func (vp *LiveVectorParty) SetGoValue(offset int, val common.GoDataValue, valid bool) {
	_ = "STUB: not implemented"
	return
}

// GetValue is the implementation from common.VectorParty
func (vp *LiveVectorParty) GetValue(row int) (val unsafe.Pointer, validity bool) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false
}

// GetMinMaxValue is not supported by list vector party.
func (vp *LiveVectorParty) GetMinMaxValue() (min, max uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

// GetDataValue is not implemented in baseVectorParty
func (vp *LiveVectorParty) GetDataValue(row int) common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

// SetDataValue
func (vp *LiveVectorParty) SetDataValue(row int, value common.DataValue,
	countsUpdateMode common.ValueCountsUpdateMode, counts ...uint32) {
	_ = "STUB: not implemented"
	return
}

// GetDataValueByRow just call GetDataValue
func (vp *LiveVectorParty) GetDataValueByRow(row int) common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

// Allocate allocate underlying storage for vector party
func (vp *LiveVectorParty) Allocate(hasCount bool) { _ = "STUB: not implemented"; return }

// Dump is for testing purpose
func (vp *LiveVectorParty) Dump(file *os.File) { _ = "STUB: not implemented"; return }

// GetHostVectorPartySlice implements GetHostVectorPartySlice in TransferableVectorParty
func (vp *LiveVectorParty) GetHostVectorPartySlice(startIndex, length int) common.HostVectorPartySlice {
	_ = "STUB: not implemented"
	// LiveVectorParty will always use startIndex = 0, length is whole VP, so startIndex is ignored here
	return *new(common.HostVectorPartySlice)
}

// SetLength is only for testing purpose, do NOT use this function in real code
func (vp *LiveVectorParty) SetLength(length int) {
	_ = "STUB: not implemented"

	// NewLiveVectorParty returns a LiveVectorParty pointer which implements ListVectorParty.
	// It's safe to pass nil HostMemoryManager.
	return
}

func NewLiveVectorParty(length int, dataType common.DataType,
	hmm common.HostMemoryManager) common.LiveVectorParty {
	_ = "STUB: not implemented"
	return *new(common.LiveVectorParty)
}
