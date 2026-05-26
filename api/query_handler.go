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

	"github.com/m3db/m3/src/x/sync"
	"github.com/uber/aresdb/cluster/topology"

	"github.com/uber/aresdb/memstore"
	"github.com/uber/aresdb/query"
	queryCom "github.com/uber/aresdb/query/common"
	"github.com/uber/aresdb/utils"

	"github.com/gorilla/mux"
	apiCom "github.com/uber/aresdb/api/common"
	"github.com/uber/aresdb/common"
)

// QueryHandler handles query execution.
type QueryHandler struct {
	shardOwner    topology.ShardOwner
	memStore      memstore.MemStore
	deviceManager *query.DeviceManager
	workerPool    sync.WorkerPool
}

// NewQueryHandler creates a new QueryHandler.
func NewQueryHandler(
	memStore memstore.MemStore,
	shardOwner topology.ShardOwner,
	cfg common.QueryConfig,
	maxConcurrentQueries int,
) *QueryHandler {
	_ = "STUB: not implemented"
	return nil
}

// GetDeviceManager returns the device manager of query handler.
func (handler *QueryHandler) GetDeviceManager() *query.DeviceManager {
	_ = "STUB: not implemented"
	return nil
}

// Register registers http handlers.
func (handler *QueryHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

// HandleAQL swagger:route POST /query/aql queryAQL
// query in AQL
//
// Consumes:
//   - application/json
//   - application/hll
//
// Produces:
//   - application/json
//
// Responses:
//
//	default: errorResponse
//	    200: aqlResponse
//	    400: aqlResponse
func (handler *QueryHandler) HandleAQL(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// default device to negative value to differentiate 0 from empty
	return
}

func (handler *QueryHandler) handleAQLInternal(aqlRequest apiCom.AQLRequest, w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Override from query parameter

// for logging purpose only

func handleQuery(memStore memstore.MemStore, shardOwner topology.ShardOwner, deviceManager *query.DeviceManager, aqlRequest apiCom.AQLRequest, aqlQuery queryCom.AQLQuery) (qc *query.AQLQueryContext, statusCode int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// Compilation error, should be bad request

// Find a device that meets the resource requirement of this query
// Use query specified device as hint

// Unable to find a device for the query.

// Unable to fulfill this request due to resource not available, clients need to try sometimes later.

// Execute.

// Report

func getReponseWriter(returnHLL bool, nQueries int) QueryResponseWriter {
	_ = "STUB: not implemented"
	return *new(QueryResponseWriter)
}

// QueryResponseWriter defines the interface to write query result and error to final response.
type QueryResponseWriter interface {
	ReportError(queryIndex int, table string, err error, statusCode int)
	ReportQueryContext(*query.AQLQueryContext)
	ReportResult(int, *query.AQLQueryContext)
	// param compress means whether client accepts compressed results
	Respond(w *utils.ResponseWriter)
	GetStatusCode() int
}

// JSONQueryResponseWriter writes query result as json.
type JSONQueryResponseWriter struct {
	response   queryCom.AQLResponse
	statusCode int
}

// NewJSONQueryResponseWriter creates a new JSONQueryResponseWriter.
func NewJSONQueryResponseWriter(nQueries int) QueryResponseWriter {
	_ = "STUB: not implemented"
	return *new(QueryResponseWriter)
}

// ReportError writes the error of the query to the response.
func (w *JSONQueryResponseWriter) ReportError(queryIndex int, table string, err error, statusCode int) {
	_ = "STUB: not implemented"
	// Usually larger status code means more severe problem.
	return
}

// ReportQueryContext writes the query context to the response.
func (w *JSONQueryResponseWriter) ReportQueryContext(qc *query.AQLQueryContext) {
	_ = "STUB: not implemented"
	return
}

// ReportResult writes the query result to the response.
func (w *JSONQueryResponseWriter) ReportResult(queryIndex int, qc *query.AQLQueryContext) {
	_ = "STUB: not implemented"
	return
}

// Respond writes the final response into ResponseWriter.
func (w *JSONQueryResponseWriter) Respond(rw *utils.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

// GetStatusCode returns the status code written into response.
func (w *JSONQueryResponseWriter) GetStatusCode() int { _ = "STUB: not implemented"; return 0 }

// HLLQueryResponseWriter writes query result as application/hll. For more inforamtion, please refer to
// https://github.com/uber/aresdb/wiki/HyperLogLog.
type HLLQueryResponseWriter struct {
	response   *queryCom.HLLQueryResults
	statusCode int
}

// NewHLLQueryResponseWriter creates a new HLLQueryResponseWriter.
func NewHLLQueryResponseWriter() QueryResponseWriter {
	_ = "STUB: not implemented"
	return *new(QueryResponseWriter)
}

// ReportError writes the error of the query to the response.
func (w *HLLQueryResponseWriter) ReportError(queryIndex int, table string, err error, statusCode int) {
	_ = "STUB: not implemented"
	return
}

// ReportQueryContext writes the query context to the response. Since the format of application/hll is not
// designed for human reading, we will ignore storing query context in response for now.
func (w *HLLQueryResponseWriter) ReportQueryContext(qc *query.AQLQueryContext) {
	_ = "STUB: not implemented"

	// ReportResult writes the query result to the response.
	return
}

func (w *HLLQueryResponseWriter) ReportResult(queryIndex int, qc *query.AQLQueryContext) {
	_ = "STUB: not implemented"
	return
}

// Respond writes the final response into ResponseWriter.
func (w *HLLQueryResponseWriter) Respond(rw *utils.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

// GetStatusCode returns the status code written into response.
func (w *HLLQueryResponseWriter) GetStatusCode() int { _ = "STUB: not implemented"; return 0 }

// for now we only eager flush when
//  1. there's only 1 query in the request
//  2. the query is non aggregate query
func canEagerFlush(queries []queryCom.AQLQuery) bool { _ = "STUB: not implemented"; return false }
