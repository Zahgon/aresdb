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

package handlers

import (
	"net/http"

	"github.com/uber/aresdb/utils"

	"github.com/gorilla/mux"
	mutatorCom "github.com/uber/aresdb/controller/mutators/common"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	forceKey = "force"
	logMsg   = "Schema request"
)

// SchemaHandlerParams defineds parameters needed to initialize schema handler
type SchemaHandlerParams struct {
	fx.In

	TableSchemaMutator mutatorCom.TableSchemaMutator
	EnumMutator        mutatorCom.EnumMutator
	Logger             *zap.SugaredLogger
}

// SchemaHandler serves schema requests
type SchemaHandler struct {
	tableSchemaMutator mutatorCom.TableSchemaMutator
	enumMutator        mutatorCom.EnumMutator
	logger             *zap.SugaredLogger
}

// NewSchemaHandler creates a new schema handler
func NewSchemaHandler(p SchemaHandlerParams) SchemaHandler {
	_ = "STUB: not implemented"
	return *new(SchemaHandler)
}

// Register adds paths to router
func (h SchemaHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

// AddTable  swagger:route POST /schema/{namespace}/tables addTable
// adds a new table
//
// Consumes:
//   - application/json
func (h SchemaHandler) AddTable(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetTable swagger:route GET /schema/{namespace}/tables/{table} getTable
// returns table schema of given table name
func (h SchemaHandler) GetTable(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetTables swagger:route GET /schema/{namespace}/tables getTables
// returns schema of all tables
func (h SchemaHandler) GetTables(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO: call membership mutator to update instance status
//if getTablesRequest.InstanceName != "" {
//	...
//}

// DeleteTable swagger:route DELETE /schema/{namespace}/tables/{table} deleteTable
// deletes an existing table
func (h SchemaHandler) DeleteTable(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// UpdateTable swagger:route PUT /schema/{namespace}/tables/{table} updateTable
// updates existing table's schema and config
//
// Consumes:
//   - application/json
func (h SchemaHandler) UpdateTable(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetHash swagger:route GET /schema/{namespace}/hash getSchemaHash
// returns hash of all table schemas in a namespace, hash change means there's a change in any of the schemas
func (h SchemaHandler) GetHash(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO: call membership mutator to update instance status
//if getTablesRequest.InstanceName != "" {
//	...
//}

// ExtendEnumCases swagger:route POST /schema/{namespace}/tables/{table}/columns/{column}/enum-cases extendEnumCases
// returns enum ids for given enum cases
func (h SchemaHandler) ExtendEnumCases(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetEnumCases swagger:route GET /schema/{namespace}/tables/{table}/columns/{column}/enum-cases getEnumCases
// returns all enum cases for given table column
func (h SchemaHandler) GetEnumCases(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
