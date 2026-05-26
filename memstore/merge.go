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
	"github.com/uber/aresdb/memstore/common"
)

// mergeContext carries all context information used during merge
type mergeContext struct {
	base   *ArchiveBatch
	patch  *archivingPatch
	merged *ArchiveBatch

	// Number of total columns (including deleted columns).
	numColumns int
	// Number of records in base patch. If base is nil, size is 0.
	baseSize int
	// Number of records in base patch and patch.
	totalSize int

	// Iterators for sorted base. One iterator per sort column.
	baseIters []sortedColumnIterator
	// Iterators for archiving patch. One iterator per sort column.
	patchIters []sortedColumnIterator

	// Following fields will only be used during preallocate stage
	// length is len(sortColumns).
	mergedLengths []int

	// Following fields will be only used during merge stage stores the indexes to
	// write to the final merged archive batch for each column length is len(all columns).
	outputBegins []int
	// It stores the accumulated counts to write to the final merged archive batch for
	// each sorted column, length is len(all columns).
	outputCounts []uint32
	// Buffer the unsorted column for writing final uncompressed values.
	unsortedColumns []int

	// Stores list of rows that have been marked as deleted
	baseRowDeleted []int
	// This is used to short circuit merge for deleted columns.
	columnDeletions []bool

	// Needed during merging.
	dataTypes []common.DataType

	defaultValues []*common.DataValue

	// Keep track of total unmanaged memory space this merge process uses.
	unmanagedMemoryBytes int64
}

// newMergeContext creates a new context for merge existing batch and archive batch
// into a new batch. It's shared by pre-allocate stage and actual merge stage.
func newMergeContext(base *ArchiveBatch, patch *archivingPatch, columnDeletions []bool, dataTypes []common.DataType,
	defaultValues []*common.DataValue, baseRowDeleted []int) *mergeContext {
	_ = "STUB: not implemented"
	return nil
}

// initIters initialize patch and base iterators.
func (ctx *mergeContext) initIters() { _ = "STUB: not implemented"; return }

// archive batch iterator

// sortedColumnIterator is the common interface to merge two sorted columns
type sortedColumnIterator interface {
	// Read values from current idx.
	read()
	// Advance the iterator.
	next()
	// Tells whether the iteration has been finished.
	done() bool
	// Returns current data value the iterator points to
	value() common.DataValue
	// Tells what's the current index of iterator in a vector
	index() int
	// Tells iterator where to start and stop in next sorted column.
	// For base iterator, it's current count.
	// For patch iterator, it's current index.
	// To make the interface consistent, we choose return uint32.
	// So for patch iterator, we need to convert it between int and uint32.
	currentPosition() uint32

	// For base iterator, it's next count.
	// For patch iterator, it's next index.
	nextPosition() uint32
	// Count of current value.
	count() uint32
	// For base iterator, it should be next end count. For patch iterator,
	// it should be next end index. This must be called before iterating
	// each slice
	setEndPosition(pos uint32)

	// list of rows that should be skipped for current value
	currentSkipRows() []int
}

type archiveBatchColumnIterator struct {
	// Column being iterated.
	vp common.ArchiveVectorParty

	// Iterator position.
	idx int

	// #
	currentCount uint32

	// rows from beginning to current value
	nextCount uint32

	// total rows of this vp
	endCount uint32
	val      common.DataValue

	// rows deleted
	rowsDeleted []int

	// rowsDeleted start position for current value
	currentRowsDeletedStart int

	// rowsDeleted end position(exclusive) for current value
	currentRowsDeletedEnd int
}

// newArchiveBatchColumnIterator creates a new iterator that iterates through
// a specific range of archive batch column.
func newArchiveBatchColumnIterator(base *ArchiveBatch, columnID int, rowsDeleted []int) sortedColumnIterator {
	_ = "STUB: not implemented"
	return *new(sortedColumnIterator)
}

func (itr *archiveBatchColumnIterator) done() bool { _ = "STUB: not implemented"; return false }

func (itr *archiveBatchColumnIterator) read() { _ = "STUB: not implemented"; return }

// Sort columns of archive batch should be either mode 3 or mode 0.

func (itr *archiveBatchColumnIterator) next() {
	_ = "STUB: not implemented"
	// move to next value
	return
}

func (itr *archiveBatchColumnIterator) index() int { _ = "STUB: not implemented"; return 0 }

func (itr *archiveBatchColumnIterator) value() common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

func (itr *archiveBatchColumnIterator) currentPosition() uint32 {
	_ = "STUB: not implemented"
	return 0
}

func (itr *archiveBatchColumnIterator) nextPosition() uint32 { _ = "STUB: not implemented"; return 0 }

func (itr *archiveBatchColumnIterator) count() uint32 { _ = "STUB: not implemented"; return 0 }

func (itr *archiveBatchColumnIterator) currentSkipRows() []int {
	_ = "STUB: not implemented"
	return nil
}

