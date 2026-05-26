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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uber/aresdb/broker/config"
	"github.com/uber/aresdb/cmd/aresd/cmd"
	"github.com/uber/aresdb/common"
	"github.com/uber/aresdb/utils"
)

func Execute(setters ...cmd.Option) { _ = "STUB: not implemented"; return }

func start(
	cfg config.BrokerConfig,
	logger common.Logger,
	queryLogger common.Logger,
	metricsCfg common.Metrics,
	httpWrapper utils.HTTPHandlerWrapper,
	middleware func(http.Handler) http.Handler,
) {
	_ = "STUB: not implemented"
	return
}

// Init common components.

// fetch and keep syncing schema

// executor

// init handlers

// start HTTP server

// Support CORS calls.

// AddFlags adds flags to command
func AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// ReadConfig populates BrokerConfig
func ReadConfig(defaultCfg map[string]interface{}, flags *pflag.FlagSet) (cfg config.BrokerConfig, err error) {
	_ = "STUB: not implemented"
	return *new(config.BrokerConfig), nil
}

// bind command flags

// set defaults

// merge in config file
