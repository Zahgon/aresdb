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
	"unsafe"

	"github.com/uber/aresdb/cgoutils"

	"github.com/uber/aresdb/memstore"
	memCom "github.com/uber/aresdb/memstore/common"
	queryCom "github.com/uber/aresdb/query/common"
	"github.com/uber/aresdb/query/expr"
)

const (
	hllQueryRequiredMemoryInMB = 10 * 1024
)

// batchTransferExecutor defines the type of the functor to transfer a live batch or a archive batch
// from host memory to device memory. hostVPs will be the columns to be released after transfer. startRow
// is used to slice the vector party.
type batchTransferExecutor func(stream unsafe.Pointer) (deviceColumns []deviceVectorPartySlice,
	hostVPs []memCom.VectorParty, firstColumn, startRow, totalBytes, numTransfers, sizeAfterPrefilter int)

// customFilterExecutor is the functor to apply custom filters depends on the batch type. For archive batch,
// the custom filter will be the time filter and will only be applied to first or last batch. For live batch,
// the custom filters will be the cutoff time filter if cutoff is larger than 0, pre-filters and time filters.
type customFilterExecutor func(stream unsafe.Pointer)

// ProcessQuery processes the compiled query and executes it on GPU.
func (qc *AQLQueryContext) ProcessQuery(memStore memstore.MemStore) {
	_ = "STUB: not implemented"
	return
}

// find out exactly what the error was and set err

// prepare geo intersection

// if no shape exist and geo check for point in shape
// no need to continue processing batch

// if no shape exist and geo check for point not in shape
// no need to do geo intersection

// query execution for last batch.

// this code snippet does the followings:
// 1. write stats to log.
// 2. allocate host buffer for result and copy the result from device to host.
// 3. clean up device status buffers if no panic.

// Copy the result to host memory.

// copy dimensions

// copy measures

func (qc *AQLQueryContext) processShard(memStore memstore.MemStore, shardID int, previousBatchExecutor BatchExecutor) BatchExecutor {
	_ = "STUB: not implemented"
	return *new(BatchExecutor)
}

// Process live batches.

// For now, dimension table does not persist min and max therefore
// we can only skip live batch for fact table.
// TODO: Persist min/max/numTrues when snapshotting.

// Process archive batches.

// Release releases all device memory it allocated. It **should only called** when any errors happens while the query is
// processed.
func (qc *AQLQueryContext) Release() {
	_ = "STUB: not implemented"
	// release device memory for processing current batch.
	return
}

// CleanUpDevice cleans up the device status including
//  1. clean up the device buffer for storing results.
//  2. clean up the cuda streams
func (qc *AQLQueryContext) cleanUpDeviceStatus() {
	_ = "STUB: not implemented"
	// clean up foreign table memory after query
	return
}

// release geo pointers

// Destroy streams

// Clean up the device result buffers.

// Clean up timezone lookup buffer.

// clean up foreign table
func (qc *AQLQueryContext) cleanUpForeignTable(table *foreignTable) {
	_ = "STUB: not implemented"
	return
}

// getGeoShapeLatLongSlice format GeoShapeGo into slices of float32 for query purpose
// Lats and Longs are stored in the format as [a1,a2,...an,a1,MaxFloat32,b1,bz,...bn]
// refer to time_series_aggregate.h for GeoShape struct
func getGeoShapeLatLongSlice(shapesLats, shapesLongs []float32, gs memCom.GeoShapeGo) ([]float32, []float32, int) {
	_ = "STUB: not implemented"
	return nil, nil, 0
}

// write place holder at start of polygon

// FLT_MAX as placeholder for each polygon

func (qc *AQLQueryContext) prepareForGeoIntersect(memStore memstore.MemStore) (shapeExists bool) {
	_ = "STUB: not implemented"
	return false
}

// geo table is not sharded

// compiler should have verified the geo column GeoShape type

// allocate memory for lats, longs (float32) and numPoints (int32) device vectors

// prepare foreign table (allocate and transfer memory) before processing
func (qc *AQLQueryContext) prepareForeignTable(memStore memstore.MemStore, joinTableID int, join queryCom.Join) {
	_ = "STUB: not implemented"
	return
}

// join only support dimension table for now
// and dimension table is not shared

// only need live store for dimension table

