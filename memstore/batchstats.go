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
	"github.com/uber/aresdb/cluster/topology"
)

// BatchStatsReporter is used to report batch level stats like row count
type BatchStatsReporter struct {
	intervalInSeconds int
	memStore          MemStore
	shardOwner        topology.ShardOwner
	stopChan          chan struct{}
}

// NewBatchStatsReporter create a new BatchStatsReporter instance
func NewBatchStatsReporter(intervalInSeconds int, memStore MemStore, shardOwner topology.ShardOwner) *BatchStatsReporter {
	_ = "STUB: not implemented"
	return nil
}

// Run is a ticker function to run report periodically
func (batchStats *BatchStatsReporter) Run() { _ = "STUB: not implemented"; return }

// Stop to stop the stats reporter
func (batchStats *BatchStatsReporter) Stop() { _ = "STUB: not implemented"; return }

func (batchStats *BatchStatsReporter) reportBatchStats() { _ = "STUB: not implemented"; return }

func (batchStats *BatchStatsReporter) reportBatchStat(batchIDs map[int]string) {
	_ = "STUB: not implemented"
	return
}
