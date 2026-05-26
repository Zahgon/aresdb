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

package api

import (
	"net/http"

	memCom "github.com/uber/aresdb/memstore/common"

	metaCom "github.com/uber/aresdb/metastore/common"
	"github.com/uber/aresdb/utils"

	"github.com/gorilla/mux"
)

// SchemaHandler handles schema http requests.
type SchemaHandler struct {
	// all write requests will go to metaStore.
	metaStore    metaCom.MetaStore
	schemaReader memCom.TableSchemaReader
}

// NewSchemaHandler will create a new SchemaHandler with memStore and metaStore.
func NewSchemaHandler(metaStore metaCom.MetaStore, schemaReader memCom.TableSchemaReader) *SchemaHandler {
	_ = "STUB: not implemented"
	return nil
}

// Register registers http handlers.
func (handler *SchemaHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

// RegisterForDebug register handlers for debug port
func (handler *SchemaHandler) RegisterForDebug(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

// ListTables swagger:route GET /schema/tables listTables
// List all table schemas
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// Responses:
//
//	default: errorResponse
//	    200: stringArrayResponse
func (handler *SchemaHandler) ListTables(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetTable swagger:route GET /schema/tables/{table} getTable
// get the table schema for specific table name
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// Responses:
//
//	default: errorResponse
//	    200: getTableResponse
func (handler *SchemaHandler) GetTable(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// AddTable swagger:route POST /schema/tables addTable
// add table to table collections
//
// Consumes:
//   - application/json
//
// Responses:
//
//	default: errorResponse
//	    200: noContentResponse
func (handler *SchemaHandler) AddTable(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// add default table configs first

// UpdateTableConfig swagger:route PUT /schema/tables/{table} updateTableConfig
// update config of the specified table
//
// Consumes:
//   - application/json
//
// Responses:
//
//	default: errorResponse
//	    200: noContentResponse
func (handler *SchemaHandler) UpdateTableConfig(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// DeleteTable swagger:route DELETE /schema/tables/{table} deleteTable
// delete table from metaStore
//
// Responses:
//
//	default: errorResponse
//	    200: noContentResponse
func (handler *SchemaHandler) DeleteTable(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO: need mapping from metaStore error to api error
/// for metaStore error might also be user error

// AddColumn swagger:route POST /schema/tables/{table}/columns addColumn
// add a single column to existing table
//
// Consumes:
//   - application/json
//
// Responses:
//
//	default: errorResponse
//	    200: noContentResponse
func (handler *SchemaHandler) AddColumn(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO: validate column
// might better do in metaStore and here needs to return either user error or server error

// TODO: need mapping from metaStore error to api error
/// for metaStore error might also be user error

// UpdateColumn swagger:route PUT /schema/tables/{table}/columns/{column} updateColumn
// update specified column
//
// Consumes:
//   - application/json
//
// Responses:
//
//	default: errorResponse
//	    200: noContentResponse
func (handler *SchemaHandler) UpdateColumn(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO: need mapping from metaStore error to api error
// for metaStore error might also be user error

// DeleteColumn swagger:route DELETE /schema/tables/{table}/columns/{column} deleteColumn
// delete columns from existing table
//
// Responses:
//
//	default: errorResponse
//	    200: noContentResponse
func (handler *SchemaHandler) DeleteColumn(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO: validate whether table exists and specified columns does not belong to primary key or time column
// might be better for metaStore to do this and return specified error type

// TODO: need mapping from metaStore error to api error
// for metaStore error might also be user error
