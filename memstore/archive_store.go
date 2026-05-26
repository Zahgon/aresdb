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

	"github.com/uber/aresdb/memstore/common"
)

// ArchiveBatch represents a archive batch.
type ArchiveBatch struct {
	common.Batch

	// Size of the batch (number of rows). Notice that compression changes the
	// length of some columns, does not change the size of the batch.
	Size int

	// Version for archive batches.
	Version uint32

	// SeqNum denotes backfill sequence number
	SeqNum uint32

	// For convenience.
	BatchID int32
	Shard   *TableShard
}

// ArchiveStoreVersion stores a version of archive batches of columnar data.
type ArchiveStoreVersion struct {
	// The mutex
	// protects the Batches map structure and the archiving cutoff field.
	// It does not protect contents within a batch. Before releasing
	// the VectorStore mutex, user should lock the batch level mutex if necessary
	// to ensure proper protection at batch level.
	sync.RWMutex `json:"-"`

	// Wait group used to prevent this ArchiveStore from being evicted
	Users sync.WaitGroup `json:"-"`

	// Each batch in the slice is identified by BaseBatchID+index.
	// Index out of bound and nil Batch for archive batches indicates that none of
	// the columns have been loaded into memory from disk.
	Batches map[int32]*ArchiveBatch `json:"batches"`

	// The archiving cutoff used for this version of the sorted store.
	ArchivingCutoff uint32 `json:"archivingCutoff"`

	// For convenience.
	shard *TableShard
}

// ArchiveStore manages archive stores versions.
// Archive store version evolves to a new version after archiving.
// Readers should follow the following locking protocol:
//
//	archiveStore.Users.Add(1)
//	// tableShard.ArchiveStore can no longer be accessed directly.
//	// continue reading from archiveStore
//	archiveStore.Users.Done()
type ArchiveStore struct {
	// The mutex protects the pointer pointing to the current version of archived vector version.
	sync.RWMutex

	PurgeManager *PurgeManager
	// Current version points to the most recent version of vector store version for queries to use.
	CurrentVersion *ArchiveStoreVersion
}

// NewArchiveStore creates a new archive store. Current version is just a place holder for test.
// It will be replaced during recovery.
func NewArchiveStore(shard *TableShard) *ArchiveStore { _ = "STUB: not implemented"; return nil }

// NewArchiveStoreVersion creates a new empty archive store version given cutoff.
func NewArchiveStoreVersion(cutoff uint32, shard *TableShard) *ArchiveStoreVersion {
	_ = "STUB: not implemented"
	return nil
}

// Destruct deletes all vectors allocated in C.
// Caller must detach the Shard first and wait until all users are finished.
func (s *ArchiveStore) Destruct() { _ = "STUB: not implemented"; return }

// RequestBatch returns the requested archive batch from the archive store version.
func (v *ArchiveStoreVersion) RequestBatch(batchID int32) *ArchiveBatch {
	_ = "STUB: not implemented"
	return nil
}

// Read version and size from MetaStore.

// WriteToDisk writes each column of a batch to disk. It happens on archiving
// stage for merged archive batch so there is no need to lock it.
func (b *ArchiveBatch) WriteToDisk() error { _ = "STUB: not implemented"; return nil }

// GetCurrentVersion returns current SortedVectorStoreVersion and does proper locking. It'v used by
// query and data browsing. Users need to call version.Users.Done() after their work.
func (s *ArchiveStore) GetCurrentVersion() *ArchiveStoreVersion {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON marshals a ArchiveStore into json.
func (s *ArchiveStore) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalJSON marshals a ArchiveStoreVersion into json.
func (v *ArchiveStoreVersion) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Avoid json.Marshal loop calls.
	return nil, nil
}

// RequestVectorParty creates(optional), pins, and returns the requested vector party.
// On creation it also asynchronously loads a vector party from disk into memory.
//
// Caller must call vp.WaitForDiskLoad() before using it,
// and call vp.Release() afterwards.
func (b *ArchiveBatch) RequestVectorParty(columnID int) common.ArchiveVectorParty {
	_ = "STUB: not implemented"
	return *new(common.ArchiveVectorParty)
}

// This is a newly added column. We need to allocate more space to put this column.

// archive batch always have archive batch

// Release lock on batch and wait for vector party to be loaded

// Set a placeholder to prevent repetitive loading,
// columnID should always be smaller than len(ValueTypeByColumn).

// TryEvict attempts to evict and destruct the specified column from the archive
// batch. It will fail fast if the column is currently in use so that host
// memory manager can try evicting other VPs immediately.
// Returns vector party evicted if succeeded.
func (b *ArchiveBatch) TryEvict(columnID int) common.ArchiveVectorParty {
	_ = "STUB: not implemented"
	return *new(common.ArchiveVectorParty)
}

// BlockingDelete blocks until all users are finished with the specified column,
// and then deletes the column from the batch.
// Returns the vector party deleted if any.
func (b *ArchiveBatch) BlockingDelete(columnID int) common.ArchiveVectorParty {
	_ = "STUB: not implemented"
	return *new(common.ArchiveVectorParty)
}

// evict attempts to evict and destruct the specified column from the archive
// batch. It will block if the blocking is set to true, other wise it will fail
// fast.
func (b *ArchiveBatch) evict(columnID int, blocking bool) common.ArchiveVectorParty {
	_ = "STUB: not implemented"
	return *new(common.ArchiveVectorParty)
}

// GetBatchForRead returns a archiveBatch for read,
// reader needs to unlock after use
func (v *ArchiveStoreVersion) GetBatchForRead(batchID int) *ArchiveBatch {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON marshals a ArchiveBatch into json.
func (b *ArchiveBatch) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// BuildIndex builds an index over the primary key columns of this archive batch and inserts the records id into the
// given primary key.
func (b *ArchiveBatch) BuildIndex(sortColumns []int, primaryKeyColumns []int, pk common.PrimaryKey) error {
	_ = "STUB: not implemented"
	return nil
}

// we need to use sortedColumnIterator to advance sort column.

// create sort column iterators if the primary key column is also a sort column.

// Prepare primary key values.

// Primary key column will not have any default values.

// Get primary key for each record.
// truncate key

// Found duplicate record in backfill is a data correctness issue,
// in which new update will only go to one of the records depending on the sort order.
// we decide for now this rare case will proceed but trigger alert
// so that user can adjust schema and backfill data when needed,
// instead of crash the server completely.

// Clone returns a copy of current batch including all references to underlying vector parties. Caller is responsible
// for holding the lock (if necessary).
func (b *ArchiveBatch) Clone() *ArchiveBatch { _ = "STUB: not implemented"; return nil }

// UnpinVectorParties unpins all vector parties in the slice.
func UnpinVectorParties(requestedVPs []common.ArchiveVectorParty) {
	_ = "STUB: not implemented"
	return
}
