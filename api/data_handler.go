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

	"github.com/uber/aresdb/memstore"
	"github.com/uber/aresdb/utils"

	"github.com/gorilla/mux"
)

// DataHandler handles data ingestion requests from the ingestion pipeline.
type DataHandler struct {
	memStore   memstore.MemStore
	workerPool sync.WorkerPool
}

// NewDataHandler creates a new DataHandler.
func NewDataHandler(memStore memstore.MemStore, maxConcurrentRequests int) *DataHandler {
	_ = "STUB: not implemented"
	return nil
}

// Register registers http handlers.
func (handler *DataHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

// PostData swagger:route POST /data/{table}/{shard} postData
// Post new data batch to a existing table shard
// Consumes:
//   - application/upsert-data
//
// Responses:
//
//	default: errorResponse
//	    200: noContentResponse
func (handler *DataHandler) PostData(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
