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

package etcd

import (
	"github.com/uber/aresdb/cluster/kvstore"
	"github.com/uber/aresdb/controller/models"
	"github.com/uber/aresdb/controller/mutators/common"
)

// NewMembershipMutator creates new MembershipMutator
func NewMembershipMutator(etcdClient *kvstore.EtcdClient) common.MembershipMutator {
	_ = "STUB: not implemented"
	return *new(common.MembershipMutator)
}

type membershipMutatorImpl struct {
	etcdClient *kvstore.EtcdClient
}

func (mm membershipMutatorImpl) Join(namespace string, instance models.Instance) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm membershipMutatorImpl) Leave(namespace, instanceName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm membershipMutatorImpl) GetInstance(namespace, instanceName string) (instance models.Instance, err error) {
	_ = "STUB: not implemented"
	return *new(models.Instance), nil
}

func (mm membershipMutatorImpl) GetInstances(namespace string) ([]models.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetHash returns hash that will be different if any instance changed
func (mm membershipMutatorImpl) GetHash(namespace string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
