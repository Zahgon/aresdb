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

	"github.com/uber/aresdb/diskstore"
	"github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/memstore/vectors"
)

const (
	ListVectorPartyHeader uint32 = 0xFADEFACF
)

// ArchiveVectorParty is the representation of list data type vector party in archive store.
// It does not support random access update. Instead updates to archiveListVectorParty can only be done
// via appending to the tail during archiving and backfill.
// It use a single value vector to store the values and validities so that it has the same
// in memory representation except archiving vp does not have cap vector.
// The only update supported is same length value in place change
type ArchiveVectorParty struct {
	baseVectorParty
	common.Pinnable
	values *vectors.Vector
	// bytesWritten should only be used when archiving or backfilling. It's used to record the current position
	// in values vector.
	bytesWritten int64
	// lengthFilled is to record the number of records appended, which can not exceed length
	lengthFilled    int
	totalValueBytes int64
}

// NewArchiveVectorParty returns a new ArchiveVectorParty.
// It should only be used during backfill or archiving when constructing a new list
// archiving vp.
// Length is the number of total rows and totalValueBytes is the total bytes used to
// store values and validities.
func NewArchiveVectorParty(length int, dataType common.DataType,
	totalValueBytes int64, locker sync.Locker) common.ArchiveVectorParty {
	_ = "STUB: not implemented"
	return *new(common.ArchiveVectorParty)
}

func newArchiveVectorParty(length int, dataType common.DataType,
	totalValueBytes int64, locker sync.Locker) *ArchiveVectorParty {
	_ = "STUB: not implemented"
	return nil
}

// Allocate allocate underlying storage for vector party. Note allocation for
// archive vp does not report host memory change. Memory reporting is done
// after switching to the new version of archive store. Before switching the memory
// managed by this vp is counted as unmanaged memory.
func (vp *ArchiveVectorParty) Allocate(hasCount bool) { _ = "STUB: not implemented"; return }

// GetBytes returns the bytes this vp occupies.
func (vp *ArchiveVectorParty) GetBytes() int64 { _ = "STUB: not implemented"; return 0 }

// SafeDestruct destructs vector party memory.
func (vp *ArchiveVectorParty) SafeDestruct() { _ = "STUB: not implemented"; return }

// AsList is the implementation from common.VectorParty
func (vp *ArchiveVectorParty) AsList() common.ListVectorParty {
	_ = "STUB: not implemented"

	// Equals is the implementation from common.VectorParty
	return *new(common.ListVectorParty)
}

func (vp *ArchiveVectorParty) Equals(other common.VectorParty) bool {
	_ = "STUB: not implemented"
	return false
}

// GetValue is the implementation from common.VectorParty
func (vp *ArchiveVectorParty) getValue(row int) (val unsafe.Pointer, validity bool) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false
}

// GetDataValue is not implemented in baseVectorParty
func (vp *ArchiveVectorParty) GetDataValue(row int) common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

// GetDataValueByRow just call GetDataValue
func (vp *ArchiveVectorParty) GetDataValueByRow(row int) common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

func (vp *ArchiveVectorParty) setValue(row int, val unsafe.Pointer, valid bool) {
	_ = "STUB: not implemented"
	return
}

// invalid in-place update, should never happen

// update offset/length

// invalid jump update, should never happen

// SetDataValue is the implentation of common.VecotrParty
func (vp *ArchiveVectorParty) SetDataValue(row int, value common.DataValue,
	countsUpdateMode common.ValueCountsUpdateMode, counts ...uint32) {
	_ = "STUB: not implemented"
	return
}

// SetListValue is the implentation of common.ListVecotrParty
func (vp *ArchiveVectorParty) GetListValue(row int) (unsafe.Pointer, bool) {
	_ = "STUB: not implemented"
	return *

	// SetListValue is the implentation of common.ListVecotrParty
	new(unsafe.Pointer), false
}

func (vp *ArchiveVectorParty) SetListValue(row int, val unsafe.Pointer, valid bool) {
	_ = "STUB: not implemented"
	return
}

// Write is the implentation of common.VecotrParty
func (vp *ArchiveVectorParty) Write(writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// length

// data type

// nonDefaultValue count, 0 for List VectorParty

// columnMode, AllValuesPresent for now

// Write 6 bytes padding.

// Write offset vector.
// Here we directly move data from c allocated memory into writer.

// value bytes

// Write value vector.

// Read reads a vector party from underlying reader. It first reads header from the reader and does
// several sanity checks. Then it reads vectors based on vector party mode.
func (vp *ArchiveVectorParty) Read(reader io.Reader, s common.VectorPartySerializer) error {
	_ = "STUB: not implemented"
	return nil
}

// non default value count

// column mode

// Read unused bytes

// Read value bytes

// Read value vector.

// Here we directly read from reader into the c allocated bytes.

// memory usage:
// 1. offset vector party: (4 bytes offset + 4 bytes length) * length
// 2. value vector is totalValueBytes of uint8

// GetCount returns cumulative count on specified offset.
func (vp *ArchiveVectorParty) GetCount(offset int) uint32 {
	_ = "STUB: not implemented"
	// Same as non mode 3 vector.
	return 0
}

// SetCount is not supported by list vector party.
func (vp *ArchiveVectorParty) SetCount(offset int, count uint32) { _ = "STUB: not implemented"; return }

// LoadFromDisk load archive vector party from disk caller should lock archive batch before using
func (vp *ArchiveVectorParty) LoadFromDisk(hostMemManager common.HostMemoryManager, diskStore diskstore.DiskStore,
	table string, shardID int, columnID, batchID int, batchVersion uint32, seqNum uint32) {
	_ = "STUB: not implemented"
	return
}

// Prune prunes vector party based on column mode to clean memory if possible
func (vp *ArchiveVectorParty) Prune() {
	_ = "STUB: not implemented"
	// Nothing to prune for list vp.
	return
}

// SliceByValue is not supported by list vector party.
func (vp *ArchiveVectorParty) SliceByValue(lowerBoundRow, upperBoundRow int, value unsafe.Pointer) (
	startRow int, endRow int, startIndex int, endIndex int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// Slice vector party to get [startIndex, endIndex) based on [lowerBoundRow, upperBoundRow)
func (vp *ArchiveVectorParty) SliceIndex(lowerBoundRow, upperBoundRow int) (
	startIndex, endIndex int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// GetHostVectorPartySlice implements GetHostVectorPartySlice in TransferableVectorParty
func (vp *ArchiveVectorParty) GetHostVectorPartySlice(startIndex, length int) common.HostVectorPartySlice {
	_ = "STUB: not implemented"
	return *new(common.HostVectorPartySlice)
}

// find first entry which has non-zero length array value, which will have valid offset
// if not found, then will start from baseAddr

// find first entry which has non-zero length array value, which will have valid offset
// if not found, then will be the end of value buffer

// Dump is for testing purpose
func (vp *ArchiveVectorParty) Dump(file *os.File) { _ = "STUB: not implemented"; return }

// CopyOnWrite clone vector party for updates, the update can only for in-place change with same length for update row
func (vp *ArchiveVectorParty) CopyOnWrite(batchSize int) common.ArchiveVectorParty {
	_ = "STUB: not implemented"
	// archive vector party should always have allUsersDone initialized correctly with batch rwlock
	return *new(common.ArchiveVectorParty)
}
