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

	memcom "github.com/uber/aresdb/memstore/common"
)

// PlayRedoLog loads data for the table Shard from disk store and recovers the Shard for serving.
func (shard *TableShard) PlayRedoLog() { _ = "STUB: not implemented"; return }

// check if this batch has already been backfilled and persisted

// for normal ingestion, will log error and keep going

// report redolog size after replay

// proactively purge redo files

func (shard *TableShard) cleanOldSnapshotAndLogs(redoLogFile int64, offset uint32) {
	_ = "STUB: not implemented"
	return
}

// snapshot won't care about the cutoff.

// delete old snapshots

// LoadMetaData loads metadata for the table Shard from metastore.
func (shard *TableShard) LoadMetaData() error { _ = "STUB: not implemented"; return nil }

// We set the archiving cutoff to the persisted value (CLW) in meta so recovery will apply
// all items in redolog that have event time > CLW. The backfill job will ignore items in
// the backfill that have associated CHW (cutoff high watermark) > persisted CHW.

// retrieve redoLog/offset checkpointed for backfill

// retrieve latest snapshot info

// loadSnapshots load snapshots for dimension tables
func (m *memStoreImpl) loadSnapshots() { _ = "STUB: not implemented"; return }

// playRedoLogs replay redo logs for all tables in parallel, and then start the data ingestion
func (m *memStoreImpl) playRedoLogs() { _ = "STUB: not implemented"; return }

// Replay all redologs

// InitShards loads/recovers data for shards initially owned by the current instance.
// It also watches Shard ownership change events and handles them in a separate goroutine.
// InitShards is only used in non sharded version which assume totalShardsInCluster to be one
func (m *memStoreImpl) InitShards(schedulerOff bool, shardOwner topology.ShardOwner) {
	_ = "STUB: not implemented"
	return
}

// tryPreload data according the column retention config and start the go routines
// to do eviction and preloading.

// Start host memory manager

// load snapshot for dimension tables

// start scheduler after we load all the metadata. This ensure we can start backfill job earlier to consume
// the backfill queue.

// Start scheduler.

// disable archiving during redolog replay

// this will start scheduler of all jobs except archiving, archiving will be started individually

// re-enable archiving after redolog replay

// watch Shard ownership change

// Shard ownership change handling

// This assumes that (certain) schema change must wait until the Shard
// is fully loaded, which may take a while.

// Unload the Shard.

// Detach first.

// Destruct.

// Do not delete the file on diskstore.

// LoadShard loads/recovers the specified Shard and attaches it to memStoreImpl for serving. If will load the metadata
// first and then replay redologs only if replayRedologs is true.
// LoadShard is only used in non sharded version, whihch assume totalShardsInCluster to be 1
func (m *memStoreImpl) LoadShard(schema *memcom.TableSchema, shard int, replayRedologs bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Add reporter for current table and Shard.

// LoadSnapshot load shard data from snapshot files
func (shard *TableShard) LoadSnapshot() error { _ = "STUB: not implemented"; return nil }

// no snapshot created yet

// find all columns in snapshot dir

//reset back the read/write record position

func (shard *TableShard) loadTableShardSnapshot(
	tableName string, shardID int,
	batchID int32, redoLogFile int64, offset uint32) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// find all columns in snapshot dir

// found the column in snapshot, read from snapshot file

func (shard *TableShard) rebuildIndexForLiveStore(batchID int32, lastRecord uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// truncate key before every read
