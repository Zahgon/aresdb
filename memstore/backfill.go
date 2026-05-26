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

// Backfill is the process of merging records with event time older than cutoff with
// archive batches.
func (m *memStoreImpl) Backfill(table string, shardID int, reporter BackfillJobDetailReporter) error {
	_ = "STUB: not implemented"
	return nil
}

// no data to backfill: checkpoint if applicable

// checkpoint backfill progress

// Wait for queries in other goroutines to prevent archiving from prematurely purging the old version.

// Archiving cutoff won't change during backfill, so it's safe to use current version's cutoff.

func (shard *TableShard) createNewArchiveStoreVersionForBackfill(
	backfillPatches map[int32]*backfillPatch, reporter BackfillJobDetailReporter, jobKey string) (err error) {
	_ = "STUB: not implemented"
	// Block column deletion
	return nil
}

// Snapshot schema

// Only those batches that are affected and changed need to be cleaned.

// Real backfill implementation.

// Batch is clean, we can copy the old batch to new version directly.

// clean pointer in cloned batch.

// Copy other batches in old version to new version.

// switch to new version

// Purge batches on disk.

// Purge columns in memory.

// Report memory usage.

// Do the nil check in case column is evicted.

// patch stores records to be patched onto a archive batch.
// The records are identified by upsert batch idx and row number within
// the upsert batch
type backfillPatch struct {
	recordIDs []memCom.RecordID
	// For convenience.
	backfillBatches []*memCom.UpsertBatch
}

