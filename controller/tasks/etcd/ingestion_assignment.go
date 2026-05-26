//	Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package etcd

import (
	"github.com/uber/aresdb/controller/models"
	mutators "github.com/uber/aresdb/controller/mutators/common"
	"github.com/uber/aresdb/controller/tasks/common"

	"github.com/m3db/m3/src/cluster/services"
	"github.com/uber-go/tally"
	"go.uber.org/zap"
)

const (
	assignmentChangeMetricName   = "ingestion_assignment_changed"
	assignmentErrorMetricName    = "ingestion_assignment_error"
	assignmentSuccessMetricName  = "ingestion_assignment_success"
	ingestionAssignmentConfigKey = "ingestionAssignmentTask"
	taskTagValue                 = "ingestionAssignmentTask"
)

type ingestionAssignmentTaskConfig struct {
	IntervalInSeconds int `yaml:"intervalInSeconds"`
}

// ingestionAssignmentTask calculates ingestion jobs assignment to subscriber instances
// given current state of the cluster
type ingestionAssignmentTask struct {
	zone        string
	environment string

	intervalSeconds int
	logger          *zap.SugaredLogger
	scope           tally.Scope
	// closing stopChan will:
	// stop ongoing leader election, which will clean up leader election states
	// stop ongoing task assignment calculation loop
	stopChan chan struct{}

	namespaceMutator   mutators.NamespaceMutator
	jobMutator         mutators.JobMutator
	schemaMutator      mutators.TableSchemaMutator
	subscriberMutator  mutators.SubscriberMutator
	assignmentsMutator mutators.IngestionAssignmentMutator

	etcdServices   services.Services
	leaderElection LeaderElector
	configHashes   map[string]configHash
}

type jobSubscriberState struct {
	namespace   string
	subscribers []models.Subscriber
	jobs        []models.JobConfig
}

type configHash struct {
	jobsHash       string
	schemaHash     string
	subscriberHash string
}

// NewIngestionAssignmentTask creates a new instance of ingestionAssignmentTask
func NewIngestionAssignmentTask(p common.IngestionAssignmentTaskParams) common.Task {
	_ = "STUB: not implemented"
	return *new(common.Task)
}

// Run starts the ingestionAssignmentTask
func (ia *ingestionAssignmentTask) Run() { _ = "STUB: not implemented"; return }

// wait random interval to avoid herd effect electing for leader on cluster reboot

// Done stops the task
func (ia *ingestionAssignmentTask) Done() { _ = "STUB: not implemented"; return }

func (ia *ingestionAssignmentTask) startIngestionAssignment(hostName string) {
	_ = "STUB: not implemented"

	// waiting for new election status change
	return
}

func (ia *ingestionAssignmentTask) isLeader() bool { _ = "STUB: not implemented"; return false }

func (ia *ingestionAssignmentTask) checkConfigHashes(namespace string) (hashes configHash, err error) {
	_ = "STUB: not implemented"
	return *new(configHash), nil
}

func (ia *ingestionAssignmentTask) readCurrentState(namespace string) (jobSubscriberState, error) {
	_ = "STUB: not implemented"
	return *new(jobSubscriberState), nil
}

func (ia *ingestionAssignmentTask) recalculateForNamespace(ns string) {
	_ = "STUB: not implemented"
	return
}

func (ia *ingestionAssignmentTask) tryRecalculateAllNamespaces() (errs int) {
	_ = "STUB: not implemented"
	return 0
}

func (ia *ingestionAssignmentTask) processIngestionAssignment(state jobSubscriberState) (changes, errs int) {
	_ = "STUB: not implemented"
	// build consistent hash ring where a ring node is a subscriber
	// and resource key is the kafka topic name.
	// this will guarantee minimum change for a topic's ingestion assignment
	// when subscribers join/leave the cluster
	return 0, 0
}

// TODO: take subscriber instance capacity into consideration when assigning jobs

// calculate starting node of task assignment base on kafka topic name

// update existing assignment

// mark not deleted

// new assignment

// add dummy assignment

func (ia *ingestionAssignmentTask) reportError(err error, errCount *int) {
	_ = "STUB: not implemented"
	return
}
