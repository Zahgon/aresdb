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

type subscriberMutator struct {
	etcdClient *kvstore.EtcdClient
}

// NewSubscriberMutator creates new subscriber mutator based on etcd
func NewSubscriberMutator(etcdClient *kvstore.EtcdClient) common.SubscriberMutator {
	_ = "STUB: not implemented"
	return *new(common.SubscriberMutator)
}

// GetSubscriber returns a subscriber
func (s *subscriberMutator) GetSubscriber(namespace, subscriberName string) (subscriber models.Subscriber, err error) {
	_ = "STUB: not implemented"
	return *new(models.Subscriber), nil
}

// GetSubscribers returns a list of subscribers
func (s *subscriberMutator) GetSubscribers(namespace string) (subscribers []models.Subscriber, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetHash returns hash of all subscribers
func (s *subscriberMutator) GetHash(namespace string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
