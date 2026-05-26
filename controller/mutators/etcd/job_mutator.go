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
	"github.com/m3db/m3/src/cluster/kv"
	pb "github.com/uber/aresdb/controller/generated/proto"
	"github.com/uber/aresdb/controller/models"
	"github.com/uber/aresdb/controller/mutators/common"
	"go.uber.org/zap"
)

// NewJobMutator creates new JobMutator
func NewJobMutator(etcdStore kv.TxnStore, logger *zap.SugaredLogger) common.JobMutator {
	_ = "STUB: not implemented"
	return *new(common.JobMutator)
}

type jobMutatorImpl struct {
	logger    *zap.SugaredLogger
	etcdStore kv.TxnStore
}

// GetJob gets job config by name
func (j *jobMutatorImpl) GetJob(namespace, name string) (job models.JobConfig, err error) {
	_ = "STUB: not implemented"
	return *new(models.JobConfig), nil
}

// always return etcd internal version, version from job payload is ignored

// GetJobs returns all jobs config
func (j *jobMutatorImpl) GetJobs(namespace string) ([]models.JobConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteJob deletes a job
func (j *jobMutatorImpl) DeleteJob(namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// found

// UpdateJob updates job config
func (j *jobMutatorImpl) UpdateJob(namespace string, job models.JobConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AddJob adds a new job
func (j *jobMutatorImpl) AddJob(namespace string, job models.JobConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// GetHash returns hash that will be different if any job changed
func (j *jobMutatorImpl) GetHash(namespace string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (j *jobMutatorImpl) readJob(namespace string, name string) (jobConfig pb.EntityConfig, version int, err error) {
	_ = "STUB: not implemented"
	return *new(pb.EntityConfig), 0, nil
}
