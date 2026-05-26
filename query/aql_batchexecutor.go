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
	"time"
	"unsafe"
)

// BatchExecutor is batch executor interface for both Non-aggregation query and Aggregation query
type BatchExecutor interface {
	// filter operation
	filter()
	// join operation
	join()
	// project of measure/select columns
	project()
	// reduce to sort and aggregate result
	reduce()
	// prepare work before execution
	preExec(lastBatch bool, start time.Time)
	// post execution after execution
	postExec(start time.Time)
}

// DummyBatchExecutorImpl is a dummy executor which do nothing
type DummyBatchExecutorImpl struct {
}

// NewDummyBatchExecutor create a dummy BatchExecutor
func NewDummyBatchExecutor() BatchExecutor { _ = "STUB: not implemented"; return *new(BatchExecutor) }

func (e *DummyBatchExecutorImpl) filter() { _ = "STUB: not implemented"; return }

func (e *DummyBatchExecutorImpl) join() { _ = "STUB: not implemented"; return }

func (e *DummyBatchExecutorImpl) project() { _ = "STUB: not implemented"; return }

func (e *DummyBatchExecutorImpl) reduce() { _ = "STUB: not implemented"; return }

func (e *DummyBatchExecutorImpl) preExec(lastBatch bool, start time.Time) {
	_ = "STUB: not implemented"
	return
}

func (e *DummyBatchExecutorImpl) postExec(start time.Time) {
	_ = "STUB: not implemented"

	// BatchExecutorImpl is batch executor implementation for original aggregation query
	return
}

type BatchExecutorImpl struct {
	qc                  *AQLQueryContext
	batchID             int32
	isLastBatch         bool
	customFilterFunc    customFilterExecutor
	stream              unsafe.Pointer
	start               time.Time
	sizeBeforeGeoFilter int
}

// NewBatchExecutor is to create a BatchExecutor.
func NewBatchExecutor(qc *AQLQueryContext, batchID int32, customFilterFunc customFilterExecutor, stream unsafe.Pointer, start time.Time) BatchExecutor {
	_ = "STUB: not implemented"
	return *new(BatchExecutor)
}

// filter
func (e *BatchExecutorImpl) filter() {
	_ = "STUB: not implemented"
	// process main table common filter
	return
}

// join
func (e *BatchExecutorImpl) join() { _ = "STUB: not implemented"; return }

// join foreign tables

// prepare foreign table recordIDs
// Note:
// RecordID {
//   int32_t batchID
// 	 uint32_t index
// }
// takes up 8 bytes

// perform hash lookup

// process filters that involves foreign table columns if any

// allocate two predicate vector for geo intersect

// evalMeasures is to fill measure values
func (e *BatchExecutorImpl) evalMeasures() {
	_ = "STUB: not implemented"
	// measure evaluation.
	return
}

// evalDimensions is to fill dimension values
func (e *BatchExecutorImpl) evalDimensions(prevResultSize int) {
	_ = "STUB: not implemented"
	// dimension expression evaluation.
	return
}

// project is to generate dimension and measure values
func (e *BatchExecutorImpl) project() {
	_ = "STUB: not implemented"
	// Prepare for dimension and measure evaluation.
	return
}

// wait for stream to clean up non used buffer before final aggregation

// reduce is to aggregate measures based on dimensions and aggregation function
func (e *BatchExecutorImpl) reduce() {
	_ = "STUB: not implemented"
	// init dimIndexVectorD for sorting and reducing
	return
}

// sort by key.

// reduce by key.

func (e *BatchExecutorImpl) preExec(isLastBatch bool, start time.Time) {
	_ = "STUB: not implemented"
	return
}

// initialize index vector.

func (e *BatchExecutorImpl) postExec(start time.Time) {
	_ = "STUB: not implemented"
	// swap result buffer before next batch
	return
}

// Only profile one batch.
