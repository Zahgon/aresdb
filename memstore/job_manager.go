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

	"github.com/uber/aresdb/memstore/common"
)

// JobManager is responsible for generating new jobs to run and manages job related stats.
type jobManager interface {
	generateJobs() []Job
	getJobDetails() interface{}
	deleteTable(table string)
	// mutator is guaranteed to be a functor by caller(scheduler).
	reportJobDetail(key string, mutator jobDetailMutator)
}

type archiveJobManager struct {
	sync.RWMutex
	// archiveJobDetails for different tables, shard. Key is {tableName}|{shardID}|archiving,
	// value is the job details.
	jobDetails map[string]*ArchiveJobDetail
	// For accessing meta data like archiving delay and interval
	memStore  *memStoreImpl
	scheduler *schedulerImpl
}

// newArchiveJobManager creates a new jobManager to manage archive jobs.
func newArchiveJobManager(scheduler *schedulerImpl) jobManager {
	_ = "STUB: not implemented"
	return *new(jobManager)
}

// generateJobs iterates each table shard from memStore and prepare list of archive jobs
// to run. A job should start to run only when newCutoff - cutoff > interval, where
// newCutoff = now - delay.
func (m *archiveJobManager) generateJobs() []Job { _ = "STUB: not implemented"; return nil }

func (m *archiveJobManager) getJobDetails() interface{} { _ = "STUB: not implemented"; return nil }

