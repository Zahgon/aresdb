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

package metastore

import (
	"sync"

	"github.com/uber/aresdb/metastore/common"
	"github.com/uber/aresdb/utils"
)

// meaningful defaults of table configurations.
const (
	DefaultBatchSize                      = 2097152
	DefaultArchivingDelayMinutes          = 1440
	DefaultArchivingIntervalMinutes       = 180
	DefaultBackfillIntervalMinutes        = 60
	DefaultBackfillMaxBufferSize    int64 = 4294967296
	DefaultBackfillThresholdInBytes int64 = 2097152
	DefaultBackfillStoreBatchSize         = 20000
	DefaultRecordRetentionInDays          = 90
	DefaultSnapshotIntervalMinutes        = 360                  // 6 hours
	DefaultSnapshotThreshold              = 3 * DefaultBatchSize // 3 batches
	DefaultRedologRotationInterval        = 10800                // 3 hours
	DefaultMaxRedoLogSize                 = 1 << 30              // 1 GB
)

// DefaultTableConfig represents default table config
var DefaultTableConfig = common.TableConfig{
	BatchSize:                DefaultBatchSize,
	ArchivingIntervalMinutes: DefaultArchivingIntervalMinutes,
	ArchivingDelayMinutes:    DefaultArchivingDelayMinutes,
	BackfillMaxBufferSize:    DefaultBackfillMaxBufferSize,
	BackfillIntervalMinutes:  DefaultBackfillIntervalMinutes,
	BackfillThresholdInBytes: DefaultBackfillThresholdInBytes,
	BackfillStoreBatchSize:   DefaultBackfillStoreBatchSize,
	RecordRetentionInDays:    DefaultRecordRetentionInDays,
	SnapshotIntervalMinutes:  DefaultSnapshotIntervalMinutes,
	SnapshotThreshold:        DefaultSnapshotThreshold,
	RedoLogRotationInterval:  DefaultRedologRotationInterval,
	MaxRedoLogFileSize:       DefaultMaxRedoLogSize,
}

// disk-based metastore implementation.
// all validation of user input (eg. table/column name and table/column struct) will be pushed to api layer,
// which is the earliest point of user input, all schemas inside system will be already valid,
// Note:
// There are four types of write calls to MetaStore, the handling of each is different:
//  1. Schema Changes
//     synchronous, return after both writing to watcher channel and reading from done channel are done
//  2. Update EnumCases
//     return after changes persisted in disk and writing to watcher channel; does not read from done channel
//  3. Adding Watchers
//     3.1 for enum cases, create channels and push existing enum cases starting from start case to channel if any
//     3.2 for table list and table schema channels, create channels and return
//  4. Update configurations
//     configurations update including updates on archiving cutoff, snapshot version, archive batch version etc,
//     these changes does not need to be pushed to memstore.
//
// Operations involves writing to watcher channels (case 1 and 2), we need to enforce the order of changes pushed into channel,
// writeLock is introduced to enforce that.
// Other operations (case 3 and 4), we only need lock to protect internal data structure, a read write lock is used.
type diskMetaStore struct {
	sync.RWMutex
	utils.FileSystem

	// writeLock is to enforce single writer at a time
	// to make sure the same order of shema change when applied to
	// MemStore through watcher channel
	writeLock sync.Mutex

	// the base path for MetaStore in disk
	basePath string

	// tableListWatcher
	tableListWatcher chan<- []string
	// tableListDone is the channel for tracking whether watcher has
	// successfully got the table list change,
	// here we adopt a synchronous model for schema change.
	tableListDone <-chan struct{}

	// tableSchemaWatcher
	tableSchemaWatcher chan<- *common.Table
	// tableSchemaDone is the channel for tracking whether watcher has
	// successfully got the table schema change
	tableSchemaDone <-chan struct{}

	// enumDictWatchers
	// maps from tableName to columnName to watchers
	enumDictWatchers map[string]map[string]chan<- string
	// tableSchemaDone are the channels for tracking whether watcher has
	// successfully got the enum case change.
	enumDictDone map[string]map[string]<-chan struct{}

	// shardOwnershipWatcher
	shardOwnershipWatcher chan<- common.ShardOwnership
	// shardOwnershipDone is used for block waiting for the consumer to finish
	// processing each ownership change event.
	shardOwnershipDone <-chan struct{}
}