// transfer primary key

// allocate device memory

// prepareTimezoneTable
func (qc *AQLQueryContext) prepareTimezoneTable(store memstore.MemStore) {
	_ = "STUB: not implemented"
	return
}

// Timezone table

// transferLiveBatch returns a functor to transfer a live batch to device memory. The size parameter will be either the
// size of the batch or num records in last batch. hostColumns will always be empty since we should not release a vector
// party of a live batch. Start row will always be zero as well.
func (qc *AQLQueryContext) transferLiveBatch(batch *memstore.LiveBatch, size int) batchTransferExecutor {
	_ = "STUB: not implemented"
	return *new(batchTransferExecutor)
}

// Allocate column inputs.

// liveBatchTimeFilterExecutor returns a functor to apply custom time filters to live batch.
func (qc *AQLQueryContext) liveBatchCustomFilterExecutor(cutoff uint32) customFilterExecutor {
	_ = "STUB: not implemented"
	return *new(customFilterExecutor)
}

// cutoff filter evaluation.
// only apply to fact table where cutoff > 0

// time filter evaluation

// prefilter evaluation

// transferArchiveBatch returns the functor to transfer an archive batch to device memory. We will need to release
// hostColumns after transfer completes.
func (qc *AQLQueryContext) transferArchiveBatch(batch *memstore.ArchiveBatch,
	isFirstOrLast bool) batchTransferExecutor {
	_ = "STUB: not implemented"
	return *new(batchTransferExecutor)
}

// Request columns, prefilter-slicing, allocate column inputs.

// Must iterate in reverse order to apply prefilter slicing properly.

// Request/pin column from disk and wait.

// prefilter slicing

// archiveBatchCustomFilterExecutor returns a functor to apply custom filter to first or last archive batch.
func (qc *AQLQueryContext) archiveBatchCustomFilterExecutor(isFirstOrLast bool) customFilterExecutor {
	_ = "STUB: not implemented"
	return *new(customFilterExecutor)
}

// helper function for copy dimension vector. Returns the total size of dimension vector.
func asyncCopyDimensionVector(toDimVector, fromDimVector unsafe.Pointer, length, offset int, numDimsPerDimWidth queryCom.DimCountsPerDimWidth,
	toVectorCapacity, fromVectorCapacity int, copyFunc cgoutils.AsyncMemCopyFunc,
	stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	return
}

// copy null bytes

// cleanupDeviceResultBuffers cleans up result buffers and resets result fields.
func (bc *oopkBatchContext) cleanupDeviceResultBuffers() { _ = "STUB: not implemented"; return }

// ok to free nil vectors even if it's not allocated.

// clean up memory not used in final aggregation (sort, reduce, hll)
// before aggregation happen
func (bc *oopkBatchContext) cleanupBeforeAggregation() { _ = "STUB: not implemented"; return }

// swapResultBufferForNextBatch swaps the two
// sets of dim/measure/hash vectors to get ready for the next batch.
func (bc *oopkBatchContext) swapResultBufferForNextBatch() { _ = "STUB: not implemented"; return }

// prepareForFiltering prepares the input and the index vectors for filtering.
func (bc *oopkBatchContext) prepareForFiltering(
	columns []deviceVectorPartySlice, firstColumn int, startRow int, stream unsafe.Pointer) {
	_ = "STUB: not implemented"
	return
}

// Allocate twice of the size to save number of allocations of temporary index vector.

// prepareForDimAndMeasureEval ensures that dim/measure vectors have enough
// capacity for bc.resultSize+bc.size.
func (bc *oopkBatchContext) prepareForDimAndMeasureEval(
	dimRowBytes int, measureBytes int, numDimsPerDimWidth queryCom.DimCountsPerDimWidth, isHLL bool, useHashReduction bool, stream unsafe.Pointer) {
	_ = "STUB: not implemented"
	return
}

// Extra budget for future proofing.

// uint32_t for index value

// uint64_t for hash value
// Note: only when aggregate function is hll, we need to reuse vector[0]

// reallocateResultBuffers reallocates the result buffer pair to size
// resultCapacity*unitBytes and copies resultSize*unitBytes from input[0] to output[0].
// this function will read and modify the device pointers in buffers
func (bc *oopkBatchContext) reallocateResultBuffers(
	buffers *[2]devicePointer, unitBytes int, stream unsafe.Pointer, copyFunc func(to, from unsafe.Pointer)) {
	_ = "STUB: not implemented"

	// copy previous pointers first
	return
}

