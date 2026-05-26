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

	"github.com/uber/aresdb/datanode/bootstrap"
	"github.com/uber/aresdb/diskstore"
	"github.com/uber/aresdb/memstore/common"
	metaCom "github.com/uber/aresdb/metastore/common"
)

// TableShard stores the data for one table shard in memory.
type TableShard struct {
	// Wait group used to prevent the stores from being prematurely deleted.
	Users sync.WaitGroup `json:"-"`

	ShardID int `json:"-"`

	// For convenience, reference to the table schema struct.
	Schema *common.TableSchema `json:"schema"`

	// For convenience.
	metaStore metaCom.MetaStore
	diskStore diskstore.DiskStore
	options   Options

	// Live store. Its locks also cover the primary key.
	LiveStore *LiveStore `json:"liveStore"`

	// Archive store.
	ArchiveStore *ArchiveStore `json:"archiveStore"`

	// The special column deletion lock,
	// see https://docs.google.com/spreadsheets/d/1QI3s1_4wgP3Cy-IGoKFCx9BcN23FzIfZGRSNC8I-1Sk/edit#gid=0
	columnDeletion sync.Mutex

	// For convenience.
	HostMemoryManager common.HostMemoryManager `json:"-"`

	// bootstrapLock protects BootstrapState
	bootstrapLock  sync.RWMutex
	BootstrapState bootstrap.BootstrapState `json:"BootstrapState"`

	// BootstrapDetails shows the details of bootstrap
	BootstrapDetails bootstrap.BootstrapDetails `json:"bootstrapDetails,omitempty"`
	// needPeerCopy mark whether the table shard need to copy data from peer
	// before own disk data is available for serve
	// default to 0 (no need for peer copy)
	needPeerCopy uint32
}

// NewTableShard creates and initiates a table shard based on the schema.
func NewTableShard(schema *common.TableSchema,
	metaStore metaCom.MetaStore,
	diskStore diskstore.DiskStore,
	hostMemoryManager common.HostMemoryManager,
	shard int,
	totalShardsInCluster int,
	options Options,
) *TableShard {
	_ = "STUB: not implemented"
	return nil
}

// Destruct destructs the table shard.
// Caller must detach the shard from memstore first.
func (shard *TableShard) Destruct() {
	_ = "STUB: not implemented"
	// TODO: if this blocks on archiving for too long, figure out a way to cancel it.
	return
}

// DeleteColumn deletes the data for the specified column.
func (shard *TableShard) DeleteColumn(columnID int) error { _ = "STUB: not implemented"; return nil }

// Delete from live store

// Delete from disk store
// Schema cannot be changed while this function is called.
// Only delete unsorted columns from disk.

// Delete from archive store

// PreloadColumn loads the column into memory and wait for completion of loading
// within (startDay, endDay]. Note endDay is inclusive but startDay is exclusive.
func (shard *TableShard) PreloadColumn(columnID int, startDay int, endDay int) {
	_ = "STUB: not implemented"
	return
}

// Only do loading if this batch does not have any data yet.
