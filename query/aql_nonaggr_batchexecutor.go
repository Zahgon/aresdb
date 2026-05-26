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

	queryCom "github.com/uber/aresdb/query/common"
)

// NonAggrBatchExecutorImpl is batch executor implementation for non-aggregation query
type NonAggrBatchExecutorImpl struct {
	*BatchExecutorImpl
}

// project for non-aggregation query will only calculate the selected columns
// dimension calculation, reduce will be skipped, once the generated result reaches limit, it will return and cancel all other ongoing processing.
func (e *NonAggrBatchExecutorImpl) project() {
	_ = "STUB: not implemented"
	// Prepare for dimension evaluation.
	return
}

// for non-aggregation query, we always write from start for dimension output

// uncompress the result from baseCount

// wait for stream to clean up non used buffer before final aggregation

func (e *NonAggrBatchExecutorImpl) prepareForDimEval(
	dimRowBytes int, numDimsPerDimWidth queryCom.DimCountsPerDimWidth, stream unsafe.Pointer) {
	_ = "STUB: not implemented"
	return
}

// only allocate dimension vector once

// Extra budget for future proofing.

func (e *NonAggrBatchExecutorImpl) expandDimensions(numDims queryCom.DimCountsPerDimWidth) {
	_ = "STUB: not implemented"
	return
}

func (e *NonAggrBatchExecutorImpl) postExec(start time.Time) {
	_ = "STUB: not implemented"
	// TODO: @shz experiment with on demand flush when next batch can not fit in buffer
	return
}

// transfer current batch result from device to host

// flush current batches results to result buffer

// Only profile one batch.

func (e *NonAggrBatchExecutorImpl) reduce() {
	_ = "STUB: not implemented"
	// nothing need to do for non-aggregation query
	return
}

// getNumberOfRecordsNeeded is a helper function
func (e *NonAggrBatchExecutorImpl) getNumberOfRecordsNeeded() (needed int) {
	_ = "STUB: not implemented"
	return 0
}
