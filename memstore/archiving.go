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
	memCom "github.com/uber/aresdb/memstore/common"
)

// liveStoreSnapshot stores a snapshot of the LiveStore structure
// for fast archiving read without mutex locking and map accessing.
// The structure is snapshotted so that new batch/vector party creation by
// ingestion will not affect archiving. The underlying data is still shared with
// ingestion in parallel, with no read/write conflict, assuming that:
//   - archiving will not read beyond lastReadRecord for newly appended records
//   - archiving will only read records older than the cutoff, while ingestion
//     will not update any record older than that (backfill will be delayed until
//     ongoing archiving completes).
type liveStoreSnapshot struct {
	// Stores the structure as [RandomBatchIndex][ColumnID].
	batches [][]memCom.VectorParty
	// For purging live batch later.
	batchIDs              []int32
	numRecordsInLastBatch int
}

// snapshot creates a snapshot of the LiveStore structure for archiving and backfill fast read.
func (s *LiveStore) snapshot() (ss liveStoreSnapshot) {
	_ = "STUB: not implemented"
	return *new(liveStoreSnapshot)
}

// Live batches are purged by archiving so all batches returned here
// should be valid.

// archivingPatch stores records to be patched onto a archive batch.
// The records are identified by recordIDs and stored in the snapshot.
// The records will be sorted according to the sortColumns.
type archivingPatch struct {
	// RecordID.BatchID here refers to the RandomBatchIndex in the snapshot.
	recordIDs   []memCom.RecordID
	sortColumns []int
	// Readonly. We won't change it during sorting archiving patch and merging
	// with archive batch.
	data liveStoreSnapshot
}

// createArchivingPatches creates an archiving patch per affected UTC day.
// The key of the returned map is the number of days since Unix Epoch.
func (ss liveStoreSnapshot) createArchivingPatches(
	cutoff uint32, oldCutoff uint32, sortColumns []int,
	reporter ArchiveJobDetailReporter, jobKey string, tableName string, shardID int,
) map[int32]*archivingPatch {
	_ = "STUB: not implemented"
	return nil
}

// Add the record for archiving

// getBatchIDsToPurge returns list of batchIDs to purge in live store if its
// max event time is less than cutoff
// We do not purge the last batch if it's partially archived (ss.numRecordsInLastBatch
// != lastBatch.Size).
func (s *LiveStore) getBatchIDsToPurge(cutoff uint32) []int32 {
	_ = "STUB: not implemented"
	return nil
}

func (ap archivingPatch) Len() int { _ = "STUB: not implemented"; return 0 }

func (ap archivingPatch) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (ap archivingPatch) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Tie, move on to next sort column.

// GetDataValue reads value from underlying columns after sorted.
func (ap *archivingPatch) GetDataValue(row, columnID int) memCom.DataValue {
	_ = "STUB: not implemented"
	return *new(memCom.DataValue)
}

// GetDataValue reads value from underlying columns after sorted. If it's missing, it will return
// passed value instead.
func (ap *archivingPatch) GetDataValueWithDefault(row, columnID int, defaultValue memCom.DataValue) memCom.DataValue {
	_ = "STUB: not implemented"
	return *new(memCom.DataValue)
}

// GetCount get number of elements for values in the specified row/column, it is only valid for Array Value
func (ap *archivingPatch) GetCount(row, columnID int) int { _ = "STUB: not implemented"; return 0 }

// Archive is the process of periodically moving stable records in fact tables from live batches to archive batches,
// and converting them to a compressed format (run-length encoding). This is a blocking call so caller need to wait
// for archiving process to finish.
func (m *memStoreImpl) Archive(table string, shardID int, cutoff uint32, reporter ArchiveJobDetailReporter) error {
	_ = "STUB: not implemented"
	return nil
}

// Emit duration metrics and report back to scheduler.

// table already deleted. stop here

// Update the archiving cutoff time high water mark so ingestion won't update records below
// the new target archiving cutoff time.

// Create a new archive store version and switch to it.

// Wait for queries in other goroutines to prevent archiving from prematurely purging the old version.

// Purge redo log files on disk.

// delete obsolete batch versions if there are no peer bootstraping job running

// Delete obsolete (merged base batch) vector parties in oldArchivedStore. We don't need any lock since all queries
// should already finish processing the old version.

// We don't need to check existence again since it should be already created if missing in merge stage.

// Purge archive batch on disk.

// Purge archive batch in memory.

// Report memory usage.

// The new merged batch is no longer unmanaged memory.

// Purge live store in memory.

func (shard *TableShard) createNewArchiveStoreVersion(cutoff uint32, reporter ArchiveJobDetailReporter, jobKey string) (
	patchByDay map[int32]*archivingPatch, oldVersion *ArchiveStoreVersion, unmanagedMemoryBytes int64, err error) {
	_ = "STUB: not implemented"

	// Block column deletion
	return nil, nil, 0, nil
}

// Snapshot schema

// Snapshot unsorted store structure, but not the data.

// Scan unsorted snapshot for stable records.

// Begin of merge.

// We need to load all columns into memory for archiving.

// Unpin columns requested in this batch to unblock eviction.

// Copy unmerged base batch into new version.
