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

// stageName represents each query stage.
type stageName string

const (
	prepareForeignTableTiming     stageName = "prepareForeignTable"
	transferTiming                          = "transfer"
	prepareForFilteringTiming               = "prepareForFiltering"
	initIndexVectorTiming                   = "initIndexVector"
	filterEvalTiming                        = "filterEval"
	prepareForeignRecordIDsTiming           = "prepareForeignRecordIDs"
	foreignTableFilterEvalTiming            = "foreignTableFilterEval"
	geoIntersectEvalTiming                  = "geoIntersectEval"
	prepareForDimAndMeasureTiming           = "prepareForDimAndMeasure"
	dimEvalTiming                           = "dimEval"
	measureEvalTiming                       = "measureEval"
	hllEvalTiming                           = "hllEval"
	sortEvalTiming                          = "sortEval"
	reduceEvalTiming                        = "reduceEval"
	hashReduceEvalTiming                    = "hashReduceEval"
	expandEvalTiming                        = "expandEval"
	cleanupTiming                           = "cleanUpEval"
	resultTransferTiming                    = "resultTransfer"
	resultFlushTiming                       = "resultFlush"
	finalCleanupTiming                      = "finalCleanUp"
)

// oopkBatchStats stores stats for a single batch execution.
type oopkBatchStats struct {
	// Store timings for each stage of a single batch.
	timings map[stageName]float64
	// totalTiming for this batch.
	totalTiming      float64
	batchID          int32
	batchSize        int
	bytesTransferred int
	numTransferCalls int
}

// oopkStageSummaryStats stores running info for each stage.
type oopkStageSummaryStats struct {
	name       stageName
	max        float64
	min        float64
	avg        float64
	count      int
	total      float64
	percentage float64
}

// MarshalJSON marshals the message to JSON in a custom way.
func (s *oopkStageSummaryStats) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type stageSummaryStatsSlice []*oopkStageSummaryStats

// Len implements sort.Sort interface for stageSummaryStatsSlice.
func (s stageSummaryStatsSlice) Len() int {
	_ = "STUB: not implemented"

	// Swap implements sort.Sort interface for stageSummaryStatsSlice.
	return 0
}

func (s stageSummaryStatsSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Less implements sort.Sort interface for stageSummaryStatsSlice.
func (s stageSummaryStatsSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// oopkQueryStats stores the overall stats for a query.
type oopkQueryStats struct {
	// stats for each stage. Sorted by total time.
	stageStats []*oopkStageSummaryStats
	// mapping from stage name to stage stats.
	Name2Stage map[stageName]*oopkStageSummaryStats `json:"stages"`

	// Total timing for all query stages **including transfer**.
	TotalTiming float64 `json:"latency"`

	// Total number of batches.
	NumBatches int `json:"batches"`
	// Total number of records processed on GPU.
	// A record could represent multiple data record if firstColumn is compressed.
	NumRecords int `json:"records"`

	// For archive batch, we skip process empty batch. For live batch, we will skip it
	// if its min or max value does not pass main table filters or time filters.
	NumBatchSkipped int `json:"numBatchSkipped"`

	// Stats for input data transferred via PCIe.
	BytesTransferred int `json:"tranBytes"`
	NumTransferCalls int `json:"tranCalls"`
}

// NumRows implements the utils.TableDataSource for stats.
func (stats oopkQueryStats) NumRows() int { _ = "STUB: not implemented"; return 0 }

// GetValue implements the utils.TableDataSource for stats. **Notes** row boundary
// are not checked!
func (stats oopkQueryStats) GetValue(row, col int) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ColumnHeaders implements the utils.TableDataSource for stats.
func (stats oopkQueryStats) ColumnHeaders() []string { _ = "STUB: not implemented"; return nil }

// reportTimingForCurrentBatch will first wait for current cuda stream if the debug mode is set and change the timing stat accordingly.
// It will add to the total timing as well. Therefore this function should only be called one time for each stage.
func (qc *AQLQueryContext) reportTimingForCurrentBatch(stream unsafe.Pointer, start *time.Time, name stageName) {
	_ = "STUB: not implemented"
	return
}

// reportTiming is similar to reportTimingForCurrentBatch except that it modifies the query stats for the
// whole query. It's usually should be called once for each stage
func (qc *AQLQueryContext) reportTiming(stream unsafe.Pointer, start *time.Time, name stageName) {
	_ = "STUB: not implemented"
	return
}

// applyStageStats applies the stage stats to the overall query stats and compute max,minCallName and total for that
// stage.
func (stats *oopkQueryStats) applyStageStats(name stageName, value float64) {
	_ = "STUB: not implemented"
	return
}

// applyBatchStats applies the current batch stats onto the overall query stats. It computes information
// like max, minCallName, average for each stage as well as the percentage.
func (stats *oopkQueryStats) applyBatchStats(batchStats oopkBatchStats) {
	_ = "STUB: not implemented"
	return
}

// writeToLog writes the summary stats for this query in a tabular format to logger.
func (stats *oopkQueryStats) writeToLog() { _ = "STUB: not implemented"; return }

// Compute average and percentage.

// Create tabular output.

// reportBatch will report OOPK batch related stats to the query logger.
func (qc *AQLQueryContext) reportBatch(isArchiveBatch bool) { _ = "STUB: not implemented"; return }
