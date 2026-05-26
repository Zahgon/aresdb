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

	"time"

	"github.com/uber/aresdb/memstore/common"
)

// SnapshotManager manages the snapshot related stats and progress.
type SnapshotManager struct {
	sync.RWMutex `json:"-"`

	// Job Trigger Condition related fields

	// Number of mutations since last snapshot. Measured as number of rows mutated.
	NumMutations int `json:"numMutations"`

	// Last snapshot time.
	LastSnapshotTime time.Time `json:"'lastSnapshotTime'"`

	// Snapshot progress related fields.

	// keep track of the redo log file of the last batch snapshotted.
	LastRedoFile int64 `json:"lastRedoFile"`

	// keep track of the offset of the last batch snapshotted
	LastBatchOffset uint32 `json:"lastBatchOffset"`

	// keep track of the record position of the last batch snapshotted
	LastRecord common.RecordID

	// keep track of the redo log file of the last batch queued
	CurrentRedoFile int64 `json:"currentRedoFile"`

	// keep track of the offset of the last batch queued
	CurrentBatchOffset uint32 `json:"currentBatchOffset"`

	// keep track of the record position when last batch queued
	CurrentRecord common.RecordID

	// Configs
	SnapshotInterval time.Duration `json:"snapshotInterval"`

	SnapshotThreshold int `json:"snapshotThreshold"`

	// for convenience.
	shard *TableShard
}

// NewSnapshotManager creates a new SnapshotManager instance.
func NewSnapshotManager(shard *TableShard) *SnapshotManager { _ = "STUB: not implemented"; return nil }

// StartSnapshot returns current redo log file ,offset
func (s *SnapshotManager) StartSnapshot() (int64, uint32, int, common.RecordID) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(common.RecordID)
}

// ApplyUpsertBatch advances CurrentRedoLogFile and CurrentBatchOffset and increments NumMutations after applying
// an upsert batch to live store.
func (s *SnapshotManager) ApplyUpsertBatch(redoFile int64, offset uint32, numMutations int, currentRecord common.RecordID) {
	_ = "STUB: not implemented"
	return
}

// QualifyForSnapshot tells whether we can trigger a snapshot job.
func (s *SnapshotManager) QualifyForSnapshot() bool { _ = "STUB: not implemented"; return false }

// updateSnapshotProgress updates snapshot progress in memory. It also subtracts NumMutations by lastNumMutations to reflect correct
// number of mutations since last snapshot.
func (s *SnapshotManager) updateSnapshotProgress(redoFile int64, offset uint32, lastNumMutations int, record common.RecordID) {
	_ = "STUB: not implemented"
	return
}

// Done updates the snapshot progress both in memory and in metastore and updates number of mutations
// accordingly.
func (s *SnapshotManager) Done(currentRedoFile int64, currentBatchOffset uint32, lastNumMutations int, currentRecord common.RecordID) error {
	_ = "STUB: not implemented"
	return nil

	// we should always record last snapshot time
}

// GetLastSnapshotInfo get last snapshot redolog file, offset and timestamp, lastRecord
func (s *SnapshotManager) GetLastSnapshotInfo() (int64, uint32, time.Time, common.RecordID) {
	_ = "STUB: not implemented"
	return 0, 0, *new(time.Time), *new(common.RecordID)
}

// SetLastSnapshotInfo update last snapshot redolog file, offset and timestamp, lastRecord
func (s *SnapshotManager) SetLastSnapshotInfo(redoLogFile int64, offset uint32, record common.RecordID) {
	_ = "STUB: not implemented"
	return
}

// MarshalJSON marshals a BackfillManager into json.
func (s *SnapshotManager) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
