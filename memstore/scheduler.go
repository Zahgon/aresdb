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
	"time"

	"github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/utils"
)

const (
	// interval for scheduler
	schedulerInterval = time.Minute
)

// jobBundle binds a result channel together with the job.
// Listening on the result channel will block until job finishes.
// Nil error indicates job runs successfully.
type jobBundle struct {
	Job
	resChan chan error
}

// Scheduler is for scheduling archiving jobs (and later backfill jobs) for table shards
// in memStore. It scans through all tables and shards to generate list of eligible jobs
// to run.
type Scheduler interface {
	Start()
	Stop()
	SubmitJob(job Job) (error, chan error)
	DeleteTable(table string, isFactTable bool)
	GetJobDetails(jobType common.JobType) interface{}
	NewBackfillJob(tableName string, shardID int) Job
	NewArchivingJob(tableName string, shardID int, cutoff uint32) Job
	NewSnapshotJob(tableName string, shardID int) Job
	NewPurgeJob(tableName string, shardID int, batchIDStart int, batchIDEnd int) Job
	EnableJobType(jobType common.JobType, enable bool)
	IsJobTypeEnabled(jobType common.JobType) bool
	utils.RWLocker
}

// newScheduler returns a new Scheduler.
func newScheduler(m *memStoreImpl) *schedulerImpl { _ = "STUB: not implemented"; return nil }

// schedulerImpl is the implementation of Scheduler interface.
type schedulerImpl struct {
	// Protecting JobRunning.
	sync.RWMutex
	// For accessing meta data like archiving delay and interval
	memStore *memStoreImpl
	// Stop main scheduler loop.
	schedulerStopChan chan struct{}
	// Channel for executing job.
	jobBundleChan chan jobBundle
	// Stop executor loop.
	executorStopChan chan struct{}
	jobManagers      map[common.JobType]jobManager
	jobEnableFlags   map[common.JobType]bool
	archivingStarted bool
}

func (scheduler *schedulerImpl) EnableJobType(jobType common.JobType, enable bool) {
	_ = "STUB: not implemented"
	return
}

func (scheduler *schedulerImpl) IsJobTypeEnabled(jobType common.JobType) bool {
	_ = "STUB: not implemented"
	return false
}

func (scheduler *schedulerImpl) reportJob(key string, mutator jobDetailMutator) {
	_ = "STUB: not implemented"
	return
}

// getIdentifier returns a unique identifier from table, shard and job type.
func getIdentifier(tableName string, shardID int, jobType common.JobType) string {
	_ = "STUB: not implemented"
	return ""
}

// GetJobDetails returns corresponding job details for given job type.
func (scheduler *schedulerImpl) GetJobDetails(jobType common.JobType) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// DeleteTable deletes the job details of a table given its name and whether it's a fact table.
func (scheduler *schedulerImpl) DeleteTable(table string, isFactTable bool) {
	_ = "STUB: not implemented"
	return
}

// GetJobManager retrieve the JobManager according to job type
func (scheduler *schedulerImpl) GetJobManager(jobType common.JobType) jobManager {
	_ = "STUB: not implemented"
	return *new(jobManager)
}

// NewArchivingJob returns a new ArchivingJob.
func (scheduler *schedulerImpl) NewArchivingJob(tableName string, shardID int, cutoff uint32) Job {
	_ = "STUB: not implemented"
	return *new(Job)
}

// NewBackfillJob returns a new BackfillJob.
func (scheduler *schedulerImpl) NewBackfillJob(tableName string, shardID int) Job {
	_ = "STUB: not implemented"
	return *new(Job)
}

// NewSnapshotJob returns a new SnapshotJob.
func (scheduler *schedulerImpl) NewSnapshotJob(tableName string, shardID int) Job {
	_ = "STUB: not implemented"
	return *new(Job)
}

// NewPurgeJob returns a new PurgeJob
func (scheduler *schedulerImpl) NewPurgeJob(tableName string, shardID, batchIDStart, batchIDEnd int) Job {
	_ = "STUB: not implemented"
	return *new(Job)
}

// Start starts the scheduler. It creates a new time.Timer every time to wait
// at least schedulerInterval time instead of running at every tick so that we
// will skip the tick if a single round takes more than one minute. This prevents
// accessing memStore (and lock) too many times during a short period.
func (scheduler *schedulerImpl) Start() { _ = "STUB: not implemented"; return }

// Scheduler loop.

// Since we already receive the event from channel,
// there is no need to stop it and we can directly reset the timer.

// It will block on waiting for executor to stop.

// Executor loop.

func (scheduler *schedulerImpl) executeJob(jb *jobBundle) { _ = "STUB: not implemented"; return }

// Set job status according to the result.

// This is a non-blocking channel sending.

// Stop stops the scheduler.
func (scheduler *schedulerImpl) Stop() { _ = "STUB: not implemented"; return }

// SubmitJob will submit a job to executor and block until it starts.
// Job submitter can decide whether to wait for job to finish and get
// the result.
func (scheduler *schedulerImpl) SubmitJob(job Job) (error, chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this check is to block request from debug handler

// run runs at every tick. It first generates a list of jobs to run based on current condition,
// then it runs every job sequentially in the same process.
func (scheduler *schedulerImpl) run() { _ = "STUB: not implemented"; return }

// Waiting for job to finish.

// Job defines the common interface for BackfillJob, ArchivingJob and SnapshotJob
type Job interface {
	JobType() common.JobType
	Run() error
	GetIdentifier() string
	String() string
}
