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

// #include "time_series_aggregate.h"
import "C"

import (
	"unsafe"

	"github.com/uber/aresdb/query/expr"
)

var bytesComma = []byte(",")

// Postprocess converts the internal dimension and measure vector in binary
// format to AQLQueryResult nested result format. It also translates enum
// values back to their string representations.
func (qc *AQLQueryContext) Postprocess() { _ = "STUB: not implemented"; return }

// skip translate enum for HLL if query is from broker

// should never be here except bug

func (qc *AQLQueryContext) initResultFlushContext() { _ = "STUB: not implemented"; return }

// flushResultBuffer reads dimension and measure data from current OOPK buffer to Results
func (qc *AQLQueryContext) flushResultBuffer() { _ = "STUB: not implemented"; return }

// don't translate enum if it's for distributed mode (DataOnly == true)

// For avg aggregation function, we only need to read first 4 bytes which is the average.

// PostprocessAsHLLData serializes the query result into HLLData format. It will also release the device memory after
// serialization.
func (qc *AQLQueryContext) PostprocessAsHLLData() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getEnumReverseDict returns the enum reverse dict of a ast node if it's a VarRef node, otherwise it will return
// a nil slice.
func (qc *AQLQueryContext) getEnumReverseDict(dimIndex int, expression expr.Expr) []string {
	_ = "STUB: not implemented"
	return nil
}

// special handling element_at for array enum type

// return validShapeUUIDs as the reverse enum dict if dimIndex match geo dimension

// ReleaseHostResultsBuffers deletes the result buffer from host memory after postprocessing
func (qc *AQLQueryContext) ReleaseHostResultsBuffers() { _ = "STUB: not implemented"; return }

// hllVectorD and hllDimRegIDCountD used for hll query only

// set geoIntersection to nil

func (qc *AQLQueryContext) ResultsRowsFlushed() int { _ = "STUB: not implemented"; return 0 }

func readMeasure(measureRow unsafe.Pointer, ast expr.Expr, measureBytes int) *float64 {
	_ = "STUB: not implemented"
	// TODO: consider converting non-zero identity values to nil.
	return nil
}

// Should never happen

// Should never happen.

// should never happen
