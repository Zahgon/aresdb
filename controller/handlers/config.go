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
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/cluster/kvstore"
	mutatorCom "github.com/uber/aresdb/controller/mutators/common"
	"github.com/uber/aresdb/utils"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigHandlerParams defines params needed to initialize ConfigHandler
type ConfigHandlerParams struct {
	fx.In

	Logger        *zap.SugaredLogger
	Scope         tally.Scope
	JobMutator    mutatorCom.JobMutator
	SchemaMutator mutatorCom.TableSchemaMutator
	EtcdClient    *kvstore.EtcdClient
}

// ConfigHandler serves requests for job configurations
type ConfigHandler struct {
	logger *zap.SugaredLogger
	scope  tally.Scope

	jobMutator    mutatorCom.JobMutator
	schemaMutator mutatorCom.TableSchemaMutator

	etcdClient *kvstore.EtcdClient
}

// NewConfigHandler creates a new ConfigHandler
func NewConfigHandler(p ConfigHandlerParams) ConfigHandler {
	_ = "STUB: not implemented"
	return *new(ConfigHandler)
}

// Register adds paths to router
func (h ConfigHandler) Register(router *mux.Router, wrappers ...utils.HTTPHandlerWrapper) {
	_ = "STUB: not implemented"
	return
}

func (h ConfigHandler) getNumShards(namespace string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetJob swagger:route GET /config/{namespace}/jobs/{job} getJob
// gets job config by name
func (h ConfigHandler) GetJob(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetJobs swagger:route GET /config/{namespace}/jobs getJobs
// returns all jobs config
func (h ConfigHandler) GetJobs(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// only return system error when error is not NonExist

// DeleteJob swagger:route DELETE /config/{namespace}/jobs/{job} deleteJob
// deletes a job
func (h ConfigHandler) DeleteJob(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// UpdateJob swagger:route PUT /config/{namespace}/jobs/{job} updateJob
// updates job config
//
// Consumes:
//   - application/json
func (h ConfigHandler) UpdateJob(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// AddJob swagger:route POST /config/{namespace}/jobs addJob
// adds a new job
//
// Consumes:
//   - application/json
func (h ConfigHandler) AddJob(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetHash swagger:route GET /config/{namespace}/hash getJobsHash
// returns hash that will be different if any job changed
func (h ConfigHandler) GetHash(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
