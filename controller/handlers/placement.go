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

	"github.com/gorilla/mux"
	"github.com/m3db/m3/src/cluster/generated/proto/placementpb"
	"github.com/m3db/m3/src/cluster/placement"
	mutatorCom "github.com/uber/aresdb/controller/mutators/common"
	"github.com/uber/aresdb/utils"
)

// PlacementHandler handles placement requests
type PlacementHandler struct {
	placementMutator mutatorCom.PlacementMutator
}

// NewPlacementHandler creates placement handler
func NewPlacementHandler(mutator mutatorCom.PlacementMutator) PlacementHandler {
	_ = "STUB: not implemented"
	return *new(PlacementHandler)
}

// Register adds paths to router
func (h PlacementHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

func newInstancesFromProto(instancepbs []placementpb.Instance) ([]placement.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Init initialize new placement
func (h *PlacementHandler) Init(rw *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Get get the current placement
func (h *PlacementHandler) Get(rw *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Add adds new instances
func (h *PlacementHandler) Add(rw *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// validate all shards available in placement before adding instance

// Replace replace existing instances within placement with new instances
func (h *PlacementHandler) Replace(rw *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// validate all shards are available before replace instance

// Remove remove instance from placement
func (h *PlacementHandler) Remove(rw *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// validate all shards are available before replace instance

// MarkNamespaceAvailable marks all instance/shards in placement as available
func (h *PlacementHandler) MarkNamespaceAvailable(rw *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// MarkInstanceAvailable marks one instance as available
func (h *PlacementHandler) MarkInstanceAvailable(rw *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func respondWithPlacement(p placement.Placement, rw *utils.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}
