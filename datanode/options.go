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

package datanode

import (
	"net/http"

	"github.com/uber/aresdb/common"
	"github.com/uber/aresdb/datanode/bootstrap"
	"github.com/uber/aresdb/utils"
)

// options is the implementation of the interface Options
type options struct {
	instrumentOpts utils.Options
	bootstrapOpts  bootstrap.Options
	httpWrappers   utils.HTTPHandlerWrapper
	middleware     func(http.Handler) http.Handler
	cfg            common.AresServerConfig
}

// NewOptions creates a new set of storage options with defaults
func NewOptions() Options { _ = "STUB: not implemented"; return *new(Options) }

func (o *options) SetInstrumentOptions(value utils.Options) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

func (o *options) InstrumentOptions() utils.Options {
	_ = "STUB: not implemented"
	return *new(utils.Options)
}

func (o *options) BootstrapOptions() bootstrap.Options {
	_ = "STUB: not implemented"
	return *new(bootstrap.Options)
}

func (o *options) SetBootstrapOptions(bootstrapOptions bootstrap.Options) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

func (o *options) SetServerConfig(cfg common.AresServerConfig) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

func (o *options) ServerConfig() common.AresServerConfig {
	_ = "STUB: not implemented"

	// HttpWrappers return HttpWrapper
	return *new(common.AresServerConfig)
}

func (o *options) HTTPWrapper() utils.HTTPHandlerWrapper {
	_ = "STUB: not implemented"
	return *

	// SetHTTPWrapper return HttpWrapper
	new(utils.HTTPHandlerWrapper)
}

func (o *options) SetHTTPWrapper(wrappers utils.HTTPHandlerWrapper) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

// Middleware return middleware
func (o *options) Middleware() func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil

	// SetMiddleware set middleware
}

func (o *options) SetMiddleware(middleware func(http.Handler) http.Handler) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}
