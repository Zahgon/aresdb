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

// Snapshot is the process to write the current content of dimension table live store in memory to disk in order to
//
//	1.Facilitate recovery during server bootstrap.
//	2.Purge stale redo logs.
func (m *memStoreImpl) Snapshot(table string, shardID int, reporter SnapshotJobDetailReporter) error {
	_ = "STUB: not implemented"
	return nil
}

// keep the current redofile and offset

// checkpoint snapshot progress

func (m *memStoreImpl) createSnapshot(shard *TableShard, redoFile int64, batchOffset uint32) error {
	_ = "STUB: not implemented"
	// Block column deletion
	return nil
}

// column deleted likely
