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

package broker

import (
	"sync"

	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/metastore/common"
)

// BrokerSchemaMutator implements metastore.TableSchemaMutator
// and memstore.TableSchemaReader, and memstore.EnumUpdater
type BrokerSchemaMutator struct {
	sync.RWMutex

	tables map[string]*memCom.TableSchema
}

func NewBrokerSchemaMutator() *BrokerSchemaMutator { _ = "STUB: not implemented"; return nil }

// ====  metastore/common.TableSchemaMutator ====
func (b *BrokerSchemaMutator) ListTables() (tables []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BrokerSchemaMutator) GetTable(name string) (table *common.Table, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BrokerSchemaMutator) CreateTable(table *common.Table) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *BrokerSchemaMutator) DeleteTable(name string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *BrokerSchemaMutator) UpdateTableConfig(table string, config common.TableConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *BrokerSchemaMutator) UpdateTable(table common.Table) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *BrokerSchemaMutator) AddColumn(table string, column common.Column, appendToArchivingSortOrder bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *BrokerSchemaMutator) UpdateColumn(table string, column string, config common.ColumnConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *BrokerSchemaMutator) DeleteColumn(table string, column string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ====  memstore/common.TableSchemaReader ====
func (b *BrokerSchemaMutator) GetSchema(table string) (*memCom.TableSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BrokerSchemaMutator) GetSchemas() map[string]*memCom.TableSchema {
	_ = "STUB: not implemented"
	return nil
}

func (b *BrokerSchemaMutator) UpdateEnum(table, column string, enumList []string) error {
	_ = "STUB: not implemented"
	return nil
}
