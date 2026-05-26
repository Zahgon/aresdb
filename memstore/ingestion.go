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
	"github.com/uber/aresdb/memstore/common"
)

// HandleIngestion logs an upsert batch and applies it to the in-memory store.
func (m *memStoreImpl) HandleIngestion(table string, shardID int, upsertBatch *common.UpsertBatch) error {
	_ = "STUB: not implemented"
	return nil
}

// Release the wait group that proctects the shard to be deleted.

// saveUpsertBatch handles data ingestion from both redolog and http
func (shard *TableShard) saveUpsertBatch(upsertBatch *common.UpsertBatch, redoLogFile int64, offset uint32, recovery, skipBackFillRows bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Put a 0 in maxEventTimePerFile in case this is redolog is full of backfill batches.

// for non-recovery and local file based redolog, need write the upsertbatch into redolog file

// change original file/offset to be local redolog file/offset

// return immediately if it does not need to wait for backfill buffer availability

// otherwise: block until backfill buffer becomes available again

// ApplyUpsertBatch applies the upsert batch to the memstore shard.
// Returns true if caller needs to wait for availability of backfill buffer
func (shard *TableShard) ApplyUpsertBatch(upsertBatch *common.UpsertBatch, redoLogFile int64, offset uint32, skipBackfillRows bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsFactTable should be immutable.

// This is the upsertbatch column index that points to the first logic column (which is
// event time for fact table).

// Validate columns in upsert batch are valid.

// For fact table ingestion, we will need to get the event time from the first column. we don't
// have to validate the column type in the upsertbatch because the loop above already handled it.

// We write insert records first so records with the same primary key in a upsert batch
// will be updated in order.

func (shard *TableShard) postUpsertBatchApplication(upsertBatch, backfillUpsertBatch *common.UpsertBatch, redoLogFile int64,
	offset uint32, numMutations int) bool {
	_ = "STUB: not implemented"
	return false
}

// add records to backfill queue if any.
// TODO: currently we're relying on LiveStore.WriterLock to guarantee the ordering of backfill batches

// caller needs to wait WHEN backfill buffer is full and there're more than 10 rows or over 5% rows are for backfill

// Per record instruction on how to read from upsert batch and write to memStore.
type recordInfo struct {
	// The row index of the record in the upsert batch.
	row int
	// The index of the to be inserted/updated record in a batch.
	index int
}

// Insert primary keys and return the records for update, insert grouped by batch.
// eventTimeColumnIndex will be used to extract the event time value per row if it >= 0.
func (shard *TableShard) insertPrimaryKeys(primaryKeyColumns []int, eventTimeColumnIndex int, redoLogFile int64,
	upsertBatch *common.UpsertBatch, skipBackfillRows bool) (
	map[int32][]recordInfo, map[int32][]recordInfo, *common.UpsertBatch, error) {
	_ = "STUB: not implemented"
	// Get primary key column indices and calculate the primary key width.
	return nil, nil, nil, nil
}

// Get primary key bytes for each record.
// truncate key

// For fact table we need to get the event time from the first column.

// event with invalid event time will be ignored
// once arrival time is older than archiving cutoff.

// Skip this record if it's out of retention

// Skip this record if its event time is latter than current time

// Update max event time so archiving won't purge redo log files that have records newer than
// archiving cut off time.

// If we get a record that is older than archiving cutoff time (exclusive) that means
// 1. during ingestion, the event should be put into a backfill queue
// 2. during recovery, the event should be ignored, because it was already put into
//    a backfill queue at ingestion time.

// mark this row as backfill row

// Update max event time for each column in this upsert batch.

// We only do it on per upsert batch level so it should be acceptable to create the scope dynamically.
// TODO: if there is any performance issue, cache the reporter at live store level.

// update ratio gauge of backfill rows/total rows

// create backfill upsertBatch if applicable

// all rows are for backfill

// some columns get pruned due to inappropriate update functions

// Read rows from a batch group and write to memStore. Batch id = 0 is for records to be inserted.
func (shard *TableShard) writeBatchRecords(columnDeletions []bool,
	upsertBatch *common.UpsertBatch, batchID int32, records []recordInfo, forUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

// We need to lock the batch for update to achieve row level consistency.

// Make sure all columns are created.

// Instead of traversing row by row, we instead do column by column to avoid making checks on each row.

// we will skip processing this column if
// 1. columnMode is AllValuesDefault
// 2. columnUpdateMode is UpdateOverwriteNotNull

// check whether the update mode is valid based on data type.

// always update

// We explicitly treat different columns by checking whether they are
// 1. Bool type
// 2. Go types
// 3. Other types
// Via doing this, we save lots of stack space to storing all related fields for different cases.

// only read oldValue when mode is one of add, min, max.

// Only need to do calculation when old value is valid, otherwise we can directly
// set what's in upsert batch.

// if the value is not updated, set the value directly using value from upsert batch.
