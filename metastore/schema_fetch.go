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

package metastore

import (
	controllerCli "github.com/uber/aresdb/controller/client"
	controllerMutatorCom "github.com/uber/aresdb/controller/mutators/common"
	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/metastore/common"
)

// SchemaFetchJob is a job that periodically pings ares-controller and updates table schemas if applicable
type SchemaFetchJob struct {
	clusterName       string
	hash              string
	intervalInSeconds int
	schemaMutator     common.TableSchemaMutator
	enumUpdater       memCom.EnumUpdater
	schemaValidator   TableSchemaValidator
	controllerClient  controllerCli.ControllerClient
	enumMutator       controllerMutatorCom.EnumMutator
	stopChan          chan struct{}
}

// NewSchemaFetchJob creates a new SchemaFetchJob
func NewSchemaFetchJob(intervalInSeconds int, schemaMutator common.TableSchemaMutator, enumUpdater memCom.EnumUpdater, schemaValidator TableSchemaValidator, controllerClient controllerCli.ControllerClient, enumMutator controllerMutatorCom.EnumMutator, clusterName, initialHash string) *SchemaFetchJob {
	_ = "STUB: not implemented"
	return nil
}

// Run starts the scheduling
func (j *SchemaFetchJob) Run() { _ = "STUB: not implemented"; return }

// Stop stops the scheduling
func (j *SchemaFetchJob) Stop() { _ = "STUB: not implemented"; return }

func (j *SchemaFetchJob) FetchSchema() { _ = "STUB: not implemented"; return }

// errors already reported, just return without updating hash

func (j *SchemaFetchJob) applySchemaChange(tables []common.Table) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// found new table

// found new table incarnation, delete previous table and data
// then create new table

// found table update

// found table deletion

// FetchEnum updates all enums
func (j *SchemaFetchJob) FetchEnum() { _ = "STUB: not implemented"; return }

func reportError(err error, isSchemaError bool, extraInfo string) {
	_ = "STUB: not implemented"
	return
}
