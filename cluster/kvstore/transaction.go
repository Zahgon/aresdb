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
package kvstore

import (
	"github.com/golang/protobuf/proto"
	"github.com/m3db/m3/src/cluster/kv"
)

// Transaction defines a Transaction
type Transaction struct {
	keys     []string
	versions []int
	values   []proto.Message
}

// NewTransaction creates a new transaction
func NewTransaction() *Transaction { _ = "STUB: not implemented"; return nil }

// AddKeyValue adds a tuple of (key, version, value) to the transaction
func (t *Transaction) AddKeyValue(key string, version int, value proto.Message) *Transaction {
	_ = "STUB: not implemented"
	return nil
}

// WriteTo writes the transaction to the transaction store
func (t *Transaction) WriteTo(store kv.TxnStore) error { _ = "STUB: not implemented"; return nil }
