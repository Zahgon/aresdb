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
	"sync"

	"github.com/uber/aresdb/controller/mutators/common"

	"github.com/m3db/m3/src/cluster/kv"
)

const maxEnumCasePerNode = 1000

type enumCache struct {
	enumCases     map[string]int
	currentNodeID int
}

type enumMutator struct {
	sync.RWMutex

	schemaMutator common.TableSchemaMutator
	// namesapce to columnID to enum cases set
	// key {namespace}/{table}/{incarnation}/{columnID}
	enumCacheMap map[string]enumCache
	txnStore     kv.TxnStore
}

// NewEnumMutator creates EnumMutator
func NewEnumMutator(store kv.TxnStore, schemaMutator common.TableSchemaMutator) common.EnumMutator {
	_ = "STUB: not implemented"
	return *new(common.EnumMutator)
}

func (e *enumMutator) ExtendEnumCases(namespace, tableName, columnName string, enumCases []string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetch enum case from etcd and update cache

// check enum cache for existing enums

// fetch enum cases from etcd and update cache

// missingID is ordered the same as newEnumCases
// which is the same order in enumIDs

func getCacheKey(namespace, tableName string, incarnation, columnID int) string {
	_ = "STUB: not implemented"
	return ""
}

func getEnumID(nodeID int, innerID int) int { _ = "STUB: not implemented"; return 0 }

func (e *enumMutator) extendEnumCase(namespace, tableName string, incarnation, columnID int, fromEnumNodeID int, newEnumCases []string, enumIDUpperBound int) ([]int, error) {
	_ = "STUB: not implemented"
	// track result resolvedEnumIDs
	return nil, nil
}

// newEnumCaseDict records the resolved resolvedEnumIDs for newEnumCases

// fetch current enum node list

// track enumNode's enum cases

// track enumNode's version

// track enumNode's key in etcd

// it is guaranteed to have at least one enum node

// fetch all enum cases from fromEnumNode to lastEnumNode
// and update newEnumCaseDict

// variables needed for write transaction to etcd

// track the last enum node

// transaction

// only when there are new enum cases not in cache, we need to write to etcd for update

// once last node is full, append the last finished node

// create new node

// advance lastEnumNodeID

// create last node transaction

// check max enum id before creating new enum ids

// append nodeList to the transaction

func (e *enumMutator) fetchEnumCases(key string) ([]string, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (e *enumMutator) updateCache(cacheKey string, nodeID int, dict map[string]int) {
	_ = "STUB: not implemented"
	return
}

// update currentNodeID for cache

// merge newEnumCaseDict to cache

func (e *enumMutator) GetEnumCases(namespace, tableName, columnName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
