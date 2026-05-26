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
	"math"
	"sync"

	"github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/redolog"
)

// BaseBatchID is the starting id of all batches.
const BaseBatchID = int32(math.MinInt32)

// LiveBatch represents a live batch.
type LiveBatch struct {
	// The common data structure holding column data.
	common.Batch

	// Capacity of the batch which is decided at the creation time.
	Capacity int

	// For convenience to access fields of live store.
	// Schema locks should be acquired after data locks.
	liveStore *LiveStore

	// maximum of arrival time
	MaxArrivalTime uint32
}

// LiveStore stores live batches of columnar data.
type LiveStore struct {
	sync.RWMutex

	// Following fields are protected by above mutex.

	// The batch id to batch map.
	Batches map[int32]*LiveBatch

	// Number of rows to create for new batches.
	BatchSize int

	// The upper bound of records (exclusive) that can be read by queries.
	LastReadRecord common.RecordID

	// This is the in memory archiving cutoff time high watermark that gets set by the archiving job
	// before each archiving run. Ingestion will not insert/update records that are older than
	// the archiving cutoff watermark.
	ArchivingCutoffHighWatermark uint32

	// Logs.
	RedoLogManager redolog.RedologManager

	// Manage backfill queue during ingestion.
	BackfillManager *BackfillManager

	// Manage snapshot related stats.
	SnapshotManager *SnapshotManager

	// For convenience. Schema locks should be acquired after data locks.
	tableSchema *common.TableSchema

	// The writer lock is to guarantee single writer to a Shard at all time. To ensure this, writers
	// (ingestion, archiving etc) need to hold this lock at all times. This lock
	// should be acquired before the VectorStore and Batch locks.
	// TODO: if spinning lock performance is a concern we may need to upgrade this
	// to a designated goroutine with a pc-queue channel.
	WriterLock sync.RWMutex

	// Following fields are protected by WriterLock.

	// Primary key table of the Shard.
	PrimaryKey common.PrimaryKey

	// The position of the next record to be used for writing. Only used by the ingester.
	NextWriteRecord common.RecordID

	// For convenience.
	HostMemoryManager common.HostMemoryManager `json:"-"`

	// Last modified time per column in live store and fact table only. Used to measure the data freshness for each column.
	// Protected by the writer lock of live store. If a column is never ingested, thhe last modified time will be zero.
	// Metrics will be emitted after each ingestion request.
	lastModifiedTimePerColumn []uint32
}

// NewLiveStore creates a new live batch.
func NewLiveStore(batchSize int, totalShards int, shard *TableShard) *LiveStore {
	_ = "STUB: not implemented"
	return nil
}

// for now dimension table is unsharded

// initial primary key buckets should consider number of shards

// MaxBufferSize should consider total number of shards

// reportBatch memory usage of backfill max buffer size.

// GetBatchIDs snapshots the batches and returns a list of batch ids for read
// with the number of records in batchIDs[len()-1].
func (s *LiveStore) GetBatchIDs() (batchIDs []int32, numRecordsInLastBatch int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// GetBatchForRead returns and read locks the batch with its ID for reads. Caller must explicitly
// RUnlock() the returned batch after all reads.
func (s *LiveStore) GetBatchForRead(id int32) *LiveBatch { _ = "STUB: not implemented"; return nil }

// GetBatchForWrite returns and locks the batch with its ID for reads. Caller must explicitly
// Unlock() the returned batch after all reads.
func (s *LiveStore) GetBatchForWrite(id int32) *LiveBatch { _ = "STUB: not implemented"; return nil }

// GetOrCreateBatch retrieve LiveBatch for specified batchID, append one if not exist
// this is only used during read snapshots, otherwise the batchID maybe cause conflict
// user should unlock the batch after use
func (s *LiveStore) getOrCreateBatch(batchID int32) *LiveBatch {
	_ = "STUB: not implemented"
	return nil
}

// AdvanceNextWriteRecord reserves space for a record that return the next available record position
// back to the caller.
func (s *LiveStore) AdvanceNextWriteRecord() common.RecordID {
	_ = "STUB: not implemented"
	return *new(common.RecordID)
}

// We only create batch when the current NextWriteRecord is pointing to nil.

// AdvanceLastReadRecord advances the high watermark of the rows to the next write record.
func (s *LiveStore) AdvanceLastReadRecord() { _ = "STUB: not implemented"; return }

// appendBatch appends a new batch. The batch is returned with its ID.
func (s *LiveStore) appendBatch(batchID int32) *LiveBatch { _ = "STUB: not implemented"; return nil }

// PurgeBatch purges the specified batch.
func (s *LiveStore) PurgeBatch(id int32) {
	_ = "STUB: not implemented"

	// Detach first.
	return
}

// Wait for readers to finish.

// SafeDestruct.

// Destruct deletes all vectors allocated in C.
// Caller must detach the Shard first and wait until all users are finished.
func (s *LiveStore) Destruct() { _ = "STUB: not implemented"; return }

// reportBatch free memory used by backfill manager.

// Delay 1 second to report memory change to wait for gc happens.

// PurgeBatches purges the specified batches.
func (s *LiveStore) PurgeBatches(ids []int32) { _ = "STUB: not implemented"; return }

// MarshalJSON marshals a LiveStore into json.
func (s *LiveStore) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Avoid json.Marshal loop calls.
	return nil, nil
}

// Following fields are protected by reader lock of liveStore.

// Following fields are protected by writer lock of liveStore.

// LookupKey looks up the given key in primary key.
func (s *LiveStore) LookupKey(keyStrs []string) (common.RecordID, bool) {
	_ = "STUB: not implemented"
	return *new(common.RecordID), false
}

// GetMemoryUsageForColumn get the live store memory usage for given data type
func (s *LiveStore) GetMemoryUsageForColumn(valueType common.DataType, columnID int) int {
	_ = "STUB: not implemented"
	return 0
}

// GetOrCreateVectorParty returns LiveVectorParty for the specified column from
// the live batch. locked specifies whether the batch has been locked.
// The lock will be left in the same state after the function returns.
func (b *LiveBatch) GetOrCreateVectorParty(columnID int, locked bool) common.LiveVectorParty {
	_ = "STUB: not implemented"
	// Ensure that columnID is not out of bound.
	return *new(common.LiveVectorParty)
}

// Ensure that the VectorParty is allocated with values and nulls.

// MarshalJSON marshals a LiveBatch into json.
func (b *LiveBatch) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
