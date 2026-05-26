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

// Purge purges out of retention data for table shard
func (m *memStoreImpl) Purge(tableName string, shardID, batchIDStart, batchIDEnd int, reporter PurgeJobDetailReporter) error {
	_ = "STUB: not implemented"
	// check if there is peer bootstraping job running, skip this purge if we can not acquire the token
	return nil
}

// detach batch from memory

// delete metadata of batches within range

// delete data file on disk of batches within range

// wait for users to finish
