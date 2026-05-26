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

	metaCom "github.com/uber/aresdb/metastore/common"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/uber/aresdb/memstore/common"
)

// preloadJob defines the job struct to preload column when preloading days is changed.
type preloadJob struct {
	tableName         string
	columnID          int
	oldPreloadingDays int
	newPreloadingDays int
}

type hostMemoryManager struct {
	sync.RWMutex
	memStore  *memStoreImpl
	metaStore metaCom.MetaStore
	// totalMemorySize is configurable uplimit.
	totalMemorySize int64
	// unManagedMemorySize and managedMemorySize reflect current memory usage status.
	unManagedMemorySize int64
	managedMemorySize   int64
	// Init maps for table -> (columnID -> columnBatchInfos) mapping
	batchInfosByColumn map[string]map[int]*columnBatchInfos
	// channel to send preloadJob.
	preloadJobChan chan preloadJob
	// channel to stop preload go routines.
	preloadStopChan chan struct{}
	// channel to send eviction job.
	// TODO: if later we find to many go routines waiting for the channel,
	// we will use sync.Cond to rewrite it.
	evictionJobChan chan struct{}
	// channel to stop eviction go routines.
	evictionStopChan chan struct{}
}

// shardBatchID is the internal data holder struct to store
// shardID and batchID which used as key in the columnBatchInfos.
type shardBatchID struct {
	shardID int
	batchID int
}

func newShardBatchID(shardID int, batchID int) shardBatchID {
	_ = "STUB: not implemented"
	return *new(shardBatchID)
}

// shardBatchIDComparator provides a basic comparison on shardBatchID
func shardBatchIDComparator(a, b interface{}) int { _ = "STUB: not implemented"; return 0 }

// columnBatchInfos is using RB-Tree data structure to hold shardBatchID to
// size mapping
type columnBatchInfos struct {
	table         string
	batchInfoByID *rbt.Tree
	sync.RWMutex
}

func newColumnBatchInfos(table string) *columnBatchInfos { _ = "STUB: not implemented"; return nil }

