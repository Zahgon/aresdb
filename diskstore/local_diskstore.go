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

package diskstore

import (
	"io"
	"time"

	"github.com/uber/aresdb/common"
	"github.com/uber/aresdb/utils"
)

// LocalDiskStore is the implementation of Diskstore for local disk.
type LocalDiskStore struct {
	rootPath        string
	diskStoreConfig common.DiskStoreConfig
}

// NewLocalDiskStore is used to init a LocalDiskStore with rootPath.
func NewLocalDiskStore(rootPath string) DiskStore {
	_ = "STUB: not implemented"
	return *new(DiskStore)
}

const timeFormatForBatchID = "2006-01-02"

// Table shard level operation

// DeleteTableShard : Completely wipe out a table shard.
func (l LocalDiskStore) DeleteTableShard(table string, shard int) error {
	_ = "STUB: not implemented"
	return nil
}

// Redo Logs

// ListLogFiles : Returns the file creation unix time in second for each log file as a sorted slice.
func (l LocalDiskStore) ListLogFiles(table string, shard int) (creationUnixTime []int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The redo log directory won't get created until the first append call.

// Failed to parse redolog file, will continue

// OpenLogFileForReplay : Opens the specified log file for replay.
func (l LocalDiskStore) OpenLogFileForReplay(table string, shard int,
	creationTime int64) (utils.ReaderSeekerCloser, error) {
	_ = "STUB: not implemented"
	return *new(utils.ReaderSeekerCloser), nil
}

// OpenLogFileForAppend : Opens/creates the specified log file for append.
func (l LocalDiskStore) OpenLogFileForAppend(table string, shard int, creationTime int64) (utils.WriteSyncCloser, error) {
	_ = "STUB: not implemented"
	return *new(utils.WriteSyncCloser), nil
}

// DeleteLogFile is used to delete a specified redolog.
func (l LocalDiskStore) DeleteLogFile(table string, shard int, creationTime int64) error {
	_ = "STUB: not implemented"
	return nil
}

// TruncateLogFile is used to truncate redolog to drop the last incomplete/corrupted upsert batch.
func (l LocalDiskStore) TruncateLogFile(table string, shard int, creationTime int64, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Snapshot files.

// ListSnapshotBatches : Returns the batch directories at the specified version.
func (l LocalDiskStore) ListSnapshotBatches(table string, shard int,
	redoLogFile int64, offset uint32) (batches []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No batches for this snapshot

// ListSnapshotVectorPartyFiles : Returns the vector party files under specific batch directory.
func (l LocalDiskStore) ListSnapshotVectorPartyFiles(table string, shard int,
	redoLogFile int64, offset uint32, batchID int) (columnIDs []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l LocalDiskStore) readVectoryPartyFiles(dir string) (columnIDs []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenSnapshotVectorPartyFileForRead : Opens the snapshot file for read at the specified version.
func (l LocalDiskStore) OpenSnapshotVectorPartyFileForRead(table string, shard int,
	redoLogFile int64, offset uint32, batchID int, columnID int) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// OpenSnapshotVectorPartyFileForWrite : Creates/truncates the snapshot file for write at the specified version.
func (l LocalDiskStore) OpenSnapshotVectorPartyFileForWrite(table string, shard int,
	redoLogFile int64, offset uint32, batchID int, columnID int) (utils.WriteSyncCloser, error) {
	_ = "STUB: not implemented"
	return *new(utils.WriteSyncCloser), nil
}

// DeleteSnapshot : Deletes snapshot directories **older than** the specified version (redolog file and offset).
func (l LocalDiskStore) DeleteSnapshot(table string, shard int, latestRedoLogFile int64, latestOffset uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Failed to parse snapshot file name, will skip.

// Failed to parse snapshot file name, will skip.

// Archived vector party files.

// ListArchiveBatchVectorPartyFiles return all vp for one batch version/seq
func (l LocalDiskStore) ListArchiveBatchVectorPartyFiles(table string, shard, batchID int,
	batchVersion uint32, seqNum uint32) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenVectorPartyFileForRead : Opens the vector party file at the specified batchVersion for read.
func (l LocalDiskStore) OpenVectorPartyFileForRead(table string, columnID int, shard, batchID int, batchVersion uint32,
	seqNum uint32) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// OpenVectorPartyFileForWrite : Creates/truncates the vector party file at the specified batchVersion for write.
func (l LocalDiskStore) OpenVectorPartyFileForWrite(table string, columnID int, shard, batchID int, batchVersion uint32,
	seqNum uint32) (utils.WriteSyncCloser, error) {
	_ = "STUB: not implemented"
	return *new(utils.WriteSyncCloser), nil
}

// DeleteBatchVersions deletes all old batches with the specified batchID that have version lower than or equal to
// the specified batch  version. All columns of those batches will be deleted.
func (l LocalDiskStore) DeleteBatchVersions(table string, shard, batchID int, batchVersion uint32, seqNum uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteBatches : Deletes all batches within [batchIDStart, batchIDEnd)
func (l LocalDiskStore) DeleteBatches(table string, shard, batchIDStart, batchIDEnd int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DeleteColumn : Deletes all batches of the specified column.
func (l LocalDiskStore) DeleteColumn(table string, columnID int, shard int) error {
	_ = "STUB: not implemented"
	return nil
}

func daysSinceEpochToTime(daysSinceEpoch int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func daysSinceEpochToTimeStr(daysSinceEpoch int) string { _ = "STUB: not implemented"; return "" }
