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
	"sync"

	memCom "github.com/uber/aresdb/memstore/common"
	metaCom "github.com/uber/aresdb/metastore/common"
)

// BackfillManager manages the records that need to be put into a backfill queue and merged with
// sorted batches directly.
type BackfillManager struct {
	sync.RWMutex `json:"-"`
	BackfillConfig

	// Name of the table.
	TableName string `json:"-"`

	// The shard id of the table.
	Shard int `json:"-"`

	// queue to hold UpsertBatches to backfill
	UpsertBatches []*memCom.UpsertBatch `json:"-"`

	// keep track of the number of records in backfill queue
	NumRecords int `json:"numRecords"`

	// keep track of the size of the buffer that holds batches to be backfilled
	CurrentBufferSize int64 `json:"currentBufferSize"`

	// keep track of the size of the buffer that holds batches being backfilled
	BackfillingBufferSize int64 `json:"backfillingBufferSize"`

	// keep track of the redo log file of the last batch backfilled
	LastRedoFile int64 `json:"lastRedoFile"`

	// keep track of the offset of the last batch backfilled
	LastBatchOffset uint32 `json:"lastBatchOffset"`

	// keep track of the redo log file of the last batch queued
	CurrentRedoFile int64 `json:"currentRedoFile"`

	// keep track of the offset of the last batch being queued
	CurrentBatchOffset uint32 `json:"currentBatchOffset"`

	AppendCond *sync.Cond `json:"-"`
}

// BackfillConfig defines configs for backfill
type BackfillConfig struct {
	// max buffer size to hold backfill data
	MaxBufferSize int64 `json:"maxBufferSize"`

	// threshold to trigger backfill
	BackfillThresholdInBytes int64 `json:"backfillThresholdInBytes"`
}

// NewBackfillManager creates a new BackfillManager instance.
func NewBackfillManager(tableName string, shard int, config BackfillConfig) *BackfillManager {
	_ = "STUB: not implemented"
	return nil
}

// WaitForBackfillBufferAvailability blocks until backfill buffer is available
func (r *BackfillManager) WaitForBackfillBufferAvailability() { _ = "STUB: not implemented"; return }

// Append appends an upsert batch into the backfill queue.
// Returns true if buffer limit has been reached and caller may need to wait
func (r *BackfillManager) Append(upsertBatch *memCom.UpsertBatch, redoFile int64, batchOffset uint32) bool {
	_ = "STUB: not implemented"
	return false
}

// advance position even if data is not for backfill

// ReadUpsertBatch reads upsert batch in backfill queue, user should not lock schema
func (r *BackfillManager) ReadUpsertBatch(index, start, length int, schema *memCom.TableSchema) (data [][]interface{}, columnNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// StartBackfill gets a slice of UpsertBatches from backfill queue and returns the
// CurrentRedoFile and CurrentBatchOffset.
func (r *BackfillManager) StartBackfill() ([]*memCom.UpsertBatch, int64, uint32) {
	_ = "STUB: not implemented"
	return nil, 0, 0
}

// no data to backfill
// but CurrentRedoFile/CurrentBatchOffset may not be checkpointed yet(live batch)

// QualifyToTriggerBackfill decides if OK to trigger size-based backfill process
func (r *BackfillManager) QualifyToTriggerBackfill() bool { _ = "STUB: not implemented"; return false }

// advanceOffset cleans up space and wakes up enqueue processes
func (r *BackfillManager) advanceOffset(redoFile int64, offset uint32) {
	_ = "STUB: not implemented"
	return
}

// GetLatestRedoFileAndOffset returns latest redofile and its batch offset
func (r *BackfillManager) GetLatestRedoFileAndOffset() (int64, uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

// MarshalJSON marshals a BackfillManager into json.
func (r *BackfillManager) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Avoid json.Marshal loop calls.
	return nil, nil
}

// Destruct set the golang object references used by backfill manager to be nil to trigger gc ealier.
func (r *BackfillManager) Destruct() { _ = "STUB: not implemented"; return }

// Done updates the backfill progress both in memory and in metastore.
func (r *BackfillManager) Done(currentRedoFile int64, currentBatchOffset uint32,
	metaStore metaCom.MetaStore) error {
	_ = "STUB: not implemented"
	return nil
}
