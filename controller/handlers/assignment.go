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
	"github.com/uber/aresdb/controller/models"
	mutatorCom "github.com/uber/aresdb/controller/mutators/common"
	"github.com/uber/aresdb/utils"
	"go.uber.org/zap"
)

// AssignmentHandler serves requests for ingestion job assignments
type AssignmentHandler struct {
	logger *zap.SugaredLogger

	assignmentMutator mutatorCom.IngestionAssignmentMutator
	schemaMutator     mutatorCom.TableSchemaMutator
	membershipMutator mutatorCom.MembershipMutator
}

// NewAssignmentHandler creates a new AssignmentHandler
func NewAssignmentHandler(logger *zap.SugaredLogger, assignmentMutator mutatorCom.IngestionAssignmentMutator, schemaMutator mutatorCom.TableSchemaMutator, membershipMutator mutatorCom.MembershipMutator) AssignmentHandler {
	_ = "STUB: not implemented"
	return *new(AssignmentHandler)
}

// Register adds paths to router
func (h AssignmentHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

// GetAssignment swagger:route GET /assignment/{namespace}/assignments/{subscriber} getAssignment
// gets assignment by subscriber name
func (h AssignmentHandler) GetAssignment(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetAssignments swagger:route GET /assignment/{namespace}/assignments getAssignments
// returns all assignments
func (h AssignmentHandler) GetAssignments(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetHash swagger:route GET /assignment/{namespace}/hash/{subscriber} getAssignmentHash
// returns hash that will be different if any thing changed for given assignment
func (h AssignmentHandler) GetHash(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func composeSingleAssignment(instances []models.Instance, assignment models.IngestionAssignment) models.IngestionAssignment {
	_ = "STUB: not implemented"
	return *new(models.IngestionAssignment)
}
