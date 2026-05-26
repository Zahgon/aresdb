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

// NamespaceHandlerParams defines parameters needed to initialize namespace handler
type NamespaceHandlerParams struct {
	fx.In

	NamespaceMutator mutatorCom.NamespaceMutator
	Logger           *zap.SugaredLogger
}

// NamespaceHandler serves namespace requests
type NamespaceHandler struct {
	namespaceMutator mutatorCom.NamespaceMutator
	logger           *zap.SugaredLogger
}

// NewNamespaceHandler creates a new namespace handler
func NewNamespaceHandler(p NamespaceHandlerParams) NamespaceHandler {
	_ = "STUB: not implemented"
	return *new(NamespaceHandler)
}

// Register adds paths to router
func (h NamespaceHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

// CreateNamespace swagger:route POST /namespaces createNamespace
// adds a new namespace
//
// Consumes:
//   - application/json
func (h NamespaceHandler) CreateNamespace(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// ListNamespaces swagger:route GET /namespaces listNamespaces
// returns all namespaces
//
// Produces:
//   - application/json
func (h NamespaceHandler) ListNamespaces(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
