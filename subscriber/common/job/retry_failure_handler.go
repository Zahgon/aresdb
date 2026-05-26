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

package job

import (
	"time"

	"github.com/uber/aresdb/controller/models"

	"github.com/uber-go/tally"
	"github.com/uber/aresdb/client"
	"github.com/uber/aresdb/subscriber/common/sink"
	"github.com/uber/aresdb/subscriber/config"
)

const defaultInitInterval = 5 * time.Second
const defaultMultiplier = float32(1.5)
const defaultMaxElapsedTime = 10 * time.Minute

// RetryFailureHandler implements
// exponential backoff retry
type RetryFailureHandler struct {
	serviceConfig  config.ServiceConfig
	scope          tally.Scope
	sink           sink.Sink
	jobName        string
	maxElapsedTime time.Duration
	elapsedTime    time.Duration
	multiplier     float32
	interval       time.Duration
}

// NewRetryFailureHandler creates a new RetryFailureHandler
func NewRetryFailureHandler(
	config models.FailureHandlerConfig,
	serviceConfig config.ServiceConfig,
	db sink.Sink,
	jobName string) *RetryFailureHandler {
	_ = "STUB: not implemented"
	return nil
}

// only support constant or increasing interval

// HandleFailure handles failure with retry
func (handler *RetryFailureHandler) HandleFailure(destination sink.Destination, rows []client.Row) (err error) {
	_ = "STUB: not implemented"
	return nil
}
