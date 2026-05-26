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

const data string = "data"
const redologs string = "redologs"
const snapshots string = "snapshots"
const archiveBatches string = "archiving_batches"

// Utils for data hierarchy layout.
// Following this wiki:
// https://github.com/uber/aresdb/wiki/data_disk_layout

// General path related utils for disk store.

// getPathForTableShard is used to get the directory to store a table shard given path prefix, table name and shard id.
func getPathForTableShard(prefix, table string, shardID int) string {
	_ = "STUB: not implemented"
	return ""
}

// Redologs Utils
// Path on disk:
//   {root_path}/data/{table_name}_{shard_id}/redologs/{creation_time}.redolog
//
// Sample:
//   /var/gForceDb/data/myTable_0/redologs/1499971253.redolog
//   /var/gForceDb/data/myTable_1/redologs/1499970221.redolog

// GetPathForTableRedologs is used to get the directory to store a table redolog given path prefix, table name and shard id.
func GetPathForTableRedologs(prefix, table string, shardID int) string {
	_ = "STUB: not implemented"
	return ""
}

// GetPathForRedologFile is used to get on disk file path given path prefix, table name, shard id and creationTime.
func GetPathForRedologFile(prefix, table string, shardID int, creationTime int64) string {
	_ = "STUB: not implemented"
	return ""
}

// Snapshot Utils
//Path on disk:
//  {root_path}/data/{table_name}_{shard_id}/snapshots/{redlo_log}_{offset}/{batchID}/{columnID}.data
//
//Sample:
//  /var/gForceDb/data/myTable_0/snapshots/1499970253_200/-2147483648/1.data
//  /var/gForceDb/data/myTable_1/snapshots/1499970221_300/-2147483648/2.data

// GetPathForTableSnapshotDir is used to get the dir path of a snapshot given path prefix, table name and shard id.
func GetPathForTableSnapshotDir(prefix, table string, shardID int) string {
	_ = "STUB: not implemented"
	return ""
}

// GetPathForTableSnapshotDirPath is used to get the dir path of a snapshot given path prefix, table name, shard id
// redo log file and offset.
func GetPathForTableSnapshotDirPath(prefix, table string, shardID int, redoLogFile int64, offset uint32) string {
	_ = "STUB: not implemented"
	return ""
}

// GetPathForTableSnapshotBatchDir is used to get the dir path of a snapshot batch given path prefix, table name,
// shard id, redo log file, offset and batchID.
func GetPathForTableSnapshotBatchDir(prefix, table string, shardID int, redoLogFile int64, offset uint32,
	batchID int) string {
	_ = "STUB: not implemented"
	return ""
}

// GetPathForTableSnapshotColumnFilePath is used to get the file path of a snapshot column given path prefix,
// table name, shard id, redo log file, offset, batchID and columnID
func GetPathForTableSnapshotColumnFilePath(prefix, table string, shardID int, redoLogFile int64, offset uint32,
	batchID, columnID int) string {
	_ = "STUB: not implemented"
	return ""
}

// Archive batches Utils
// Path on disk:
//   {root_path}/data/{table_name}_{shard_id}/archiving_batches/{batch_id}_{batch_version}
//   {root_path}/data/{table_name}_{shard_id}/archiving_batches/{batch_id}_{batch_version}/{columnID}.data
// Note:
//   batch_id is UTC date
// 	 batch_version is the cutoff seconds in unix time.
//
// Sample:
//   /var/gForceDb/data/myTable_0/archiving_batches/2017-07-19_1499971253/1.data
//   /var/gForceDb/data/myTable_0/archiving_batches/2017-07-19_1499971253/2.data

// GetPathForTableArchiveBatchRootDir is used to get root directory path for archive batch given path prefix, table name and shard id.
func GetPathForTableArchiveBatchRootDir(prefix, table string, shardID int) string {
	_ = "STUB: not implemented"
	return ""
}

// GetPathForTableArchiveBatchDir is used to get the dir path of an archive batch version given path prefix, table name, shard id, batch id and batch version.
func GetPathForTableArchiveBatchDir(prefix, table string, shardID int, batchID string, batchVersion uint32, seqNum uint32) string {
	_ = "STUB: not implemented"
	return ""
}

// largest uint32 indicates there's no seqNum

// GetPathForTableArchiveBatchColumnFile is used to get the file path of a column inside an archive batch version given path prefix, table name, shard id, batch id, batch version and column id.
func GetPathForTableArchiveBatchColumnFile(prefix, table string, shardID int, batchID string, batchVersion uint32, seqNum uint32, columnID int) string {
	_ = "STUB: not implemented"
	return ""
}

// ParseBatchIDAndVersionName will parse a batchIDAndVersion into batchID and batchVersion+seqNum.
func ParseBatchIDAndVersionName(batchIDAndVersion string) (string, uint32, uint32, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}
