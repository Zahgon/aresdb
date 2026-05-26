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

	"github.com/uber/aresdb/cluster/topology"
	"github.com/uber/aresdb/datanode/bootstrap"

	"github.com/uber/aresdb/diskstore"
	"github.com/uber/aresdb/memstore/common"
	metaCom "github.com/uber/aresdb/metastore/common"
)

// TableShardMemoryUsage contains memory usage for column memory and primary key memory usage
type TableShardMemoryUsage struct {
	ColumnMemory     map[string]*common.ColumnMemoryUsage `json:"cols"`
	PrimaryKeyMemory uint                                 `json:"pk"`
}

// MemStore defines the interface for managing multiple table shards in memory. This is for mocking
// in unit tests
type MemStore interface {
	common.TableSchemaReader
	bootstrap.Bootstrapable

	// GetMemoryUsageDetails
	GetMemoryUsageDetails() (map[string]TableShardMemoryUsage, error)
	// GetScheduler returns the scheduler for scheduling archiving and backfill jobs.
	GetScheduler() Scheduler
	// GetHostMemoryManager returns the host memory manager
	GetHostMemoryManager() common.HostMemoryManager
	// AddTableShard add a table shard to the memstore
	AddTableShard(table string, shardID int, totalShards int, needPeerCopy bool, needPurge bool)
	// GetTableShard gets the data for a pinned table Shard. Caller needs to unpin after use.
	GetTableShard(table string, shardID int) (*TableShard, error)
	// RemoveTableShard removes table shard from memstore
	RemoveTableShard(table string, shardID int)
	// FetchSchema fetches schema from metaStore and updates in-memory copy of table schema,
	// and set up watch channels for metaStore schema changes, used for bootstrapping mem store.
	FetchSchema() error
	// InitShards loads/recovers data for shards initially owned by the current instance.
	InitShards(schedulerOff bool, shardOwner topology.ShardOwner)
	// HandleIngestion logs an upsert batch and applies it to the in-memory store.
	HandleIngestion(table string, shardID int, upsertBatch *common.UpsertBatch) error
	// Archive is the process moving stable records in fact tables from live batches to archive
	// batches.
	Archive(table string, shardID int, cutoff uint32, reporter ArchiveJobDetailReporter) error

	// Backfill is the process of merging records with event time older than cutoff with
	// archive batches.
	Backfill(table string, shardID int, reporter BackfillJobDetailReporter) error

	// Snapshot is the process to write the current content of dimension table live store in memory to disk.
	Snapshot(table string, shardID int, reporter SnapshotJobDetailReporter) error

	// Purge is the process to purge out of retention archive batches
	Purge(table string, shardID, batchIDStart, batchIDEnd int, reporter PurgeJobDetailReporter) error
}

// memStoreImpl implements the MemStore interface.
type memStoreImpl struct {
	// memStoreImpl mutex is used to protect the TableShards and TableSchemas maps.
	//
	// For Shard access:
	//   Readers/writers must call TableShard.liveStore.Users.Add(1)
	//   before releasing this mutex, and call
	//   TableShard.liveStore.Users.Done() after their businesses.
	//
	//   Table Shard deleter must detach the Shard first, and then call
	//   TableShard.liveStore.Users.Wait() before deleting the Shard.
	//
	// For schema access:
	//   User should lock the TableSchema before releasing this mutex.
	//
	sync.RWMutex
	// Table name and Shard ID as the map keys.
	TableShards map[string]map[int]*TableShard
	// Schema for all tables in the system. Schemas are not deleted for simplicity
	TableSchemas map[string]*common.TableSchema

	HostMemManager common.HostMemoryManager

	// reference to metaStore for registering watchers,
	// fetch latest schema and store Shard versions.
	metaStore metaCom.MetaStore
	diskStore diskstore.DiskStore
	options   Options

	// each MemStore should only have one scheduler instance.
	scheduler Scheduler
}

func getTableShardKey(tableName string, shardID int) string { _ = "STUB: not implemented"; return "" }

// NewMemStore creates a MemStore from the specified MetaStore.
func NewMemStore(metaStore metaCom.MetaStore, diskStore diskstore.DiskStore, options Options) MemStore {
	_ = "STUB: not implemented"
	return *new(MemStore)
}

// Create HostMemoryManager

func (m *memStoreImpl) GetMemoryUsageDetails() (map[string]TableShardMemoryUsage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// primary key memory usage

// archive memory usage

// live store memory usage

func (shard *TableShard) getLiveMemoryUsageByColumns(columnMemory map[string]*common.ColumnMemoryUsage) {
	_ = "STUB: not implemented"
	return
}

// GetTableShard gets the data for a pinned table Shard. Caller needs to unpin after use.
func (m *memStoreImpl) GetTableShard(table string, shardID int) (*TableShard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSchema returns schema for a table.
func (m *memStoreImpl) GetSchema(table string) (*common.TableSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSchemas returns all table schemas. Callers need to hold a reader lock to access this function.
func (m *memStoreImpl) GetSchemas() map[string]*common.TableSchema {
	_ = "STUB: not implemented"
	return nil

	// GetScheduler returns the scheduler instance bound to the MemStore.
}

func (m *memStoreImpl) GetScheduler() Scheduler {
	_ = "STUB: not implemented"

	// TryEvictBatchColumn tries to evict a column from a given table/Shard/batchID.
	// Return values are the check for column is deleted or not and error.
	return *new(Scheduler)
}

func (m *memStoreImpl) TryEvictBatchColumn(table string, shardID int, batchID int32, columnID int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *memStoreImpl) AddTableShard(table string, shardID int, totalShards int, needPeerCopy bool, needPurge bool) {
	_ = "STUB: not implemented"
	return
}

// table might get deleted at this point

// create new shard

// purge to make sure disk space is clean for new table shard when it is added

func (m *memStoreImpl) RemoveTableShard(table string, shardID int) {
	_ = "STUB: not implemented"
	return

	// Detach first.
}

// Destruct.

// preloadAllFactTables preloads recent days data for all columns of all table shards into memory.
// The number of preloading days is defined at each column level. This call will happen at
// shard initialization stage.
func (m *memStoreImpl) preloadAllFactTables() { _ = "STUB: not implemented"; return }

// snapshot (tableName, shardID)s.

// Table shard may have already been removed from this node.

func (m *memStoreImpl) GetHostMemoryManager() common.HostMemoryManager {
	_ = "STUB: not implemented"
	return *new(common.HostMemoryManager)
}
