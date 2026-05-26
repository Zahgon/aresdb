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

package cmd

import (
	"net/http"

	"github.com/uber-go/tally"
	"github.com/uber/aresdb/common"
	"github.com/uber/aresdb/utils"
)

// Options represents options for executing command
type Options struct {
	DefaultCfg   map[string]interface{}
	ServerLogger common.Logger
	QueryLogger  common.Logger
	Metrics      common.Metrics
	HTTPWrapper  utils.HTTPHandlerWrapper
	Middleware   func(http.Handler) http.Handler
}

// Option is for setting option
type Option func(*Options)

// AresD is a wrapper of original functions for code reuse
type AresD struct {
	// server configuration and options
	cfg     common.AresServerConfig
	options *Options
	// server is http server instance started inside AresD
	server *http.Server
	// StartedChan is used to notify other route that the server is started
	StartedChan chan struct{}
}

// NewAresD create singleton of AresD
func NewAresD(cfg common.AresServerConfig, options *Options) *AresD {
	_ = "STUB: not implemented"
	return nil
}

// Start start aresd server
func (aresd *AresD) Start() { _ = "STUB: not implemented"; return }

// Shutdown stop aresd server
func (aresd *AresD) Shutdown() { _ = "STUB: not implemented"; return }

// Execute executes command with options
func Execute(setters ...Option) { _ = "STUB: not implemented"; return }

// start is the entry point of starting ares.
func (aresd *AresD) start(
	cfg common.AresServerConfig,
	logger common.Logger,
	queryLogger common.Logger,
	metricsCfg common.Metrics,
	httpWrapper utils.HTTPHandlerWrapper,
	middleware func(http.Handler) http.Handler,
) {
	_ = "STUB: not implemented"
	return
}

// Check whether we have a correct device running environment

// Pause profiler util requested

// Init common components.

// TODO keep this path for non-distributed mode, and to aovid code break
// should be removed later after distributed mode is mature

// Create MetaStore.

// Create DiskStore.

// fetch schema from controller and start periodical job

// immediate initial fetch

// Create MemStore.

// Read schema.

// create schema handler

// create enum handler

// create query hanlder.
// static shard owner with non distributed version

// create health check handler.

// Start HTTP server for debugging.

// Init shards.

// Start serving.

// Support CORS calls.

// notify other routes that the server is up

// waiting for the server to stop

// start datanode in distributed mode
func startDataNode(
	cfg common.AresServerConfig,
	logger common.Logger,
	scope tally.Scope,
	httpWrapper utils.HTTPHandlerWrapper,
	middleware func(http.Handler) http.Handler,
) {
	_ = "STUB: not implemented"
	return
}

// preparing

// start serving traffic