// set buffers to null device pointer

// make sure input pointers are cleaned up

// reallocate new device buffers

// doProfile checks the corresponding profileName against query parameter
// and do cuda profiling for this action if name matches.
func (qc *AQLQueryContext) doProfile(action func(), profileName string, stream unsafe.Pointer) {
	_ = "STUB: not implemented"
	return
}

// explicit waiting for cuda stream to avoid profiling previous actions.

// explicit waiting for cuda stream to wait for completion of current action.

// processBatch allocates device memory and starts async input data
// transferring to device memory. It then invokes previousBatchExecutor
// asynchronously to process the previous batch. When both async operations
// finish, it prepares for the current batch execution and returns it as
// a function closure to be invoked later. customFilterExecutor is the executor
// to apply custom filters for live batch and archive batch.
func (qc *AQLQueryContext) processBatch(table string, shardID int,
	batch *memCom.Batch, batchID int32, batchSize int, transferFunc batchTransferExecutor,
	customFilterFunc customFilterExecutor, previousBatchExecutor BatchExecutor, needToUnlockBatch bool) BatchExecutor {
	_ = "STUB: not implemented"
	return *new(BatchExecutor)
}

// Finish executing previous batch first to avoid timeline overlapping

// reset stats.

// Async transfer.

// Async execute the previous batch.

// find out exactly what the error was and set err

// Wait for data transfer of the current batch.

// only archive vector party will be returned after transfer function

// Wait for execution of the previous batch.

// column data transfer for current batch is done
// need release current batch's column data before panic

// if the query is already satisfied in the middle, we can skip next batch and return

// no prefilter slicing in livebatch, startRow is always 0

// prefilterSlice does the following:
// 1. binary search for prefilter values following the matched sort column order
// 2. record matched index range on these matched sort columns
// 3. binary search on unmatched compressed columns for the row number range
// 4. index slice on uncompressed columns for the row number range
// 5. align/pad all slices to be pushed
func (qc *AQLQueryContext) prefilterSlice(vp memCom.ArchiveVectorParty, prefilterIndex, startRow, endRow int) (int, int, memCom.HostVectorPartySlice) {
	_ = "STUB: not implemented"
	return 0, 0, *new(memCom.HostVectorPartySlice)
}

// matched equality filter

// matched range filter
// lower bound

// treat as unmatchedColumn when there is one range filter missing

// SliceByValue of upperBound

// treat as unmatchedColumn when there is one range filter missing

// unmatched columns, simply slice based on row number range

// calculateMemoryRequirement estimate memory requirement for batch data.
func (qc *AQLQueryContext) calculateMemoryRequirement(memStore memstore.MemStore) int {
	_ = "STUB: not implemented"
	// keep track of max requirement for batch
	return 0
}

//TODO(jians): hard code hll query memory requirement here for now,
//we can track memory usage
//based on table, dimensions, duration to do estimation

// estimate live batch memory usage

// find first non null batch and estimate.

// estimate archive batch memory usage

// estimateLiveBatchMemoryUsage estimate the GPU memory usage for live batches
func (qc *AQLQueryContext) estimateLiveBatchMemoryUsage(batch *memstore.LiveBatch) int {
	_ = "STUB: not implemented"
	return 0
}

// for array live vp, we need to use offset size + pool size, and remove cap size for device

// estimateArchiveBatchMemoryUsage estimate the GPU memory usage for archive batch
func (qc *AQLQueryContext) estimateArchiveBatchMemoryUsage(batch *memstore.ArchiveBatch, isFirstOrLast bool) int {
	_ = "STUB: not implemented"
	return 0
}

// max number of rows after pre-filtering. used for non-agg query

// TODO(cdavid): only read metadata when estimate query memory requirement.

// estimateMemUsageForBatch calculates memory usage including:
// * Index vector
// * Predicate vector
// * Dimension
// * Measurement
// * Sort (hash/index)
// * Reduce
func (qc *AQLQueryContext) estimateMemUsageForBatch(firstColumnSize, columnMemUsage, maxSizeAfterPreFilter int) (memUsage int) {
	_ = "STUB: not implemented"
	// 1. columnMemUsage
	return 0
}