// createBackfillPatches groups records in upsert batches by day and put them into backfillPatches.
// Records in each backfillPatch are identified by RecordID where BatchID is the upsert batch index
// index is the row within the upsert batch.
func createBackfillPatches(backfillBatches []*memCom.UpsertBatch, reporter BackfillJobDetailReporter, jobKey string) (map[int32]*backfillPatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// backfillContext carries all context information used during backfill for a single day.
type backfillContext struct {
	// temporary live store to hold data to be later on merged with archive batch.
	backfillStore *LiveStore
	base          *ArchiveBatch
	patch         *backfillPatch

	// new archive batch after backfill.
	new *ArchiveBatch

	// snapshot of table schema.
	columnDeletions   []bool
	sortColumns       []int
	primaryKeyColumns []int
	defaultValues     []*memCom.DataValue

	// keep track of which columns have been forked already.
	columnsForked []bool

	// keep track of which row in base batch has been deleted and added to backfill store.
	baseRowDeleted []int

	dataTypes []memCom.DataType

	// columns need to be purged under two cases:
	// 	1. old columns are forked for updating in place.
	//  2. old columns in base batch for merging.
	// if there are no columns to be purged, it means the base batch is clean and therefore we don't
	// need to advance the version for this batch.
	columnsToPurge []memCom.ArchiveVectorParty

	// keep track of how much unmanaged memory bytes this day uses.
	// this does not include the temp primary key.
	unmanagedMemoryBytes int64

	// If we invoked a merge process,we can early unpin it before writing to disk.
	okForEarlyUnpin bool
}

func newBackfillStore(tableSchema *memCom.TableSchema, hostMemoryManager memCom.HostMemoryManager, initBuckets int) *LiveStore {
	_ = "STUB: not implemented"
	return nil
}

func newBackfillContext(baseBatch *ArchiveBatch, patch *backfillPatch, tableSchema *memCom.TableSchema, columnDeletions []bool,
	sortColumns []int, primaryKeyColumns []int, dataTypes []memCom.DataType, defaultValues []*memCom.DataValue,
	hostMemoryManager memCom.HostMemoryManager) backfillContext {
	_ = "STUB: not implemented"
	return *new(backfillContext)
}

// allocate more space for insertion.

// column deletion will be blocked during backfill, so we are safe to get column deletions from schema without
// lock.

// we can simply copy all the columns of base batch without lock since all columns already have been requested
// and pinned.

// release releases the resource hold by the backfillContext.
func (ctx *backfillContext) release() {
	_ = "STUB: not implemented"
	// release both the batch and primary key resources.
	return
}

// createArchivingPatch create an archiving patch for a single day. This assume all records in the snapshot is within
// the same calendar day bucket.
func (ss liveStoreSnapshot) createArchivingPatch(sortColumns []int) *archivingPatch {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *backfillContext) backfill(reporter BackfillJobDetailReporter, jobKey string) error {
	_ = "STUB: not implemented"
	// build index on archive batch.
	return nil
}

// reuse the space for primary primaryKeyValues of each row.

// newRecords: records that does not exist in base
// inplaceUpdateRecords: records that modifies unsortedColumns and can be updated inplace
// deleteThenInsertRecords: records that modifies sortedColumns and needs to be deleted from base and inserted again into temp live store
// noEffectRecords: records that does not modify any column

// We will do backfill row by row in patch.

// record id to apply to temp live store.

// truncate key

// new row in live store!

// changedPatchRow converts patch values in upsert batch to a slice of data values. The length is the number
// of columns in base batch. If the corresponding column does not exist in the upsert batch or the column is
// deleted, it will be nil. This changed row is used in several places:
//  1. when this patch row is a new row or an update on existing row in temp live store, we apply the row to
// temp live store directly.
//  2. when this patch row contains updates on sort column of base batch, we will first get the whole column
// from base batch and apply changes from this patch changed row and then write into temp live store
//  3. when this patch row only contains updates on unsort column of base batch, we will apply the changed patch
// row to forked column.

// get the data value from upsertBatch

// record is already in base batch.

// first detect if there are any changes to sort columns or array columns.

// we should write to live store.

// sorted column or array column size changed

// update the primary key pointing to new record id.

// only unsorted columns are changed, or array column has value change while size not changed

// in case we fork the column but does not invoke the merge procedure (which also call column.Prune()).
// column.Prune is idempotent so it's safe to call multiple times.

// original baseRowDeleted is not sorted.

// We can early unpin it only if we invoked a merge process.

// merge merges the records in temp live store with the new batch.
func (ctx *backfillContext) merge(reporter BackfillJobDetailReporter, jobKey string) {
	_ = "STUB: not implemented"
	return
}

// getChangedPatchRow get the upsert batch row as a slice of pointer of data value format to be consistent with changed
// base row. Note an upsert batch row may not have values for all columns so some of the data value may be nil.
func (ctx *backfillContext) getChangedPatchRow(patchRecordID memCom.RecordID, upsertBatch *memCom.UpsertBatch) ([]*memCom.DataValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getChangedBaseRow get changed row from base batch if there are any changes to sort columns or array columns. It will fetch the whole
// row in base batch and apply patch value to it.
// for array columns, only when array size change will trigger the base batch change, value change while size not change will be covered in
// the unsorted column change
func (ctx *backfillContext) getChangedBaseRow(baseRecordID memCom.RecordID, changedPatchRow []*memCom.DataValue) []*memCom.DataValue {
	_ = "STUB: not implemented"
	return nil
}

// loop through sorted columns and array columns

// there's change in sorted column or size change in array column

// Mark deletion for this row.

// Copy the whole row in base batch to changed row and apply the change.

// column changed, get from patch

// column unchanged, get from base

// writePatchValueForUnsortColumn writes the patch value to forked columns if value changes.
// this function return false if no column update happens
func (ctx *backfillContext) writePatchValueForUnsortedColumn(baseRecordID memCom.RecordID, changedPatchRow []*memCom.DataValue) (updated bool) {
	_ = "STUB: not implemented"
	return false
}

// For updates to unsorted columns, if the value changes, fork the column, and update in place
// in the forked copy, this will make sure that ongoing queries do not see this change.

// For the forked columns, we will always allocate space for value vector and null vector despite
// of the mode of the original vector.

// will use original size for array vp

// Report before allocation.

// applyChangedRowToLiveStore applies changes in changedRow to temp live store.
func (ctx backfillContext) applyChangedRowToLiveStore(recordID memCom.RecordID, changedRow []*memCom.DataValue) {
	_ = "STUB: not implemented"
	return
}
