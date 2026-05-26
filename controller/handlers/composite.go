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

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServerParams defines params needed to init server
type ServerParams struct {
	fx.In

	Config            config.Provider
	Logger            *zap.SugaredLogger
	HealthHandler     HealthHandler
	ConfigHandler     ConfigHandler
	SchemaHandler     SchemaHandler
	NamespaceHandler  NamespaceHandler
	MembershipHandler MembershipHandler
	AssignmentHandler AssignmentHandler
	PlacementHandler  PlacementHandler
	UIHandler         UIHandler
	WrapperProvider   utils.MetricsLoggingMiddleWareProvider
}

// NewCompositeHandler is the provider for http server
func NewCompositeHandler(p ServerParams) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