// SetManagedObject is used to add a new batch/update an existing batch.
// Returns the bytes changes during this operation. For new batch, it's
// same as bytes value. For update batch, it's the value of
// bytesChanges = (currentBytes - oldBytes).
func (a *columnBatchInfos) SetManagedObject(shard, batchID int, bytes int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// deleteManagedObject is used to delete a batch.
// Returns the bytes got changed.
func (a *columnBatchInfos) DeleteManagedObject(shard, batchID int) int64 {
	_ = "STUB: not implemented"
	return 0
}

// GetArchiveMemoryUsageByShard returns memory usage [preload, non-preload] by shard
func (a *columnBatchInfos) GetArchiveMemoryUsageByShard(preloadDays int) map[int]*common.ColumnMemoryUsage {
	_ = "STUB: not implemented"
	return nil
}

// NewHostMemoryManager is used to init a HostMemoryManager.
func NewHostMemoryManager(memStore *memStoreImpl, totalMemorySize int64) common.HostMemoryManager {
	_ = "STUB: not implemented"
	return *new(common.HostMemoryManager)
}

// All the following three functions trigger preloading and eviction
// asynchrounously. In addition, as time goes on, reloading and eviction can
// also be triggered automatically.

// ReportUnmanagedSpaceUsageChange : Increase/Decrease bytes to the unmanaged space
// usage (for live batches and PKeys).
// Positive bytes number means to increase UnmanagedSpaceUsage, negative number
// means to decrease UnmanagedSpaceUsage.
func (h *hostMemoryManager) ReportUnmanagedSpaceUsageChange(bytes int64) {
	_ = "STUB: not implemented"
	return
}

// ReportManagedObject : Report space usage for a managed object (archive batch vector party).
func (h *hostMemoryManager) ReportManagedObject(table string, shard, batchID, columnID int, bytes int64) {
	_ = "STUB: not implemented"
	return
}

// Start will do a blocking preloading first and then start the go routines to do
// data preloading and eviction.
func (h *hostMemoryManager) Start() { _ = "STUB: not implemented"; return }

// Preloader execution loop.

// Evictor execution loop.

// Stop stops the gom rountines to do data preloading and eviction. It's a
// blocking call.
func (h *hostMemoryManager) Stop() { _ = "STUB: not implemented"; return }

// TriggerPreload will handle the column preloading days config change and
// trigger the column preloading if necessary. It's a asynchronous call.
func (h *hostMemoryManager) TriggerPreload(tableName string, columnID int,
	oldPreloadingDays int, newPreloadingDays int) {
	_ = "STUB: not implemented"
	return
}

// TriggerEviction triggers the eviction. It's a asynchronous call.
func (h *hostMemoryManager) TriggerEviction() { _ = "STUB: not implemented"; return }

func (h *hostMemoryManager) getUnmanagedSpaceUsage() int64 { _ = "STUB: not implemented"; return 0 }

func (h *hostMemoryManager) getManagedSpaceUsage() int64 { _ = "STUB: not implemented"; return 0 }

// GetArchiveMemoryUsageByTableShard get the managed memory details by table shard and column
func (h *hostMemoryManager) GetArchiveMemoryUsageByTableShard() (map[string]map[string]*common.ColumnMemoryUsage, error) {
	_ = "STUB: not implemented"
	return nil,

		// tableName_shardID -> columnName -> columnMemoryUsage
		nil
}

// ignore deleted table

// managedObjectExists : Return whether the corresponding managed object exists in managed memory.
func (h *hostMemoryManager) managedObjectExists(table string, shard, batchID, columnID int) bool {
	_ = "STUB: not implemented"
	return false
}

// AddOrUpdateManagedObject : Report space usage increase or update for a managed object (archive batch vector party).
func (h *hostMemoryManager) addOrUpdateManagedObject(table string, shard, batchID, columnID int, bytes int64) {
	_ = "STUB: not implemented"
	return
}

// deleteManagedObject : Report space usage reduce for a managed object (archive batch vector party).
func (h *hostMemoryManager) deleteManagedObject(table string, shard, batchID, columnID int) {
	_ = "STUB: not implemented"
	return
}

// handleColumnPreloadingDaysChange handles the preloading config change for a column.
func (h *hostMemoryManager) handleColumnPreloadingDaysChange(j preloadJob) {
	_ = "STUB: not implemented"
	return
}

// snapshot shardIDs.

// Table shard may have already been removed from this node.

// tryEviction : try to trigger eviction once
// unManagedMem + managedMem > totalAssignedMem. This method will pop batches
// from the per column holder data structure, calculate global priority
// based on column metadata, then push the batch into a priority queue.
// Eviction will happen through all the populated batches until memory usage
// decreases to a certain level. All failed eviction batches will be
// reinserted.
func (h *hostMemoryManager) tryEviction() {
	_ = "STUB: not implemented"
	// Check if eviction should be triggered
	return
}

// Init all columnar priority batches.

// Pop from globalPriorityQueueWithLock and do eviction

// Adding the corresponding next batch into priority queue.

// Still cannot meet the memory constraints even after evictions.

// pushBatchIntoGlobalPriorityQueue will generate a globalPriority object then
// push it into globalPriorityQueueWithLock.
func (gpq *globalPriorityQueue) pushBatchIntoGlobalPriorityQueue(h *hostMemoryManager,
	columnBatchInfos *columnBatchInfos, columnID int, columnIt rbt.Iterator) {
	_ = "STUB: not implemented"

	// batchInfo := batchInfoInterface.(*archiveBatchInfo)
	return
}

// initialGlobalPriorityQueue will initialize a globalPriorityQueueWithLock and fetch
// one batch for each table column from batchInfosByColumn.
func (h *hostMemoryManager) initialGlobalPriorityQueue() *globalPriorityQueue {
	_ = "STUB: not implemented"
	return nil
}

// globalPriority is the holding struct for all the neccesarry fields to
// compare batch priority in a global view.
type globalPriority struct {
	shardID  int
	columnID int

	// globalPriority comparison is based on the below 4 fields.
	isPreloading   bool
	columnPriority int64
	batchID        int
	size           int64
}

// globalPriorityComparator provides a basic comparison on globalPriority
func globalPriorityComparator(a, b interface{}) int { _ = "STUB: not implemented"; return 0 }

func createBatchPriority(shardID, columnID int, isPreloading bool, columnPriority int64, batchID int, size int64) *globalPriority {
	_ = "STUB: not implemented"
	return nil
}

// globalPriorityQueue definition START.
// A globalPriorityQueue implements heap.Interface and holds many Item pointers.
// Sorting method is based on globalPriorityComparator for struct globalPriority.

type globalPriorityItem struct {
	value    *columnBatchInfos
	it       rbt.Iterator
	priority *globalPriority
}

type globalPriorityQueue []*globalPriorityItem

func (gpq globalPriorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (gpq globalPriorityQueue) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return false
}

func (gpq globalPriorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (gpq *globalPriorityQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (gpq *globalPriorityQueue) push(item *globalPriorityItem) { _ = "STUB: not implemented"; return }

func (gpq *globalPriorityQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (gpq *globalPriorityQueue) pop() *globalPriorityItem { _ = "STUB: not implemented"; return nil }

func newGlobalPriorityQueue() *globalPriorityQueue { _ = "STUB: not implemented"; return nil }

func (gpq *globalPriorityQueue) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (gpq *globalPriorityQueue) size() int {
	_ = "STUB: not implemented"

	// globalPriorityQueue definition END.
	return 0
}

// isPreloadingBatch will check if a given batchID falling into the preloading
// zone.
// batchID is daysSinceEpoch value.
func isPreloadingBatch(batchID, preloadingDays int) bool { _ = "STUB: not implemented"; return false }