// ListTables list existing table names
func (dm *diskMetaStore) ListTables() ([]string, error) {
	_ = "STUB: not implemented"
	return nil,

		// GetTable return the table schema stored in metastore given tablename,
		// return ErrTableDoesNotExist if table not exists.
		nil
}

func (dm *diskMetaStore) GetTable(name string) (*common.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEnumDict gets the enum cases for given tableName and columnName
func (dm *diskMetaStore) GetEnumDict(tableName, columnName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetArchivingCutoff gets the latest archiving cutoff for given table and shard.
func (dm *diskMetaStore) GetArchivingCutoff(tableName string, shard int) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DeleteTableShard deletes all table shard level metadata
func (dm *diskMetaStore) DeleteTableShard(tableName string, shard int) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSnapshotProgress gets the latest snapshot progress for given table and shard
func (dm *diskMetaStore) GetSnapshotProgress(tableName string, shard int) (int64, uint32, int32, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

// UpdateArchivingCutoff updates archiving cutoff for given table (fact table), shard
func (dm *diskMetaStore) UpdateArchivingCutoff(tableName string, shard int, cutoff uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateSnapshotProgress update snapshot version for given table (dimension table), shard.
func (dm *diskMetaStore) UpdateSnapshotProgress(tableName string, shard int, redoLogFile int64, upsertBatchOffset uint32, lastReadBatchID int32, lastReadBatchOffset uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Updates the latest redolog/offset that have been backfilled for the specified shard.
func (dm *diskMetaStore) UpdateBackfillProgress(table string, shard int, redoFile int64, offset uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieve the latest redolog/offset that have been backfilled for the specified shard.
func (dm *diskMetaStore) GetBackfillProgressInfo(table string, shard int) (int64, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Update ingestion commit offset, used for kafka like streaming ingestion
func (dm *diskMetaStore) UpdateRedoLogCommitOffset(table string, shard int, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

// No sanity check here for schema/fact table/directory, assuming all should be passed before calling this func

// Get ingestion commit offset, used for kafka like streaming ingestion
func (dm *diskMetaStore) GetRedoLogCommitOffset(table string, shard int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// No sanity check here for schema/fact table/directory, assuming all should be passed before calling this func

// Update ingestion checkpoint offset, used for kafka like streaming ingestion
func (dm *diskMetaStore) UpdateRedoLogCheckpointOffset(table string, shard int, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

// No sanity check here for schema/fact table/directory, assuming all should be passed before calling this func

// Get ingestion checkpoint offset, used for kafka like streaming ingestion
func (dm *diskMetaStore) GetRedoLogCheckpointOffset(table string, shard int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// No sanity check here for schema/fact table/directory, assuming all should be passed before calling this func

// WatchTableListEvents register a watcher to table list change events,
// should only be called once,
// returns ErrWatcherAlreadyExist once watcher already exists
func (dm *diskMetaStore) WatchTableListEvents() (events <-chan []string, done chan<- struct{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// WatchTableSchemaEvents register a watcher to table schema change events,
// should be only called once,
// returns ErrWatcherAlreadyExist once watcher already exists
func (dm *diskMetaStore) WatchTableSchemaEvents() (events <-chan *common.Table, done chan<- struct{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// WatchEnumDictEvents register a watcher to enum cases change events for given table and column,
// returns
//
//	ErrTableDoesNotExist, ErrColumnDoesNotExist, ErrNotEnumColumn, ErrWatcherAlreadyExist.
//
// if startCase is larger than the number of current existing enum cases, it will be just as if receiving from
// latest.
func (dm *diskMetaStore) WatchEnumDictEvents(table, column string, startCase int) (events <-chan string, done chan<- struct{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// if start is larger than length of existing enum cases
// will treat as if sending from latest

// WatchTableSchemaEvents register a watcher to table schema change events,
// should be only called once,
// returns ErrWatcherAlreadyExist once watcher already exists
func (dm *diskMetaStore) WatchShardOwnershipEvents() (events <-chan common.ShardOwnership, done chan<- struct{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CreateTable creates a new Table,
// returns
//
//	ErrTableAlreadyExist if table already exists
func (dm *diskMetaStore) CreateTable(table *common.Table) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// append enum case for enum column with default value

// UpdateTable update table configurations
// return
//
//	ErrTableDoesNotExist if table does not exist
func (dm *diskMetaStore) UpdateTableConfig(tableName string, config common.TableConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// UpdateTable updates table schema and config
// table passed in should have been validated against existing table schema
func (dm *diskMetaStore) UpdateTable(table common.Table) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// append enum case for enum column with default value for new columns

// DeleteTable deletes a table
// return
//
//	ErrTableDoesNotExist if table does not exist
func (dm *diskMetaStore) DeleteTable(tableName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AddColumn adds a new column
// returns
//
//	ErrTableDoesNotExist if table does not exist
//	ErrColumnAlreadyExist if column already exists
func (dm *diskMetaStore) AddColumn(tableName string, column common.Column, appendToArchivingSortOrder bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// UpdateColumn deletes a column.
// return
//
//	ErrTableDoesNotExist if table does not exist.
//	ErrColumnDoesNotExist if column does not exist.
func (dm *diskMetaStore) UpdateColumn(tableName string, columnName string, config common.ColumnConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// DeleteColumn deletes a column
// return
//
//	ErrTableDoesNotExist if table not exist
//	ErrColumnDoesNotExist if column not exist
func (dm *diskMetaStore) DeleteColumn(tableName string, columnName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ExtendEnumDict extends enum cases for given table column
func (dm *diskMetaStore) ExtendEnumDict(table, columnName string, enumCases []string) (enumIDs []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PurgeArchiveBatches deletes the archive batches' metadata with batchID within [batchIDStart, batchIDEnd)
func (dm *diskMetaStore) PurgeArchiveBatches(tableName string, shard, batchIDStart, batchIDEnd int) error {
	_ = "STUB: not implemented"
	return nil
}

// OverwriteArchiveBatchVersion overwrites batch version
func (dm *diskMetaStore) OverwriteArchiveBatchVersion(tableName string, shard, batchID int, version uint32, seqNum uint32, batchSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// AddArchiveBatchVersion adds a new version to archive batch.
func (dm *diskMetaStore) AddArchiveBatchVersion(tableName string, shard, batchID int, version uint32, seqNum uint32, batchSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func (dm *diskMetaStore) writeArchiveBatchVersionWithMode(tableName string, shard, batchID int, version uint32, seqNum uint32, batchSize int, mode int) error {
	_ = "STUB: not implemented"
	return nil
}

func (dm *diskMetaStore) GetArchiveBatches(table string, shard int, batchIDStart, batchIDEnd int32) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we'll skip invalid batchID

// GetArchiveBatchVersion gets the latest version <= given archiving/live cutoff
// all cutoff and batch versions are sorted in file per batch
// sample:
//
//		/root_path/metastore/{$table}/shards/{$shard_id}/batches/{$batch_id}
//	 version,size
//	 1-0,10
//	 2-0,20
//	 2-1,26
//	 4-0,20
//	 5-0,20
//	 5-1,25
//	 5-2,38
//
// if given cutoff 6, returns 5-2,38
// if given cutoff 4, returns 4-0,20
// if given cutoff 0, returns 0-0, 0
func (dm *diskMetaStore) GetArchiveBatchVersion(table string, shard, batchID int, cutoff uint32) (uint32, uint32, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

// do binary search to find the first cutoff that is larger than the specified cutoff

// backward compatible: sequence number may not exist for old version

// this should never happen

// all cutoffs larger than given cutoff

func (dm *diskMetaStore) pushSchemaChange(table *common.Table) { _ = "STUB: not implemented"; return }

func (dm *diskMetaStore) pushShardOwnershipChange(tableName string) {
	_ = "STUB: not implemented"
	return
}

// listTable lists the table
func (dm *diskMetaStore) listTables() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (dm *diskMetaStore) removeTable(tableName string) error { _ = "STUB: not implemented"; return nil }

// close all related enum dict watchers
// make sure all producer have done producing and detach

// drain done channels for related enum watchers
// to make sure all previous changes are done

func (dm *diskMetaStore) addColumn(table *common.Table, column common.Column, appendToArchivingSortOrder bool) error {
	_ = "STUB: not implemented"
	return nil
}

// if enum column, append a enum case for default value

func (dm *diskMetaStore) updateColumn(table *common.Table, columnName string, config common.ColumnConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// continue looking since there could be reused column name
// with different column id.

func (dm *diskMetaStore) removeColumn(table *common.Table, columnName string) error {
	_ = "STUB: not implemented"
	return nil
}

// continue looking since there could be reused column name
// with different column id

// trying to delete timestamp column from fact table

func (dm *diskMetaStore) getTableDirPath(tableName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getEnumDirPath(tableName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getEnumFilePath(tableName, columnName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getSchemaFilePath(tableName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getShardsDirPath(tableName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getShardDirPath(tableName string, shard int) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getShardVersionFilePath(tableName string, shard int) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getArchiveBatchVersionFilePath(tableName string, shard, batchID int) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getArchiveBatchDirPath(tableName string, shard int) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getRedoLogVersionAndOffsetFilePath(tableName string, shard int) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *diskMetaStore) getSnapshotRedoLogVersionAndOffsetFilePath(tableName string, shard int) string {
	_ = "STUB: not implemented"
	return ""
}

// Get file path which stores the ingestion commit offset, mainly used for kafka or other streaming based ingestion
func (dm *diskMetaStore) getIngestionCommitOffsetFilePath(tableName string, shard int) string {
	_ = "STUB: not implemented"
	return ""
}

// Get file path which stores the ingestion checkpoint offset, mainly used fo kafka or other streaming based ingestion
func (dm *diskMetaStore) getIngestionCheckpointOffsetFilePath(tableName string, shard int) string {
	_ = "STUB: not implemented"
	return ""
}

// readEnumFile reads the enum cases from file.
func (dm *diskMetaStore) readEnumFile(tableName, columnName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// writeEnumFile append enum cases to existing file
func (dm *diskMetaStore) writeEnumFile(tableName, columnName string, enumCases []string) error {
	_ = "STUB: not implemented"
	return nil
}

// readSchemaFile reads the schema file for given table.
func (dm *diskMetaStore) readSchemaFile(tableName string) (*common.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// writeSchemaFile reads the schema file for given table.
func (dm *diskMetaStore) writeSchemaFile(table *common.Table) error {
	_ = "STUB: not implemented"
	return nil
}

// readVersion reads the version from a given version file.
func (dm *diskMetaStore) readVersion(file string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// readRedoLogFileAndOffset reads the redo log file and offset from the file.
func (dm *diskMetaStore) readRedoLogFileAndOffset(filePath string) (int64, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// readSnapshotRedoLogFileAndOffset reads the redo log file and offset from the file.
func (dm *diskMetaStore) readSnapshotRedoLogFileAndOffset(filePath string) (int64, uint32, int32, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

// writeArchivingCutoff writes the version to a given file.
func (dm *diskMetaStore) writeArchivingCutoff(file string, version uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// writeRedoLogVersionAndOffset writes redolog&offset to a given file.
func (dm *diskMetaStore) writeRedoLogVersionAndOffset(file string, redoLogFile int64, upsertBatchOffset uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// writeSnapshotRedoLogVersionAndOffset writes redolog&offset and last record position to a given file.
func (dm *diskMetaStore) writeSnapshotRedoLogVersionAndOffset(file string, redoLogFile int64, upsertBatchOffset uint32, lastReadBatchID int32, lastReadBatchOffset uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// closeEnumWatcher try to close enum watcher and delete enum file
func (dm *diskMetaStore) removeEnumColumn(tableName, columnName string) {
	_ = "STUB: not implemented"
	return
}

// drain up done channel for enum column
// to make sure previous changes are processed

//TODO: log an error and alert.

// tableExists checks whether table exists,
// return ErrTableDoesNotExist.
func (dm *diskMetaStore) tableExists(tableName string) error { _ = "STUB: not implemented"; return nil }

func (dm *diskMetaStore) getColumnByName(tableName, columnName string) (*common.Column, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// continue since column name can be reused
// with different id

// enumColumnExists checks whether column exists and it is a enum column,
// return ErrTableDoesNotExist, ErrColumnDoesNotExist, ErrNotEnumColumn.
func (dm *diskMetaStore) enumColumnExists(tableName string, columnName string) error {
	_ = "STUB: not implemented"
	return nil
}

// NewDiskMetaStore creates a new disk based metastore
func NewDiskMetaStore(basePath string) (common.MetaStore, error) {
	_ = "STUB: not implemented"
	return *new(common.MetaStore), nil
}