// 2. index vector memory usage (4 bytes each)

// 3. predicate memory usage (1 byte each)

// 4. record id vector for foreign table (8 bytes each recordID)

// 5. expression eval memory (max scratch space)

// 6. geoPredicateVector

// 7. max(memUsageBeforeAgg, sortReduceMemoryUsage)

// 8. Dimension vector memory usage (input + output)

// For hash reduction, need hash table with int64_t key (8 bytes) and measureBytes value.
// The capacity is 2 * size.

// For sort based reduction, need to allocate space for hash vectors (8bytes each) and
// dim index vectors (4 bytes each)

// 9. Measure vector memory usage (input + output)

// memory usage duration expression (filter, dimension, measure) evaluation
func (qc *AQLQueryContext) estimateExpressionEvaluationMemUsage(inputSize int) (memUsage int) {
	_ = "STUB: not implemented"
	// filter expression evaluation
	return 0
}

// dimension expression evaluation

// measure expression evaluation

// Note: we only calculate Sort memory usage
// since sort memory usage is larger than reduce
// and we only care about the maximum
func estimateSortReduceMemUsage(inputSize int) (memUsage int) {
	_ = "STUB: not implemented"
	// dimension index vector
	// 4 byte for uint32
	// 2 vectors for input and output
	return 0
}

// hash vector
// 8 byte for uint64 hash value
// 2 vectors for input and output

// we sort dim index values as value, and hash value as key

// estimateScratchSpaceMemUsage calculates memory usage for an expression
func estimateScratchSpaceMemUsage(exp expr.Expr, firstColumnSize int, isRoot bool) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// calculateForeignTableMemUsage returns how much device memory is needed for foreign table
func (qc *AQLQueryContext) calculateForeignTableMemUsage(memStore memstore.MemStore) int {
	_ = "STUB: not implemented"
	return 0
}

// join only support dimension table for now
// and dimension table is not shared

// only need live store for dimension table

// primary key

// VPs

// FindDeviceForQuery calls device manager to find a device for the query
func (qc *AQLQueryContext) FindDeviceForQuery(memStore memstore.MemStore, preferredDevice int,
	deviceManager *DeviceManager, timeout int) {
	_ = "STUB: not implemented"
	return
}

func (qc *AQLQueryContext) runBatchExecutor(e BatchExecutor, isLastBatch bool) {
	_ = "STUB: not implemented"
	return
}

// copyHostToDevice copy vector party slice to device vector party slice
func copyHostToDevice(vps memCom.HostVectorPartySlice, deviceVPSlice deviceVectorPartySlice, stream unsafe.Pointer, device int) (bytesCopied, numTransfers int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// for array data type
// copy offset-length

// copy value

func hostToDeviceColumn(hostColumn memCom.HostVectorPartySlice, device int) deviceVectorPartySlice {
	_ = "STUB: not implemented"
	return *new(deviceVectorPartySlice)
}

// fnon-array type

// shouldSkipLiveBatch will determine whether we can skip processing a live batch by checking time filter and
// eligible main table common filters. The batch must be non nil.
func (qc *AQLQueryContext) shouldSkipLiveBatch(b *memstore.LiveBatch) bool {
	_ = "STUB: not implemented"
	return false
}

// shouldSkipLiveBatchWithFilter will check max and min for the corresponding column against the filter express and
// determines whether we should skip processing this live batch.
// Following constraints apply:
//  1. Filter must be on main table.
//  2. Filter must be a binary expression.
//  3. OPs must be one of (EQ, GTE,GE,LTE,LE).
//  4. One side of the expr must be VarRef
//  5. Another side of the xpr must be NumericalLiteral
//  6. ColumnType must be UInt32
func shouldSkipLiveBatchWithFilter(b *memstore.LiveBatch, filter expr.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

// First try lhs VarRef, rhs Num.

// Then try rhs VarRef, lhs Num.

// Swap column to the left and number to right.

// Invert the OP.

// Time filters and main table filters are guaranteed to be on main table.

func (qc *AQLQueryContext) initializeNonAggResponse() { _ = "STUB: not implemented"; return }

// non eager flush
