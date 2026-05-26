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
	"github.com/uber/aresdb/controller/mutators/common"

	"github.com/m3db/m3/src/cluster/kv"
)

// NewIngestionAssignmentMutator creates new IngestionAssignmentMutator
func NewIngestionAssignmentMutator(etcdStore kv.TxnStore) common.IngestionAssignmentMutator {
	_ = "STUB: not implemented"
	return *new(common.IngestionAssignmentMutator)
}

type ingestionAssignmentMutatorImpl struct {
	etcdStore kv.TxnStore
}

// GetIngestionAssignment gets IngestionAssignment config by name
func (j *ingestionAssignmentMutatorImpl) GetIngestionAssignment(namespace, name string) (ingestionAssignment models.IngestionAssignment, err error) {
	_ = "STUB: not implemented"
	return *new(models.IngestionAssignment), nil
}

// GetIngestionAssignments returns all IngestionAssignments config
func (j *ingestionAssignmentMutatorImpl) GetIngestionAssignments(namespace string) (ingestionAssignments []models.IngestionAssignment, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteIngestionAssignment deletes a IngestionAssignment
func (j *ingestionAssignmentMutatorImpl) DeleteIngestionAssignment(namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateIngestionAssignment updates IngestionAssignment config
func (j *ingestionAssignmentMutatorImpl) UpdateIngestionAssignment(namespace string, ingestionAssignment models.IngestionAssignment) error {
	_ = "STUB: not implemented"
	return nil
}

// AddIngestionAssignment adds a new IngestionAssignment
func (j *ingestionAssignmentMutatorImpl) AddIngestionAssignment(namespace string, ingestionAssignment models.IngestionAssignment) error {
	_ = "STUB: not implemented"
	return nil
}

// GetHash returns hash that will be different if ingestionAssignment for subscriber changed
func (j *ingestionAssignmentMutatorImpl) GetHash(namespace, subscriber string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
