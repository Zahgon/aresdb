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
	"github.com/uber/aresdb/cluster/kvstore"

	"github.com/m3db/m3/src/cluster/kv"
	pb "github.com/uber/aresdb/controller/generated/proto"
	"github.com/uber/aresdb/controller/mutators/common"
	metaCom "github.com/uber/aresdb/metastore/common"
	"go.uber.org/zap"
)

// NewTableSchemaMutator returns a new TableSchemaMutator
func NewTableSchemaMutator(store kv.TxnStore, logger *zap.SugaredLogger) common.TableSchemaMutator {
	_ = "STUB: not implemented"
	return *new(common.TableSchemaMutator)
}

type tableSchemaMutator struct {
	txnStore kv.TxnStore
	logger   *zap.SugaredLogger
}

func (m *tableSchemaMutator) ListTables(namespace string) (tableNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *tableSchemaMutator) GetTable(namespace, name string) (table *metaCom.Table, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *tableSchemaMutator) CreateTable(namespace string, table *metaCom.Table, force bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// table get recreated

func (m *tableSchemaMutator) DeleteTable(namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// found table

// delete enums in background

func (m *tableSchemaMutator) deleteEnum(namespace string, table *metaCom.Table) {
	_ = "STUB: not implemented"
	return
}

func (m *tableSchemaMutator) UpdateTable(namespace string, table metaCom.Table, force bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// always use old table's incarnation for update table operation will not modify incarnation

// merge existing table and column level configs if not specified in the input

// TODO: remove this when upstream supports archive sort columns

// for new columns, pre-create enum nodes

func preCreateEnumNodes(txn *kvstore.Transaction, namespace string, table *metaCom.Table, startColumnID int, endColumnID int) {
	_ = "STUB: not implemented"
	return
}

// default value be first enum case

// enum node list

// first node for enum column

func (m *tableSchemaMutator) GetHash(namespace string) (hash string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *tableSchemaMutator) readSchema(namespace string, name string) (schemaProto pb.EntityConfig, version int, err error) {
	_ = "STUB: not implemented"
	return *new(pb.EntityConfig), 0, nil
}
