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
	"github.com/uber/aresdb/controller/mutators/common"
)

// NewNamespaceMutator returns a new NamespaceMutator
func NewNamespaceMutator(etcdStore kv.TxnStore) common.NamespaceMutator {
	_ = "STUB: not implemented"
	return *new(common.NamespaceMutator)
}

type namespaceMutatorImpl struct {
	txnStore kv.TxnStore
}

func (m *namespaceMutatorImpl) CreateNamespace(namespace string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// pre create schema, job, job assignments list key

func (m *namespaceMutatorImpl) ListNamespaces() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
