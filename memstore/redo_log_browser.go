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
	"github.com/uber/aresdb/diskstore"
	"github.com/uber/aresdb/memstore/common"

	"github.com/uber/aresdb/utils"
)

// RedoLogBrowser is the interface to list redo log files, upsert batches and read upsert batch data.
type RedoLogBrowser interface {
	ListLogFiles() ([]int64, error)
	ListUpsertBatch(creationTime int64) ([]int64, error)
	ReadData(creationTime int64, upsertBatchOffset int64, start int, length int) (
		[][]interface{}, []string, int, error)
}

// redoLogBrowser is the implementation of RedoLogBrowser.
type redoLogBrowser struct {
	tableName string
	shardID   int
	diskStore diskstore.DiskStore
	schema    *common.TableSchema
}

// ListLogFiles lists all log files of a given table Shard.
func (rb *redoLogBrowser) ListLogFiles() ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListUpsertBatches opens corresponding redo log file given creation time and returns starting offsets of upsert batches
// in this file.
func (rb *redoLogBrowser) ListUpsertBatch(creationTime int64) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read magic header.

// Offset starts from magical header.

// ReadData first locates the upsert batch using creationTime and upsertBatchOffset. It then returns data
// starting from given start and has length rows along with number of total rows and column names in this
// upsert batch.
func (rb *redoLogBrowser) ReadData(creationTime int64, upsertBatchOffset int64, start int, length int) (
	data [][]interface{}, columnNames []string, numRows int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// readUpsertBatch reads an upsert batch from current offset of a stream.
func (rb *redoLogBrowser) readUpsertBatch(f utils.ReaderSeekerCloser) (*common.UpsertBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewRedoLogBrowser creates a RedoLogBrowser using field from Shard.
func (shard *TableShard) NewRedoLogBrowser() RedoLogBrowser {
	_ = "STUB: not implemented"
	return *new(RedoLogBrowser)
}
