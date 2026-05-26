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
	"github.com/uber-go/tally"
	mutatorCom "github.com/uber/aresdb/controller/mutators/common"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// MembershipHandlerParams defineds parameters needed to initialize schema handler
type MembershipHandlerParams struct {
	fx.In

	MembershipMutator mutatorCom.MembershipMutator
	Logger            *zap.SugaredLogger
	Scope             tally.Scope
}

// MembershipHandler serves schema requests
type MembershipHandler struct {
	membershipMutator mutatorCom.MembershipMutator
	logger            *zap.SugaredLogger
	scope             tally.Scope
}

// NewMembershipHandler creates a new schema handler
func NewMembershipHandler(p MembershipHandlerParams) MembershipHandler {
	_ = "STUB: not implemented"
	return *new(MembershipHandler)
}

// Register adds paths to router
func (h MembershipHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

// Join adds a instance
func (h MembershipHandler) Join(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetInstance swagger:route GET /membership/{namespace}/instances/{instance} getInstance
// returns an instance
func (h MembershipHandler) GetInstance(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetInstances swagger:route GET /membership/{namespace}/instances getInstances
// returns all instances
func (h MembershipHandler) GetInstances(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Leave deletes an existing instance
func (h MembershipHandler) Leave(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetHash swagger:route GET /membership/{namespace}/hash getMembershipHash
// returns hash of all instances in a namespace
func (h MembershipHandler) GetHash(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