func (m *archiveJobManager) reportJobDetail(key string, jobMutator jobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

func (m *archiveJobManager) reportArchiveJobDetail(key string, jobMutator ArchiveJobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

// caller needs to hold the write lock.
func (m *archiveJobManager) getJobDetail(key string) *ArchiveJobDetail {
	_ = "STUB: not implemented"
	return nil
}

// deleteTable deletes metadata for the table in archiveJobManager.
func (m *archiveJobManager) deleteTable(table string) { _ = "STUB: not implemented"; return }

// ArchivingJob defines the structure that an archiving job needs.
type ArchivingJob struct {
	// table to archive
	tableName string
	// shard to archive
	shardID int
	// new cut off
	cutoff uint32
	// for calling archiving function in memStore
	memStore MemStore
	// for reporting job detail changes
	reporter ArchiveJobDetailReporter
}

// Run starts the archiving process and wait for it to finish.
func (job *ArchivingJob) Run() error { _ = "STUB: not implemented"; return nil }

// GetIdentifier returns a unique identifier of this job.
func (job *ArchivingJob) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

// String gives meaningful string representation for this job
func (job *ArchivingJob) String() string { _ = "STUB: not implemented"; return "" }

// JobType return job type
func (job *ArchivingJob) JobType() common.JobType {
	_ = "STUB: not implemented"
	return *new(common.JobType)
}

type backfillJobManager struct {
	sync.RWMutex
	// backfillJobDetails for different tables, shard. Key is {tableName}|{shardID}|backfill,
	jobDetails map[string]*BackfillJobDetail
	// For accessing meta data like archiving delay and interval
	memStore  *memStoreImpl
	scheduler *schedulerImpl
}

// newBackfillJobManager creates a new jobManager to manage backfill jobs.
func newBackfillJobManager(scheduler *schedulerImpl) jobManager {
	_ = "STUB: not implemented"
	return *new(jobManager)
}

// generateJobs iterates each table shard from memStore and prepare list of backfill jobs
// to run.
func (m *backfillJobManager) generateJobs() []Job { _ = "STUB: not implemented"; return nil }

// size based strategy

// timer based strategy

// the job detail has just been initialized.

// enqueue backfill job

func (m *backfillJobManager) getJobDetails() interface{} { _ = "STUB: not implemented"; return nil }

func (m *backfillJobManager) reportJobDetail(key string, jobMutator jobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

func (m *backfillJobManager) reportBackfillJobDetail(key string, jobMutator BackfillJobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

// caller needs to hold the write lock.
func (m *backfillJobManager) getJobDetail(key string) *BackfillJobDetail {
	_ = "STUB: not implemented"
	return nil
}

// deleteTable deletes metadata for the table in backfillJobManager.
func (m *backfillJobManager) deleteTable(table string) { _ = "STUB: not implemented"; return }

// BackfillJob defines the structure that a backfill job needs.
type BackfillJob struct {
	// table to backfill
	tableName string
	// shard to backfill
	shardID int
	// for calling backfill function in memStore
	memStore MemStore
	// for reporting JobDetail changes.
	reporter BackfillJobDetailReporter
}

// Run starts the backfill process and wait for it to finish.
func (job *BackfillJob) Run() error { _ = "STUB: not implemented"; return nil }

// GetIdentifier returns a unique identifier of this job.
func (job *BackfillJob) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

// String gives meaningful string representation for this job
func (job *BackfillJob) String() string { _ = "STUB: not implemented"; return "" }

// JobType return job type
func (job *BackfillJob) JobType() common.JobType {
	_ = "STUB: not implemented"
	return *new(common.JobType)
}

type snapshotJobManager struct {
	sync.RWMutex
	// snapshotJobDetails for different tables, shard. Key is {tableName}|{shardID}|snapshot,
	jobDetails map[string]*SnapshotJobDetail
	// For accessing meta data like archiving delay and interval
	memStore  *memStoreImpl
	scheduler *schedulerImpl
}

// newSnapshotJobManager creates a new jobManager to manage snapshot jobs.
func newSnapshotJobManager(scheduler *schedulerImpl) jobManager {
	_ = "STUB: not implemented"
	return *new(jobManager)
}

// generateJobs iterates each table shard from memStore and prepare list of snapshot jobs
// to run.
func (m *snapshotJobManager) generateJobs() []Job { _ = "STUB: not implemented"; return nil }

// the job detail has just been initialized.

func (m *snapshotJobManager) getJobDetails() interface{} { _ = "STUB: not implemented"; return nil }

// deleteTable deletes metadata for the table in snapshotJobManager.
func (m *snapshotJobManager) deleteTable(table string) { _ = "STUB: not implemented"; return }

func (m *snapshotJobManager) reportJobDetail(key string, jobMutator jobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

func (m *snapshotJobManager) reportSnapshotJobDetail(key string, jobMutator SnapshotJobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

// caller needs to hold the write lock.
func (m *snapshotJobManager) getJobDetail(key string) *SnapshotJobDetail {
	_ = "STUB: not implemented"
	return nil
}

// SnapshotJob defines the structure that a snapshot job needs.
type SnapshotJob struct {
	// table to snapshot
	tableName string
	// shard to snapshot
	shardID int
	// for calling snapshot function in memStore
	memStore MemStore
	// for reporting snapshot JobDetail changes
	reporter SnapshotJobDetailReporter
}

// Run starts the snapshot process and wait for it to finish.
func (job *SnapshotJob) Run() error { _ = "STUB: not implemented"; return nil }

// GetIdentifier returns a unique identifier of this job.
func (job *SnapshotJob) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

// String gives meaningful string representation for this job
func (job *SnapshotJob) String() string { _ = "STUB: not implemented"; return "" }

// JobType return job type
func (job *SnapshotJob) JobType() common.JobType {
	_ = "STUB: not implemented"
	return *new(common.JobType)
}

type purgeJobManager struct {
	sync.RWMutex
	// purge job details for different tables, shard. Key is {tableName}|{shardID}|purge,
	jobDetails map[string]*PurgeJobDetail
	memStore   *memStoreImpl
	scheduler  *schedulerImpl
}

// newPurgeJobManager creates a new jobManager to manage purge jobs.
func newPurgeJobManager(scheduler *schedulerImpl) jobManager {
	_ = "STUB: not implemented"
	return *new(jobManager)
}

// generateJobs iterates each table shard from memStore and prepare list of purge jobs
// to run.
func (m *purgeJobManager) generateJobs() []Job { _ = "STUB: not implemented"; return nil }

func (m *purgeJobManager) getJobDetails() interface{} { _ = "STUB: not implemented"; return nil }

func (m *purgeJobManager) getJobDetail(key string) *PurgeJobDetail {
	_ = "STUB: not implemented"
	return nil
}

func (m *purgeJobManager) reportJobDetail(key string, jobMutator jobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

// deleteTable deletes metadata for the table in purgeJobManager.
func (m *purgeJobManager) deleteTable(table string) { _ = "STUB: not implemented"; return }

func (m *purgeJobManager) reportPurgeJobDetail(key string, jobMutator PurgeJobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

// PurgeJob defines the structure that a purge job needs.
type PurgeJob struct {
	tableName string
	shardID   int
	// max batch id to purge
	batchIDStart int
	batchIDEnd   int
	memStore     MemStore
	reporter     PurgeJobDetailReporter
}

// Run starts the purge process and wait for it to finish.
func (job *PurgeJob) Run() error { _ = "STUB: not implemented"; return nil }

// GetIdentifier returns a unique identifier of this job.
func (job *PurgeJob) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

// String gives meaningful string representation for this job
func (job *PurgeJob) String() string { _ = "STUB: not implemented"; return "" }

// JobType return job type
func (job *PurgeJob) JobType() common.JobType {
	_ = "STUB: not implemented"
	return *new(common.JobType)
}
