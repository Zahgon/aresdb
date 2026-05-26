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

	"github.com/uber/aresdb/diskstore"
	"github.com/uber/aresdb/memstore/common"
)

// archiveVectorParty is the implementation of ArchiveVectorParty
type archiveVectorParty struct {
	cVectorParty
	common.Pinnable
}

// Prune judges column mode first and sets the mode to vector party.
// Afterwards it purges unnecessary vectors based on the column mode.
func (vp *archiveVectorParty) Prune() { _ = "STUB: not implemented"; return }

// GetCount implements GetCount interface function in archiveVectorParty.
func (vp *archiveVectorParty) GetCount(offset int) uint32 { _ = "STUB: not implemented"; return 0 }

// SetCount implements SetCount interface function in archiveVectorParty.
func (vp *archiveVectorParty) SetCount(offset int, count uint32) { _ = "STUB: not implemented"; return }

// CopyOnWrite clone vector party for updates
// Only work for uncompressed archive vector party, Mode 3 vector party (has count) cannot be cloned for write
func (vp *archiveVectorParty) CopyOnWrite(batchSize int) common.ArchiveVectorParty {
	_ = "STUB: not implemented"
	return *new(common.ArchiveVectorParty)
}

// archive vector party should always have allUsersDone initialized correctly with batch rwlock

// All values present, we need to set all bits to 1.

// LoadFromDisk load archive vector party from disk
// caller should lock archive batch before using
func (vp *archiveVectorParty) LoadFromDisk(hostMemManager common.HostMemoryManager, diskStore diskstore.DiskStore, table string, shardID int, columnID, batchID int, batchVersion uint32, seqNum uint32) {
	_ = "STUB: not implemented"
	return
}

// newArchiveVectorParty creates a archive store vector party,
// archiveVectorParty use c allocated memory
func newArchiveVectorParty(length int, dataType common.DataType, defaultValue common.DataValue, locker sync.Locker) *archiveVectorParty {
	_ = "STUB: not implemented"
	return nil
}
