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
	"github.com/spf13/cobra"
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/cluster/kvstore"
	"github.com/uber/aresdb/cmd/aresd/cmd"
	"github.com/uber/aresdb/controller/handlers"
	"github.com/uber/aresdb/utils"
	cfgfx "go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// default config file
	cfgFile = "config/ares-controller.yaml"
	// default port
	port = 6708

	Module = fx.Provide(Init)
)

type Params struct {
	fx.In

	LifeCycle fx.Lifecycle
}

// Result contains all
type Result struct {
	fx.Out

	ConfigProvider         cfgfx.Provider
	ZapLogger              *zap.SugaredLogger
	Scope                  tally.Scope
	EtcdClient             *kvstore.EtcdClient
	MetricsLoggingProvider utils.MetricsLoggingMiddleWareProvider
}

// AddFlags adds flags to command
func AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func Execute(setters ...cmd.Option) { _ = "STUB: not implemented"; return }

func Init() Result { _ = "STUB: not implemented"; return *new(Result) }

func runServer(logger *zap.SugaredLogger, handlerParams handlers.ServerParams) {
	_ = "STUB: not implemented"
	return
}
