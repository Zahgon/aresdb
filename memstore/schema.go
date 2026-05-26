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
	metaCom "github.com/uber/aresdb/metastore/common"
)

// FetchSchema fetches schema from metaStore and updates in-memory copy of table schema,
// and set up watch channels for metaStore schema changes, used for bootstrapping mem store.
func (m *memStoreImpl) FetchSchema() error { _ = "STUB: not implemented"; return nil }

// watch table addition/modification

// watch table deletion

// watch enum cases appending

func (m *memStoreImpl) fetchTable(tableName string) error { _ = "STUB: not implemented"; return nil }

// watch enumCases will setup watch channels for each enum column.
func (m *memStoreImpl) watchEnumCases(tableName, columnName string, startCase int) error {
	_ = "STUB: not implemented"
	return nil
}

// handleTableListChange handles table deletion events from metaStore.
func (m *memStoreImpl) handleTableListChange(tableListChangeEvents <-chan []string, done chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

func (m *memStoreImpl) applyTableList(newTableList []string) { _ = "STUB: not implemented"; return }

// detach shards and schema from map
// to prevent new usage

// only one table deletion at a time

// handleTableSchemaChange handles table schema change event from metaStore including new table schema.
func (m *memStoreImpl) handleTableSchemaChange(tableSchemaChangeEvents <-chan *metaCom.Table, done chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

func (m *memStoreImpl) applyTableSchema(newTable *metaCom.Table) { _ = "STUB: not implemented"; return }

// default start watching from first enumCase

// new table

// default value is already appended, start watching from 1

// new deletions only

// default value is already appended, start watching from 1

// always set default value after enum map creation

// preloading will be triggered if
// 1. this is a new column and PreloadingDays > 0
// 2. this is a old column and PreloadingDays > oldPreloadingDays

// May block for extended amount of time during archiving

// handleEnumDictChange handles enum dict change event from metaStore for specific table and column.
func (m *memStoreImpl) handleEnumDictChange(tableName, columnName string, enumDictChangeEvents <-chan string, done chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

func (m *memStoreImpl) applyEnumCase(tableName, columnName string, newEnumCase string) {
	_ = "STUB: not implemented"
	return
}