func (itr *archiveBatchColumnIterator) setEndPosition(pos uint32) {
	_ = "STUB: not implemented"

	// see to the first valid value
	return
}

type archivingPatchColumnIterator struct {
	patch    *archivingPatch
	columnID int

	endIdx int
	// Iterator position.
	idx     int
	nextIdx int
	val     common.DataValue
}

// newArchivingPatchColumnIterator creates a new iterator that iterates through a specific
// range of archive batch column.
func newArchivingPatchColumnIterator(patch *archivingPatch, columnID int) sortedColumnIterator {
	_ = "STUB: not implemented"
	return *new(sortedColumnIterator)
}

func (itr *archivingPatchColumnIterator) done() bool { _ = "STUB: not implemented"; return false }

func (itr *archivingPatchColumnIterator) read() { _ = "STUB: not implemented"; return }

// Read current value.

// Find next value != current value.

func (itr *archivingPatchColumnIterator) next() {
	_ = "STUB: not implemented"
	// Jump directly to index of next different value.
	return
}

func (itr *archivingPatchColumnIterator) value() common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

func (itr *archivingPatchColumnIterator) index() int { _ = "STUB: not implemented"; return 0 }

func (itr *archivingPatchColumnIterator) currentPosition() uint32 {
	_ = "STUB: not implemented"
	return 0
}

func (itr *archivingPatchColumnIterator) nextPosition() uint32 { _ = "STUB: not implemented"; return 0 }

func (itr *archivingPatchColumnIterator) count() uint32 { _ = "STUB: not implemented"; return 0 }

func (itr *archivingPatchColumnIterator) currentSkipRows() []int {
	_ = "STUB: not implemented"
	return nil
}

func (itr *archivingPatchColumnIterator) setEndPosition(pos uint32) {
	_ = "STUB: not implemented"
	return
}

// merge an live batch with a archive batch and store the merged data into a new archive batch.
// It has two stages:
//  1. preallocate: attempt to merge two batches but only calculate how much space merged data needs.
//  2. merge: based on the calculated size, allocate space and do actual merge.
//
// This algorithm will do merge on sorted columns first and based on the positions of last sorted column,
// it will copy non-sorted columns data into final result.
// The parameters will be used as cutoff and seqNum for the merged batch.
func (ctx *mergeContext) merge(cutoff uint32, seqNum uint32) {
	_ = "STUB: not implemented"
	// We preallocate space in 1st pass to avoid allocate unnecessary memory.
	return
}

// Allocate space for merged archive batch.

// Reset iterators to begin 2nd pass.

// Do actual merge and write to output vector party.

// If sort columns is empty, we need to dump all values in batch first and then dump patch values.

// Write base.

// Write patch.

// Scan through all columns for mode 0 and 1 columns and remove unnecessary vectors.

// allocate space for merged archive batch based on calculated mergedLengths.
func (ctx *mergeContext) allocate(cutoff uint32, seqNum uint32) { _ = "STUB: not implemented"; return }

// Need to create batch in advance otherwise vector party's allUsersDone will have nil value.

// Sort columns.

// Non-sort columns.

// array will never appear in sort columns

// calculate value vector bytes needed for Array ArchiveParty
func (ctx *mergeContext) calculateArrayVectorPartyBytes(columnID int, dataType common.DataType) int64 {
	_ = "STUB: not implemented"
	return 0

	// bytes for patch
}

// bytes for original archive patch

// skip the deleted row

// common function signature for both preallocate and merge.
type mergeAction func(baseIter, patchIter sortedColumnIterator, sortColIdx, compareRes int, mergedVP common.ArchiveVectorParty)

// preAllocate called during first pass.
func (ctx *mergeContext) preAllocate(baseIter, patchIter sortedColumnIterator, sortColIdx, compareRes int, mergedVP common.ArchiveVectorParty) {
	_ = "STUB: not implemented"
	return
}

func (ctx *mergeContext) writeUnsortedColumns(start, end int, reader common.BatchReader, skipRows []int) {
	_ = "STUB: not implemented"
	// Base batch is possible to be nil for a particular day.
	return
}

// We will skip writing to deleted columns so that it will have all null values.

func (ctx *mergeContext) writeOutput(baseIter, patchIter sortedColumnIterator, sortColIdx, compareRes int, mergedVP common.ArchiveVectorParty) {
	_ = "STUB: not implemented"
	return
}

// if patch value is less, we read from patch

// Set value on the mergedVP.

// mergeRecursive does merge on base and patch iterators on a given sort column. baseEndPos is the end count
// for base iter to stop. patchEndPos is the end index for patch iter to stop. Any end pos with 0 value means
// for this iterator it's an empty range.
func (ctx *mergeContext) mergeRecursive(sortColIdx int, baseEndPos uint32, patchEndPos int, f mergeAction) {
	_ = "STUB: not implemented"
	return
}

// Only used by merge stage during preallocate stage ctx.merged will be nil

// ignore values if the count is 0.

// New value from patch.

// This is equal to compareRes < 0.

// This is equal to compareRes > 0.
